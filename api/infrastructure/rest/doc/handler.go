package doc

import (
	"bytes"
	"github.com/go-chi/chi/v5"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

func renderMarkdown(filePath string) ([]byte, error) {
	source, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Table,
			extension.Typographer,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	if err := md.Convert(source, &buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type Handler interface {
	HandleDocRequest(w http.ResponseWriter, r *http.Request)
}

type RestHandler struct {
	docBasePathDir string
}

func NewHandler(docBasePathDir string) Handler {
	return &RestHandler{docBasePathDir: docBasePathDir}
}

func (rh RestHandler) HandleDocRequest(w http.ResponseWriter, r *http.Request) {
	var file string

	if filename := chi.URLParam(r, "filename"); filename != "" {
		cleanFilename := filepath.Clean(filename)
		if !strings.HasSuffix(cleanFilename, ".md") {
			http.Error(w, "Invalid file type", http.StatusForbidden)
			return
		}
		file = filepath.Clean(rh.docBasePathDir + cleanFilename)
		if !strings.HasPrefix(file, rh.docBasePathDir) {
			http.Error(w, "File not allowed", http.StatusForbidden)
			return
		}
	} else {
		file = "docs/api.md"
	}

	html, err := renderMarkdown(file)
	if err != nil {
		http.Error(w, "Error when processed file: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`<!DOCTYPE html>
<html lang="es">
<head>
	<meta charset="UTF-8">
	<title>Documentación</title>
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<style>
		body {
			font-family: 'Segoe UI', sans-serif;
			background-color: #1e1e1e;
			color: #d4d4d4;
			margin: 0;
			padding: 2rem;
			max-width: 800px;
			margin-left: auto;
			margin-right: auto;
			line-height: 1.6;
		}
		h1, h2, h3 {
			color: #569cd6;
			border-bottom: 1px solid #333;
			padding-bottom: 0.3rem;
		}
		a {
			color: #4fc1ff;
		}
		code {
			background-color: #2d2d2d;
			padding: 2px 4px;
			border-radius: 4px;
			font-family: monospace;
		}
		pre {
			background-color: #2d2d2d;
			padding: 1rem;
			overflow-x: auto;
			border-radius: 8px;
		}
		table {
			width: 100%;
			border-collapse: collapse;
			margin-top: 1rem;
		}
		th, td {
			border: 1px solid #444;
			padding: 8px;
			text-align: left;
		}
		th {
			background-color: #333;
		}
		tr:nth-child(even) {
			background-color: #2a2a2a;
		}
		hr {
			border: 0;
			border-top: 1px solid #444;
			margin: 2rem 0;
		}
	</style>
</head>
<body>`))
	w.Write(html)
	w.Write([]byte(`</body></html>`))
}
