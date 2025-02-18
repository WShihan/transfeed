package form

type FeedUpdateForm struct {
	ID                   int    `body:"id"`
	Url                  string `body:"url"`
	Title                string `body:"title"`
	Logo                 string `body:"logo"`
	TranslatorId         int    `body:"translatorId"`
	TranslateTitle       bool   `body:"translateTitle"`
	TranslateDescription bool   `body:"translateDescription"`
	Public               bool   `body:"public"`
	Description          string `body:"description"`
	FromLang             string `body:"fromLang"`
	ToLang               string `body:"toLang"`
}
