package main

import (
	"encoding/json"
	"fmt"
	mux "github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	fmt.Fprintf(w, "<h1 style=\"font-family: Helvetica;\">Hello, welcome to track service</h1>")
}

func TrackIndex(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	tracks := FindAllTracks()

	if err := json.NewEncoder(w).Encode(tracks); err != nil {
		panic(err)
	}
}

func TrackShow(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	id, err := strconv.Atoi(ps.ByName("trackId"))
	HandleError(err)
	track := FindTrack(id)
	if err := json.NewEncoder(w).Encode(track); err != nil {
		panic(err)
	}

}

func TrackCreate(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var track Track
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	HandleError(err)
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Save JSON to Track struct
	if err := json.Unmarshal(body, &track); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	CreateTrack(track)
	w.Header().Set("Content-Type", "application/json: charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
