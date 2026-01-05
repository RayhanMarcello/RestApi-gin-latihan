package entity

type Model struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}