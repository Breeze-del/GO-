package tempconv

import "errors"

type Function interface {
	Push(v interface{})
	Pop() (interface{}, error)
	Top() (interface{}, error)
}

//切片实现Stack 栈
type Stack []interface{}

func (stack *Stack) Push(v interface{}) {
	*stack = append(*stack, v)
}

//弹出栈顶元素，栈收缩
func (stack *Stack) Pop() (interface{}, error) {
	if len(*stack) == 0 {
		return nil, errors.New("stack empty")
	}
	v := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return v, nil
}

func (stack *Stack) Top() (interface{}, error) {
	if len(*stack) == 0 {
		return nil, errors.New("stack empty")
	}
	return (*stack)[len(*stack)-1], nil
}
