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

func main() {
	t := template.New("a test")
	t, err := t.Parse("Hello Template {{.Name}} who's age is {{.Age}}.")
	if err != nil {
		panic(err)
	}
	data := person{Name: "Charles", Age: 18}
	var buffer bytes.Buffer

	t.Execute(&buffer, data)
	fmt.Print("result=", buffer.String())
}
