package models

type Customer struct {
	Id          int64   `gorm:"primaryKey"`
	Material    string  `gorm:"column:material"`
	Latitude    float64 `gorm:"column:latitude; type:float"`
	Longitude   float64 `gorm:"column:longitude; type:float"`
	FloorArea   float32 `json:"floor_area; type:float"`
	PhoneNumber string  `json:"phone_number; type:TEXT"`
}

func (c *Customer) TableName() string {
	return "customers"
}
