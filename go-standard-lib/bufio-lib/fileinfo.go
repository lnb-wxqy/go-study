package bufio_lib

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读写文件及复制文件
// 一、 读取文件
// 步骤：1.打开文件  2. 读取文件  3.关闭文件
var FileName = "E:\\lnb\\go\\src\\golang-study\\ex\\test.txt"
var _filepath = "E:\\lnb\\go\\src\\golang-study"

func ReadFile() {
	// 1. 打开文件

	f, e := os.Open(FileName)

	// 3.关闭文件
	defer f.Close()
	if e != nil {
		fmt.Println("打开文件有误：", e)
		return
	}

	// 2. 读/写
	//从file对应的文件中读取量最多len(bs)个数据，存入到bs切片中，n是实际读取到的量
	sli := make([]byte, 1024, 1024)
	n := -1
	for {
		n, e = f.Read(sli)
		if n == 0 || e == io.EOF {
			fmt.Println("文件读取完..")
			break
		}
		fmt.Println(string(sli[:n]))
	}

}

// 写入文件
// 步骤：1.打开或创建文件  2.写入文件  3.关闭文件
func WriteFile() {
	//1.打开文件
	f, e := os.OpenFile(FileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if e != nil {
		fmt.Println("打开文件有误：", e)
		return
	}

	//3.关闭文件
	defer f.Close()

	//2.写入文件
	n, e := f.Write([]byte("dfjadjajdflkj"))
	fmt.Println(n)
	fmt.Println(e)

	n, e = f.WriteString("进口量进口量积累级联")
	fmt.Println(n)
	fmt.Println(e)

}

//复制文件
// 示例1： 实现文件的拷贝，返回值是拷贝的总数量（字节），错误
func copyFile1(srcFile, dstFile string) (int, error) {
	rf, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	defer rf.Close()

	wf, err := os.OpenFile(dstFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer wf.Close()

	// 拷贝数据（读取数据）
	bs := make([]byte, 1024, 1024)
	n := -1
	total := 0
	for {
		n, err = rf.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("拷贝完毕!")
			break
		}
		total += n
		if err != nil {
			fmt.Println("err: ", err)
			return total, err
		} else {
			wf.Write(bs[:n])
		}
	}

	return total, nil
}

// 示例2:
func copyFile2(srcFile, dstFile string) (int64, error) {
	rf, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	defer rf.Close()
	wf, err := os.OpenFile(dstFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer wf.Close()
	w, err := io.Copy(wf, rf)
	return w, err
}

// ioutil 包
// ioutil包核心函数
func iouiltFunction() {
	//1.ReadFile()
	//读取⽂件中的所有的数据，返回读取的字节数组
	data, e := ioutil.ReadFile(FileName)
	if e != nil {
		fmt.Println("读取文件异常： ", e)
	} else {
		fmt.Println(string(data))
	}
	//2、WriteFile()
	//向指定⽂件写⼊数据，如果⽂件不存在，则创建⽂件，写入数据之前清空文件
	e = ioutil.WriteFile(FileName, []byte("ioutil包"), os.ModePerm)
	if e != nil {
		fmt.Println("写入文件异常: ", e)
	}
	//3、ReadDir()
	//读取⼀个⽬目录下的⼦子内容：⼦⽂件和⼦⽬目录，但是仅有一层
	dirs, e := ioutil.ReadDir(_filepath)
	if e != nil {
		fmt.Println("读取文件目录异常: ", e)
	}
	for _, dir := range dirs {
		fmt.Println("name: ", dir.Name(), "isDir: ", dir.IsDir())
	}
	//4、TempDir(）
	//在当前目录下，创建⼀个以指定字符串为前缀的临时⽂件夹，并返回⽂件夹路径
	name, e := ioutil.TempDir(_filepath, "pre")
	if e!=nil{
		fmt.Println("创建临时文件夹异常： ",e)
	}
	fmt.Println(name)
	//5、TempFile()
	//在当前⽬录下，创建⼀个以指定字符串串为前缀的⽂文件，并以读写模式打开
	//⽂件，并返回os.File指针对象
	f, e := ioutil.TempFile(_filepath, "pre")
	fmt.Println(f.Name())
}
