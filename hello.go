/*package main

import "fmt"

func main() {
	fmt.Print("Hello World")
}*/
package main

import (
	"fmt"
	//"link"
)

type People struct {
	Name string
}

type Teacher struct {
	People
	Department string
}

type Speaker interface {
	SayHi()
}

func (p People) SayHi() {
	fmt.Printf("%s\n", p.Name)
}

func main() {
	people := People{"zhang shan"}
	teacher := Teacher{People{"zheng zhi"}, "CS"}
	var is Speaker
	is = people
	is.SayHi()
	is = teacher
	is.SayHi()
}
