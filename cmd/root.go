// Package cmd wires the dither command line interface up to the dithering
// package.
package cmd

import (
	"errors"
	"fmt"
	"image"
	"os"

	"github.com/spf13/cobra"
	"github.com/theMagicalKarp/dithering/pkg"
)

// maxFactor is the largest quantization factor that still maps onto distinct
// 8-bit grayscale levels.
const maxFactor = 255

var (
	errMissingFiles = errors.New("please specify input/output files")
	errFactorRange  = errors.New("factor must be any value between 1-255")
	errOutputFormat = errors.New("not a valid output format, must be (jpeg|png)")
)

// options is a validated view of the command line flags.
type options struct {
	inFileName  string
	outFileName string
	factor      int
	encode      pkg.EncodeFunc
}

// RunE dithers the image named by the --in flag and writes the result to the
// file named by --out.
func RunE(cmd *cobra.Command, _ []string) error {
	inFileName, outFileName, err := parseFileNames(cmd)
	if err != nil {
		return err
	}

	factor, err := parseFactor(cmd)
	if err != nil {
		return err
	}

	encode, err := parseEncoder(cmd)
	if err != nil {
		return err
	}

	return dither(options{
		inFileName:  inFileName,
		outFileName: outFileName,
		factor:      factor,
		encode:      encode,
	})
}

// Execute builds the root command and runs it.
func Execute() error {
	//nolint:exhaustruct // cobra.Command is configured field by field, not built whole.
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

	err := rootCmd.Execute()
	if err != nil {
		return fmt.Errorf("CMD failure: %w", err)
	}

	return nil
}

// parseFileNames reads the --in and --out flags, both of which are required.
func parseFileNames(cmd *cobra.Command) (string, string, error) {
	inFileName, err := cmd.Flags().GetString("in")
	if err != nil {
		return "", "", fmt.Errorf("failed to read 'in' flag: %w", err)
	}

	outFileName, err := cmd.Flags().GetString("out")
	if err != nil {
		return "", "", fmt.Errorf("failed to read 'out' flag: %w", err)
	}

	if inFileName == "" || outFileName == "" {
		return "", "", errMissingFiles
	}

	return inFileName, outFileName, nil
}

// parseFactor reads the --factor flag and checks it against the range of
// quantization levels the dither supports.
func parseFactor(cmd *cobra.Command) (int, error) {
	factor, err := cmd.Flags().GetInt("factor")
	if err != nil {
		return 0, fmt.Errorf("failed to read 'factor' flag: %w", err)
	}

	if factor < 0 || factor > maxFactor {
		return 0, errFactorRange
	}

	return factor, nil
}

// parseEncoder resolves the --out-format flag to the encoder that writes it.
func parseEncoder(cmd *cobra.Command) (pkg.EncodeFunc, error) {
	outFormat, err := cmd.Flags().GetString("out-format")
	if err != nil {
		return nil, fmt.Errorf("failed to read 'out-format' flag: %w", err)
	}

	encode, ok := pkg.GetSupportedEncodings()[outFormat]
	if !ok {
		return nil, fmt.Errorf("(%s) is %w", outFormat, errOutputFormat)
	}

	return encode, nil
}

// dither runs the decode/dither/encode pipeline described by opts.
func dither(opts options) error {
	// The paths come from the user's own command line, so reaching any file
	// the user can already open is the intended behaviour.
	inFile, err := os.Open(opts.inFileName)
	if err != nil {
		return fmt.Errorf("failed to read 'in' file: %w", err)
	}

	defer func() { _ = inFile.Close() }()

	inImg, _, err := image.Decode(inFile)
	if err != nil {
		return fmt.Errorf("failed to decode 'in' file: %w", err)
	}

	grayScale := pkg.NewGreyScale(inImg.Bounds())
	pkg.ReadGreyScale(inImg, grayScale)
	pkg.ApplyDither(grayScale, opts.factor)

	outImg := image.NewNRGBA(inImg.Bounds())
	pkg.WriteGrayScale(grayScale, outImg)

	outFile, err := os.Create(opts.outFileName)
	if err != nil {
		return fmt.Errorf("failed to open 'out' file: %w", err)
	}

	defer func() { _ = outFile.Close() }()

	err = opts.encode(outFile, outImg)
	if err != nil {
		return fmt.Errorf("failed to encode 'out' file: %w", err)
	}

	// Closed explicitly as well as deferred, because a write that only fails
	// on flush would otherwise be swallowed and a truncated file reported as a
	// success. The deferred close then becomes a no-op whose error is moot.
	err = outFile.Close()
	if err != nil {
		return fmt.Errorf("failed to close 'out' file: %w", err)
	}

	return nil
}
