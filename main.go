package main

import di "github.com/goforbroke/practice-di/internal"

func main() {
	place := "tokyo"
	uf := di.NewContainer()
	uf.UsecaseQuery(place)
}
