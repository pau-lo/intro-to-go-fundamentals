package main

import (
	"fmt"
	"reflect"
)

type Animal struct {

	// setting the required tag with a max length of chars
	Name   string `required max: "100"`
	Origin string
}

func main() {

	// getting the type of the object using the reflect package
	t := reflect.TypeOf(Animal{})
	// getting the field from that type and passing it the Name
	field, _ := t.FieldByName("Name")
	// get the field tag
	fmt.Println(field.Tag)
}
