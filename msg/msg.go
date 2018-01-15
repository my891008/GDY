package msg

import (
	"github.com/name5566/leaf/network/json"
)


// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
var Processor = json.NewProcessor()

func init() {
	// 这里我们注册 JSON 消息
	Processor.Register(&Hello{})
	Processor.Register(&GameRequest{})
	Processor.Register(&GameResponce{})
	Processor.Register(&GamePlay{})
	Processor.Register(&EnterRoom{})
	Processor.Register(&EnterRoomResponce{})

}

// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Hello
type Hello struct {
	Name string
}

//进入房间msg
type EnterRoom struct{
	Uid    int
	Name   string
	Roomid int  //房间号
}

type EnterRoomResponce struct{
	Status  int  //1  进入成功  0  进入失败
	Msg     string
	Uid     int
	Players interface{}
}

type GameRequest struct {
	Roomid  int  //房间号
	Type  int  //1 发牌   2 补牌
}

type GameResponce struct {
	Uid     int
	Players interface{}
}

//出牌
type GamePlay struct{
	Roomid int  //房间号
	Uid    int
	Poke   interface{}
}