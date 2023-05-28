package function

import (
	"StudyDemo/LoginRegisterCheck/modules"
	"StudyDemo/LoginRegisterCheck/utils"
	"log"
)

func CheckUserByName(name string) (*modules.User, error) {
	sqlStr := "select id,name,password from usertable where name = ?"
	row := utils.Db.QueryRow(sqlStr, name)
	user := &modules.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		return user, err
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

func SaveUser(name string, password string) error {
	sqlStr := "insert into usertable(name,password) values (?,?)"
	_, err := utils.Db.Exec(sqlStr, name, password)
	if err != nil {
		return err
	}
	return nil
}
