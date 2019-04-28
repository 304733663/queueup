package queue

import (
	"math"
	"queueup/commom"
	"queueup/libs/config"
	"time"
)

type queue struct {
	maxNum     int
	count      int
	currentNum int
	waitList   map[string]*udata
}

type udata struct {
	uid string
	pos int
}

var ch = make(chan *udata, config.Get("queue").Get("maxNum").MustInt())
var switchch = make(chan int)
var ins *queue

func New() *queue {
	if ins == nil {
		ins = &queue{maxNum: config.Get("queue").Get("maxNum").MustInt(), count: 0, currentNum: 0, waitList: map[string]*udata{}}
	}
	return ins
}

/**
 * 废弃
 */
func (self *queue) Run() {
	for {
		if !commom.NewOnline().IsFull() && self.QueueLength() > 0 {
			info := <-ch
			commom.NewOnline().AddOnline(info.uid)
			self.currentNum = info.pos
			delete(self.waitList, info.uid)
		}
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}

/**
 * 利用chan阻塞协程
 */
func (self *queue) Run2() {
	for {
		<-switchch
		//计算出目前还能进入几个人
		num := int(commom.NewOnline().EnterNum())
		chNum := self.QueueLength()
		diff := math.Min(float64(num), float64(chNum))
		for i := 0; i < int(diff); i++ {
			info := <-ch
			commom.NewOnline().AddOnline(info.uid)
			self.currentNum = info.pos
			delete(self.waitList, info.uid)
		}
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}

func (self *queue) WriteCh() {
	switchch <- 1
}

/**
 * 入队列
 * @param userid 用户Id
 */
func (self *queue) Push(userid string) {
	if self.QueueLength() < self.maxNum {
		self.count++
		info := &udata{uid: userid, pos: self.count}
		self.waitList[userid] = info
		ch <- info
	}
}

/**
 * 队列是否满员
 */
func (self *queue) IsFull() bool {
	return self.QueueLength() >= self.maxNum
}

/**
 * 队列长度
 */
func (self *queue) QueueLength() int {
	return len(ch)
}

/**
 * 是否在排队中
 */
func (self *queue) Exists(userid string) bool {
	if _, ok := self.waitList[userid]; ok {
		return true
	}
	return false
}

/**
 * 计算出当前等待位置
 */
func (self *queue) GetUserPos(userid string) int {
	return self.waitList[userid].pos - self.currentNum
}
