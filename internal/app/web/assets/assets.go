package assets

import "embed"

var (
	//go:embed *.html
	HTML embed.FS
	//go:embed static/js/*
	JS embed.FS
	//go:embed static/css/*
	CSS embed.FS
	//go:embed static/img/*
	IMG embed.FS
)
