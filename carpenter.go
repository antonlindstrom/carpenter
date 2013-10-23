package main

import (
	"flag"
	"fmt"
	"github.com/antonlindstrom/carpenter/image"
	"log"
	"os"
)

func main() {

	hostname := flag.String("hostname", "www.example.com", "Hostname for new server")
	file := flag.String("image", "/var/image.qcow2", "Image to scrub")
	flag.Parse()

	if flag.NFlag() != 2 {
		fmt.Printf("USAGE: %s -hostname=HOSTNAME -image=/path/to/image.qcow2\n", os.Args[0])
		os.Exit(2)
	}

	mountpoint := "/mnt/" + *hostname

	os.Mkdir(mountpoint, 0644)

	i := image.Image{*hostname, mountpoint, "/dev/nbd0", "p1", *file}

	log.Printf("Mounting..\n")
	i.Mount()
	log.Printf("Cleaning host keys..\n")
	i.CleanHostKeys()
	log.Printf("Changing hostname..\n")
	i.ChangeHostname(*hostname)
	log.Printf("Unmounting..\n")
	i.UnMount()
	os.Remove(mountpoint)
}
