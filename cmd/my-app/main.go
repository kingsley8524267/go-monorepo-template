package main

import "go-monorepo-template/apps/my-app"

func main() {
	app, err := my_app.New()
	if err != nil {
		panic(err)
	}

	app.Run()
}
