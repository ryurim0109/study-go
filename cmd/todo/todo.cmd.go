package todo

import (
	"time"
	"strconv"
	"log"

	"github.com/gofiber/fiber/v2"
	mariadb "github.com/ryurim0109/study-go/db"
)

// Create 함수
func Create(c *fiber.Ctx) error {
	db := mariadb.InitDB();

	log.Println("Logging")

	// 요청 본문에서 text 추출
	type Request struct {
		Text string `json:"text"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// 현재 시간으로 생성/업데이트 시간 설정
	now := time.Now()

	// SQL 쿼리문 준비
	statement, err := db.Prepare("INSERT INTO edel_todo (text, createdAt, updatedAt) VALUES (?, ?, ?)")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer statement.Close()

	// SQL 쿼리문 실행
	res, err := statement.Exec(req.Text, now, now)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// 삽입된 레코드의 ID 반환
	todoId, err := res.LastInsertId()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"id": todoId}})
}


// read 
func Get(c *fiber.Ctx) error{
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var notes []

}