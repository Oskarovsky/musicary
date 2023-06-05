package main

import "time"

type Artist struct {
	name string
}

type Track struct {
	title string
	Artist
	year    int
	tempo   int
	version string
	genre   string
}

type Mix struct {
	tracks    []Track
	artist    Artist
	timestamp time.Time
}

func main() {

}
