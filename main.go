package main

import "github.com/bhoriuchi/terraform-remote-backend/pkg/command"

func main() {
	if err := command.NewCommand().Execute(); err != nil {
		panic(err)
	}
}
