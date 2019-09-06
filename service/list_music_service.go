package service

import (
	"singo/model"
	"singo/serializer"
)

type ShowMusicService struct {
}

func (service *ShowMusicService) Show(id string) serializer.Response {
	var music model.Music
	err := model.DB.First(&music, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "音乐不存在",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildMusic(music),
	}
}
