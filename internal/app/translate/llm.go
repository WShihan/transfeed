package translator

import (
	"fmt"
	"sync"
	"transfeed/internal/app/model"
	"transfeed/internal/app/store"

	"github.com/go-resty/resty/v2"
)

type Translator interface {
	Execute(text string) (string, error)
	Translate(sem chan struct{}, wg *sync.WaitGroup, feed *model.Feed, entry *model.Entry)
}

// 大语言翻译结果结构
type LLMResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     any    `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens        int `json:"prompt_tokens"`
		CompletionTokens    int `json:"completion_tokens"`
		TotalTokens         int `json:"total_tokens"`
		PromptTokensDetails struct {
			CachedTokens int `json:"cached_tokens"`
		} `json:"prompt_tokens_details"`
		PromptCacheHitTokens  int `json:"prompt_cache_hit_tokens"`
		PromptCacheMissTokens int `json:"prompt_cache_miss_tokens"`
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
}

type LLMTranslator struct {
	Agent model.Translator
}

func (translator *LLMTranslator) Execute(text string) (string, error) {
	client := resty.New()
	postData := map[string]interface{}{
		"model": translator.Agent.Role,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": translator.Agent.Prompt,
			},
			{
				"role":    "system",
				"content": "好的，请发送文字",
			},
			{
				"role":    "user",
				"content": text,
			},
		},
	}
	llmData := LLMResponse{}
	resp, err := client.R().
		SetResult(&llmData).
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", translator.Agent.Key)).
		SetBody(postData).
		Post(translator.Agent.Url)

	if err != nil {
		return "", err
	}
	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("request failed")
	}
	if len(llmData.Choices) > 0 {
		translator.Agent.Comsume += int64(llmData.Usage.TotalTokens)
		store.DB.Save(&translator.Agent)
		return llmData.Choices[0].Message.Content, nil
	} else {
		return "", fmt.Errorf("no response")
	}
}

func (translator *LLMTranslator) Translate(sem chan struct{}, wg *sync.WaitGroup, feed *model.Feed, entry *model.Entry) {
	defer func() {
		wg.Done()
		sem <- struct{}{} // 获取信号量
		<-sem             // 释放信号量
	}()
	fmt.Printf("----Translate %s---\n", entry.Title)
	if feed.TranslateTitle {
		title := entry.Title
		newTitle, err := translator.Execute(title)
		if err != nil {
			return
		}
		entry.Title += fmt.Sprintf("｜%s", newTitle)
	}

	if feed.TranslateDescription {
		desc := entry.Summary
		newDesc, err := translator.Execute(desc)
		if err != nil {
			return
		}
		entry.Summary = fmt.Sprintf("%s\n%s", newDesc, entry.Summary)
	}
}
