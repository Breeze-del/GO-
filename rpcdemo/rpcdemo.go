package rpcdemo

import "github.com/pkg/errors"

// 可以放入服务需要的配置
type DemoService struct {
}

type Args struct {
	A, B int
}

// 一个服务
func (d DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}
