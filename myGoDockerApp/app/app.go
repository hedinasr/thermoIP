package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"bufio"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"strconv"
	"encoding/json"
	"time"
	"math/rand"
)

type Temperature struct {
	Value float64
	Unit string
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "This is the RESTful api !!!!!")
	fmt.Fprintf(w, "Go to /temp if you want the current temperature")
	fmt.Fprintf(w, "Go to /lum if you want the current luminosity")
}

func tempHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := ioutil.ReadFile("/tmp/temp.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(data))
}

func lumHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Feature not yet implemented !")
}

// http://aliafshar.github.io/blog/posts/2013-07-26-json-http-client-golang.html
func getTemp() Temperature {
	resp, err := http.Get("http://192.168.1.177")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	temp := Temperature{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&temp)

	if err != nil {
		log.Println(err)
	}

	return temp
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	c := make(chan float64)
	f, err := os.Create("/tmp/temp.txt")

	check(err)

	defer f.Close()

	go func() {
		for i := 1; i < 10; i++ {
			//currentTemp := getTemp()
			//c <- currentTemp.Value
			c <- rand.Float64() * 10
			time.Sleep(time.Second)
		}
	}()

	go func() {
		w := bufio.NewWriter(f)
		for {
			temp := <-c
			_, err := w.WriteString(FloatToString(temp) + "\n")
			check(err)
			w.Flush()
		}
	}()

	router := httprouter.New()
	router.GET("/", indexHandler)
	router.GET("/temp", tempHandler)
	router.GET("/lum", lumHandler)

	// print env
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	http.ListenAndServe(":8080", router)
}
