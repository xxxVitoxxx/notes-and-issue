package main

import "fmt"

// interface polymorphism
// 多型

type Animal interface {
	Run()
	Eat()
}

type Dog struct{}

func (d *Dog) Run() {
	fmt.Println("dog is running")
}

func (d *Dog) Eat() {
	fmt.Println("dog is eating")
}

type Cat struct{}

func (c *Cat) Run() {
	fmt.Println("cat is running")
}

func (c *Cat) Eat() {
	fmt.Println("cat is eating")
}

// Dog 跟 Cat 都實作了 Animal 的方法
// Dog 跟 Cat 擁有同樣的方法，但結果卻不一樣，這就是多型，相同方法卻有不一樣的結果

// 以下是可以將實作 Animal 介面的 object 當成變數傳入函式
func animal(a Animal) {
	a.Run()
	a.Eat()
}

func main() {
	d := &Dog{}
	c := &Cat{}
	animal(d)
	animal(c)
}
