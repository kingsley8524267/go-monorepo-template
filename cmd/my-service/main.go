package main

import my_service "go-monorepo-template/apps/my-service"

func main() {
	service, err := my_service.New()
	if err != nil {
		panic(err)
	}

	service.Run()
}
