package translator

import (
	"testing"
	"transfeed/internal/app/model"
)

func TestLLMTranslator(t *testing.T) {
	llm := model.Translator{
		Name:   "deepseek",
		Role:   "deepseek-chat",
		Key:    "",
		Prompt: "你是一个智能助手",
		Url:    "https://api.deepseek.com/chat/completions",
	}
	ds := LLMTranslator{
		Agent: llm,
	}
	msg, err := ds.Execute("你好")
	if err != nil {
		t.Logf(err.Error())
	}
	t.Logf(msg)
}
