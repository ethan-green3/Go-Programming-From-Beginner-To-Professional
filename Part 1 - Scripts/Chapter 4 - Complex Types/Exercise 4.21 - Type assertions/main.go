package main

import (
	"errors"
	"fmt"
)

// Can use the any keyword to achieve the same thing as interface{}
func doubler(v interface{}) (string, error) {
	if i, ok := v.(int); ok {
		return fmt.Sprint(i * 2), nil
	}
	if s, ok := v.(string); ok {
		return s + s, nil
	}

	return "", errors.New("Unsupported type passed")
}

func main() {
	res, _ := doubler(5)
	fmt.Println("5 :", res)
	result, _ := doubler("yum")
	fmt.Println("yum: ", result)
	_, err := doubler(true)
	fmt.Println("true:", err)

}
