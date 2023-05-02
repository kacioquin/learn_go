package structs

type person struct {
	name string
	age  int
}

func NewPerson(name string, age int) person {

	var person = person{
		name: name,
		age:  age,
	}

	return person
}
