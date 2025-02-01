package html

import (
	"bytes"
	"context"
	"os"
	"strings"

	"golang.org/x/net/html"

	"io"

	"github.com/carlmjohnson/requests"
	"github.com/shammianand/rtt/utils/logger"
)

func ParseHTML(ctx context.Context, url string) error {
	logger.Log.Info("Parsing HTML from URL: ", url)
	var content string
	err := requests.URL(url).ToString(&content).Fetch(ctx)
	if err != nil {
		return err
	}

	text, err := extractBodyText(content)
	if err != nil {
		return err
	}
	os.WriteFile("rtt-html.txt", []byte(text), 0644)

	return nil
}

// a simple function to extract body text from a url
func extractBodyText(content string) (string, error) {
	r := strings.NewReader(content)
	z := html.NewTokenizer(r)

	var buffer strings.Builder
	inBody := false

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			if z.Err() == io.EOF {
				break
			}
			return "", z.Err()
		}

		switch tt {
		case html.StartTagToken:
			name, _ := z.TagName()
			if bytes.Equal(name, []byte("body")) {
				inBody = true
			}
		case html.EndTagToken:
			name, _ := z.TagName()
			if bytes.Equal(name, []byte("body")) {
				inBody = false
			}
		case html.TextToken:
			if inBody {
				text := strings.TrimSpace(string(z.Text()))
				if text != "" {
					buffer.WriteString(text)
					buffer.WriteString("\n")
				}
			}
		}
	}

	return buffer.String(), nil
}
