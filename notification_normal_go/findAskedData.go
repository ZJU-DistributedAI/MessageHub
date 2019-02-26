package main

import (
	"github.com/garyburd/redigo/redis"
	"strings"
	"./utils"
)

func askDataClient(conn redis.Conn,metaData string){

	//请求数据方的同意,这里暂时先标记到redis，然后让数据方自己请求看哪些数据被访问了
	datas:=utils.SmembersFromRedis(conn,"metadata")

	for i:=0;i<len(datas);i++{
		hashes:=strings.Split(datas[i],":")
		if metaData==hashes[1]{
			utils.Sadd2Redis(conn,hashes[0],"",metaData)
		}
	}


}


func askComputing(conn redis.Conn,computing string){

	//请求运算方的同意,这里也先暂时标记到redis,然后让运算方看那些运算资源被请求了
	computings:=utils.SmembersFromRedis(conn,"computing")

	for i:=0;i<len(computings);i++{
		hashes:=strings.Split(computings[i],":")
		if computing==hashes[1]{
			utils.Sadd2Redis(conn,hashes[0],"",computing)
		}
	}
}