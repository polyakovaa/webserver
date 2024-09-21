package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/polyakovaa/standartserver/internal/app/api"
)

var (
	configPath string
)

func init() {
	//скажем, что наше приложение будет на этапе запуска получать путь дл конфиг файла извне
	flag.StringVar(&configPath, "path", "configs/app.toml", "path to config path in .toml")
}

func main() {
	//в этот момент инициализируется переменная configpath
	flag.Parse()
	log.Println("it works")

	//server instance initialization
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config) //десериализация томл
	if err != nil {
		log.Println("Can not find configs file, using default:", err)
	}

	server := api.New(config)
	//api server start
	/*if err := server.Start(); err != nil {
		log.Fatal(err)
	}
		ИЛИ */

	log.Fatal(server.Start())
}
