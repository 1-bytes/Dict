package config

// FaStarsModel fa_stars 数据表.
type FaStarsModel struct {
	ID                       int    `gorm:"type:int(11); primaryKey; autoIncrement"`
	CreatorId                uint64 `gorm:"type:bigint(20)"`
	CreatorName              string `gorm:"type:varchar(255)"`
	CreatorNickname          string `gorm:"type:varchar(255)"`
	ContactMails             string `gorm:"type:varchar(255)"`
	ContactWhatsapps         string `gorm:"type:varchar(255)"`
	ContactCountryCodes      string `gorm:"type:varchar(100)"`
	Avatar                   string `gorm:"type:text"`
	IsVerified               int    `gorm:"type:tinyint(2)"`
	Region                   string `gorm:"type:varchar(255)"`
	ProductCategories        string `gorm:"type:varchar(255)"`
	Whatsappswitch           int    `gorm:"type:int(11)"`
	Mailswitch               int    `gorm:"type:int(11)"`
	FollowerCnt              int    `gorm:"type:int(11)"`
	FollowerTopGender        int    `gorm:"type:int(11)"`
	FollowerTopGenderShare   string `gorm:"type:varchar(255)"`
	FollowerTopAgeGroup      int    `gorm:"type:int(11)"`
	FollowerTopAgeGroupShare string `gorm:"type:varchar(255)"`
	VideoAvgViewCnt          int    `gorm:"type:int(11)"`
	VideoPubCnt              int    `gorm:"type:int(11)"`
	ProductCnt               int    `gorm:"type:int(11)"`
	IsFirst                  int    `gorm:"type:(2)"`
	Createtime               int64  `gorm:"type:int(11)"`
	Updatetime               int64  `gorm:"type:int(11)"`
}
