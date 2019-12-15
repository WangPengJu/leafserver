package msg

import (
	"github.com/name5566/leaf/network/json"
)

// var Processor network.Processor
var Processor = json.NewProcessor()

func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	Processor.Register(&Hello{})
}

type Hello struct {
	Name string
}
