package boot

import (
	"fmt"
	"github/vamshi1997/projects/go/youtube_app/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitApp() {
	InitConfig()
	InitDb()
}

func GetDB() *gorm.DB {
	return db
}

func InitDb() {
	cfg = GetConfig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%v sslmode=%v TimeZone=Asia/Kolkata",
		cfg.AppConfig.DB.Host,
		cfg.AppConfig.DB.UserName,
		cfg.AppConfig.DB.Password,
		cfg.AppConfig.DB.DBName,
		cfg.AppConfig.DB.Port,
		cfg.AppConfig.DB.SslMode,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Not able connect to database")
		panic(err)
	}
	db.AutoMigrate(&models.Video{})
	fmt.Println("Application connected to database successfully ...")
}
