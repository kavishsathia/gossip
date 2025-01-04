package models

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Thread struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Body        string `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
	Likes       uint   `gorm:"default:0"`
	Image       string
	Comments    uint `gorm:"default:0"`
	Shares      uint `gorm:"default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	ThreadTags []ThreadTag `gorm:"foreignKey:ThreadID"`
	User       User        `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ThreadTag struct {
	ThreadID uint   `gorm:"primaryKey;autoIncrement:false"`
	Tag      string `gorm:"primaryKey;autoIncrement:false"`

	Thread Thread `gorm:"foreignKey:ThreadID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"not null;unique;type:varchar(255)"`
	PasswordHash string `gorm:"not null;type:varchar(255)"`
	ProfileImage string
	Posts        uint
	Comments     uint
	Aura         uint
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Threads []Thread
}

type ThreadLike struct {
	UserID   uint `gorm:"primaryKey;autoIncrement:false"`
	ThreadID uint `gorm:"primaryKey;autoIncrement:false"`

	User   User   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Thread Thread `gorm:"foreignKey:ThreadID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ThreadComment struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	ThreadID  uint   `gorm:"not null"`
	Comment   string `gorm:"not null"`
	Likes     uint
	Comments  uint
	ParentID  *uint
	CreatedAt time.Time
	UpdatedAt time.Time

	User   User           `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Thread Thread         `gorm:"foreignKey:ThreadID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Parent *ThreadComment `gorm:"foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ThreadCommentLike struct {
	UserID    uint `gorm:"primaryKey;autoIncrement:false"`
	CommentID uint `gorm:"primaryKey;autoIncrement:false"`

	User    User          `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comment ThreadComment `gorm:"foreignKey:CommentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func Migrate() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL not found in environment variables")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(&Thread{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	err = db.AutoMigrate(&ThreadLike{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	err = db.AutoMigrate(&ThreadComment{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	err = db.AutoMigrate(&ThreadCommentLike{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	err = db.AutoMigrate(&ThreadTag{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
