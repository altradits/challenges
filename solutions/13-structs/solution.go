package piscine

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

func (p Person) Greet() string {
	return "Hi, I'm " + p.Name
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}
