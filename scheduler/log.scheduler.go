package logScheduler

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func SecondCron() {
	// gocron 스케줄러 설정
	scheduler := gocron.NewScheduler(time.Local)

	// 매 초마다 실행될 작업 추가 (예시)
	scheduler.Every(1).Second().Do(func() {
		log.Println("매 초마다 실행됩니다.")
		// 여기서 필요한 작업을 수행합니다
	})

	// 스케줄러 시작
	scheduler.StartAsync()

}
