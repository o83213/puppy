package puppy

import (
	"fmt"

	"github.com/o83213/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From11(){
	fmt.Println("I am from version 1.1.0")
}

func From12(){
	fmt.Println("I am from version 1.2.0")
}