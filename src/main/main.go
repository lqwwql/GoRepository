package main

import (
	"fmt"
	"grutils/grmath"
	_ "net/http/pprof"
	"reflect"
	"strconv"
	"strings"
	"time"
)


type StringS struct {
	Name string `json:"name"`
}
func main() {
	//network.HttpDo("C://123.jpg")
	//fmt.Println(time.Now().Format("20060102"))
	//str := fmt.Sprintf("已兑换%v积分，若未收到红包请联系客服！",0.01)
	//fmt.Println(str)
	//fmt.Println(fmt.Sprintf("%.2f",13.005))

	//numString := "0.4.2.5.5"
	//index := strings.Index(numString,".")
	//fmt.Printf("index = %x \n",index)
	//fmt.Printf("subStr := %s \n",grcommon.SubStrOnBytes(numString,0,index+3))

	//goroutine.GoroutineMain()

	//testScanf()

	//test()

	//channels.TestChannel()
	//getFileName("C:\\Users\\admin\\Desktop\\error.txt")

	//isImage("error.PNGs")
	//ifContainKey("cageageahea","b","a")
	//testTimeFormat()
	//testFmtFormat()
	//network.NetWorkProgram()
	//var i = 11
	//fmt.Printf("i = %3d0",i)
	//result := fmt.Sprintf("%03d",i)
	//
	//fmt.Println(result)

	//imsiHeader := ""
	//imsiSql := "AND 1=1 "
	//if imsiHeader != "" {
	//	imsiSql = strings.Join([]string{imsiSql," AND `imsi` LIKE '",imsiHeader,"%'"},"")
	//}
	//fmt.Printf("imsiSql = %s",imsiSql)


	//myreflect.ReflectTest()
	//stringToInt("12345")

	//myFlag.TestFlagPackage()
	//dbconn.InitDB()

	testStringTo01()
}

func testStringTo01(){
	str := "nmsl"
	fmt.Println([]byte(str))
}

func testSlice(){
	slicea := []int{1,2,3,4}
	fmt.Println(slicea)
	changeSlice(&slicea)
	fmt.Println(slicea)
}

func changeSlice(slicea *[]int){
	sliceb := []int{1,2,3}
	*slicea = sliceb
}

func stringToInt(str string){
	var result int
	for i:=0;i<len(str);i++ {
		s,_:= strconv.Atoi(string(str[i]))
		fmt.Println(s)
		result += s
	}
	fmt.Println(result)
}

func testFmtFormat(){
	val := []int{1,2,3}
	fmt.Printf("%v,%T \n",val,val)
	fmt.Println(reflect.TypeOf(val))
	fmt.Printf("%p,%p,%p,%p \n",val,val[0],val[1],val[2])
	fmt.Printf("%p,%p,%p,%p \n",&val,&val[0],&val[1],&val[2])


}


func testTimeFormat() {
	                                           //转化为时间戳 类型是int64
	//时间戳转日期
	dataTimeStr := time.Unix(1420041600, 0).Format(grmath.DATETIME_FORMAT) //设置时间戳 使用模板格式化为日期字符串
	fmt.Println(dataTimeStr)
}

func ifContainKey(key ... string) {
	fmt.Printf("lenth = %d ", len(key))
}

func isImage(fileName string) bool {
	start := strings.Index(fileName, ".") + 1
	end := len(fileName)
	suffix := substring(fileName, start, end)

	fmt.Println(suffix)
	if strings.EqualFold(suffix, "png") || strings.EqualFold(suffix, "jpg") || strings.EqualFold(suffix, "jpeg") {
		fmt.Println("is image")
	}
	strings.EqualFold(suffix, "png")
	fmt.Println("not a image")
	return false
}

func getFileName(originName string) {
	start := strings.LastIndex(originName, "\\") + 1
	end := len(originName)
	result := substring(originName, start, end)
	fmt.Println(result)
}

func substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 || end == 0 && end == length {
		return source
	}

	return string(r[start:end])
}

func testPointer() {
	var b int = 10
	a := &b
	var c *int
	c = a
	var d *int = &b

	fmt.Println("a = ", *a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", *c)
	fmt.Println("d = ", *d)
}

func testScanf() {
	var n int
	var m int
	fmt.Println("程序运行中.....")
	for {
		fmt.Scanf("%d%d%s", &n, &m)
		if n > m {
			fmt.Println("over.....")
			break
		}

		for i := n; i < m; i++ {
			if isPrime(i) == true {
				fmt.Printf("%d\n", i)
				continue
			}
		}
	}
}

//判断素数
func isPrime(value int) bool {
	if value <= 3 {
		return value >= 2
	}
	if value%2 == 0 || value%3 == 0 {
		return false
	}
	for i := 5; i*i <= value; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {
			return false
		}
	}
	return true
}
