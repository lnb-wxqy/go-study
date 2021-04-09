package main

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"go-study/httpdemo/model"
	"go-study/utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var workerHttpClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:        20,
		MaxIdleConnsPerHost: 5,
		MaxConnsPerHost:     5,
		IdleConnTimeout:     5 * time.Minute,
	},
	Timeout: 5 * time.Minute, //粗粒度 时间计算包括从连接(Dial)到读完response body
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}, //禁止重定向
}

var url ="http://172.16.129.124:14001/VIID/Subscribes"
func main() {
	subscribe := make([]*model.Subscribe, 0)
	sb :=&model.Subscribe{
		SubscribeID:           "",
		Title:                 "",
		SubscribeDetail:       "",
		ResourceURI:           "",
		ApplicantName:         "",
		ApplicantOrg:          "",
		BeginTime:             "",
		EndTime:               "",
		ReceiveAddr:           "",
		ReportInterval:        0,
		Reason:                "",
		OperateType:           0,
		SubscribeStatus:       0,
		SubscribeCancelOrg:    "",
		SubscribeCancelPerson: "",
		CancelTime:            "",
		CancelReason:          "",
		ResourceClass:         0,
		ResultFeatureDeclare:  0,
		ResultImageDeclare:    "",
		TabID:                 "",
	}
	subscribe = append(subscribe, sb)
	subscribeModel := &model.SubscribeModel{SubscribeListObject: &model.SubscribeListObject{SubscribeObject: subscribe}}

	request, header, i, err := Request(url, http.MethodPost, map[string]string{
		"User-Identify": "12354678797898987",
		"Content-Type":  "application/VIID+JSON",
	}, subscribeModel)

	fmt.Println(request,header,i,err)
}

//http请求
func Request(url, method string, headMap map[string]string, body interface{}) ([]byte, http.Header, int, error) {
	if !strings.HasPrefix(url, "http://") {
		url = `http://` + url
	}
	var bodyBytes []byte
	var resBytes []byte
	var header http.Header
	var code int
	if body != nil {
		bodyBytes, _ = jsoniter.Marshal(body)
		fmt.Println("======bodyBytes: ", utils.BytesString(bodyBytes))
	}
	err := Retry(func() error {
		req, err := http.NewRequest(method, url, bytes.NewReader(bodyBytes))
		if err != nil {
			return err
		}
		if headMap != nil {
			for key, value := range headMap {
				req.Header.Set(key, value)
			}
		}
		res, err := workerHttpClient.Do(req)
		if err != nil {
			return err
		}
		defer func() {
			err := res.Body.Close()
			if err != nil {
				fmt.Println("关闭res失败", err)
			}
		}()
		resBytes, err = ioutil.ReadAll(res.Body)
		header = res.Header
		code = res.StatusCode
		if err != nil {
			return err
		}
		//_ = res.Body.Close()
		//if !(res.StatusCode == http.StatusOK || res.StatusCode == http.StatusUnauthorized ||
		//	res.StatusCode == http.StatusCreated||res.StatusCode == http.StatusNoContent){
		//	return errors.New(string(resBytes))
		//}
		return nil
	}, 3, 3*time.Second)
	if err != nil {
		return resBytes, header, code, err
	}

	return resBytes, header, code, nil
}

/**
重试
*/
func Retry(f func() error, times int, space time.Duration) error {
	if times == 0 || times == 1 {
		return f()
	}
	var err error
	var i = 0
	for {
		err = f()
		if err == nil {
			return nil
		}
		i++
		if times > 1 && i >= times {
			break
		}
		time.Sleep(space)
	}
	return err
}