package dbconn

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const(
	username = "root"
	password = ""
	ip = "172.16.0.216"
	port = "3306"
	dbname = "world"
)

var DB *sql.DB

func InitDB (){
	//构建链接
	dbPath := strings.Join([]string{username,":",password,"@tcp(",ip,":",port,")/",dbname,"?charset=utf8"},"")

	fmt.Printf("dbPath = %s \n",dbPath)

	//打开数据库
	DB , err := sql.Open("mysql",dbPath)
	if err != nil {
		fmt.Printf("open mysql err = %s \n",err)
	}
	//设置最大连接数
	DB.SetConnMaxLifetime(100)
	//设置最大闲置连接数
	DB.SetMaxIdleConns(10)
	if err := DB.Ping();err != nil {
		fmt.Printf("ping mysql err = %s \n",err)
		return
	}
	fmt.Println("connect mysql success")
}
