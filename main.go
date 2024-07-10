package main

import (
	"github.com/AswinJoseOpen/websocket/server"
	"github.com/AswinJoseOpen/websocket/utils"
)

func main() {
	utils.LoadEnv()
	redisClient := utils.InitializeRedisClient()
	go server.StartAuth(redisClient)
	server.StartVoteServer(redisClient)

}
