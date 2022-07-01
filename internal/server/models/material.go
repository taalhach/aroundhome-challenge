package models

type Material struct {
	Id   int64  `gorm:"primaryKey"`
	Name string `gorm:"column:name;not null;unique"`
}

func (m *Material) TableName() string {
	return "materials"
}
