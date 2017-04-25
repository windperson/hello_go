package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type person struct {
	Name string
	Age  uint
}

func nilValueTemplate() (string, error) {
	t := template.New("a test")
	t, err := t.Parse("A Template without any binding")
	if err != nil {
		panic(err)
	}
	var buffer bytes.Buffer

	t.Execute(&buffer, nil)
	return buffer.String(), err
}

func main() {
	t := template.New("a test")
	t, err := t.Parse("Hello Template {{.Name}} who's age is {{.Age}}.")
	if err != nil {
		panic(err)
	}
	data := person{Name: "Charles", Age: 18}
	var buffer bytes.Buffer

	t.Execute(&buffer, data)
	fmt.Println("result=", buffer.String())
	nonBinding, err := nilValueTemplate()
	if err != nil {
		panic(err)
	}
	fmt.Println("nil binding=", nonBinding)
}
