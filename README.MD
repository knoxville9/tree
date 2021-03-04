#select
第一种:

500毫秒内没有信号,则打印end ,循环

```go
package main
import (

"fmt"
"time"
)
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
```

***
第二种:

一秒后 跳出循序
```go
package main
import (

"fmt"
"time"
)
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

10000个goroutines 相继传值给左边
```go
package main
import "fmt"
func main() {
	theleft := make(chan int)
	left := theleft
	right := theleft
	for i := 0; i < 10000; i++ {
		right = make(chan int)
		go f(left, right)
		left=right
	}
	go func(c chan int) { right <- 1 }(right)
	fmt.Println(<-theleft)

}

func f(left, right chan int) {
	left <- 1 + <-right
}

```

**

模拟搜索,有三个搜索服务器,如果有超时则返回,取最快的一个

```go
package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
)

type Result struct {
	video string
}

func fakeSeach() <-chan Result {
	c := make(chan Result)

	go func() {
		for {
			rand.Seed(time.Now().Unix())
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			result := Result{
				video: uuid.NewV4().String(),
			}
			c <- result
		}
	}()
	return c

}

func main() {
	fanin(fakeSeach(), fakeSeach(), fakeSeach())

}

func fanin(ch1, ch2, ch3 <-chan Result) {
	for {
		select {
		case a := <-ch1:
			fmt.Println(a)
		case a := <-ch2:
			fmt.Println(a)
		case a := <-ch3:
			fmt.Println(a)
		case <-time.After(time.Duration(rand.Intn(1e3)) * time.Millisecond):
			fmt.Println("try one more time")
			return
		}

	}

}



```