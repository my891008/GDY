package internal
import (
	"github.com/name5566/leaf/gate"
)

type GDYPlayer struct {
	Uid     int  //用户id
	Name    string  //用户昵称
	Roomid  int   //玩家所在房间
	Score   int//当前分数
	WinNum  int//胜数
	LoseNum int//输数
	Card    []Card //当前牌
	Status  int//状态
	IsReady int//是否准备
	IsFZ    int//是否为房主
	operatime int64//操作时间
	agent     gate.Agent     //消息通道
}