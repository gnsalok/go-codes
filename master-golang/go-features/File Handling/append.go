package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    //storing string as byte 
    mydata := []byte("All the data I wish to write to a file\n")

    // the WriteFile method returns an error if unsuccessful
    err := ioutil.WriteFile("myfile.data", mydata, 0777)

    if err != nil {
        fmt.Println(err)
    }

    data, err := ioutil.ReadFile("myfile.data")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Print(string(data))

    f, err := os.OpenFile("myfile.data", os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
		panic(err)
	}
	
    defer f.Close()

    if _, err = f.WriteString("new data that wasn't there originally\n"); err != nil {
        panic(err)
	}

	data, err = ioutil.ReadFile("myfile.data")
	
    if err != nil {
        fmt.Println(err)
    }

    fmt.Print(string(data))

}