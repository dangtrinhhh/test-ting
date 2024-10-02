package Initializers

import (
	"log"
	"os"
	"time"

	"wan-api-kol-event/Models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	//Get database url from environment variables (defined in .env file)
	var dsn string = os.Getenv("DB_URL")

	//Connect with postgres
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             50 * time.Millisecond, // Slow SQL threshold
				LogLevel:                  logger.Warn,           // Log level
				IgnoreRecordNotFoundError: false,                 // Dont ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      false,                 // Include params in the SQL log
				Colorful:                  true,                  // Disable color
			},
		),
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Automatically create table KOL if not exist
	if !isTableExists("KOL") {
		log.Println("Table KOL not exist, creating...")
		err = DB.AutoMigrate(&Models.Kol{})
		if err != nil {
			log.Fatal("Failed to migrate database:", err)
		}
	} else {
		log.Println("KOL is already exist")
		var kols []Models.Kol
		if err := DB.Find(&kols).Error; err != nil {
			log.Fatalf("Error get data from KOL table: %v", err)
		}
	}
}

// Hàm kiểm tra sự tồn tại của bảng
func isTableExists(tableName string) bool {
	var exists bool
	err := DB.Raw("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = ?)", tableName).Scan(&exists).Error
	if err != nil {
		log.Fatalf("Error when check table exist: %v", err)
	}
	return exists
}
