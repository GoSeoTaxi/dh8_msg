package main

import (
	"context"
	"fmt"

	"github.com/GoSeoTaxi/dh8_msg/cmd"
	"github.com/fatih/color"
)

func main() {
	fmt.Println(color.HiMagentaString("Preparing..."))
	ctx, cancel := context.WithCancel(context.Background())

	cmd.RunService(ctx)

	fmt.Println(color.HiRedString("Shutting down..."))
	cancel()
}
