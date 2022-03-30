package model

const TableDictWordSentence = "dict_word_sentence"

// DictWordSentenceModel 单词对应的例句表
type DictWordSentenceModel struct {
	ID       int    `gorm:"type:int(11); primaryKey; autoIncrement; unsigned; not null" json:"id"`
	WID      int    `gorm:"type:int(11);unsigned;not null" json:"wid"`
	Sentence string `gorm:"type:varchar(255);not null" json:"sentence"`
	TransCN  string `gorm:"type:varchar(255);not null" json:"trans_cn"`
}
