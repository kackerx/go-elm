package vars

type PageParams struct {
	PageSize int `form:"page_size" json:"page_size" binding:"min=1"`
	PageNum  int `form:"page_num" json:"page_num" binding:"min=1"`
}
