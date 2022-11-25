package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func Time(w http.ResponseWriter, req *http.Request) {
	time := time.Now()
	fmt.Fprintf(w, "%dh%d", time.Hour(), time.Minute())
}

func Dice(w http.ResponseWriter, req *http.Request) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Fprintf(w, "%04d", r1.Intn(1000)+1)
}

func Dices(w http.ResponseWriter, req *http.Request) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	t := 0
	i := 0
	face := 0
	var faces = [8]int{2, 4, 6, 8, 10, 12, 20, 100}
	style := req.URL.Query().Get("style")
	face, _ = strconv.Atoi(style)
	if face != 0 {
		t = 1
		face, _ = strconv.Atoi(style)
	}
	for i < 15 {
		if t == 0 {
			face = faces[r1.Intn(8)]
		}
		switch {
		case face == 2 || face == 4 || face == 6 || face == 8:
			fmt.Fprintf(w, "%d ", r1.Intn(face)+1)
		case face == 10 || face == 12 || face == 20:
			fmt.Fprintf(w, "%02d ", r1.Intn(face)+1)
		case face == 100:
			fmt.Fprintf(w, "%03d ", r1.Intn(face)+1)
		}
		i++
	}
}

func main() {
	http.HandleFunc("/", Time)
	http.HandleFunc("/dice", Dice)
	http.HandleFunc("/dices", Dices)
	http.ListenAndServe(":4567", nil)
}
