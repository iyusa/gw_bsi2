package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // sengaja
	"github.com/ussidev/bsi2/common"
)

var db *gorm.DB

// Initialize database
func Initialize() {
	fmt.Printf("Will initialize model: %s ...\n", common.Config.DB)
	if db != nil {
		return
	}

	db, err := gorm.Open("mysql", common.Config.DB)
	if err != nil {
		log.Printf("%v\n", err)
		panic(err)
	}

	db.LogMode(false)
	fmt.Printf("model initialized db is %v...\n", db)
}

func Reconnect() (err error) {
	if db == nil {
		db, err = gorm.Open("mysql", common.Config.DB)
		if err != nil {
			return
		}

		if db == nil {
			err = fmt.Errorf("unknown db error")
			return
		}
	}
	return
}
func Close() {
	if db != nil {
		db.Close()
		db = nil
	}
}
