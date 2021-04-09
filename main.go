package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-study/utils"
	"io"
	"os"
)

type Demo struct {
}

var EQUIPMENT_GBID_EQ = make(map[string]string)

func (d *Demo) test() {
	fmt.Println("d: ", d)
	fmt.Println(d == nil)
}

func GetFile() []byte {
	url := "/api/dag/v2/resource/attachment/2021/3/24/6dc9497cf300403796920300ed137b75.json"

	centerIp := "172.16.129.124"
	url = "http://" + centerIp + ":8080" + url

	fmt.Println("url: ", url)

	data, err := utils.DownLoadFile(url)
	if err != nil {
		fmt.Println("download config file err. ", err)
		return nil
	}
	return data
}

var deviceCodeMapping = make(map[string]string, 0)

func initMapping(data []byte) error {
	//deviceCodeMapping := make(map[string]string, 0)
	if data != nil && len(data) > 0 {
		//读取的数据为json格式，需要进行解码
		err := json.Unmarshal(data, &deviceCodeMapping)
		if err != nil {
			fmt.Println("unmarshal failed！", err)
			return err
		}
	}

	fmt.Println("d.deviceCodeMapping: ", deviceCodeMapping)
	return errors.New("未开启映射内容，映射内容未空.")
}

func AB() (dd *Demo) {
	fmt.Println(dd)
	return nil
}

type NetBarFace struct {
	//id->网吧编码->网吧名称->行政区划->地址->上网人员姓名->证件类型->证件号码->网监->国籍->上机时间->ip->下机时间->区划编码
	id               string
	netBarDeviceID   string
	netBarName       string
	orgCode          string
	addr             string
	name             string
	zhengjianleixing string
	zjhm             string
	wangjian         string
	guoji            string
	shangjishijian   string
	ip               string
	xiajishijian     string
	quhuabianma      string
}

