package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	controllers "TheAdidasTM/Controllers"
	models "TheAdidasTM/Models"

	"github.com/gin-gonic/gin"
)

var configPath string = "config.json"

func main() {
	//подгрузка конфига
	info, err := os.Stat("config.json")
	if err != nil {
		fmt.Printf("error while to get info about config file - %v\n", err)
		return
	}

	//читаем файл
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("error while trying to read %v - %v\n", info.Name(), err)
		return
	}

	//парсим байты в джсон
	var config models.Config
	if err = json.Unmarshal(data, &config); err != nil {
		fmt.Printf("error while trying to parse bytes to struct (config) - %v\n", err)
		return
	}
	var ip string = config.Ip
	var port string = config.Port

	//валидация конфига
	validIp := net.ParseIP(ip)
	var buildUtl string = fmt.Sprintf("%v%v", validIp, port)

	r := gin.Default()

	//эндпоинты
	r.POST("/adidas", controllers.RequestFromIlya)
	//эндпоинты

	//порт из конфига
	if err = r.Run(buildUtl); err != nil {
		fmt.Printf("error while trying to start server - %v\n", err)
		return
	}
}
