

# Go语言学习笔记

#### 一、变量与常量

##### （一）变量概念

1. 变量时计算机中存储数据的抽象概念。
2. Go语言的变量名有字母、数字、下划线组成，首个字符不能为数字，变量命名建议使用`camelCasing`格式。

##### （二）变量声明

1. var 变量名 变量类型

例1：

```
var s string
```

例2：批量声明

```go
var (
	i int
	s string
	f []float32
	fn func() bool
)
```

2. var 变量名 变量类型 = 表达式
3. var 变量名 表达式
4. 变量名 := 表达式（**短变量声明，建议使用这种命名方式**）

```go
s := "golang"
```

- 变量类型由编译器自动推导；只能用于局部变量的声明，不能用于全局变量声明。

##### （三）变量多重赋值和匿名变量

1. 多重赋值

```go
i, s, f := 7, "go", 12.32
fmt.Println(i, s, f)
```

2. 变量交换

```go
m :=12
n :=22
fmt.Println("m,n交换前的值分别为：",m,n) // 输出结果为：12,22
m,n=n,m
fmt.Println("m,n交换后的值分别为：",m,n) // 输出结果为：22,12
```

3. 匿名变量

- 匿名变量不占用命名空间，不分配内存，用 "_" 代替，对应返回值用不上的可以使用匿名变量。

```go
_ = "匿名变量"
```

#### 二、数据类型

- 基本数据类型：整形、浮点型、布尔型、字符串、字节byte和rune

- 复合数据类型：指针（pointer）、数组（array）、切片（slice）、字典（map）、函数（function）、结构体（struct）、通道（channel）

##### （一）整型（分为两大类）

1. 按长度分：`int8、int、int16、int32、int64`
2. 无符号整形：`uint8、uint、uint16、uint32、uint64`
3. 其中uint8就是byte类型，int16对应java中的short类型

##### （二）浮点型

- Go语言支持4中浮点型数：float32、float64、complex64（32位实数和虚数）、complex128 (64位实数和虚数)

##### （三）字符串和字符

1. 字符串在Go语言总是以基本数据类型出现的。
2. 字符串中的每一个元素叫做 "字符"，定义字符使用单引号。GO语言字符有两种
3. byte类型：其实是uint8。代表一个ASCII码的一个字符。
4. rune类型：其实是int32。代表一个utf-8字符。当需要处理中文等Unicode字符集时会用到rune类型

```go
var b byte = "b"
var br rune = '中'
```

#### 三、数据类型转换

##### （一）转换格式

T(表达式)

T表示要转换的类型

```go
i :=479
fmt.Printf("%T\n",i) // 输出结果： int

s :=string(i)
fmt.Printf("%T\n",s) // 输出结果：string
```



#### 四、常量

##### （一）声明方式

1. const 标识符 [类型] = 值
2. 可以省略类型说明符[类型]，编译器会根据值自动推导出类型

```go
const B Sting
const C = "值"
```

3. 常量定义后未被使用，不会在编译时出错。

##### （二）常量用于枚举（常用组）

例如：

```go
const {
    Female = 0  // 位置性别
    Male = 1    // 女星
    Unkwon 2  未知
}
```

1. 常量组中如果不指定类型和初始值，则与上一行非空常量的值相同。

   ```go
   const {
   	a = 10
   	b
   	c
   }
   fmt.println(a,b,c) //输出结果为： 10,10,10 
   ```

   

##### （三）iota 

1. iota,特殊常量值，是一个系统定义的可以被修改的常量值。iota只能用在常量赋值中。

2. 在每一个const关键字出现时，被重置为0，然后没出现一个常量，iota所代表的数值会自动增加1。

3. iota可以被用作枚举值：

   ```go
   const (
   	a = iota
   	b = iota
   	c = iota
   )
   fmt.println(a,b,c) //输出结果为： 0,1,2
   ```

   - 第一个`iota`的值等于0，每当 `iota`在新的一行使用时，它的值就指定加1，所以 `a=0,b=1,c=2`

4. 常量组中如果不指定类型和初始值，则与上一行非空常量的值相同，所以上述的枚举可以简写为：

   ```
   const (
   	a = iota
   	b 
   	c 
   )
   fmt.println(a,b,c) //输出结果为： 0,1,2
   ```

   - 示例：

   ```go
   const (
   	i = 1 << iota
   	j = 3 << iota
   	k
   	l
   )
   
   func main() {
   	fmt.Println(i, j, k, l) // 输出结果：1 6 12 24
   }
   ```

#### 五、运算符

Go语言内置运算符有：

- 算术运算符
- 关系运算符
- 逻辑运算符
- 位运算符
- 赋值运算符
- 其他运算符

##### （一）算术运算符

1. 如下表，假设A的值为1，B的值为2

   | 运算符 | 描述 |       实例        |
   | :----: | :--: | :---------------: |
   |   +    |  加  | A + B 输出结果 3  |
   |   -    |  减  | A - B 输出结果 -1 |
   |   *    |  乘  | A * B 输出结果 2  |
   |   /    |  除  | A / B 输出结果 0  |
   |   %    | 取余 |  A%B 输出结果 1   |
   |   ++   | 自增 |  A++ 输出结果 2   |
   |   --   | 自减 |  A-- 输出结果 0   |

   - 代码示例

     ```go
     	var a int =1
     	var b int =2
     	var c int
     
     	c = a + b
     	fmt.Printf("第一⾏ - c 的值为 %d\n", c )
     	c = a - b
     	fmt.Printf("第二行 - c 的值为 %d\n", c )
     	c = a * b
     	fmt.Printf("第三行 - c 的值为 %d\n", c )
     	c = a / b
     	fmt.Printf("第四行 - c 的值为 %d\n", c )
     	c = a % b
     	fmt.Printf("第五行 - c 的值为 %d\n", c )
     	a++
     	fmt.Printf("第六行 - a 的值为 %d\n", a )
     	a=1 // 为了了⽅方便便测试，a 这⾥里里重新赋值为 1
     	a--
     	fmt.Printf("第七行 - a 的值为 %d\n", a )
     ```

     

   - 输出结果

     ```go
     	第一⾏ - c 的值为 3
         第二行 - c 的值为 -1
         第三行 - c 的值为 2
         第四行 - c 的值为 0
         第五行 - c 的值为 1
         第六行 - a 的值为 2
         第七行 - a 的值为 0
     ```

##### （二）关系运算符

| 运算符 |                            描述                             |
| :----: | :---------------------------------------------------------: |
|   ==   |     判断两个值是否相等，如果相等返回True，否则返回false     |
|   !=   |   判断两个值是否不相等，如果不相等返回True，否则返回false   |
|   >    |   判断左边值是否大于右边值，如果是返回True，否则返回false   |
|   <    |   判断左边值是否小于右边值，如果是返回True，否则返回false   |
|   >=   | 判断左边值是否大于等于右边值，如果是返回True，否则返回false |
|   <=   | 判断左边值是否小于等于右边值，如果是返回True，否则返回false |

- 代码示例

  ```go
  	var a int = 1
  	var b int = 2
  	var c bool
  
  	c = a == b
  	fmt.Println(" a==b 的结果为：", c)
  	c = a != b
  	fmt.Println(" a!=b 的结果为：", c)
  	c = a > b
  	fmt.Println(" a>b 的结果为：", c)
  	c = a < b
  	fmt.Println(" a<b 的结果为：", c)
  	c = a >= b
  	fmt.Println(" a>=b 的结果为：", c)
  	c = a <= b
  	fmt.Println(" a<=b 的结果为：", c)
  ```

  

- 输出结果

  ```go
  	 a==b 的结果为： false
       a!=b 的结果为： true
       a>b 的结果为： false
       a<b 的结果为： true
       a>=b 的结果为： false
       a<=b 的结果为： true
  ```

##### （三）逻辑运算符

1. 如下表，假定 A 的值为 True，B 的值为 False

   | 运算符 | 描述                                                         |
   | ------ | ------------------------------------------------------------ |
   | &&     | 逻辑AND运算符。如果两边的操作数都为True，则结果为True，否则为False |
   | \|\|   | 逻辑OR运算符。如果两边的操作数有一个True，则结果为True，否则为False |
   | !      | 逻辑NOT运算符。如果条件为True，则逻辑 NOT 条件为 False，否则为 True |

   - 代码示例

     ```go
     a := true
     b := false
     
     fmt.Println("a && b的结果为：", a && b)
     fmt.Println("a || b的结果为：", a || b)
     fmt.Println("!a    的结果为：", !a)
     ```

   - 输出结果

     ```go
     a && b的结果为： false
     a || b的结果为： true
     !a    的结果为： false
     ```

##### （四）位运算符

1. 概述

   - 位运算符对整数

     在内存中二进制进行操作。

   - 位运算符比一般的算术运算符速度要快，而且可以实现一些算术运算符不能实现的功能。如果要开发高效率程序，位运算符是必不可少的。位运算符用来对二进制进行操作，包括：按位与（&）、按位或（|），按位异或（^）、按位左移（<<）、按位右移（>>)

   - 假定 A = 60; B = 13; 其二进制数转换为：
     A = 0011 1100
     B = 0000 1101
     …………………………

     A&B = 0000 1100
     A|B  =  0011 1101
     A^B  = 0011 0001

2. Go语言支持的位运算符如下表所示：假定 A 为 60，B 为13：

   | 位运算符 |                             描述                             |
   | :------: | :----------------------------------------------------------: |
   |    &     |                     按位与，双目运算符。                     |
   |    \|    |                     按位或，双目运算符。                     |
   |    ^     |   按位异或，双目运算符。当"^"两边的二进制位不同时，结果为1   |
   |    <<    | 左移，双目运算符。左移n位就是乘以2的n次方。"<<"左边的运算数的各二进制位全部左移若干位。 |
   |    >>    | 右移，双目原始符。右移n位就是除以2的n次方。">>"左边的运算数的各二进制位全部右移若干位。 |

   - 代码示例

     ```go
     	var a uint = 60 /* 60 = 0011 1100 */
     	var b uint = 13 /* 13 = 0000 1101 */
     	var c uint = 0
     	c = a & b /* 12 = 0000 1100 */
     	fmt.Printf("a & b的结果为： %d\n", c )
     	c = a | b /* 61 = 0011 1101 */
     	fmt.Printf("a | b的结果为： %d\n", c )
     	c = a ^ b /* 49 = 0011 0001 */
     	fmt.Printf("a ^ b的结果为： %d\n", c )
     	c = a << 2 /* 240 = 1111 0000 */
     	fmt.Printf("a << b的结果为： %d\n", c )
     	c = a >> 2 /* 15 = 0000 1111 */
     	fmt.Printf("a >> b的结果为： %d\n", c )
     ```

   - 输出结果

     ```go
     a & b的结果为： 12
     a | b的结果为： 61
     a ^ b的结果为： 49
     a << b的结果为： 240
     a >> b的结果为： 15
     ```

##### （五）赋值运算符

- Go 语言赋值运算符如下表

  | 运算符 | 描述                           |
  | :----: | :----------------------------- |
  |   =    | 将一个表达式的值赋值给一个左值 |
  |   +=   | 相加后赋值                     |
  |   -=   | 相减后赋值                     |
  |   *=   | 相乘后赋值                     |
  |   /=   | 相除后赋值                     |
  |   %=   | 求余后赋值                     |
  |  <<=   | 左移后赋值                     |
  |  >>=   | 右移后赋值                     |
  |   &=   | 按位与后赋值                   |
  |   ^=   | 按位异或后赋值                 |
  |  \|=   | 按位或后赋值                   |

  - 代码示例

    ```go
    a := 7
    var c int
    
    c = a
    fmt.Printf(" = 运算符实例，c 值为： %d\n", c)
    c += a
    fmt.Printf(" += 运算符实例，c 值为： %d\n", c)
    c -= a
    fmt.Printf(" -= 运算符实例，c 值为： %d\n", c)
    c *= a
    fmt.Printf(" *= 运算符实例，c 值为： %d\n", c)
    c /= a
    fmt.Printf(" /= 运算符实例，c 值为： %d\n", c)
    c <<= a
    fmt.Printf(" <<= 运算符实例，c 值为： %d\n", c)
    c >>= a
    fmt.Printf(" >>= 运算符实例，c 值为： %d\n", c)
    c &= a
    fmt.Printf(" &= 运算符实例，c 值为： %d\n", c)
    c ^= a
    fmt.Printf(" ^= 运算符实例，c 值为： %d\n", c)
    c |= a
    fmt.Printf(" |= 运算符实例，c 值为： %d\n", c)
    ```

  - 输出结果

    ```go
     = 运算符实例，c 值为： 7
     += 运算符实例，c 值为： 14
     -= 运算符实例，c 值为： 7
     *= 运算符实例，c 值为： 49
     /= 运算符实例，c 值为： 7
     <<= 运算符实例，c 值为： 896
     >>= 运算符实例，c 值为： 7
     &= 运算符实例，c 值为： 7
     ^= 运算符实例，c 值为： 0
     |= 运算符实例，c 值为： 7
    ```

