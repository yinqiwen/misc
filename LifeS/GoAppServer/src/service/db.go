package service

import (
	"github.com/ziutek/mymysql/mysql"
	"log"
	//"os"
	_ "github.com/ziutek/mymysql/native" // Native engine
	//_ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
)

var userDbConn mysql.Conn
var orderDbConn mysql.Conn

func getUserDBConn() mysql.Conn{
	if nil != userDbConn {
		err := userDbConn.Ping()
		if nil == err {
			return userDbConn
		}
		closeUserDBConn()
	}
	userDbConn = mysql.New("tcp", "", "127.0.0.1:3306", "root", "wqy123", "asp_user")
	userDbConn.Register("set names utf8")
	err := userDbConn.Connect()
	if err != nil {
		userDbConn = nil
		log.Printf("Faield to connect user db")
	}
	return userDbConn
}

func closeUserDBConn() {
	if nil != userDbConn {
		userDbConn.Close()
		userDbConn = nil
	}
}

func getOrderDBConn() mysql.Conn{
	if nil != orderDbConn {
		err := orderDbConn.Ping()
		if nil == err {
			return orderDbConn
		}
		closeOrderDBConn()
	}
	orderDbConn = mysql.New("tcp", "", "127.0.0.1:3306", "root", "wqy123", "asp_order")
	orderDbConn.Register("set names utf8")
	err := orderDbConn.Connect()
	if err != nil {
		orderDbConn = nil
		log.Printf("Faield to connect user db")
	}
	return orderDbConn
}

func closeOrderDBConn() {
	if nil != orderDbConn {
		orderDbConn.Close()
		orderDbConn = nil
	}
}
