package bd

import (
	"fmt"
	"initial/models"
	"initial/tools"

	_ "github.com/go-sql-driver/mysql"
)

func SingUp(sig models.SingUp) error {
	fmt.Println("Starting to register on Db")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	query := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.DateMySQL() + "')"
	fmt.Println(query)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println("Error executing query: " + err.Error())
		return err
	}
	return nil
}
