package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func DownLoadFile(path string) ([]byte, error) {
	var buf []byte
	err := Retry(func() error {
		//_, err := url.ParseRequestURI(path)
		//uri, err := url.ParseRequestURI(path)
		//if err != nil {
		//	fmt.Println("文件地址错误：" + path)
		//	return err
		//}
		res, err := http.Get(path)
		//res, err := http.Get("http://"+uri.Host + uri.Path)
		if err != nil {
			return err
		}
		defer func() {
			err := res.Body.Close()
			if err != nil {
				fmt.Println("下载文件,关闭res失败：url - "+path, err)
			}
		}()
		fmt.Println("res.ContentLength: ", res.ContentLength)
		buf = make([]byte, 0)
		if res.ContentLength <= 0 {
			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println("ReadAll读取文件流失败: ", err)
				return err
			}
			buf = data
		} else {
			buff := bytes.NewBuffer(make([]byte, 0, res.ContentLength))
			_, err = buff.ReadFrom(res.Body)
			if err != nil {
				fmt.Println("ReadFrom读取图片流失败: ", err)
				return err
			}
			buf = buff.Bytes()
			buff.Reset()
		}
		return nil
	}, 3, 1*time.Second)
	return buf, err
}