func main() {

	fmt.Println(1^0)

	//timer := time.NewTimer(5 * time.Second)
	//defer timer.Stop()
	//
	//for {
	//	select {
	//	case <-timer.C:
	//		fmt.Printf("现在是：%d, 定时任务结束！", time.Now().Unix())
	//		fmt.Println()
	//	}
	//}

	//str1 :="20210401"
	//str2 :="20210402"
	//
	//fmt.Println(str1<str2)

	//now := time.Now()
	//strTime := utils.Time2StrF(time.Now(), utils.DateFormatYMD)
	//fmt.Println(strTime)
	//fmt.Println(time.Since(now))

	//str := "37C9787003\t37020001000336\t网缘电竞馆（东海）\t320700\t东海县宏达花园门面房107号\t陈标\t111\t320722199709080039\t连云港网监\tCN\t20201117205924\t\t\t0000\n"
	//
	//st := strings.Split(str, "\t")
	////fmt.Println("st: ",st)
	//for i, s := range st {
	//	fmt.Printf("[%d]：[%s]", i, s)
	//	fmt.Println()
	//}
	//fmt.Println("st.len: ", len(st))
	//
	//var nf NetBarFace
	//err := json.Unmarshal([]byte(str), &nf)
	//fmt.Println("err: ",err)
	//fmt.Println("nf: ",nf)

	//nums := []int{2, 7, 11, 15}
	//target := 9
	//result := leetcode.TwoSum(nums, target)
	//fmt.Println("result: ", result)

	//nums2 := []int{2, 8, 11, 15,7,1,3,6}
	//target2 := 9
	//result2 := leetcode.TwoSum(nums2, target2)
	//fmt.Println("result2: ", result2)
	//AB()
	//buf :=make([]byte,-1)
	//fmt.Println(buf)

	//fileData := GetFile()
	//initMapping(fileData)

	//str := utils.Stamp2Str(time.Now().UnixNano()/1e6, utils.DateFormatCNSSS)
	//fmt.Println("str: ",str)
	//
	//url :="http://10.24.113.117:8082/ctm01facebigpic/v1/pic/getPic?picUrl=aHR0cDovLzEzLjIyNS4yLjU2OjYxMjAvcGljPzJkZDk0NHo2OC09czk1OTE2ODg2MzY4MWUtPXQxaTdtKj1wMXA4aT1kMXMqaTdkMTQqOGQ2PSo0YjhpNTVjYzM5MDdhYzNkZTEtYTBiMzE3LTc1Mmk1NjkqZTZlN2kxNT0="
	//
	//
	//subUrl := strings.Replace(url, "http://", "", 1)
	//i := strings.Index(subUrl, "/")
	//newUrl := "http://172.16.129.124:6551"+ subUrl[i:]
	//fmt.Println(newUrl)

	//s :="01"
	//i, _ := strconv.Atoi(s)
	//fmt.Println("i: ",i)
	//
	//t := "2021-02-02 09:38:51"
	//t2 := utils.Str2Stamp(t)
	//
	//t3 := "2021-02-04 02:08:09"
	//t4 := utils.Str2Stamp(t3)
	//fmt.Println(t2)
	//fmt.Println(t4)

	// 600000
	// 300000
	//url :="http://10.24.113.117:8082/ctm01facebigpic/v1/pic/getPic?picUrl=aHR0cDovLzEzLjIyNS4yLjU2OjYxMjAvcGljPzJkZDk0NHo2OC09czk1OTE2ODg2MzY4MWUtPXQxaTdtKj1wMXA4aT1kMXMqaTdkMTQqOGQ2PSo0YjhpNTVjYzM5MDdhYzNkZTEtYTBiMzE3LTc1Mmk1NjkqZTZlN2kxNT0="
	//url2 :="http://10.24.113.117:8082/ctm01facebigpic/v1/pic/getPic?picUrl=aHR0cDovLzEzLjE3MS4xLjE3Nzo4MTAxL2Nkcy92MS9maWxlL2NvbW1vbio4NjMxMzg2NzdfOTQ2NTMvVmFzUmVhbFZpZGVvL1JlYWxWaWRlby1ZZEZCUTZDV2FxRFkvMjAyMS0wMS0yOC92YXNAUmVhbFZpZGVvLVlkRkJRNkNXYXFEWUA5OTk0OUAxNjExNzY4MDY1LmpwZz9kZXY9Y2R2c2VydmVyLTAxJmZpZD04MTk5OTMtMTI4LTQ3MDMyMDAwMTYtMzlCOUItMzlGNjY="
	//fmt.Println(len(url))
	//fmt.Println(len(url2))
	////
	////fmt.Println(1 ^ 1)
	////fmt.Println(2 ^ 2)
	//////ui.Qoui_demo()
	////readMp4()
	////fmt.Println(time.Now().UnixNano() / 1e6)
	////
	//////maxInt64 := math.MaxInt64
	//////fmt.Println(maxInt64)
	////
	////stamp := utils.Str2Stamp("20210104123222")
	////fmt.Println(stamp)
	////_,b := EQUIPMENT_GBID_EQ["abc"]
	////_,b2 := EQUIPMENT_GBID_EQ[""]
	////fmt.Println(b)
	////fmt.Println(b2)
	////var d *Demo
	////d.test()
	////tTime :="20201111121222"
	////tmpBeginTime, _ := strconv.ParseInt(tTime, 10, 64)
	////fmt.Printf("v: type: %T\n",tmpBeginTime)
	////fmt.Println(reflect.TypeOf(tmpBeginTime))
	//
	////var d Demo
	////fmt.Println(d)
	//
	////s :="111111111111233"
	////ss := strings.Split(s, ";")
	////fmt.Println(ss)
	//
	////t := time.Now().UnixNano() / 1e6
	////fmt.Println(t)
	////
	////stamp := utils.Str2Stamp("20201208112222")
	////fmt.Println(stamp)
	//
	////pathName :="E:/work/项目支持/内蒙/造数据小组件/Desktop/"
	////
	////infos, _ := ioutil.ReadDir(pathName)
	////for _,f :=range infos{
	////	name := f.Name()
	////	fmt.Println(name)
	////}
	//
	////data := []string{"one", "two", "three","four"}
	////for _, v := range data {
	////	//go func(s string) {
	////	//	fmt.Println(s)
	////	//}(v)
	////
	////	go func() {
	////		fmt.Println(v)
	////	}()
	////
	////}
	////time.Sleep(3 * time.Second)
}

func readMp4() {
	mp4path := "E:\\lnb\\java基础\\go\\4天掌握GO语言密码学-用实践验证理论视频\\day1\\01-课程介绍.mp4"
	f, err := os.OpenFile(mp4path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Print("open file error: ", err)
		return
	}
	defer f.Close()
	bs := make([]byte, 1024*1024*100, 1024*1024*100)
	n := 1
	for {
		n, err = f.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("read finish!")
			break
		}
		if err != nil {
			fmt.Println("read file error: ", err)
			break
		} else {
			fmt.Println("1111111111111")
		}

	}
}
