package pkg

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

type EncodeFunc func(w io.Writer, m image.Image) error

func GetSupportedEncodings() map[string]EncodeFunc {
	return map[string]EncodeFunc{
		"jpeg": JPEGEncode,
		"png":  PNGEncode,
	}
}

func JPEGEncode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, &jpeg.Options{Quality: 95})
}

func PNGEncode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}
