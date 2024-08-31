package struct_learn

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Grade string
}

func (s Student) Introduce() {
	fmt.Printf("Hi, I'm %s, %d years old, and I'm in %s grade\n", s.Name, s.Age, s.Grade)
}

func (s *Student) Promote() {
	switch s.Grade {
	case "9th":
		s.Grade = "10th"
	case "10th":
		s.Grade = "11th"
	case "11th":
		s.Grade = "12th"
	case "12th":
		s.Grade = "Graduated"
	default:
		fmt.Println("Grade not recognized or already graduated.")
	}
}

type School struct {
	Name     string
	Students []*Student
}

func (s *School) AddStudent(student *Student) {
	s.Students = append(s.Students, student)
}

func (s School) ListStudents() {
	fmt.Printf("School: %s\n", s.Name)
	for _, student := range s.Students {
		student.Introduce()
	}
	/*for i := 0; i < len(s.Students); i++ {
		s.Students[i].Introduce()
	}
	*/
}
