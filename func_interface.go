package main

import "fmt"

/*
step 1: define the interface with a single method
step 2: define a function named {interface_name}Func type which takes the same inputs as the interface method
step 3: make the function implements the interface
step 4: create the implementations of the {interface_name}Func
*/

//step 1
type Animale interface {
	Action()
}

//step 2
type AnimaleFunc func()

//step 3
func (a AnimaleFunc) Action() {
	a()
}

//step 4
func run() {
	fmt.Println("run")
}

func swim() {
	fmt.Println("swim")
}

func fly() {
	fmt.Println("fly")
}

func main() {
	AnimaleFunc(run).Action()
	AnimaleFunc(swim).Action()
	AnimaleFunc(fly).Action()
}
