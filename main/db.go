package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

var currentTrackId int
var currentArtistId int

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	HandleError(err)
	return c
}

// Give us some seed data
func init() {
	CreateTrack(Track{
		Artist: Artist{
			Name: "DJ Hazel",
		},
		Title:   "I love Poland",
		Year:    2010,
		Tempo:   128,
		Genre:   "Club",
		Version: "Original Mix",
	})

	CreateTrack(Track{
		Artist: Artist{
			Name: "Tiesto",
		},
		Title:   "Traffic",
		Year:    2005,
		Tempo:   140,
		Genre:   "Trance",
		Version: "Radio Mix",
	})

}

func FindAllTracks() Tracks {
	var tracks Tracks
	c := RedisConnect()
	defer func(c redis.Conn) {
		err := c.Close()
		if err != nil {
			fmt.Println("Error occurred for FindAllTracks")
		}
	}(c)

	keys, err := c.Do("KEYS", "post:*")
	HandleError(err)

	for _, k := range keys.([]interface{}) {
		var track Track
		reply, err := c.Do("GET", k.([]byte))
		HandleError(err)
		if err := json.Unmarshal(reply.([]byte), &track); err != nil {
			panic(err)
		}
		tracks = append(tracks, track)

	}
	return tracks
}

func FindTrack(id int) Track {
	var track Track
	c := RedisConnect()
	defer func(c redis.Conn) {
		err := c.Close()
		if err != nil {
			fmt.Println("Error occurred for FindTrack")
		}
	}(c)
	reply, err := c.Do("GET", "post:"+strconv.Itoa(id))
	HandleError(err)
	fmt.Println("GET OK")

	if err = json.Unmarshal(reply.([]byte), &track); err != nil {
		panic(err)
	}
	return track
}

func CreateTrack(t Track) {
	currentTrackId += 1
	currentArtistId += 1

	t.Id = currentTrackId
	t.Artist.Id = currentArtistId

	c := RedisConnect()
	defer func(c redis.Conn) {
		err := c.Close()
		if err != nil {
			fmt.Println("Error occurred for CreateTrack")
		}
	}(c)

	b, err := json.Marshal(t)
	HandleError(err)

	// Save JSON blob to Redis
	reply, err := c.Do("SET", "post:"+strconv.Itoa(t.Id), b)
	HandleError(err)
	fmt.Println("GET ", reply)
}
