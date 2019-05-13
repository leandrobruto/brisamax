package main

import (
	"brisamax/app/controllers/programa"
	"os"
)

func main() {
	a := programa.App{}
	a.Initialize(
		os.Getenv("postgres"),
		os.Getenv("admin"),
		os.Getenv("brisamax"))

	a.Run(":8000")
}