##### （六）其他运算符

- Go 语言的其他运算符

  | 运算符 | 描述             |                          |
  | ------ | ---------------- | ------------------------ |
  | &      | 返回变量存储地址 | &a; 将给出变量的实际地址 |
  | *      | 指针变量         | *a; 是一个支持变量       |

  - 代码示例

    ```go
    a := 4
    var b int32
    var c float32
    var ptr *int
    /* 运算符实例例 */
    fmt.Printf("a 变量类型为 = %T\n", a)
    fmt.Printf("b 变量类型为 = %T\n", b)
    fmt.Printf("c 变量类型为 = %T\n", c)
    /* & 和 * 运算符实例例 */
    ptr = &a /* 'ptr' 包含了了 'a' 变量量的地址 */
    fmt.Printf("a 的值为 %d\n", a)
    fmt.Printf("&a 的值为 %v\n", ptr)
    fmt.Printf("*ptr 为 %d\n", *ptr)
    ```

    

  - 输出结果

    ```go
    a 变量类型为 = int
    b 变量类型为 = int32
    c 变量类型为 = float32
    a 的值为 4
    &a 的值为 0xc00000a200
    *ptr 为 4
    ```

##### （七）运算符优先级

- 运算符优先级如下表，右上至下代表优先级由高到低：

  | 优先级 | 运算符           |
  | ------ | ---------------- |
  | 5      | * / % << >> & &^ |
  | 4      | + - \| ^         |
  | 3      | == != , <= > >=  |
  | 2      | &&               |
  | 1      | \|\|             |

#### 六、Go语言流程控制

##### （一）Go语言流程控制概述

1. 流程控制是每种编程语言控制逻辑走向和执行次序的重要部分，流程控制是一们语言的经脉；
2. 流程控制有条件判断语句、条件分支语句及循环语句；
3. Go语言的基本流程控制语句包括：
   - if 条件判断语句
   - switch 分支语句
   - for 循环语句
   - goto 跳转语句及break和continue循环控制语句

##### （二）if 条件判断语句

1. if 语句语法

   if 布尔表达式 {

   ​	 /* 在布尔表达式为 true 时执行*/

   }

2. if...else语句语法

   if 布尔表达式 {

   ​	 /* 在布尔表达式为 true 时执行*/

   }else {

   ​	/* 在布尔表达式 false 时执行 */

   }

3. if...else if ... else 语句的语法如下：

   if 布尔表达式 {

   ​	 /* 在布尔表达式为 true 时执行*/

   }else if{

   ​	/*在布尔表达式 true 时执行*/

   }else {

   ​	/* 在布尔表达式 false 时执行 */

   }

   

