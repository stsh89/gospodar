package config

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"fmt"
)

var DatabaseHost string
var DatabaseSslmode string
var DatabaseUser string
var DatabaseName string

var SmtpUserName string
var SmtpPassword string
var SmtpDomain string
var SmtpPort string

var Port string

func InitGlobalVars() {
	raw, err := ioutil.ReadFile("../gospodar/app/config/secrets.json")
	json := string(raw)
	fmt.Println("Raw json")
	fmt.Println(json)

	if err != nil {
		panic(err)
	}

	value := gjson.Get(json, "database.host")
	DatabaseHost = value.String()
	value = gjson.Get(json, "database.sslmode")
	DatabaseSslmode = value.String()
	value = gjson.Get(json, "database.user")
	DatabaseUser = value.String()
	value = gjson.Get(json, "database.name")
	DatabaseName = value.String()
	value = gjson.Get(json, "smtp.userName")
	SmtpUserName = value.String()
	value = gjson.Get(json, "smtp.password")
	SmtpPassword = value.String()
	value = gjson.Get(json, "smtp.domain")
	SmtpDomain = value.String()
	value = gjson.Get(json, "smtp.port")
	SmtpPort = value.String()
	value = gjson.Get(json, "port")
	Port  = value.String()

	fmt.Println("Init global vars")
	fmt.Println(DatabaseHost)
}
