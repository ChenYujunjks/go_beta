package struct_learn

func Struct_jjj() {
	school := School{Name: "Greenwood High"}
	student1 := &Student{Name: "John", Age: 16, Grade: "10th"}
	student2 := &Student{Name: "Jane", Age: 15, Grade: "9th"}
	student3 := &Student{Name: "Mike", Age: 17, Grade: "11th"}

	p_school := &school
	school.AddStudent(student1)
	p_school.AddStudent(student2)
	school.AddStudent(student3)
	student3.Promote()
	// 列出所有学生
	school.ListStudents()
	// School: Greenwood High
	// Name: John, Age: 16, Grade: 10th
	// Name: Jane, Age: 15, Grade: 9th
	// Name: Mike, Age: 17, Grade: Gra
}
