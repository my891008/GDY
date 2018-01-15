package internal

import (
	"sort"
)

type Card struct {
	Val     int  //用户id
	//Type    int  //花色 1 黑桃 2 红桃 3 梅花 4 方块  5小王  6 大王
}

type CardArray struct {
	Poke []Card
}

type CardType int

// iota 初始化后会自动递增
const (
	Dan CardType = iota // value --> 0
	Dui              // value --> 1
	Shun            // value --> 2
	Zha           // value --> 3
)


func (ca CardArray) Len() int           { return len(ca.Poke) }
func (ca CardArray) Less(i, j int) bool { return ca.Poke[i].Val < ca.Poke[j].Val }
func (ca CardArray) Swap(i, j int)      { ca.Poke[i].Val, ca.Poke[j].Val = ca.Poke[j].Val, ca.Poke[i].Val }

func (ca *CardArray) GetAllCard() CardArray{
	po := CardArray{}
	cards :=make([]Card,0)
	for i:=0 ; i < 4; i++ {
		for m := 3; m < 16; m++ {
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


//检测牌型
func getCardType(myCards []Card) CardType{

	var cardType CardType

	if len(myCards) != 0 {
		// 大概率事件放前边，提高命中率
		if (isDan(myCards)) {
			cardType = Dan;
		} else if (isDuiZi(myCards)) {
			cardType = Dui;
		} else if (isZhadan(myCards)) {
			cardType = Shun;
		} else if (isShunZi(myCards)) {
			cardType = Zha;
		}
	}

	return cardType
}


func isDan(myCards []Card) bool{
	if len(myCards) == 1 {
		return true
	}else{
		return false
	}
}

func isDuiZi(myCards []Card) bool{
	if len(myCards) == 2 {
		val1:= myCards[0].Val
		val2:= myCards[0].Val
		if val1==77 || val2==78 || val1 ==78 || val2 ==77{
			return true
		}
		if (val1%16) == (val2%16) {
			return true
		}else{
			return false
		}
	}else{
		return false
	}
}

func isShunZi(myCards []Card) bool{
	var flag = true
	if len(myCards) >2  {
		var tmca CardArray
		tempcard:=make([]Card,0)
		for _,card:=range myCards{
			if card.Val > 70{
				tempcard=append(tempcard,card)
			}else{
				card.Val=card.Val%60
				tempcard=append(tempcard,card)
			}
		}
		tmca.Poke=tempcard

		sort.Sort(tmca)

		for i:=0;i<len(tmca.Poke)-1;i++ {
			prev := tmca.Poke[i].Val
			next := tmca.Poke[i+1].Val

			if prev > 70 || next >70{
				continue
			}

			if prev == 15 || next == 15{  //2 不在顺子里
				flag = false;
				break;
			}

			if (prev - next != -1) {
				flag = false;
				break;
			}

		}

	}else{
		flag=false
	}
	return flag

}

func isZhadan(myCards []Card) bool{
	var flag = true
	if len(myCards) > 2 {
		var val = 70
		for _,card:=range myCards{
			if card.Val >70 {
				continue
			}

			v:=card.Val%60
			if val == 70 {
				val=v
			}else if val != v{
				flag=false
				break
			}

		}

	}else{
		flag= false
	}
	return flag
}