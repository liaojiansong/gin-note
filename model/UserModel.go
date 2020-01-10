package model

import (
	"database/sql"
	"gin/initDB"
	"log"

	//"gin/initDB"
	//"log"
)

type UserModel struct {
	Id       int64  `form:"id"`
	Email    string `form:"email" binding:"email"`
	Password string `form:"password" `
	// sql 包的一种类型
	Avatar sql.NullString
}

// 保存数据
func (user *UserModel) Save() int64 {
	result, e := initDB.Db.Exec("insert into gin.user (email, password) values (?,?);", user.Email, user.Password)
	if e != nil {
		log.Panicln("user insert error", e.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}
	return id
}

// 查询
func (user *UserModel) QueryByEmail() UserModel {
	temp := UserModel{}
	row := initDB.Db.QueryRow("select user.id,user.email,user.password from gin.user where user.email= ?;", user.Email)
	err := row.Scan(&temp.Id,&temp.Email, &temp.Password)
	if err != nil {
		log.Panicf("读取失败:%s", err.Error())
	}
	return temp
}

// 根据id 查找
func (user UserModel) QueryById(id int) (UserModel, error) {
	empty := UserModel{}
	row := initDB.Db.QueryRow("select * from gin.user where user.id=?;", id)
	err := row.Scan(&empty.Id, &empty.Email, &empty.Password, &empty.Avatar)
	if err != nil {
		log.Panicf("扫描失败")
	}
	return empty, err

}

func (user *UserModel) Update(id int64) error {
	stmt, err := initDB.Db.Prepare("update user set user.password=?,user.avatar=? where user.id=?")
	if err != nil {
		log.Panicf("更新sql发生错误:%s",err.Error())
	}
	_, err = stmt.Exec(user.Password, user.Avatar.String, user.Id)
	if err != nil {
		log.Panicf("更新失败:%s",err.Error())
	}
	return err

}
/**
func (user *UserModel) Update(id int) error {
	var stmt, e = initDB.Db.Prepare("update user set password=?,avatar=?  where id=? ")
	if e != nil {
		log.Panicln("发生了错误", e.Error())
	}
	_, e = stmt.Exec(user.Password, user.Avatar.String, user.Id)
	if e != nil {
		log.Panicln("错误 e", e.Error())
	}

	return e
}
 */