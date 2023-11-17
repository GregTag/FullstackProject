package entity

type Genre struct {
	MediaID uint   `gorm:"primaryKey"`
	Genre   string `gorm:"primaryKey"`
}
