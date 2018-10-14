package models

type Event struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Date   string   `json:"date"`
	Photos []string `json:"photos"`
}
