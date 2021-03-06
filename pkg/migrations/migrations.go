// Code generated by go-bindata.
// sources:
// ../../migrations/20180423103837-images_table_added.sql
// ../../migrations/20180426193723-resize_jobs_table_added.sql
// DO NOT EDIT!

package migrations

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

var _migrations20180423103837Images_table_addedSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xce\x31\x6e\x83\x30\x18\x47\xf1\xdd\xa7\xf8\x8f\xa0\x96\x13\x30\xb9\xc5\x55\x51\x8d\x8d\x2c\x5b\x88\x2e\xd1\x17\xb0\x88\x07\x27\xc8\x58\x41\xb9\x7d\xa4\x64\x62\x7d\x6f\xf9\xb1\xaa\xc2\x47\x0c\x4b\xa2\xec\xe1\x56\xf6\x6d\x04\xb7\x02\x96\x7f\x49\x81\x10\x69\xf1\x1b\x0a\x06\x00\x61\xc6\x9d\xd2\x74\xa1\x84\xde\xb4\x1d\x37\x23\xfe\xc4\xf8\xf9\x7a\x89\x76\x9c\x1f\xd9\x13\x94\xb6\x50\x4e\xca\x77\x9f\x92\xa7\xec\xe7\x13\x65\xe4\x10\xfd\x96\x29\xae\x18\x5a\xfb\xab\x9d\x85\x6d\x3b\x81\x7f\xad\x04\x1a\xf1\xc3\x9d\xb4\x50\x7a\x28\x4a\x56\xd6\xec\xa0\x6a\x6e\xfb\x95\x35\x46\xf7\x07\x55\xcd\x9e\x01\x00\x00\xff\xff\x40\xb3\x99\xf3\xbb\x00\x00\x00")

func migrations20180423103837Images_table_addedSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20180423103837Images_table_addedSql,
		"migrations/20180423103837-images_table_added.sql",
	)
}

func migrations20180423103837Images_table_addedSql() (*asset, error) {
	bytes, err := migrations20180423103837Images_table_addedSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20180423103837-images_table_added.sql", size: 187, mode: os.FileMode(420), modTime: time.Unix(1524754374, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20180426193723Resize_jobs_table_addedSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd0\xd1\x4e\xc2\x30\x18\x05\xe0\xfb\x3e\xc5\xb9\x84\x08\x4f\xc0\x55\xdd\x7e\xe2\xe2\x68\xb1\xb4\x21\x78\x43\x0a\xfb\xb3\xd5\x38\x20\x6d\x75\xd1\xa7\x37\x0b\x6a\x34\x7a\xd9\x7c\x39\xe9\xf9\x8f\x98\xcf\x71\xd3\x87\x36\xfa\xcc\x70\x17\x51\x18\x92\x96\x60\xe5\x6d\x4d\x88\x9c\xc2\x3b\xef\x9f\xce\x87\x84\x89\x00\x42\x83\x43\x68\x13\xc7\xe0\x9f\xb1\x36\xd5\x4a\x9a\x1d\xee\x69\x37\x1b\xad\xf7\x2d\xef\x43\x83\x57\x1f\x8f\x9d\x8f\x30\xb4\x24\x43\xaa\xa0\xcd\xd5\x12\xb4\x42\x49\x35\x59\x42\x21\x37\x85\x2c\x69\xcc\x0d\xa1\xc9\x1d\xc2\x29\x73\xcb\x11\x4a\x5b\x28\x57\xd7\xa3\x74\x1c\xda\x2e\xff\x4b\x29\xfb\xfc\x92\xbe\xbf\xfa\x49\xd1\x0f\x38\xbc\x65\xf6\xe3\xe3\x18\xd9\x67\x6e\xf6\x3e\x23\x87\x9e\x53\xf6\xfd\x05\xdb\xca\xde\x69\x67\x61\xab\x15\xe1\x51\x2b\x42\x49\x4b\xe9\x6a\x0b\xa5\xb7\x93\xe9\x98\x73\xaa\x7a\x70\x84\xc9\xd7\x51\xb3\x6b\xcd\xd9\x67\xa7\xa9\x98\x2e\x7e\x0d\x57\x9e\x87\x93\x28\x8d\x5e\xff\x1d\x6e\x21\x3e\x02\x00\x00\xff\xff\xb1\x93\x0d\xe5\x63\x01\x00\x00")

func migrations20180426193723Resize_jobs_table_addedSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20180426193723Resize_jobs_table_addedSql,
		"migrations/20180426193723-resize_jobs_table_added.sql",
	)
}

func migrations20180426193723Resize_jobs_table_addedSql() (*asset, error) {
	bytes, err := migrations20180426193723Resize_jobs_table_addedSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20180426193723-resize_jobs_table_added.sql", size: 355, mode: os.FileMode(420), modTime: time.Unix(1524761364, 0)}
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
	"migrations/20180423103837-images_table_added.sql":      migrations20180423103837Images_table_addedSql,
	"migrations/20180426193723-resize_jobs_table_added.sql": migrations20180426193723Resize_jobs_table_addedSql,
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
	"migrations": &bintree{nil, map[string]*bintree{
		"20180423103837-images_table_added.sql":      &bintree{migrations20180423103837Images_table_addedSql, map[string]*bintree{}},
		"20180426193723-resize_jobs_table_added.sql": &bintree{migrations20180426193723Resize_jobs_table_addedSql, map[string]*bintree{}},
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
