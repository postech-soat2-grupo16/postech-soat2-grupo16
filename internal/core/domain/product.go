package domain

type Product struct {
	ID          uint32 `gorm:"primary_key;auto_increment"`
	Name        string `gorm:"size:255;not null;"`
	Category    string `gorm:"size:100;not null;"`
	Description string `gorm:"size:255;not null;"`
}

func (p *Product) IsCategoryValid() bool {
	return false
}
