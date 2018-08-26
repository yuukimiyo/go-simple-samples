package main

import (
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config 設定ファイルのための構造体
type Config struct {
	Mysql struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Db   string `yaml:"db"`
	}
}

func main() {

	var configFile = "C:/usr/workspace/go-simple-samples/yaml/config.yml"
	var conf Config

	// 設定ファイルを読み込む
	confYaml, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	// 設定ファイルの値をconfへ格納
	err = yaml.Unmarshal(confYaml, &conf)
	if err != nil {
		panic(err)
	}
}
