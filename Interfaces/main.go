package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	side float64
}

type shape interface {
	area() float64
}

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	jarvis := englishBot{}
	javier := spanishBot{}
	printGreeting(jarvis)
	printGreeting(javier)

	t := triangle{3, 23.21}
	sq := square{2.34}

	printArea(t)
	printArea(sq)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func (t triangle) area() float64 {
	return t.base * t.height / 2
}

func (sq square) area() float64 {
	return sq.side * sq.side
}

func printArea(s shape) {
	fmt.Println("Area of this shape is:", s.area())
}
