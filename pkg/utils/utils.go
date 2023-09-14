package utils

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

// Download the file
func DownlowdFile(url, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("there is errr : %v", err)
		return err
	}
	defer resp.Body.Close()

	// create a file
	os.RemoveAll(path)
	os.Mkdir("dataset", 0750)

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Errorf("there is error in opening file: %v", err)
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Debugf("there is error while reading response: %v", err)
		return err
	}
	file.Write(data)
	log.Infof("written output to file: %s", file.Name())
	file.Close()
	return nil
}

// extract file
func ExtractFile(filename string, dst string) {
	gzipstream, err := os.Open(filename)
	if err != nil {
		log.Errorf("error while reading tar : %v", err)
		return
	}
	defer gzipstream.Close()

	uncompressedStream, err := gzip.NewReader(gzipstream)
	if err != nil {
		log.Errorf("ExtractTarGz: NewReader failed")
	}

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Errorf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeReg:
			f, err := os.Create(dst + "/" + header.Name)
			if err != nil {
				log.Errorf("create failed : %v", err)
			}
			io.Copy(f, tarReader)
			f.Close()
		}
	}
}
