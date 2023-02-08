package models

import "gorm.io/gorm"

type ShortUrl struct {
	gorm.Model
	Url      string `json:"url" gorm:"not null"`
	ShortUrl string `json:"short_url" gorm:"not null"`
	Uid      uint   `json:"uid" gorm:"not null"`
}

func (su *ShortUrl) CreateShortUrl() (*ShortUrl, error) {
	err := DB.Create(&su).Error

	if err != nil {
		return nil, err
	}

	return su, nil
}

func GetUrl(shortUrl string) (string, error) {
	su := &ShortUrl{}

	err := DB.Model(ShortUrl{}).Where("short_url = ?", shortUrl).Take(su).Error
	if err != nil {
		return "", err
	}

	return su.Url, nil
}

func GetUrlListByUid(uid uint) ([]ShortUrl, error) {
	var listSU []ShortUrl

	err := DB.Model(ShortUrl{}).Where("uid = ?", uid).Find(&listSU).Error
	if err != nil {
		return listSU, err
	}

	return listSU, nil
}
