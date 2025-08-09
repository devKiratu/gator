package main

import (
	"fmt"

	"github.com/devKiratu/gator/internal/config"
)

func main(){
	data, err := config.Read()
	if err != nil {
		fmt.Println("Error reading file: %w", err)
	}
	fmt.Println("Found some data: ")
	fmt.Printf("Current user: %s\n", data.CurrentUserName)
	fmt.Printf("Database Url: %s\n", data.DbUrl)
	data.SetUser("tom")
	fmt.Println("==============Read after Write==================")
		data, err = config.Read()
	if err != nil {
		fmt.Println("Error reading file: %w", err)
	}
	fmt.Println("Found some data: ")
	fmt.Printf("Current user: %s\n", data.CurrentUserName)
	fmt.Printf("Database Url: %s\n", data.DbUrl)

}
