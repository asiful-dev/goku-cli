# Goku CLI

Goku is a small Go-based CLI for converting between JSON and YAML. This repository currently covers **Phase 1** only: file conversion with automatic output-file creation.

## Supported Platforms

- macOS
- Linux

Windows is not supported for this project setup.

## Requirements

- Go 1.26 or newer if you want to build from source or install with `go install`
- A terminal on macOS or Linux

## Installation

### 1. Install with Go

If you want to install it directly from source using Go:

```bash
go install github.com/asiful-dev/goku@latest
```

After installation, make sure your Go bin directory is in `PATH`, then run:

```bash
goku --help
```

### 2. Install from a Binary Release

If you publish release assets, download the archive for your platform and extract it.

Example for Linux:

```bash
tar -xzf goku-linux-amd64.tar.gz
chmod +x goku
sudo mv goku /usr/local/bin/
```

Example for macOS:

```bash
tar -xzf goku-darwin-amd64.tar.gz
chmod +x goku
sudo mv goku /usr/local/bin/
```

### 3. Build from Source

Clone the repository and build the binary locally:

```bash
git clone https://github.com/asiful-dev/goku.git
cd goku
go build -o goku .
```

Then run the local binary:

```bash
./goku --help
```

## Phase 1 Features

- Convert JSON to YAML
- Convert YAML to JSON
- Save converted output to a file automatically
- Use the input filename by default for the output file
- Override the output filename with `-n` or `--name`

## Usage

```bash
goku -i input.json -o yaml
goku -i input.yaml -o json
goku -i input.json -o yaml -n custom-output-filename
```

## Examples

Convert JSON to YAML:

```bash
goku -i test.json -o yaml
```

Convert YAML to JSON:

```bash
goku -i test.yaml -o json
```

## Output File Behavior

- If you do not provide an output name, Goku uses the input filename and changes the extension.
- Example: `test.json` becomes `test.yaml`.
- If you provide `-n custom-output`, Goku writes `custom-output.yaml` or `custom-output.json` depending on the target format.

## Project Status

Phase 1 is complete. Phase 2 (PostgreSQL CRUD commands) is planned separately.
