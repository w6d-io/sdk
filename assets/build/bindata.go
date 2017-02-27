// Code generated by go-bindata.
// sources:
// etc/build/Makefile
// etc/build/etc/Dockerfile.build.alpine
// etc/build/etc/Dockerfile.build.debian
// etc/build/make/bootstrap.mk
// etc/build/make/functions.mk
// etc/build/make/rules.mk
// DO NOT EDIT!

package build

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _makefile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x50\xbb\x6e\xdb\x40\x10\xec\xef\x2b\x06\xb0\x0a\x11\x30\x25\x20\x65\x80\x18\x30\xe0\xc0\x85\x93\x34\x4e\x7a\x1e\x79\x7b\xd6\x85\xf7\xca\xde\x52\xb1\xff\x3e\x38\x2a\x12\x99\x14\x56\xb9\x3b\xb3\xb3\x33\x73\x83\xef\x07\x57\xe0\x0a\xe4\x40\xa0\x28\xfc\x96\x93\x8b\x82\x64\xe7\x4d\xdf\x7b\x5b\x0e\x08\x7a\x24\x3c\x3f\x3c\xdd\xce\xcb\xe7\x87\x27\x94\x43\x9a\xbc\x41\x4f\x28\x92\x98\x0c\xb4\xa8\x1b\x74\xc5\x8c\x3e\x0d\x5a\x5c\x8a\xdd\x2d\x5e\x28\x12\x6b\x21\x03\x7a\xa5\x61\x12\x17\x5f\xd0\x9d\x24\xdb\x62\x46\x64\xa6\xac\x99\xda\x7e\x72\xde\x74\x3b\x7c\xd5\xf1\x0d\xc6\x59\x4b\x4c\xb1\xea\x1d\x35\x17\xe8\x82\xee\xcb\xfd\xb7\xc7\x1f\xf7\x8f\x9f\x3b\x68\x26\x30\xfd\x9a\x1c\x93\xa9\x76\x5c\xa9\x2c\xa7\x7b\x4f\x65\x06\xe9\x55\x58\x0f\xf3\x2f\xcb\x29\x54\xc7\xd5\x59\xd0\xd1\x59\x2a\xb2\x93\x14\x7c\x07\x2d\x73\x14\x4e\x49\xa0\xa3\x99\x87\xcc\xe9\x27\x0d\x02\x61\x1d\x8b\xd7\xb3\x84\x13\x48\x82\x3e\x35\xe0\xe2\xe0\x27\x53\xe5\xb4\x2c\x8a\x1d\xa6\xf2\x7f\xb2\x05\x93\x94\xfc\x4e\xa9\x73\x2b\xf8\x84\xcd\xb6\x1c\xc8\x7b\xe4\xdf\xa6\x51\xab\xc2\x66\xe8\x3c\x34\xfb\x5d\x31\xa3\x3a\xeb\x9c\xce\x16\x6e\xb3\xaf\x86\xf6\x97\x50\x61\x54\x12\xf2\xbf\x4f\xd6\x6c\x09\x59\x9d\xff\x86\xd1\x38\x46\x9b\xb1\xd9\xae\x6e\x9a\x46\xa9\x55\x80\x8f\x8b\xcf\x21\x85\x50\x2b\x6a\x8f\x58\x11\x3e\xdc\x61\x6f\xe8\xb8\x8f\x93\xf7\x8d\x72\xd6\x90\x5d\xc1\x0a\xc0\x45\x60\xb3\x5d\x80\xe6\x52\x0d\xda\x36\xb3\x8b\xd2\x52\x3c\xe2\x0e\x9b\xed\x19\x68\x1a\x45\xd1\x38\xab\xd4\xdf\xbe\xaf\x46\x7f\x8f\x67\xa7\x38\xd4\xb1\x5c\x23\xf6\x29\x49\x11\xd6\xf9\x1a\x91\x27\x4f\x55\xed\x4f\x00\x00\x00\xff\xff\x24\x9b\xb7\x29\x3c\x03\x00\x00")

func makefileBytes() ([]byte, error) {
	return bindataRead(
		_makefile,
		"Makefile",
	)
}

