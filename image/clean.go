package image

import (
	"log"
	"os"
	"path/filepath"
)

func (i *Image) CleanHostKeys() error {
	match, err := filepath.Glob(i.Mountpoint + "/etc/ssh/*_host_*")

	if err != nil {
		log.Printf("Error in looking up path: %s\n", err)
	}

	for _, path := range match {
		log.Printf("Cleaning: %s\n", path)
		err = os.Remove(path)

		if err != nil {
			log.Printf("Could not remove file %s: %s\n", path, err)
		}
	}

	return err
}
