package domain

type Product struct {
	ID          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `gorm:"size:255;not null;" json:"name"`
	Category    string `gorm:"size:100;not null;" json:"category"`
	Description string `gorm:"size:255;not null;" json:"description"`
}
