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

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // Set headless flag to false
	)

	allocatorCtx, cancelAllocator := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelAllocator()

	taskCtx, cancelTask := chromedp.NewContext(allocatorCtx)
	defer cancelTask()

	//	var ccmc string
	// var perpetual string
	var content string
	err = chromedp.Run(
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
		// chromedp.OuterHTML("html", &perpetual, chromedp.ByQuery),
		//		chromedp.Navigate("https://www.messenger.com/t/6843762519014249"),
		//		chromedp.WaitVisible("span.x1lliihq.x1plvlek.xryxfnj.x1n2onr6", chromedp.ByQuery),
		//		chromedp.OuterHTML("html", &ccmc, chromedp.ByQuery),
		chromedp.Navigate("https://www.messenger.com/e2ee/t/9514674638608261"),
		chromedp.WaitVisible("span.x1lliihq.x1plvlek.xryxfnj.x1n2onr6", chromedp.ByQuery),
		chromedp.Click("div.x9f619.x1ja2u2z.xzpqnlu.x1hyvwdk.x14bfe9o.xjm9jq1.x6ikm8r.x10wlt62.x10l6tqk.x1i1rx1s+button"),
		chromedp.Click("div.x1i10hfl.x1ejq31n.x18oe1m7.x1sy0etr.xstzfhl.x9f619.x1ypdohk.x3ct3a4.xdj266r.x14z9mp.xat24cr.x1lziwak.x16tdsg8.x1hl2dhg.xggy1nq.x87ps6o.x1lku1pv.x1a2a7pz.x6s0dn4.x1rl75mt.x19t5iym.xz7t8uv.x13xmedi.x972fbf.x10w94by.x1qhh985.x14e42zd.x78zum5.xl56j7k.xexx8yu.xyri2b.x18d9i69.x1c1uobl.x1n2onr6.x10w6t97.x1td3qas.x10ltxyv"),
		chromedp.Click("div.x1i10hfl.xjbqb8w.x1ejq31n.x18oe1m7.x1sy0etr.xstzfhl.x972fbf.x10w94by.x1qhh985.x14e42zd.x1ypdohk.x3ct3a4.xdj266r.x14z9mp.xat24cr.x1lziwak.xexx8yu.xyri2b.x18d9i69.x1c1uobl.x16tdsg8.x1hl2dhg.xggy1nq.x1fmog5m.xu25z0z.x140muxe.xo1y3bh.x87ps6o.x1lku1pv.x1a2a7pz.x9f619.x3nfvp2.xdt5ytf.xl56j7k.x1n2onr6.xh8yej3[aria-label=\"Don't restore messages\"]"),
		chromedp.Evaluate(`
			const inputContainer = document.querySelector('p.xat24cr.xdj266r');
			let inputText = document.createElement('span');
			inputText.classList.add('x3jgonx');
			inputText.dataset.lexicalText = 'true';
			inputText.innerText = 'hi';
			inputContainer.innerHTML = '';
			inputContainer.appendChild(inputText);
		`, nil),
		chromedp.Click("div.x1i10hfl.x1qjc9v5.xjbqb8w.xjqpnuy.xc5r6h4.xqeqjp1.x1phubyo.x13fuv20.x18b5jzi.x1q0q8m5.x1t7ytsu.x972fbf.x10w94by.x1qhh985.x14e42zd.x9f619.x1ypdohk.xdl72j9.x2lah0s.x3ct3a4.xdj266r.xat24cr.x2lwn1j.xeuugli.x1n2onr6.x16tdsg8.x1hl2dhg.xggy1nq.x1ja2u2z.x1t137rt.x1fmog5m.xu25z0z.x140muxe.xo1y3bh.x3nfvp2.x1q0g3np.x87ps6o.x1lku1pv.x1a2a7pz.x13fj5qh.x1xegmmw.x1y1aw1k.xwib8y2.x1pixwil.x1bjonze[role=\"button\"]"),
//		chromedp.OuterHTML("html", &content, chromedp.ByQuery),
	)
	if err != nil {
		panic(err)
	}
	//	fmt.Println(ccmc)
	fmt.Println(content)
}
