package main

type Config struct {
	Id      string            `json:"Id"`
	Entries map[string]string `json:"Entries"`
}
