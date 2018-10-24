package main

import (
	"fmt"
)

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

	test()
}

func test(){
	var b int = 10
	a := &b
	var c *int
	c = a
	var d *int = &b

	fmt.Println("a = ",*a)
	fmt.Println("b = ",b)
	fmt.Println("c = ",*c)
	fmt.Println("d = ",*d)
}

func testScanf(){
	var n int
	var m int
	fmt.Println("程序运行中.....")
	for {
		fmt.Scanf("%d%d%s",&n,&m)
		if n > m {
			fmt.Println("over.....")
			break
		}

		for i:=n;i<m;i++{
			if isPrime(i) == true {
				fmt.Printf("%d\n",i)
				continue
			}
		}
	}
}

//判断素数
func isPrime(value int)bool{
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
