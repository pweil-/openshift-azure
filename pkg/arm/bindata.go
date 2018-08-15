// Code generated by go-bindata. DO NOT EDIT. @generated
// sources:
// data/azuredeploy.json
// data/master-startup.sh
// data/node-startup.sh
package arm

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

var _azuredeployJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x1b\xdb\x6e\xdb\x38\xf6\xbd\x5f\x41\x68\x0b\xa4\x2d\xe2\x5b\x92\x76\x17\x01\xf6\x21\x97\x76\x62\x34\x69\x8c\x28\x9d\x7d\x18\x04\x05\x4d\x1e\xdb\x9c\x48\xa4\x96\xa4\x9c\xba\x5e\xff\xfb\x82\xba\xd8\x92\x4d\x99\x92\xed\x41\x3b\x9c\x87\x78\xa4\x73\x0e\x79\xee\x17\xaa\xf3\x57\x08\x21\xe4\xbd\x56\x64\x02\x21\xf6\xce\x91\x37\xd1\x3a\x52\xe7\x9d\x4e\xfa\xa4\x1d\x62\x8e\xc7\x10\x02\xd7\x6d\xfc\x23\x96\xd0\x26\x22\xcc\xde\xa9\xce\x49\xb7\xf7\xbe\xd5\xed\xb5\xba\xbd\x0e\x85\x28\x10\x33\x03\xf7\x08\x61\x14\x60\x0d\xed\x3f\x95\xe0\xff\xf0\x8e\xd3\x1d\x88\xe0\x1a\xb8\xfe\x1d\xa4\x62\x82\x9b\x8d\x7a\xed\xae\xf9\x2f\x07\x88\xb0\xc4\x21\x68\x90\xca\x3b\x47\xf3\x45\xf6\x74\x8a\x25\xc3\xc3\x00\x4a\x0f\x25\x28\x11\x4b\x92\x3c\xfc\xe3\xd5\x7c\xde\x42\x6c\x84\xb8\xd0\xe8\x0d\xe3\x14\xbe\xa3\xf6\x95\xe0\x1a\x33\x0e\xd2\x07\x39\x65\x04\xda\x03\x29\x22\x90\x9a\x81\x6a\x5f\x8c\x81\xeb\x81\x10\xc1\x40\x8a\x11\x0b\x40\xa1\xee\xdb\xf6\xef\x1c\xb4\x1f\x0f\x39\xe8\xfe\x35\x5a\x2c\x92\x6d\xcc\x9a\x2f\x7f\x25\x1b\xeb\x59\x04\xe6\xe8\x77\x8c\x48\xa1\xc4\x48\xb7\xbf\x80\x7e\x11\xf2\xb9\x33\x65\x52\xc7\x38\xc8\xfe\x57\x65\x4c\x2d\x11\x71\xc4\x0a\x9c\x9f\x74\x7b\x1f\x5a\xdd\xd3\xd6\x69\x77\x1d\x2e\x10\x04\xeb\x0c\x6a\x3e\xb7\x30\x72\x9b\x01\xa0\xff\xa1\x3f\x15\x5a\x2c\xd6\x09\x70\x1c\x26\x27\x9c\x72\xd0\xeb\xef\xa2\xa5\x10\x8c\x30\x4b\xef\xd2\x43\x52\x2a\x41\x29\x3f\xc2\x04\xac\x10\x45\xa8\x81\x84\x11\xfb\x9e\xa9\xc0\x06\x98\x00\xf7\xba\xa9\x92\x3b\xff\xf2\xac\x40\x4f\x1b\x4f\x17\xc7\x9b\x07\x53\x89\x66\xaa\xb7\xb2\x9f\x14\x15\xe5\x41\x61\x84\xe3\x60\x5d\x24\x25\x50\x87\x78\x4a\xb0\x25\x21\x24\xc6\x9c\x33\x7a\x72\x66\xe7\x34\xe1\xcd\xfa\x66\xf3\x69\x59\x2a\xab\xf7\x8b\xe3\xc4\xd8\x81\x53\x63\xa3\xf3\x79\xe7\x1d\x7a\xbc\xbf\xbe\x3f\x47\x2f\x80\x42\x3c\x43\x43\x40\xc6\x57\x90\x16\x08\x4f\x05\xa3\x48\x4f\x00\xf5\x07\x08\x73\x8a\x6e\x2f\xd1\x04\x24\x18\x4f\x79\x01\xa4\xb4\x88\x50\xac\x18\x1f\xa3\x77\x9d\x9c\xd6\x35\x86\x50\x70\x1f\xb4\x42\x23\x21\x91\x7f\xfd\xa5\x8d\xd0\xe3\x04\x38\xba\xc3\x4a\x83\x44\x5c\x50\x50\x48\x4d\x44\x1c\x50\x44\x44\x08\x28\x8e\x10\x56\xe8\x01\x30\x9d\x15\x08\xe1\x58\x8b\x10\x6b\x46\x70\x10\xcc\x8e\x93\xed\x27\xc0\x09\x98\x03\x02\xd7\x20\x81\x22\xc6\xb5\x40\x44\x70\xc5\x28\xc8\xd4\xa4\xcd\xa6\xb8\x40\xe6\x56\x60\x7a\x89\x03\xcc\x09\x48\x94\x39\x00\x1a\x49\xc1\xb5\x39\xb7\xe1\xed\x62\xd0\x47\x0a\xe4\x14\x64\x8a\x96\x0b\xaa\xb6\xdb\x46\xf1\x30\x60\xa4\x3f\xb8\x48\xf5\x09\x3f\xdf\x71\x59\xd4\xc2\x11\x4b\x99\x6a\xea\xc0\x4b\x6e\x82\xfc\x38\x77\xa0\x27\x82\x1a\xba\xd7\x33\x8e\x43\x46\x2c\x0e\xe0\x31\x1a\xc0\x23\x0b\x41\xc4\xba\xcf\xef\x18\x8f\x75\xb2\x41\xef\xbd\x05\x96\x72\xe5\x83\x36\x0a\xa8\xf6\x12\x8f\x8a\x10\x33\xfe\x05\x87\x70\x8b\x87\x10\x14\x64\x32\x62\xe3\x76\x6a\x4c\xb7\x97\x57\x06\x20\xf5\x20\x23\x8f\xcd\x48\x50\xf6\x82\x35\x59\xa8\xe7\xd8\x2e\x84\x5c\x90\x97\x58\x31\xe2\x55\x7b\x52\xfe\xb3\xb6\xad\x04\x05\x73\xdc\xd9\x4e\x28\x44\xc0\xa9\xba\xe7\xd6\x60\xe6\xfd\x91\xa7\xb6\x3e\x7d\x73\x54\xc3\x5c\x8f\x8e\xd1\x51\xd1\x62\x8e\xde\x3e\x95\x59\x7e\xfa\xab\xec\x34\x18\xee\x6e\xa7\x43\x4c\x9e\x81\xd3\x8c\x0b\x93\x8b\xf7\x8a\xed\x19\x39\x7b\xdc\xb5\x44\x57\x8b\x59\x27\x51\x05\x38\xed\x0f\x52\x23\x8d\xd3\x90\xb4\xd7\xb1\x72\x9a\x87\xca\x39\x91\x64\x53\xac\xa1\xa9\x7b\x97\x69\x94\x0d\xc8\xb9\x29\x4a\xe3\x83\xd9\xe1\xe0\xa6\xb9\xbe\xec\xd9\xb1\xfa\x4d\x3d\xcd\x32\x3e\x14\x31\xa7\x5f\xb0\x5e\xda\xd9\x76\xb0\x87\x38\xad\x36\xad\x60\xab\x18\xc0\xf8\x78\x09\xb9\xab\x85\x44\x42\xea\xd6\xd9\xd9\xe9\xa1\x2c\x64\xd3\xaf\x1a\x29\x98\x08\x4e\xb0\x7e\xb3\x5d\xcf\xa5\x28\x68\x74\x5c\x0c\x04\x47\x6f\x8f\xd1\x51\xc7\xe2\xde\xf9\x33\xb7\x11\x38\x0c\x38\xa3\x33\x10\x52\x7b\xe7\xe8\xec\xec\xd4\x01\x0f\xdc\x54\x44\x9f\x02\x81\x4d\xce\xea\x0f\xbc\x73\x34\xc2\x81\x02\x07\x5a\x45\x3c\xf8\x29\xe2\xac\x8a\x4d\xcb\x17\x7b\x0b\x35\x27\x54\x5b\xaa\x0d\x2a\x86\x12\x9e\xe1\xf6\x9a\x29\x2d\xd9\x30\xce\xb3\xd0\xb5\xb3\x32\x47\x99\x1f\x0c\xab\xbb\x92\xb5\xd3\x1d\x54\xfe\xc9\xce\xaa\x93\x3b\xeb\xde\xd2\x8e\xa4\xd0\x82\x24\xbe\xe9\x3d\x92\xe8\x00\xed\x82\x25\x50\x89\x58\xd7\x0a\x68\x29\x73\xbf\x52\x10\x63\xa6\x41\x98\xe2\xa0\xcf\x7d\x20\x82\x53\x83\xe2\xb2\x2b\x1e\x87\x43\x90\xf7\xa3\x41\xce\xcd\x89\x4b\x07\x75\x2d\xfd\xf0\xca\xfa\x45\xaa\x5a\x5f\x0b\x89\xc7\xd0\x51\xe9\xdf\x0b\x42\x44\xcc\xb5\xbb\xae\x7d\xdf\xea\x7e\x68\xf5\xde\xff\x65\xfd\x4f\xa1\x51\x78\x80\xb1\x89\x15\x33\xbf\x74\xc4\x2a\x02\xce\xe9\x46\x8a\xfe\x98\x49\xc2\xd7\x98\x53\x2c\xe9\xb7\xdb\x07\xbf\x52\x9e\x49\xbb\x2d\x31\x1f\x03\x7a\x8d\xa3\x08\x9d\xff\xbb\xe9\x6c\x69\xb1\x4b\xa7\xc1\xd3\xbf\x3e\x90\x58\x32\x3d\xfb\x4d\x8a\x38\x3a\x54\x67\xfa\x7a\x8b\x6a\xaa\xb5\xc2\xd5\xb8\x65\x90\x71\x14\xb5\x4d\xdb\xb6\x83\xfc\x55\xc6\xcd\xde\x55\x13\x0e\x02\xf1\xf2\x4d\xa9\xc9\xc1\x86\x39\x84\xa4\xb5\xb0\x67\x2a\xeb\x17\x57\x26\xa2\xa0\x88\x64\x51\x2e\xd3\x04\x07\xf9\xfe\x0d\xd2\x12\x8f\x46\xee\x1a\x9c\x82\xd2\x8c\x27\x12\xbf\x58\x1f\x23\xbd\x6b\x80\x6c\xb2\xf5\x83\xb1\xcd\x44\xff\x27\xad\x93\x13\x27\x32\x93\x40\xf2\x73\xf7\xd3\x8a\xd7\x9d\x77\x99\x30\x6a\x33\xe9\xbd\xdb\x6b\x18\x2d\x1d\xe0\x69\x62\x6e\x2e\x84\x14\xaf\xc4\xff\xbb\xe6\x91\x39\x9b\x1b\xc3\x7f\x53\xb3\x7e\x10\x01\x20\x2f\x4c\x86\x13\x5e\xc9\x6f\x8b\xab\xb6\x7d\x26\x63\xf4\x5f\xc9\x42\x6f\x1e\x1f\x07\xfe\x4f\xb5\xd1\xb3\xb3\x53\x47\xa1\x80\x0e\x62\xa5\xce\xe4\xff\x77\xb3\xd2\x6c\xe0\xbb\xfe\x72\xcb\x98\x38\xff\xe9\x4c\x39\x57\x22\x8c\x62\x0d\xf9\xfd\xc5\x1d\x26\x13\xc6\xc1\x27\x38\x00\x1f\x6a\x94\x03\xff\x6c\xf5\x4e\x5a\xdd\xde\x1e\x49\xa7\x22\x9d\x97\x07\x65\x85\x3b\x9e\xc4\x59\xab\xee\x6b\x96\xe8\x8e\x69\xc5\xda\x75\x8d\x29\xfc\xa7\x1c\xb4\xa9\xed\xd7\x66\xec\x5b\x63\x44\xe3\x7d\x1d\xed\xc6\xda\xee\x4d\xa9\x5b\xeb\x06\xb3\x8b\x25\x7d\x3b\x27\x85\x79\x38\x53\xca\x95\xf9\x2b\x8b\x56\xcd\x40\x16\x6b\x2d\xdb\xf0\x99\xe0\x08\x93\xd4\x73\xbd\x7c\x9f\xab\x2d\x75\x1e\x5a\x2b\x16\x53\x73\xb8\xf3\xd9\x0f\x58\x62\xac\x97\xd8\x05\x2d\xe6\xb5\x65\x3f\xc4\x63\x78\xc8\xa4\x99\xf0\xe5\x6d\x28\xd4\x8b\x02\x6c\x6f\xf9\x4b\x07\x28\x91\xf4\x3f\x7f\xdd\x76\xf0\x64\x4a\xa6\x26\xa9\x58\x36\x90\x07\xf9\xdb\xad\x24\xa4\xa0\x31\xd1\x56\x02\xf7\xa3\x51\x01\xd9\x26\x06\x8b\x65\x39\x2b\x37\x31\x05\x19\x49\x31\x65\x99\xdf\x57\x0c\x50\xbc\x38\x1a\x4b\x4c\x61\x20\x02\x46\x66\xd5\xb7\x03\xa1\xa0\x69\x10\xc2\x3c\xc6\x81\x65\xee\x6f\x21\x5d\x0e\x4f\x59\x75\x5d\xbd\x85\x50\x2e\x10\x94\xde\xe2\x85\x8c\x7f\x55\x20\x73\x75\x92\x40\xc4\xb4\x15\xab\x8d\x91\x76\x09\x8d\xa4\x31\x53\xae\xee\x2f\x8a\xb6\x98\x39\x49\x6b\x1b\x85\x80\xf1\xf8\x7b\xb3\xc9\x92\x47\x99\xc2\xc3\x00\x06\x58\xa9\x17\x21\xe9\x45\xac\x27\xc0\x35\x5b\xc6\x59\x2d\x63\xd7\x54\xcb\xd4\xcc\xb5\x26\x28\xe9\x30\xf7\x33\xcc\xb6\x5f\xef\x16\x97\x9b\xea\x92\xfa\x33\xcc\xae\xb1\xc6\x99\xd0\x7c\xff\x66\x90\x6f\x77\xa1\x7c\x2d\x19\x1f\xaf\xcc\xda\xf7\x6f\x3e\xc3\xac\xbd\x84\xb0\x3b\x45\x35\x23\x58\x1b\x96\xbd\xce\x44\x84\xd0\x59\xa9\xb7\xd3\x56\x6a\xd2\xc1\xb1\x9e\x08\xc9\x7e\x00\xfd\xf6\x6c\x78\xad\x45\xb7\x7a\x4c\x9d\xaf\xcd\x4b\xed\x7a\xf8\x15\x25\x80\x9d\x5d\x2f\x6b\xdc\x6b\x19\x3a\x4b\xa3\xdd\x08\x24\xf0\xec\x6e\x7f\xc7\x98\xb8\x41\x5a\x98\x90\x53\x23\x18\xd5\xb9\x98\xd8\x2b\x30\x96\x85\x93\xa4\xa4\xa6\xe1\xb9\x44\x62\xba\xaa\x73\x36\xc8\x64\x35\xd0\x2a\xd2\x26\xa1\x35\x50\xe0\x14\x97\xe5\x2e\x65\x83\x7a\xae\x82\x24\x7f\xe7\x7b\x98\x34\xbe\x59\xb6\x25\x8a\x4d\x52\x7c\x25\x95\x44\x91\x39\x11\x93\xf4\xb7\x14\x18\xf9\xda\x32\xc2\xf4\x84\xba\x66\xea\xd9\x1d\xaf\x48\x12\xab\xc7\x86\xdd\x07\xc0\xf4\x3f\x92\x69\x70\xc9\x9c\x48\xc0\x1a\xee\x97\x9d\xcb\x27\x29\xc2\x84\x19\x17\x62\xfa\x95\x12\xad\x75\x32\x54\xf0\x9e\x8b\xf2\x50\x68\x20\x21\x64\x71\xb8\x39\x13\x5a\x5f\xdb\x9c\x78\xa7\xbe\x32\x39\x14\xc5\x1a\x1b\x16\xdc\x51\xb7\x06\x87\x94\xa9\x67\x53\x15\xfd\x76\xe9\x9d\xa3\x53\x47\x4f\x84\x6c\xd2\xff\x18\x46\x7a\x56\x23\xda\x7a\x41\x6c\xe0\xbb\x3b\x4a\xec\xc9\x65\x91\x55\x11\x30\x2b\x77\x6b\x45\xc0\x0c\xb6\xcf\x35\xc8\x11\x26\x50\xf3\x9a\x37\x5f\x35\xe4\xbd\x1c\x95\x39\xbb\x6b\xd4\xb0\xfd\x2f\xe0\xb0\x10\xcb\x59\xad\x64\xbf\x44\x4a\x6f\xc0\xfa\x83\x4f\x42\xbe\x60\x49\x53\x97\x6c\x80\x6f\x6b\x2a\x6a\x1f\x19\xd5\xbe\x3f\xde\xab\x77\xa9\x5a\x8e\xab\x98\xd5\x09\xa3\x86\xf6\x50\x5c\xf5\x25\x81\x50\xe9\x43\x1f\x92\xec\xd9\xa0\x9c\x41\x3b\x1a\xce\x1a\x7e\x73\x23\x2a\x11\x48\xbf\xc6\xdb\x69\x73\xb4\x32\x87\xf9\x9c\x8d\x36\x7b\xf9\xc5\x62\x3e\xb7\x3e\x34\xc9\x75\xb1\xa8\x77\xa7\x57\xd9\xdc\x1f\xa3\xa3\x4e\xf6\x2d\x61\x27\xfb\x20\xf0\xe8\xed\xd3\x7c\x0e\x9c\xda\x3e\x40\x72\xad\x9a\xd6\x55\x62\x3e\xfb\x4c\x22\xca\xc6\x49\xcd\x6f\x97\xad\x54\x57\x36\xd5\xd0\x9a\x56\xe7\xda\xcf\xaa\x96\x74\x2a\x6f\x86\x77\x22\xe9\xae\xb5\x37\x31\x76\x4e\xbf\x55\xcb\x2b\x0e\x6d\x2e\x1b\x7c\xba\xe4\x5a\x7b\x89\xf9\x97\xf9\x62\xa3\x6a\x35\xd7\x9d\xb3\x10\xd8\x7f\xab\x7a\x90\xdb\x1b\x38\x37\x9d\x6d\xb5\x8e\x15\xbe\xa2\xc0\x81\xef\x1a\xb8\x69\x34\x6a\x95\x38\x4b\xe8\x83\x96\x33\x44\xb9\x8a\x6f\xb4\x6b\x39\x83\x63\x2d\xbe\xa6\x33\xa3\x3b\xc6\x85\x5c\x4d\x96\x1b\x94\x27\x91\x14\x1a\x88\x06\xea\xfc\x2c\xd5\x8a\x9e\x5e\x94\x64\x4d\xde\x25\x56\xf0\xe1\xec\x23\x27\x82\x02\x7a\xe3\x6b\x2c\x75\x1c\x15\xc2\xc8\x5b\xeb\x77\xaa\xb6\x55\xb7\xf0\x28\xb5\xbd\x2b\xd7\xbd\x48\xfe\xa5\xc5\xc7\x95\x42\x6b\x92\x53\x05\x19\xd4\x3d\x42\x7e\x27\x70\x15\x2b\x2d\x42\x3f\x95\x47\x03\xdc\x1b\xcc\x69\x00\xb2\x78\x2d\xd0\xee\xba\xa5\x74\x60\x0f\xda\x9c\x22\x56\x5d\x8e\xac\x07\x99\x6c\xee\xed\x89\x58\x47\xb1\x4e\x45\xf7\x6a\xf1\xea\xff\x01\x00\x00\xff\xff\x1c\xd2\x91\xd7\x19\x33\x00\x00")

