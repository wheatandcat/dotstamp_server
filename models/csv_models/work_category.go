package csvModels

// WorkCategory body内のjson設計
type WorkCategory struct {
	ID   int
	Name string
}

func (w WorkCategory) getFileName() string {
	return "work_category.csv"
}
