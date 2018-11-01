package _interface

// 表示任意表达式
type Expr interface{}

// 变量
type Var string

// 浮点型常量
type literal float64

// 运算符号
type unary struct {
	op rune // + - 其中之一
	x  Expr
}
type binary struct {
	op   rune // + - * / 其中之一
	x, y Expr
}

// 表示对一个函数的调用
type call struct {
	fn   string // pow sin "sqrt" 之一
	args []Expr
}

// 环境变量 将名字映射成值
type Env map[string]float64
