package main

import (
	"github.com/ayhangenc/ejder/db"
	"github.com/ayhangenc/ejder/video"
)

func main() {
	db.DBQuery()
	video.Tool()
	video.Bool()
	db.SQL()
}
