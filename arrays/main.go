package main

import "fmt"

func main() {
	// grades := [...]int{97, 85, 93}
	// fmt.Printf("Grades: %v", grades)

	// var students [3]string
	// fmt.Printf("Students: %v\n", students)
	// students[0] = "Taufik"
	// students[1] = "Fauzi"
	// students[2] = "Doni"
	// fmt.Printf("Student #1: %v\n", students[0])
	// fmt.Printf("Number of Students: %v", len(students))

	// var identifyMatrix [3][3]int = [3][3]int{
	// 	{1, 0, 0},
	// 	{0, 1, 0},
	// 	{0, 0, 1},
	// }
	// var identifyMatrix [3][3]int
	// identifyMatrix[0] = [3]int{1, 0, 0}
	// identifyMatrix[1] = [3]int{0, 1, 0}
	// identifyMatrix[2] = [3]int{1, 0, 1}
	// fmt.Println(identifyMatrix)

	// a := [...]int{1, 2, 3}
	// b := &a
	// b[1] = 5
	// fmt.Println(a)
	// fmt.Println(b)

	// ===== slices =====
	// a := []int{1, 2, 3}
	// b := a
	// b[1] = 5
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	// a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// b := a[:]
	// c := a[3:]
	// d := a[:6]
	// e := a[3:6]
	// a[5] = 42
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	// a := make([]int, 3, 100)
	// a := []int{}
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, 1)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a ))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, []int{2, 3, 4, 5}...)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := append(a[:2], a[3:]...)
	fmt.Println(b)
	fmt.Println(a)
}
