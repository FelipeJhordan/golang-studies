package main

// Você tem um canal e você irá distribuir em outros canais
func main() {
	c := generate(5, 10)

	// fan out
	d1 := divide(c)
	d2 := divide(c)
}

func generate(numbers ...int) chan int {
	channel := make(chan int)

	go func() {
		for _, n  range numbers{
			channel <- n
		}
		close(channel)
	}()

	return channel
}


func divide(input chan int) chan int {
	channel := make(chan int)
	go func() {
		for number := range input {
			channel <- number /2 
		}
	}()

	return channel
}