4. if 语句中的注意细节

   - 不需要使用括号将条件包含起来
   - 大括号{}必须存在，即使只有一行语句
   - 左括号{ 必须在if 或 else 的同一行
   - 在 if 之后，条件语句之前，可以添加变量初始化语句，使用 ";" 进行分隔

5. 示例

   - 用if语言判断数据奇数偶数

     ```go
     num :=20
     if num % 2 ==0{
         fmt.Println(num, " 偶数")
     }else {
          fmt.Println(num, " 奇数")
     }
     ```

     

   - 判断学生平均成绩。有优、良、中、及格、不及格等五档

     ```go
     	score :=88
     	if score>=90{
     		fmt.Println("优秀")
     	}else if score>=80{
     		fmt.Println("良好")
     	}else if score>=70{
     		fmt.Println("中等")
     	}else if score>60{
     		fmt.Println("及格")
     	}else{
     		fmt.Println("不及格")
     	}
     ```

6. if 语句特殊写法

- if statement;condition {
       // 代码块

  }

- 示例： 判断一个数是奇数还是偶数

  ```go
  if num :=10;num % 2 =0{
      fmt.Println(num," 偶数")
  }else {
      fmt.Println(num," 奇数")
  }
  ```

- 注意：num定义在 if 里，所以只能够在改 if...else语句块中使用，否则编译器会报错。

##### （三）if 嵌套语句

- 略...

##### （四）switch分支语句

1. 语法结构

   ```go
   switch var1 {
       case val1:
       ...
       case val2:
       ...
       default:
       ...
   }
   ```

2. switch 语句中的注意细节
   - switch 语句执行的过程自上而下，直到找到case匹配项，匹配项中无需使用break，因为Go语言中的switch默认给每个case自带break，因此匹配成功后不会向下执行其他的case分支，而是跳出整个switch。
   - case 后的值不能重复
   - 可以<font color="red">同时测试多个符合条件的值</font>，也就是说case后可以有多个值，这些值之间使用逗号分隔，例如：case val1,val2,val3
   - Go 语言中switch后的表达式可以省略，则默认是switch true
   - Go语言中的switch case 因自带break，所以匹配某个case后不会自动向下执行其他case，如需要贯通后续的case，可以添加 fallthrough ，强制执行后面的case分支
   - fallthrough 必须放在case分支的最后一行。

3. 示例

   - 判断学生平均成绩。有优、良、中、及格、不及格五档。

   ```go
   grade := ""
   score := 78.5
   
   switch { //switch 后面省略不写，默认相当于： switch true
   case score >= 90:
   	grade = "A"
   case score >= 80:
   	grade = "B"
   case score >= 70:
   	grade = "C"
   case score >= 60:
   	grade = "D"
   
   default:
   	grade = "E"
   }
   
   fmt.Printf("你的等级是： %s\n", grade)
   fmt.Print("最终评价是：")
   
   switch grade {
   case "A":
   	fmt.Println("优秀！")
   case "B":
   	fmt.Println("良好")
   case "C":
   	fmt.Println("中等")
   case "D":
   	fmt.Println("及格")
   default:
   	fmt.Println("差")
   }
   ```

   - 判断某年某月的天数

     ```go
     year := 2021
     month := 6
     days := 0
     
     switch month {
     case 1, 3, 5, 7, 8, 10, 12:
     	days = 31
     case 4, 6, 9, 11:
     	days = 30
     case 2:
     	if (year%4 == 0 && year%100 == 0) || year%400 == 0 {
     		days = 29
     	} else {
     		days = 28
     	}
     default:
     	days = -1
     
     }
     fmt.Printf("%d 年 %d 月的天数为：%d\n", year, month, days)
     ```

4. Type switch

   - switch 语句还可以被用于type-switch来判断某个interface变量中实际存储的变量类型

   - 语法结构

     ```go
     switch x.(type){
         case type:
         statement(s)
         case type:
         statement(s)
         
         default:
         statement(s)
     }
     ```

     - 示例

       ```go
       	var x interface{}
       	switch i := x.(type) {
       	case nil:
       		fmt.Printf(" x 的类型 :%T", i)
       	case int:
       		fmt.Printf("x 是 int 型")
       	case float64:
       		fmt.Printf("x 是 float64 型")
       	case func(int) float64:
       		fmt.Printf("x 是 func(int) 型")
       	case bool, string:
       		fmt.Printf("x 是 bool 或 string 型")
       	default:
       		fmt.Printf("未知型")
       	}
       ```

     - 输出结果

       ```go
        x 的类型 :<nil>
       ```

       

##### （五）select 语句（后续讲解）

- select 语句类似与 switch 语句，但是 select 会随机执行一个可运行的 case。如果没有case 可运行，它将阻塞，直到有case可运行。

- 示例代码

  ```go
  var c1,c2,c3 chan int
  var i1,i2 int
  select {
      case i1 = <-c1:
      	fmt.Printf("received ",i1," from c1\n")
      case c2 <- i2:
      	fmt.Printf("sent ",i2, " to c2\n")
  	case i3,ok :=(<-c3): //same as: i3,ok :=<-c3
      if ok {
          fmt.Printf("received ",i3," from c3\n")
      }else {
          fmt.Printf("c3 is closed \n")
      }
      
      default:
      	fmt.Printf("no communication\n")
  }
  ```

##### （六）循环语句

1. 概述

   - 在不少实际问题中有许多具有规律性的重复操作，因此在程序中就需要重复执行某些语句。循环有包括 循环处理语句及循环控制语句。
   - 循环语句有： for循环 和 嵌套循环
   - 循环控制语句有：break语句、continue语句、goto语句

2. for 循环语句

   - for 是Go语言中唯一的循环语句，Go 没有 while、do...while 循环

   - for 循环中for关键字后不能加小括号

   - 语法形式一

     for 初始语句init; 条件表达式condition; 结束语句post {

     ​	// 循环体代码

     }

   ​     示例代码：

   ```go
   for i :=0;i<=10;i++{
   	fmt.Printf("%d\n",i)
   }
   ```

   - 语法形式二

   - for 循环条件condition { }

     效果类似其它编程语言中的while循环

     示例代码：

     ```go
     var i int
     for i<=10{
       fmt.Println(i)
       i++
     }
     ```

   - 语法形式三

     for {}

     效果与其它编程语言的 for(;;){} 一致，此时for执行无限循环

   - 示例代码

     ```go
     var i int
     for {
         if (i>10){
             break
         }
         fmt.Println(i)
         i++
     }
     ```

   - 语法形式四（<font color="red">for ...range</font>）

     ```go
     for key,value :=range map {
     	fmt.Println(key,value)
     }
     ```

     - 案例1： 遍历字符串，获得字符

       ```go
       str := "golang"
       for k, v := range str {
           fmt.Printf("第 %d 位的ASCII值 %d, 字符是 %c\n", k, v, v)
       }
       ```

     - 输出结果

       ```go
       第 0 位的ASCII值 103, 字符是 g
       第 1 位的ASCII值 111, 字符是 o
       第 2 位的ASCII值 108, 字符是 l
       第 3 位的ASCII值 97, 字符是 a
       第 4 位的ASCII值 110, 字符是 n
       第 5 位的ASCII值 103, 字符是 g
       ```

##### （七）for 嵌套循环语句

1. 语法结构

```go
for [condition | ( init; condition; increment ) | Range] {
	for [condition | ( init; condition; increment ) | Range] {
		statement(s);
	}
	statement(s);
}		
```

2. 案例代码

   - 案例1： 打印直角三角形

     ```go
     	// 定义⾏数
     	lines := 8
     	for i := 0; i < lines; i++ {
     		for n := 0; n < 2*i+1; n++ {
     			fmt.Print("* ")
     		}
     		fmt.Println()
     	}
     ```

     

   - 运行结果

     ![1623993422802](E:\lnb\study\笔记\go语言学习笔记.assets\1623838543318.png)

   - 案例2：使用循环嵌套输出 2 到 100 间 的素数

     ```go
     func printPrimeNumber() {
     	fmt.Print("2 - 100的素数：")
     	var a, b int
     	for a = 2; a <= 100; a++ {
     		for b = 2; b <= (a / b); b++ {
     			if a%b == 0 {
     				break
     			}
     		}
     		if b > (a / b) {
     			fmt.Printf("%d\t", a)
     		}
     	}
     }
     ```

     

   - 输出结果

     `2 - 100的素数：2	3	5	7	11	13	17	19	23	29	31	37	41	43	47	53	59	61	67	71	73	79	83	89	97	`
     
   - 案例3：冒泡排序
   
     ```go
     // 冒泡排序
     func BuffleSort(arr []int) {
     
     	for i := 0; i < len(arr)-1; i++ {
     		for j := 0; j < len(arr)-i-1; j++ {
     			if arr[j] > arr[j+1] {
     				// 如果前者比后者大，则执行交换
     				arr[j], arr[j+1] = arr[j+1], arr[j]
     			}
     		}
     	}
     }
     
     // 调用方法
     arr :=[]int{12,3,44,8,33}
     fmt.Println("排序前：",arr)
     
     go_base.BuffleSort(arr)
     fmt.Println("排序后：",arr)
     ```
   
     - 输出结果
   
       ```go
       排序前： [12 3 44 8 33]
       排序后： [3 8 12 33 44]
       ```

##### （八）循环控制语句

1. break

   break：跳出循环体。break语句用于终止当前for循环，并开始执行循环之后的语句。

   - 示例代码

     ```go
     func main() {
         for i := 1; i <= 10; i++ {
             if i > 5 {
             break // 如果i > 5，则循环终止
             }
             fmt.Printf("%d ", i)
         }
         fmt.Printf("\nline after for loop")
     }
     ```

2. continue

   continue: <font color="red">跳过</font>当前循环执行下一次循环语句。

   - 示例代码

     ```go
     func main() {
     	for i := 1; i <= 10; i++ {
     		if i%2 == 0 {
     		continue
     		}	
     		fmt.Printf("%d ", i)
     	}
     }
     ```

     【备注】：break，continue的区别

     - break语句将无条件跳出并结束当前的循环，然后执行循环体后的语句；

     - continue语句是跳出当前的循环语句，而开始执行下一次循环。

       示例代码：

       ```go
       // 1. break终止循环
       for i :=0;i<10;i++{
       	if i==5{
       		break
       	}
           fmt.Print(i)
       }
       
       // 运行结果： 01234
       
       // 2. continue跳出某次循环
       for i :=0;i<10;i++{
       	if i==5{
       		continue
       	}
           fmt.Print(i)
       }
       
       // 运行结果：012346789
       ```

3. goto

   - Go语言的goto语句可以无条件地转移到程序指定的行

   - goto语句通常与条件语句配合使用。可以用来实现条件转移，构成循环，跳出循环体等功能。

     但是，在结构化程序设计中一般不主张使用goto语句，以免造成程序流程的混乱，是理解和调试程序都产生困难

   - goto 语法格式如下：

     ```go
     LABEL: statement
     goto LABEL
     ```

     案例：输出1-50 之间不包含 4 的数

     ```go
     func gotoTest() {
         // 定义局部变量量
         num := 0
         /* 跳过迭代 */
         LOOP:
         for num < 50 {
             num++
              if num%10 == 4 || num/10%10 == 4 {
             goto LOOP
         }
         fmt.Printf("%d\t", num)
         }
     }
     ```

   - break 也支持结合label的用法

     语法格式如下：

     ```go
     LABEL: statement
     break LABEL
     ```

#### 七、 函数

- 函数是组织好的、可重复使用的执行特定任务的代码块。它可以提高应用程序的模块性和代码的重复利用率。

##### （一）函数定义

1. Go语言函数定义格式如下：

````
func function_name (param1 type1,param2 type2,...) (ret1 type1,ret2 type2){
	// 函数体
	return ret1,ret2
}
````

函数定义解析：

- func：函数由 func 开始声明
- function_name：函数名称，参数列表和返回值类型构成了函数签名。函数名由字母、数字和下划线组成。函数名的第一个字母不能为数字。在同一个包内，函数名称不能重复。
- param1 type1,param2 type2：参数列表，参数是可选的，也就是说函数也可以不包含参数。
- ret1 type1,ret2 type2：返回值列表。
  - 返回值返回函数的结果，结束函数的执行
  - Go语言的函数可以返回多个值
  - 返回值可以是：返回数据的数据类型，或者是：变量名+变量类型的组合
  - 函数声明时有返回值，必须在函数体中使用return语句提供返回值列表
  - 如果只有一个返回值且不声明返回值变量，那么可以省略包括返回值的括号
  - return 后的数据，要保持和声明的返回值类型、数量、顺序一致
  - 如果函数没有声明返回值，函数中也可以使用return关键字，用于强制结束函数
- 函数体：函数定义的代码集合。

##### （二）变量作用域

1. 概述

   Go语言中变量可以在三个地方声明：

   - 函数内定义的变量成为局部变量，作用域只在函数体内；
   - 函数外定义的变量称为全局变量，作用域全局甚至外包使用；
   - 函数中定义的参数称为形式参数，形式参数会作为函数的局部变量来使用。

- 示例

  ```go
  // 声明全局变量
  var a1 int = 7
  var b1 int = 9
  
  func main() {
  
  	// main函数中声明局部变量
  	a1, b1, c1 := 10, 20, 0
  
  	fmt.Printf("main 函数中 a1 =%d\n", a1)
  	fmt.Printf("main 函数中 b1 =%d\n", b1)
  	fmt.Printf("main 函数中 c1 =%d\n", c1)
  
    	// 函数调用  
  	c1 = sum(a1, b1)
  	fmt.Printf("main 函数中 c1 =%d\n", c1)
  }
  
  // 定义两个函数相加
  func sum(a1, b1 int) (c1 int) {
  	a1++
  	b1 += 2
  	c1 = a1 + b1
  
  	fmt.Printf("sum 函数中 a1 =%d\n", a1)
  	fmt.Printf("sum 函数中 b1 =%d\n", b1)
  	fmt.Printf("sum 函数中 c1 =%d\n", c1)
  	return
  }
  ```

- 输出结果：

  ```go
  main 函数中 a1 =10
  main 函数中 b1 =20
  main 函数中 c1 =0
  sum 函数中 a1 =11
  sum 函数中 b1 =22
  sum 函数中 c1 =33
  main 函数中 c1 =33
  ```

##### （三）函数变量（函数作为值）

- 在Go语言中，函数也是一种类型，可以和其他类型一样被保存在变量中。
- 可以通过 type 来定义一个自定义类型。函数的参数完全相同（包括：参数类型、个数、顺序），函数返回值相同。

1. 案例代码1：

   ```go
   func main() {
   	result := stringToLower("AbcdefGHijklMNOPqrstUVWxyz", fnCase) // 函数当做值来传递
   	fmt.Println(result)
   	result = stringToLower2("AbcdefGHijklMNOPqrstUVWxyz", fnCase)
   	fmt.Println(result)
   }
   
   func stringToLower(s string, fn func(s string) string) string {
   	fmt.Printf("%T\n", fn)
   	return fn(s)
   }
   
   // 处理字符串，奇数偶数依次显示为大小写
   func fnCase(str string) string {
   	result := ""
   	for i, v := range str {
   		if i%2 == 0 {
   			result += strings.ToUpper(string(v))
   		} else {
   			result += strings.ToLower(string(v))
   		}
   	}
   
   	return result
   }
   
   // 声明了一个函数类型，通过type关键字
   type caseFunc func(s string) string
   
   func stringToLower2(s string, fn caseFunc) string {
   	fmt.Printf("%T\n", fn)
   	return fn(s)
   }
   
   // 运行结果：
   func(string) string
   AbCdEfGhIjKlMnOpQrStUvWxYz
   main.caseFunc
   AbCdEfGhIjKlMnOpQrStUvWxYz
   ```

##### （四）匿名函数

1. 概念

- Go语言支持匿名函数，即在需要使用函数时，再定义函数，匿名函数没有函数名，只有函数体，函数可以被作为一种类型被赋值给变量，匿名函数也往往以变量的方式被传递。
- 匿名函数经常被用于实现回调函数、闭包等。

2. 定义格式

   ```
   func(参数列列表) （返回参数列列表） {
   	//函数体
   }
   ```

3. 定义匿名函数

- 在定义是调用匿名函数

```go
func main(){
    func(data int){
        fmt.Println("hello ",data)
    }(100)
}
```

- 将匿名函数赋值给变量

  ```go
  func main(){
      f :=func(data string){
          fmt.Println(data)
      }
      f("欢迎学习Go语言")
  }
  ```

4. 匿名函数的用法——作为回调函数

   ```go
   func main(){
   	// 调用函数，对每个元素进行平方根操作
       arr :=[]float64{1,9,16,25,30}
       visit(arr,func(v float64){
           v = math.Sqrt(v)
           fmt.Printf("%.2f\n",v)
       })
       
       // 调用函数，对每个元素进行平方操作
       arr :=[]float64{1,9,16,25,30}
       visit(arr,func(v float64){
           v = math.Pow(v,2)
           fmt.Printf("%.0f\n",v)
       })
   }
   
   // 定义一个函数，遍历切片元素，对每个元素进行处理
   func visit(arr []float64,f func(float64)){
       for _,value :=range arr {
           f(value)
       }
   }
   
   ```

##### （五）可变参数

1. 如果一个函数的参数类型一致，但个数不定，可以使用函数的可变参数。

2. 语法：

   ```go
   func function_name(parmas ...type) (return_list){
   	// 函数体
   	return return_list
   }
   ```

   - 当要传递若干个值到不定参数函数中的时候，可以手动书写每个参数，也可以将一个slice传递给改函数，通过"..."可以将slice中的参数对应的传递给函数。

3. 代码示例：计算学员考试总成绩及平均成绩

   ```go
   func main() {
   
   	// 1.传递n个参数
   	sum, avg, count := GetScore(90, 82.5, 73, 64.8)
   	fmt.Printf("学员共有%d门成绩，总成绩为：%.2f，平均成绩为：%.2f\n", count, sum, avg)
   
   	// 2.传递切片作为参数
   	scores := []float64{92, 72.5, 93, 74.5, 89, 87, 74}
   	sum, avg, count = GetScore(scores...)
   
   	fmt.Printf("学员共有%d门成绩，总成绩为：%.2f，平均成绩为：%.2f\n", count, sum, avg)
   	
   }
   
   // 累加求和，参数个数不定，参数个数从0-n
   func GetScore(scores ...float64) (sum, avg float64, count int) {
   	for _, v := range scores {
   		sum += v
   		count++
   	}
   	avg = sum / float64(count)
   	return
   }
   ```

4. 可变参数注意细节：

   - 一个函数最多只能有一个可变参数
   - 参数列表中还有其它类型参数，则可变参数写在所有参数的最后

##### （六）递归函数

1. 在函数内部，可以调用其它函数。如果一个函数在内部调用自身，这个函数就是递归函数。递归函数必须满足以下两个条件：

   - 在每一次调用自己时，必须更接近于解
   - 必须有一个终止处理或计算的准则

2. 示例代码：求阶乘

   - 分析：计算阶乘 n!=1x2x3 ... x n,用函数 fact(n) 表示，可以看出：fact(n) = 1x2x3 ... x(n-1) x n = fact(n-1) x n.

     所以，fact(n) 可以表示为 n x fact(n-1)，只有 n=1 是需要特殊处理

   ```go
   func factorial(n int) int{
   	if n == 0 {
   		return 1
   	}
   	return n * factorial(n-1)
   }
   ```

   - 注意：使用递归函数需要注意防止栈溢出。在计算机中，函数调用是通过栈（stack）这种数据结构实现的，每当进入一个函数调用，栈就会加一层，每当函数返回，栈就会减一层。由于栈的大小不是无限的，所以，递归调用的次数过多，会导致栈溢出。

#### 八、指针

##### （一）指针概述

1. 指针是存储另一个变量的内存地址的变量。
   - 变量是一种使用方便的占位符，变量都指向计算机的内存地址。
   - 一个指针变量可以指向任何一个值的内存地址。
   - 如下图：变量b的值 156 ，存储在内存中的地址为 0x1040a124。变量a持有b的地址，则a被认为指向b。

![1623993422802](E:\lnb\study\笔记\go语言学习笔记.assets\1623993422802.png)

2. 获取变量的地址

   - Go 语言的 <font color="red">取值地址&</font> ，一个变量前使用 &，会返回改变量的地址。

   ```go
   func main(){
   	a :=10
   	fmt.Printf("变量的地址：%x \n",&a)
   }
   
   // 运行结果：
   变量地址： c420014050
   ```

3. Go 语言指针的特点

   - 指针不能运算（不同于c语言）

   - 在Go 语言中如果对指针进行运算会报错： navlid operation: p++ (non-numeric type *int)

     

##### （二）声明指针

1. var 指针变量名 <font color="red">*指针类型</font> 

   - var ip *int  // 指向整型的指针

2. 如何使用指针

   - 定义指针变量
   - 为指针变量赋值
   - 访问指针变量中指向地址的值
   - 获取指针的值：在指针类型的变量前加上 <font color="red">* 号（前缀）来获取指针锁指向的内容</font> 

3. 示例代码一：

   ```go
   func main() {
   
   	// 声明实际变量
   	var a int = 120
   
   	// 声明指针变量
   	var ip *int
   
   	// 给指针变量赋值，将变量a的地址赋值给ip
   	ip = &a
   
   	// 打印a的类型和值
   	fmt.Printf("a的类型是：%T，值是： %v\n", a, a)
   
   	// 打印&a的类型和值
   	fmt.Printf("&a的类型是：%T，值是： %v\n", &a, &a)
   
   	// 打印ip的类型和值
   	fmt.Printf("ip的类型是：%T，值是： %v\n", ip, ip)
   
   	// 打印*ip的类型和值
   	fmt.Printf("*ip的类型是：%T，值是： %v\n", *ip, *ip)
   
   	// 打印*&a的类型和值
   	fmt.Printf("*&a的类型是：%T，值是： %v\n", *&a, *&a)
   }
   
   // 运行结果
   a的类型是：int，值是： 120
   &a的类型是：*int，值是： 0xc00000a200
   ip的类型是：*int，值是： 0xc00000a200
   *ip的类型是：int，值是： 120
   *&a的类型是：int，值是： 120
   ```

4. 示例代码二

   ```go
   func main() {
   
   	var s1 = Student{"wxqy", 20, false, 0}
   	var s2 = Student{"cherry", 30, true, 1}
   
   	var a *Student = &s1 // 将s1的内存地址赋值给Student指针变量a
   	var b *Student = &s2 // 将s2的内存地址赋值给Student指针变量b
   
   	fmt.Printf("s1类型为：%T，值为：%v\n", s1, s1)
   	fmt.Printf("s2类型为：%T，值为：%v\n", s2, s2)
   
   	fmt.Printf("a类型为：%T，值为：%v\n", a, a)
   	fmt.Printf("b类型为：%T，值为：%v\n", b, b)
   
   	fmt.Printf("*a类型为：%T，值为：%v\n", *a, *a)
   	fmt.Printf("*b类型为：%T，值为：%v\n", *b, *b)
   }
   
   // 运行结果
   s1类型为：main.Student，值为：{wxqy 20 false 0}
   s2类型为：main.Student，值为：{cherry 30 true 1}
   a类型为：*main.Student，值为：&{wxqy 20 false 0}
   b类型为：*main.Student，值为：&{cherry 30 true 1}
   *a类型为：main.Student，值为：{wxqy 20 false 0}
   *b类型为：main.Student，值为：{cherry 30 true 1}
   ```

   

##### （三） 空指针

1. Go 空指针
   - 当指针定义后没有分配到任何变量时，它的值为nil，nil指针也成为空指针
   - 指针变量通常缩写为 ptr
2. 空指针判断
   - if (ptr == nil)
   - if (ptr != nil)

##### （四）示例代码

1. 示例代码一：操作指针改变变量的数值

   ```go
   package main
   
   import "fmt"
   
   func main() {
   	b := 123
   	a := &b
   
   	fmt.Println("b 的地址： ", a)
   	fmt.Println(" *a 的值： ", *a)
   
   	*a++
   	fmt.Println("b 的新值为：", b)
   }
   
   // 运行结果
   b 的地址：  0xc000062090
   *a 的值：  123
   b 的新值为： 124
   ```

2. 示例代码二： 使用指针作为函数的参数

   ```go
   package main
   
   import "fmt"
   
   func main() {
   	a :=58
   	fmt.Println("调用函数前 a 的值： ",a)
   	fmt.Printf("%T \n",a)
   	fmt.Printf("%x \n",&a)
   
   	b :=&a
   	change(b)
   	fmt.Println("调用函数后 a 的值： ",a)
   }
   
   func change(val *int)  {
   	*val = 15
   }
   
   // 运行结果
   调用函数前 a 的值：  58
   int 
   c00000a0b8 
   调用函数后 a 的值：  15
   ```

   

#### 九、数组和切片

##### （一）什么是数组

1. 数组是具有相同类型的一组长度固定的数据序列
2. 数组元素可以通过索引下标来读取或者修改元素数据，索引从0开始
3. 数组一旦定义后，大小不能更改

##### （二）数组的语法

1. 声明数组

   Go 语言声明数组需指定 元素类型和元素个数，语法格式如下：

```go
var 变量名 [数组长度] 数据类型
```

2. 初始化数组

- var nums = [5]{1,2,3,4,5}
- 初始化数组中 {} 中元素个数不能大于 [] 中的数字
- 如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素个数来设置数组的大小
- 可以忽略声明中数组的长度并将其替换为 ... ，编译器会自动计算长度。
- var nums = [...]{1,2,3,4,5}
- 该实例与上面的示例是一样的

##### （三） 数组的长度

1. 通过 len() 函数获取数组的长度

##### （四）遍历数组

```go
func main() {
	a := [4]float64{67.7, 89.8, 21, 78}
	b := [...]int{2, 3, 5}
	
	// 遍历数组方式1
	for i :=0;i< len(a);i++{
		fmt.Print(a[i],"\t")
	}
	
	// 遍历数组方式2
	for _,v :=range b{
		fmt.Print(v,"\t")
	}
}
```

##### （五）二维数组

1. 二维数组定义

   ```go
   var arrayName[x][y] varaible_type
   ```

   例：

   ```go
   a =[3][4]int{
       {0,1,2,3},
       {4,5,6,7},
       {8,9,10,11}
   }
   ```

2. 二维数组遍历

```go
func main() {
	a := [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	// 遍历数组方式1
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Printf("a[%d][%d] = %d \n", i, j, a[i][j])
		}

	}
}
```

##### （六）数组是值类型

1. Go语言中的数组是值类型，而不是引用类型。这意味着当他们被分配给一个新变量时，将把原始的数组的副本分配给新变量。如果对新数组进行了更改，则不会在原始数组中反映。
2. 将数组传递个函数作为参数时，它们将通过值传递，而原始数组将保持不变。
3. 示例代码：

```go
func main() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}

// 运行结果
a:  [USA China India Germany France]
b:  [Singapore China India Germany France]
```

##### （七）什么是切片

1. Go 语言切片是对数组的抽象
   - Go 数组的长度不可改变，但在特定场景中这样的集合就不太适用，Go 中提出了一种灵活、功能强悍的内置类型切片（“动态数组”）
   - 与数组相比，切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大
   - 切片本身没有任何数据，它们只是对现有数组的引用
   - 切片与数组相比，不需要设定长度，在 [] 中不用设置值，相对比较灵活
   - 从概念上来说，切片（slice）像一个结构体，这个结构体包含三个元素：
     - 指针，指向数组中slice 指定的开始位置
     - 长度，即slice的长度
     - 最大长度，也就是 slice 开始位置到数组的最后位置的长度

##### （八）切片语法

1. 声明切片

   1）声明一个未指定长度的数组来定义切片

   - var sli []type
   - 切片不需要说明长度
   - 该声明方式，且未初始化的切片为空切片。该切片默认为 nil，长度为 0

   2） 使用make() 函数来创建切片

   - var slice1 []type = make([]type,len)
   - 简写为： slice1 :=make([]type,len)
   - 可以指定容量，其中capacity为可选参数： make([]type,length,capacity) 

2. 初始化
   - 直接初始化：  s :=[]int{1,2,3}
   - 通过数组截取来初始化切片： 
     - 数组：arr :=[5]int{1,2,3,4,5}
     - s :=arr[:]
     - s :=arr [startIndex:endIndex] 将arr中下标startIndex到endIndex-1 下的元素创建为一个新的切片（前闭后开），长度为 endIndex - startIndex
3. 切片截取示例

``` go
func main() {

	// 创建切片
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(numbers)

	// 打印原始切片
	fmt.Println("numbers ==", numbers)

	/* 打印⼦子切⽚片从索引 1 (包含) 到索引 4 (不不包含) */
	fmt.Println("numbers[1D4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])
	/* 默认上限为 len ( s ) */
	fmt.Println("numbers[4:] ==", numbers[4:])

	/* 打印⼦子切⽚片从索引 0 (包含) 到索引 2 (不不包含) */
	number2 := numbers[:2]
	printSlice(number2)
	/* 打印⼦子切⽚片从索引 2 (包含) 到索引 5 (不不包含) */
	number3 := numbers[2:5]
	printSlice(number3)


}

func printSlice(x []int) {
	fmt.Printf("len= %d cap=%d slice=%v\n", len(x), cap(x), x)
}

// 运行结果
len= 10 cap=10 slice=[0 1 2 3 4 5 6 7 8 9]
numbers == [0 1 2 3 4 5 6 7 8 9]
numbers[1D4] == [1 2 3]
numbers[:3] == [0 1 2]
numbers[4:] == [4 5 6 7 8 9]
len= 2 cap=10 slice=[0 1]
len= 3 cap=8 slice=[2 3 4]
```

##### （九）len() 和 cap() 函数

1. 切片的长度是切片中元素的数量
2. 切片的容量是从创建切片的索引开始的底层数组中元素的数量
3. 切片是可以索引的，并且可以有 len() 方法获取长度，切片提供了计算容量的方法cap() ，可以测量切片最长可以达到多少。（数组计算cap() 结果和 len() 相同）
4. 切片实际的是获取数组的某一部分， len 切片 <= cap 切片 <= len 数组
5. cap() 的结果决定了切片截取的注意细节

##### （十）切片是引用类型

1. slice 没有自己的任何数据。它只是底层数组的一个引用。对slice所做的任何修改都将反映在底层数组中
2. 数组是值类型，切片是引用了类型
3. 数组与切片区别的示例代码

```go
func main() {

	a := [4]float64{67.7, 89.8, 21, 78}
	b := []int{2, 3, 5}
	fmt.Printf("变量量a —— 地址：%p ， 类型：%T，数值：%v，⻓长度：%d\n", &a, a, a, len(a))
	fmt.Printf("变量量b —— 地址：%p ， 类型：%T，数值：%v，⻓长度：%d\n", &b, b, b, len(b))
	
    c := a
	d := b
	fmt.Printf("变量量c —— 地址：%p ， 类型：%T，数值：%v，⻓长度：%d\n", &c, c, c, len(c))
	fmt.Printf("变量量d —— 地址：%p ， 类型：%T，数值：%v，⻓长度：%d\n", &d, d, d, len(d))
	
    a[1] = 200
	fmt.Println("a=", a, "c=", c) // a= [67.7 200 21 78] c= [67.7 89.8 21 78]
	d[0] = 100
	fmt.Println("b=", b, "d=", d) // b= [100 3 5] d= [100 3 5]

}

// 运行结果
变量量a —— 地址：0xc0000123a0 ， 类型：[4]float64，数值：[67.7 89.8 21 78]，⻓长度：4
变量量b —— 地址：0xc0000044c0 ， 类型：[]int，数值：[2 3 5]，⻓长度：3
变量量c —— 地址：0xc000012440 ， 类型：[4]float64，数值：[67.7 89.8 21 78]，⻓长度：4
变量量d —— 地址：0xc000004540 ， 类型：[]int，数值：[2 3 5]，⻓长度：3
a= [67.7 200 21 78] c= [67.7 89.8 21 78]
b= [100 3 5] d= [100 3 5]
```

【注】：数组是值类型，故a[1]=200时，只有数组a中对应的元素值被修改，数组c不变

​			   切片是引用类型，故 d[0]=100 时，切片b和d中对应的元素值都被修改，因为b和d指向的底层数组是一样的。

4. 修改切片数值

   - 当多个切片共享相同底层数组时，每个元素所做的更改将在数组中反应出来
   - 示例代码：

   ```go
   func main() {
   
   	// 定义数组
   	arr := [3]int{1, 2, 3}
   	// 根据数组截取切⽚片
   	nums1 := arr[:]
   	nums2 := arr[:]
   	fmt.Println("arr=", arr) // [ 1 2 3 ]
   	nums1[0] = 100
   	fmt.Println("arr=", arr) // [ 100 2 3 ]
   	nums2[1] = 200
   	fmt.Println("arr=", arr) // [ 100 200 3 ]
   
   }
   
   // 运行结果
   arr= [1 2 3]
   arr= [100 2 3]
   arr= [100 200 3]
   ```

   

##### （十一）append()函数和 copy() 函数

1. append() 函数

   - 往切片中追加新元素，可以追加一个或者多个元素，也可以追加一个切片
   - append() 函数会改变slice所引用的数组的内容，从而影响到引用统一数组的其他slice
   - 当时用append() 追加元素到切片时，如果容量不够（即 cap-len==0）,Go 会创建一个新的内存地址来存储元素

2. copy() 函数

   - 复制切片元素，将元切片中的元素赋值到目标切片中，返回反之的元素个数
   - copy方法是不会建立元切片与目标切片之间的联系。也及时两个切片不存在联系，一个修改不影响另一个

   【备注】：以上两个方法不适用于数组

3. 代码示例

```go
func main() {

	fmt.Println("1、------------------")
	numbers := make([]int, 0, 20)
	printSlices("numbers:", numbers)

	/* 向切⽚片添加⼀一个元素 */
	numbers = append(numbers, 1) // [ 0 1 ]
	printSlices("numbers:", numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4, 5, 6, 7) // [ 0 1 2 3 4 5 6 7 ]
	printSlices("numbers:", numbers)

	fmt.Println("2、------------------")
	// 追加⼀一个切⽚片
	s1 := []int{100, 200, 300, 400, 500, 600, 700}
	numbers = append(numbers, s1...)
	printSlices("numbers:", numbers)

	fmt.Println("3、------------------")
	// 切⽚片删除元素
	// 删除第⼀一个元素
	numbers = numbers[1:]
	printSlices("numbers:", numbers)
	// 删除最后⼀一个元素
	numbers = numbers[:len(numbers)-1]
	printSlices("numbers:", numbers)
	// 删除中间⼀一个元素
	a := int(len(numbers) / 2)
	fmt.Println("中间数：", a)
	numbers = append(numbers[:a], numbers[a+1:]...)
	printSlices("numbers:", numbers)
	fmt.Println("4、========================")
	/* 创建切⽚片 numbers1 是之前切⽚片的两倍容量量 */
	//numbers1 := make ([] int , 0 , ( cap ( numbers )) *2 )
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	/* 拷⻉贝 numbers 的内容到 numbers1 */
	count := copy(numbers1, numbers)
	fmt.Println("拷⻉贝个数：", count)
	printSlices("numbers1:", numbers1)
	numbers[len(numbers)-1] = 99
	numbers1[0] = 100
	/*numbers1 与 numbers 两者不不存在联系， numbers 发⽣生变化时，
	  numbers1 是不不会随着变化的。也就是说 copy ⽅方法是不不会建⽴立两个切⽚片的
	  联系的
	*/
	printSlices("numbers1:", numbers1)
	printSlices("numbers:", numbers)

}

func printSlices(name string, x []int) {
	fmt.Print(name, "\t")
	fmt.Printf("addr:%p \t len=%d \t cap=%d \t slice=%v\n", x, len(x),
		cap(x), x)
}

// 运行结果
1、------------------
numbers:	addr:0xc000096000 	 len=0 	 cap=20 	 slice=[]
numbers:	addr:0xc000096000 	 len=1 	 cap=20 	 slice=[1]
numbers:	addr:0xc000096000 	 len=7 	 cap=20 	 slice=[1 2 3 4 5 6 7]
2、------------------
numbers:	addr:0xc000096000 	 len=14 	 cap=20 	 slice=[1 2 3 4 5 6 7 100 200 300 400 500 600 700]
3、------------------
numbers:	addr:0xc000096008 	 len=13 	 cap=19 	 slice=[2 3 4 5 6 7 100 200 300 400 500 600 700]
numbers:	addr:0xc000096008 	 len=12 	 cap=19 	 slice=[2 3 4 5 6 7 100 200 300 400 500 600]
中间数： 6
numbers:	addr:0xc000096008 	 len=11 	 cap=19 	 slice=[2 3 4 5 6 7 200 300 400 500 600]
4、========================
拷⻉贝个数： 11
numbers1:	addr:0xc000098000 	 len=11 	 cap=38 	 slice=[2 3 4 5 6 7 200 300 400 500 600]
numbers1:	addr:0xc000098000 	 len=11 	 cap=38 	 slice=[100 3 4 5 6 7 200 300 400 500 600]
numbers:	addr:0xc000096008 	 len=11 	 cap=19 	 slice=[2 3 4 5 6 7 200 300 400 500 99]
```

##### （十二） 冒泡排序

1. 原理
   - 比较相邻的元素。如果第一个比第二个大，就交换他们两个
   - 对每一对相邻元素做同样的工作，从开始第一对到结尾的最后一对。在这一点，最后的元素应该会是最大的数
   - 针对所有的元素重复以上的步骤，除了最后一个
   - 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较
2. 冒泡分析

![1624011788049](E:\lnb\study\笔记\go语言学习笔记.assets\1624011788049.png)

![1624011797700](E:\lnb\study\笔记\go语言学习笔记.assets\1624011797700.png)

![1624011804632](E:\lnb\study\笔记\go语言学习笔记.assets\1624011804632.png)

3. 代码示例（核心代码）

```go
// 从⼩小到⼤大排列列
func BubbleSort(arr []int) {
	iCount := 0 // 记录内循环次数
	jCount := 0 // 记录外循环的次数
	// 双层 for 循环
	for i := 0; i < len(arr)-1; i++ {
		// 定义⼀个标记，判断本轮是否有两两换位。如果没有换位，那说明排序结束，可以跳出循环。
		// 例例如： 12345 ，当执行第一轮循环后，所有相邻的两个数值都⽆无需换位，那说明排序正常，⽆需排序。不⽤用执行后续的循环。
		flag := true
		for j := 0; j < len(arr)-1-i; j++ {
			// 判断相邻两个数据的⼤小
			if arr[j] > arr[j+1] {
				// 如果前者⽐后者大，则执行数据交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// 如果出现换位，说明排序还没有结束，需要继续循环执行。
				flag = false
			}
			iCount++
		}
		jCount++
		// 如果本轮没有换位，那说明排序结束，则跳出循环
		if flag {
			break
		}
	}
	fmt.Println("i循环次数=", jCount)
	fmt.Println("j循环次数=", iCount)
}
```

#### 十、map

##### （一）什么是map

1. map 是Go中的内置类型，是一种无序的、键值对形式的集合。
   - 特点：无序、key唯一、长度不固定

##### （二）map 的用法

1. map声明 （两种方式）

   - var  var_name map[key_type]value_type
   - var_name = make(map[key_type]value_type)

2. 往map中存放键值对

   - 示例代码2

   ```go
   var country = map[string]string{
    "China": "Beijing",
    "France": "Paris",
    "Italy": "Rome",
    "Japan": "Tokyo",
    //"Japan": "Tokyo2",//key名不能重复
   }
   ```

   - 示例代码2

   ```go
   map1 := map[string]string{
    "element": "div",
    "width": "100px",
    "height": "200px",
    "border": "solid",
    "color": "red",
    "background": "none", }
   ```

   - 示例代码3： 先声明map，再赋值

   ```go
   countryCapitalMap := make(map[string]string)
   /* map 插⼊ key-value 对，各个国家对应的⾸都 */
   countryCapitalMap["China"] = "Beijing"
   countryCapitalMap["France"] = "Paris"
   countryCapitalMap["Italy"] = "Rome"
   countryCapitalMap["Japan"] = "Tokyo"
   countryCapitalMap["India"] = "New Delhi"
   ```

3. map 遍历

   - 示例代码

   ```go
   // a. 遍历key与value
   for k, v := range countryCapitalMap {
   	fmt.Println("国家是：", k, "⾸都：", v) 
   }
   
   //b.只遍历value
   for _, v := range countryCapitalMap {
    	fmt.Println("国家是：?", "⾸都：", v) 
   }
   
   // c.只遍历key
   for k := range countryCapitalMap {
   	fmt.Println("国家是：", k, "⾸都：", countryCapitalMap[k])
   }
   ```

4. 查看元素在集合中是否存在

   - 可以通过key 获取 map中对应的value的值。语法为： map[key]，但是 当key不存在的时候，map会返回该value值类型的默认值，比如string类型得到空字符串，int类型得到0。但是程序不会报错。
   - 所以可以通过 value,ok :=map[key] ，获取key /value 是否存在。ok 是bool类型，如果 OK为true，则存在，否则不存在。
   - 示例代码

   ```go
   captial, ok := countryCapitalMap["United States"]
   /* 如果 ok 是 true, 则存在，否则不存在 */
   if ok {
    fmt.Println("⾸都是：", captial)
   } else {
    fmt.Println("该国家的⾸都没有列出！")
   }
   ```

##### （三）delete() 函数

1. delete(map,key) 函数用于删除集合中的某个元素，参数为map 和其对应的key。删除函数不返回任何值

```go
func main() {

	// 1. 创建并初始化map
	m := map[string]string{
		"name":   "golang",
		"age":    "13",
		"sex":    "male",
		"height": "12",
		"length": "6",
	}

	// 2. 根据key删除
	fmt.Println("删除前map： ", m)
	if _, ok := m["sex"]; ok {
		delete(m, "sex")
	}
	fmt.Println("删除后map： ", m)

	//	清空map
	m = make(map[string]string)
	fmt.Println("清空后的map: ", m, len(m))

}

// 输出结果
删除前map：  map[age:13 height:12 length:6 name:golang sex:male]
删除后map：  map[age:13 height:12 length:6 name:golang]
清空后的map:  map[] 0
```

##### （四）清空map

- Go 语言没有为 map 提供任何清空map的函数，清空map的唯一办法就是重新make一个map
- 不用担心垃圾回收的效率，Go 语言垃圾回收比写一个清空map的函数更高效

##### （五）map是引用类型

- 与切片相似，map是引用类型。当将map分配给一个新变量时，他们都指向相同的内部数据结果。因此，一个的变化会反映另一个。

#### 十一、 结构体

##### （一）定义结构体

1. 什么是结构体

   结构体是由一系列具有相同类型或不通类型的数据构成的数据集合。

2. 结构体的定义格式

   type 结构体名称 struct {

   ​	element1 type1

   ​	element2 type2

   ​	element3 ,element4 type3

   }

   - 结构体名称：在同一包内不能重复
   - element：结构体中的成员属性，也叫字段，必须唯一
   - 同类型的成员属性可以写在一行

   示例代码

   ```go
   // 定义一个结构体
   type Teacher struct{
       name string
       age int8
       sex byte
   }
   
   ```

##### （二） 实例化结构体—— 为结构体分配内存并初始化

1. 声明方式实例化结构体，初始化方式为：对象.属性 = 值

   ``` go
   var t1 Teacher
   t1.name = "Andy"
   t1.age = 20
   t1.sex = 1
   ```

2. 变量短声明格式实例化结构体，初始化方式为：对象.属性 = 值

   ``` go
   t2 :=Teacher{}
   t2.name = "Cherry"
   t2.age = 18
   t2.sex = 0
   ```

3. 变量简短声明格式实例化结构体，声明时初始化。初始化方式为：属性:值

   ```go
   t3 :=Teacher{
       name: "zyt",
       age: 22,
       sex: 0
   }
   ```

4. 变量简短声明格式实例化，声明时初始化，不写属性名，按属性顺序只写属性值

   ```go
   t4 :=Teacher{"cy",24,0}
   ```

5. 创建指针类型的结构体

   - 使用内置函数new() 对结构体进行实例化，结构体实例化后行程指针类型的结构体
   - new 内置函数会分配内存。第一个参数是类型，而不是值，返回的值是指针向该类型新分配的零值的指针。

   ```go
   t5 :=new(Teacher)
   (*t5).name = "ty"
   (*t5).age = 26
   t5.sex = 0 // 语法简写形式，语法糖
   ```

##### （三）结构体是值类型

```go
package main

import "fmt"

type Human struct {
	name string
	age  int8
	sex  byte
}

func main() {
	// 1. 初始化Human
	h1 := Human{"zyt", 23, 0}
	fmt.Printf("h1: %T,%v,%p\n", h1, h1, &h1)
	fmt.Println("----------------------------")

	// 2.将结构体对象进行拷贝
	h2 := h1
	h2.name = "cy"
	h2.age = 20
	fmt.Printf("h2修改后=%T,%v,%p\n", h2, h2, &h2)
	fmt.Printf("h1: %T,%v,%p\n", h1, h1, &h1)
	fmt.Println("----------------------------")

	// 将结构体对象作为参数传递
	changeName(h1)
	fmt.Printf("h1: %T,%v,%p\n", h1, h1, &h1)
	fmt.Println("----------------------------")

}

func changeName(h Human) {
	h.name = "wxqy"
	h.sex = 1
	fmt.Printf("函数内h修改后=%T , %v , %p\n", h, h, &h)
}

// 运行结果
h1: main.Human,{zyt 23 0},0xc0000044c0
----------------------------
h2修改后=main.Human,{cy 20 0},0xc000004540
h1: main.Human,{zyt 23 0},0xc0000044c0
----------------------------
函数内h修改后=main.Human , {wxqy 23 1} , 0xc0000045e0
h1: main.Human,{zyt 23 0},0xc0000044c0
----------------------------
```

##### （四）结构体的深拷贝和浅拷贝

- 值类型是深拷贝
- 引用类型是浅拷贝

1. 示例代码

   ```go
   package main
   
   import "fmt"
   
   type Dog struct {
   	name  string
   	color string
   	age   int8
   	kind  string // 品种
   }
   
   func main() {
   	// 1. 实现结构体的深拷贝
   	// struct的数据类型：值类型，所以默认的复制就是深拷贝
   	d1 := Dog{"豆豆", "black", 2, "二哈"}
   	fmt.Printf("d1：%T , %v , %p \n", d1, d1, &d1) // main.Dog , {豆豆 black 2 二哈} , 0xc00003e080
   
   	d2 := d1  //深拷贝 dog
   	fmt.Printf("d2：%T , %v , %p \n", d2, d2, &d2) // main.Dog , {豆豆 black 2 二哈} , 0xc000048100
   
   	// 修改 d2 ， d1 是否也发⽣生变化？
   	d2.name = "毛毛"
   	fmt.Println("d2修改后=", d2) // d2修改后= {毛毛 black 2 二哈}
   	fmt.Println("d1=", d1)    // d1= {豆豆 black 2 二哈}
   	fmt.Println("------------------------")
   
   	//2 、实现结构体的浅拷贝：直接拷贝指针地址实现浅拷贝
   	d3 := &d1
   	fmt.Printf("d3：%T , %v , %p \n", d3, d3, d3) // &{豆豆 black 2 二哈} , 0xc00003e080
   	d3.kind = "萨摩耶"
   	d3.color = "白色"
   	d3.name = "球球"
   	fmt.Println("d3修改后=", d3) // d3修改后= &{球球 白色 2 萨摩耶}
   	fmt.Println("d1=", d1) //d1= {球球 白色 2 萨摩耶}
   	fmt.Println("------------------------")
   
   	//3 、实现结构体的浅拷贝
   	// 拷贝通过 new 函数实例例化的对象
   	d4 := new(Dog) //*Dog
   	d4.name = "多多"
   	d4.color = "棕色"
   	d4.age = 1
   	d4.kind = "巴哥⽝犬"
   	d5 := d4 //*Dog
   	fmt.Printf("d4：%T , %v , %p \n", d4, d4, d4) // *main.Dog , &{多多 棕色 1 巴哥⽝犬} , 0xc00003e340
   	fmt.Printf("d5：%T , %v , %p \n", d5, d5, d5) // *main.Dog , &{多多 棕色 1 巴哥⽝犬} , 0xc00003e340
   
   	// 修改 d5 ， d4 是否也发⽣生变化？
   	d5.color = "金色"
   	d5.kind = "金毛"
   	fmt.Println("d5修改后=", d5) // d5修改后= &{多多 金色 1 金毛}
   	fmt.Println("d4=", d4) //d4= &{多多 金色 1 金毛}
   	fmt.Println("------------------------")
   }
   
   // 输出结果
   d1：main.Dog , {豆豆 black 2 二哈} , 0xc00003e080 
   d2：main.Dog , {豆豆 black 2 二哈} , 0xc00003e140 
   d2修改后= {毛毛 black 2 二哈}
   d1= {豆豆 black 2 二哈}
   ------------------------
   d3：*main.Dog , &{豆豆 black 2 二哈} , 0xc00003e080 
   d3修改后= &{球球 白色 2 萨摩耶}
   d1= {球球 白色 2 萨摩耶}
   ------------------------
   d4：*main.Dog , &{多多 棕色 1 巴哥⽝犬} , 0xc00003e340 
   d5：*main.Dog , &{多多 棕色 1 巴哥⽝犬} , 0xc00003e340 
   d5修改后= &{多多 金色 1 金毛}
   d4= &{多多 金色 1 金毛}
   ------------------------
   ```

##### （五）匿名结构体和匿名字段

1. 匿名结构体

   （1）概念

   - 名义名字的结构体。无需通过type关键字定义就可以直接使用

   - 在创建匿名结构体时，同时要创建对象

   - 匿名结构体由结构体定义和键值对初始化两部分组成。

   - 语法格式

     变量名 :=struct {

     ​	// 定义成员属性

     }{ // 初始化成员属性  }

     

   （2）示例代码

   ```go
   func main() {
   	// 匿名函数
   	res := func(a, b float64) float64 {
   		return math.Pow(a, b)
   	}(2, 3)
   
   	fmt.Println(res)
   
   	// 匿名结构体
   	addr := struct {
   		province, city string
   	}{"陕西省", "西安市"}
   	fmt.Println(addr)
   
   	cat := struct {
   		name, color string
   		age         int8
   	}{
   		name:  "喵米",
   		color: "black",
   		age:   2,
   	}
   	fmt.Println(cat)
   }
   
   ```

2. 结构体的匿名字段

   （1）概念

   - 结构体中的字段没有名字，只包含一个没有字段名的类型。这些字段被称为匿名字段
   - 如果字段没有名字，那么默认使用类型作为字段名
   - 注意：同一个类型只能写一个
   - 结构体嵌套中采用匿名结构体字段可以模拟继承关系

   （2）示例代码

   ```go
   type User struct {
   	string
   	byte
   	int8
   	float64
   }
   
   func main() {
   	// 实例化结构体
   	user :=User{"andy",'f',22,177.2}
   	fmt.Println(user)
   
   	// 如果想依次获取姓名、年龄、性别、身高，可以这样写：
   	fmt.Printf("姓名：%s,性别：%c,年龄：%d,身高：%.2f",user.string,user.byte,user.int8,user.float64)
   }
   
   // 输出结果
   {andy 102 22 177.2}
   姓名：andy,性别：f,年龄：22,身高：177.20
   ```

##### （六）结构体嵌套

1. 概念
   - 将一个结构体作为另一个结构体的属性（字段），这种结构就是结构体嵌套
   - 结构体嵌套可以模拟对象中的两种关系
     - 聚合关系：一个类作为另一个类的属性
     - 继承关系：一个类作为另一个类的子类。子类和父类
2. 示例代码
   - 结构体嵌套模拟聚合关系

```go
type Address struct {
	province, city string
}

type Person struct {
	name string
	age  int
	addr Address
}

func main() {
	// 模拟对象之间的聚合关系
	p := Person{}
	p.name = "Andy"
	p.age = 23

	// 赋值方式1
	addr := Address{}
	addr.province = "陕西"
	addr.city = "西安"
	p.addr = addr
	fmt.Println(p)
	fmt.Println("姓名:", p.name)
	fmt.Println("年年龄:", p.age)
	fmt.Println("省:", p.addr.province)
	fmt.Println("市:", p.addr.city)
	fmt.Println("---------------------")

	// 修改Person对象的数据，是否影响Address对象？
	p.addr.city = "咸阳市"
	fmt.Println("姓名:", p.name)
	fmt.Println("年年龄:", p.age)
	fmt.Println("省:", p.addr.province)
	fmt.Println("市:", p.addr.city)
	fmt.Println("---------------------")
	fmt.Println("市:", addr.city) // 没有影响

	// 修改 Address 对象的数据，是否影响 Person 对象？
	addr.city = "彬州市"
	fmt.Println("姓名:", p.name)
	fmt.Println("年龄:", p.age)
	fmt.Println("省:", p.addr.province)
	fmt.Println("市:", p.addr.city)// 没有影响
	fmt.Println("---------------------")

	// 赋值方式2：
	p.addr = Address{
		province: "陕西省",
		city: "西安市",
	}

	fmt.Println(p)
	fmt.Println("姓名:", p.name)
	fmt.Println("年年龄:", p.age)
	fmt.Println("省:", p.addr.province)
	fmt.Println("市:", p.addr.city)
	fmt.Println("---------------------")

}
```

（2）结构体嵌套模拟继承关系

- 继承是传统面向对象编程中三大特征之一。用于描述两个类之间的关系。一个类（子类、派生类）继承于另一个类（父类、超类）

- 子类可以有自己的属性和方法，也可以重写父类已有的方法

- 子类可以直接访问父类所有的属性和方法

- <font color="red">提升字段</font>

  - 在结构体中属于匿名结构体的字段称为提升字段，因为它们可以被访问，就好像它们属于用用匿名结构字段的结构一样。
  - 换句话说，父类中的字段就是提升字段

- 继承的意义：

  - 扩展类的功能
  - 避免重复代码

- 采用匿名字段的形式就是模拟继承关系。而模拟聚合关系时一定要采用有名的结构体作为字段。

- 示例代码

  ```go
  package main
  
  import "fmt"
  
  type Person struct {
  	Name string
  	Age int
  	Sex string
  }
  
  type Student struct {
  	Person // 采用匿名结构体字段模拟继承关系
  	SchoolName string
  }
  
  func main() {
  	// 1. 初始化Person
  	p1 :=Person{"Andy",24,"male"}
  	fmt.Println(p1)
  	fmt.Printf("p1: %T , %+v \n", p1, p1)
  	fmt.Println("----------------------")
  
  	// 2. 初始化Student
  	// 写法1：
  	s1 :=Student{p1,"西安交大"}
  	fmt.Println(s1)
  	fmt.Printf("s1: %T , %+v \n", s1, s1)
  	fmt.Println("----------------------")
  
  	// 写法2：
  	s2 :=Student{Person{"Cherry",26,"female"},"港大"}
  	fmt.Println(s2)
  	fmt.Printf("s2: %T , %+v \n", s2, s2)
  	fmt.Println("----------------------")
  
  	// 写法 3 ：
  	s3 := Student{Person: Person{
  		Name: "zwt",
  		Age: 19,
  		Sex: "男",
  	},
  		SchoolName: "头大",
  	}
  	fmt.Println(s3)
  	fmt.Printf("s3: %T , %+v \n", s3, s3)
  	fmt.Println("----------------------")
  
  	// 写法 4 ：
  	s4 := Student{}
  	s4.Name = "Jerry"
  	s4.Sex = "男"
  	s4.Age = 12
  	s4.SchoolName = "耳朵大"
  	fmt.Println(s4)
  	fmt.Printf("s4: %T , %+v \n", s4, s4)
  	fmt.Println("----------------------")
  }
  ```

  

#### 十二、方法

##### （一）方法的语法格式

```go
func (接受者变量 接受者类型) 方法名(参数列表) （返回值列表）{
    // 方法体
}
```

示例代码

```go
type Employee struct {
	name, currency string
	salary         int
}

func main() {
	emp1 := Employee{
		name:     "Andy Wang",
		salary:   2000,
		currency: "$",
	}
	// 调⽤用⽅方法
	emp1.displaySalary()
	// 调⽤用函数
	displaySalary(emp1)
}

//displaySalary () ⽅方法，接受者类型为 Employee
func (e Employee) displaySalary() {
	fmt.Printf("员⼯工姓名：%s ，薪资： %s%d \n", e.name, e.currency, e.salary)
}

//displaySalary () 函数，参数为 Employee 类型对象
func displaySalary(e Employee) {
	fmt.Printf("员⼯工姓名：%s ，薪资： %s%d \n", e.name, e.currency, e.salary)
}
```

##### （二）指针作为接受者

1. 若接受者不是指针，实际只是获取了一个copy，而不能真正改变接受者中原来的数据

   ```go
   type Rectangle struct {
   	width, height float64
   }
   
   func main() {
   	r1 := Rectangle{5, 8}
   	r2 := r1
   	// 打印内存地址
   	fmt.Printf("r1的地址：%p \n", &r1)
   	fmt.Printf("r2的地址：%p \n", &r2)
   	r1.setVal()
   	fmt.Println("r1.height=", r1.height)
   	fmt.Println("r2.height=", r2.height)
   	fmt.Println("--------------------")
   	r1.setVal2()
   	fmt.Println("r1.height=", r1.height)
   	fmt.Println("r2.height=", r2.height)
   }
   func (r Rectangle) setVal() {
   	fmt.Printf("setVal()⽅方法中r的地址：%p \n", &r)
   	r.height = 10
   }
   func (r *Rectangle) setVal2() {
   	fmt.Printf("setVal2()⽅方法中r的地址：%p \n", r)
   	r.height = 20
   }
   
   // 输出结果
   r1的地址：0xc00000a0d0 
   r2的地址：0xc00000a0e0 
   setVal()⽅方法中r的地址：0xc00000a110 
   r1.height= 8
   r2.height= 8
   --------------------
   setVal2()⽅方法中r的地址：0xc00000a0d0 
   r1.height= 20
   r2.height= 8
   ```

   ##### （三）method继承

   1. method 是可以继承的，如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该匿名结构体中的method

      示例代码

      ```go
      type Human struct {
      	name, phone string
      	age         int
      }
      type Student struct {
      	Human  // 匿匿名字段
      	school string
      }
      type Employee struct {
      	Human   // 匿匿名字段
      	company string
      }
      
      func main() {
      	s1 := Student{Human{"Andy", "15012345678", 13}, "港大"}
      	e1 := Employee{Human{"Cherry", "17812345678", 35}, "交大"}
      	s1.SayHi()
      	e1.SayHi()
      }
      func (h *Human) SayHi() {
      	fmt.Printf("大家好!我是 %s ，我%d岁，我的联系⽅方式是：%s\n", h.name, h.age,
      		h.phone)
      }
      
      // 输出结果
      大家好!我是 Andy ，我13岁，我的联系⽅方式是：15012345678
      大家好!我是 Cherry ，我35岁，我的联系⽅方式是：17812345678
      ```

      ##### （三）method重写

      1. 方法是可以继承和重写的。存在继承关系时，按照就近原则进行调用

      2. 示例代码

         ```go
         type Human struct {
         	name, phone string
         	age         int
         }
         type Student struct {
         	Human  // 匿匿名字段
         	school string
         }
         type Employee struct {
         	Human   // 匿匿名字段
         	company string
         }
         
         func main() {
         	s1 := Student{Human{"zwt", "15012345678", 13}, "交大"}
         	e1 := Employee{Human{"cy", "17812345678", 35}, "人大"}
         	s1.SayHi()
         	e1.SayHi()
         }
         func (h *Human) SayHi() {
         	fmt.Printf("大家好! 我是 %s ，我%d岁，我的联系方式是：%s\n", h.name, h.age,
         		h.phone)
         }
         
         //Student 的 method 重写 Human 的 method
         func (s *Student) SayHi() {
         	fmt.Printf("大家好! 我是 %s ，我%d岁，我在%s上学，我的联系方式是：%s\n", s.name,
         		s.age, s.school, s.phone)
         }
         
         //Employee 的 method 重写 Human 的 method
         func (e *Employee) SayHi() {
         	fmt.Printf("大家好! 我是 %s ，我%d岁，我在%s工作，我的联系方式是：%s\n", e.name,
         		e.age, e.company, e.phone)
         }
         
         // 输出结果
         大家好! 我是 zwt ，我13岁，我在交大上学，我的联系方式是：15012345678
         大家好! 我是 cy ，我35岁，我在人大工作，我的联系方式是：17812345678
         ```

         

#### 十三 、并发编程goroutine

##### （一）什么是Goroutine

1. 在执行多任务时，为了保证每个任务能及时被分配到CPU上运行处理，同时避免多个任务频繁地在线程间切换执行而损失效率，常使用协程池；但是如果面对随时随地可能发送的并发和线程处理需求，线程池不是非常直观和方便；
2. 使用者分配足够多的任务，系统能自动帮助使用者把任务分配到CPU上，让这些任务尽量并发运行。这种机制在Go语言中被称为Goroutine。Goroutine类似于线程，Goroutine由Go程序运行runtime调度和管理，Go程序会智能的将Goroutine中的任务合理的分配给每个CPU。
3. Go程序从main包的main()函数开始，在程序启动时，Go程序会为main() 函数创建一个默认的Goroutine

##### （二）Goroutine 在线程上的优势

1. Go中使用Goroutine 来实现并发concurrently
2. Goroutine 是与其它函数或方法同时运行的函数或方法
3. Goroutine 在线程上的优势
   - 与线程相比，Goroutine 成本很小。他们只是堆栈大小的几个kb，堆栈可以根据应用程序的需要增长和收缩，而在线程的情况下，堆栈大小必须指定并且是固定。由于创建Goroutine的成本很小。因此，Go应用程序可以并发运行数千个Goroutine
   - 当时用Goroutine 访问共享内存时，通过设计的通道可以防止竞态条件发生。通道可以被认为是Goroutine通信的管道

##### （三）使用普通函数创建goroutine

1. 在函数或方法调用前面加上go关键字，将会同时运行一个新的goroutine
   - 使用go关键字创建goroutine时，被调用的函数往往没有返回值，如果有返回值也会被忽略
   - 如果需要在goroutine中返回数据，必须使用通道channel，通过通道channel把数据从goroutine中作为返回值传出
2. 示例代码1：

```go
func main() {
	go hello()
	//time.Sleep(1*time.Second)
	fmt.Println("main function")
}

func hello()  {
	fmt.Println("Hello golang goroutine")
}

// 运行结果
main function
```

3. 代码分析
   - Go程序执行的过程是：创建和启动goroutine，初始化操作，执行main() 函数，当main() 函数结束，主goroutine随之结束，程序结束。
   - 被启动的goroutine叫做子goroutine
   - 如果main 的goroutine终止了，程序将被终止，而其他Goroutine将不在运行。换句话说，所有goroutine在main（）函数结束时会一同结束。
4. 修改后的代码

```go
func main() {
	go hello()
	time.Sleep(1*time.Second)
	fmt.Println("main function")
}

func hello()  {
	fmt.Println("Hello golang goroutine")
}

// 运行结果
Hello golang goroutine
main function
```

- 在上面的程序中，调用time.sleep()方法，它会在执行过程中休眠。在这种情况下，main的goroutine被用来休眠1s。现在调用go hello() 有足够的时间在main Goroutine终止之前执行。这个程序首先打印Hello golang goroutine，等待1s，然后打印main函数

5. 示例代码2

```go
func main() {
	go running()
	var input string
	fmt.Scanln(&input)
}

func running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)
	}
}
```

- 程序运行效果，控制台不断输出tick，同时还可以接收用户输入，两个环节同时运行。
- 该案例中，go程序在启动时，运行是runtime默认为main()函数创建一个goroutine。在main() 函数的goroutine执行到go running()语句时，归属于running()函数的goroutine被创建，running()函数开始自己的goroutine中执行。
- 此时，main() 函数继续执行，两个goroutine通过Go程序的调度机制同时运行

##### （四）调整并发的运行性能Gomaxprocs

1. 在Go程序运行时runtime实现了一个小型的任务调度器。这套调度器的工作原理类似于操作系统调度线程，Go程序调度器可以高效的将CPU资源分配给每一个任务。

   传统逻辑中，开发需要维护线程池中的线程与CPU核心数量的对应关系。在Go语言中可以通过runtime.GOMAXPROCS()函数做到。

   语法为：runtime.GOMAXPROCS(逻辑CPU数量)

2. 逻辑CPU数量有如下几种数值：

   - <1，不修改任何数值
   - =1，单核执行
   - `>1` 多核并发执行

   Go1.5 版本之前，默认使用单核执行。Go1.5版本开始，默认执行：runtime.GOMAXPROCOS(逻辑CPU数量)，让代码并发执行，最大效率地利用CPU。

#### 十四、并发编程channel

多个goroutine间通信的通道channel

##### （一）通道的概述

1. 使用通道的意义
   - 单纯的将函数并发执行没有意义。函数与函数间需要交换数据才能体现出并发执行的意义。
   - 虽然可以使用共享内存进行数据交换，但是共享内存在不同goroutine中容易发生竞态问题，必须使用互斥对内存进行枷锁，所以造成性能问题。
   - Go语言中提倡使用通道channel的方式代替共享内存。也就是说，Go语言中主张，应该通过数据传递来实现共享内存，而不是通过共享内存来实现消息传递。
   - 排队的目的是避免拥堵、插队造成的资源使用和交换过程低效问题。多个goroutine为了争抢数据，势必造成低效，使用排队的方式是最高效的，channel就是一种队列结构。
2. 什么是通道
   - Go语言中channel是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出first in first out 的规则，保证收发数据的顺序。
   - 通道可以被认为是Goroutine通信的管道。类似于管道中的水从一端到另一端的流动，数据可以从一端发送到另一端，通过通道接收。
   - 每个通道都有与其相关的类型。该类型是通道允许传输的数据类型。（通道的零值为nil）nil通道没有任何用处，因此通道必须使用类似于map和slice的方法来定义。

##### （二）声明通道类型

1. 语法

```go
var 通道变量 chan 通道类型
```

2. chan 类型的空值是nil，声明后需要配合使用make才能使用

##### （三）创建通道

1. 通道是引用类型，需要使用make进行创建

   语法： 通道示例： =make(chan 数据类型)

2. 例如：

   ```go
   ch1 :=make(chan int) // 创建一个整数类型的通道
   ch2 :=make(chan interface{}) // 创建一个空接口类型的通道，可以存放任意数据
   
   type Equip struct{ /* 属性 */ }
   ch3 :=make(chan *Equip) 
   ```

##### （四）通道发送数据

1. 使用通道发送数据的格式

   通道发送数据使用特殊的操作符 "<-"，将数据通过通道发送的语法为：channel_v <- value

   - 通道发送的值可以是变量、常量、表达式或者函数返回值等。值的类型必须与ch通道的元素类型一致
   - 把数据往通道中发送，如果接收方一直没有接收，那么发送操作将继续阻塞。此时所有的goroutine，包括main的goroutine都出于等待状态。
   - 运行是会提示报错：fatal error：all goroutine are alseep -deadlock!

2. 死锁deadlock

   - 使用通道时要考虑的一个重要因素是死锁
   - 如果Goroutine在一个通道发送数据，那么预计其他的Goroutine应该接收数据。如果这种情况不发送，那么程序将在运行时出现死锁。
   - 类似的，如果Goroutine正在等待从通道接收数据，那么另一些Goroutine将会在该通道上写入数据，否则程序会死锁。

3. 示例代码：

```go
func mian(){
    // 创建一个空指针型通道
    ch :=make(chan interface{})
    // 将0通过通道发送
    ch <- 0
    // 发送字符创
    ch <- "channel"
}
```

##### （五）阻塞

1. 一个通道发送和接收数据默认是阻塞的。
   - 当一个数据被发送到通道时，在发送语句中被阻塞，知道另一个Goroutine 从该通道读取数据。类似的，当从通道读取数据时，读取被阻塞，直到一个Goroutine将数据写入该通道。
   - 这些通道的特性是帮助Goroutine 有效的进行通信，而无需像使用其他编程语言中非常常见的显式锁或条件变量
2. 示例代码：

```go
func main() {
	var ch1 chan int
	fmt.Println(ch1)        // nil
	fmt.Printf("%T\n", ch1) // chan int

	ch1 = make(chan int)
	fmt.Println(ch1)

	ch2 := make(chan bool)

	go func() {
		fmt.Println("子goroutine....")
		data, ok := <-ch1 // 阻塞式，ch1中读取数据
		time.Sleep(1 * time.Second)
		fmt.Println("子goroutine从通道中读取到main传来的数据是：", ok, data)
		ch2 <- true // 向通道中写入数据，表示结束

	}()

	ch1 <-100 // 阻塞式，main goroutine向通道中写入数据
	<- ch2 // 目的是防止main goroutine先执行完毕后退出。因为如果main的goroutine终止，程序将终止，而其他Goroutine将不再运行
	fmt.Println("main... over")

}

// 运行结果
<nil>
chan int
0xc0000180c0
子goroutine....
子goroutine从通道中读取到main传来的数据是： true 100
main... over
```

##### （六）通道接收数据

1. 使用通道接收数据的格式

   通道接收同样使用特殊符号“<-”

   - 通道变量 <- 值
   - 通道收发操作在不同的两个goroutine间进行
   - 接收操作将持续阻塞，直到发送方发送数据
   - 每次接收一个元素

##### （七）通道接收数据的四种写法

1. 阻塞接收数据

   - data :=<-ch
   - 执行该语句时将会阻塞，直到接收到数据并赋值给data变量

2. 阻塞接收数据的完整写法

   - data,ok := <-ch
   - data: 表示接收到的数据。未接收到数据时，data为通道类型的零值
   - ok： 表示是否接收到数据
   - 通过ok值可以判断当前通道是否关闭

3. 接收任意数据，忽略接收到数据

   - <- ch
   - 执行该语句是将会阻塞
   - 其目的不在于接收通道中的数据，而是为了阻塞goroutine

4. 循环接收数据

   - 循环接收数据，需要配合使用关闭通道

   - 借助普通的for循环和for...range 语句循环接收多个元素

   - 遍历通道，遍历的结果就是接收到 的数据，数据类型就是通道的数据类型

   - 普通for循环接收通道数据，需要有break循环的条件；for...range 会自动判断出通道已关闭，而无需通过判断来终止循环

   - 循环接收数据的三种方式

     ```go
     func main() {
     	ch1 := make(chan string)
     	go sendData(ch1)
     
     	// 1. 循环接收数据方式1
     	for {
     		data := <-ch1
     		if data == "" {
     			break
     		}
     		fmt.Println("从通道中读取数据方式1：", data)
     	}
     
     	// 2. 循环接收数据方式2
     	for {
     		data, ok := <-ch1
     		if !ok {
     			break
     		}
     		fmt.Println("从通道中读取数据方式2：", data)
     	}
     
     	// 3. 循环接收数据方式3
     	// for...range 会自动判断通道是否关闭，自动break循环
     	for ch := range ch1 {
     		fmt.Println("从通道中读取数据方式3：", ch)
     	}
     }
     
     func sendData(ch1 chan string) {
     	for i := 1; i <= 10; i++ {
     		ch1 <- fmt.Sprintf("发送数据%d\n", i)
     	}
     
     	fmt.Println("发送数据完毕。。。")
     	// 显示调用close() 实现关闭通道
     	close(ch1)
     
     }
     ```

     

##### （八）关闭通道

1. 发送方如果数据写入完毕，需要关闭通道，用于通知接收方数据传递完毕。一般都是发送方关闭通道

2. 如何判断一个channel是否已经关闭？可以在读取的时候使用多个返回值的方式。如果返回false，则表示通道已经关闭

3. 如果往关闭的通道中写入数据，会报错：panic: send on closed channel

4. 示例代码：

   ```go
   func main() {
   	// 通道关闭后是否可以写入和读取数据呢？
   	ch1 := make(chan int)
   	go func() {
   		ch1 <- 100
   		ch1 <- 200
   		close(ch1)
   		//ch1 <- 10 // 关闭的通道无法写入数据
   	}()
   
   	data,ok :=<- ch1
   	fmt.Println("main读取数据：",data,ok)
   
   	data,ok =<- ch1
   	fmt.Println("main读取数据：",data,ok)
   
   	data,ok =<- ch1
   	fmt.Println("main读取数据：",data,ok)
   
   	data,ok =<- ch1
   	fmt.Println("main读取数据：",data,ok)
   
   
   	data,ok =<- ch1
   	fmt.Println("main读取数据：",data,ok)
   
   }
   
   // 运行结果
   main读取数据： 100 true
   main读取数据： 200 true
   main读取数据： 0 false
   main读取数据： 0 false
   main读取数据： 0 false
   ```

   修改代码：

   ```go
   func main() {
   	// 通道关闭后是否可以写入和读取数据呢？
   	ch1 := make(chan int)
   	go func() {
   		ch1 <- 100
   		ch1 <- 200
   		close(ch1)
   		ch1 <- 10 // 关闭的通道无法写入数据
   	}()
   
   	data,ok :=<- ch1
   	fmt.Println("main读取数据：",data,ok)
   
   	data,ok =<- ch1
   	fmt.Println("main读取数据：",data,ok)
   
   	data,ok =<- ch1
   	fmt.Println("main读取数据：",data,ok)
   
   	data,ok =<- ch1
   	fmt.Println("main读取数据：",data,ok)
   
   
   	data,ok =<- ch1
   	fmt.Println("main读取数据：",data,ok)
   
   }
   
   // 运行结果
   main读取数据： 100 true
   main读取数据： 200 true
   panic: send on closed channel
   
   goroutine 6 [running]:
   main.main.func1(0xc0000180c0)
   	E:/lnb/go/src/go-study/main.go:12 +0x80
   created by main.main
   	E:/lnb/go/src/go-study/main.go:8 +0x74
   ```



缓冲通道和定向通道

##### （九）缓冲通道

1. 缓冲通道：自带一块缓冲区，可以暂时存储数据，如果缓冲区满了，那么才会阻塞；

2. 非缓冲通道：默认创建的通道都是非缓冲通道，读写都是即时阻塞

3. 示例代码：

   ```go
   func main() {
   	// 1. 非缓冲通道
   	ch1 := make(chan int)
   	fmt.Println("非缓冲通道：", len(ch1), cap(ch1)) // 0,0
   
   	go func() {
   		data := <-ch1 // 阻塞
   		fmt.Println("获取数据： ", data)
   	}()
   
   	ch1 <- 100 // 阻塞
   	time.Sleep(1 * time.Second)
   	fmt.Println("写入数据OK -------------")
   
   	// 2. 缓冲通道
   	ch2 := make(chan string, 5)
   	fmt.Printf("%T\n", ch2)
   
   	go sendData(ch2)
   
   	for data := range ch2 {
   		fmt.Println("\t读取数据：", data)
   	}
   
   	// 读取数据完毕
   	fmt.Println("读取数据完毕")
   
   }
   
   func sendData(ch chan string) {
   	for i := 1; i <= 10; i++ {
   		ch <- fmt.Sprintf("data%d", i)
   		fmt.Println("写入数据：", i)
   		fmt.Println(len(ch), cap(ch))
   	}
   	close(ch)
   }
   ```

4. 缓冲通道模拟生产者和消费者

   ```go
   func main() {
   	ch1 := make(chan int, 5)
   	ch2 := make(chan bool) // 判断结束
   	rand.Seed(time.Now().UnixNano())
   	// 写⼊入数据：⽣生产者
   	go func() {
   		for i := 1; i <= 20; i++ {
   			ch1 <- i
   			fmt.Println("写⼊入数据：", i)
   			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
   		}
   		close(ch1)
   	}()
   	// 读取数据：消费者
   	go func() {
   		for data := range ch1 {
   			fmt.Println("\t1号消费者：", data)
   			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
   		}
   		ch2 <- true
   	}()
   	go func() {
   		for data := range ch1 { //1
   			fmt.Println("\t2号消费者：", data)
   			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
   		}
   		ch2 <- true
   	}()
   	<-ch2
   	fmt.Println("main...over...")
   }
   ```

##### （十）定向通道

1. 通道默认都是双向的，即可写入数据，又可读取数据

2. 定向通道：也叫单向通道，只读，或只写

   - 只读： make(<- chan type)，只能读取数据，不能写入数据

     ​			<- chan

   - 只写：make(chan <- type)，只能写入数据，不能读取数据

     ​			chan<- data

3. 创建通道时，采用单向通道，没有意义的。都是双向通道。

   - 将通道作为参数传递的时候使用单向通道

     定义函数，只有写入数据功能，或者只有读取数据功能

4. 定向通道的意义：在语法级别，保证通道的操作安全

5. 示例代码

   ```go
   func main() {
   	// 1. 双向通道
   	ch1 := make(chan string)
   	go fun1(ch1)
   
   	data := <-ch1
   	fmt.Println("main,接收到数据：", data)
   	ch1 <- "Go 语言好学么？"
   	ch1 <- "区块链好学么？"
   
   	go fun2(ch1)
   	go fun3(ch1)
   
   	time.Sleep(1 * time.Second)
   	fmt.Println("main over!")
   
   }
   
   func fun1(ch1 chan string) {
   	ch1 <- "我是Andy"
   	data := <-ch1
   	data2 := <-ch1
   	fmt.Println("回应：", data, data2)
   }
   
   // 功能：值写入数据
   func fun2(ch chan<- string) {
   	// 只能写入
   	ch <- "Are U OK？"
   	//<-ch // Invalid operation: <-ch (receive from send-only type chan<- string)
   }
   
   // 功能：只能读取数据
   func fun3(ch <-chan string) {
   	data := <-ch
   	fmt.Println("只读：", data)
   	// ch <- "hello" // Invalid operation: ch <- "hello" (send to receive-only type <-chan string)
   }
   ```



##### （十一）select 分支语句

- select 语句类似于switch 语句，但是select会随机执行一个可运行的case。如果没有case可以运行，它将阻塞，知道case可运行
- 每个case都必须是一个通道
- 所有channel表达式都会被求值
- 所有被发送的表达式都会被求值
- 如果任意某个通道可以进行，它就执行；其他被忽略

1. 执行流程

   - 如果多个case都可以运行，select会随机公平的选出一个执行，其他不会执行
   - 如果有default子句，则执行语句
   - 如果没有deafult子句，select将阻塞，直到某个通道可以运行
   - Go不会重新对channel或值进行求值

2. 示例代码1

   ```go
   func main() {
   
   	ch1 := make(chan int)
   	ch2 := make(chan int)
   	go func() {
   		time.Sleep(1 * time.Second)
   		ch1 <- 100
   	}()
   	go func() {
   		time.Sleep(1 * time.Second)
   		ch2 <- 200
   	}()
   	select {
   	case data := <-ch1:
   		fmt.Println("ch1中读取数据了了:", data)
   	case data := <-ch2:
   		fmt.Println("ch2中读取数据了了：", data)
   	default:
   		fmt.Println("执⾏行行了了default。。。")
   	}
   }
   
   // 运行结果
   执⾏行行了了default。。。
   
   ```

   示例代码2：

   ```go
   func main() {
   	ch1 := make(chan int)
   	ch2 := make(chan int)
   	go func() {
   		time.Sleep(1 * time.Second)
   		data := <-ch1
   		fmt.Println("ch1：", data)
   	}()
   	go func() {
   		time.Sleep(2 * time.Second)
   		data := <-ch2
   		fmt.Println("ch2：", data)
   	}()
   	select {
   	case ch1 <- 100: // 阻塞
   		close(ch1)
   		fmt.Println("ch1中写⼊入数据。。")
   	case ch2 <- 200: // 阻塞
   		close(ch2)
   		fmt.Println("ch2中写⼊入数据。。")
   	case <-time.After(2 * time.Second): // 阻塞
   		fmt.Println("执⾏行行延时通道")
   		//default:
   		// fmt.Println ( "default.." )
   	}
   	time.Sleep(4 * time.Second)
   	fmt.Printf("main over ")
   }
   
   // 运行结果
   ch1： 100
   ch1中写⼊入数据。。
   main over 
   ```

   

