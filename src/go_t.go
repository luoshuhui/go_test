package main

import (
	"encoding/json"
	"fmt"
	"runtime"

	//	"io/ioutil"
)
type t_base struct{
	i int32
}

type t_son struct{
	//匿名组合,==继承
	t_base
	j int32
}
a
func  make_son(i int32,j int32)(ret * t_son){
	var b t_son
	b.i =i
	b.j =j
	return &b 
}
func (base t_base) base_func()(){
	fmt.Printf("base %d\r\n",base.i)
}


func (son t_son) son_func()(){
	fmt.Printf("son %d\r\n",son.j)
}
func inherit_test1(){
	s:=make_son(1,2)
	s.base_func()
	s.son_func()
}
//------------------------------------------------------------------------------
func run(i int, c chan int) {
	fmt.Printf("run %d\r\n", i)
	fmt.Printf("run %d\r\n", i)
	//进入syscall或者调用Gosched会导致调度。
	runtime.Gosched()
	fmt.Printf("run %d\r\n", i)
	c <- i
	
}

func runtime_test1() {
	fmt.Println("go test:", runtime.NumCPU())
}
//------------------------------------------------------------------------------
func chan_test1() {
	chans := make([]chan int, 10)
	//设置可以并行的goroutine数量，其实就是系统线程数
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i = i + 1 {
		chans[i] = make(chan int)
		go run(i, chans[i])
	}
	for _, ch := range chans {
		<-ch
	}

}
//------------------------------------------------------------------------------
type Book struct {
	Title string
	Auth  string
	Age   int
}

func encoding_json_test1() {
	var book Book = Book{"aaa", "sb", 1111}
	var d *Book = new(Book)
	d.Age = 100
	b, e := json.Marshal(book)
	if e == nil {
		for _, ch := range b {
			fmt.Printf("%c\r\n", ch)
		}

	} else {
		fmt.Println("err ", e.Error)
	}

	var book2 Book
	json.Unmarshal(b, &book2)
	fmt.Println("book2=", book2)
}
//------------------------------------------------------------------------------
func slice_test1() {
	var slice []int
	slice = make([]int, 5, 10)
	slice[1] = 111
	slice[2] = 222
	slice[3] = 333
	slice[4] = 444
	fmt.Printf("len %d,cap %d slice[1] %d\r\n", len(slice), cap(slice), slice[1])
}
//------------------------------------------------------------------------------
type Inter int

func (a Inter) Less(b Inter) bool {
	return a < b
}

func Select_test1() {
	ch := make(chan int, 1)
	for {
		select {
			case ch <- 0:
			case ch <- 1:
		}
		i := <-ch
		fmt.Println("Value received:", i)

	}
}

func main() {
	//runtime_test1()
	chan_test1()
	//encoding_json_test1()

	//slice_test1()
	//Select_test1()
	inherit_test1()

}
