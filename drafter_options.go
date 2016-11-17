package main

type DrafterOptions struct {
	sourcemap bool
	drafterFormat DrafterFormat
}

func NewDrafterFormat(sourcemap bool, drafterFormat DrafterFormat) DrafterOptions {
	return DrafterOptions{sourcemap, drafterFormat}
}
