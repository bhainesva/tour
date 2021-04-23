package main

import (
	"net/http"
	"strings"
	"tour/ajax"
	"tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch<-t.Value
	Walk(t.Right, ch)
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i:=0;i<10;i++ {
		x, y := <-ch1, <-ch2
		if x != y {
			return false
		}
	}

	return true
}

func getUsersList(users map[string]struct{}) []string {
	var ul []string
	for k, _ := range users {
		ul = append(ul, k)
	}
	return ul
}

func juliaHandler(ch chan interface{}) {
	for {
		msg := <-juliaChan
		fields := strings.Fields(msg.Msg)
		if fields[0] != "julia:" {
			continue
		}

		ch<-Chat{ Who: "julia", Msg: msg.Who + ": <img src=\"/julia?p=320+240+(-2%2b1i)+(4-4i)+" + strings.ReplaceAll(fields[1], "+", "%2b") + "+100\">"}
	}
}

func parrotHandler(ch chan interface{}) {
	knowledge := map[string]string{}

	for {
		msg := <-parrotChan
		fields := strings.Fields(msg.Msg)
		if fields[0] != "parrot:" {
			continue
		}

		if fields[2] == "=" && fields[3] != "" {
			knowledge[fields[1]] = fields[3]
			ch <- Chat{Who: "parrot", Msg: msg.Who + ": " + strings.Join(fields[1:], " ")}
		}
		if fields[1] == "tell" && fields[3] == "about" {
			if _, ok := knowledge[fields[4]]; !ok {
				ch <- Chat{Who: "parrot", Msg: "I don't know about " + fields[4] + "!"}
				continue
			}

			ch <- Chat{Who: "parrot", Msg: fields[2] + ": " + knowledge[fields[4]]}
		}
	}
}

func managerRoutine(ch chan interface{}) {
	type Status string
	users := map[string]struct{}{}

	for {
		switch m := (<-ch).(type) {
		case Join:
			ajax.Chan <- Status("[+" + strings.Join(append([]string{m.Who}, getUsersList(users)...), ", ") + "]")
			users[m.Who] = struct{}{}
		case Chat:
			ajax.Chan <- m
			parrotChan <- m
			juliaChan <- m
		case Exit:
			delete(users, m.Who)
			ajax.Chan <- Status("[-" + strings.Join(append([]string{m.Who}, getUsersList(users)...), ", ") + "]")
		}
	}
}

type Chat struct {
	Who string
	Msg string
}

type JoinHandler chan interface{}
func (j  JoinHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	j <- Join{r.FormValue("id")}
}

type SayHandler chan interface{}
func (j  SayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	j <- Chat{r.FormValue("id"), r.FormValue("msg")}
}

type ExitHandler chan interface{}
func (j  ExitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	j <- Exit{r.FormValue("id")}
}

type Join struct {
	Who string
}

type Exit struct {
	Who string
}

