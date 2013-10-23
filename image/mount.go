package image

// Relying on http://en.wikibooks.org/wiki/QEMU/Images#Mounting_an_image_on_the_host

import (
	"github.com/antonlindstrom/carpenter/cmd"
	"log"
)

func (i *Image) Mount() error {
	err := cmd.Run("qemu-nbd", "-c", i.Device, i.File)

	if err != nil {
		log.Printf("qemu-nbd -c %s %s\n", i.Device, i.File)
		log.Fatal("Error, could not execute qemu-nbd\n", err)
	}

	cmd.Run("partprobe", i.Device)

	err = cmd.Run("mount", i.Device+i.Partition, i.Mountpoint)

	if err != nil {
		log.Printf("Error during mount, let's not care..\n")
		return err
	}

	return err
}

func (i *Image) UnMount() error {
	err := cmd.Run("umount", i.Mountpoint)

	if err != nil {
		log.Fatal("Error, could not unmount (umount).\n")
		return err
	}

	err = cmd.Run("qemu-nbd", "-d", i.Device)

	if err != nil {
		log.Fatal("Error, could not unmount (qemu-nbd -d).\n")
	}

	return err
}
