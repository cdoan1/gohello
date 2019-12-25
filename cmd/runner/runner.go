package runner

import (
	"context"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// Run is a public function to trigger the Chrome runner
func Run(name string, url string, username string, password string) {

	// create options
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("ignore-certificate-errors", "1"),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// capture screenshot of an element
	var buf []byte

	// create output directory if not exists
	newpath := filepath.Join("./output/", name)
	if _, err := os.Stat(newpath); os.IsNotExist(err) {
		os.MkdirAll(newpath, 0755)
	}

	if err := chromedp.Run(ctx, fullScreenshot(url, 100, &buf)); err != nil {
		log.Fatal(err)
	}
	outpath := filepath.Join(newpath, "fullScreenshot.png")
	if err := ioutil.WriteFile(outpath, buf, 0644); err != nil {
		log.Fatal(err)
	}

	log.Printf("login screen - arguments")
	err := chromedp.Run(ctx, login1(url, username, password, 100, &buf))
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(newpath, "login-filled.png"), buf, 0644); err != nil {
		log.Fatal(err)
	}

	log.Printf("login screen - welcome")
	err = chromedp.Run(ctx, login2(url, username, password, 100, &buf))
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(newpath, "welcome.png"), buf, 0644); err != nil {
		log.Fatal(err)
	}

	log.Printf("overview view -")
	err = chromedp.Run(ctx, overview(url, 100, &buf))
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(newpath, "overview.png"), buf, 0644); err != nil {
		log.Fatal(err)
	}

	log.Printf("cluster view -")
	err = chromedp.Run(ctx, view3(url, 100, &buf))
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(newpath, "clusters.png"), buf, 0644); err != nil {
		log.Fatal(err)
	}

	log.Printf("policies view -")
	err = chromedp.Run(ctx, view4(url, 100, &buf))
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(newpath, "policies.png"), buf, 0644); err != nil {
		log.Fatal(err)
	}
}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(sel, chromedp.ByID),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByID),
	}
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Liberally copied from puppeteer's source.
//
// Note: this will override the viewport emulation settings.
func fullScreenshot(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}

func login1(urlstr string, username string, password string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(`//input[@id='password']`),
		chromedp.SendKeys(`//input[@id='password']`, password),
		chromedp.SendKeys(`//input[@id='username']`, username),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}

// login screen
// fill form
// submit form @ name=loginButton
func login2(urlstr string, username string, password string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(`//input[@id='password']`),
		chromedp.SendKeys(`//input[@id='password']`, password),
		chromedp.SendKeys(`//input[@id='username']`, username),
		// chromedp.Submit(`//name=loginButton`),
		chromedp.Submit("//button[@type='submit']"),
		// chromedp.WaitNotPresent(`//input[@id='password']`),
		// chromedp.WaitVisible(`//*[@id='hamburger']`),
		// chromedp.WaitVisible(`//*[@id='page']`),
		chromedp.WaitVisible(`//*[@class='welcome']`),
		// chromedp.Click(`name=loginButton`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}

func overview(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	var overview = urlstr + `/multicloud/overview`
	log.Printf("urlstr: %s", overview)
	return chromedp.Tasks{
		chromedp.Navigate(overview),
		chromedp.WaitVisible(`//*[@class='overview-page']`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}

func view3(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	var view = urlstr + `/multicloud/clusters`
	log.Printf("view3 url: %s", view)
	return chromedp.Tasks{
		chromedp.Navigate(view),
		chromedp.WaitVisible(`//*[@class='page-content-container']`),
		chromedp.Sleep(1 * time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}

func view4(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	var view = urlstr + `/multicloud/policies`
	log.Printf("urlstr: %s", view)
	return chromedp.Tasks{
		chromedp.Navigate(view),
		chromedp.WaitVisible(`//*[@class='page-content-container']`),
		chromedp.Sleep(1 * time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}
