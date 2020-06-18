package main

import (
	"fmt"
	"os"
	"os/user"
)

var test = ""

// Color definitions
var (
  Black   = Color("\033[1;30m%s\033[0m")
  Red     = Color("\033[1;31m%s\033[0m")
  Green   = Color("\033[1;32m%s\033[0m")
  Yellow  = Color("\033[1;33m%s\033[0m")
  Purple  = Color("\033[1;34m%s\033[0m")
  Magenta = Color("\033[1;35m%s\033[0m")
  Teal    = Color("\033[1;36m%s\033[0m")
  White   = Color("\033[1;37m%s\033[0m")
)


func Color(colorString string) func(...interface{}) string {
  sprint := func(args ...interface{}) string {
    return fmt.Sprintf(colorString,
      fmt.Sprint(args...))
  }
  return sprint
}


func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return name
}

func GetUsername() string {
	user, err := user.Current()
	if err!= nil {
		panic(err)
	}
	
	return user.Username
}

func GetPath() string {
	user, err := user.Current()
	if err!= nil {
		panic(err)
	}
	return user.HomeDir
}

func GetTerminal() {
	fmt.Print(Teal(GetUsername()))
	fmt.Print(Green("@"))
	fmt.Print(Magenta(GetHostname()))
	fmt.Print(Teal(":~"))
	fmt.Print(Yellow(GetPath()))
	fmt.Print(" " + Red("\u25B6"))
}

func main() {
    fmt.Println("Welcome", Purple(GetUsername()), "^-^")
    GetTerminal()
}

