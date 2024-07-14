package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//Save and load.

func SaveTasks(filename string) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadTasks(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil 
		}
		return err
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return err
	}
	return nil
}
