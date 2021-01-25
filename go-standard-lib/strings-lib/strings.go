package strings_lib

import (
	"fmt"
	"strings"
	"unicode"
)

// strings包函数
//func EqualFold(s, t string) bool
//func HasPrefix(s, prefix string) bool
//func HasSuffix(s, suffix string) bool
//func Contains(s, substr string) bool
//func ContainsRune(s string, r rune) bool
//func ContainsAny(s, chars string) bool
//func Count(s, sep string) int
//func Index(s, sep string) int
//func IndexByte(s string, c byte) int
//func IndexRune(s string, r rune) int
//func IndexAny(s, chars string) int
//func IndexFunc(s string, f func(rune) bool) int
//func LastIndex(s, sep string) int
//func LastIndexAny(s, chars string) int
//func LastIndexFunc(s string, f func(rune) bool) int
//func Title(s string) string
//func ToLower(s string) string
//func ToLowerSpecial(_case unicode.SpecialCase, s string) string
//func ToUpper(s string) string
//func ToUpperSpecial(_case unicode.SpecialCase, s string) string
//func ToTitle(s string) string
//func ToTitleSpecial(_case unicode.SpecialCase, s string) string
//func Repeat(s string, count int) string
//func Replace(s, old, new string, n int) string
//func Map(mapping func(rune) rune, s string) string
//func Trim(s string, cutset string) string
//func TrimSpace(s string) string
//func TrimFunc(s string, f func(rune) bool) string
//func TrimLeft(s string, cutset string) string
//func TrimLeftFunc(s string, f func(rune) bool) string
//func TrimPrefix(s, prefix string) string
//func TrimRight(s string, cutset string) string
//func TrimRightFunc(s string, f func(rune) bool) string
//func TrimSuffix(s, suffix string) string
//func Fields(s string) []string
//func FieldsFunc(s string, f func(rune) bool) []string
//func Split(s, sep string) []string
//func SplitN(s, sep string, n int) []string
//func SplitAfter(s, sep string) []string
//func SplitAfterN(s, sep string, n int) []string
//func Join(a []string, sep string) string
//type Reader
//func NewReader(s string) *Reader
//func (r *Reader) Len() int
//func (r *Reader) Read(b []byte) (n int, err error)
//func (r *Reader) ReadByte() (b byte, err error)
//func (r *Reader) UnreadByte() error
//func (r *Reader) ReadRune() (ch rune, size int, err error)
//func (r *Reader) UnreadRune() error
//func (r *Reader) Seek(offset int64, whence int) (int64, error)
//func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
//func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
//type Replacer
//func NewReplacer(oldnew ...string) *Replacer
//func (r *Replacer) Replace(s string) string
//func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)

