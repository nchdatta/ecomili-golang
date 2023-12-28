package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"type:varchar(200);not null"`
	Phone       string         `json:"phone" gorm:"type:varchar(12)"`
	Email       string         `json:"email" gorm:"uniqueIndex;type:varchar(200);not null"`
	Password    string         `json:"password" gorm:"type:varchar(250);not null"`
	Avatar      sql.NullString `json:"avatar" gorm:"default:null"`
	OTP         sql.NullInt32  `json:"otp" gorm:"default:null"`
	OTPVerified bool           `json:"otp_verified" gorm:"default:false"`
	RoleID      uint           `json:"role_id" gorm:"index;default:null"`
	Role        Role
	Infobite    []Infobite `json:"infobites" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Category    []Category `json:"categories" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Role struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name   string `json:"name" gorm:"uniqueIndex;type:varchar(100);not null"`
	Status Status `json:"status" gorm:"type:enum('active', 'inactive');default:'active'"`
	Users  []User `json:"users" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Infobite struct {
	gorm.Model
	ID      uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Title   string         `json:"title" gorm:"uniqueIndex;type:varchar(150);not null"`
	Picture sql.NullString `json:"picture" gorm:"type:longtext;default:null"`
	UserID  uint           `json:"user_id" gorm:"index"`
	User    User
	Status  Status `gorm:"default:active"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Category struct {
	gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"uniqueIndex;type:varchar(200);not null"`
	Icon      sql.NullString `json:"icon" gorm:"type:longtext;default:null"`
	Tags      []CategoryTag  `json:"tags" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsSpecial bool           `json:"is_special" gorm:"default:false"`
	ParentID  uint           `json:"parent_id" gorm:"default:null"`
	Status    Status         `json:"status" gorm:"default:active"`
	News      []News         `json:"news" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uint           `json:"user_id" gorm:"index"`
	User      User

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CategoryTag struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"uniqueIndex;type:varchar(150);not null"`
	CategoryID uint   `json:"category_id" gorm:"index;default:null"`
	Category   Category

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type News struct {
	gorm.Model
	ID               uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Title            string         `json:"title" gorm:"type:varchar(250);not null"`
	Slug             string         `json:"slug" gorm:"uniqueIndex;type:varchar(250);not null"`
	Category         Category       `json:"category"`
	CategoryID       uint           `json:"category_id" gorm:"index;default:null"`
	AuthorID         uint           `json:"author_id" gorm:"default:null"`
	FeaturedImage    sql.NullString `json:"featured_image"`
	ThumbnailURL     sql.NullString `json:"thumbnail_url" gorm:"type:longtext"`
	ShortDesc        sql.NullString `json:"short_desc" gorm:"type:longtext"`
	Description      sql.NullString `json:"description" gorm:"type:longtext"`
	ImgSourceURL     sql.NullString `json:"img_source_url" gorm:"type:longtext"`
	OriginalNewsURL  sql.NullString `json:"original_news_url"`
	NewsSource       NewsSource     `json:"news_source" gorm:"default:reuters"`
	PublishTimestamp time.Time      `json:"publish_timestamp"`
	ViewCount        int            `json:"view_count" gorm:"type:int;default:0"`
	ShareCount       int            `json:"share_count" gorm:"type:int;default:0"`
	IsFeatured       bool           `json:"is_featured" gorm:"default:false"`
	URI              sql.NullString `json:"uri" gorm:"uniqueIndex;type:varchar(250)"`
	VersionCreated   time.Time      `json:"version_created"`
	Language         *string        `json:"language" gorm:"type:varchar(20)"`
	Type             *string        `json:"type" gorm:"type:varchar(20)"`
	Tags             []Tag          `json:"tags" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Images           []NewsImage    `json:"images" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status           NewsStatus     `json:"status" gorm:"default:draft"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type NewsImage struct {
	gorm.Model
	ID     uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Image  sql.NullString `json:"image" gorm:"type:longtext"`
	NewsID uint           `json:"news_id" gorm:"index"`
	News   News

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Tag struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name   string `json:"name" gorm:"type:varchar(200);not null"`
	NewsID uint   `json:"news_id" gorm:"index;not null"`
	News   News
	Status Status `json:"status" gorm:"default:active"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Author struct {
	gorm.Model
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"type:varchar(200);not null"`
	Designation sql.NullString `json:"designation" gorm:"type:varchar(150);default:null"`
	Bio         sql.NullString `json:"bio" gorm:"type:longtext;default:null"`
	Avatar      sql.NullString `json:"avatar" gorm:"type:longtext;default:null"`
	Twitter     sql.NullString `json:"twitter" gorm:"type:varchar(200);default:null"`
	Linkedin    sql.NullString `json:"linkedin" gorm:"type:varchar(200);default:null"`
	Email       sql.NullString `json:"email" gorm:"type:varchar(200);default:null"`
	News        []News         `json:"news" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status      Status         `json:"status" gorm:"default:active"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