func makefile() (*asset, error) {
	bytes, err := makefileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Makefile", size: 828, mode: os.FileMode(420), modTime: time.Unix(1488075079, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcDockerfileBuildAlpine = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\x5d\x4f\xdb\x30\x14\x7d\x26\xbf\xe2\x2a\x8b\xd0\x36\xc9\x31\x14\xd6\x50\xa4\x3d\x04\x9a\x75\xd5\x20\x41\x21\xe5\x69\x52\xe4\xd8\xce\x87\xea\xc4\x99\xed\x40\x07\xe2\xbf\x4f\x49\x4a\x4b\x25\xb4\x27\xe7\xde\x73\xee\xb9\xf6\x39\xf9\x11\x47\xb7\xe0\xbc\xcc\xa3\xeb\x5f\x41\x9c\x5e\xad\x96\x37\xf3\x34\xf4\x93\xe5\x43\x90\x2e\x6f\xfd\x45\xf0\x6a\x59\x41\xf8\x00\x8b\xe8\xc6\x0f\x17\xe9\x7d\x7c\x9d\xae\xe2\x1b\x28\x8d\x69\xf5\x25\xc6\x85\x14\xa4\x29\x5c\xa9\x0a\xcc\x04\x2e\xa4\xf3\x12\xaf\xc2\x64\x79\x1b\xa4\x8b\x28\x7d\x08\xe2\xfb\x65\x14\xbe\xba\x5a\x51\xd7\x10\xe5\x16\xcf\x96\xf5\x09\x72\x25\xeb\xbd\x40\x65\xca\x2e\x73\xa9\xac\x31\x93\x74\xcd\x15\x12\x55\xa6\x88\xfa\xbb\x95\xc6\x99\x90\x19\xae\x89\x36\x5c\xe1\x53\xf7\x02\x13\xd1\x56\x0d\xc7\xf3\x81\x9c\x57\x82\x5b\x96\x3f\x9f\xef\xf4\x14\x79\x72\x47\xcd\x4e\x73\x45\x65\x63\x78\x63\xfe\x23\x7f\x7a\x36\xa1\xcc\x3b\xf1\xa6\x17\xfc\x2c\xa3\x93\xe9\x6c\x76\x32\xe1\xe7\xd4\xcb\xbe\x79\xb3\xc9\xc9\x59\x3e\x9d\x32\x3a\xcb\xa7\xe7\xef\x77\x37\x12\xb5\x15\x75\x5b\x62\x68\x09\xd8\xb2\xe2\x55\x08\x9a\x1b\x40\x7c\x03\xbf\xad\xa3\xe3\x63\x20\xed\x1a\x08\x63\x80\x50\x23\x11\x25\xb4\xe4\x80\xd0\x63\xa5\x4c\x47\x04\xb8\x59\x57\x09\x86\x18\x6f\x75\x4f\x3f\xca\x88\x2e\x87\x8f\x82\xd2\xe1\xac\x3b\x2d\x10\xe3\x8f\x43\x21\x5b\xde\x68\x2d\x46\x82\xec\x8f\x71\x05\xdf\xb4\x52\x19\x58\x44\x71\x14\x25\xe9\x55\x14\x25\xf7\x49\xec\xdf\x7d\xb7\x9d\xcf\x85\x04\xde\x3c\x6e\xa1\x2f\xf6\x7e\xe6\xa9\xe8\x6f\xf9\x07\x6c\xe7\x30\x4e\x1b\x50\x04\xdb\x28\xc7\x9c\xb6\xef\x30\x44\x01\xba\x06\xdc\x69\x85\x85\xa4\x44\x00\xda\x3c\xe7\x1f\x52\x55\xfd\x61\x9b\xb2\x77\xd3\xb8\x90\x58\x2b\xba\x85\x46\xff\x50\x3b\x01\x54\xc1\xa1\xa9\x23\xc1\xc5\x35\x59\x73\xf7\xcd\x9f\xdd\x1e\xa4\x72\xc0\x5f\x0f\xa8\xbd\xe1\x8c\x1f\x78\xfb\xf6\xdf\xde\xf9\xc9\x4f\xc0\x85\x1c\xca\xa1\x70\xc6\x26\xce\xaa\xe6\xf2\xf0\x72\x7d\xc7\xe9\xb1\x31\xd4\x7a\xcd\x2a\x05\xa8\x1d\x0c\x1b\x46\xb4\xa2\xf6\xbe\xca\xaa\xc6\x86\xfe\x91\x65\x2d\x19\xa0\x18\x3c\xcf\xdb\xa1\xf6\xbf\x00\x00\x00\xff\xff\xbb\xef\xcb\x99\x5a\x03\x00\x00")

