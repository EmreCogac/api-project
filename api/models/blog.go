package models

type Blog struct {
	Pid      uint   `gorm:"primaryKey"`
	Header   string `gorm:"type:varchar(30);not null"  json:"header"`
	Content1 string `gorm:"type:varchar(200);not null"  json:"content1"`
	Code1    string `gorm:"type:varchar(1000)"  json:"code1"`
	// Content2 *string `gorm:"type:varchar(400)" json:"content2"`
	// Code2    *string `gorm:"type:varchar(1000)"  json:"code2"`
	// Content3 *string `gorm:"type:varchar(800)" json:"content3"`
	// Code3    *string `gorm:"type:varchar(2000)"  json:"code3"`
	// Content4 *string `gorm:"type:varchar(1000)" json:"content4"`
	// Code4    *string `gorm:"type:varchar(5000)"  json:"code4"`
	Userid uint
	User   *User `gorm:"foreignKey:Userid"`
}

type Categories struct {
	Cateid   uint   `gorm:"primaryKey"`
	Catename string `gorm:"type:varchar(30)"`
}
type Catepostrel struct {
	Catepostid uint `gorm:"primaryKey"`
	Catfkpost  uint
	Postfkcat  uint
	Blog       Blog       `gorm:"foreignKey:Postfkcat"`
	Categories Categories `gorm:"foreignKey:Catfkpost"`
}
