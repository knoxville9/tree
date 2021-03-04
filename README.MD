#select
第一种:
````
func main() {
	w1 := make(chan bool)
	w2 := make(chan bool)
	one := bb("1", w1)
	two := bb("2", w2)
	for {
		select {
		case w := <-one:
			fmt.Println(w)
		case l := <-two:
			fmt.Println(l)
		case <-time.After(time.Millisecond*500):
			fmt.Println("end")
		}
	}

}
func bb(msg string, wait chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- msg
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return c
}
````

***
第二种:
```
func main() {
	w1 := make(chan bool)
	w2 := make(chan bool)
	one := bb("1", w1)
	two := bb("2", w2)
	after := time.After(time.Second)
	for {
		select {
		case w := <-one:
			fmt.Println(w)
		case l := <-two:
			fmt.Println(l)
		case <-after:
			fmt.Println("end")
			return
		}
	}

}
func bb(msg string, wait chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- msg
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return c
}
```