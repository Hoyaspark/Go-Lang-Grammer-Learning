package main

func main() {
	//for i := range countTo(10) {
	//	fmt.Println(i)
	//}

}

func countTo(max int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < max; i++ {
			ch <- i
		}
	}()

	return ch
}
