package interface-test

import (
	"fmt"
)

type Flyer interface {
	fly() string
}

type Walker interface {
	walk() string
}

type Bird struct {
	Name string
}

func (b *Bird) getName() string {
	return b.Name
}

func (b *Bird) fly() string {
	return b.getName() + " is Flying..."
}

func (b *Bird) walk() string {
	return b.getName() + " is Walking..."
}

func isFlyer(f Flyer) {
	fmt.Println("I am ... ", f.fly())
}

func isWalker(w Walker) {
	fmt.Println("I am ... ", w.walk())
}

func main() {
	b := Bird{"Chirper"}

	fmt.Println(b.fly())  // Flying...
	fmt.Println(b.walk()) // Walking...
	isFlyer(b)
	isWalker(&b)

}
