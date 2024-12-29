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
    ID          uint      `gorm:"primaryKey"`
    Title       string    `gorm:"not null"`
    Description string
    Body        string    `gorm:"not null"`
    UserID      uint      `gorm:"not null"`
    Likes       uint      `gorm:"default:0"`
    Image       string
    Comments    uint      `gorm:"default:0"`
    Shares      uint      `gorm:"default:0"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    ThreadTags  []ThreadTag `gorm:"foreignKey:ThreadID"`
}

type ThreadTag struct {
    ThreadID     uint        `gorm:"primaryKey;autoIncrement:false"`
    Tag          string      `gorm:"primaryKey;autoIncrement:false"`
}

type User struct {
    ID           uint      `gorm:"primaryKey"`
    Username     string    `gorm:"not null;unique;type:varchar(255)"`
    PasswordHash string    `gorm:"not null;type:varchar(255)"`
    ProfileImage string 
    Threads      []Thread
    Posts        uint
    Comments     uint
    Aura         uint  
    CreatedAt    time.Time
    UpdatedAt    time.Time
}

type UserThreadLikes struct {
    UserID   uint `gorm:"primaryKey;autoIncrement:false"`
    ThreadID uint `gorm:"primaryKey;autoIncrement:false"`
}

type UserThreadComments struct {
    ID                   uint `gorm:"primaryKey"`
    UserID               uint
    ThreadID             uint 
    Comment              string `gorm:"not null"`
    Likes                uint
    Comments             uint
    UserThreadCommentsID uint 
    CreatedAt            time.Time
    UpdatedAt            time.Time
}

type UserThreadCommentLikes struct {
    UserID    uint `gorm:"primaryKey;autoIncrement:false"`
    CommentID uint `gorm:"primaryKey;autoIncrement:false"`
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

    err = db.AutoMigrate(&UserThreadLikes{})
    if err != nil {
        return fmt.Errorf("failed to migrate database: %w", err)
    }

    err = db.AutoMigrate(&UserThreadComments{})
    if err != nil {
        return fmt.Errorf("failed to migrate database: %w", err)
    }

    err = db.AutoMigrate(&UserThreadCommentLikes{})
    if err != nil {
        return fmt.Errorf("failed to migrate database: %w", err)
    }

    err = db.AutoMigrate(&ThreadTag{})
    if err != nil {
        return fmt.Errorf("failed to migrate database: %w", err)
    }

    return nil
}


