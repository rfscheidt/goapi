package main

import "log"

type Student struct {
	name   string
	grades []int
	age    int
}

func (s Student) averageGrade() float32 {
	sum := 0
	for _, item := range s.grades {
		sum += item
	}

	return float32(sum) / float32(len(s.grades))
}

func averageGrade(grades []int) float32 {
	sum := 0
	for _, item := range grades {
		sum += item
	}

	return float32(sum) / float32(len(grades))
}

func funcMethod() {
	student1 := Student{"Rafael", []int{70, 80, 80}, 19}
	log.Println(student1.averageGrade())

	log.Println(averageGrade(student1.grades))
}
