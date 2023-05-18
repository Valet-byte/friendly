package model

type Point struct {
	UserId        string  `json:"user_id"`
	X             float32 `json:"x"`
	Y             float32 `json:"y"`
	BatteryCharge int8    `json:"battery_charge"`
	Speed         float32 `json:"speed"`
	IsInHome      bool    `json:"is_in_home"`
}
