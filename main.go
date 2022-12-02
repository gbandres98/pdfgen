package main

import (
	"image"
	"log"
	"os"
	"path"

	"github.com/signintech/gopdf"
)

func main() {
	joinImages("./in", "./out.pdf")
}

func joinImages(in string, out string) {
	files, err := os.ReadDir(in)
	if err != nil {
		log.Println(err)
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	for _, f := range files {
		imgPath := path.Join(in, f.Name())
		r, err := os.Open(imgPath)
		if err != nil {
			log.Println(err)
			continue
		}

		i, _, err := image.Decode(r)
		if err != nil {
			log.Println(err)
			continue
		}

		pageSize := &gopdf.Rect{W: float64(i.Bounds().Dx()), H: float64(i.Bounds().Dy())}

		pdf.AddPageWithOption(gopdf.PageOption{PageSize: pageSize})
		pdf.Image(imgPath, 0, 0, pageSize)
	}

	err = pdf.WritePdf(out)
	if err != nil {
		log.Println(err)
	}
}
