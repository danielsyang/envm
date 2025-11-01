package main

import (
	controller "envm/controller"
	db "envm/database"
)

func main() {
	db.InitDatabase()

	controller.Execute()

}
