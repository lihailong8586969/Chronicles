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