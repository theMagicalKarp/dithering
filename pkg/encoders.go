package pkg

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

// jpegQuality is high enough that JPEG's block artifacts do not smear the
// dither pattern, which is all high-frequency detail.
const jpegQuality = 95

// EncodeFunc writes an image to w in a particular file format.
type EncodeFunc func(w io.Writer, m image.Image) error

// GetSupportedEncodings maps each output format name accepted on the command
// line to the encoder that writes it.
func GetSupportedEncodings() map[string]EncodeFunc {
	return map[string]EncodeFunc{
		"jpeg": JPEGEncode,
		"png":  PNGEncode,
	}
}

// JPEGEncode writes m to w as JPEG.
func JPEGEncode(w io.Writer, m image.Image) error {
	err := jpeg.Encode(w, m, &jpeg.Options{Quality: jpegQuality})
	if err != nil {
		return fmt.Errorf("failed to encode jpeg: %w", err)
	}

	return nil
}

// PNGEncode writes m to w as PNG.
func PNGEncode(w io.Writer, m image.Image) error {
	err := png.Encode(w, m)
	if err != nil {
		return fmt.Errorf("failed to encode png: %w", err)
	}

	return nil
}
