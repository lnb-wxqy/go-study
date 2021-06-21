package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Main() {
	dUrl := "http://quote.eastmoney.com/center/"
	err := crawler(dUrl)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

}

func crawler(dUrl string) error {
	req, err := http.NewRequest(http.MethodGet, dUrl, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println("body: ", string(body))
	return nil
}


