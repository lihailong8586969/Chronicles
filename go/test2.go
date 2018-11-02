package main

import (
	"fmt"
	"log"
	"io"
	"os"
	"net/http"
	"time"
)


func main() {
	start := time.Now()
	log.Println("Starting")
	links := map[string]string{
		"Python-3.7.0a2.tgz": "http://ipv4.download.thinkbroadband.com/1GB.zip",
		"Python-3.6.0.tgz": "http://ipv4.download.thinkbroadband.com/1GB.zip",
		"Python-3.5.0.tgz": "http://ipv4.download.thinkbroadband.com/1GB.zip",
		"Python-3.4.0.amd64.msi": "http://ipv4.download.thinkbroadband.com/1GB.zip",
		"python-3.7.0a1.exe": "http://ipv4.download.thinkbroadband.com/1GB.zip",
		"python-2.7.9rc1.amd64.msi": "http://ipv4.download.thinkbroadband.com/1GB.zip",
		"python-2.7.9-macosx10.6.pkg": "http://ipv4.download.thinkbroadband.com/1GB.zip",
		"python-3.6.2rc2-macosx10.6.pkg": "http://ipv4.download.thinkbroadband.com/1GB.zip",
		"python-3.6.2rc1.exe": "http://ipv4.download.thinkbroadband.com/1GB.zip",
		"python-3.6.2.exe": "http://ipv4.download.thinkbroadband.com/1GB.zip",

	}

	for n, link :=range(links){
		resp, err := http.Get(link)
		if err != nil {
			log.Println("Could not download link", link)
		}
		defer resp.Body.Close()


		//Create file object 
		file, err := os.Create("/tmp/" + n)
		if err != nil {
			log.Fatal("Could not create file", n)
		}

		log.Println("Downloading File", n)
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Fatal(err)

		}
		file.Close()
		log.Println("File downloaded")

	}
	
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("total time to run the program", elapsed)
}