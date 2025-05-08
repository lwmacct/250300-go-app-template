package app

import "embed"

//go:embed embed/*
var Embed embed.FS

var Flag *TsFlag = &TsFlag{}
