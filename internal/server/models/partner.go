package models

type Partner struct {
	Id        int64   `gorm:"primaryKey"`
	Latitude  float64 `gorm:"column:latitude; type:float"`
	Longitude float64 `gorm:"column:longitude; type:float"`
	Radius    int     `gorm:"column:radius; type:int;"`
	Rating    float32 `gorm:"column:rating; type:NUMERIC(2,1)"`
	Geom      string  `gorm:"->;type:geography; default:ST_GeomFromText('POINT(longitude latitude)', 4326)::geography; NOT NULL"`
}

func (p *Partner) TableName() string {
	return "partners"
}
