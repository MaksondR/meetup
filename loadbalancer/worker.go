package main

type Worker struct {
	i int
	requests chan Request
	pending int
}

func (w *Worker) work(done chan *Worker) {
	for {
		req := <- w.requests
		req.c <- req.fn()
		done <- w
	}
}
