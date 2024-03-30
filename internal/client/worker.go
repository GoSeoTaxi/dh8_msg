package client

import (
	"context"
	"log"
	"time"

	"github.com/GoSeoTaxi/dh8_msg/internal/config"
	"github.com/GoSeoTaxi/dh8_msg/internal/sender"
)

func StartWorker(ctx context.Context, cfg *config.Config) {

	res, err := PingUrlWithTimeOutAndCount(cfg.URL, 10, time.Duration(5*time.Second))
	if err != nil {
		sender.SendResp(cfg)
	}
	// Тут нужен какой-то конвертер для отправки результатов
	log.Println(res)

}
