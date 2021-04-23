package main

import "tour/distributed/nodes"

var parrotChan = make(chan Chat, 100)
var juliaChan = make(chan Chat, 100)

func main() {
	//fmt.Println(Sqrt(100))
	//fmt.Println(Sqrt(23523))
	//fmt.Println(CubeRt(complex(2.0, 0), 1))

	//wc.Serve(WordCount)
	//pic.Serve(Pic)

	// lesson 2
	//http.Handle("/string", String("I'm a frayed knot."))
	//http.Handle("/struct", &Struct{"Hello",":","USENIX!"})
	//http.Handle("/hello", helloHandler{})
	//http.ListenAndServe(":4000", nil)

	//http.Handle("/", fractal.MainPage)
	//http.Handle("/mandelbrot", FractalHandler{parser: mandlebrotParser, colorer: fractal.Cycle})
	//http.Handle("/julia", FractalHandler{parser: juliaParser})
	//http.ListenAndServe(":4000", nil)
	//pic.ServeImage(Pic2)

	// lesson 3
	//fmt.Println(Same(tree.New(1), tree.New(1)))
	//fmt.Println(Same(tree.New(1), tree.New(2)))

	//type Log string
	//http.Handle("/", ajax.LogPage)
	//go func() {
	//	for i := 0; ; i++ {
	//		ajax.Chan <- Log(fmt.Sprintf("log - %d", i))
	//		time.Sleep(time.Second)
	//	}
	//}()

	//chatRoom := make(chan interface{})
	//http.Handle("/", chat.ChatPage)
	//http.Handle("/join", JoinHandler(chatRoom))
	//http.Handle("/say", SayHandler(chatRoom))
	//http.Handle("/exit", ExitHandler(chatRoom))
	//http.Handle("/julia", FractalHandler{parser: juliaParser})
	//
	//go managerRoutine(chatRoom)
	//go parrotHandler(chatRoom)
	//go juliaHandler(chatRoom)
	//http.ListenAndServe(":4000", nil)

	// lesson 4
	master := nodes.StartMaster("./gopher_complete.png", 20, 20)
	master.
}

