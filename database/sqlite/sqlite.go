package sqlite

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/muzammil-cyber/golang-gin/database"
)

type SQLiteDB struct {
	DB *gorm.DB
}

func getPath() string {
	homeDir := os.Getenv("DB_PATH")
	if homeDir == "" {
		homeDir = "./data/sqlite.db"
	}
	return homeDir
}

func NewSQLiteDB() (database.Database, error) {
	dbPath := getPath()
	dir := filepath.Dir(dbPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, fmt.Errorf("failed to create directory: %w", err)
		}
	}

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	gormDB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize gorm with sqlite: %w", err)
	}

	return &SQLiteDB{DB: gormDB}, nil
}

func (s *SQLiteDB) GetDB() *gorm.DB {
	return s.DB
}
