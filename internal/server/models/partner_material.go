package models

type PartnerMaterial struct {
	Id         int64 `gorm:"primaryKey"`
	MaterialID int64
	Material   Material `gorm:"foreignKey:MaterialID"`
	PartnerID  int64
	Partner    Partner `gorm:"foreignKey:PartnerID"`
}

func (pm *PartnerMaterial) TableName() string {
	return "partner_materials"
}