func azuredeployJsonBytes() ([]byte, error) {
	return bindataRead(
		_azuredeployJson,
		"azuredeploy.json",
	)
}

func azuredeployJson() (*asset, error) {
	bytes, err := azuredeployJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "azuredeploy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\xfd\x73\x1a\x39\x96\x3f\x6f\xff\x15\x3a\x3c\x7b\x99\x49\xa5\x01\x3b\x4e\x26\xe9\x89\xb7\x8a\x60\x9c\x50\x21\x36\x0b\x24\x73\x73\x7b\x53\x2e\xd1\x7a\x80\xc6\xdd\x52\x47\x52\x13\x33\x49\xfe\xf7\x2b\x7d\x74\xa3\x86\xc6\xc6\x76\x76\xaf\xee\xea\x32\x55\x1e\x90\xde\x97\xf4\x9e\x9e\xde\x47\x37\x07\xff\xd6\x9a\x52\xd6\x9a\x62\xb9\x40\x21\x5c\x07\xc1\x01\x9a\x5c\x9c\x5e\x44\xa8\x05\x2a\x6e\x11\x26\x53\x2c\x3f\x35\x49\x8b\x0b\x3a\xa7\x2c\xcc\x33\xa9\x04\xe0\x34\x24\x4c\x36\x63\xce\x66\x88\x4a\x14\xe7\x42\x00\x53\xc9\x0a\x2d\xb0\x20\x31\x27\x40\x7e\x41\x54\x05\x07\x28\x13\x7c\x8a\xa7\xc9\x0a\xc9\x05\xcf\x13\xc2\x1e\x29\x34\x85\x20\x18\xf7\x46\x1f\xfb\xdd\xde\xe5\xe4\xb7\x61\xef\xc4\x52\x0e\xe8\x0c\xfd\x03\x85\x33\xd4\x30\x8c\xe5\x4a\x6a\xea\x74\xde\xc2\x8a\xa7\x34\x0e\x79\x06\x4c\x2e\xe8\x4c\x85\x8c\x13\x68\xa0\xdf\x7f\x41\x6a\x01\x2c\x40\x08\xa1\x0a\xb9\x4d\xf8\x60\x46\x83\x20\x4f\x79\xce\x14\x6a\xa5\x4c\xb5\x04\x48\x9e\x8b\x18\xd0\xd7\xaf\x48\x89\x1c\x82\xf4\x6a\x26\x9b\xd7\x33\xa9\xb9\xb7\x08\x2c\x5b\x84\xca\xab\x16\xfe\x33\x17\x50\x02\x87\x19\x16\xea\x30\x80\x78\xc1\xd1\xa3\x9b\x81\x10\x6a\x2d\xb1\x68\x25\x74\xda\x22\x3c\xbe\x02\x81\x90\x26\x8e\xe6\x22\xfb\x94\x73\x85\x11\x6a\xa3\xf6\x23\xf4\xb7\xbf\x99\x85\xce\xa4\xc2\xd3\x40\xae\xa4\x82\x34\x56\x09\x92\x8a\x67\xc8\xe2\x35\x25\x88\x25\x8d\x21\x70\xc2\x57\xa9\x06\x02\xa4\xe2\x02\x62\xce\x50\x38\xda\x9a\xf5\x29\x62\xa1\x36\x49\x96\x6a\x8e\x39\x93\x94\x80\x40\x33\x1c\x2b\xa4\x16\x58\x6d\xed\x81\x8c\x25\x3d\x6c\x25\x39\x6b\xef\x52\xb6\xd6\x34\x16\x8a\x2a\xca\xd9\x0d\xe8\xbf\x78\xdc\x72\x95\x0b\x40\x52\x09\xac\x60\xbe\x42\x33\x2e\x90\x00\x49\xff\x04\x89\xe8\x2c\x38\x40\x0c\x80\x00\x29\xc5\x64\x5c\x79\x9c\x29\x53\xc0\x08\x65\x73\x83\xa7\x16\x54\x22\x19\x0b\x9a\x29\xa4\x38\x9a\x02\x12\x00\x4c\x09\xcc\xd4\x13\x34\xcd\x95\x01\x92\x78\x06\x6a\xa5\x09\x73\x85\x64\x06\x31\x9d\xad\x34\x01\x5f\xf9\x0b\x10\x9e\x35\xdc\xb0\x0d\x85\xe5\xd4\x9b\x83\x07\xb8\x56\x0b\xa8\x98\x38\x43\x20\x30\xc3\x79\xa2\x64\x9d\x21\x6c\xa8\x5a\x63\xed\x52\xb4\x99\xb3\x22\x34\x5e\x5f\x5c\x4c\xc6\x93\x51\x67\x78\xd9\xbd\x38\x3f\xeb\xbf\xb9\x3c\xef\xbc\xef\x9d\xe8\x93\x12\xda\x63\x14\xa6\x58\x2a\x10\x8d\x82\xdb\xfa\x7c\xfd\xf0\xc5\x3f\x3e\xdf\xcc\xf1\x0a\x02\x09\x04\x85\x14\x85\x80\x1a\xf2\xe0\xb4\xf7\xfa\xc3\x9b\xcb\xc1\xc5\x9b\x41\xef\x63\x6f\x70\x72\xb4\x39\x70\x7c\xd0\x40\x7b\x51\xd5\x8a\x20\x52\x21\xca\x90\x8a\xb3\x27\x47\x4f\x5f\xb4\x7f\x41\x84\x07\x07\x5b\x13\x3f\xbf\x2c\x21\xcc\x87\x17\xc7\xc7\x4f\x8b\x0f\xc7\xf6\x43\xfb\xd9\x53\x94\x13\xf7\x41\x8f\xbc\x6c\xbf\xb4\xe4\xfe\x92\x09\xae\xf8\xc9\x0f\x5f\x88\x54\x7f\xfd\xeb\x93\xc7\xdf\x82\xbf\x64\x5c\x28\x3b\x70\x70\xf0\xf8\xc9\xb7\xe0\x2f\x34\x53\x78\x9a\x80\x44\x61\x07\x5d\x8c\x2f\xcf\xfa\xa3\xde\xaf\x9d\xc1\xe0\xb2\x33\x18\x5c\xfc\x8a\xc2\x0c\xfd\x60\x88\xa0\x30\xd5\x07\x48\x01\x0a\x43\xfb\xff\xf3\xde\xaf\x7a\xb0\x98\x0e\x89\x26\x8d\x7e\x30\x7f\xc3\x3f\x50\xa7\xdb\xed\x0d\x27\x01\xe1\x0c\x82\xa0\x60\x12\x4a\xbc\x04\xb4\xb9\xf3\xc5\x6c\x10\x88\x14\x85\x62\x66\xf7\x50\xab\xb5\xf5\xd8\x7e\xb6\xae\xb1\x65\x75\xd7\x7a\x1c\x04\x53\x2c\xe1\xf9\x31\x0a\x09\x7a\xf5\xea\x15\xfa\xf2\x05\x75\x41\xa8\x8e\x7c\xbd\x52\x20\x51\xb3\x6b\xe8\x36\xf5\x18\x9d\xd1\x18\x2b\x90\xcd\x9e\x8a\x49\x17\x9b\x31\xf4\x15\xbd\x36\xf8\x3d\xa6\x8f\x2d\xfa\xf6\xcd\x89\x64\x58\xc6\xb8\x19\x0b\x75\x4f\x0e\x63\x10\x4b\x10\x7b\x70\x91\x16\xb0\x96\xd3\x50\xd0\x25\x56\xf0\x0e\x56\xfb\xf2\x7b\x07\xab\xbd\xd8\x5d\xc1\xea\x9e\x0b\x1b\xc2\x5e\xcb\xca\xe0\x3b\x2c\xca\xf0\xba\x75\x49\x86\x95\x5e\xd0\x36\xaf\xdf\x70\x9a\xbc\xc7\x42\x2e\x70\x52\x72\xe9\x90\x94\xb2\x77\xf9\x14\xac\xd1\xed\xa4\xed\x4c\x4d\x9f\x53\xf3\xa7\x79\x55\xe2\xdc\x63\xeb\x6e\xb3\x38\x9f\x9b\x33\xbc\x20\xbd\x22\x54\xe8\x93\x57\x63\xfa\x0c\xa7\x40\xfe\x49\xd6\x5f\xe5\x64\xff\xd7\xd4\x5b\x1d\xde\xfb\x48\xdc\x8d\xe5\x2e\x36\x7b\x9a\x4e\x17\xdf\x68\x34\x5b\xbc\xee\x77\x18\x2e\x8a\x78\xaa\xcb\x99\xe4\x09\xdc\x65\x81\x46\x7b\xad\xd8\x21\x3e\x64\xad\x5b\x52\xec\xbf\xf2\xaa\x10\xf7\xdb\x84\x33\xc1\x99\x1a\x0a\x7e\xbd\xba\x9b\x86\x67\x1a\x2f\xcc\x34\xe2\xfd\x8d\x6a\x6c\x43\xb7\x31\x9d\x33\xca\xe6\x77\x13\xc0\x85\x7d\xa1\xa4\x73\xf6\x40\x4f\xb5\x25\xc6\xfe\x2a\xd8\x90\xe2\xfe\x5e\xb9\x9b\x50\x60\xea\xde\xc7\xda\x62\x3f\xd4\x5d\x3b\x21\xf6\x5f\x7e\x8d\x0c\xf7\xdb\x82\xce\x7c\x2e\x60\x8e\x15\x17\x6b\x83\xbc\xcb\x66\xe0\x12\x3f\xf4\x2c\xf3\x41\x1b\x52\x2b\xd2\xfe\x5b\xb3\x43\xa2\xfb\x6d\xcf\x7b\x43\x53\xdf\x79\x09\xa8\x7b\x9b\xca\x95\xc5\xff\x1e\xd6\x52\x27\xd0\x9d\xcd\x66\x43\x9e\x87\x6c\x8d\xf5\x60\xf7\xdd\x18\xe7\xc6\xbe\xd7\xb6\xf8\xc2\xdc\x79\x53\x2a\xb2\x3c\x64\x4b\xf6\x09\x60\x6b\x25\xf8\x0e\x01\x6d\x45\x82\x3b\x6f\xc1\x4d\x21\xee\x30\x9f\x26\x34\xae\xe1\xef\x9c\x78\x27\x8e\x75\xbe\xf9\x0e\x56\xcd\x12\xf4\x6e\xbe\x1c\x5b\x0a\xb2\x99\x19\xfc\x1d\x62\xec\xdc\x87\x2d\x39\xee\xcb\xdd\x72\xa8\x67\x5f\x30\x7b\xab\x86\x58\xca\xcf\x64\x4f\x1e\x0b\x95\x19\xf0\xfd\xe2\xec\xf5\x19\xdf\x2f\xd0\x76\x4c\xd6\x65\x2d\xef\x98\xbb\xb0\x3b\x88\xb1\xf2\x23\x7f\xfd\xc7\xd6\xdc\x5e\xbd\xea\x5d\x9c\x05\xbd\x49\xf7\xf4\xb2\x73\xfa\xb1\x37\x9a\xf4\xc7\xbd\xcb\xee\xa0\xdf\x3b\x9f\x5c\x7e\x18\x0d\xc6\x27\x0b\xa5\x32\x19\xb5\x5a\x3f\xfc\xb8\xe0\x52\xe9\xd0\xe7\xa7\x48\x27\xd5\x16\xa7\xdb\x1b\x4d\x2e\xcf\xfa\x83\xde\x49\x6d\x62\x66\x61\x2c\x35\x03\xda\xf9\x30\x79\x7b\x62\x0a\x1e\x66\xea\xb4\x33\xe9\x5c\x9e\xf6\x47\x27\xd5\x52\x84\x99\xeb\x0d\x7a\xdd\x49\xff\xe2\xfc\x72\xd2\x7f\xdf\xbb\xf8\x30\x39\x39\x7a\xd6\x6e\xdb\xa9\xb7\xbd\xce\x68\xf2\xba\xd7\x99\x5c\xf6\xcf\x27\xbd\xd1\xc7\xce\xe0\xa4\x9c\xeb\x9f\xf7\x27\xfd\xce\xc0\x5b\xcd\xb0\xd7\x1b\xdd\xb4\x96\x17\x1b\x98\xdd\xc1\x87\xf1\xa4\x37\x3a\xb1\xdb\x18\xb6\xcd\xbf\x12\xb7\x32\x6a\xb0\x9f\xf8\x43\x87\xb5\x80\x87\xdb\x80\x47\xb5\x80\x47\x9e\x3c\xef\x7a\xbf\xed\xd8\x5a\x6d\x9b\x06\x64\xd0\x1f\x4f\x7a\xe7\xb5\xfa\x6a\x37\xcd\x7f\x9e\xae\x1c\xf0\xf6\x76\xac\x41\x0b\xd6\xa6\xe6\xe3\xed\x92\x1d\x35\x98\x75\x1a\x2f\x73\x56\x0f\x6c\xb7\xd2\xcd\x7c\xcd\xe2\xca\x74\x74\x0d\x35\x19\x69\x55\x9c\x5e\x76\x3b\x9b\xc0\x2e\xf6\x35\xa0\x7f\xff\x70\x31\xe9\x5c\xbe\xee\x74\xdf\xf5\xce\x4f\x2f\x5f\xff\x36\xe9\x8d\x4f\x8e\x8f\x5e\x1e\xbf\x7c\xfe\xf3\xd1\xcb\xe7\x16\xe6\x76\x4a\x17\x67\x41\x10\x57\x53\x46\x2f\xa9\xac\x19\x37\x57\x45\x11\x84\xdf\x86\x99\x5d\xd1\x56\x8c\x43\x25\x72\xa9\x5a\xb6\xc6\xdb\xc2\x2c\x5e\x70\x21\xbd\x93\xeb\x88\xe5\x19\xc1\x0a\xc2\x02\xde\x3f\xbe\x75\x8e\xdb\x15\xe6\x9a\x2b\x9c\x26\xee\x40\x63\x92\x52\x29\x29\x67\xd6\xa7\x44\x01\x42\x59\x92\xcf\xa9\xf7\x1d\xa1\x4e\xf2\x19\xaf\xe4\x30\x4f\x92\x7e\x8a\xe7\x20\xed\x28\x42\x96\x5c\x2e\xb0\xa2\x9c\x15\x83\x08\x5d\x51\x46\x22\x74\x6a\x4b\x8e\x9d\x2a\x83\x12\x08\x67\xf4\x23\x08\x3d\x11\xa1\xe5\x61\x39\x4c\xa8\xc4\xd3\x04\x22\x34\xc3\x89\x04\x33\xfc\x3a\xa7\x09\x71\xd4\x6e\x63\xbd\x83\xaa\x95\xa8\x42\xc8\x13\xc7\x8c\x5f\x2c\x41\x08\x4a\x6e\x5d\xdc\xed\x1c\x4a\x4a\x1e\x8b\x21\x27\x43\x01\x12\xd4\x43\xa8\xdf\xb0\xa3\xa5\x65\x34\x29\x6f\x19\x25\x0d\x79\x42\xe3\xd5\xfd\xd8\xc1\x35\xc4\xb9\x86\x1c\xe5\xc9\x7a\x43\x10\x0a\x51\x8a\x55\xbc\x30\xf4\x3b\x8c\x71\x65\xc8\x79\x00\x1a\xe4\x0a\x56\x11\xa2\xc6\x4e\x9a\x15\xb1\x08\xb0\x55\x58\x92\xf6\x70\x10\x5a\xe2\x24\x87\x08\x35\xf4\xd9\x6f\x78\x33\xda\xa7\x44\x6b\x71\x42\x02\x8c\x02\xf1\x00\x38\x1b\xb9\x46\xc8\x86\x14\x45\x7f\x24\x42\x19\x27\x72\xc7\xd4\x54\xab\xcb\x9f\x14\xf0\x07\xc4\x2a\xb2\x95\xf6\xf5\xb0\xbc\xa2\xd9\x85\xe1\x94\x18\x39\xce\x30\x4d\x72\x01\x1b\x70\x56\x49\xde\xe6\x3b\xfd\xac\x13\x0d\xef\x8c\xad\x23\xd0\x3e\x9b\x71\x2b\x7b\x0c\x42\x9d\x51\x6d\xfc\x37\x64\x4b\x86\x13\xac\x6e\x84\xd3\xbe\x11\x67\x74\x00\x4b\x48\x64\x14\x84\x5a\xb7\x1b\xaa\xc6\xb9\x5a\xac\xc5\x11\xf0\x29\x07\xa9\xde\x02\x26\x20\x9c\x30\x46\xb8\x6e\x27\x42\x35\x95\x04\x0f\x80\xa7\x29\x67\xe7\x38\x2d\x14\x10\xee\x10\x2a\xb0\x86\xa5\x04\xb6\x5c\x86\x02\x66\xf4\x7a\x8d\xf5\x1f\xe1\x08\x52\xae\x20\xec\x69\x98\xd0\x8c\xce\x05\xcf\x33\x0b\xbe\x0d\xf7\x46\x4f\x9a\xc1\x5c\x82\xd0\x96\xb2\x0b\xf2\x83\x04\x11\xc4\x9c\x29\xc1\x93\x04\x3c\x2d\x40\x02\xf1\xfa\x40\x24\x3c\xbe\x3a\x37\x06\xb7\x19\x21\x85\x6b\x64\x6d\x2d\x2e\x12\x34\x81\x24\x9b\xeb\xb8\xda\x12\xb0\x35\x86\xf2\xc8\x95\xda\xac\xa9\x84\x38\x8b\x29\xf4\x58\x53\xa5\xf0\x58\x46\xa8\xf1\xb8\x11\xc4\x5c\xc8\x4e\x92\xf0\xcf\x40\x2e\x8c\x6f\x97\x51\x40\x98\x5c\xaf\x66\x4a\x19\xe9\x10\x22\x40\xca\x08\x15\x57\xf5\x8b\xf6\xb3\xa7\x6e\xee\x1c\xd4\x67\x2e\xae\x22\xa4\xe2\xec\x38\x80\xb2\x9c\x50\x18\x60\x8c\x23\x54\x53\x8a\xf4\x57\xb2\xa3\xa4\xe1\xad\x64\x47\xc1\x01\xa1\x5c\x24\x46\x33\x21\xda\x19\x28\x6a\xa4\xb1\xe2\x02\xcf\x61\xbd\x2a\x1d\x9c\x0a\x06\x0a\xa4\x9b\xb2\x86\x13\x79\x13\x4d\xca\xeb\x00\xab\x9e\x4d\xeb\x74\xac\x75\xba\x41\xc6\x77\x51\x35\x60\x3e\x11\xe3\xd4\xd6\x92\xcd\xb8\x48\xb1\x8a\xfc\x70\xbf\xbf\x86\x38\x33\xb3\xe8\x2b\x02\x19\xe3\x4c\x07\xe3\x16\xdf\x77\x0d\x9a\x0a\x65\x4a\x5b\x6f\x32\x82\x39\x95\x4a\xac\xde\xba\x3d\x89\x5c\xfb\x34\x14\x6e\xa2\xe9\x7a\x78\x4d\xb9\x8c\xa3\x67\xed\x76\x3b\xb0\x0e\xc7\x66\x01\xce\xd7\x5c\xf9\x09\xbf\xaf\xd8\xdd\xca\xac\x29\x3a\x6c\xeb\xb3\xa6\x12\x80\x50\xc6\x85\x8a\xd0\x61\xfb\xe8\x59\x3b\x58\xef\xbe\x2f\x8f\xe6\x8e\x33\x6a\xf3\xcc\x8e\x98\xe7\x29\xb0\xe2\x0e\x8f\x13\x9e\x13\x17\x91\x14\x47\xd6\x8f\x5c\xcc\x7c\x26\xf8\x92\x12\x10\xb6\xcd\x69\x32\x11\x0f\xb9\x98\x2d\x3d\x8f\x06\x32\x9f\x45\xce\x14\x4d\x61\x83\xbc\x04\xa5\x28\x9b\xcb\xe6\xd5\x0b\x6d\x34\xad\xe5\x21\x4e\xb2\x05\x3e\x3c\x29\xfd\xb8\xb4\x5a\x0f\xa7\x38\xbe\x02\x46\x0a\x44\x6d\x99\x4f\x2b\x00\x29\x10\x8a\x43\xb5\xca\xa0\x64\x9e\x65\x89\xce\xae\x29\x67\xad\x25\x23\x4d\xcf\x3e\x4d\xdf\x6e\x9a\x6b\xd1\xd7\xc7\xfa\x5f\xb9\x1d\x71\x92\x1b\x3f\x26\x6d\x05\x35\xd4\x46\x10\xce\xb4\x82\x6b\x38\x55\x5b\x04\x75\xe8\x57\xb0\xda\x03\xdb\x1a\x89\xfd\xde\x1f\x46\xe8\xf0\xe8\x67\xe3\x94\x0e\x6f\xbf\xff\x76\x95\x7d\x2a\x4e\x73\x57\x3d\x06\x21\x19\x2f\x80\xe4\xa5\xab\xb7\xe0\x75\x29\x7d\x01\xd7\xfc\x43\x9a\x88\xc4\xf9\x61\x39\xce\xa7\x0c\xb4\x6d\xff\x7c\xd4\x7c\x6a\x1c\x69\xeb\xf0\x79\x60\xb1\xac\xd4\x46\x6b\xa5\xef\x18\x70\x9e\x69\x93\xe9\xba\x3b\x91\x31\x7b\xb1\x6c\x04\x95\x38\x8e\x21\xd3\xd3\x0a\x98\x9a\xac\x32\x90\xd1\x3e\x66\xf3\xc4\x87\x71\x92\x22\x34\xcd\x85\x54\x11\x7a\xde\x6e\x07\x2e\xc2\x2b\xa8\xee\x45\xd4\x20\x7d\xca\x64\x84\x9e\x1a\x0a\x5b\x6b\x79\x97\x4f\x0b\x67\xb7\x75\x21\xfa\x25\x03\x3b\x62\x6b\x39\x1f\x46\x03\xe3\x0f\x33\x41\x99\x42\x8d\xc2\xd3\x37\x8c\x83\x54\x98\x32\x5b\x73\xa2\x31\x34\x87\x82\x67\x20\x14\x05\xd9\xbc\x10\xf1\x02\xcc\x63\x19\x5c\x0c\x05\xd7\x96\x65\xba\x2e\x63\xd7\x75\xd1\x9e\xd5\xd2\x2f\x9c\x23\xfa\x8a\x3e\xe5\x5c\x19\xaf\xca\xec\xb5\xb6\x76\x38\xce\x5a\xdd\x75\xe7\x2e\x9d\x98\x12\xa1\x7d\x55\xf3\xf0\xe8\x85\xd5\xe7\xb1\xd9\x01\x7d\x05\x59\x6d\x0f\x80\xcd\xd5\x22\x42\x2f\x03\x13\xa7\x18\xa7\xdc\x1f\x3a\x2a\xdd\xfe\xe9\xc8\x51\x72\x37\x6b\x4b\x6f\x9a\xe3\x3d\x34\x19\x93\x8d\x1d\x04\x90\x05\x56\x5e\xaa\xc6\x97\x32\x94\x86\xc3\xda\xc0\x3c\xaa\x9b\x46\xc6\xab\x91\x19\x96\x12\xd4\xbf\x78\x73\x1b\x45\xcb\xaa\xd5\xf0\x37\x5a\xc7\x63\x98\x29\x3f\x31\x4c\x41\x2d\x38\x89\x10\xce\x95\xbe\x3c\x29\x01\xa6\xa8\x5a\x0d\x9d\x23\x72\x3b\x96\xf0\x39\x65\x5e\xb0\x9c\xe2\x2c\xa3\x6c\xfe\xde\x21\xc7\x09\xa6\x69\xb0\x0e\xf7\x3b\xda\x6d\xa1\xce\xa9\x19\xaa\xfa\xb4\x1d\x19\x8b\xa1\xe0\x25\x00\x90\x62\x9a\xf8\x69\x8b\x19\x28\xbf\x53\xe2\xcf\xc9\x7c\x1a\x54\xf2\x0d\x6f\x4e\x7f\x2f\xbf\x66\x02\x66\x20\x04\x90\x0f\x2e\xde\xf4\x21\x73\x46\x3f\xe5\x70\xe9\x21\x58\x8f\xd4\x3f\x2d\x03\x84\x9d\x2a\x72\x43\x43\x41\x59\x4c\x33\x9c\x14\x6a\x72\xfe\xf1\xb4\xaa\x84\x35\xed\x31\xc4\x02\xd4\xfd\xe9\x5b\xfc\x6d\xea\x36\xae\xd0\x46\xd2\x3f\xed\x6f\xa8\xd4\x81\x14\xd1\x9c\x53\x4a\xae\x16\x5c\xd0\x3f\xa1\xce\x3e\x8d\xf6\x9b\x29\x8d\x05\x97\x7c\xa6\x38\x4b\x28\xd3\x57\x57\xea\x2c\x57\x1b\xe0\x04\x18\x36\x2b\x6d\xb4\x8c\xfd\x1f\xb5\x4a\x92\x8d\x6d\xf9\x10\x52\xfc\x0a\xd8\xf7\x63\x66\xc8\x6d\x30\x0a\x51\xbc\xc0\x49\x02\x6c\xee\xe7\x79\x77\x34\xe5\x01\x8f\x71\x82\x4c\x4d\x97\x0b\xb2\xbf\x41\xcf\x76\xdd\x58\x65\x81\xd8\xd7\xd4\xdb\x89\x2d\x32\x0f\x1d\x9f\x1a\x9d\xb9\xab\xab\xe3\x05\x82\xff\xd3\x7e\xbb\x90\xe0\x7e\xbc\xcf\xfe\x7e\x7a\x5e\xa5\x26\x61\xa3\x8e\x55\x0e\xbd\xc7\xd7\x9d\x39\x8c\xf5\x85\x45\xf4\x7d\x57\x5c\x99\x6e\xda\xfa\x6d\x29\x99\x3f\x68\xcf\x86\xdc\x1d\x3a\x58\xb0\x50\x5a\x38\x53\x52\x0b\x9c\x61\xfa\x22\xe8\x2b\x5f\xca\x89\x1e\xde\x10\xe3\xc5\xf3\x63\x27\x47\x69\xeb\x75\x60\xcf\xda\xed\x20\x13\xfc\x0f\x88\x3d\xcf\xeb\x52\x81\x73\x4e\x60\x6c\x92\x58\x2e\x22\x64\x1e\xbd\x13\xe6\x49\x03\x3f\x35\x6a\xc5\x3c\xcd\x72\x05\x45\x90\x2b\x21\xce\x05\x55\x2b\x9d\x4b\xc6\x5a\x75\xce\x95\xc7\xb2\x1c\x19\x61\x63\xf5\xb2\x1d\xb5\x8e\x8a\xc9\x01\x9e\x42\x22\x87\xa6\x79\x66\x8b\x24\xcf\x6c\xf2\x4d\xc9\x26\xde\x61\xbb\xf8\x17\x1e\xbe\x2c\xfe\xb5\xcc\x68\x20\x78\xae\x43\xf0\xf5\x52\x64\x3e\x25\x3c\xc5\xd4\x1e\xe7\x1f\x29\x23\x70\xfd\x70\xb3\x1b\xf1\xdc\xf6\xf9\xf4\xac\x44\xed\x9f\x9c\x21\x8e\x0b\x6e\xbe\xed\xc8\x4a\x47\x68\x2d\x5a\x8a\x19\x9e\x03\x29\x8b\x1b\x61\xb1\xef\xe6\xb3\x29\x1e\x99\xc3\xa5\xc7\xb3\x84\xaf\x76\x9d\xb4\xac\xec\x45\x55\xb2\xfd\xda\x4e\x12\x42\x59\xd1\x15\xd3\xc0\x8e\xef\x0d\x9d\x2f\x69\xab\x10\x45\x14\x5d\x5b\x04\x38\x3e\xae\xaf\x01\xd4\x84\xdc\x5e\x4f\xc6\xaf\x01\x95\x6b\xd9\x8c\xbe\xbd\x46\x83\x5e\xfb\xf5\xc8\x56\x92\x64\x9f\x9d\x25\x74\xbe\x50\xd6\x80\xcb\x0a\xd3\x84\xa6\xc0\x73\xb5\x79\x16\xcd\x83\x32\x7e\x83\xd2\xd8\x64\xe8\x89\xb7\xd7\x23\x3e\xd5\xa2\xca\x5e\x0f\xe4\x94\xd7\x7e\x79\x9d\x85\xb7\xde\xa7\x0f\xf3\x7b\x4b\x9e\xe4\xa9\x57\x46\x20\x2b\x86\x53\x1a\x1b\x67\xad\x5d\x0a\x65\xf3\x1e\xc3\xd3\x04\x88\xbb\x67\x6c\xc3\x61\x47\x61\xbf\xce\x0f\xa1\x57\xaf\x1e\xf5\x2e\xce\x1e\x6d\xd6\xfc\xec\x6d\x31\xae\x38\xb8\xc0\x21\x46\x41\x68\xfc\x90\xbe\x39\x6c\xe2\x50\x29\x6c\x54\x90\x3a\xb9\x5a\x6c\xf6\x1a\x2b\xae\x18\x58\x2c\x56\xd9\xcd\x44\x7a\x2c\xbe\x81\xc6\xcd\x6b\xae\xa4\x6d\xe5\x6a\xbf\x04\x08\x35\xd6\x4b\x6e\x44\xa8\xb1\x3c\x6c\x3c\xd1\xa3\x7a\xe5\xfa\xbb\x2d\xbe\xd8\xb1\x4c\x00\xb1\xc6\xd6\x88\xd0\x3f\x8c\xf2\xbf\x38\x13\x68\x68\x9d\x69\xf8\x73\xfe\xd1\x68\xeb\x3f\x39\x33\x1a\x4b\x68\xac\x6c\x7d\xfa\xdb\x93\x7a\x8c\xf7\xf8\xba\xf7\x7a\xfc\xd1\xe9\x38\x67\xb7\x83\xbf\xe9\xf6\x86\xa7\x77\x41\x30\xf1\xf1\x29\x95\x57\x77\x40\x52\xf1\xa2\xcf\xb4\x47\xe4\xa4\x33\x9b\x51\x46\xd5\xea\x66\x94\x73\xae\x39\xec\xb7\xe6\x37\xc0\x40\xe8\xb0\xb2\xdc\xd0\x1b\xc1\x87\x9c\x4c\x78\x02\x42\x43\xea\x1b\x6c\x82\x29\x53\xb7\xe0\x74\x17\x10\x5f\x69\xe0\xf7\x90\x72\xb1\x1a\x6a\xff\x96\x0b\xd8\x13\x49\x2f\xe5\x0e\x28\x76\x5f\x5f\x53\xf3\xb8\x7f\x3d\x3c\x76\xc5\x9b\x46\x54\x8e\x21\xd4\x28\x2e\x92\x62\x87\xfd\x49\x84\x1a\x89\xb9\x42\x4b\x6b\x2b\xc7\x05\xcc\xb5\xbd\x7a\x83\xbf\x97\x9f\x8b\x98\xd7\x09\xe0\x89\x3a\xf2\xb0\x34\xd4\xef\xce\xaa\x29\x17\x54\xbb\xa9\xdd\x56\x5d\x84\x0c\xe3\x4c\x00\x26\x43\x8b\x61\x4f\x85\x81\xfb\x0c\xda\x73\x37\x22\x74\x78\xe3\x5e\x6d\x1a\xd4\xbd\x09\x0d\x00\x4b\xe5\x2e\x0e\xb8\xbf\x3c\xaf\x71\x82\x59\x0c\xa4\xe8\x08\xb9\xb0\x44\x6f\xd2\x5d\x49\x69\xa3\x19\x9a\x74\xaf\xb3\xe4\x94\x0c\x39\x91\x37\x89\x65\x22\x9b\xdb\xe8\x3d\x78\x9b\xcc\x39\x71\x27\x87\x72\x76\x67\x3a\x37\xdb\x2c\x53\xf4\x46\xbb\xd5\x12\xfc\xc9\x99\xd7\x9e\xdb\x6d\x9a\xda\x5d\xd6\x88\x75\xb4\xb6\xd5\xc0\xb9\x78\xfb\x36\x8b\xb9\x7f\x4d\x20\x81\x0e\x9f\xbf\x68\x3e\x7f\xda\x3c\x3c\x7a\xd9\x3c\x7c\xfe\xa8\xe6\xd1\x6d\x01\x92\x27\x4b\x5b\xfa\xac\x7d\x7c\xbb\x52\x28\x5d\xbf\x74\x64\xde\xd8\xa1\x12\x91\xdc\x96\xc6\x80\x3c\x41\x98\x11\xf3\xfa\xd1\x23\x69\x5e\xd1\xa1\x04\x70\xb2\x7d\xeb\xec\xaa\xbc\x96\x17\x8f\xb2\x09\x25\xa9\x5c\x76\x65\x96\xe9\x07\x9a\xf9\xd4\xbe\x33\x44\x39\xdb\x00\x1f\xfb\x53\x15\x24\x8c\x8b\x3e\x0d\xf9\xbe\xa5\x84\x92\xf0\xf7\xaf\x23\x60\x4c\x26\xfb\x6e\x4a\xd1\x8e\x35\xfd\xbc\x0a\xf4\xc8\x9f\xf1\x51\x8a\x73\xbd\x43\xe4\x81\x9b\xae\xc6\xf8\x36\xeb\x31\xb4\x6c\xc6\xc7\xe4\x3c\x74\xb9\x51\x90\x09\x9a\x62\xb1\x1a\xc7\x38\x81\x31\xa8\x22\x25\x2c\xe7\x97\xa9\x2d\xac\x2e\x53\x29\xad\xe5\x16\x76\x45\xd9\x12\xa4\xa2\x73\xac\x00\xa9\x05\xa0\x30\x4c\x31\xa3\x33\x90\x2a\xcc\x45\x82\xdc\x13\x98\x28\xc3\x02\xa7\xa0\x40\x18\xa3\x93\x00\x88\xce\x10\x55\x28\xd5\x67\x23\x38\x40\x0b\x48\x32\x94\x4b\x84\x15\xc2\x49\x8d\x11\x1a\xdb\xcf\x38\x91\xf6\x91\x2c\xff\x09\x8e\xba\x20\x6f\xc8\x49\x90\x82\xc2\x04\x2b\x1c\x05\x65\x23\x5d\xc5\xc4\x7d\x91\x19\x8e\xc1\x36\xd1\x42\xfb\xde\x5e\x20\x33\x88\x23\xd7\xa2\x30\xfb\xe9\x92\x0f\x2c\xe6\x65\x4f\xf5\xab\x3b\xd4\x12\x14\x0a\xb1\xfb\xd2\x44\x35\x0f\x8c\xb9\x39\xb8\x86\xb8\xe0\xab\x49\xa7\x29\x5e\xb7\x55\xcc\x6b\xa0\x72\xe1\xbe\x85\xb1\xf9\x60\xba\x64\x15\x3b\xb0\x1d\xa5\x9e\x8a\x89\xe9\xaf\x6d\x16\x84\x6c\x5b\x2d\x4f\x12\xf7\xc8\x83\x7b\x5c\xc5\xd6\x6b\xe8\x12\x18\x48\x39\x14\x7c\x5a\xd6\xeb\xb4\x4c\xeb\xe2\x55\x45\x26\x54\xb6\x7b\x62\x95\x78\x23\x61\x18\x63\xd3\xf2\xf0\xc6\xb6\x9e\x0c\xaa\x80\x17\x1d\x96\x5a\x84\xf2\x39\x28\x1f\xa5\xe8\xaa\xec\xc6\x28\x13\x16\x87\x01\x8c\x64\x5c\x07\x4f\xde\xe8\xce\xee\xea\x1a\xa4\x68\xe7\x2c\x00\x27\x6a\xe1\x26\xb4\xe7\xa7\x38\x39\x85\x04\xaf\xca\xf4\xec\xf8\x99\x57\xcc\x2a\xb5\x58\x9c\x25\xd3\xc8\xb8\x2e\x1f\x69\xd1\x69\x2c\x4d\x60\x5e\x26\x2e\x7a\xd0\x66\x3b\xef\x4d\xda\x5a\xa8\xdd\xbc\x70\x38\xc4\x6a\x11\xad\x17\xe8\x68\xac\x39\xb9\x56\x98\x1b\xd7\x81\xca\x05\x4b\x56\x1e\xe5\x2a\x9d\xca\x63\x81\x5b\xb4\xf4\x31\x30\xa3\x3a\xf9\xa5\x6c\x7e\x4a\xc5\x36\x8e\xde\xaf\x75\x7e\x6c\xd9\x58\xf1\x8b\xc6\x35\x97\x96\x5d\xb1\xe0\x9a\x25\xd4\x2d\x60\x27\xe6\x96\xd0\x9b\x22\xd7\xa7\x41\x6b\x5f\x80\x33\xba\x47\xc6\xb7\xe9\x0c\x6c\x1c\x6a\x65\xa9\x3c\xa1\xa3\xbd\x1d\x67\xc0\x54\x84\x70\x46\x4b\xb7\xb1\xfe\x7c\x4f\xaf\x61\x5e\x07\x2e\x94\x66\x0e\xb3\xfb\x62\x29\xbb\x13\x63\xb6\xeb\x64\xaf\x87\xd7\x4a\xa4\x84\xcf\x13\x58\x42\x72\x72\x5c\xe7\x5f\xd6\x2f\x64\xd7\x3b\x96\xae\x6d\xc6\x0e\x13\xcc\xe0\xfb\xb8\x16\x7d\xfa\xde\xac\x9f\xf2\x2a\x34\x6d\x4f\xda\x9f\xeb\x51\xd3\x39\xb7\x65\x19\xe7\x50\xe3\x05\xe8\xbd\x7e\x3b\x99\x0c\xc7\x7b\x1c\x49\x84\xd4\x46\x29\xe5\xb0\xed\x99\x50\xb1\xb1\xfa\xdc\xd0\x3b\x4a\xd9\xd2\x48\xab\xef\x21\xab\x13\x69\xa7\xac\xdf\xdb\x91\x54\xcc\xa6\xe2\x05\x2a\x26\xb4\xaf\x4f\xd9\x19\x54\xd6\x51\xae\x74\xe0\x77\x71\x78\x88\x8b\xd9\x5e\x5c\xfd\xd2\xf6\x21\xb2\xbd\x9c\x9b\x16\x73\x9b\x17\xf2\x1e\x55\xba\x87\x37\xb2\x9c\xab\x4f\x58\xfd\x33\x7c\x4d\x95\xc3\x43\x7c\x0e\x95\x0a\xd8\xd6\x43\xd1\xc7\xc7\xc7\xff\x57\xdc\xd2\xf1\x0d\x47\xbd\x4e\x5d\xff\x7f\x92\xff\x37\x9d\xe4\xe2\x77\x16\x84\xf9\xa1\x85\x1f\xd1\x63\x9b\x7a\x47\xe8\xa7\xe6\xe3\x83\xff\x3a\xac\x09\x20\x8b\xdf\x5a\xd8\xff\x3d\x92\x07\xf3\xa8\x7d\x29\xfc\x00\xbd\xed\x74\xdf\x69\xeb\xcb\x56\x68\x63\xd2\xfc\x0c\x07\xe7\x4a\x2a\x81\x33\x7f\x5c\x72\xfb\x13\x23\xa5\xb8\xee\xd7\x49\x34\x3e\xca\xb5\xa0\x94\x99\xcc\x4d\xae\x58\x1c\x1c\x20\x82\x21\xe5\x4c\x27\x38\x9f\x69\x92\x98\xc2\xc0\x0c\xd3\xc4\xe5\x6d\xca\x80\xda\x05\x5b\x12\x36\x9a\x42\x31\x17\x02\x62\x95\xac\x9a\xb5\x8f\xf8\x6f\x4a\xbb\x05\x50\x27\x7b\x60\x7f\x3b\x04\x22\x54\xf7\xdb\x16\x28\x16\x58\x2e\x50\xc2\x79\x26\x51\xce\x14\x4d\x0a\xb9\xa8\x44\x79\xe6\xfd\x1c\x0b\x98\x7e\x42\x2d\x91\xf2\xd7\x59\x36\x7f\xbc\xe5\x26\x60\xf4\xef\xfe\xab\xf3\x82\x73\xd5\x32\x52\x6f\xae\xfc\x76\x43\xf1\xb1\x5b\x6e\xd1\xff\x1d\x00\x00\xff\xff\xde\x42\xfa\x60\x1f\x48\x00\x00")

