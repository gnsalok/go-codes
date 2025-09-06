package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.Open("test.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")

	body, readErr := ioutil.ReadAll(data)
	fmt.Println(string(body))

	if readErr != nil {
		log.Fatal(readErr)
	}

	fmt.Println("Encoding....")
	sEnc := b64.StdEncoding.EncodeToString([]byte(body))
	fmt.Println(sEnc)

	fmt.Println("Decoding....")
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

	if strings.Compare(string(body), string(sDec)) == 0 {
		fmt.Println("Both the strings match.")
	} else {
		fmt.Println("The strings do not match.")
	}

	defer data.Close()

}
