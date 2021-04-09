package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/mastermeng/goweb/models"
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

	md := goldmark.New(goldmark.WithExtensions(extension.GFM, extension.Typographer),
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
	ctx.HTML(http.StatusOK, "details.html", gin.H{"Content": template.HTML(buf.String())})
}

func handleFavicon(ctx *gin.Context) {
	http.ServeFile(ctx.Writer, ctx.Request, "./static/favicon.ico")
}

func handlePost(ctx *gin.Context) {
	article := &models.Atricles{}
	if err := ctx.ShouldBindJSON(article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
}
