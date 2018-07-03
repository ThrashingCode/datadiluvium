package main

import "go/types"

type DataStore struct {
	DataStore  string      `json:"datastore"`
	Connection string      `json:"connection"`
	Generate   types.Array `json:"generate"`
	GenRows    int         `json:"genrows"`
}

type DataStores struct {
	Collection []DataStore
}
