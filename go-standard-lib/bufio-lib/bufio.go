package bufio_lib

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// bufio.Reader的所有方法

func BufioMethod() {
	//type Reader
	//func NewReader(rd io.Reader) *Reader
	//func NewReaderSize(rd io.Reader, size int) *Reader
	//func (b *Reader) Reset(r io.Reader)
	//func (b *Reader) Buffered() int
	//func (b *Reader) Peek(n int) ([]byte, error)
	//func (b *Reader) Read(p []byte) (n int, err error)
	//func (b *Reader) ReadByte() (c byte, err error)
	//func (b *Reader) UnreadByte() error
	//func (b *Reader) ReadRune() (r rune, size int, err error)
	//func (b *Reader) UnreadRune() error
	//func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
	//func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
	//func (b *Reader) ReadBytes(delim byte) (line []byte, err error)
	//func (b *Reader) ReadString(delim byte) (line string, err error)
	//func (b *Reader) WriteTo(w io.Writer) (n int64, err error)
	//type Writer
	//func NewWriter(w io.Writer) *Writer
	//func NewWriterSize(w io.Writer, size int) *Writer
	//func (b *Writer) Reset(w io.Writer)
	//func (b *Writer) Buffered() int
	//func (b *Writer) Available() int
	//func (b *Writer) Write(p []byte) (nn int, err error)
	//func (b *Writer) WriteString(s string) (int, error)
	//func (b *Writer) WriteByte(c byte) error
	//func (b *Writer) WriteRune(r rune) (size int, err error)
	//func (b *Writer) Flush() error
	//func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)
	//type ReadWriter
	//func NewReadWriter(r *Reader, w *Writer) *ReadWriter
	//type SplitFunc
	//func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
	//func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
	//func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
	//func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
	// type Scanner
	//func NewScanner(r io.Reader) *Scanner
	//func (s *Scanner) Split(split SplitFunc)
	//func (s *Scanner) Scan() bool
	//func (s *Scanner) Bytes() []byte
	//func (s *Scanner) Text() string
	//func (s *Scanner) Err() error

	// 示例：使用bufio包创建文件并写入内容
	bufioWrite()
	bufioReader()

}

func bufioWrite() {
	//os.OpenFile("打开的文件"，os.O_CREATE（如果没有这个文件就创建这个文件）|os.O_WRONLY（以写的方式打开），文件权限)
	file, err := os.OpenFile(FileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("open file error： ", err)
		return
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		_, _ = w.WriteString("Hello bufio...\n")
	}
	_ = w.Flush()
}

func bufioReader() {
	file, err := os.Open(FileName)
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	for {
		// 读取字符创，这里是字节类型，要用单引号'\n'
		line, err := rd.ReadString('\n')
		// 判断文件是不是结尾，读取到文件结尾回返回一个io.EOF的错误
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}
		fmt.Printf("read string successful: %s\n", line)
	}
}
