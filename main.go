package main

import "mitton/gin-handson/router"

func main() {
	r := router.SetupRouter()

	r.Run(":3000")
}
