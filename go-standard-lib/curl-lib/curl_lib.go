package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

const URL_REGIST = "/VIID/System/Register"

func main() {

	url :="http://32.73.148.130:7320/storage/v1/image?cluster_id=WJSPW_MTS_1587290067&uri_base64=a3Y6Ly9rdi1pbWFnZS1mbG93X2ZhY2UvNDhlNDYwMTkwNzQ5MDAwMDFjMDAxOWNhYTgyNV8xNjA0MDYxNzc3LWRkYmNjYjYwLTZhM2EtNDU5MC05ODkwLTgxZmM0MzA2Y2ExMA=="
	url = base64.URLEncoding.EncodeToString([]byte(url))+"&encode=1"
	fmt.Println(url)

	url2 :="aHR0cDovLzMyLjczLjE0OC4xMzA6NzMyMC9zdG9yYWdlL3YxL2ltYWdlP2NsdXN0ZXJfaWQ9V0pTUFdfTVRTXzE1ODcyOTAwNjcmdXJpX2Jhc2U2ND1hM1k2THk5cmRpMXBiV0ZuWlMxbWJHOTNYMlpoWTJVdk5EaGxORFl3TVRrd056UTVNREF3TURGak1EQXhPV05oWVRneU5WOHhOakEwTURZeE56YzNMV1JrWW1OallqWXdMVFpoTTJFdE5EVTVNQzA1T0Rrd0xUZ3habU0wTXpBMlkyRXhNQT09&encode=1"

	fmt.Println(url==url2)
	//
	//t :="2021-02-02 09:38:51"
	//stamp := utils.Str2Stamp(t)
	//fmt.Println(stamp)
	//
	//t2 :="1612229931000"
	//stamp = utils.Str2Stamp(t2)
	//fmt.Println(stamp)


	//realm :=" real viid api"
	//trim := strings.Trim(realm, " ")
	//fmt.Println(trim)

	//, qop="auth", nonce="MTYxMjI0MTIzNjgxMDpiYThjMmJkZWU5NjhjZjRjZmUyNGIyYzMyYjRjYmM2Ng=="
	//param := strings.ReplaceAll(strings.ReplaceAll(s, `"`, ""), `'`, "")
	//fmt.Println("param: ", param)
	//
	//var authorizationMap = make(map[string]string)
	//for _, argument := range strings.Split(s, ",") {
	//	//parm := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(argument, `"`, ""), `'`, ""), " ", "")
	//	// TODO 测试
	//	parm := strings.ReplaceAll(strings.ReplaceAll(argument, `"`, ""), `'`, "")
	//	if strings.Index(parm, "=") == -1 {
	//		fmt.Println("header 参数错误", parm)
	//		continue
	//	}
	//	authorizationMap[parm[:strings.Index(parm, "=")]] = parm[strings.Index(parm, "=")+1:]
	//}
	//
	//fmt.Println("authorizationMap: ", authorizationMap)
	//
	//// HA1:  35ed0a6073d36c979aafe7d5ad73fa25
	//HA1 := utils.MD5("admin" + ":" + "real viid api" + ":" + "admin123")
	//fmt.Println("HA1: ", HA1)
	//
	////e981db7569f065fda647d8236af6f901
	//HA2 := utils.MD5("POST" + ":" + URL_REGIST)
	//fmt.Println("HA2: ", HA2)
	//
	//sss :="realm viid 123"
	//fmt.Println(strings.Contains(sss,"realm"))

	//	args := `
	// "http://172.16.129.124:14000/VIID/System/Register" -H "accept:application/viid+json;charset=UTF-8" -H "Content-Type: application/json" -d "{ "RegisterObject": { "DeviceID": "1234567890"}}" --digest -u admin:admin123
	//`
	//	curl := exec.Command("curl", "-X", "POST",args)
	//	out, err := curl.Output()
	//	if err != nil {
	//		fmt.Println("erorr", err)
	//		return
	//	}
	//
	//	fmt.Println(utils.BytesString(out))

	//engine := gin.Default()
	//
	////注册
	//engine.POST(URL_REGIST, func(c *gin.Context) {
	//	regist(c)
	//})
	//
	//engine.Run(":8899")
}

func regist(c *gin.Context) {

	data := make([]interface{}, 0)
	_ = jsoniter.NewDecoder(c.Request.Body).Decode(data)
	header := c.Request.Header
	fmt.Println(data)
	fmt.Println(header)
}
