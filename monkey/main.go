package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

// type User struct {
// 	Name  string
// 	Email *string
// }

// func (u User) String() string {
// 	email := ""
// 	if u.Email != nil {
// 		email = *u.Email
// 	}
// 	return fmt.Sprintf(`Name:"%s", Email:"%s"`, u.Name, email)
// }

// func To[T any](v T) *T {
// 	return &v
// }

func main() {
	// u := User{Name: "Chad"}
	// fmt.Println(u)

	// u.Email = To("a@b.c")
	// fmt.Println(u)

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