func etcDockerfileBuildAlpineBytes() ([]byte, error) {
	return bindataRead(
		_etcDockerfileBuildAlpine,
		"etc/Dockerfile.build.alpine",
	)
}

func etcDockerfileBuildAlpine() (*asset, error) {
	bytes, err := etcDockerfileBuildAlpineBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/Dockerfile.build.alpine", size: 858, mode: os.FileMode(420), modTime: time.Unix(1488150182, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcDockerfileBuildDebian = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x41\x6b\x83\x30\x1c\xc5\xcf\xf3\x53\xfc\x11\xf1\x16\x73\x19\x13\x7a\x73\x35\x73\x32\x4d\x46\xaa\xee\x32\x08\x56\x69\x2a\x8b\xa6\x44\x85\xd2\xd2\xef\x3e\x16\xb7\x95\x42\x8f\xef\xf7\xf2\xf2\xde\xff\x85\xb3\x1c\xbc\x73\xcc\xd6\x6f\x84\x8b\xe7\x32\xcd\x62\x41\xa3\x22\xad\x88\x48\xf3\x28\x21\x17\xc7\x21\xb4\x82\x84\x65\x11\x4d\x44\xcc\x3e\x68\xc6\xa2\x58\x94\x3c\x83\xfd\x34\x1d\xc6\x15\xc6\x52\xab\x7a\x90\x81\x36\x12\xb7\x0a\x4b\xed\x9d\x79\x49\x8b\x34\x27\x22\x61\xa2\x22\x7c\x93\x32\x7a\x09\x54\x37\xcc\x47\x54\xf7\xed\xd3\x63\x30\xd5\x26\x90\x27\xc7\xe1\x25\x85\x66\x36\x0a\xd0\x6e\xdc\x64\xe0\x7a\x77\x6a\x5c\x40\x1a\x7e\x2b\x96\x1c\x7c\x3a\x0f\xbe\x0f\x53\x6d\x00\xad\x01\xcf\xa3\xc1\x4a\x37\xb5\x02\x74\x3c\xed\xee\x3e\x35\xfd\x2d\xfe\xbb\xe9\x3d\x2a\x5e\x01\x4b\x6d\xa5\x15\xde\x02\xf1\xb6\x1b\x56\xd7\x9f\xb1\xd4\x96\x78\x3f\xde\x32\xbb\xff\x6a\x3b\x03\xe8\x60\x47\xdb\xc8\x68\x1a\xf7\xaa\xb6\xdd\xe0\x82\xef\x43\xb3\xef\x75\x0b\x88\x43\x18\x86\xff\xae\xfb\x1d\x00\x00\xff\xff\x64\xa0\x99\x0d\x76\x01\x00\x00")

func etcDockerfileBuildDebianBytes() ([]byte, error) {
	return bindataRead(
		_etcDockerfileBuildDebian,
		"etc/Dockerfile.build.debian",
	)
}

