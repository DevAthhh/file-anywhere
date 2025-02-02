package main

import "github.com/DevAthhh/fileanywhere/internal/handler"

func main() {
	router := handler.Handle()
	router.Run(":8000")
}
