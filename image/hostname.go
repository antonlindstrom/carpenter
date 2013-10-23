package image

import (
	"io/ioutil"
)

func (i *Image) ChangeHostname(hostname string) error {
	err := ioutil.WriteFile(i.Mountpoint+"/etc/hostname", []byte(hostname), 0644)

	return err
}
