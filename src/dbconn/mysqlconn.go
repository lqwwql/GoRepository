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

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ProfileId int    `json:"profile_id"`
}


var DB *sql.DB

func InitDB (){
	//构建链接
	dbPath := strings.Join([]string{username,":",password,"@tcp(",ip,":",port,")/",dbname,"?charset=utf8"},"")

	fmt.Printf("dbPath = %s \n",dbPath)

	//打开数据库
	DB , _ = sql.Open("mysql",dbPath)

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

func (user *User) InsertUserToDb()bool{
	//开启事务
	tx,err := DB.Begin()
	if err!=nil{
		fmt.Printf("db.begin err = %s \n",err)
		return false
	}

	//sql
	stmt,err := tx.Prepare("INSERT INTO `user` (`name`,`profile_id`) VALUES (?,?)")
	if err!=nil{
		fmt.Printf("prepare sql err = %s \n",err)
		return false
	}

	res,err := stmt.Exec(user.Name,user.ProfileId)
	if err!=nil{
		fmt.Printf("stmt exec err = %s \n",err)
		return false
	}

	//事务提交
	tx.Commit()
	//返回自增ID
	fmt.Println(res.LastInsertId())
	return true
}

func DeleteUserFromDB(id int)bool{
	tx,err := DB.Begin()
	if err!=nil {
		fmt.Printf("db.begin err= %s \n",err)
		return false
	}

	stmt,err := tx.Prepare("DELETE FROM `user` WHERE `id` = ?")
	if err!=nil {
		fmt.Printf("prepare sql err = %s \n",err)
		return false
	}

	res,err := stmt.Exec(id)
	if err!=nil {
		fmt.Printf("exec sql err = %s \n ",err)
		return false
	}

	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}

func (user *User)UpdateUserToDB()bool{
	tx,err := DB.Begin()
	if err!=nil {
		fmt.Printf("db.begin err= %s \n",err)
		return false
	}

	stmt,err := tx.Prepare("UPDATE `user` SET `name`=? ,`profile_id`=? WHERE `id` = ?")
	if err!=nil {
		fmt.Printf("prepare sql err = %s \n",err)
		return false
	}

	res,err := stmt.Exec(user.Name,user.ProfileId,user.Id)
	if err!=nil {
		fmt.Printf("exec sql err = %s \n ",err)
		return false
	}

	tx.Commit()
	fmt.Println(res.LastInsertId())

	return true
}

func QueryUserBYID(id int)(user User){
	err := DB.QueryRow("SELECT * FROM `user` WHERE `id`=? ",id).Scan(&user.Id,&user.Name,&user.ProfileId)
	if err!=nil{
		fmt.Printf("query sql errr = %s \n",err)
	}
	return
}

func QueryAllUser()(userList []User){
	rows,err := DB.Query("SELECT * FROM `user`")
	if err!=nil{
		fmt.Printf("query sql errr = %s \n",err)
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id,&user.Name,&user.ProfileId)
		if err!=nil {
			fmt.Printf("rows.scan err = %s \n",err)
		}
		userList = append(userList, user)
	}
	return
}