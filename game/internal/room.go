package internal

type Room struct {
	RoomId     int
	OnlineNum  int
	Status     int  //-1  等待中   1  游戏中  2 满员
	PlayerList []GDYPlayer
	MsgSet     chan string  // 消息队列
}


type GDYRoom struct {
	RoomEntity     Room
	Multiple int //倍数
	CurUid   int //当前操作人
	CurCard	 CardArray//当前出的牌
	//牌型
	Players  GDYPlayer//用户的牌
	//已准备的人数
	//上一轮出的牌
	Boom int//炸弹
	Operatime int64//时间
}

func (room *Room) CreateRoom(){

}

func (room *Room) EnterRoom(){

}

func (room *Room) CloseRoom(){

}

func (room *Room) GetRoomInfo(){

}



