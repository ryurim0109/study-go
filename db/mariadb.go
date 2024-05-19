package mariadb

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)



func GetConnector() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(".env 파일을 찾을 수 없습니다.")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbDatabase := os.Getenv("DB_DATABASE")
	password := os.Getenv("DB_PASSWORD")

	cfg := mysql.Config{
		User:                 dbUser,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 dbHost+":"+dbPort,
		Collation:            "utf8mb4_general_ci",
		Loc:                  time.UTC,
		MaxAllowedPacket:     4 << 20.,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		DBName:							dbDatabase ,
		ParseTime: true,
	}
	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		panic(err)
	}
	db := sql.OpenDB(connector)
	return db
	
}

func InitDB() *sql.DB {
	db := GetConnector()
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

// func Init() {
	


// 	var dbErr error

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		dbUser, password, dbHost, dbPort, dbName)

// 	DBConn, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if dbErr != nil {
// 		log.Fatal("failed to connect database")
// 	}
// }