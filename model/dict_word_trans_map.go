package model

const TableDictWordTransMap = "dict_word_trans_map"

// DictWordTransMapModel 中英文对照表
type DictWordTransMapModel struct {
	ID      int    `gorm:"type:int(11); primaryKey; autoIncrement; unsigned; not null" json:"id"`
	Wid     int    `gorm:"type:int(11);unsigned;not null" json:"wid"`
	TransCN string `gorm:"type:varchar(255);not null" json:"trans_cn"`
	Pos     string `gorm:"type:varchar(255);not null" json:"pos"`
	Grammar string `gorm:"type:varchar(255);not null" json:"grammar"`
}
