package main

import (
	"log"
	api "test/kitex_gen/api/test"
)

func main() {
	svr := api.NewServer(new(TestImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
