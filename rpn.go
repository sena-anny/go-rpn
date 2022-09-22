package go_rpn

import "fmt"

func Rpn(input []interface{}) (int, error) {
	var stack []int
	for _, value := range input {
		switch value.(type) {
		case string:
			switch value {
			case "+":
				top, sec := stack[0], stack[1]
				stack = append([]int{top + sec}, stack[2:]...)
			case "-":
				top, sec := stack[0], stack[1]
				stack = append([]int{sec - top}, stack[2:]...)
			case "/":
				top, sec := stack[0], stack[1]
				stack = append([]int{sec / top}, stack[2:]...)
			case "*":
				top, sec := stack[0], stack[1]
				stack = append([]int{top * sec}, stack[2:]...)
			default:
				return 0, fmt.Errorf("%v is invalid string", value)
			}
		case int:
			stack = append(stack, value.(int))
		default:
			return 0, fmt.Errorf("%v is invalid value", value)
		}
	}
	return stack[0], nil
}
