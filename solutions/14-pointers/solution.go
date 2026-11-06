package piscine

type Person struct {
	Name string
	Age  int
}

func Double(n *int) {
	*n *= 2
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}
