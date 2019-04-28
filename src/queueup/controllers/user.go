package controllers

import (
	"fmt"
	"net/http"
	"queueup/commom"
	"queueup/libs/config"
	"queueup/libs/queue"
)

type UserController struct {
}

/**
 * 用户检测获取游戏链接
 */
func (self *UserController) CheckLink(w http.ResponseWriter, request *commom.Httprequest) {
	userId := request.GetString("userId")
	if commom.CheckValueEmpty(userId) {
		fmt.Fprintln(w, "您还没有登陆，无法得到游戏区服链接")
	}
	if commom.NewOnline().IsFull() && !commom.NewOnline().IsOnline(userId) { //如果在线人数满了，就进入队列
		//先看用户是不是在队列中
		if queue.New().Exists(userId) {
			str := fmt.Sprintf("你在继续排队中，你当前的位置%d", queue.New().GetUserPos(userId))
			fmt.Fprintln(w, str)
		} else if !queue.New().IsFull() {
			queue.New().Push(userId)
			str := fmt.Sprintf("你在排队中，你当前的位置%d", queue.New().GetUserPos(userId))
			fmt.Fprintln(w, str)
		} else {
			fmt.Fprintln(w, "现在人员太多，您等会再来吧")
		}
	} else if !commom.NewOnline().IsOnline(userId) {
		commom.NewOnline().AddOnline(userId)
		fmt.Fprintln(w, config.Get("gameUrl").MustString())
		fmt.Println(userId, commom.NewOnline().GetCurrentNum())
	} else {
		fmt.Fprintln(w, config.Get("gameUrl").MustString())
	}
}

/**
 * 用户离线
 */
func (self *UserController) OffOnline(w http.ResponseWriter, request *commom.Httprequest) {
	userId := request.GetString("userId")
	if commom.CheckValueEmpty(userId) {
		fmt.Fprintln(w, "用户ID错误")
	}
	if commom.NewOnline().IsOnline(userId) {
		commom.NewOnline().OffOnline(userId)
		queue.New().WriteCh()
		fmt.Println(userId, "离开，目前在线人数", commom.NewOnline().GetCurrentNum())
	}
}

func (self *UserController) Index(w http.ResponseWriter, request *commom.Httprequest) {
	fmt.Fprint(w, "欢迎进来")
}
