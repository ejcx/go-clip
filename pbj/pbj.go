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
        naanfile := getnaanname(*num)
        _, err := os.Stat(naanfile)
        if err == nil {
                reader, _:= os.Open(naanfile)
                paste, _:= ioutil.ReadAll(reader)
                copy := exec.Command("pbcopy")
                in, _:= copy.StdinPipe()
                copy.Start()
                in.Write(paste)
                in.Close()
                copy.Wait()
        } else {
                os.Exit(1)
        }
}
