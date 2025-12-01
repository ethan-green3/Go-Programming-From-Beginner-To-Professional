package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func checkType(v any) string {
	switch v.(type) {
	case int, int32, int64:
		return "int"
	case float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

func main() {
	data := []any{28, 2.87, "hello", true, person{"Ethan", "Green"}}

	for _, d := range data {
		fmt.Printf("%v is %v\n", d, checkType(d))
	}
}
