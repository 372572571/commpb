package main

import (
	"fmt"
	"time"

	"github.com/372572571/commpb"
	easyjson "github.com/mailru/easyjson"
)

// 消耗很低
func main() {
	time_test(func() {
		str := "{\"code\":1,\"message\":\"SUCCESS OK\"}"
		var state = &commpb.Status{}
		easyjson.Unmarshal([]byte(str), state)
		fmt.Println(state)
	},
	)
}

func time_test(call func()) {
	s := time.Now().UnixMilli()
	defer func() {
		fmt.Println(time.Now().UnixMilli() - s)
	}()
	call()
}
