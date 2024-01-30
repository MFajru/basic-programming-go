package solution

import "fmt"

type student struct {
	id    string
	name  string
	age   int
	grade float32
}

func addStudent(id string, name string, age int, grade float32) student {
	newStudent := student{
		id:    id,
		name:  name,
		age:   age,
		grade: grade,
	}

	return newStudent
}

func Pretest() {
	var id, name string
	var age int
	var grade float32
	var gradeTotal float32

	students := []student{
		{id: "132131", name: "John", age: 20, grade: 100},
		{id: "121321", name: "Bro", age: 21, grade: 20},
	}

	fmt.Print("ID: ")
	fmt.Scanln(&id)
	fmt.Print("Name: ")
	fmt.Scanln(&name)
	fmt.Print("Age: ")
	fmt.Scanln(&age)
	fmt.Print("Grade: ")
	fmt.Scanln(&grade)
	fmt.Print("\n")

	students = append(students, addStudent(id, name, age, grade))

	for i, student := range students {
		fmt.Printf("Student %d\n", i+1)
		fmt.Printf("Student ID: %s\n", student.id)
		fmt.Printf("Student name: %s\n", student.name)
		fmt.Printf("Student age: %d\n", student.age)
		fmt.Printf("Student name: %1f\n\n", student.grade)
		gradeTotal += student.grade
	}

	avg := gradeTotal / float32(len(students))
	fmt.Printf("Average student grade: %2f\n", avg)

}
