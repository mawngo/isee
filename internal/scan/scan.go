package scan

import (
	"fmt"
	_ "golang.org/x/image/bmp"  // Enable support for bmp.
	_ "golang.org/x/image/webp" // Enable support for webp.
	"image"
	_ "image/jpeg" // Enable support for jpeg.
	_ "image/png"  // Enable support for png.
	"log/slog"
	"os"
	"path/filepath"
)

func Img(dir string) <-chan DecodedImage {
	ch := make(chan DecodedImage, 1)
	info, err := os.Stat(dir)
	if err != nil {
		slog.Error("Err scanning file(s)", slog.String("path", dir), slog.Any("err", err))
		close(ch)
		return ch
	}

	go func() {
		defer close(ch)
		if !info.IsDir() {
			img, err := decode(dir)
			if err != nil {
				slog.Error("Err decoding image", slog.String("path", dir), slog.Any("err", err))
				return
			}
			ch <- img
			return
		}

		files, err := os.ReadDir(".")
		if err != nil {
			slog.Error("Err scanning dir", slog.String("path", dir), slog.Any("err", err))
			return
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}
			path := filepath.Join(dir, file.Name())
			img, err := decode(path)
			if err != nil {
				slog.Error("Not a image", slog.String("path", path), slog.Any("err", err))
				continue
			}
			ch <- img
		}
	}()
	return ch
}

func ProcessImg(dirs []string, concurrency int, processFn func(img DecodedImage)) {
	con := make(chan struct{}, concurrency)
	for _, arg := range dirs {
		for img := range Img(arg) {
			con <- struct{}{}
			go func() {
				defer func() {
					<-con
				}()
				processFn(img)
			}()
		}
	}
	for range concurrency {
		con <- struct{}{}
	}
}

func decode(path string) (DecodedImage, error) {
	img := DecodedImage{
		Path: path,
	}
	f, err := os.Open(path)
	if err != nil {
		return img, err
	}
	defer f.Close()

	config, _, err := image.DecodeConfig(f)
	if err != nil {
		return img, err
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	slog.Debug("Decoding image", slog.String("path", path), slog.String("dimension", fmt.Sprintf("%dx%d", config.Width, config.Height)))
	imageData, _, err := image.Decode(f)
	if err != nil {
		return img, err
	}
	img.Image = imageData

	return img, nil
}

type DecodedImage struct {
	image.Image
	Path string
}

func (i DecodedImage) Width() int {
	return i.Bounds().Dx()
}

func (i DecodedImage) Height() int {
	return i.Bounds().Dy()
}
