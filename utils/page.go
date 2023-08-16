package utils

func GetPageOffset(pageSize, pageNum int) int {
	if pageNum <= 0 {
		pageNum = 1
	}
	return (pageNum - 1) * pageSize
}
