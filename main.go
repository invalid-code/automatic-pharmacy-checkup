package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	// "sync"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"github.com/joho/godotenv"
	// "github.com/robfig/cron/v3"
)

func sendCheckupMessage(domain string, opts []chromedp.ExecAllocatorOption, cookies []string) {
	allocatorCtx, cancelAllocator := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelAllocator()

	// chromedpCtx, cancelChromedpCtx := chromedp.NewContext(allocatorCtx)
	// defer cancelChromedpCtx()

	// taskCtx, cancelTask := context.WithTimeout(chromedpCtx, time.Hour/2)
	taskCtx, cancelTask := chromedp.NewContext(allocatorCtx)
	defer cancelTask()

	err := chromedp.Run(
		taskCtx,
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
		// chromedp.Click("div[aria-label=\"Close\"][role=\"button\"]"),
		// chromedp.Click("div[aria-label=\"Don't restore messages\"][role=\"button\"][tabindex=\"0\"]"),
		// chromedp.SendKeys("div[aria-label=\"Message\"]", msg+kb.Enter),

		// chromedp.Navigate("https://www.messenger.com/t/6843762519014249"),
		chromedp.Navigate("https://www.messenger.com/t/888200397021574/"),
		chromedp.WaitVisible("span.x1lliihq.x1plvlek.xryxfnj.x1n2onr6", chromedp.ByQuery),
		chromedp.Click("div[aria-label=\"Close\"][role=\"button\"]"),
		chromedp.Click("div[aria-label=\"Don't restore messages\"][role=\"button\"][tabindex=\"0\"]"),
		chromedp.SendKeys("div[aria-label=\"Message\"]", "Good afternoon pila na sales?"+kb.Enter),
		chromedp.WaitVisible("span.xdj266r.x14z9mp.xat24cr.x1lziwak.xexx8yu.xyri2b.x18d9i69.x1c1uobl.x1hl2dhg.x16tdsg8.x1vvkbs.x1xf6ywa"),
		chromedp.Poll(`
			(function() {
				const msgCond = document.querySelector("span.xdj266r.x14z9mp.xat24cr.x1lziwak.xexx8yu.xyri2b.x18d9i69.x1c1uobl.x1hl2dhg.x16tdsg8.x1vvkbs.x1xf6ywa");
				return msgCond.innerText === "Sent";
			})()`, nil, chromedp.WithPollingInterval(1*time.Second)),
		chromedp.Sleep(3*time.Second),
		chromedp.WaitVisible("div.html-div.xdj266r.xat24cr.xexx8yu.xyri2b.x18d9i9.x1c1uobl.x6s0dn4.xmg6eyc.xa4qsjk.xwnhzmj.x4hg4is.x1iuwi03.xr9e8f9.x1e4oeot.x1ui4y5.x6en5u8.78zum5.xqu0tyb.xm2jcoa.x1mpyi22.x51ohtg.x1xwhvez", chromedp.ByQuery),
	)
	if err != nil {
		panic(err)
	}

	fmt.Scanln()

	// var wg sync.WaitGroup
	// wg.Add(2)

	// go func() {
	// 	defer wg.Done()

	// 	ccmcCtx, cancelCcmcCtx := chromedp.NewContext(taskCtx)
	// 	defer cancelCcmcCtx()

	// 	err := chromedp.Run(
	// 		ccmcCtx,
	// 		chromedp.ActionFunc(func(ctx context.Context) error {
	// 			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
	// 			for i := 0; i < len(cookies); i += 2 {
	// 				err := network.SetCookie(cookies[i], cookies[i+1]).
	// 					WithExpires(&expr).
	// 					WithDomain(domain).
	// 					WithPath("/").
	// 					WithHTTPOnly(false).
	// 					WithSecure(false).
	// 					Do(ctx)
	// 				if err != nil {
	// 					return err
	// 				}
	// 			}
	// 			return nil
	// 		}),
	// 		chromedp.Navigate("https://www.messenger.com/t/6843762519014249"),
	// 		chromedp.WaitVisible("span.x1lliihq.x1plvlek.xryxfnj.x1n2onr6", chromedp.ByQuery),
	// 		chromedp.Click("div[aria-label=\"Close\"][role=\"button\"]"),
	// 		chromedp.Click("div[aria-label=\"Don't restore messages\"][role=\"button\"][tabindex=\"0\"]"),
	// 		chromedp.SendKeys("div[aria-label=\"Message\"]", msg+kb.Enter),
	// 	)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }()

	// go func() {
	// 	defer wg.Done()

	// 	perpetualCtx, cancelPerpetualCtx := chromedp.NewContext(taskCtx)
	// 	defer cancelPerpetualCtx()

	// 	err := chromedp.Run(
	// 		perpetualCtx,
	// 		chromedp.ActionFunc(func(ctx context.Context) error {
	// 			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
	// 			for i := 0; i < len(cookies); i += 2 {
	// 				err := network.SetCookie(cookies[i], cookies[i+1]).
	// 					WithExpires(&expr).
	// 					WithDomain(domain).
	// 					WithPath("/").
	// 					WithHTTPOnly(false).
	// 					WithSecure(false).
	// 					Do(ctx)
	// 				if err != nil {
	// 					return err
	// 				}
	// 			}
	// 			return nil
	// 		}),
	// 		chromedp.Navigate("https://www.messenger.com/t/7808053889217350"),
	// 		chromedp.WaitVisible("span.x1lliihq.x1plvlek.xryxfnj.x1n2onr6", chromedp.ByQuery),
	// 		chromedp.Click("div[aria-label=\"Close\"][role=\"button\"]"),
	// 		chromedp.Click("div[aria-label=\"Don't restore messages\"][role=\"button\"][tabindex=\"0\"]"),
	// 		chromedp.SendKeys("div[aria-label=\"Message\"]", msg+kb.Enter),
	// 	)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// }()

	// wg.Wait()
}

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

	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // Set headless flag to false
	)

	// c := cron.New()

	// _, err = c.AddFunc("0 7 * * *", func() {
	// 	sendCheckupMessage(domain, opts, cookies, "Good morning pila na sales?")
	// })
	// if err != nil {
	// 	panic(fmt.Sprintf("Error scheduling task: %v", err))
	// }

	// _, err = c.AddFunc("0 15 * * *", func() {
	// 	sendCheckupMessage(domain, opts, cookies, "Good afternoon pila na sales?")
	// })
	// if err != nil {
	// 	panic(fmt.Sprintf("Error scheduling task: %v", err))
	// }
	sendCheckupMessage(domain, opts, cookies)

	// c.Start()
}
