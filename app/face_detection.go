package app

import (
	"log"

	"autolock/vars"

	"github.com/Kagami/go-face"
)

var (
	locked = false
	count  = 0
)

func FaceDetect(image string) error {
	det := detect(image)

	if det && locked {
		count = 0
		locked = false
		if vars.Config.Verbose {
			log.Println("face detected")
		}
	} else if !det && count < vars.Config.Threshold && !locked {
		count++
		if vars.Config.Verbose {
			log.Println(count)
		}
	} else if !det && count >= vars.Config.Threshold && !locked {
		count = 0
		locked = true
		if vars.Config.Verbose {
			log.Println("Locked")
		}
		if err := RunCommand(); err != nil {
			return err
		}
	} else {
		count = 0
	}

	return nil
}

func detect(image string) bool {
	rec, err := face.NewRecognizer(vars.Config.Path.Models)
	if err != nil {
		if vars.Config.Verbose {
			log.Println(err.Error())
		}
		return false
	}
	defer rec.Close()

	faces, err := rec.RecognizeFile(image)
	if err != nil {
		if vars.Config.Verbose {
			log.Println(err.Error())
		}
		return false
	}

	if len(faces) == 0 || len(faces) > 1 {
		return false
	}

	return true
}
