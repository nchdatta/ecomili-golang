package validations

import (
	"time"

	_ "github.com/go-playground/validator/v10"
	"github.com/nchdatta/ecomili-golang/internal/models"
)

type NewsCreate struct {
	Title            string            `json:"title" validate:"required,max=250"`
	Slug             string            `json:"slug" validate:"required"`
	CategoryID       uint              `json:"category_id" validate:"required,gte=1"`
	AuthorID         uint              `json:"author_id" validate:"required,gte=1"`
	Language         string            `json:"language"`
	Type             string            `json:"type"`
	FeaturedImage    string            `json:"featured_image"`
	Thumbnail        string            `json:"thumbnail"`
	Excert           string            `json:"excert"`
	Description      string            `json:"description"`
	ImgSourceURL     string            `json:"img_source_url"`
	OriginalNewsURL  string            `json:"originial_news_url"`
	NewsSource       models.NewsSource `json:"news_source" validate:"oneof=local reuters"`
	PublishTimestamp time.Time         `json:"publish_timestamp"`
	IsFeatured       bool              `json:"is_featured"`
	Tags             []string          `json:"tags"`
	Images           []string          `json:"images"`
}

type NewsUpdate struct {
	Title            string            `json:"title" validate:"required,max=250"`
	Slug             string            `json:"slug" validate:"required"`
	CategoryID       uint              `json:"category_id" validate:"required,gte=1"`
	AuthorID         uint              `json:"author_id" validate:"required,gte=1"`
	Language         string            `json:"language"`
	Type             string            `json:"type"`
	FeaturedImage    string            `json:"featured_image"`
	Thumbnail        string            `json:"thumbnail"`
	Excert           string            `json:"excert"`
	Description      string            `json:"description"`
	ImgSourceURL     string            `json:"img_source_url"`
	OriginalNewsURL  string            `json:"originial_news_url"`
	NewsSource       models.NewsSource `json:"news_source" validate:"oneof=local reuters"`
	PublishTimestamp time.Time         `json:"publish_timestamp"`
	IsFeatured       bool              `json:"is_featured"`
	Tags             []string          `json:"tags"`
	Images           []string          `json:"images"`
}
