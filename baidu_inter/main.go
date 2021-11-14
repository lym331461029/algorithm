package main

import "fmt"



func main() {
	c := make(chan int)
	go func() {

		for {
			char, ok := <-c
			if ok {
				fmt.Printf("%c", byte(char))
				c <- char
			} else {
				break
			}

		}
	}()

	diff := 'a' - '1'
	i := int('a')
	for ; i <= int('d'); i++ {
		fmt.Printf("%c", byte(i))
		c <- i - int(diff)

		<-c
	}
	fmt.Println()
	close(c)
}

