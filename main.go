package main

import (
	"github.com/devKiratu/gator/internal/config"
)

func main(){
	config.StartGator()
	// data, err := config.Read()
	// if err != nil {
	// 	fmt.Println("Error reading file: %w", err)
	// }
	// state := config.InitState(&data)

	// args := os.Args
	// if len(args) < 3 {
	// 	fmt.Println("not enough arguments were provided")
	// 	os.Exit(1)
	// }

	// currentCommand := args[1]
	// currentArg := args[2]
	// commands := config.GetCommands()
	// cmd, ok := commands.Data[currentCommand]
	// if  !ok {
	// 	fmt.Println("unkown command")
	// 	os.Exit(1)
	// }
	// cmd.run





	
	// fmt.Println("Found some data: ")
	// fmt.Printf("Current user: %s\n", data.CurrentUserName)
	// fmt.Printf("Database Url: %s\n", data.DbUrl)
	// // data.SetUser("tom")
	// fmt.Println("==============Read after Write==================")
	// 	data, err = config.Read()
	// if err != nil {
	// 	fmt.Println("Error reading file: %w", err)
	// }
	// fmt.Println("Found some data: ")
	// fmt.Printf("Current user: %+v\n", data)
	// fmt.Printf("Database Url: %s\n", data.DbUrl)

}
