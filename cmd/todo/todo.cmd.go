package todo

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	mariadb "github.com/ryurim0109/study-go/db"
	errorHandler "github.com/ryurim0109/study-go/error"
	models "github.com/ryurim0109/study-go/model"
)

// Create 함수
func Create(c *fiber.Ctx) error {
	conn := mariadb.InitDB()

	// 요청 본문에서 text 추출
	type Request struct {
		Text string `json:"text"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return errorHandler.SendJSONError(c, "cannot parse JSON")
	}

	// 현재 시간으로 생성/업데이트 시간 설정
	now := time.Now()

	// SQL 쿼리문 준비
	statement, err := conn.Prepare("INSERT INTO edel_todo (text, createdAt, updatedAt) VALUES (?, ?, ?)")
	if err != nil {
		return errorHandler.SendJSONError(c, err.Error())
	}
	defer statement.Close()

	// SQL 쿼리문 실행
	res, err := statement.Exec(req.Text, now, now)
	if err != nil {
		return errorHandler.SendJSONError(c, err.Error())
	}

	// 삽입된 레코드의 ID 반환
	todoId, err := res.LastInsertId()
	if err != nil {
		return errorHandler.SendJSONError(c, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"id": todoId}})
}

// Read 함수
func Get(c *fiber.Ctx) error {
	conn := mariadb.InitDB()

	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	rows, err := conn.Query("SELECT * FROM edel_todo WHERE isDel = 0 LIMIT ? OFFSET ?", intLimit, offset)
	if err != nil {
		return errorHandler.SendJSONError(c, err.Error())
	}
	defer rows.Close() // 함수 종료 시 닫아서 추가 열거를 방지

	var todoList []models.ToDo

	for rows.Next() {
		var todo models.ToDo
		err := rows.Scan(
			&todo.Id,
			&todo.Text,
			&todo.IsDel,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return errorHandler.SendJSONError(c, err.Error())
		}
		todoList = append(todoList, todo)
	}

	if err = rows.Err(); err != nil {
		return errorHandler.SendJSONError(c, err.Error())
	}

	// 응답 반환
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"todoList": todoList}})

}

// Update 함수
func Update(c *fiber.Ctx) error {
	conn := mariadb.InitDB()
	defer conn.Close()
	tx, err := conn.Begin()
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	todoId := c.Params("todoId")
	type Request struct {
		Text string `json:"text"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return errorHandler.SendJSONError(c, "cannot parse JSON")
	}

	text := req.Text
	// 현재 시간으로 업데이트 시간 설정
	now := time.Now()
	custom := now.Format("2006-01-02 15:04:05")

	query := fmt.Sprintf("UPDATE edel_todo SET text = '%s', updatedAt = '%s' WHERE id= '%s'",
		text, custom, todoId)
	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		panic(err)
	} else {
		tx.Commit()
	}
	// 응답 반환
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"status": "success"})
}

// Delete 함수
func Delete(c *fiber.Ctx) error {
	conn := mariadb.InitDB()

	defer conn.Close()
	tx, err := conn.Begin()
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	todoId := c.Params("todoId")

	query := fmt.Sprintf("UPDATE edel_todo SET isDel = 1 WHERE id = '%s'",
		todoId)
	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		panic(err)
	} else {
		tx.Commit()
	}
	// 응답 반환
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"status": "success"})
}
