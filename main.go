package main

import (
	"fmt"
)

type Human struct {
	Name string
	Age  int
	Prof string
}

type Action struct {
	Human
	Skill string
}

func (human Human) Greeting() {
	fmt.Printf("Hey there! I'am %s %d years old. I do %s.\n", human.Name, human.Age, human.Prof)
}

func (action Action) ShowSkill() {
	fmt.Printf("My super power is %s.\n", action.Skill)
}

func main() {
	toPrint := Action{Human: Human{Name: "Rustam", Age: 8, Prof: "operation systems"}, Skill: "Math"}

	toPrint.Greeting()
	toPrint.ShowSkill()
}
