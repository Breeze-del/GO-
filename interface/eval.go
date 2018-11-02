package _interface

import (
	"fmt"
	"math"
	"testing"
)

// 表示任意表达式
type Expr interface {
	// 求职器 根据绑定environment变量返回表达式的值
	Eval(env Env) float64
}

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

func (v Var) Eval(env Env) float64 {
	return env[string(v)]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.Eval(env)
	case '-':
		return -u.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", c.fn))
}

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F ‐ 32)", Env{"F": -40}, "‐40"},
		{"5 / 9 * (F ‐ 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F ‐ 32)", Env{"F": 212}, "100"},
	}
	var prevExpr string
	for _, test := range tests {
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Sprintf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}
