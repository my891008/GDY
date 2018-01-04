package internal

type Card struct {
	Val     int  //用户id
	//Type    int  //花色 1 黑桃 2 红桃 3 梅花 4 方块  5小王  6 大王
}

type CardArray struct {
	Num int
	Poke []Card
}


func (ca *CardArray) GetAllCard() CardArray{
	po := CardArray{Num:54}
	cards :=make([]Card,0)
	for i:=0 ; i < 4; i++ {
		for m := 0; m < 13; m++ {
			var card Card
			card.Val = int(byte(i)<<4 | byte(m))
			cards=append(cards,card)
		}
	}
	cards=append(cards,Card{int(byte(0x4d))})  //大王
	cards=append(cards,Card{int(byte(0x4e))})	 //小王
	po.Poke=cards
	return po
}