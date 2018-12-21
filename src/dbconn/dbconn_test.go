package dbconn

import (
	"fmt"
	"testing"
)

func Test_Dbconn(t *testing.T){
	user := User{}
	user.Id = 1
	user.Name = "去死吧"
	user.ProfileId = 666

	InitDB()
	//result :=user.InsertUserToDb()
	//result := DeleteUserFromDB(2)
	//result := user.UpdateUserToDB()
	//result := QueryUserBYID(1)
	//result := QueryAllUser()
	result := TestTimeBetween()
	fmt.Println(result)
}