package commom

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Httprequest struct {
	req    *http.Request
	params map[string]interface{}
}

func NewHttpRequest() *Httprequest {
	self := new(Httprequest)
	return self
}

func (self *Httprequest) SetParms(req *http.Request) {
	self.req = req
	self.analysisParam()
}

func (self *Httprequest) analysisParam() {
	self.params = make(map[string]interface{})
	self.req.ParseForm()
	for k, v := range self.req.Form {
		self.params[k] = v[0]
	}
	result, _ := ioutil.ReadAll(self.req.Body)
	if len(result) > 0 {
		var f interface{}
		json.Unmarshal(result, &f)
		m := f.(map[string]interface{})
		for k, v := range m {
			self.params[k] = v
		}
	}
}

func (self *Httprequest) GetParam(key string) interface{} {
	return self.params[key]
}

func (self *Httprequest) GetString(key string) string {
	return self.params[key].(string)
}
