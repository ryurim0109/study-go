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
		Addr:                 dbHost + ":" + dbPort,
		Collation:            "utf8mb4_general_ci",
		Loc:                  time.UTC,
		MaxAllowedPacket:     4 << 20.,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		DBName:               dbDatabase,
		ParseTime:            true,
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
	if err != nil || db.Ping() != nil { // 핑을 통해 네트워크 연결 및 DB 사용 가능한지 체크
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3) // 유휴 연결 시간 설정 (5분 미만으로 잡는 게 좋다고 함)
	db.SetMaxOpenConns(10)                 // 최대 연결 수 설정
	db.SetMaxIdleConns(10)                 // 최대 유휴 연결 수 설정
	return db
}
