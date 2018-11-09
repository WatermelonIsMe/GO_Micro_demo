package db

import (
	"database/sql"
	"user-srv/entity"
)

// user表的查
func SelectUserById(id int32) (*entity.User, error) {

	user := new(entity.User)
	if err := db.Get(user, " SELECT name,address,phone FROM user WHERE id = ?", id); err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		return nil, nil
	}
	return user, nil
}

// user表的增
func InsertUser(user *entity.User) (int64,error) {

	rep, err := db.Exec("INSERT INTO `user`(`name`,`address`,`phone`) VALUE (?,?,?)",user.Name,user.Address,user.Phone)

	if err != nil{
		return 0,err
	}

	return rep.LastInsertId()
}

// user表的改
func ModifyUser(user *entity.User) (error) {

	_, err := db.Exec("UPDATE  `user` set `name` = ? ,`phone` = ?, `address` = ? where `id` = ?",user.Name,user.Phone,user.Address,user.Id)

	if err != nil{
		return err
	}

	return nil
}

// user表的删除
func DeleteUser(id int32) (error){

	_,err := db.Exec("DELETE FROM `user` WHERE `id` = ?",id)
	if err != nil{
		return  err
	}

	return nil
}