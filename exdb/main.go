package main

import (
	"fmt"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
)


func main(){

	username	:="root"
	password	:="en98as"
	database	:="testdb"
	proto		:="tcp"
	addr		:="127.0.0.1:3306"

	db	:=mysql.New(proto,"",addr,username,password,database)

	db.Connect()

	_,err := db.Start("drop table if exists test_table")

	if (err != nil){
		fmt.Println(err)
	}


	db.Query("create table test_table (id varchar(50) not null primary key,username varchar(50) not null)")

	db.Query("insert into test_table values ('zsc','zsc')")
	db.Query("insert into test_table values ('korben','korben')")


	fmt.Println("End")
}
