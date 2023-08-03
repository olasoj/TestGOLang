package main

import (
	"fmt"
	"sort"
)

type Ordered interface {
	~int | ~float64 | ~string
}

func addStudent[T any](students []T, student T) []T {
	return append(students, student)
}

type OrderedSlice[T Ordered] []T // T must implement < and >

func (s OrderedSlice[T]) Len() int {
	return len(s)
}
func (s OrderedSlice[T]) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s OrderedSlice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// end of group for SortType
// Finally, we define a function, PerformSort, that uses SortType as follows:

func PerformSort[T any](slice []T, compare func(T, T) bool) {
	sort.Sort(SortType[T]{slice, compare})
}

func main() {

	students := []string{}
	result := addStudent(students, "Michael")
	result = addStudent(result, "Jennifer")
	result = addStudent(result, "Elaine")

	sort.Sort(OrderedSlice[string](result))
	fmt.Println(result)

	students1 := []int{}
	result1 := addStudent[int](students1, 78)
	result1 = addStudent[int](result1, 64)
	result1 = addStudent[int](result1, 45)
	sort.Sort(OrderedSlice[int](result1))
	fmt.Println(result1)

	students2 := []Student{}
	result2 := addStudent[Student](students2, Student{"John", 213, 17.5})
	result2 = addStudent[Student](result2, Student{"James", 111, 18.75})
	result2 = addStudent(result2, Student{"Marsha", 110, 16.25})

	PerformSort[Student](result2, func(s1, s2 Student) bool {
		return s1.Age < s2.Age // comparing two Student values
	})

	fmt.Println(result2)
}
