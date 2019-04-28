package config

import (
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"strings"
)

//var ConfigPath string = os.Getenv("DELIVERY_CONFIG")
var ConfigPath string = "/Users/homework/data/work/gowork/queueup/conf"
var configJson *simplejson.Json

func init() {
	LoadConfig()
}

func LoadConfig() *simplejson.Json {
	c, err := ioutil.ReadFile(strings.TrimRight(ConfigPath, "/") + "/" + "config.json")
	if err != nil {
		panic(err)
	}

	json, err := simplejson.NewJson(c)
	configJson = json
	return configJson
}

func Get(key string) *simplejson.Json {
	return configJson.Get(key)
}
