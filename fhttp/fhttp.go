package fhttp

import (
	"github.com/jung-kurt/gofpdf"
	"net/http"
)

// RegisterRemoteImage registers a remote image. Downloading the image from the
// provided URL and adding it to the PDF but not adding it to the page. Use
// Image() with the same URL to add the image to the page.
func RegisterRemoteImage(f *gofpdf.Fpdf, urlStr, tp string) (info *gofpdf.ImageInfoType) {
	resp, err := http.Get(urlStr)

	if err != nil {
		f.SetError(err)
		return
	}

	defer resp.Body.Close()

	if tp == "" {
		tp = f.ImageTypeFromMime(resp.Header["Content-Type"][0])
	}

	return f.RegisterImageReader(urlStr, tp, resp.Body)
}