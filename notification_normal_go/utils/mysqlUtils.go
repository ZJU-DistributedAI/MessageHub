package utils

import (
	"../bean"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"
)

const (
	userName = "root"
	password = "12345678"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "distributeai"
)

var db *sql.DB

func InitMysqlConnection() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"

	for {
		if db == nil {

			mutex.Lock()
			if db == nil {
				path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName,
					"?charset=utf8"}, "")

				//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
				conn, err := sql.Open("mysql", path)
				if err != nil {
					log.Println(err)
				}
				db = conn
				//设置数据库最大连接数
				db.SetConnMaxLifetime(100)
				//设置上数据库最大闲置连接数
				db.SetMaxIdleConns(10)
				//验证连接
				if err := db.Ping(); err != nil {
					log.Println("opon database fail", err)
					time.Sleep(500) //每隔0.5秒尝试重新连接
				} else {
					break
				}
			}
			mutex.Unlock()

		} else {
			break
		}
	}

}

func GetDockerStatus(useraccount string) (byte){

	db = GetMysqlConnect()

	var userdockerstatus bean.UserDockerStatus
	var user bean.User

	err := db.QueryRow("select id from users where useraccount = ?",
		useraccount).Scan(&user.Id)

	if err != nil {
		log.Println("获取用户信息失败", err)
	}

	err = db.QueryRow("select dockerstatus from user_dockerstatus where userid = ?",
		user.Id).Scan(&userdockerstatus.DockerStatus)

	if err != nil {
		log.Println("获取docker状态失败", err)
		return 0
	}

	return userdockerstatus.DockerStatus

}

func UpdateDockerStatus(useraccount string, dockerstatus int) {

	db = GetMysqlConnect()

	tx, err := db.Begin()
	if err != nil {
		log.Println("start transaction fail: ", err)
		return
	}
	var user bean.User
	err = db.QueryRow("select id from users where useraccount = ?", useraccount).Scan(&user.Id)
	if err != nil {
		ptmt, err := db.Prepare("insert into users (`useraccount`) values (?)")

		if err != nil {
			log.Println("prepare fail: ", err)
		}
		result, err := ptmt.Exec(useraccount)
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
		ptmt, err = db.Prepare("insert into user_dockerstatus (`userid`,`dockerstatus`) values (?,?)")
		if err != nil {
			log.Println("prepare fail: ", err)
		}
		lastid, _ := result.LastInsertId()
		_, err = ptmt.Exec(lastid, 0)
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}

	} else {
		ptmt, err := db.Prepare("update user_dockerstatus set dockerstatus = ? where userid = ?")
		if err != nil {
			log.Println("prepare fail: ", err)
		}
		_, err = ptmt.Exec(dockerstatus, user.Id)
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}

}

func GetMysqlConnect() (*sql.DB) {
	if db != nil {
		return db
	} else {
		InitMysqlConnection()
		return db
	}
}
