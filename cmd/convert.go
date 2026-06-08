package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/asiful-dev/goku/internal/converter"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
    inputFile    string
    outputFormat string
    outputFileName string
)

var convertCmd = &cobra.Command{
    Use:   "convert",
    Short: "Convert between JSON and YAML formats",
    RunE:  runConvert,
}

func init() {
    // Allow using flags at root level as well
    rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Path to the input file (required)")
    rootCmd.Flags().StringVarP(&outputFormat, "output", "o", "", "Output format: json or yaml (required)")
    rootCmd.Flags().StringVarP(&outputFileName, "name", "n", "", "Optional output file name")
    _ = rootCmd.MarkFlagRequired("input")
    _ = rootCmd.MarkFlagRequired("output")
    rootCmd.RunE = runConvert
    rootCmd.AddCommand(convertCmd)
}

func runConvert(cmd *cobra.Command, args []string) error {
    ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(inputFile), "."))
    if ext == "yml" {
        ext = "yaml"
    }
    outFmt := strings.ToLower(outputFormat)

    if ext == outFmt {
        return fmt.Errorf("requested output format (%s) must be different from input format (%s)", outFmt, ext)
    }

    if ext != "json" && ext != "yaml" {
        return fmt.Errorf("unsupported input format: .%s (only .json or .yaml/.yml allowed)", ext)
    }
    if outFmt != "json" && outFmt != "yaml" {
        return fmt.Errorf("unsupported output format: %s (use 'json' or 'yaml')", outFmt)
    }

    data, err := os.ReadFile(inputFile)
    if err != nil {
        return fmt.Errorf("could not read file '%s': %w", inputFile, err)
    }

    var result []byte
    if ext == "json" && outFmt == "yaml" {
        result, err = converter.JSONToYAML(data)
    } else {
        result, err = converter.YAMLToJSON(data)
    }
    if err != nil {
        return err
    }

    outputPath := buildOutputPath(inputFile, outputFileName, outFmt)
    if err := os.WriteFile(outputPath, result, 0o644); err != nil {
        return fmt.Errorf("could not write output file '%s': %w", outputPath, err)
    }

    color.Green("Conversion successful! Input: .%s → Output: .%s\n", ext, outFmt)
    color.Green("Saved to: %s\n\n", outputPath)
    color.Cyan("--- Result ---\n")
    fmt.Println(string(result))
    return nil
}

func buildOutputPath(inputPath, requestedName, outputFormat string) string {
    inputDir := filepath.Dir(inputPath)
    inputBase := filepath.Base(inputPath)
    inputExt := filepath.Ext(inputBase)
    baseName := strings.TrimSuffix(inputBase, inputExt)

    fileName := requestedName
    if fileName == "" {
        fileName = baseName + "." + outputFormat
    } else {
        ext := filepath.Ext(fileName)
        if ext == "" {
            fileName += "." + outputFormat
        }
    }

    if filepath.IsAbs(fileName) {
        return fileName
    }
    return filepath.Join(inputDir, fileName)
}
