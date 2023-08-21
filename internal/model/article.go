package model

type ArticleContent struct {
	AutoId  int    `gorm:"primarykey" db:"auto_id" json:"auto_id"`
	Id      string `gorm:"unique;not null" db:"id" json:"id"`
	Content string `gorm:"not null" db:"content" json:"content"`
}

func (m *ArticleContent) TableName() string {
	return "articles_content"
}

type Articles struct {
	Id           int    `gorm:"primarykey" db:"id" json:"id"`
	Title        string `gorm:"not null" db:"title" json:"title"`
	DiyDate      string `gorm:"not null" db:"diy_date" json:"diy_date"`
	DiyQihao     string `gorm:"not null" db:"diy_qihao" json:"diy_qihao"`
	DiyData      string `gorm:"not null" db:"diy_data" json:"diy_data"`
	DiyShengxiao string `gorm:"not null" db:"diy_shengxiao" json:"diy_shengxiao"`
	DiyTema      string `gorm:"not null" db:"diy_tema" json:"diy_tema"`
	// CreatedAt    time.Time `gorm:"not null" db:"created_at" json:"created_at"`
	PublishTime int64  `gorm:"not null" db:"publish_time" json:"publish_time"`
	ImgUrl      string `gorm:"not null" db:"img_url" json:"img_url"`
	Cid         string `gorm:"not null" db:"cid" json:"cid"`
}

const (
	ArticleTypeShengXiao = "shengxiao"
	AritcleTypeNumber    = "number"
)
