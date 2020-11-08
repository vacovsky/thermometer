package main

import (
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/vacovsky/greef/sensors/dht11"
)

var (
	storage = TempHumData{}
	mux     = sync.Mutex{}
)

func main() {

	go monitorAmbientAir()
	time.Sleep(time.Second * 2)
	go Serve()

	db := DataService().DB

	for {
		spew.Dump(storage)

		db.Save(&storage)
		time.Sleep(time.Minute * 5)
	}

}

func monitorAmbientAir() {
	airPin := 4
	for {
		hum, temp := dht11.GetAmbientInfoF(airPin)
		mux.Lock()
		storage.Humidity, storage.Temperature = hum, temp
		storage.Timestamp = time.Now()
		mux.Unlock()
		time.Sleep(time.Minute * 5)
	}
}
