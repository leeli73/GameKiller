package main

import (
	"encoding/json"
	"io/ioutil"
	"path"
)

type ConfigStruct struct {
	Address    string `json:Address`
	Port       int    `json:Port`
	CacheFile  string `json:CacheFile`
	TTL        int    `json:TTL`
	WebAddress string `json:WebAddress`
	WebPort    int    `json:WebPort`
	WebPasswd  string `json:WebPasswd`
}

var Config ConfigStruct

func InitConfig() error {
	data, err := ioutil.ReadFile(path.Join("conf", "config.json"))
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &Config)
	if err != nil {
		return err
	}
	return nil
}
