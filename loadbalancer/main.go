package main

const nRequests  = 100
const nWorkers = 10

func main() {
	work := make(chan Request)

	for i := 0; i < nRequests; i++ {
		go requester(work)
	}

	NewBalancer().balance(work)
}
