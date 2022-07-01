package models

type Partner struct {
	Id        int64   `gorm:"primaryKey"`
	Name      string  `gorm:"column:name; type:TEXT"`
	Latitude  float64 `gorm:"column:latitude; type:float"`
	Longitude float64 `gorm:"column:longitude; type:float"`
	Radius    int     `gorm:"column:radius; type:int;"`
	Rating    float32 `gorm:"column:rating; type:NUMERIC(2,1)"`
	Geom      string  `gorm:"->;type:geography GENERATED ALWAYS AS (ST_MakePoint(longitude,latitude, 4326)::geography) STORED"`
}

func (p *Partner) TableName() string {
	return "partners"
}
