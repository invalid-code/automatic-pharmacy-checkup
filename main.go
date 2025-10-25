package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
	msgr_acct := os.Getenv("MESSENGER_ACCT")
	msgr_pass := os.Getenv("MESSENGER_PASS")
	// opts := append(
	// 	chromedp.DefaultExecAllocatorOptions[:],
	// 	chromedp.Flag("no-sandbox", true),
	// 	chromedp.Flag("headless", true),
	// )
	// allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	// defer cancel()

	ctx, cancle := chromedp.NewContext(context.Background())
	defer cancle()

	var htmlContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.messenger.com/"),
		chromedp.Sleep(2*time.Second),
		chromedp.SendKeys("#email", msgr_acct),
		chromedp.SendKeys("#pass", msgr_pass),
		chromedp.Click("#loginbutton"),
		chromedp.Sleep(2*time.Second),
		chromedp.OuterHTML("html", &htmlContent),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(htmlContent)
}
