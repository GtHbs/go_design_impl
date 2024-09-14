package go_base_3

import (
	"errors"
	"fmt"
)

type Error interface {
	Caller() []CallerInfo // 存储调用链路
	Wrapped() []Error
	Code() int
	error
}

type CallerInfo struct {
	FuncName string // 调用函数名
	FileName string // 文件名
	FileLine int    // 出错行
}

func UsingPanic() {
	defer func() {
		if err := recover(); err != nil {
			switch x := err.(type) {
			case error:
				err = x
			case string:
				err = errors.New(x)
			default:
				err = fmt.Errorf("unknown panic: %v", x)
			}
		}
	}()
	//defer recover()	// 这种情况无法捕获异常，必须和panic发生的栈帧隔一个
	panic(123) // 异常抛出避免使用nil作为入参
}
