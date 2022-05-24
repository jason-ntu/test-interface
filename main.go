package main

import (
	"fmt"

	. "github.com/jason-ntu/test-interface/animal"
)

func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
	for _, a := range animals {
		fmt.Println(a.Speak())
	}
}
