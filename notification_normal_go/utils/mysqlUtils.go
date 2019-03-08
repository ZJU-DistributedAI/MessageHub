package utils

import (
	"database/sql"
	"log"
	"strings"
	"time"
	"../bean"
)

const (
	userName = "root"
	password = "123456"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "loginserver"
)



var db *sql.DB

func InitMysqlConnection(){
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"

	for{
		if db == nil {

			mutex.Lock()
			if db == nil {
				path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

				//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
				db, _ = sql.Open("mysql", path)
				//设置数据库最大连接数
				db.SetConnMaxLifetime(100)
				//设置上数据库最大闲置连接数
				db.SetMaxIdleConns(10)
				//验证连接
				if err := db.Ping(); err != nil{
					log.Println("opon database fail", err)
					time.Sleep(500)
				}else{
					break
				}
			}
			mutex.Unlock()

		}else{
			break
		}
	}

}



func UpdateDockerStatus(userAccount string, userId int, dockerstatus int){

	db = GetMysqlConnect()

	tx, err := db.Begin()
	if err != nil {
		log.Println("start transaction fail: ", err)
		return
	}
	var user bean.User
	err = db.QueryRow("select id from user_dockerstatus where userid = ?", userId).Scan(&user.Id)
	if err != nil {
		ptmt, err := db.Prepare("insert into users (`useraccount`) values (?)")

		if err != nil{
			log.Println("prepare fail: ", err)
		}
		_ , err = ptmt.Exec(userAccount)
		if err != nil{
			tx.Rollback()
		}else{
			tx.Commit()
		}
		ptmt, err = db.Prepare("insert into user_dockerstatus (`userid`,`dockerstatus`) values (?,?)")
		if err != nil{
			log.Println("prepare fail: ", err)
		}
		_ , err = ptmt.Exec(userId,0)
		if err != nil{
			tx.Rollback()
		}else{
			tx.Commit()
		}

	}else{
		ptmt, err := db.Prepare("update user_dockerstatus set dockerstatus = ? where userid = ?")
		if err != nil{
			log.Println("prepare fail: ", err)
		}
		_, err = ptmt.Exec(dockerstatus, userId)
		if err != nil{
			tx.Rollback()
		}else{
			tx.Commit()
		}
	}

}





func GetMysqlConnect()(*sql.DB){
	if db != nil{
		return db
	}else{
		InitMysqlConnection()
		return db
	}
}
