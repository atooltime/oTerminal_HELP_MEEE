package main

import (
    "fmt"
    "strings"
    "log"
    "path/filepath"
    "os"
    "bufio"
    "os/exec"
)
//search the database

var result string 
func z(lmao string){
    f, err := os.Open("tosearch")
    if err != nil {
        print(err)
    }
    defer f.Close()

    // Splits on newlines by default.
        scanner := bufio.NewScanner(f)
        line := 1
        for scanner.Scan() {
            if strings.Contains(scanner.Text(), lmao) {
                result=scanner.Text()
                break
        }else{
            line++
}}
print("")
print(result+"\n")
fmt.Println("Install this? y/n")
var first string
fmt.Scanln(&first)
if first == "y" {
    findpacker()
}

}


//Parse the url 
func parser(){
    for i:=1; i<len(os.Args); i++ {
        print(os.Args[i]+"\n")
        z(os.Args[i]) 
    }

}


//Find Packer.init file 
func findpacker() {
    dir,goddamnit:=os.Getwd()
    if goddamnit !=nil{
        fmt.Println("Veryy weird path dude make sure to place this inside your .config/nvim")
    }
    var files []string 
    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

        if err != nil {

            fmt.Println(err)
            return nil
        }

        if !info.IsDir() && filepath.Ext(path) == ".lua" {
            files = append(files, path)
        }

        return nil
    })

    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        if (strings.Contains(file,"packer-config"))==true{
            fmt.Println(file)
        }
    }
    toexec := "sed -i '2i use \""+result+"\"'"+" "+files[0]
    out,err :=exec.Command("bash","-c",toexec).Output()
    if err !=nil{
        fmt.Println(err)
    }
    fmt.Println(string(out))
}
func completed(){
out,err :=exec.Command("bash","-c"," nvim --headless -c 'autocmd User PackerComplete quitall' -c 'PackerSync'").Output()
if err!=nil{
    fmt.Println("failed")
}else{
    fmt.Println("successfully done")
    fmt.Println(out)
}
}
func main(){
    
    parser()
    completed()
    //findpacker()
}
