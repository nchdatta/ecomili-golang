package services

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/nchdatta/ecomili-golang/internal/app/validations"
	"github.com/nchdatta/ecomili-golang/internal/database"
	"github.com/nchdatta/ecomili-golang/internal/models"
	"gorm.io/gorm"
)

func GetAllNews() (*[]models.News, error) {
	newsList := []models.News{}

	if err := database.DBConn.Preload("Tags").Preload("Images").Find(&newsList).Error; err != nil {
		return nil, err
	}

	return &newsList, nil
}
func GetNewsByID(id int) (*models.News, error) {
	news := &models.News{}

	if errFound := database.DBConn.Where("id = ?", id).First(&news).Error; errFound != nil {
		if errFound == gorm.ErrRecordNotFound {
			return nil, errors.New("news not found")
		}
		return nil, errFound
	}
	return news, nil
}
func CreateNews(newsCreate *validations.NewsCreate) (*models.News, error) {
	var existingNews models.News
	existErr := database.DBConn.Where("slug = ?", strings.ToLower(newsCreate.Slug)).Find(&existingNews).Error

	if existingNews.Slug != "" {
		return nil, errors.New("News already exists with the slug: " + newsCreate.Slug)
	}
	if existErr != nil {
		return nil, existErr
	}

	news := &models.News{
		Title:            newsCreate.Title,
		Slug:             newsCreate.Slug,
		CategoryID:       newsCreate.CategoryID,
		AuthorID:         newsCreate.AuthorID,
		Language:         newsCreate.Language,
		Type:             newsCreate.Type,
		FeaturedImage:    sql.NullString{String: newsCreate.FeaturedImage, Valid: newsCreate.FeaturedImage != ""},
		Thumbnail:        sql.NullString{String: newsCreate.Thumbnail, Valid: newsCreate.FeaturedImage != ""},
		IsFeatured:       newsCreate.IsFeatured,
		Excert:           newsCreate.Excert,
		Description:      newsCreate.Description,
		ImgSourceURL:     newsCreate.ImgSourceURL,
		OriginalNewsURL:  newsCreate.OriginalNewsURL,
		NewsSource:       newsCreate.NewsSource,
		PublishTimestamp: newsCreate.PublishTimestamp,
	}

	if err := database.DBConn.Create(&news).Error; err != nil {
		return nil, err
	}

	if len(newsCreate.Tags) > 0 {
		for _, val := range newsCreate.Tags {
			tag := &models.Tag{
				Name:   val,
				NewsID: news.ID,
			}
			if err := database.DBConn.Create(tag).Error; err != nil {
				return nil, err
			}
		}
	}

	if len(newsCreate.Images) > 0 {
		for _, val := range newsCreate.Images {
			newsImage := &models.NewsImage{
				Image:  sql.NullString{String: val, Valid: val != ""},
				NewsID: news.ID,
			}
			if err := database.DBConn.Create(newsImage).Error; err != nil {
				return nil, err
			}
		}
	}

	return news, nil
}

func UpdateNews(id int, newsUpdate *validations.NewsUpdate) (*models.News, error) {
	news := &models.News{}
	if err := database.DBConn.Where("id = ?", id).First(&news).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("news not found")
		}
		return nil, err
	}

	news.Title = newsUpdate.Title
	news.Slug = newsUpdate.Slug
	news.CategoryID = newsUpdate.CategoryID
	news.AuthorID = newsUpdate.AuthorID
	news.Language = newsUpdate.Language
	news.Type = newsUpdate.Type
	news.FeaturedImage = sql.NullString{String: newsUpdate.FeaturedImage, Valid: newsUpdate.FeaturedImage != ""}
	news.Thumbnail = sql.NullString{String: newsUpdate.Thumbnail, Valid: newsUpdate.FeaturedImage != ""}
	news.IsFeatured = newsUpdate.IsFeatured
	news.Excert = newsUpdate.Excert
	news.Description = newsUpdate.Description
	news.ImgSourceURL = newsUpdate.ImgSourceURL
	news.OriginalNewsURL = newsUpdate.OriginalNewsURL
	news.NewsSource = newsUpdate.NewsSource
	news.PublishTimestamp = newsUpdate.PublishTimestamp

	for i := range newsUpdate.Tags {
		tagName := newsUpdate.Tags[i]
		var existingTag models.Tag
		existErr := database.DBConn.First(&existingTag, "name = ? AND news_id = ?", tagName, news.ID).Error

		if existErr == nil {
			existingTag.Name = tagName
			if err := database.DBConn.Save(&existingTag).Error; err != nil {
				return nil, err
			}
		}

		// Create new tag if it doesn't exist
		if errors.Is(existErr, gorm.ErrRecordNotFound) {
			tag := &models.Tag{
				Name:   tagName,
				NewsID: news.ID,
			}
			if err := database.DBConn.Create(tag).Error; err != nil {
				return nil, err
			}
		}
	}

	if len(newsUpdate.Images) > 0 {
		for _, val := range newsUpdate.Images {
			newsImage := &models.NewsImage{
				Image:  sql.NullString{String: val, Valid: val != ""},
				NewsID: news.ID,
			}
			if err := database.DBConn.Create(newsImage).Error; err != nil {
				return nil, err
			}
		}
	}

	if err := database.DBConn.Save(&news).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func DeleteNews(id int) (*models.News, error) {
	news := &models.News{}

	if err := database.DBConn.Where("id = ?", id).First(&news).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("news not found")
		}
		return nil, err
	}

	if err := database.DBConn.Where("news_id = ?", id).Delete(&models.Tag{}).Error; err != nil {
		return nil, err
	}
	if err := database.DBConn.Where("news_id = ?", id).Delete(&models.NewsImage{}).Error; err != nil {
		return nil, err
	}

	if err := database.DBConn.Delete(&news).Where("id=?", id).Error; err != nil {
		return nil, err
	}
	return nil, nil
}
