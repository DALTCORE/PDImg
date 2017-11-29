package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/cli"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	app := cli.NewApp()

	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Ramon Smit",
			Email: "rsmit@daltcore.com",
		},
	}
	app.Name = "pdimg"
	app.HelpName = "pdimg"
	app.Usage = "Convert PDF's to JPEG, PNG or WebP"
	app.UsageText = "pdimg convert <filename.pdf> jpeg"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:  "convert",
			Usage: "Convert PDF to ",
			Action: func(c *cli.Context) error {

				if c.Args().Get(0) == "" {
					color.Red("First parameter (pdf) is not set!")
					os.Exit(1)
				}
				pdf := c.Args().Get(0)

				if c.Args().Get(1) == "" {
					color.Red("First parameter (pdf) is not set!")
					os.Exit(1)
				}
				imgType := c.Args().Get(1)

				fmt.Println("Loading PDF:", pdf)

				imagick.Initialize()
				defer imagick.Terminate()
				mw := imagick.NewMagickWand()
				defer mw.Destroy()

				mw.SetResolution(200, 200)

				mw.ReadImage(pdf)

				totalPages := mw.GetNumberImages()

				fmt.Println("Total pages: ", totalPages)

				for i := 0; i < int(totalPages); i++ {
					fmt.Println("Transforming page: ", i)

					mw.SetIteratorIndex(int(i)) // This being the page offset
					mw.SetImageFormat(imgType)
					mw.SetCompressionQuality(60)
					mw.SetCompression(0)
					mw.SetImageCompressionQuality(60)

					pw := imagick.NewPixelWand()
					defer pw.Destroy()
					pw.SetColor("white")
					mw.SetImageBackgroundColor(pw)
					mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_REMOVE)

					mw.SetImageResolution(50, 50)
					fmt.Print(fmt.Sprintf("> page-%v.jpg\n", i))
					errorr := mw.WriteImage(fmt.Sprintf("page-%v.jpg", i))
					if errorr != nil {
						fmt.Print(errorr)
					}
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}

func resize() {

}
