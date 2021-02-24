package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Content": "hello world"})
}

func content(ctx *gin.Context) {
	var buf bytes.Buffer

	input, err := ioutil.ReadFile(filepath.Clean("./articles/playground.md"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	md := goldmark.New(goldmark.WithExtensions(extension.GFM,extension.Typographer),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
			parser.WithBlockParsers(),
		),
		goldmark.WithRendererOptions(
			// html.WithUnsafe(),
			html.WithHardWraps(),
		),
	)

	if err = md.Convert(input, &buf); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.HTML(http.StatusOK, "index.html", gin.H{"Content": template.HTML(buf.String())})
}
