package app

import (
	"errors"
	"log"
	"os"
	"time"

	"autolock/vars"

	"github.com/radovskyb/watcher"
)

func Run(imagesPath string) error {
	if _, err := os.Stat(imagesPath); os.IsNotExist(err) {
		return errors.New(imagesPath + " not found")
	}

	w := watcher.New()
	w.FilterOps(watcher.Write)

	go func() {
		for {
			select {
			case event := <-w.Event:
				if !event.IsDir() {
					if err := FaceDetect(imagesPath); err != nil {
						if vars.Config.Verbose {
							log.Fatal(err)
						}
					}
				}
			case err := <-w.Error:
				if vars.Config.Verbose {
					log.Fatal(err)
				}
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Add(imagesPath); err != nil {
		return err
	}

	go func() {
		w.Wait()
	}()

	return w.Start(time.Second)
}
