package models

type Material struct {
	Id   int64  `gorm:"primaryKey"`
	Name string `json:"column:name; UNIQUE"`
}

func (m *Material) TableName() string {
	return "materials"
}
