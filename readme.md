# link2xray

[![Go Reference](https://pkg.go.dev/badge/github.com/lunardoesdev/link2xray.svg)](https://pkg.go.dev/github.com/lunardoesdev/link2xray)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

`link2xray` is a Go library designed to bridge the gap between parsing Xray shared links and generating runtime-safe `libXray` configuration structures.

It addresses a specific, undocumented behavior in `libXray` where the `SendThrough` field must be explicitly handled (Issue #90) to prevent crashes during initialization.

## Features

*   **Simple Parsing:** Converts Xray shared links (VMess, VLESS, Trojan, etc.) directly into `*conf.Config` structs.
*   **Metadata Extraction:** Returns the subscription name (remark) alongside the configuration.
*   **Safety Handling:** Automatically sanitizes the configuration by handling the `SendThrough` field requirement described in [XTLS/libXray#90](https://github.com/XTLS/libXray/issues/90).
*   **Drop-in Replacement:** Returns the standard configuration objects used by `libXray`.

## Installation

```bash
go get github.com/lunardoesdev/link2xray
```

## Usage

Use the `SharedLinkToXrayConfig` function to parse a link and obtain a configuration object that is safe to pass to `libXray`.

```go
package main

import (
	"fmt"
	"os"

	"github.com/lunardoesdev/link2xray"
)

const exampleLink = "ss://Y2hhY2hhMjAtaWV0Zi1wb2x5MTMwNTpjdklJODVUclc2bjBPR3lmcEhWUzF1@45.87.175.187:8080#%3E%3E%40FreakConfig%3A%3AXX"

func main() {
	name, config, err := link2xray.SharedLinkToXrayConfig(exampleLink)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Configration found: %s\n", name)
	fmt.Printf("%+v\n", config.OutboundConfigs[0])
}
```

## API Reference

### `func SharedLinkToXrayConfig(link string) (name string, config *conf.Config, err error)`

Parses a shared link string and returns a sanitized configuration ready for `libXray`.

*   **`link`**: The shared link string (e.g., `vmess://...`, `vless://...`).
*   **`name`**: The "remark" or subscription name extracted from the link.
*   **`config`**: A pointer to the `conf.Config` struct. **Note:** This object is pre-processed to handle the `SendThrough` field issue, preventing runtime panics reported in [libXray#90](https://github.com/XTLS/libXray/issues/90).
*   **`err`**: An error if the link format is invalid or parsing fails.

## Motivation

When using the official `XTLS/libXray` library, simply parsing a configuration is often not enough. 

As reported in [XTLS/libXray#90](https://github.com/XTLS/libXray/issues/90), if the `SendThrough` field is not explicitly set to `nil` before initializing the instance, the core may fail to start properly. This behavior is currently undocumented in the core library. 

`link2xray` abstracts this complexity away, ensuring that developers don't have to manually patch their configuration structs every time they parse a link.

## Contributing

Contributions are welcome! Please feel free to open a Pull Request.

## License

This project is licensed under the MIT License
