package main

import "time"

type Artist struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Track struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Artist  Artist `json:"artist"`
	Year    int    `json:"year"`
	Tempo   int    `json:"tempo"`
	Version string `json:"version"`
	Genre   string `json:"genre"`
}

type Mix struct {
	Id        int       `json:"id"`
	Tracks    []Track   `json:"tracks"`
	Artist    Artist    `json:"artist"`
	Timestamp time.Time `json:"timestamp"`
}

type Tracks []Track
type Artists []Artist
type Mixes []Mix
