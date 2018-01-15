package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/gate"
	"reflect"
	"server/msg"
	"time"

	"math/rand"
	"fmt"
)
var cards CardArray
var rooms map[string]*Room
var tmp int
var leftNum int = 54


func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	//发牌和摸牌
	handler(&msg.GameRequest{}, handleGameRequest)
	//进入房间和创建房间
	handler(&msg.EnterRoom{}, handleEnterGame)
	//出牌
	handler(&msg.GamePlay{}, handlePlayGame)



	players:=make([]GDYPlayer,0)

	rooms=make(map[string]*Room)
	room := new(Room)
	room.RoomId=1
	room.PlayerList=players

	rooms["1"] = room

	cards=cards.GetAllCard()




	fmt.Println(cards)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleEnterGame(args []interface{}){
	// 收到的 Hello 消息
	m := args[0].(*msg.EnterRoom)
	// 消息的发送者
	a := args[1].(gate.Agent)

	//获取房间信息
	k:=fmt.Sprint(m.Roomid)

	room:=rooms[k]

	players:=room.PlayerList

	//进入房间

	fmt.Println("房间信息：",room)


	//房间是否满员
	if len(players) >4{
		a.WriteMsg(&msg.EnterRoomResponce{
			Status: 0,
			Msg:"房间人数已满",
		})
	}else{

		//发送消息给其他玩家 提醒进入房间
		for _,p:=range players{
			p.agent.WriteMsg(&msg.EnterRoomResponce{
				Status: 1,
				Uid:m.Uid,
				Msg:"进入房间成功",
			})
		}


		play := new (GDYPlayer)
		play.Uid=m.Uid
		play.IsReady=1  //进入自动准备
		play.Name =m.Name
		play.agent = a

		if len(players) == 0 {
			play.IsFZ = 1
		}

		players=append(players,*play)

		room.PlayerList=players
		room.OnlineNum++

		rooms[k] = room


		fmt.Println(players)
		a.WriteMsg(&msg.EnterRoomResponce{
			Status: 1,
			Msg:"进入房间成功",
			Players:players,
		})
	}
}


func handleGameRequest(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.GameRequest)

	//获取房间信息
	k:=fmt.Sprint(m.Roomid)

	room:=rooms[k]

	players:=room.PlayerList

	// 输出收到的消息的内容
	log.Debug("发牌和摸牌 %v", m.Roomid)

	if m.Type == 1 {
		players=getInitPoker(players,5)

		// 发牌给每个用户  每个用户收到的消息只为自己的
		for _,player := range players{
			player.agent.WriteMsg(&msg.GameResponce{
				Players:player,
			})
		}

	}

}


func getInitPoker(players []GDYPlayer,num int) (b []GDYPlayer){

	b = players

	rand.Seed(time.Now().UnixNano())

	for i:=0;i<(num+1);i++{
		for j:=0;j<len(players);j++{
			if (54-leftNum) <= len(players) * num {
				cardId := rand.Intn(leftNum)
				b[j].Card = append(b[j].Card,cards.Poke[cardId])
				cards.Poke[cardId]=cards.Poke[leftNum-1]
				leftNum--
			}

		}
	}

	return b
}



func handlePlayGame(args []interface{}){
	// 收到的 Hello 消息
	m := args[0].(*msg.GamePlay)

	//获取房间信息
	k:=fmt.Sprint(m.Roomid)

	room:=rooms[k]

	players:=room.PlayerList

	// 输出收到的消息的内容
	log.Debug("出牌 %v", m.Roomid)


	for _,player := range players{
		if player.Uid != m.Uid{

			//mycards:=m.Poke.([]Card)
			//cardtype:=getCardType(mycards)

			//checkCards(cardtype,player.Card)

			player.agent.WriteMsg(m)
		}

	}
}


