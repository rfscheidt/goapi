package main

import (
	"fmt"

	"github.com/rfscheidt/goapi/animals"
	"github.com/rfscheidt/goapi/circus"
)

func interfaces() {
	dog1 := animals.Dog{"Everest"}
	dog2 := animals.Dog{"Chase"}

	cat1 := animals.Cat{"Marrie"}
	cat2 := animals.Cat{"Cherrei"}

	//fmt.Println(dog1.Speaks())
	//fmt.Println(dog2.Speaks())

	//fmt.Println(cat1.Speaks())
	//fmt.Println(cat2.Speaks())

	fmt.Println(circus.Speaks(dog1))
	fmt.Println(circus.Speaks(dog2))

	fmt.Println(circus.Speaks(cat1))
	fmt.Println(circus.Speaks(cat2))
}
