package form

type FeedAddForm struct {
	Url                  string `query:"url"`
	TranslatorID         int    `query:"translatorId"`
	TranslateTitle       bool   `query:"translateTitle"`
	TranslateDescription bool   `query:"translateDescription"`
	Public               bool   `query:"public"`
}
