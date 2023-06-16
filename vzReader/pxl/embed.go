package pxl

import (
	"embed"
)

//go:embed *.pxl
var _ embed.FS
