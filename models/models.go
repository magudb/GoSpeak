package models

type Qoute struct {
	ID     int
	Text   string
	Author string
}

func (this *Qoute) TableName() string {
	return "quotes"
}
