package main

import (
	"day4/config"
	"day4/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}

func TestMain(fail bool) {
	if e != nil {
		return fail
	}
	else {
		return not fail
	}
}
