package config

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", Health).Methods("GET")
	router.HandleFunc("/advisory", BartServiceAdvisoryHandler).Methods("GET")
	router.HandleFunc("/advisory/count", CountHandler).Methods("GET")
	router.HandleFunc("/advisory/elevator", ElevatorHandler).Methods("GET")
	router.HandleFunc("/realtime", RealTimeHandler).Methods("GET")
	router.HandleFunc("/route/info", RouteInfoHandler).Methods("GET")
	router.HandleFunc("/route", RoutesHandler).Methods("GET")
	return router
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Health Check"))
}

//Adivisory API Handlers
func BartServiceAdvisoryHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://api.bart.gov/api/bsa.aspx?cmd=bsa&key=MW9S-E7SL-26DU-VV8V&json=y"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func CountHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://api.bart.gov/api/bsa.aspx?cmd=count&key=MW9S-E7SL-26DU-VV8V&json=y"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func ElevatorHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://api.bart.gov/api/bsa.aspx?cmd=elev&key=MW9S-E7SL-26DU-VV8V&json=y"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "033c8163-0e70-4a90-8875-0cd3e3f54994")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func HelpAPIHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://api.bart.gov/api/bsa.aspx?cmd=help&key=MW9S-E7SL-26DU-VV8V&json=y"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "3888abba-c799-4cde-acb4-d58ab57ca687")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

//Real Time Information - API

func RealTimeHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://api.bart.gov/api/etd.aspx?cmd=etd&orig=24th&key=MW9S-E7SL-26DU-VV8V&json=y"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "528894ff-6019-4467-9a84-9efacb23cca5")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

//Route Information - API

func RoutesHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://api.bart.gov/api/route.aspx?cmd=routes&key=MW9S-E7SL-26DU-VV8V&json=y"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "e6b4fff7-18d7-4980-8f9e-d46d5c207e89")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func RouteInfoHandler(w http.ResponseWriter, r *http.Request) {

	url := "http://api.bart.gov/api/route.aspx?cmd=routeinfo&route=9&key=MW9S-E7SL-26DU-VV8V&json=y"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "5e7d1e78-169f-43b1-83dc-367955455b70")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// Schedule Information API
