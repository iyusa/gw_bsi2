package common

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("Initializing rand seed ...")
	rand.Seed(time.Now().UnixNano())
}

// PrintJSON print s as json
func PrintJSON(title string, s interface{}) error {
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n%v\n", title, string(b))
	return nil
}

// CurrentTime in HHMMSS
func CurrentTime() string {
	dt := time.Now()
	return dt.Format("150405")
}

// TimeStamp time stamp in finnet format
func TimeStamp() string {
	dt := time.Now()
	return dt.Format("02-01-2006 15:04:05")
}

const letters = "1234567890"

// const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// RandomString generate random string
func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// JSON return json string
func JSON(s interface{}) string {
	b, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(b)
}
