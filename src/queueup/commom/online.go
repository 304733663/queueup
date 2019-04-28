package commom

import (
	"math"
	"queueup/libs/config"
)

type Online struct {
	currentNum int
	maxNum     int
	onlineList map[string]int
}

var ins *Online = &Online{0, config.Get("onlineMax").MustInt(), make(map[string]int)}

func NewOnline() *Online {
	return ins
}

func (self *Online) GetCurrentNum() int {
	return self.currentNum
}

func (self *Online) EnterNum() float64 {
	return math.Max(float64(self.maxNum-self.GetCurrentNum()), float64(0))
}

func (self *Online) IsFull() bool {
	return self.currentNum >= self.maxNum
}

func (self *Online) AddOnline(userid string) {
	self.onlineList[userid] = 1
	self.currentNum++
}

func (self *Online) OffOnline(userid string) {
	self.currentNum = int(math.Max(float64(self.currentNum)-float64(1), float64(0)))
	delete(self.onlineList, userid)
}

func (self *Online) IsOnline(userid string) bool {
	if _, ok := self.onlineList[userid]; ok {
		return true
	}
	return false
}
