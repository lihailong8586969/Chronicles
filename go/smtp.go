package main

import (
	"fmt"
	"github.com/go-gomail/gomail"
)

func main () {
	fmt.Printf("Hello")
	m := gomail.NewMessage()
	m.SetHeader("From", "abc@abc.com")
	m.SetHeader("To", "abc@abc.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer("email-smtp.com", 25, "username", "password")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}


}

func dump_response_to_file(body []byte){
	result := bytes.Split(body, []byte("filename="))
		len_divisions := len(result)
		i:=1
		for (i < len_divisions) {
			main_content  := bytes.Split(result[i], []byte("Content-"))
			filename := string(bytes.Trim(main_content[0],"\"\r\n"))
			index := bytes.Index(main_content[1],[]byte("\r\n"))
			data_to_be_iterated := main_content[1]
			split_content := data_to_be_iterated[index+1:]
			file_content  := bytes.Split(split_content, []byte("------WebKit"))
			file_data := file_content[0] //This is where the data lies

			//File Writer
			err = ioutil.WriteFile(filename, file_data, 0666)
			if err != nil {
				panic(err)
			}
			i++
		}
}