package grayscale

import (
	"fmt"
	"github.com/mawngo/isee/internal/scan"
	"github.com/spf13/cobra"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"log/slog"
	"math"
	"path/filepath"
	"runtime"
)

const DefaultRamps = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "

func NewCmd() *cobra.Command {
	f := flags{
		Ramps:         "",
		VerticalRatio: 35,
		Width:         80,
	}

	command := cobra.Command{
		Use:     "grayscale",
		Aliases: []string{"gray"},
		Short:   "Generate grayscale ascii art",
		Args:    cobra.MinimumNArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			scan.ProcessImg(args, runtime.NumCPU(), func(img scan.DecodedImage) {
				process(f, img)
			})
		},
	}

	command.Flags().StringVar(&f.Ramps, "ramps", f.Ramps, "Character ramps used to generate image")
	command.Flags().UintVarP(&f.Width, "width", "w", f.Width, "Width of the ascii art")
	command.Flags().IntVarP(&f.VerticalRatio, "vertical-ratio", "v", f.VerticalRatio, "Shrinking height of the image (to % of the original size)")
	return &command
}

type flags struct {
	Ramps         string
	Width         uint
	Contrast      int
	VerticalRatio int
}

func process(f flags, img scan.DecodedImage) {
	slog.Info("Processing",
		slog.String("img", filepath.Base(img.Path)),
		slog.String("dimension", fmt.Sprintf("%dx%d", img.Width(), img.Height())),
		slog.Uint64("width", uint64(f.Width)),
	)
	img = shrink(f, img)
	ramp := f.Ramps
	if ramp == "" {
		ramp = DefaultRamps
	}

	for y := range img.Height() {
		for x := range img.Width() {
			c := grayscale(img.At(x, y))
			fmt.Print(string(ramp[(len(ramp)-1)-len(ramp)*c/65536]))
		}
		fmt.Println()
	}
}

func grayscale(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return int(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
}

func shrink(f flags, img scan.DecodedImage) scan.DecodedImage {
	ratio := float64(f.Width) / float64(img.Width())
	height := int(math.Round(float64(img.Height()) * ratio))
	height = int(float64(height) * float64(f.VerticalRatio) / 100.0)
	resized := image.NewRGBA(image.Rect(0, 0, int(f.Width), height))
	draw.CatmullRom.Scale(resized, resized.Bounds(),
		img.Image,
		img.Image.Bounds(),
		draw.Src,
		nil)
	return scan.DecodedImage{
		Path:  img.Path,
		Image: resized,
	}
}
