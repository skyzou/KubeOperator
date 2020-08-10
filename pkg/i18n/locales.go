// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package i18n generated by go-bindata.// sources:
// locales/en-US/home.yml
// locales/zh-CN/home.yml
package i18n

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _localesEnUsHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x55\xcd\x6e\xe3\x36\x10\xbe\xf3\x29\xc6\xf2\xb5\xe8\x03\xf8\xc6\x4a\x4c\xac\xae\x2d\x09\xfa\xd9\x36\xbd\x08\x34\x39\x5e\xb3\x91\x49\x81\xa4\x90\xa6\xb7\xbe\x57\xdf\xa9\xaf\x50\x50\xa6\x1d\xef\xc6\xc1\xba\x68\x2e\x81\x01\x7e\x33\xdf\x37\xdf\x37\xa3\xa5\x30\xc7\xa3\xd1\xa4\xa0\x5b\xd6\xb3\x5f\xf3\xa6\x6d\x56\x90\x14\xfc\x88\xc0\x07\x8b\x5c\xbe\x02\xfe\xa1\x9c\x77\x09\x21\x4b\x31\x4c\xce\xa3\x25\xe9\xa6\x6b\x5a\x56\xf7\x19\xdb\xb0\x96\xf5\x0f\x34\xdf\xb0\x6c\x05\x89\xe0\x1a\xb4\xf1\x20\x71\x40\x8f\x10\x9f\x83\xd2\x20\x26\x6b\x51\x7b\x70\x9e\x7b\x4c\xc8\x52\x58\x94\xa8\xbd\xe2\x03\xf9\xaa\x48\x5f\xb3\xa6\xec\xea\x94\xad\x20\x79\xe0\x6a\x40\x09\xde\xc4\x7a\x0b\x68\x0f\x68\x11\x94\x03\xae\x81\x3b\x67\x84\xe2\x1e\x25\x1c\x8c\xf3\x30\x69\x89\x16\xfc\x41\x39\x78\xc6\xd7\xe4\x83\xb2\xfd\x6f\x65\xf1\x9f\x6a\xff\x69\x34\xde\xa8\xfd\x40\xbb\x4d\xdb\xa7\x35\xcb\x58\xd1\xe6\x74\xd3\xa7\xb4\xe8\x8b\xb2\x8d\x23\x59\x41\x92\xe1\x9e\x4f\x83\x87\x37\xa5\x20\xb8\x0e\xd3\xd9\x61\x6c\x2a\xc3\x4c\x03\x79\xb2\x2e\x9b\xb6\xa7\x9b\x9a\xd1\xec\xe9\xcd\x85\x75\xd0\xf5\xad\x0b\x51\xd7\x8c\xb8\x0c\xfe\xa6\x9c\xd3\x5c\x82\xa2\x58\xe2\x4a\xd6\x8b\xf2\x07\xf0\x87\x8b\x47\xb7\xea\xf6\x3f\x3d\xf5\x55\x5d\xfe\xcc\xd2\xf6\x7f\xb5\x18\xad\xf9\x1d\x85\x4f\x48\xf3\xd4\xb4\x6c\xdb\xe7\xd5\x3c\xa9\x87\xb2\x2b\x02\xf7\x6a\x40\xee\x10\xf6\x6a\x18\x42\x52\x02\x62\x30\x82\x0f\x90\x57\xb0\x57\xd6\xf9\x7f\xfe\xfe\x2b\x0c\x6a\x72\x68\x49\x45\x9b\xe6\x97\xb2\xce\xe6\x0a\x5b\xda\xa6\xeb\x15\x24\x81\xc8\xc8\x9d\x7b\x31\x56\x06\x32\x4a\x0b\x63\xed\xdc\xb2\xac\xf3\xc7\xbc\xa0\x9b\x77\xef\x8d\x55\x5f\x94\xe6\xc3\x47\xc0\xae\x61\x75\x9f\x37\x33\x8e\xa6\x6d\xfe\x99\x45\x60\xa0\x11\xde\x06\x27\x51\xf3\xdd\x80\x72\x01\x51\x83\x30\xda\x73\xe1\x67\x0d\x5c\x1e\x95\x56\xce\x5b\xee\x8d\x5d\xc4\x82\xd7\xba\x0b\x03\x6e\x12\x87\xb9\xe0\x2c\x91\x66\xdb\xbc\x78\x1f\xa4\xd0\x54\xc6\x30\xcd\x45\x4f\x14\xde\x85\x69\xf1\x35\xe9\x9a\x6d\x68\xcb\xb2\x2b\x07\xbb\x00\x3b\xf0\x40\xfd\xda\xa7\x68\xcf\x0d\x0a\x5d\x95\xd1\xfb\x28\xa0\x54\x27\x06\x84\x2c\x2d\x7e\x51\x46\x9f\xf3\x54\xb3\xc7\xbc\x2c\xee\xdd\x6e\x38\x81\xbf\x97\xa8\xb0\x94\x21\x12\xe1\xff\xb9\x51\x58\xec\xbb\xdb\xcc\x5b\xfd\xbd\xd8\x0e\x5c\x87\x26\x66\x44\xed\x3c\x17\xcf\xe4\x91\xb5\x67\x3d\xac\xae\xcb\xfa\x64\x62\xa4\xbc\x37\x93\x9e\x17\x3a\xce\x73\x8b\xc7\x1d\xda\x8b\x25\x34\xcb\xae\x2d\xd8\x21\x6a\xe0\x52\xe2\x35\xe4\x72\x57\xa2\x67\x1f\x1f\x95\x08\xb8\x75\x51\xce\xd8\x35\x6d\xfa\x78\xa7\xc3\x92\x45\xc0\x95\xd0\xb8\xf9\x0b\x48\x6f\x24\x89\x90\xa5\x36\x12\x49\x51\x66\xec\x72\x98\xea\xae\x28\xf2\xe2\xb1\x6f\x69\xf3\x69\x05\x09\x95\x12\xc2\x23\x30\xf6\x7c\xf1\xe7\x9f\xe7\xa1\xda\x49\x6b\xa5\xbf\xfc\x30\x9e\x96\xe3\x85\x2b\x0f\xca\x83\x34\x1a\x7f\x0c\xaa\x77\x5c\x3c\x4f\x23\x15\xc2\x4c\xda\x93\x8a\xd6\x74\xdb\xb3\x6d\xd5\x3e\xad\x20\xc9\xb5\x9b\xf6\x7b\x25\x54\xf8\x68\x8c\xdc\xf2\x23\x7a\xb4\x2e\x21\x61\x1c\x4d\x57\x55\x65\x3d\x47\x5a\xbb\x69\x1c\x8d\x0d\x7a\xfc\xeb\x88\x09\x49\xd7\x2c\xfd\xf4\x76\x18\x3f\xa3\x55\x7b\x25\xb8\x9f\x2d\x9a\xb3\x30\x0f\x7c\xe0\xfa\x9b\x2f\xc4\x1d\xc7\x2e\xa0\xee\x3e\x76\xf3\xdf\xbf\x01\x00\x00\xff\xff\x16\x5c\xcd\xd1\x5e\x07\x00\x00")

func localesEnUsHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesEnUsHomeYml,
		"locales/en-US/home.yml",
	)
}

func localesEnUsHomeYml() (*asset, error) {
	bytes, err := localesEnUsHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/en-US/home.yml", size: 1886, mode: os.FileMode(420), modTime: time.Unix(1597052511, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesZhCnHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xd1\x4e\xe3\x46\x14\x7d\xf7\x57\x44\xe1\xb5\xea\x07\xf0\x36\xb5\x07\x70\x37\xb1\xad\xf1\xa4\x2d\x7d\x19\x65\x83\xd5\xd2\x05\x1b\x85\xf0\xd2\xa7\x06\x0a\x9b\xb0\x49\xa1\xbb\xa1\x85\x6e\x54\x92\x36\x88\x68\x77\x13\xb3\x50\x11\xc0\x40\x7f\xc6\x33\xe3\x3c\xf1\x0b\xd5\xd8\x89\x6b\x48\x51\x79\x1d\xdd\x7b\xce\xb9\xe7\xcc\xbd\x53\x05\x67\x79\xd9\xb1\x25\x0d\x64\x21\x81\x5f\xa9\x26\x36\xa7\x53\x69\xba\x5b\xe7\xc7\x27\x74\x70\x4a\x7b\xfb\xb4\xd9\x4d\x4b\xd2\x54\x61\x69\x6d\xb5\x64\x15\x25\x39\x93\x33\x31\x44\x44\x81\x19\x88\x21\x99\x01\x6a\x06\x2a\xd3\xa9\x34\xfb\xb5\xc5\xce\xf6\x68\xa5\x35\x3c\xe8\xd0\x9b\x37\xb4\x5a\xe7\xdb\xe7\xec\x87\x32\xff\xed\xc7\xe1\xdb\x2d\x7e\xdb\x09\x41\x8a\xd6\x82\x65\x97\x16\xf3\x4b\xd2\xbd\x7e\x82\xa0\xa9\xe7\x90\x0c\x05\x77\x04\xd1\xf9\x18\xfc\x75\x74\x77\x5d\x0e\xdc\x23\x7e\xbc\x3f\x7c\x7d\xe4\x5f\xbc\x62\xcd\x2a\xdd\x3c\x0b\xca\x0d\xff\xc2\x63\xcd\xab\xf4\x23\x20\xe4\x6b\x5d\x7b\x2a\x12\xdd\x71\x79\xa3\x4b\x6b\x21\xd8\x0c\xc8\x65\x30\x91\x11\x54\xa0\x86\x55\x90\x21\x32\xd0\x88\xa6\xe3\xd1\xb0\xd3\xa9\xf4\xd0\xdb\x0f\xfa\x1d\xfa\xb2\xc7\xea\x7d\xff\xa2\x1e\x6c\xdc\x44\x24\x77\xd7\x65\x31\xdf\xb7\xce\x6a\x49\x9a\xd3\x4d\x4c\x40\x06\x41\xa0\xcc\xff\x6b\x69\x24\x39\x61\xe9\x48\x7b\x58\x1d\xbb\x38\x29\x39\xee\xe3\xde\x4e\x24\x79\x6c\xe7\x24\x00\xf9\x6c\x9e\x18\x48\xff\x1c\xca\xf8\xa9\x58\xed\x4b\xfe\xb6\x1f\xaa\x37\xe7\x4d\x0c\xb3\x44\x35\xc2\x89\x67\xf4\x9c\x26\x04\x05\xee\x80\x6e\x56\x68\xfb\x3d\xdd\x3a\x60\xcd\x0f\xac\x79\xa5\x1a\xa3\x61\xd7\x56\xad\xa2\x64\x00\xd3\xfc\x52\x47\x4a\xd8\x94\x05\x58\x9e\x13\xcc\xee\x16\x6f\x95\x87\x8d\x83\xc0\x75\xd3\x92\x8e\xd4\x59\x55\x03\x99\xfb\x25\x3f\x1d\xde\xaf\xca\x99\x10\x11\xd5\x0c\x8b\x80\x8c\xd5\x2f\x84\xdd\xbc\xd1\x65\x95\x01\x6b\xbe\xa3\xbb\x22\xa6\x70\x8a\x41\x50\x6e\xf0\x33\x8f\xf7\xdb\x7c\x77\x8b\xfe\xbc\x1f\xaa\x09\xbb\x93\xba\xc5\x77\xec\x75\xa2\xfe\xb0\x02\x28\x59\x55\x7b\x2c\xd0\x54\x7e\x61\x79\xd1\x4e\x45\xe5\x51\xae\xc1\x1f\xef\x13\xd1\x26\xd5\x21\x98\x01\x18\x2a\x09\xa7\x47\x32\x4f\xdb\xf1\xb7\x8a\x7c\x7d\xc8\x9a\x33\x14\x10\xb2\x82\x09\x3a\xff\xef\x3e\x6b\x5c\x46\xce\x4a\x53\x45\xeb\x9b\x45\xc7\x1e\x27\x8c\xe0\xac\xaa\x6b\x4f\x5a\x15\x5a\xbb\xa2\x87\x87\xc9\x84\x13\x1f\x5c\x9a\xfa\xde\xb1\xad\x31\xaa\x58\x92\xa7\x61\x8e\x11\xee\x7d\x9c\x8d\x2e\xbf\x39\x0d\xfa\x6d\x5a\x79\x2d\x90\x9d\x15\xcb\x5e\x2d\xe5\x0b\x2f\xa4\x59\x88\xc7\x8a\x21\x42\x3a\x12\x61\x54\x6f\xfd\x8b\x3a\xad\x9c\x44\xf2\x44\xfd\x4a\xd1\xf9\xce\x2a\x94\xb2\xd6\xf2\x73\xab\x18\xdb\x0b\x14\x25\xb6\x33\x62\x63\x03\x8f\x6e\xb7\x12\x1d\xf1\x9a\x8e\xec\x7f\x2c\xd2\x28\x81\x89\x1d\x1d\x77\xcd\x01\x93\x8c\x4e\x99\x68\x09\x8b\x93\xeb\x75\x77\x5d\xfe\x8f\xfd\xb6\x9d\x05\x4b\xd2\x74\x05\xc6\xfb\x8d\x72\x9a\xa6\x6a\xb3\x04\x03\xf3\x99\x70\x6f\xf3\xdc\xf7\x7e\x09\xb6\xd7\xf9\xfa\x25\xdb\x3b\x19\xbe\xdc\x61\x6f\xea\xfe\x4d\x93\xf5\xfe\xa4\xcd\x2e\xab\x1e\x07\xed\xda\x27\xa9\xc0\x1d\xf0\x5e\x95\xde\x6e\xd2\xfe\x86\xef\x7d\x88\x9e\x69\xbf\xc6\xdc\xbd\x4f\x05\xcd\xf3\x7c\xe1\xc5\xda\x0a\x28\x14\x9c\x35\xbb\x24\x19\x00\x81\x2c\x81\x59\x03\xcf\x0b\x86\x9d\x75\xb6\x77\x22\xb4\x9d\x9f\xa5\x25\x31\xb8\x99\x33\x0c\x1d\xe1\xf0\xc6\xd4\x59\xc3\x65\x35\x71\x74\xf9\x47\x8f\xfe\xfe\x2a\x2d\xc9\x73\x50\x7e\x96\x38\xd3\xad\xf6\xf0\x5d\x2d\x8e\x37\x34\x76\x29\x6f\x3f\xb8\xa4\xff\x73\x47\x92\xe1\x4f\x5e\x13\xf1\x83\xff\x09\x00\x00\xff\xff\x04\x89\xf7\x52\x57\x06\x00\x00")

func localesZhCnHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesZhCnHomeYml,
		"locales/zh-CN/home.yml",
	)
}

func localesZhCnHomeYml() (*asset, error) {
	bytes, err := localesZhCnHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/zh-CN/home.yml", size: 1623, mode: os.FileMode(420), modTime: time.Unix(1597052511, 0)}
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
	"locales/en-US/home.yml": localesEnUsHomeYml,
	"locales/zh-CN/home.yml": localesZhCnHomeYml,
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
	"locales": &bintree{nil, map[string]*bintree{
		"en-US": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesEnUsHomeYml, map[string]*bintree{}},
		}},
		"zh-CN": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesZhCnHomeYml, map[string]*bintree{}},
		}},
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
