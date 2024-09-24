package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/szh/naturewatch-fetcher/pkg/util"
)

func ListPhotos() ([]string, error) {
	url := constructURL("/data/photos")
	return listFiles(url)
}

func ListVideos() ([]string, error) {
	url := constructURL("/data/videos")
	return listFiles(url)
}

func DownloadFile(path string) ([]byte, error) {
	url := constructURL("/data/" + path)
	httpClient := http.DefaultClient
	req := &http.Request{
		Method: http.MethodGet,
		URL:    &url,
		Header: http.Header{
			"Accept": {"application/octet-stream"},
		},
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Cannot download file: %v", err)
	}
	// Read response body as a byte array
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot read response body: %v", err)
	}
	return data, nil
}

func DeleteFile(path string) error {
	url := constructURL("/data/" + path)
	httpClient := http.DefaultClient
	req := &http.Request{
		Method: http.MethodDelete,
		URL:    &url,
	}
	_, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("Cannot delete file: %v", err)
	}
	return nil
}

func listFiles(url url.URL) ([]string, error) {
	httpClient := http.DefaultClient
	req := &http.Request{
		Method: http.MethodGet,
		URL:    &url,
		Header: http.Header{
			"Accept": {"application/json"},
		},
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Cannot fetch files: %v", err)
	}
	// Read response body as a JSON array of strings (filepaths)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot read response body: %v", err)
	}
	var files []string
	err = json.Unmarshal(body, &files)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse response body: %v", err)
	}
	return files, nil
}

// Normalizes the URL by parsing the base URL and appending the subpath
func constructURL(subpath string) url.URL {
	url, err := url.Parse(util.Config.NatureWatchURL)
	if err != nil {
		panic(fmt.Errorf("Cannot parse URL: %v", err))
	}
	url.Path = subpath
	return *url
}
