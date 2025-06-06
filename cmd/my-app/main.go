package main

import "go-monorepo-template/apps/my-app"

func main() {
	service, err := my_app.New()
	if err != nil {
		panic(err)
	}

	service.Run()
}