func etcDockerfileBuildDebian() (*asset, error) {
	bytes, err := etcDockerfileBuildDebianBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/Dockerfile.build.debian", size: 374, mode: os.FileMode(420), modTime: time.Unix(1488075079, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _makeBootstrapMk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\x5d\x6b\xdb\x30\x14\x7d\xae\x7f\xc5\xa1\xcd\x83\x0d\xf6\xc6\x9e\x06\x85\x32\xd2\x46\xcb\xcc\xda\x64\x38\x69\xd8\x9b\x2b\x5b\xd7\xf1\xa5\x8e\x94\x49\xb2\xcb\x60\x3f\x7e\x38\x89\x93\xb4\x1d\x7b\x94\x74\xcf\xc7\x3d\x47\x57\xe8\xa4\x65\x59\x34\xe4\x50\x10\xeb\x35\x58\x97\x4d\xab\x48\xa1\xb2\x66\x03\x5f\x13\x9e\x36\x52\x73\x45\xce\x7f\xd8\x3c\x3f\x05\xf7\xe3\xd9\xf4\x71\x3c\x15\xf8\x72\x13\x64\x8f\xb3\x65\xfa\x20\xf2\xf9\xe2\xfc\x34\x1b\x2f\xd3\x95\xc8\x57\x22\x5b\xa4\xf3\xd9\xf9\xcb\x74\x7e\x7e\x1b\x5c\xc1\x6c\x3d\x1b\x2d\x9b\x93\x8b\x60\x92\xa5\x2b\x91\xe5\x13\xb1\xca\x7f\x64\xe2\x6b\xfa\x13\xd7\x37\x50\xd4\x0d\x0f\x27\x02\x8c\xc2\x77\xc3\x51\x32\x0a\x5d\x4d\x4d\x83\x35\x7b\x58\xea\x92\xad\xb4\x8e\xf0\x4d\x8c\x27\xf8\x83\xb2\xf5\x48\xca\x4f\xc9\xe7\x28\x08\x26\xf3\xbb\xef\x22\xcb\xd3\x87\xfd\x32\x28\x8a\xa6\x72\xf5\xc7\x51\x38\x6c\x18\x25\xca\x72\x47\xf6\xd5\xe4\xa0\x2f\x26\x7b\x07\xa5\x6c\x1a\x90\x2b\xe5\x96\x72\x65\xca\x67\xb2\xb9\x97\xeb\x78\x14\x9e\x83\xa2\xeb\xa3\xd5\x03\x3c\x8a\x06\xd6\xdb\xc7\xf4\x7e\x32\x64\x76\x34\xf3\x06\x9f\x14\x2d\x37\xea\x35\xe4\x40\xf8\x7f\x48\xf2\xc2\xbe\x4e\xd6\xa6\x0f\x5b\x51\xc5\x9a\x14\x0a\xaa\x65\xc7\xa6\xb5\xa8\x8c\xc5\x6e\xcc\x81\xb5\x63\x45\xf0\x56\x76\xec\x92\x92\x03\xae\x34\xfd\x42\x38\x0a\x8d\xe5\x35\x6b\xdc\xa5\x51\x8c\x56\x1f\x48\xa2\x00\x00\xae\xc0\x15\x5e\x08\xd2\xd2\x40\x70\x97\xc6\xe8\xc8\x16\xc6\x11\xd8\x81\x74\xdf\xaa\x42\xf1\xbb\x97\x97\x6d\xe3\x83\x8b\x95\xc8\x6e\xe7\x0b\xd1\xf7\xea\x6d\x4b\x01\x69\xc5\x55\x6f\x90\x2b\x2c\xb3\xf1\x2a\x5d\xe4\xcb\xf1\xf4\x68\xf7\x4d\xef\xec\x60\x3a\xb2\x96\x15\xa9\x93\xc9\x13\x2e\x8a\xb1\x37\xf7\x06\x77\xdd\x27\x74\x36\xf6\x4a\xf6\xb0\x83\x36\x1e\xac\xd1\x17\xb8\xfb\xf9\xdb\xd6\xd5\xbd\xa0\x62\xb7\x5b\x23\xe0\x6a\x2f\x57\xb1\x75\xfe\xc5\x58\x85\x51\xe8\xda\xc2\x79\x24\x31\xe2\xf7\x25\x47\xf1\x3f\x3f\x69\x14\x5c\xf4\xd4\x03\x2d\x6e\x70\xb9\x93\x3a\x5e\xf4\xc5\x28\xea\xa8\x31\xdb\x0d\x69\xdf\x27\xea\xd8\x68\x77\xb9\x77\xfd\x37\x00\x00\xff\xff\x2b\x2b\x66\x5c\xb6\x03\x00\x00")

func makeBootstrapMkBytes() ([]byte, error) {
	return bindataRead(
		_makeBootstrapMk,
		"make/bootstrap.mk",
	)
}

