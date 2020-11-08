package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var buffer = make(map[int][]byte)
var frame []byte
var mutex = &sync.Mutex{}

// Serve serves the server serve
func Serve() {
	port := os.Getenv("GREENHOUSE_WEB_PORT")
	if port == "" {
		port = "3000"
	}

	host := "0.0.0.0:" + port

	// http.HandleFunc("/data", getData)
	http.HandleFunc("/tempchartdata", getTemperatureChartData)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/"+r.URL.Path[1:])
	})
	log.Fatal(http.ListenAndServe(host, nil))
}

func getTemperatureChartData(w http.ResponseWriter, r *http.Request) {
	ep := GetTemperaturesChartData()
	js, err := json.Marshal(ep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// GetTemperaturesChartData stuff
func GetTemperaturesChartData() Chart {
	chartHours := 96
	rawdata := []TempHumData{}
	labels := []int64{}

	ambientTempStr := []string{}
	ambientHumStr := []string{}

	DataService().DB.Where("timestamp > ?",
		time.Now().Add(-time.Duration(chartHours)*time.Hour)).Find(&rawdata).Order("timestamp desc")

	var curAT float64
	var curAH float64

	for _, v := range rawdata {
		labels = append(labels, v.Timestamp.Unix()*1000)

		// remove dud readings from the chart to smooth it out
		if v.Humidity != 0 {
			curAH = v.Humidity
		}
		if v.Temperature != 32 {
			curAT = v.Temperature
		}

		ambientTempStr = append(ambientTempStr, fmt.Sprintf("%.2f", curAT))
		ambientHumStr = append(ambientHumStr, fmt.Sprintf("%.0f", curAH))

	}
	chart := Chart{
		Series: []string{"Ambient Air Temperature", "Ambient Humidity"},
		Labels: labels,
		Data:   [][]string{ambientTempStr, ambientHumStr},
	}
	return chart
}
