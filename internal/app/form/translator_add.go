package form

type TranslatorAddForm struct {
	Name   string `query:"name"`
	Role   string `query:"role"`
	Url    string `query:"url"`
	Key    string `query:"key"`
	Prompt string `query:"prompt"`
	Lang   string `query:"lang"`
}
