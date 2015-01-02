package main

import (
	"fmt"
	"time"
	"container/list"
	"os/exec"
	"os"
)

const naandir = "/tmp/naanfile/"
const naanbase = "naan"
const THRESHOLD=50
func getnaanname(count int) string {
        return fmt.Sprintf("%s%s.%d", naandir, naanbase, count)
}
func saveList(clipped *list.List) {
        os.Create(getnaanname(0)) 
        file, err := os.OpenFile(getnaanname(0), os.O_WRONLY, 0644)
        if err == nil {
                file.WriteString(string(clipped.Front().Value.([]byte)))
        }else {
                fmt.Println(err)
        }

        //Cleanup the last
        if _,err:=os.Stat(getnaanname(THRESHOLD)) ; err == nil {
                os.Remove(getnaanname(THRESHOLD))
        }

        //Cleanup. Move everything one more
        for num:=clipped.Len() ; num > 0 ; num-- {
                os.Rename(getnaanname(num-1), getnaanname(num))
        }
}

func main(){
        fi, err := os.Stat("/tmp/naan")
        if fi == nil {
                os.Mkdir(naandir, 0644)
        } else if !fi.IsDir() || err != nil{
                os.Exit(1)
        }
	clipped := list.New()
	for {
		current_clip, _ := exec.Command("pbpaste", "", "").Output()
		head:=clipped.Front()
		if (head == nil && current_clip != nil) || string(head.Value.([]byte)) != string(current_clip) {
                        if clipped.Len() >= 10 {
                                clipped.Remove(clipped.Back())
                        }
			clipped.PushFront(current_clip)
                        if clipped.Len() > 0 {
                                saveList(clipped)
                        } 
		} 

		time.Sleep(300 * time.Millisecond)
	}
}
