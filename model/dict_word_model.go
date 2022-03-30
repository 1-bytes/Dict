package model

const TableDictWord = "dict_word"

// DictWordModel 单词表
type DictWordModel struct {
	ID       int    `gorm:"type:int(11); primaryKey; autoIncrement; unsigned; not null" json:"id"`
	ParentID int    `gorm:"type:int(11);unsigned;not null" json:"parent_id"`
	Word     string `gorm:"type:varchar(255);not null" json:"word"`
	PronEN   string `gorm:"type:varchar(255);not null" json:"pron_en"`
	PronAM   string `gorm:"type:varchar(255);not null" json:"pron_am"`
	Type     int    `gorm:"type:int(11);unsigned;not null" json:"type"`
}
