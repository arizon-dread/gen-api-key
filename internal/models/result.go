// Package models has the structs for the CLI app.
package models

type Result struct {
	APIKey  string `json:"api-key"`
	Hash    string `json:"hash"`
	Prefix  string `json:"prefix,omitempty"`
	AppName string `json:"app-name,omitempty"`
	Env     string `json:"env,omitempty"`
}
