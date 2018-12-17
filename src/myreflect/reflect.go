package myreflect

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func ReflectTest(){
	//t := reflect.TypeOf(3)
	//fmt.Println(t.String())
	//fmt.Println(t)


	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))

}







