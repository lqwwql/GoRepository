package myFlag

import (
	"flag"
	"fmt"
)

//命令行读取参数
func TestFlagPackage(){
	username := flag.String("name","","Input your name")
	age := flag.Int("age",0,"Input your age")
	flag.Parse()
	fmt.Printf("name = %s , age = %d \n",*username,*age)
}
