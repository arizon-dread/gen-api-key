// Package models has the structs for the CLI app.
package models

type Result struct {
	APIKey  string `json:"api-key"`
	Hash    string `json:"hash"`
	Prefix  string `json:"prefix"`
	AppName string `json:"app-name"`
	Env     string `json:"env"`
}
