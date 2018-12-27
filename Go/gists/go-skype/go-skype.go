// 原链接： https://gist.github.com/vodolaz095/9faa709d1cb78ee9606dafd2521eef73

package main

// Accessing last 10 messages in Skype using Golang.

//HOW TO RUN this application:

// 1. Install Go - https://golang.org/doc/install.

// 2. Instal dependencies by running in console
// go get github.com/jinzhu/gorm
// go get github.com/jinzhu/gorm/dialects/sqlite

//3. Update `databasePath` variable to actual skype path

//4. Run application by `go run go-skype.go

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //is is ok
)

type Message struct {
	ID           int64
	Author       string `gorm:"column:author"`
	FromDispname string `gorm:"column:from_dispname"`
	Timestamp    int64  `gorm:"column:timestamp"`
	BodyXML      string `gorm:"column:body_xml"`
}

func (Message) TableName() string {
	return "messages"
}

func main() {

	//TODO - change this path to location of main.db on windows!!!
	//this path works for me for Skype v 4.3.0.37 running under Fedora 23 linux

	databasePath := "/home/vodolaz095/.Skype/nowak.anatolij/main.db"
	//Sorry, I no idea where does skype stores database on windows(
	//for windows it will be something like c:\skype\main.db

	fmt.Printf("Using skype database from file of %s\n", databasePath)

	db, err := gorm.Open("sqlite3", databasePath)
	if err != nil {
		panic(err)
	}

	//verify that database is opened
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	//initialize array of messages
	var messages []Message

	//extract last 10 messages
	err = db.Order("id DESC").Limit(10).Find(&messages).Error
	if err != nil {
		panic(err)
	}

	//Output messages to screen in readable format
	for _, v := range messages {
		timestampPrepared := time.Unix(int64(v.Timestamp), 0)
		fmt.Printf("#%v FROM: %s TO: %s Timestamp: %v\n", v.ID, v.Author, v.FromDispname, timestampPrepared.Format("15:04:05"))
		fmt.Println(v.BodyXML)
	}
}