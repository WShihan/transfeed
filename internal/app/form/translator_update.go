package form

type TranslatorUpdateForm struct {
	ID     int    `body:"id"`
	Name   string `body:"name"`
	Role   string `body:"role"`
	Url    string `body:"url"`
	Key    string `body:"key"`
	Prompt string `body:"prompt"`
	Lang   string `body:"lang"`
}
