/* MIT License
#
# Copyright (c) 2020 Ferhat Geçdoğan All Rights Reserved.
# Distributed under the terms of the MIT License.
#
# */

package main

import (
	"fmt"
	"os"
	"os/user"
	"os/exec"
    	"strings"
)

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


func ListDir() {
    file, err := os.Open(".")
    if err != nil {
        fmt.Println("failed opening directory: %s", err)
    }
    defer file.Close()
 
    list,_ := file.Readdirnames(0)
    for _, name := range list {
        fmt.Println(name)
    }
}


func Color(colorString string) func(...interface{}) string {
  sprint := func(args ...interface{}) string {
    return fmt.Sprintf(colorString,
      fmt.Sprint(args...))
  }
  return sprint
}


func GetInput(input string) string {
	fmt.Print(" ")
	fmt.Scanln(&input)
	return input
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

func execInput(input string) error {
    input = strings.TrimSuffix(input, "\n")
    cmd := exec.Command(input)
    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout
    return cmd.Run()
}
func main() {
	var input string
    	fmt.Println("Welcome", Purple(GetUsername()), "^-^")
    	for {
    	    	GetTerminal()
    		input = GetInput(input)
    		if input == "help" {
    		} else if input == "printlnf" {
    		} else if input == "ls" {
    			ListDir()
    		} else if input == "cls" {
    			os.Exit(3)
    		} else if input == "fr" || input == "cd" {
    		} else {
    			execInput(input)
    		}
    		input = ""
    	}
}

