// Package models has the structs for the CLI app.
package models

type Result struct {
	APIKey string `json:"api-key"`
	Hash string `json:"hash"`
	Env string `json:"env"`
	AppName string `json:"app-name"`
}
