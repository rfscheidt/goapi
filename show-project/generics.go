package main

import "log"

func addInts(list []int) int {
	var sum int
	for _, item := range list {
		sum += item
	}

	return sum
}

func addFloats(list []float32) float32 {
	var sum float32
	for _, item := range list {
		sum += item
	}

	return sum
}

type NumberType interface {
	int | float32 | float64
}

func addList[T NumberType](list []T) T {
	var sum T
	for _, item := range list {
		sum += item
	}

	return sum
}

func generics() {
	log.Println("Tests generics")

	//log.Println(addInts([]int{1, 2, 3}))
	//log.Println(addFloats([]float32{1, 2, 3}))

	log.Println(addList([]int{1, 2, 3}))
	log.Println(addList([]float32{1, 2, 3}))
	log.Println(addList([]float64{1, 2, 3}))
}
