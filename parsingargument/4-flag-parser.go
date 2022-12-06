package main

import (
	"fmt"
	"os"

	//"github.com/labstack/echo"
	"gopkg.in/alecthomas/kingpin.v2"
)


var app = kingpin.New("App", "Simple App")
var commandSomething = app.Command("something", "do something")
var commandSomethingArgX = commandSomething.Flag("x", "arg x").String()
var commandSomethingFlagY = commandSomething.Flag("y", "flag y").String()

// Add
var commandAdd = app.Command("add", "add new user")
var commandAddFlagOverride = commandAdd.Flag("override", "override existing user").Short('o').Bool()
var commandAddArgUser = commandAdd.Arg("user", "username").Required().String()

// Update
var commandUpdate = app.Command("update", "update user")
var commandUpdateOldUser = commandUpdate.Arg("old", "old username").Required().String()
var commandUpdateArgNewUser = commandUpdate.Arg("new", "new username").Required().String()

// delete
var commandDelete = app.Command("delete", "delete user")
var commandDeleteFlagForce = commandDelete.Flag("force", "force deletion").Short('f').Bool()
var commandDeleteArgUser = commandDelete.Arg("user", "username").Required().String()

func main() {
	commandAdd.Action(func(ctx *kingpin.ParseContext) error {
		// code here
		user := *commandAddArgUser
		override := *commandAddFlagOverride
		fmt.Printf("adding user %s, override %t \n", user, override)


		return nil
	})

	commandUpdate.Action(func(ctx *kingpin.ParseContext) error {
		// code here
		oldUser := *commandUpdateOldUser
		newUser := *commandUpdateArgNewUser
		fmt.Printf("updating user %s %s \n", oldUser, newUser)

		return nil
	})

	commandDelete.Action(func(ctx *kingpin.ParseContext) error {
		// code here
		user := *commandDeleteArgUser
		force := *commandDeleteFlagForce
		fmt.Printf("Deleting user %s, force %t \n", user, force)

		return nil
	})

	kingpin.MustParse(app.Parse(os.Args[1:]))
}