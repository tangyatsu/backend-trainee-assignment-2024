package main

import (
	"backend-trainee-assignment-2024/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
	// TODO: init logger: slog

	// TODO: init storage: postgresql

	// TODO: init cache: redis

	// TODO: init router: chi, "chi render"

	// TODO: run server
}
