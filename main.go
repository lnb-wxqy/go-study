package main

import (
	"fmt"
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

func main() {

	fmt.Println(1 ^ 1)
	fmt.Println(2 ^ 2)
	//ui.Qoui_demo()
	//readMp4()
	//fmt.Println(time.Now().UnixNano() / 1e6)
	//
	////maxInt64 := math.MaxInt64
	////fmt.Println(maxInt64)
	//
	//stamp := utils.Str2Stamp("20210104123222")
	//fmt.Println(stamp)
	//_,b := EQUIPMENT_GBID_EQ["abc"]
	//_,b2 := EQUIPMENT_GBID_EQ[""]
	//fmt.Println(b)
	//fmt.Println(b2)
	//var d *Demo
	//d.test()
	//tTime :="20201111121222"
	//tmpBeginTime, _ := strconv.ParseInt(tTime, 10, 64)
	//fmt.Printf("v: type: %T\n",tmpBeginTime)
	//fmt.Println(reflect.TypeOf(tmpBeginTime))

	//var d Demo
	//fmt.Println(d)

	//s :="111111111111233"
	//ss := strings.Split(s, ";")
	//fmt.Println(ss)

	//t := time.Now().UnixNano() / 1e6
	//fmt.Println(t)
	//
	//stamp := utils.Str2Stamp("20201208112222")
	//fmt.Println(stamp)

	//pathName :="E:/work/项目支持/内蒙/造数据小组件/Desktop/"
	//
	//infos, _ := ioutil.ReadDir(pathName)
	//for _,f :=range infos{
	//	name := f.Name()
	//	fmt.Println(name)
	//}

	//data := []string{"one", "two", "three","four"}
	//for _, v := range data {
	//	//go func(s string) {
	//	//	fmt.Println(s)
	//	//}(v)
	//
	//	go func() {
	//		fmt.Println(v)
	//	}()
	//
	//}
	//time.Sleep(3 * time.Second)
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
