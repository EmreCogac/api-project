package models

type Emote struct {
	Eid   uint   `gorm:"primaryKey"`
	Emote string `gorm:"type:varchar(30)"`
}
