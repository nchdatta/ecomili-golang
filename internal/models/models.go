package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID     `json:"id" gorm:"type:varchar(36)"`
	Name        string        `json:"name" gorm:"uniqueIndex;type:varchar(200);not null"`
	Phone       *string       `json:"phone" gorm:"uniqueIndex;type:varchar(12)"`
	Email       string        `json:"email" gorm:"uniqueIndex;type:varchar(200);not null"`
	Password    string        `json:"password" gorm:"type:varchar(250);not null"`
	Avatar      *string       `json:"avatar" gorm:"default:null"`
	OTP         *int          `json:"otp" gorm:"default:null"`
	OTPVerified bool          `json:"otp_verified" gorm:"default:false"`
	Role        *Role         `json:"role" gorm:"foreignKey:RoleID;default:null"`
	RoleID      uuid.NullUUID `json:"role_id" gorm:"type:varchar(36);default:null"`
	Infobite    []Infobite    `json:"infobites"`
	Category    []Category    `json:"categories"`
	AddedNews   []News        `json:"added_news" gorm:"foreignKey:AddedByID;"`
	UpdatedNews []News        `json:"updated_news" gorm:"foreignKey:UpdatedByID;"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type Role struct {
	gorm.Model
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36)"`
	Name      string    `json:"name" gorm:"uniqueIndex;type:varchar(100);not null"`
	Status    Status    `json:"status" gorm:"type:enum('active', 'inactive');default:'active'"`
	Users     []User    `json:"users"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Infobite struct {
	gorm.Model
	ID          uuid.UUID     `json:"id" gorm:"type:varchar(36)"`
	Title       string        `json:"title" gorm:"uniqueIndex;type:varchar(150);not null"`
	Picture     *string       `json:"picture" gorm:"type:longtext;default:null"`
	UpdatedBy   *User         `json:"updated_by" gorm:"type:varchar(36);foreignKey:UpdatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UpdatedByID uuid.NullUUID `json:"updated_by_id" gorm:"default:null"`
	Status      Status        `gorm:"default:active"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type Category struct {
	gorm.Model
	ID           uuid.UUID     `json:"id" gorm:"type:varchar(36)"`
	Name         string        `json:"name" gorm:"uniqueIndex;type:varchar(200);not null"`
	Icon         *string       `json:"icon" gorm:"type:longtext;default:null"`
	Tags         []CategoryTag `json:"tags"`
	IsSpecial    bool          `json:"is_special" gorm:"default:false"`
	ParentID     uuid.NullUUID `json:"parent_id" gorm:"type:varchar(36);default:null"`
	Status       Status        `json:"status" gorm:"default:active"`
	CatAddedBy   *User         `json:"cat_added_by" gorm:"foreignKey:CatAddedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CatAddedByID uuid.NullUUID `json:"cat_added_by_id" gorm:"type:varchar(36);default:null"`
	News         []News        `json:"news"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

type CategoryTag struct {
	gorm.Model
	ID         uuid.UUID     `json:"id" gorm:"type:varchar(36)"`
	Name       string        `gorm:"uniqueIndex;type:varchar(150);not null"`
	Category   *Category     `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CategoryID uuid.NullUUID `json:"cat_id" gorm:"type:varchar(36);default:null"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
}

type News struct {
	gorm.Model
	ID               uuid.UUID     `json:"id" gorm:"type:varchar(36)"`
	Title            string        `json:"title" gorm:"type:varchar(250);not null"`
	Slug             string        `json:"slug" gorm:"uniqueIndex;type:varchar(250);not null"`
	Category         *Category     `json:"category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CategoryID       uuid.NullUUID `json:"category_id" gorm:"type:varchar(36);default:null"`
	Author           *Author       `json:"author" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	AuthorID         uuid.NullUUID `json:"author_id" gorm:"type:varchar(36);default:null"`
	FeaturedImage    *string       `json:"featured_image"`
	ThumbnailURL     *string       `json:"thumbnail_url" gorm:"type:longtext"`
	ShortDesc        *string       `json:"short_desc" gorm:"type:longtext"`
	Description      *string       `json:"description" gorm:"type:longtext"`
	ImgSourceURL     *string       `json:"img_source_url" gorm:"type:longtext"`
	OriginalNewsURL  *string       `json:"original_news_url"`
	NewsSource       NewsSource    `json:"news_source" gorm:"default:reuters"`
	PublishTimestamp time.Time     `json:"publish_timestamp"`
	ViewCount        int           `json:"view_count" gorm:"type:int;default:0"`
	ShareCount       int           `json:"share_count" gorm:"type:int;default:0"`
	IsFeatured       bool          `json:"is_featured" gorm:"default:false"`
	URI              *string       `json:"uri" gorm:"uniqueIndex;type:varchar(250)"`
	VersionCreated   time.Time     `json:"version_created"`
	Language         *string       `json:"language" gorm:"type:varchar(20)"`
	Type             *string       `json:"type" gorm:"type:varchar(20)"`
	AddedBy          *User         `json:"added_by" gorm:"foreignKey:AddedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UpdatedBy        *User         `json:"updated_by" gorm:"type:varchar(36);foreignKey:UpdatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	AddedByID        uuid.NullUUID `json:"added_by_id" gorm:"type:varchar(36);"`
	UpdatedByID      uuid.NullUUID `json:"updated_by_id" gorm:"type:varchar(36);"`
	Tags             []Tag         `json:"tags"`
	Images           []NewsImage   `json:"images"`
	Status           NewsStatus    `json:"status" gorm:"default:draft"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
}

type NewsImage struct {
	gorm.Model
	ID     uuid.UUID `json:"id" gorm:"type:varchar(36)"`
	Image  *string   `json:"image" gorm:"type:longtext"`
	News   News      `json:"news" gorm:"foreignKey:NewsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	NewsID uuid.UUID `json:"news_id" gorm:"type:varchar(36)"`
}

type Tag struct {
	gorm.Model
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36)"`
	Name      string    `json:"name" gorm:"type:varchar(200);not null"`
	News      News      `json:"news" gorm:"foreignKey:NewsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	NewsID    uuid.UUID `json:"news_id" gorm:"type:varchar(36);not null"`
	Status    Status    `json:"status" gorm:"default:active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Author struct {
	gorm.Model
	ID          uuid.UUID `json:"id" gorm:"type:varchar(36)"`
	Name        string    `json:"name" gorm:"type:varchar(200);not null"`
	Designation *string   `json:"designation" gorm:"type:varchar(150);default:null"`
	Bio         *string   `json:"bio" gorm:"type:longtext;default:null"`
	Avatar      *string   `json:"avatar" gorm:"type:longtext;default:null"`
	Twitter     *string   `json:"twitter" gorm:"type:varchar(200);default:null"`
	Linkedin    *string   `json:"linkedin" gorm:"type:varchar(200);default:null"`
	Email       *string   `json:"email" gorm:"type:varchar(200);default:null"`
	News        []News    `json:"news"`
	Status      Status    `json:"status" gorm:"default:active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
