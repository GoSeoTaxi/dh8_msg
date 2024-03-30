package client

import (
	"context"
	"log"
	"time"

	"github.com/GoSeoTaxi/dh8_msg/internal/config"
)

func StartWorker(ctx context.Context, cfg *config.Config) (err error) {

	if cfg.CheckPing {
		err = pingUrlWithTimeOutAndCount(cfg.URL, 10, time.Duration(5*time.Second))
		if err != nil {
			log.Println(err)
		}
	}
	if cfg.CheckFile {
		err = checkFile(cfg.URLFile, cfg.Sha256)
	}

	if cfg.CheckLogin {
		err = checkLogin(cfg.URLLogin)
	}

	return err
}
