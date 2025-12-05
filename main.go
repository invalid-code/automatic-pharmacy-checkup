package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
	cUser := os.Getenv("C_USER")
	datr := os.Getenv("DATR")
	locale := os.Getenv("LOCALE")
	psL := os.Getenv("PS_L")
	psN := os.Getenv("PS_N")
	sb := os.Getenv("SB")
	wd := os.Getenv("WD")
	xs := os.Getenv("XS")
	cookies := []string{"c_user", cUser, "datr", datr, "locale", locale, "ps_l", psL, "ps_n", psN, "sb", sb, "wd", wd, "xs", xs}
	host := "https://www.messenger.com/"
	parsedURL, err := url.Parse(host)
	if err != nil {
		panic("invalid host URL")
	}
	domain := parsedURL.Hostname()

	ctx, cancle := chromedp.NewContext(context.Background())
	defer cancle()

	var ccmc string
	// var perpetual string
	err = chromedp.Run(
		ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			for i := 0; i < len(cookies); i += 2 {
				err := network.SetCookie(cookies[i], cookies[i+1]).
					WithExpires(&expr).
					WithDomain(domain).
					WithPath("/").
					WithHTTPOnly(false).
					WithSecure(false).
					Do(ctx)
				if err != nil {
					return err
				}
			}
			return nil
		}),
		// chromedp.Navigate("https://www.messenger.com/t/7808053889217350"),
		// chromedp.WaitVisible("span.x1lliihq.x1plvlek.xryxfnj.x1n2onr6", chromedp.ByQuery),
		// chromedp.OuterHTML("html", &perpetual, chromedp.ByQuery),
		chromedp.Navigate("https://www.messenger.com/t/6843762519014249"),
		chromedp.WaitVisible("span.x1lliihq.x1plvlek.xryxfnj.x1n2onr6", chromedp.ByQuery),
		chromedp.OuterHTML("html", &ccmc, chromedp.ByQuery),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(ccmc)
}
