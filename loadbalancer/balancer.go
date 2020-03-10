package main

import "container/heap"

type Balancer struct {
	pool Pool
	done chan *Worker
}

func NewBalancer() *Balancer {
	done := make(chan *Worker, nWorkers)
	b := &Balancer{
		pool: make(Pool, 0, nWorkers),
		done: done,
	}

	for i := 0; i <nWorkers; i++ {
		w := &Worker{requests: make(chan Request, nRequests)}
		heap.Push(&b.pool, w)
		go w.work(b.done)
	}

	return b
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <- work:
			b.dispatch(req)
		case w := <- b.done:
			b.completed(w)
		}
	}
}

func (b *Balancer) dispatch(request Request) {
	w := heap.Pop(&b.pool).(*Worker)
	w.requests <- request
	w.pending++
	heap.Push(&b.pool, w)
}

func (b *Balancer) completed(w *Worker) {
	w.pending--
	heap.Remove(&b.pool, w.i)
	heap.Push(&b.pool, w)
}
