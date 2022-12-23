package function

import (
	"StudyDemo/LoginAndRegister/utils"
	"StudyDemo/LoginRegisterCheck/modules"
	"log"
)

//func CheckUserByNameAndPassword(username string, password string) (*store.User, error) {
//	sqlStr := "select id,username,password,email from users where username = ? and password = ?"
//	row := utils.Db.QueryRow(sqlStr, username, password)
//	user := &store.User{}
//	err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
//	if err != nil {
//		log.Println("数据库数据拿取成功")
//	}
//	return user, nil
//}

func CheckUserByName(name string) (*modules.User, error) {
	sqlStr := "select id,name,password from usertable where name = ?"
	row := utils.Db.QueryRow(sqlStr, name)
	user := &modules.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		log.Println(err)
	}
	return user, nil
}

func CheckUserByID(ID uint) (*modules.User, error) {
	sqlStr := "select id,name,password from usertable where id = ?"
	row := utils.Db.QueryRow(sqlStr, ID)
	user := &modules.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		log.Println(err)
	}
	return user, nil
}

//func CheckUserByName(username string) (*store.User, error) {
//	sqlStr := "select id,username,password,email from users where username = ?"
//	row := utils.Db.QueryRow(sqlStr, username)
//	user := &store.User{}
//	err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
//	if err != nil {
//		log.Println("密码或者用户名错误")
//	}
//	return user, nil
//}

func SaveUser(name string, password string) error {
	sqlStr := "insert into usertable(name,password) values (?,?)"
	_, err := utils.Db.Exec(sqlStr, name, password)
	if err != nil {
		return err
	}
	return nil
}
