package entity

type Event struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
