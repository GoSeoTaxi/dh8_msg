package cmd

import (
	"context"
	"log"

	"github.com/GoSeoTaxi/dh8_msg/internal/client"
	"github.com/GoSeoTaxi/dh8_msg/internal/config"
	"github.com/fatih/color"
)

func RunService(ctx context.Context) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("parcing config err = %v", err)
	}

	log.Println(color.HiRedString("Starting..."))

	client.StartWorker(ctx, cfg)

}
