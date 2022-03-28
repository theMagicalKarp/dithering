package cmd

import (
	"fmt"
	"image"
	"os"

	"github.com/spf13/cobra"
	"github.com/theMagicalKarp/dithering/pkg"
)

func RunE(cmd *cobra.Command, args []string) error {
	inFileName, err := cmd.Flags().GetString("in")
	if err != nil {
		return fmt.Errorf("Failed to read 'in' flag: %w", err)
	}

	outFileName, err := cmd.Flags().GetString("out")
	if err != nil {
		return fmt.Errorf("Failed to read 'out' flag: %w", err)
	}

	if inFileName == "" || outFileName == "" {
		return fmt.Errorf("Please specify input/output files")
	}

	factor, err := cmd.Flags().GetInt("factor")
	if err != nil {
		return fmt.Errorf("Failed to read 'factor' flag: %w", err)
	}

	if factor < 0 || factor > 255 {
		return fmt.Errorf("Factor must be any value between 1-255")
	}

	validOutEncodings := pkg.GetSupportedEncodings()
	outFormat, err := cmd.Flags().GetString("out-format")
	if err != nil {
		return fmt.Errorf("Failed to read 'out-format' flag: %w", err)
	}

	encodeFunc, ok := validOutEncodings[outFormat]
	if !ok {
		return fmt.Errorf("(%s) is not valid output format must be (jpeg|png)", outFormat)
	}

	inFile, err := os.Open(inFileName)
	if err != nil {
		return fmt.Errorf("Failed to read 'in' file: %w", err)
	}
	defer inFile.Close()

	inImg, _, err := image.Decode(inFile)
	if err != nil {
		return fmt.Errorf("Failed to decode 'in' file: %w", err)
	}

	grayScale := pkg.NewGreyScale(inImg.Bounds())
	pkg.ReadGreyScale(inImg, grayScale)
	pkg.ApplyDither(grayScale, factor)
	outImg := image.NewNRGBA(inImg.Bounds())
	pkg.WriteGrayScale(grayScale, outImg)

	outFile, err := os.Create(outFileName)
	if err != nil {
		return fmt.Errorf("Failed to open 'out' file: %w", err)
	}
	defer outFile.Close()

	err = encodeFunc(outFile, outImg)
	if err != nil {
		return fmt.Errorf("Failed to encode 'out' file: %w", err)
	}

	return nil
}

func Execute() error {
	rootCmd := &cobra.Command{
		Use:   "dither",
		Short: "Apply basic Floyd–Steinberg technique to a given image.",
		RunE:  RunE,
	}

	rootCmd.Flags().StringP("in", "i", "", "Input file to use as source")
	rootCmd.Flags().StringP("out", "o", "", "Output file to use as destination")
	rootCmd.Flags().StringP("out-format", "u", "png", "Out Format to use when writing to destination")
	rootCmd.Flags().IntP(
		"factor", "f", 1,
		"The factor level of quantization (e.g. 1=2bits of color/2=3bits of color/etc...)",
	)

	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("CMD failure: %w", err)
	}

	return nil
}
