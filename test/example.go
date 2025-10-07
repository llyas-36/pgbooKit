package main

// import (
// 	"context"
// 	pgbookit "pgbooKit"
// )

// func GetUserByEmail(email string) (string, error) {
// 	dbpool := pgbookit.DB()
// 	var name string
// 	err := dbpool.QueryRow(context.Background(), "selet name from users where email=$1", email).Scan(&name)
// 	if err != nil{
// 		return "", err
// 	}

// 	return name, err
// }

// func StoreUserToDB(name, email, password string) error{
// 	dbpool := pgbookit.DB()
// 	_, err := dbpool.Exec(context.Background(), "INSERT INTO users(name, email, password) values($1, $2, $3)", name, email, password)

// 	if err != nil{
// 		return err
// 	}

// 	return nil
// }
