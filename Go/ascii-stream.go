// 原地址 ：https://github.com/HFO4/plus1s.live/blob/master/stream.go


package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const MAX_FRAME = 360

var frames []string = make([]string, 0, MAX_FRAME)

func leftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func plus1s(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("User-Agent"), "curl") {
		http.Redirect(w, r, "https://github.com/HFO4/plus1s.live", http.StatusFound)
		return
	}
	i := 1
	loop := 0
	for {
		if i == MAX_FRAME {
			i = 0
			loop++
		}
		if loop > 20 {
			write(w, "You've waste too much time, we have to stop you.")
			return
		}
		write(w, frames[i])
		i++
	}
}

func write(w http.ResponseWriter, s string) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("Expected http.ResponseWriter to be an http.CloseNotifier")
	}
	fmt.Fprintf(w, s)
	fmt.Fprintf(w, "\033[2J\033[H")
	time.Sleep(100000000)
	flusher.Flush()
}

func prepare() {
	fmt.Println("Prepare frames...")
	for i := 0; i < MAX_FRAME; i++ {
		b, err := ioutil.ReadFile("pic/" + leftPad2Len(strconv.Itoa(i+1), "0", 3) + ".txt")
		if err != nil {
			fmt.Print(err)
		}
		str := string(b)
		frames = append(frames, str)
	}
	fmt.Println("Frames are ready.")
}

func main() {
	prepare()
	fmt.Println("Ready to start server")
	http.HandleFunc("/", plus1s)
	err := http.ListenAndServe(":1926", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}