func makeBootstrapMk() (*asset, error) {
	bytes, err := makeBootstrapMkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "make/bootstrap.mk", size: 950, mode: os.FileMode(420), modTime: time.Unix(1488107348, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _makeFunctionsMk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\x4d\x0e\x82\x30\x10\xc5\xf1\x7d\x4f\xf1\x12\x58\x40\xd2\x2e\xdc\x72\x19\x32\xb4\x83\x12\xcb\x8c\xe9\x47\xbc\xbe\x51\x51\x49\x70\x3b\x79\xf9\xff\xa6\x01\x67\x4f\x37\x1e\x83\xfa\x2b\xa7\xb1\xd0\x79\xbb\xc0\x6b\x54\x81\xbf\x50\x42\x51\x50\x8c\x7a\x47\xcd\x0c\xc2\x7b\x8b\xe7\x96\x32\x52\x8d\x6c\x02\xcf\x8b\xf0\x31\x66\xda\x2e\xd7\x29\x17\x0c\xd6\x39\xdb\x76\xa7\xbe\x37\x2c\x81\x67\x63\x1a\x54\x39\xe2\x24\x5b\x24\xec\x99\xa2\x98\xf8\xa5\x2f\xf2\x7b\xc0\xeb\xba\x92\x84\x0f\xfe\x27\xf7\xe5\x9d\xb3\xc3\x9e\x7f\x04\x00\x00\xff\xff\xff\x00\xb3\xb9\xf8\x00\x00\x00")

func makeFunctionsMkBytes() ([]byte, error) {
	return bindataRead(
		_makeFunctionsMk,
		"make/functions.mk",
	)
}

func makeFunctionsMk() (*asset, error) {
	bytes, err := makeFunctionsMkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "make/functions.mk", size: 248, mode: os.FileMode(420), modTime: time.Unix(1488075079, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _makeRulesMk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x5b\x6f\xdb\x36\x14\x7e\x16\x7f\xc5\x81\x2b\xa3\x22\x1a\xd9\x2b\xfa\xe6\xcc\x6b\xdc\x5a\x75\x8c\xf9\x12\xf8\x92\x6d\x68\x06\x81\x96\x28\x9b\x30\x25\x39\x22\xe5\xa6\xd8\xf6\xdf\x07\x8a\xba\x26\x8e\xb6\xbd\xec\x29\xca\xf1\x77\x2e\x3c\xe7\x3b\x1f\xf9\x69\x3b\x9d\x8d\xdd\xe9\x18\x06\x43\x30\x2d\x71\xa0\x9c\x83\x4f\x24\x85\x77\x9d\x6e\x68\x77\x7d\xbb\xfb\x9b\xdb\xbd\x75\xbb\x73\xb7\xbb\xee\x60\xa4\xe1\x77\xa3\xcd\xad\x76\xe0\xb1\x47\x24\x8b\x23\xdc\xdf\xa5\x8c\xfb\x08\xbd\x01\x3f\xf6\x8e\x34\x81\x24\x8d\x24\x0b\x29\x78\x71\x18\x92\xc8\x17\x68\xbc\xfc\xfc\xb3\xb3\x72\x3f\xcf\xc7\xf0\x71\x98\xa3\x0a\x63\x16\x56\x99\x4d\xab\x82\x61\xd0\x31\x73\xcb\x6a\xbb\x78\x89\x48\xd2\x08\x6c\x3b\x09\xc1\x96\x05\x6e\x33\x9a\xbc\xc4\x49\xb2\x2f\x7e\xbf\xdb\xae\x6f\x5f\x02\x4e\xa9\x38\xa0\xfc\x78\xf7\xcb\xd9\x76\xee\xb8\x9b\xd1\x6a\xe2\x6c\x14\xb4\x1f\x9f\x64\xdf\x4f\xd8\x99\x26\x7d\x91\x78\xfd\x26\x2e\xeb\xc6\xc7\x46\x37\x50\x91\xec\xcb\x74\xe6\xb8\x65\xa6\xe9\x7c\x34\x71\xdc\x7b\x67\xb5\x9e\x2e\x17\xce\x18\x3f\xf3\xea\x8f\xb3\xa6\x04\x8c\xd3\xcb\xfe\x3a\xed\x78\x35\xbd\x2f\x82\xe5\x21\x84\x7f\xac\xa2\x50\xe9\xd5\x22\xf5\xb2\x26\xf6\x4c\x6b\xb5\x5d\x6c\xa6\x73\xc7\x5d\xae\x71\x5b\xf4\xc5\x68\x33\xbd\x77\x1a\xd1\x2f\x15\xd8\x2b\xe7\xcd\x99\x90\x10\x07\xc0\x42\xb2\xa7\x02\x64\x9c\x8f\x2d\x67\x96\x8a\x33\x6c\xcd\xd1\x7a\xbc\x57\x7b\xa7\x52\xef\x63\x4e\xa2\xfd\x4b\xaa\x4d\x96\x19\xcd\x86\xb0\x8f\xd5\xf7\xc6\x59\x6f\x40\x1d\x44\xdb\x31\x48\x2a\x24\xd8\x67\xf5\xdb\x6c\xfc\x65\x36\x9a\xac\x61\x08\xf6\xaf\x10\x12\x16\xf5\xce\x34\x11\x2c\x8e\x54\xcd\xba\x92\x3c\x27\x2e\x11\xd9\xf9\x86\xa6\x55\xec\x4e\x56\x4b\x66\x04\x1a\x9d\x59\x12\x87\x34\x92\x70\x26\x09\x23\x3b\x4e\x45\xde\x89\xed\xda\x59\xa9\x7e\xee\x76\x3c\x10\x87\xc2\x38\xcd\x69\xaf\x37\x8f\xf9\x60\xa7\x60\x5a\x0a\x8b\x8b\x6d\x1b\xad\x26\x6b\xf8\x38\x44\x8d\xde\xe5\x6b\x54\xb6\x67\xb5\x5d\x60\x78\x40\x46\xe6\x5e\x25\xc4\x83\xf2\xbf\xe9\x58\x03\xce\x25\xa0\xc6\xdf\x0a\xd7\x20\x7f\xe6\xd1\x36\xbc\x62\x65\xf2\x56\xfd\x8f\x55\x65\x1e\x93\xa5\x86\xf5\xf7\xf1\x85\x4a\x1b\x4c\x52\x33\x62\x01\xdc\x3b\xab\x4f\xcb\xb5\x03\x4c\x40\x1a\x09\x2a\x0b\xb9\xd2\xe3\x63\x02\xe8\x13\xf5\x52\x49\x7d\x60\x11\x3c\xa6\x4c\x52\x08\x63\x9f\x22\x16\xd0\x47\xb0\x4c\x2b\x4e\xd8\x9e\x45\x45\x18\x7c\x05\x69\xe4\xd3\x80\x45\xd4\xc7\xc8\xa8\xcd\xeb\xdd\x10\xec\x47\x44\x23\x9f\x05\x08\x8d\x66\xb3\xe5\x2f\xce\xd8\x9d\x2e\x5c\x5d\xa0\x5a\x3c\x18\xaa\x92\x67\xa3\xc5\x64\x3b\x9a\x38\xea\xbb\xd8\xd0\xbc\xc3\x39\xef\xa0\x30\x4f\x96\xa5\xe9\xa1\xc8\x95\x91\xaa\xa2\x52\x4d\x9e\x1f\x90\x51\xdf\x1c\xb8\xbc\x46\xf0\xea\x68\x55\xbf\x24\x4d\x42\xf0\x62\x1e\x27\x02\x0e\x94\x9f\x68\x22\x90\xe7\x25\xd4\x1f\xc2\x03\xfd\xfa\xc3\xf5\x87\xf7\x21\xf2\xbc\xef\x94\xf3\xf8\x5b\x61\xfa\x10\x66\x10\x41\xa5\xb6\x84\x2a\xd0\x37\x0a\xf4\xe9\x14\x27\x12\xe4\x81\x96\x9b\xa1\xb4\x82\x28\x5f\xb5\x36\x22\xdd\x09\x79\x05\xd9\x1f\x26\x53\x49\x0b\xa8\x50\x93\x90\x07\x8a\xde\x40\x25\x3f\x02\xe9\x78\x08\x11\xce\x07\xb9\xe2\xa0\x82\x2d\x9a\x14\xc8\x08\x8f\x3e\x4b\xc0\x3e\x41\xe3\x87\x0a\xa7\x99\x31\x40\xc6\x8d\x69\xd1\x33\xe1\xb0\x71\xe6\x77\x99\x2a\xd6\xaf\xc3\xf0\x28\x69\x78\xd2\x61\x64\x78\x2a\xd5\x10\x63\x64\xdc\x64\x6e\x9d\xa2\x7e\x78\x6b\x5a\x41\x9c\x50\xe2\x1d\xe0\x7c\x65\x5a\x17\x07\x8f\xaf\x1e\x4c\xf3\x0f\xd3\x3a\xe3\xbf\xf0\x5b\xf8\xb1\xda\x16\x2d\xc7\x37\x18\x7e\x02\xd3\x2a\x4b\xc1\x9d\xe7\xc4\xc6\xe5\x79\x14\xd5\x30\xd8\x12\x4c\xcb\x23\x9c\x43\x1a\x51\xe1\x91\x13\x75\x35\xab\x5d\x49\xf6\x57\x2a\xa0\x1d\x34\x22\x42\x0f\x21\xa5\x81\x03\xf8\x13\xce\x84\xb3\xec\xc2\x57\x06\x3b\x22\x92\x9d\xf3\x6f\x7d\xe1\xa1\x9a\x7d\xf0\x5c\xab\x9b\x5a\x60\x14\x65\x55\x1a\x85\x21\x24\xc7\x46\x6c\x9b\x45\x92\x26\x11\xe1\xba\x84\x3c\x8b\xaa\xa4\x31\xa4\xd6\x4c\xad\x57\x46\x55\x46\x25\x4a\xf5\x32\x74\xc2\xcb\x65\x94\xd6\x9c\x52\x2f\x6a\x36\x3c\x1f\x34\xf4\x5a\x0b\x4e\x7e\xbd\x60\xe8\xf5\x7b\xbd\x1e\x42\x99\x9b\x3a\x4d\xdd\x3f\xff\x47\x3b\xb6\x5c\x69\x75\x9f\xc1\x7f\xe8\x47\x4b\xe7\x2f\x1f\x03\xd5\x0b\xfa\xc7\xde\x37\x6f\xe4\x7a\xc4\x96\x5e\xd7\x33\xbc\x92\xb8\xea\xf6\xc5\xbe\x56\x4f\x40\xb0\x6d\xee\x07\x9c\xec\x85\x5a\xb0\xea\xd6\xc6\x6f\xc1\x8e\x9b\xb5\xe7\xcf\x34\xe8\xa9\x38\x48\xbd\xea\x0a\x75\x30\x6e\x58\x00\x5f\xc1\xb4\x94\xd1\x67\x42\x29\x90\x8f\xe1\xf7\x6b\x25\x2e\x91\xca\x6a\x50\xef\x10\x83\x4d\xa1\x63\x5a\x99\xcc\xe1\x67\x60\x6d\x16\x54\xe2\xce\xb5\x76\x78\x62\x12\xde\x67\xdf\x01\x43\xb5\x3d\xdd\x8c\x26\xb8\x75\x29\x5f\xe3\x40\x76\xb7\x19\xff\xda\x13\xe3\x01\x27\x8a\xc0\xb5\xdc\xea\x81\xdb\xf6\x70\x6a\x47\x96\x01\x51\xa1\x0c\x03\xe5\xa1\x5f\x2d\xb6\xf0\x8f\x18\x76\x71\x2c\x85\x4c\xc8\x09\x6c\xdb\x4f\xbe\xdb\x49\x1a\x21\xe4\x71\x4a\xa2\x01\x32\xd4\x5b\x3c\x09\x9a\x43\xf9\x3b\x00\x00\xff\xff\x23\x5f\x98\xc3\x65\x0c\x00\x00")

func makeRulesMkBytes() ([]byte, error) {
	return bindataRead(
		_makeRulesMk,
		"make/rules.mk",
	)
}

func makeRulesMk() (*asset, error) {
	bytes, err := makeRulesMkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "make/rules.mk", size: 3173, mode: os.FileMode(420), modTime: time.Unix(1488155413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"Makefile": makefile,
	"etc/Dockerfile.build.alpine": etcDockerfileBuildAlpine,
	"etc/Dockerfile.build.debian": etcDockerfileBuildDebian,
	"make/bootstrap.mk": makeBootstrapMk,
	"make/functions.mk": makeFunctionsMk,
	"make/rules.mk": makeRulesMk,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"Makefile": &bintree{makefile, map[string]*bintree{}},
	"etc": &bintree{nil, map[string]*bintree{
		"Dockerfile.build.alpine": &bintree{etcDockerfileBuildAlpine, map[string]*bintree{}},
		"Dockerfile.build.debian": &bintree{etcDockerfileBuildDebian, map[string]*bintree{}},
	}},
	"make": &bintree{nil, map[string]*bintree{
		"bootstrap.mk": &bintree{makeBootstrapMk, map[string]*bintree{}},
		"functions.mk": &bintree{makeFunctionsMk, map[string]*bintree{}},
		"rules.mk": &bintree{makeRulesMk, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

