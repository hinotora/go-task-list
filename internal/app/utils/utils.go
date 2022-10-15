package utils

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

var Console_delimeter string = "======================================"
var Data_folder string = ".tasklist"
var Data_registry_file string = "registry"

func StoreData(filename string, data interface{}) {
	dataFile, _ := os.Create(fmt.Sprintf("%s/%s", Prepare_dir(), filename))

	dataEncoder := gob.NewEncoder(dataFile)
	dataEncoder.Encode(data)

	dataFile.Close()
}

func Prepare_dir() string {
	home_dir, _ := os.UserHomeDir()

	data_dir := fmt.Sprintf("%s/%s", home_dir, Data_folder)

	_, err := os.Stat(data_dir)

	if os.IsNotExist(err) {
		os.MkdirAll(data_dir, 0777)
	}

	return data_dir
}

func GetClearDateFromUNIX(unix_ts int64) string {
	return time.Unix(unix_ts, 0).Format("02.01.2006 15:04")
}
