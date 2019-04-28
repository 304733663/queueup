package server

import (
	"fmt"
	"net/http"
	"queueup/commom"
	"queueup/controllers"
	"queueup/libs/config"
	"queueup/libs/logger"
	"reflect"
	"strings"
	"time"
)

type server struct {
	server *http.Server
}

var regStruct map[string]interface{}
var methodMap = map[string]map[string]string{}

func New() *server {
	self := &server{
		server: &http.Server{
			Addr:           config.Get("server_addr").MustString(),
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
	return self
}

func (self *server) handler(w http.ResponseWriter, req *http.Request) {
	pathInfo := strings.Trim(req.URL.Path, "/")
	spt := strings.Split(pathInfo, "/")
	var action string = "index"
	var actionName string = "Index"
	var ctr string = "User"
	if len(spt) > 1 {
		ctr = strings.Title(spt[0])
		action = strings.ToLower(spt[1])
	}
	if regStruct[ctr] != nil {
		controller := reflect.ValueOf(regStruct[ctr])
		methodList := methodMap[ctr]
		if _, ok := methodList[action]; ok {
			actionName = methodList[action]
		}
		method := controller.MethodByName(actionName)
		if !method.IsValid() {
			method = controller.MethodByName(strings.Title("index"))
		}
		responseValue := reflect.ValueOf(w)
		//设置参数
		httprequest := commom.NewHttpRequest()
		httprequest.SetParms(req)
		method.Call([]reflect.Value{responseValue, reflect.ValueOf(httprequest)})
	}
}

func init() {
	registCtrl()   //注册入口
	registMethod() //注册方法
}

/**
 * 注册contrl
 */
func registCtrl() {
	regStruct = make(map[string]interface{})
	regStruct["User"] = &controllers.UserController{}
}

/**
 * 注册所有contrl类的方法
 */
func registMethod() {
	for ctrName, ctr := range regStruct {
		v := reflect.ValueOf(ctr)
		t := reflect.TypeOf(ctr)
		mm := make(map[string]string)
		for i := 0; i < v.NumMethod(); i++ {
			name := t.Method(i).Name
			mm[strings.ToLower(name)] = name
		}
		methodMap[ctrName] = mm
	}
}

func (self *server) Run() {
	self.httpd()
}

func (self *server) httpd() bool {
	http.HandleFunc("/", self.handler)
	err := self.server.ListenAndServe()
	if err != nil {
		fmt.Println("监听失败：", err.Error())
		logger.New().Info("START", "监听失败"+err.Error())
		return false
	}
	return true
}
