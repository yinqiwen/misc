package common

import (
	"log"
	"strings"
)

var MailCc []string = make([]string, 0)
var MailTo []string = make([]string, 0)

var ListenAddress string = "0.0.0.0:8080"
var EnableDaemon bool

func InitConfig(conf string) bool {
	cfg, err := LoadIniFile(conf)
	if nil != err {
		log.Fatalf("Failed to load config file for reason:%s\n", err.Error())
		return false
	}
	if addrs, exist := cfg.GetProperty("Mail", "OrderTo"); exist {
		MailTo = strings.Split(addrs, ",")
	} else {
		log.Printf("Empty mail to address")
	}

	if addrs, exist := cfg.GetProperty("Mail", "OrderCc"); exist {
		MailCc = strings.Split(addrs, ",")
	} else {
		log.Printf("Empty mail cc address")
	}

	if addr, exist := cfg.GetProperty("Server", "Listen"); exist {
		ListenAddress = addr
	}
	
	if enable, exist := cfg.GetIntProperty("Server", "Daemon"); exist {
		EnableDaemon = enable == 1
	}
	return true
}
