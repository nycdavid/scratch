package main

import (
	"log"
	"net/http"
)

func main() {
	_1_unintended_variable_shadowing(true)
}

func _1_unintended_variable_shadowing(tracing bool) error {
	httpClient := func() (*http.Client, error) {
		return &http.Client{}, nil
	}

	var client *http.Client

	if tracing {
		client, err := httpClient()
		if err != nil {
			return err
		}
		log.Println(client)
	} else {
		client, err := httpClient()
		if err != nil {
			return err
		}
		log.Println(client)
	}

	log.Println(client)
	return nil

	/*
		The `log.Println` in either the if or the else clause prints out a valid http.Client object
		The `log.Println` on line 30 prints out nil.
	*/
}
