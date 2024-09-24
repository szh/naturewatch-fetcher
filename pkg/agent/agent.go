package agent

import (
	"log"
	"os"
	"time"

	"github.com/szh/naturewatch-fetcher/pkg/api"
	"github.com/szh/naturewatch-fetcher/pkg/util"
)

func Start() {
	for {
		err := doFetch()
		if err != nil {
			log.Printf("[ERROR] %v", err)
		}

		// Sleep the configured number of seconds
		time.Sleep(time.Duration(util.Config.FetchIntervalSeconds) * time.Second)
	}
}

func doFetch() error {
	photos, err := api.ListPhotos()
	if err != nil {
		return err
	}

	log.Printf("Listed %d photos", len(photos))

	for _, photo := range photos {
		err := processFile("photos/" + photo)
		if err != nil {
			return err
		}
	}

	videos, err := api.ListVideos()
	if err != nil {
		return err
	}

	log.Printf("Listed %d vidoes", len(photos))

	for _, video := range videos {
		err := processFile("videos/" + video)
		if err != nil {
			return err
		}
	}

	return nil
}

func processFile(path string) error {
	data, err := api.DownloadFile(path)
	if err != nil {
		return err
	}

	log.Printf("Downloaded %s", path)

	err = saveFile(path, data)
	if err != nil {
		return err
	}

	log.Printf("Saved %s", path)

	err = api.DeleteFile(path)
	if err != nil {
		return err
	}

	log.Printf("Deleted %s", path)

	return nil
}

func saveFile(path string, data []byte) error {
	outputPath := util.Config.OutputPath + "/" + path
	return os.WriteFile(outputPath, data, 0666)
}
