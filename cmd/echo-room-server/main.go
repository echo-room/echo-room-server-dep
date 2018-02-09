package main

import (
	"github.com/echo-room/echo-room-server/pkg/api"
	"github.com/echo-room/echo-room-server/pkg/db"
)

func main() {
	db.Test()
	api.Start()
}
