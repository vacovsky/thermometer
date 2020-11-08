package main

import (
	"log"
	"sync"
	"time"

	"github.com/jinzhu/gorm"

	// This is how the documentation indicated to do it.
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ORM struct {
	DB   *gorm.DB
	lock sync.Mutex
}

var instance *ORM
var once sync.Once

// DataService returns data service in a singleton
func DataService() *ORM {
	once.Do(func() {
		instance = &ORM{
			lock: sync.Mutex{},
		}
		instance.setup()
	})
	return instance
}

func (d *ORM) sqliteConnect() {
	var err error

	d.DB, err = gorm.Open("sqlite3", "greenhouse.db")

	if err != nil {
		log.Println(err)
	}
}

func (d *ORM) setup() {
	d.sqliteConnect()
	d.DB.AutoMigrate(
		&TempHumData{},
	)
}

// TempHumData for storing stuff
type TempHumData struct {
	// ID          int `gorm:"primaryKey;autoIncrement:true"`
	Timestamp   time.Time
	Temperature float64
	Humidity    float64
}

type Chart struct {
	Series []string
	Labels []int64
	Data   [][]string
}
