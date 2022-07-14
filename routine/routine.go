package main

func main() {
	runThing(make(chan int), make(chan int))
}

func process(val int) int {
	return 3
}

func runThing(in <-chan int, out chan<- int) {
	go func() {
		for val := range in {
			result := process(val)
			out <- result
		}
	}()
}
