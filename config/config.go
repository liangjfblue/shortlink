/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package config

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

type config struct {
	Server    server    `yaml:"server"`
	Mysql     mysql     `yaml:"mysql"`
	Redis     redis     `yaml:"redis"`
	ShortLink shortLink `yaml:"shortLink"`
	Log       log       `yaml:"log"`
	Etcd      etcd      `yaml:"etcd"`
}

type server struct {
	RunMode     string `yaml:"runMode"`
	Port        int    `yaml:"port"`
	RpcPort     int    `yaml:"rpcPort"`
	Network     string `yaml:"network"`
	PingTime    int    `yaml:"pingTime"`
	PingUrl     string `yaml:"pingUrl"`
	OpenIpLimit int    `yaml:"openIpLimit"`
}

type mysql struct {
	Host        string `yaml:"host"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
}

type redis struct {
	Host string `yaml:"host"`
}

type shortLink struct {
	ShortUrl string `yaml:"shortUrl"`
}

type log struct {
	ReportCaller bool `yaml:"reportCaller"`
	Level        int  `yaml:"level"`
}

type etcd struct {
	Host string `yaml:"host"`
}

var (
	onceDo  sync.Once
	_config config
)

func Init(filePath string) {
	onceDo.Do(func() {
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		if err := yaml.Unmarshal(data, &_config); err != nil {
			panic(err)
		}
	})
}

func Config() *config {
	return &_config
}
