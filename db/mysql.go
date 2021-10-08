package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"laji/v1/config"
	"sync"
)

var (
	ConnectOnce sync.Once
	mysql       *Mysql
	err         error
	Host        string
	User        string
	Pass        string
	Port        string
	Dbname      string
	Charset     string
)

type Mysql struct {
	DB *gorm.DB
}

func newMysql() *Mysql {
	return &Mysql{}
}

/**
 * @author lidong
 * @description 获取mysql
 * @date 16:44 2021/9/9
 * @param
 * @return
 **/
func GetMysql() (*Mysql, error) {
	ConnectOnce.Do(func() {
		mysql, _ = prepareConnect()
	})
	return mysql, nil
}

/**
 * @author lidong
 * @description //mysql连接
 * @date 16:43 2021/9/9
 * @param
 * @return
 **/
func prepareConnect() (*Mysql, error) {
	cfg, _ := config.GetConfig()
	Host = cfg.Mysql["host"]
	User = cfg.Mysql["user"]
	Pass = cfg.Mysql["pass"]
	Port = cfg.Mysql["port"]
	Dbname = cfg.Mysql["dbname"]
	Charset = cfg.Mysql["charset"]
	mysql := newMysql()
	fmt.Println(User + ":" + Pass + "@tcp(" + Host + ":" + Port + ")/" + Dbname + "?charset=" + Charset + "&parseTime=True&loc=Local")
	mysql.DB, err = gorm.Open("mysql", User+":"+Pass+"@tcp("+Host+":"+Port+")/"+Dbname+"?charset="+Charset+"&parseTime=True&loc=Local")
	if err != nil {
		log.Error("mysql: Mysql connection failed!")
		return nil, err
	}
	log.Print("mysql: Mysql connection success!")
	return mysql, nil
}
