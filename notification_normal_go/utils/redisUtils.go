package utils

import (

	"fmt"
	"time"
	"sync"
	"github.com/garyburd/redigo/redis"

)

var mutex sync.Mutex

func Connect2Redis() (conn redis.Conn) {

	var connect_redis redis.Conn

	for {

		conn, err := redis.Dial("tcp", "212.64.85.208:6379")

		if err != nil {
			fmt.Println("connect to redis fail! reconnecting.....", err)
			time.Sleep(100) //等待0.1秒后重新尝试连接
		} else {
			connect_redis = conn
			break
		}
	}
	return connect_redis

}

func Sadd2Redis(conn redis.Conn, key string,from string,value string) bool {

	mutex.Lock()
	var err error
	if from==""{
		_, err = conn.Do("sadd", key,  value)
	}else{
		_, err = conn.Do("sadd", key,  from+":"+value)
	}
	mutex.Unlock()

	if err != nil {
		fmt.Println("fail to sadd the value ", err)
		return false
	}

	//持久化到本地
	//result := save2local(conn)
	//if !result {
	//	return false
	//}
	return true;
}

func Save2local(conn redis.Conn) bool {

	mutex.Lock()
	_, err := conn.Do("save")
	mutex.Unlock()

	if err != nil {
		fmt.Println("fail store whole data to local", err)
		return false
	}

	return true
}


func SmembersFromRedis(conn redis.Conn,key string)[]string{
	var keys []string
	result,err:=redis.Values(conn.Do("smembers",key))

	if err!=nil{
		fmt.Println("fail to hkeys the value ",err)
		return nil
	}

	for _,v:=range result{
		keys=append(keys,string(v.([]byte)))
	}

	return keys
}




func HgetKeyField(conn redis.Conn,key string,field string)string{

	result,err:=redis.String(conn.Do("hget",key,field))

	if err!=nil{
		fmt.Println("fail to hkeys the value ",err)
		return ""
	}

	return result

}

/*
func smembersFromRedis(conn redis.Conn,key string)[]string{

	result,err:=conn.Do("smember",key)

	if err!=nil{
		fmt.Println("fail to smember the value ",err)
		return nil
	}

	return result
}*/
