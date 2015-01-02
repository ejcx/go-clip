package main 

import (
        "os"
        "os/exec"
        "fmt"
        "io/ioutil"
        "flag"
)
const naandir = "/tmp/naanfile/"
const naanbase = "naan"
func getnaanname(count string) string {
        return fmt.Sprintf("%s%s.%s", naandir, naanbase, count)
}
func main(){
        num := flag.String("n", "", "The copy you want")
        flag.Parse()
        if num == nil || len(*num) == 0 {
                fmt.Println("-n copy number required")
                os.Exit(1)
        }
        //Save the bytes
        naanfile := getnaanname(*num)
        _, err := os.Stat(naanfile)
        if err == nil {
                os.Remove(naanfile) 
        }
        bytes, _ := ioutil.ReadAll(os.Stdin)
        ioutil.WriteFile(naanfile, bytes, 0644)
       
         //Copy to clipboard 
        copy := exec.Command("pbcopy")
        in, _:= copy.StdinPipe()
        in.Write(bytes)
        in.Close()
}
