package form

type Pagination struct {
	Keyword string `query:"keyword"`
	Page    int    `query:"page"`
	Limit   int    `query:"limit"`
}
