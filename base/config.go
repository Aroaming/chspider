package base

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"sync"
)

type Log struct {
	File string `xml:"file"`
}

type Config struct {
	ConfigName string `xml:"configure"`
	Log        Log    `xml:"log"`
}

var (
	conf      *Config
	conf_once sync.Once
)

func Confige(name string) *Config {
	conf_once.Do(func() {
		if err := conf.init(name); err != nil {
			fmt.Println(err)
		}
	})
	return conf
}

func (c *Config) init(name string) error {
	content, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
	}
	err = xml.Unmarshal(content, &conf)
	if err != nil{
		fmt.Println(err)
	}
	return nil
}

func GetConfig() *Config {
	return conf
}
