package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := foo()
	if err != nil {
		fmt.Printf("%+v", err)
	}
}

func foo() error {
	// 因为sql.ErrNoRows是原生error的一个别名,不是包装error, 遇到问题不好定位位置，所以用pkg/errors包装一下，可以输出堆栈,方便定位
	return errors.WithStack(sql.ErrNoRows)
}
