package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {

	t :=time.Now().UnixNano() / 1e6
	fmt.Println(t)

	//filePath := "C:/Users/lanningbo/Desktop/image/"
	//imageData := readImage(filePath)
	//fmt.Println(imageData)
}

func readImage(filePath string) []string {
	imageDates := make([]string, 0)
	fileInfos, _ := ioutil.ReadDir(filePath)
	for _, f := range fileInfos {
		fileName := f.Name()
		fmt.Println(fileName)
		file, e := os.Open(filePath + fileName)
		if e != nil {
			panic(e)
		}
		bys, e := ioutil.ReadAll(file)
		if e != nil {
			panic(e)
		}
		base64Str := base64.StdEncoding.EncodeToString(bys)
		fmt.Println("base64Str: ", base64Str)
		e = file.Close()
		if e != nil {
			fmt.Println("file close error")
		}
		imageDates = append(imageDates, string(base64Str))
	}

	return imageDates
}