func masterStartupShBytes() ([]byte, error) {
	return bindataRead(
		_masterStartupSh,
		"master-startup.sh",
	)
}

func masterStartupSh() (*asset, error) {
	bytes, err := masterStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x71\x4f\x23\xb7\x13\xfd\xdf\x9f\x62\x7e\xf0\x53\xd3\x9e\xba\x6b\xb8\x9e\x50\xcb\x01\x12\x84\x80\x10\x1c\x89\x48\x7a\xd2\xa9\xaa\x90\xd7\x9e\x10\x2b\xbb\xb6\x6f\x3c\x1b\xb1\x05\xbe\x7b\xb5\xbb\x21\x24\x81\xe3\x68\xfb\x5f\x34\xf3\xe6\xcd\xdb\x19\xfb\x39\x9b\xff\x93\x99\x75\x32\x53\x71\x02\x09\xde\x0a\xb1\x09\xa3\xfe\x71\x7f\x17\x24\xb2\x96\xc6\xc5\x42\xc5\xaf\xa9\x91\x9e\xec\x8d\x75\x49\x19\x22\x13\xaa\x22\x31\x2e\xa6\xda\xbb\x31\xd8\x08\xba\x24\x42\xc7\x79\x05\x13\x45\x46\x7b\x83\xe6\x23\x58\x16\x9b\x10\xc8\x67\x2a\xcb\x2b\x88\x13\x5f\xe6\xc6\x75\x18\x32\x14\x62\xd8\xbb\xfa\x7c\xd6\xed\x5d\x8f\xbe\x0c\x7a\xfb\x2d\xb3\xb0\x63\xf8\x03\x92\x31\x6c\x34\x8d\x63\x15\x6b\x76\x7b\x23\x15\xfb\xc2\xea\xc4\x07\x74\x71\x62\xc7\x9c\x38\x6f\x70\x03\xfe\xfc\x08\x3c\x41\x27\x00\x00\x56\xe8\xd6\xf1\x62\x6c\x85\x28\x0b\x5f\x3a\x06\x59\x38\x96\x84\xd1\x97\xa4\x11\xee\xef\x81\xa9\x44\x51\x4c\xc7\x31\xbd\x1d\xc7\xba\xbb\x34\x38\x93\xc6\xc6\xa9\x54\x7f\x95\x84\x0b\x70\x12\x14\xf1\xb6\x40\x3d\xf1\xd0\x79\x1d\x04\x20\x67\x8a\x64\x6e\x33\x69\xbc\x9e\x22\x01\xd4\xe4\x70\x43\xe1\x6b\xe9\x59\x01\x6c\xc1\x56\x07\x0e\x0e\x9a\x0f\x1d\x47\x56\x99\x88\x55\x64\x2c\x34\xe7\x10\xd9\x07\x68\xeb\xd2\x88\x34\xb3\x1a\xc5\x5c\xfc\x2a\xab\x20\x8c\xec\x09\xb5\x77\x90\x5c\x3d\xcb\x2e\x33\x2a\xe2\x75\xca\xf9\x97\x1c\xf5\xfb\xa3\xe1\xe8\xea\x70\x70\xdd\xed\x5f\x9e\x9c\x9d\x5e\x5f\x1e\x7e\xea\xed\xd7\x13\x4e\xda\xf1\x27\x77\x77\x90\xf6\x6e\x99\x54\x7a\xe5\x73\x84\x87\x87\x85\xf2\xa7\x15\xfd\xff\x6e\x79\x03\x0f\xcd\x86\x84\x88\x68\x20\xb1\x90\x20\x6c\xc4\xcd\xe3\xde\xd1\xef\xa7\xd7\x17\xfd\xd3\x8b\xde\xe7\xde\xc5\xfe\xfb\xf5\xc0\x87\xcd\x0d\x78\x13\x2b\x15\x90\xd0\xb8\xc5\x22\x6b\x23\xdf\xb5\xbf\xdb\x53\x24\x0b\x15\x19\x49\xbe\x13\x22\x53\x11\x77\x3e\x40\x62\x60\x6f\x6f\x0f\xee\xee\xe0\xa8\x09\xf4\x5c\x7d\x3e\xe1\xc7\x2f\xaa\xc8\x3f\x29\x8a\x13\x95\x43\xda\x6d\x3a\xa6\x97\xde\xe0\x91\xf7\x1c\x99\x54\x38\x2f\x33\x6c\x95\xfc\x04\x0f\x0f\x70\xb0\xdc\xa5\x96\x22\xb3\x47\x64\x3a\x5d\x40\xbf\xd7\xb5\x8b\xc4\x87\xf1\xa8\x62\x8c\x8b\xae\x75\xcc\x8e\xad\x56\x8c\x71\x55\x42\x93\xfa\x46\xf7\x66\x47\x0b\x09\x01\x29\xd5\xc4\xdf\x6b\x3f\x20\x3b\x53\x8c\xe7\x58\xfd\x03\x11\xe7\x58\xbd\x59\xc3\x14\x2b\xa1\x27\x85\x37\xb0\xb5\xb3\xb5\x05\x6f\xab\x78\x0e\x7b\x71\xb4\xff\x79\xb6\x5d\xf5\xda\x40\xb5\x6a\x26\xa8\xc3\x73\x39\x6d\xaa\x8d\x87\xa9\x95\x5a\x25\x4c\x65\x64\xd9\x5e\x7b\xa9\x9c\x9e\x78\x8a\xf2\xc9\xa3\xe6\x64\x65\x30\x8a\x31\x79\xc4\x3f\xde\x3a\xa7\x0a\xac\x2f\x22\x12\x6c\xef\xfc\x9a\xee\xfc\x92\x6e\xbf\xff\x2d\xdd\xde\xe9\xbc\x20\xab\xf6\x96\x7c\xd6\x58\xad\x28\xa6\xc6\x12\x24\xab\x0a\x75\xee\x4b\x13\xc8\xcf\xac\x41\x7a\xb2\x6f\x9e\xd8\x58\x7b\xb3\x29\x43\xde\x7c\xbf\xf9\x19\x94\x33\xc0\x13\xc5\x9d\x08\xce\x33\x58\x83\x2a\x17\x5a\xf1\x6a\xdb\x15\xc2\xd6\xe1\x5a\xa7\xdf\xdb\xeb\xf4\xfa\x27\x1d\xc1\xe8\x94\xe3\x33\xb3\x5b\xef\xe0\x71\xd2\xa3\x36\x78\x0c\xf7\x50\xbb\x5c\xed\x14\x22\x96\x59\xd4\x64\x03\x5b\xef\xd6\xe0\xc3\xe5\xd4\x4a\x91\x52\xa6\x9b\x5b\x5c\x69\xc0\xca\x3a\xa4\x61\xeb\x5d\xe9\x80\x7c\xa8\xf7\x8a\x31\x9d\x87\x06\x64\x9d\xb6\x41\xe5\x03\xf2\x63\x9b\x63\x3a\x67\x78\x99\x78\x88\x9a\x90\xff\x3d\x79\x5b\xbf\x46\x3d\x7a\xeb\x50\x1e\x1f\x8b\x53\xf2\x65\x58\x41\x5f\x2d\x67\x96\x4b\x72\xaf\x55\x3d\xa8\x6f\x48\xbe\x98\xa7\x57\x46\x8f\xba\x24\xcb\x55\xc3\x75\xa9\x0a\xdc\x05\x17\x6f\x12\xed\x8b\x50\x32\x8a\x40\xb6\x50\x54\x0d\xb5\xca\x71\x88\xdc\x02\x62\x5c\xe4\x67\xc5\xa8\x0a\xb8\x0b\xb3\x22\x46\xd1\xeb\x9f\xd4\xe7\xca\x79\xc6\x5d\x78\xc9\x95\x41\x53\xfd\xf7\x21\xf7\x3e\x44\x28\x1d\xdb\x1c\x5a\x1f\xae\x0f\x60\x19\x96\xde\x22\x74\x2a\xcb\xf1\x45\x92\xc5\xd3\xb4\xfe\x72\xbd\x06\x86\x1f\xc4\xdf\x01\x00\x00\xff\xff\xe5\x33\xaa\xbb\xc1\x08\x00\x00")

func nodeStartupShBytes() ([]byte, error) {
	return bindataRead(
		_nodeStartupSh,
		"node-startup.sh",
	)
}

func nodeStartupSh() (*asset, error) {
	bytes, err := nodeStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"azuredeploy.json":  azuredeployJson,
	"master-startup.sh": masterStartupSh,
	"node-startup.sh":   nodeStartupSh,
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
	"azuredeploy.json":  {azuredeployJson, map[string]*bintree{}},
	"master-startup.sh": {masterStartupSh, map[string]*bintree{}},
	"node-startup.sh":   {nodeStartupSh, map[string]*bintree{}},
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
