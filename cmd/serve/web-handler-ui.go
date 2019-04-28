package serve

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"mime"
	"net/http"
	nurl "net/url"
	"os"
	fp "path/filepath"
	"strconv"
	"strings"
	"bytes"

	"github.com/julienschmidt/httprouter"
)

// serveFiles serve files
func (h *webHandler) serveFiles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := serveFile(w, r.URL.Path)
	checkError(err)
}

// serveIndexPage is handler for GET /
func (h *webHandler) serveIndexPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check token
	err := h.checkToken(r)
	if err != nil {
		redirectPage(w, r, "/login")
		return
	}

	err = serveFile(w, "index.html")
	checkError(err)
}

// serveSubmitPage is handler for GET /submit
func (h *webHandler) serveSubmitPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := serveFile(w, "submit.html")
	checkError(err)
}

// serveLoginPage is handler for GET /login
func (h *webHandler) serveLoginPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check token
	err := h.checkToken(r)
	if err == nil {
		redirectPage(w, r, "/")
		return
	}

	err = serveFile(w, "login.html")
	checkError(err)
}

// serveBookmarkCache is handler for GET /bookmark/:id
func (h *webHandler) serveBookmarkCache(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get bookmark ID from URL
	strID := ps.ByName("id")
	id, err := strconv.Atoi(strID)
	checkError(err)

	// Get bookmarks in database
	bookmarks, err := h.db.GetBookmarks(true, id)
	checkError(err)

	if len(bookmarks) == 0 {
		panic(fmt.Errorf("No bookmark with matching index"))
	}

	// Create template
	funcMap := template.FuncMap{
		"html": func(s string) template.HTML {
			return template.HTML(s)
		},
		"hostname": func(s string) string {
			parsed, err := nurl.ParseRequestURI(s)
			if err != nil || len(parsed.Scheme) == 0 {
				return s
			}

			return parsed.Hostname()
		},
	}

	tplCache, err := createTemplate("cache.html", funcMap)
	checkError(err)

	bt, err := json.Marshal(&bookmarks[0])
	checkError(err)

	// Execute template
	strBt := string(bt)
	err = tplCache.Execute(w, &strBt)
	checkError(err)
}

// serveSearchPage is handler for GET /search?tag=test&keyword=rss
func (h *webHandler) serveSearchPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// There may be simple way to do this, but I don't know how, I am new to golang
	queryValues := r.URL.Query()
	tag := queryValues.Get("tag")
	keyword := queryValues.Get("keyword")
	serachText := "search: '"
	if tag != "" {
		tag = strings.Replace(tag, ",", " #", -1)
		serachText += "#" + tag + " "
	}
	if keyword != "" {
		serachText += keyword
	}
	serachText += "',"
	
	// Open file
	path := "index.html"
	src, err := assets.Open(path)
	checkError(err)
	defer src.Close()

	// Get content type
	ext := fp.Ext(path)
	mimeType := mime.TypeByExtension(ext)
	if mimeType != "" {
		w.Header().Set("Content-Type", mimeType)
	}

	// Serve file
	buf := new(bytes.Buffer)
	buf.ReadFrom(src)
	content := strings.Replace(buf.String(), "search: '',", serachText, 1)
	_, err = io.Copy(w, strings.NewReader(content))
	checkError(err)
}

// serveThumbnailImage is handler for GET /thumb/:id
func (h *webHandler) serveThumbnailImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get bookmark ID from URL
	id := ps.ByName("id")

	// Open image
	imgPath := fp.Join(h.dataDir, "thumb", id)
	img, err := os.Open(imgPath)
	checkError(err)
	defer img.Close()

	// Get image type from its 512 first bytes
	buffer := make([]byte, 512)
	_, err = img.Read(buffer)
	checkError(err)

	mimeType := http.DetectContentType(buffer)
	w.Header().Set("Content-Type", mimeType)

	// Serve image
	img.Seek(0, 0)
	_, err = io.Copy(w, img)
	checkError(err)
}

func serveFile(w http.ResponseWriter, path string) error {
	// Open file
	src, err := assets.Open(path)
	if err != nil {
		return err
	}
	defer src.Close()

	// Get content type
	ext := fp.Ext(path)
	mimeType := mime.TypeByExtension(ext)
	if mimeType != "" {
		w.Header().Set("Content-Type", mimeType)
	}

	// Serve file
	_, err = io.Copy(w, src)
	return err
}