func StringsMethod() {
	// 判断两个utf-8编码字符串（将Unicode大写、小写、标题三种格式字符视为相同）是否相同
	//func EqualFold(s, t string) bool
	fmt.Println(strings.EqualFold("Go", "go")) // true

	// 判断s是否有前缀字符串prefix。
	//func HasPrefix(s, prefix string) bool
	fmt.Println(strings.HasPrefix("wxqy", "w")) //true
	fmt.Println(strings.HasPrefix("wxqy", "m")) //false

	//判断s是否有后缀字符串prefix。
	//func HasSuffix(s, suffix string) bool
	fmt.Println(strings.HasSuffix("wxqy", "w")) //false
	fmt.Println(strings.HasSuffix("wxqy", "y")) //true

	//判断字符串s是否包含子串substr。
	//func Contains(s, substr string) bool
	fmt.Println(strings.Contains("wxqy", "x"))   //true
	fmt.Println(strings.Contains("wxqy", "xq"))  //true
	fmt.Println(strings.Contains("wxqy", "xqw")) //false

	//判断字符串s是否包含utf-8码值r。
	//func ContainsRune(s string, r rune) bool
	fmt.Println(strings.ContainsRune("wxqy", 'x')) //true

	//判断字符串s是否包含字符串chars中的任一字符。
	//func ContainsAny(s, chars string) bool
	fmt.Println(strings.ContainsAny("wxqy", "xw")) //true
	fmt.Println(strings.ContainsAny("wxqy", "xm")) //true

	//返回字符串s中有几个不重复的sep子串。
	//func Count(s, sep string) int
	fmt.Println(strings.Count("wxxqyzz", "z"))     //2
	fmt.Println(strings.Count("wxxqyzz", "zz"))    //1
	fmt.Println(strings.Count("wxxqyzzxabx", "x")) //4

	//子串sep在字符串s中第一次出现的位置，不存在则返回-1。
	//func Index(s, sep string) int
	fmt.Println(strings.Index("wxqy", "q")) // 2
	fmt.Println(strings.Index("wxqy", "p")) // -1

	//字符c在s中第一次出现的位置，不存在则返回-1。
	//func IndexByte(s string, c byte) int
	fmt.Println(strings.IndexByte("wxqy", 'w')) //0

	//unicode码值r在s中第一次出现的位置，不存在则返回-1。
	//func IndexRune(s string, r rune) int
	fmt.Println(strings.IndexByte("wwxqy", 'w')) //0

	//字符串chars中的任一utf-8码值在s中第一次出现的位置，如果不存在或者chars为空字符串则返回-1。
	//func IndexAny(s, chars string) int

	//s中第一个满足函数f的位置i（该处的utf-8码值r满足f(r)==true），不存在则返回-1。
	//func IndexFunc(s string, f func(rune) bool) int
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c) // 是否是汉字
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))    //7
	fmt.Println(strings.IndexFunc("Hello, world", f)) //-1

	// 子串sep在字符串s中最后一次出现的位置，不存在则返回-1。
	//func LastIndex(s, sep string) int

	// 字符串chars中的任一utf-8码值在s中最后一次出现的位置，如不存在或者chars为空字符串则返回-1。
	//func LastIndexAny(s, chars string) int

	// s中最后一个满足函数f的unicode码值的位置i，不存在则返回-1。
	//func LastIndexFunc(s string, f func(rune) bool) int

	//返回s中每个单词的首字母都改为标题格式的字符串拷贝。
	//BUG: Title用于划分单词的规则不能很好的处理Unicode标点符号。
	//func Title(s string) string
	fmt.Println(strings.Title("life is short")) //Life Is Short

	//返回将所有字母都转为对应的小写版本的拷贝。
	//func ToLower(s string) string
	fmt.Println(strings.ToLower("LIFe is SHOrt")) //life is short

	//使用_case规定的字符映射，返回将所有字母都转为对应的小写版本的拷贝。
	//func ToLowerSpecial(_case unicode.SpecialCase, s string) string

	//返回将所有字母都转为对应的大写版本的拷贝。
	//func ToUpper(s string) string
	fmt.Println(strings.ToUpper("LIFe is SHOrt")) // LIFE IS SHORT

	//返回将所有字母都转为对应的大写版本的拷贝。
	//func ToUpperSpecial(_case unicode.SpecialCase, s string) string

	//返回将所有字母都转为对应的标题版本的拷贝。
	//func ToTitle(s string) string
	fmt.Println(strings.ToTitle("wxqy qq"))  // WXQY QQ
	fmt.Println(strings.ToTitle("x pai 06")) // X PAI 06

	//使用_case规定的字符映射，返回将所有字母都转为对应的标题版本的拷贝。
	//func ToTitleSpecial(_case unicode.SpecialCase, s string) string

	//返回count个s串联的字符串。
	//func Repeat(s string, count int) string
	fmt.Println("ba" + strings.Repeat("na", 2)) // banana

	// 返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
	//func Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace("oink oink oink", "k", "ok", 2))      //oinok oinok oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // moo moo moo

	//将s的每一个unicode码值r都替换为mapping(r)，返回这些新码值组成的字符串拷贝。
	//如果mapping返回一个负值，将会丢弃该码值而不会被替换。（返回值中对应位置将没有码值）
	//func Map(mapping func(rune) rune, s string) string
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brilling and the slithy gopher...")) //'Gjnf oevyyvat naq gur fyvgul tbcure...

	//返回将s前后端所有cutset包含的utf-8码值都去掉的字符串。
	//func Trim(s string, cutset string) string
	fmt.Printf("【%q】\n", strings.Trim(" !!! wxqy! wxqy !!!!", "!"))  // " !!! wxqy! wxqy "】
	fmt.Printf("【%q】\n", strings.Trim(" !!! wxqy! wxqy !!!!", "! ")) // "wxqy! wxqy"】

	//返回将s前后端所有空白（unicode.IsSpace指定）都去掉的字符串。
	//func TrimSpace(s string) string
	fmt.Println(strings.TrimSpace("\t\n think cherry \n\t\r\n")) //think cherry

	//返回将s前后端所有满足f的unicode码值都去掉的字符串。
	//func TrimFunc(s string, f func(rune) bool) string

	// 返回将s前端所有cutset包含的utf-8码值都去掉的字符串。
	//func TrimLeft(s string, cutset string) string

	//返回将s前端所有满足f的unicode码值都去掉的字符串。
	//func TrimLeftFunc(s string, f func(rune) bool) string

	// 返回去除s可能的前缀prefix的字符串。
	//func TrimPrefix(s, prefix string) string

	// 返回将s后端所有cutset包含的utf-8码值都去掉的字符串。
	//func TrimRight(s string, cutset string) string

	// 返回将s后端所有满足f的unicode码值都去掉的字符串。
	//func TrimRightFunc(s string, f func(rune) bool) string

	// 返回去除s可能的后缀suffix的字符串。
	//func TrimSuffix(s, suffix string) string

	// 返回将字符串按照空白（unicode.IsSpace确定，可以是一到多个连续的空白字符）分割的多个字符串。
	// 如果字符串全部是空白或者是空字符串的话，会返回空切片。
	//func Fields(s string) []string
	fmt.Printf("Fields are: %q\n", strings.Fields(" boo bar baz")) // ["boo" "bar" "baz"]

	//类似Fields，但使用函数f来确定分割符（满足f的unicode码值）。
	// 如果字符串全部是分隔符或者是空字符串的话， 会返回空切片。
	//func FieldsFunc(s string, f func(rune) bool) []string

	// 用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，
	// 即使两个sep相邻，也会进行两次切割）。
	// 如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。
	//func Split(s, sep string) []string
	fmt.Println(strings.Split("god,is,girl", " ")) // [god,is,girl]

	//用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，
	// 即使两个sep相邻，也会进行两次切割）。
	//如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。参数n决定返回的切片的数目：
	//n > 0 : 返回的切片最多n个子字符串；最后一个子字符串包含未进行切割的部分。
	//n == 0: 返回nil
	//n < 0 : 返回所有的子字符串组成的切片
	//func SplitN(s, sep string, n int) []string
	fmt.Println(strings.SplitN("god,is,girl,girl,girl", ",", 3)) //[god is girl,girl,girl]
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)

	//用从s中出现的sep后面切断的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，
	// 即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。
	//func SplitAfter(s, sep string) []string

	//用从s中出现的sep后面切断的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，
	// 即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。参数n决定返回的切片的数目：
	//
	//n > 0 : 返回的切片最多n个子字符串；最后一个子字符串包含未进行切割的部分。
	//n == 0: 返回nil
	//n < 0 : 返回所有的子字符串组成的切
	//func SplitAfterN(s, sep string, n int) []string

	//将一系列字符串连接为一个字符串，之间用sep来分隔。
	//func Join(a []string, sep string) string
	strJoin := strings.Join([]string{"dog", "cat", "tiger"}, "|") //dog|cat|tiger
	fmt.Println(strJoin)

	//Reader类型通过从一个字符串读取数据，实现了io.Reader、io.Seeker、io.ReaderAt、io.WriterTo、io.ByteScanner、
	// io.RuneScanner接口。
	//type Reader
	//type Reader struct {
	//	内含隐藏或非导出字段
	//}

	//NewReader创建一个从s读取数据的Reader。本函数类似bytes.NewBufferString，但是更有效率，且为只读的。
	//func NewReader(s string) *Reader

	//Len返回r包含的字符串还没有被读取的部分。
	//func (r *Reader) Len() int

	//func (r *Reader) Read(b []byte) (n int, err error)
	//func (r *Reader) ReadByte() (b byte, err error)
	//func (r *Reader) UnreadByte() error
	//func (r *Reader) ReadRune() (ch rune, size int, err error)
	//func (r *Reader) UnreadRune() error
	//Seek实现了io.Seeker接口。
	//func (r *Reader) Seek(offset int64, whence int) (int64, error)
	//func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
	//WriteTo实现了io.WriterTo接口。
	//func (r *Reader) WriteTo(w io.Writer) (n int64, err error)

	//Replacer类型进行一系列字符串的替换
	//type Replacer
	//type Replacer struct {
	//
	//}

	//使用提供的多组old、new字符串对创建并返回一个*Replacer。替换是依次进行的，匹配时不会重叠。
	//func NewReplacer(oldnew ...string) *Replacer
	r := strings.NewReplacer(",", "M", ".", "N", "。", "L", "!", "x")
	fmt.Println(r.Replace("老夫聊发少年狂,不私聊,不私聊!嘻嘻。"))

	//Replace返回s的所有替换进行完后的拷贝。
	//func (r *Replacer) Replace(s string) string

	// WriteString向w中写入s的所有替换进行完后的拷贝。
	//func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)

}
