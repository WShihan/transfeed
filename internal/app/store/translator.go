package store

import (
	"transfeed/internal/app/form"
	"transfeed/internal/app/model"
)

func GetTranslator(user *model.User, id string) (*model.Translator, error) {
	trans := model.Translator{}
	err := DB.Where("user_id = ?", user.ID).First(&trans, id).Error
	return &trans, err

}

func CreateTranslator(user *model.User, payload form.TranslatorAddForm) (*model.Translator, error) {

	trans := model.Translator{
		Url:    payload.Url,
		Name:   payload.Name,
		Role:   payload.Role,
		Key:    payload.Key,
		Prompt: payload.Prompt,
		Lang:   payload.Lang,
	}
	err := DB.Model(user).Association("Translators").Append(&trans)
	return &trans, err

}

func UpdateTranslator(user *model.User, payload form.TranslatorUpdateForm) (*model.Translator, error) {
	trans := model.Translator{}
	err := DB.Where("user_id = ?", user.ID).First(&trans, payload.ID).Error
	if err != nil {
		return nil, err
	}
	if trans.Url != payload.Url && trans.Url != "" {
		trans.Url = payload.Url
	}
	if trans.Name != payload.Name && trans.Name != "" {
		trans.Name = payload.Name
	}
	if trans.Role != payload.Role && trans.Role != "" {
		trans.Role = payload.Role
	}
	if trans.Key != payload.Key && trans.Key != "" {
		trans.Key = payload.Key
	}
	if trans.Prompt != payload.Prompt && trans.Prompt != "" {
		trans.Prompt = payload.Prompt
	}
	if trans.Lang != payload.Lang && trans.Lang != "" {
		trans.Lang = payload.Lang
	}
	err = DB.Save(&trans).Error
	return &trans, err
}

func GetTranslators(user *model.User) ([]*model.Translator, error) {
	err := DB.Preload("Translators").First(user, user.ID).Error
	return user.Translators, err
}
