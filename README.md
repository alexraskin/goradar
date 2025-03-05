# GoRadar

A command-line interface for tracking aircraft using the adsb.lol API.

## Features

- Search aircraft by various criteria:
  - Registration number
  - ICAO hex code
  - Aircraft type
  - Squawk code
  - Military aircraft
  - PIA (Privacy ICAO Address) aircraft
  - LADD (Limiting Aircraft Data Displayed) aircraft
- Pagination support for large result sets
- Colored output for better readability
- FlightRadar24 integration

## Installation

### macOS (via Homebrew)

```bash
brew tap alexraskin/homebrew-tap
brew install goradar
```

### Linux

Download the latest release for your architecture from the [GitHub releases page](https://github.com/alexraskin/goradar/releases):

```bash
# For x86_64
wget https://github.com/alexraskin/goradar/releases/latest/download/goradar_Linux_x86_64.tar.gz
tar -xzf goradar_Linux_x86_64.tar.gz
sudo mv goradar /usr/local/bin/

# For ARM64
wget https://github.com/alexraskin/goradar/releases/latest/download/goradar_Linux_arm64.tar.gz
tar -xzf goradar_Linux_arm64.tar.gz
sudo mv goradar /usr/local/bin/
```

### Windows

Download the latest release for Windows from the [GitHub releases page](https://github.com/alexraskin/goradar/releases):

1. Download `goradar_Windows_x86_64.zip`
2. Extract the zip file
3. Add the extracted directory to your system's PATH environment variable

## Usage

### Search by Registration

```bash
goradar registration G-KELS
```

### Search by ICAO Hex

```bash
goradar hex 4CA87C
```

### Search by Aircraft Type

```bash
goradar type A320
```

### Search by Squawk Code

```bash
goradar squawk 7700
```

### List Military Aircraft

```bash
goradar military
```

### List PIA Aircraft

```bash
goradar pia
```

### List LADD Aircraft

```bash
goradar ladd
```

### Pagination

All commands support pagination using the `--limit` and `--offset` flags:

```bash
# Show 5 results per page
goradar military --limit 5

# Show 5 results starting from offset 10
goradar military --limit 5 --offset 10
```

## Development

### Building from Source

```bash
git clone https://github.com/alexraskin/goradar.git
cd goradar
go build
```

### Running Tests

```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 