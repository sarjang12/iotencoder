// Code generated by go-bindata. DO NOT EDIT.
// sources:
// sql/20180525115614_create_device_table.down.sql (112B)
// sql/20180525115614_create_device_table.up.sql (493B)
// sql/20180526232618_add_streams_table.down.sql (37B)
// sql/20180526232618_add_streams_table.up.sql (354B)
// sql/20181202133704_add_operations.down.sql (45B)
// sql/20181202133704_add_operations.up.sql (50B)
// sql/20190306164350_remove_broker_col.down.sql (43B)
// sql/20190306164350_remove_broker_col.up.sql (39B)
// sql/20190306170548_add_certificate_table.down.sql (34B)
// sql/20190306170548_add_certificate_table.up.sql (106B)
// sql/20190308144957_rename_policy_id.down.sql (62B)
// sql/20190308144957_rename_policy_id.up.sql (62B)
// sql/20190315170620_add_uuid_column_to_stream.down.sql (39B)
// sql/20190315170620_add_uuid_column_to_stream.up.sql (125B)
// sql/20190315225536_change_stream_unique_index.down.sql (161B)
// sql/20190315225536_change_stream_unique_index.up.sql (163B)
// sql/20190512204433_add_device_label.down.sql (47B)
// sql/20190512204433_add_device_label.up.sql (71B)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __20180525115614_create_device_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x49\x2d\xcb\x4c\x4e\x2d\x56\x70\x76\x0c\x76\x76\x74\x71\xb5\xe6\xe2\x82\x28\x8b\x0c\x40\x56\x95\x5a\x51\x90\x5f\x5c\x5a\x94\x8a\xae\xcc\x35\x22\xc4\xd5\x2f\xd8\xd3\xdf\x0f\x49\x6d\x41\x7a\x72\x51\x65\x41\x49\xbe\x35\x20\x00\x00\xff\xff\xa4\x12\x3b\x91\x70\x00\x00\x00")

func _20180525115614_create_device_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20180525115614_create_device_tableDownSql,
		"20180525115614_create_device_table.down.sql",
	)
}

func _20180525115614_create_device_tableDownSql() (*asset, error) {
	bytes, err := _20180525115614_create_device_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20180525115614_create_device_table.down.sql", size: 112, mode: os.FileMode(420), modTime: time.Unix(1542929607, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3b, 0xff, 0x6c, 0x58, 0xd0, 0xed, 0x18, 0xff, 0x61, 0x91, 0x4d, 0x9, 0xd6, 0x88, 0xbb, 0xcd, 0xac, 0x24, 0x38, 0x5, 0xb4, 0xbc, 0x9c, 0x47, 0xd, 0x9f, 0x45, 0xb, 0x15, 0x44, 0xa1, 0x8e}}
	return a, nil
}

var __20180525115614_create_device_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xcd\x6e\xc2\x30\x10\x84\xef\x79\x8a\xb9\x01\x12\x4f\x00\x27\x17\x16\xd5\x6a\xe2\xa4\xb1\x2d\x42\x2f\x88\x62\x0b\x45\x54\x36\x32\x4e\x4b\xdf\xbe\x4a\x28\x69\xfa\x73\xe8\x6d\xec\x9d\xd9\xd5\x7c\x8b\x92\x98\x22\x50\xa5\x48\x48\x9e\x0b\xf0\x15\x44\xae\x40\x15\x97\x4a\xe2\x74\xd8\x87\xf7\x53\xf4\xf3\x24\xf9\x74\xaa\x4d\x41\xb0\x97\x93\x3f\x37\xc1\x82\x49\x90\xd0\x19\xc6\xa3\xc6\x1d\x9d\x7f\x73\xa3\x29\x46\xb5\x33\xde\x87\x56\xf9\x26\x76\x72\x32\xc8\xb3\xbb\x94\x7e\x5c\x31\xf6\xb5\xde\xdb\x33\xc6\x09\x50\x1b\x48\x2a\x39\x4b\x51\x94\x3c\x63\xe5\x06\x0f\xb4\x99\x26\xc0\x73\xf0\x47\x1b\xa0\xa8\x52\x5d\x56\xe8\x34\x6d\xff\xaf\xe1\x6d\xf4\x47\xeb\x7e\x4f\x5f\xbc\x3b\xd4\xb1\x31\x16\xcb\x5c\xb7\x97\x8b\x92\x16\xbc\x6b\xfa\xcd\xb6\x8b\xff\x70\xf5\xb5\x7b\x71\x9b\x62\x49\x2b\xa6\x53\x85\x9e\xc3\x6c\x76\x33\xb5\xc9\x7d\xb0\xbb\x68\xcd\x76\x17\xa1\x78\x46\x52\xb1\xac\xc0\x9a\xab\xfb\xee\x89\xa7\x5c\x50\xbf\x42\xe4\xeb\xf1\x24\x19\x20\xd3\x82\x3f\x6a\x02\x17\x4b\xaa\xfe\x26\x77\x6d\xbf\xad\xcd\x25\x01\x72\xf1\x05\x74\x08\x67\x32\xff\x08\x00\x00\xff\xff\x32\xbc\xcf\x9a\xed\x01\x00\x00")

func _20180525115614_create_device_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20180525115614_create_device_tableUpSql,
		"20180525115614_create_device_table.up.sql",
	)
}

func _20180525115614_create_device_tableUpSql() (*asset, error) {
	bytes, err := _20180525115614_create_device_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20180525115614_create_device_table.up.sql", size: 493, mode: os.FileMode(420), modTime: time.Unix(1542929607, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa5, 0x2b, 0x77, 0x99, 0x65, 0x10, 0xae, 0xfa, 0x52, 0x1f, 0x37, 0x2b, 0x4c, 0xc, 0x80, 0x1, 0x8c, 0x65, 0x8b, 0x6b, 0xf0, 0xd0, 0x1e, 0x9b, 0x65, 0xdf, 0xca, 0xbd, 0xd2, 0x9b, 0x1, 0xa}}
	return a, nil
}

var __20180526232618_add_streams_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2e\x29\x4a\x4d\xcc\x2d\x56\x70\x76\x0c\x76\x76\x74\x71\xb5\x06\x04\x00\x00\xff\xff\xa6\x34\x43\x4f\x25\x00\x00\x00")

func _20180526232618_add_streams_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20180526232618_add_streams_tableDownSql,
		"20180526232618_add_streams_table.down.sql",
	)
}

func _20180526232618_add_streams_tableDownSql() (*asset, error) {
	bytes, err := _20180526232618_add_streams_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20180526232618_add_streams_table.down.sql", size: 37, mode: os.FileMode(420), modTime: time.Unix(1542206352, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x76, 0x96, 0x77, 0xdd, 0x2, 0x53, 0x31, 0x1d, 0x8e, 0x44, 0x5b, 0x3f, 0x38, 0x8b, 0x5f, 0xed, 0x94, 0x30, 0x7a, 0x61, 0xe1, 0x1a, 0x55, 0x2, 0x76, 0x3d, 0xea, 0xf2, 0xb1, 0x75, 0xe7, 0x47}}
	return a, nil
}

var __20180526232618_add_streams_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xcd\x4e\xeb\x30\x10\x85\xf7\x7e\x8a\xb3\x4c\xa4\xbe\x41\x57\x6e\x3b\xb9\xd7\xc2\x71\x8a\x3d\x51\x13\x36\x56\x88\xbd\xb0\x5a\x28\x6a\x02\xa2\x6f\x8f\x12\x95\x00\x82\xe5\xcc\xf9\xd1\x7c\xb3\xb5\x24\x99\xc0\x72\xa3\x09\xaa\x80\xa9\x18\xd4\x28\xc7\x0e\xc3\x78\x89\xdd\xd3\x80\x4c\x00\x29\xc0\x91\x55\x52\x63\x6f\x55\x29\x6d\x8b\x3b\x6a\x57\x02\x08\xf1\x2d\xf5\xd1\xa7\x00\x65\x98\xfe\x91\x9d\x1b\x4c\xad\x35\x2c\x15\x64\xc9\x6c\xc9\xdd\x5c\x43\x96\x42\x3e\x85\x5e\x5e\x1f\x4f\xa9\xf7\xc7\x78\x05\x53\xc3\x4b\x64\xd6\xce\xa7\xd4\x5f\xa7\xc2\x5f\xd2\x78\x3e\xc6\x67\x6c\x5a\x26\xf9\x63\xdf\x5f\x62\x37\xc6\xe0\xbb\x11\xac\x4a\x72\x2c\xcb\x3d\x0e\x8a\xff\xcf\x23\x1e\x2a\x43\xd8\x51\x21\x6b\x3d\xf5\x1d\xb2\x5c\xe4\x6b\x21\x6e\xe4\xb5\x51\xf7\x35\x41\x99\x1d\x35\x7f\x3f\xc0\x2f\x8c\xfe\xeb\x70\x9f\xc2\xbb\x00\x2a\xf3\xe9\xca\x16\xd7\xea\x1b\x5f\xbe\x16\x1f\x01\x00\x00\xff\xff\x54\x2c\xaa\xbe\x62\x01\x00\x00")

func _20180526232618_add_streams_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20180526232618_add_streams_tableUpSql,
		"20180526232618_add_streams_table.up.sql",
	)
}

func _20180526232618_add_streams_tableUpSql() (*asset, error) {
	bytes, err := _20180526232618_add_streams_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20180526232618_add_streams_table.up.sql", size: 354, mode: os.FileMode(420), modTime: time.Unix(1542929607, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc8, 0x23, 0xef, 0x1d, 0x3, 0x67, 0xfa, 0x95, 0xba, 0xd7, 0xd6, 0xe0, 0x66, 0xb1, 0xcc, 0x11, 0x3d, 0xe, 0x73, 0x6c, 0x48, 0xde, 0xb1, 0x1f, 0xb1, 0x5, 0x58, 0x59, 0xae, 0x2e, 0xf0, 0xec}}
	return a, nil
}

var __20181202133704_add_operationsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x2e\x29\x4a\x4d\xcc\x2d\xe6\x52\x50\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\x2f\x48\x2d\x4a\x2c\xc9\xcc\xcf\x2b\xb6\x06\x04\x00\x00\xff\xff\x57\x1c\xaa\xf8\x2d\x00\x00\x00")

func _20181202133704_add_operationsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20181202133704_add_operationsDownSql,
		"20181202133704_add_operations.down.sql",
	)
}

func _20181202133704_add_operationsDownSql() (*asset, error) {
	bytes, err := _20181202133704_add_operationsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20181202133704_add_operations.down.sql", size: 45, mode: os.FileMode(420), modTime: time.Unix(1552671827, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xef, 0x17, 0xa5, 0xa6, 0x56, 0xf, 0xce, 0xd8, 0xfd, 0xdc, 0xfd, 0x79, 0xf4, 0x17, 0x51, 0x29, 0xb1, 0xcf, 0xb9, 0x55, 0xb1, 0x2e, 0x58, 0x16, 0x69, 0xe3, 0xa8, 0x21, 0xf5, 0xa0, 0xe6, 0x41}}
	return a, nil
}

var __20181202133704_add_operationsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x2e\x29\x4a\x4d\xcc\x2d\xe6\x52\x50\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\x2f\x48\x2d\x4a\x2c\xc9\xcc\xcf\x2b\x56\xf0\x0a\xf6\xf7\x73\xb2\x06\x04\x00\x00\xff\xff\x97\xbc\x02\xc2\x32\x00\x00\x00")

func _20181202133704_add_operationsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20181202133704_add_operationsUpSql,
		"20181202133704_add_operations.up.sql",
	)
}

func _20181202133704_add_operationsUpSql() (*asset, error) {
	bytes, err := _20181202133704_add_operationsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20181202133704_add_operations.up.sql", size: 50, mode: os.FileMode(420), modTime: time.Unix(1552671827, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0x81, 0x8c, 0xf7, 0x44, 0xa5, 0x90, 0xca, 0x30, 0x21, 0xb8, 0x6a, 0x65, 0xb6, 0x9, 0x89, 0x3f, 0x34, 0x99, 0x2b, 0xc5, 0x3a, 0xd3, 0x82, 0x2c, 0xae, 0xce, 0xb3, 0xf5, 0x28, 0x88, 0x16}}
	return a, nil
}

var __20190306164350_remove_broker_colDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x49\x2d\xcb\x4c\x4e\x2d\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x2a\xca\xcf\x4e\x2d\x52\x08\x71\x8d\x08\xb1\x06\x04\x00\x00\xff\xff\xb8\xa4\xe3\x27\x2b\x00\x00\x00")

func _20190306164350_remove_broker_colDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20190306164350_remove_broker_colDownSql,
		"20190306164350_remove_broker_col.down.sql",
	)
}

func _20190306164350_remove_broker_colDownSql() (*asset, error) {
	bytes, err := _20190306164350_remove_broker_colDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20190306164350_remove_broker_col.down.sql", size: 43, mode: os.FileMode(420), modTime: time.Unix(1552671827, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x73, 0x12, 0x9a, 0xd9, 0xd2, 0x13, 0x16, 0x44, 0x33, 0x66, 0x1d, 0xa6, 0xc6, 0x3c, 0xf2, 0x1c, 0x8d, 0xda, 0x6a, 0xa5, 0x22, 0x83, 0x0, 0xeb, 0xd0, 0x94, 0x7d, 0xef, 0xb4, 0x40, 0xfb, 0x20}}
	return a, nil
}

var __20190306164350_remove_broker_colUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x49\x2d\xcb\x4c\x4e\x2d\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x2a\xca\xcf\x4e\x2d\xb2\x06\x04\x00\x00\xff\xff\x42\x2d\xf7\x17\x27\x00\x00\x00")

func _20190306164350_remove_broker_colUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20190306164350_remove_broker_colUpSql,
		"20190306164350_remove_broker_col.up.sql",
	)
}

func _20190306164350_remove_broker_colUpSql() (*asset, error) {
	bytes, err := _20190306164350_remove_broker_colUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20190306164350_remove_broker_col.up.sql", size: 39, mode: os.FileMode(420), modTime: time.Unix(1552671827, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x27, 0xcf, 0x6e, 0x61, 0xcd, 0x94, 0xbb, 0x70, 0xd3, 0x57, 0xe, 0x83, 0xf3, 0xc2, 0xec, 0x4d, 0xc8, 0xd5, 0x19, 0x54, 0xfa, 0xa4, 0x57, 0x95, 0x8c, 0x59, 0xcf, 0x8d, 0xba, 0x43, 0xa8, 0x5a}}
	return a, nil
}

var __20190306170548_add_certificate_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x4e\x2d\x2a\xc9\x4c\xcb\x4c\x4e\x2c\x49\x2d\xb6\x06\x04\x00\x00\xff\xff\x9b\x6a\xf7\x60\x22\x00\x00\x00")

func _20190306170548_add_certificate_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20190306170548_add_certificate_tableDownSql,
		"20190306170548_add_certificate_table.down.sql",
	)
}

func _20190306170548_add_certificate_tableDownSql() (*asset, error) {
	bytes, err := _20190306170548_add_certificate_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20190306170548_add_certificate_table.down.sql", size: 34, mode: os.FileMode(420), modTime: time.Unix(1552671827, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3d, 0x85, 0xef, 0x15, 0xa1, 0x51, 0x74, 0x22, 0x6b, 0x2f, 0xde, 0x28, 0x99, 0xb5, 0x60, 0xd6, 0xe8, 0x10, 0x23, 0xa7, 0x48, 0x63, 0xf2, 0xc4, 0x3c, 0xca, 0x83, 0x1f, 0xb4, 0x65, 0xad, 0x98}}
	return a, nil
}

var __20190306170548_add_certificate_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\xf0\xf3\x0f\x51\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x4e\x2d\x2a\xc9\x4c\xcb\x4c\x4e\x2c\x49\x2d\x56\xd0\xe0\x52\x50\xc8\x4e\xad\x54\x08\x71\x8d\x08\x01\x2b\xf2\x0b\xf5\xf1\x51\x08\x08\xf2\xf4\x75\x0c\x8a\x54\xf0\x76\x8d\xd4\xe1\x52\x40\xd6\xa1\xe0\x14\x19\xe2\xea\x08\x57\xc9\xa5\x69\x0d\x08\x00\x00\xff\xff\x2d\x4d\xb2\x71\x6a\x00\x00\x00")

func _20190306170548_add_certificate_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20190306170548_add_certificate_tableUpSql,
		"20190306170548_add_certificate_table.up.sql",
	)
}

func _20190306170548_add_certificate_tableUpSql() (*asset, error) {
	bytes, err := _20190306170548_add_certificate_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20190306170548_add_certificate_table.up.sql", size: 106, mode: os.FileMode(420), modTime: time.Unix(1552671827, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x66, 0x3c, 0x3, 0x6a, 0x8c, 0x5a, 0x0, 0xe2, 0xca, 0x24, 0x4b, 0xf0, 0x4b, 0x55, 0xb2, 0xc4, 0x3f, 0x19, 0x75, 0x20, 0x4f, 0xd3, 0x4d, 0xc6, 0xa6, 0x9b, 0xbb, 0xc1, 0x94, 0x70, 0xbc, 0x38}}
	return a, nil
}

var __20190308144957_rename_policy_idDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x2e\x29\x4a\x4d\xcc\x2d\xe6\x52\x50\x08\x72\xf5\x73\xf4\x75\x55\x70\xf6\xf7\x09\xf5\xf5\x53\x48\xce\xcf\xcd\x2d\xcd\xcb\x2c\xa9\x8c\xcf\x4c\x51\x08\xf1\x57\x28\xc8\xcf\xc9\x4c\x06\x71\xac\x01\x01\x00\x00\xff\xff\xe7\x3c\x58\x88\x3e\x00\x00\x00")

func _20190308144957_rename_policy_idDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20190308144957_rename_policy_idDownSql,
		"20190308144957_rename_policy_id.down.sql",
	)
}

func _20190308144957_rename_policy_idDownSql() (*asset, error) {
	bytes, err := _20190308144957_rename_policy_idDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20190308144957_rename_policy_id.down.sql", size: 62, mode: os.FileMode(420), modTime: time.Unix(1552671827, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x41, 0x9f, 0xd1, 0x62, 0xa4, 0x15, 0x9e, 0x20, 0x98, 0xca, 0x4f, 0x1c, 0xd9, 0xe4, 0xe7, 0xe3, 0x30, 0xc1, 0xc6, 0xc8, 0x9f, 0xd7, 0x6d, 0xce, 0x36, 0xfe, 0xa7, 0xa5, 0xa7, 0x59, 0xf7, 0x6f}}
	return a, nil
}

var __20190308144957_rename_policy_idUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x2e\x29\x4a\x4d\xcc\x2d\xe6\x52\x50\x08\x72\xf5\x73\xf4\x75\x55\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xc8\xcf\xc9\x4c\xae\x8c\xcf\x4c\x51\x08\xf1\x57\x48\xce\xcf\xcd\x2d\xcd\xcb\x2c\x01\xf1\xad\x01\x01\x00\x00\xff\xff\x69\x65\xa3\xeb\x3e\x00\x00\x00")

func _20190308144957_rename_policy_idUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20190308144957_rename_policy_idUpSql,
		"20190308144957_rename_policy_id.up.sql",
	)
}

func _20190308144957_rename_policy_idUpSql() (*asset, error) {
	bytes, err := _20190308144957_rename_policy_idUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20190308144957_rename_policy_id.up.sql", size: 62, mode: os.FileMode(420), modTime: time.Unix(1552671827, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcc, 0x5b, 0x7d, 0xd, 0x38, 0x77, 0xf, 0xcd, 0x16, 0x40, 0xd8, 0x41, 0xc2, 0x4b, 0x7b, 0x81, 0x98, 0xd3, 0xb6, 0x5c, 0x1d, 0x87, 0xdb, 0x42, 0x76, 0x2b, 0xe6, 0x7c, 0xc8, 0x39, 0x2, 0x85}}
	return a, nil
}

var __20190315170620_add_uuid_column_to_streamDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x2e\x29\x4a\x4d\xcc\x2d\xe6\x52\x50\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\x2d\xcd\x4c\xb1\x06\x04\x00\x00\xff\xff\x98\x01\x3c\xa4\x27\x00\x00\x00")

func _20190315170620_add_uuid_column_to_streamDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20190315170620_add_uuid_column_to_streamDownSql,
		"20190315170620_add_uuid_column_to_stream.down.sql",
	)
}

func _20190315170620_add_uuid_column_to_streamDownSql() (*asset, error) {
	bytes, err := _20190315170620_add_uuid_column_to_streamDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20190315170620_add_uuid_column_to_stream.down.sql", size: 39, mode: os.FileMode(420), modTime: time.Unix(1552671827, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xee, 0x5c, 0x86, 0x4a, 0x1b, 0x17, 0xf9, 0xa1, 0x4, 0xc7, 0x18, 0x12, 0xf7, 0x4f, 0xc3, 0x45, 0x6f, 0xc3, 0x64, 0xd3, 0x16, 0x1c, 0x9a, 0x61, 0x66, 0xd3, 0x64, 0x55, 0xbb, 0xc, 0xdc, 0xf1}}
	return a, nil
}

var __20190315170620_add_uuid_column_to_streamUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcc\xb1\x0a\xc2\x50\x0c\x85\xe1\xfd\x3e\xc5\x19\xf5\x19\x3a\xc5\x26\x42\x20\xe6\x62\x9b\x40\xb7\x22\xd4\xa1\x83\x8b\xb5\xe0\xe3\x4b\x85\xbb\x7e\x9c\xf3\x93\x85\x0c\x08\xba\x98\x60\xfb\xbc\x9f\x8f\xd7\x56\x00\x62\x46\x5f\x2d\x6f\x8e\x7d\x5f\x17\x64\x2a\xc3\x6b\xc0\xd3\xac\x2b\xa5\x1f\x84\x42\x90\xae\xf7\x14\xa8\xb3\x4c\xd0\xeb\x7f\x21\x93\x8e\x31\xb6\xd6\x7c\xdc\xe7\x75\xf9\x16\xa0\x7a\x53\x9c\x0e\x3e\x77\xbf\x00\x00\x00\xff\xff\x8e\x65\xf6\x21\x7d\x00\x00\x00")

func _20190315170620_add_uuid_column_to_streamUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20190315170620_add_uuid_column_to_streamUpSql,
		"20190315170620_add_uuid_column_to_stream.up.sql",
	)
}

func _20190315170620_add_uuid_column_to_streamUpSql() (*asset, error) {
	bytes, err := _20190315170620_add_uuid_column_to_streamUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20190315170620_add_uuid_column_to_stream.up.sql", size: 125, mode: os.FileMode(420), modTime: time.Unix(1552671827, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x63, 0xe8, 0xa6, 0xbe, 0x10, 0xc0, 0x22, 0x47, 0x38, 0x51, 0x1a, 0x88, 0x1e, 0x34, 0x16, 0xa7, 0xf, 0x0, 0x7f, 0xb0, 0xd1, 0x7f, 0xc5, 0x90, 0xec, 0x9f, 0x38, 0xd4, 0x9b, 0xfa, 0xf8, 0x37}}
	return a, nil
}

var __20190315225536_change_stream_unique_indexDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\xf0\xf4\x73\x71\x8d\x50\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2e\x29\x4a\x4d\xcc\x2d\x8e\x4f\x49\x2d\xcb\x4c\x4e\x8d\xcf\x4c\x89\x4f\xce\xcf\xcd\x2d\xcd\xcb\x2c\xa9\x04\x71\x32\x53\x2a\xac\xb9\xb8\x9c\x83\x5c\x1d\x43\x5c\x15\x42\xfd\x3c\x03\x43\x5d\x11\x46\xf8\xf9\x87\xe0\x36\xa6\xa0\x34\x29\x27\x33\x39\x3e\x3b\x15\x64\x4e\x05\x97\x82\x82\xbf\x1f\x4c\x95\x06\x5c\x95\x8e\x02\x42\x99\xa6\x35\x20\x00\x00\xff\xff\xbf\xf2\x66\xc2\xa1\x00\x00\x00")

func _20190315225536_change_stream_unique_indexDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20190315225536_change_stream_unique_indexDownSql,
		"20190315225536_change_stream_unique_index.down.sql",
	)
}

func _20190315225536_change_stream_unique_indexDownSql() (*asset, error) {
	bytes, err := _20190315225536_change_stream_unique_indexDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20190315225536_change_stream_unique_index.down.sql", size: 161, mode: os.FileMode(420), modTime: time.Unix(1552691070, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x61, 0xed, 0xd4, 0x8a, 0xa3, 0x5d, 0xcc, 0x12, 0x7f, 0x4, 0xe0, 0x25, 0x98, 0xb, 0x9a, 0x9f, 0x1d, 0xe2, 0xb, 0x43, 0x18, 0x9f, 0x92, 0xc1, 0xb8, 0x2c, 0x34, 0x1d, 0xa7, 0x51, 0xef, 0xb}}
	return a, nil
}

var __20190315225536_change_stream_unique_indexUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\xf0\xf4\x73\x71\x8d\x50\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2e\x29\x4a\x4d\xcc\x2d\x8e\x4f\x49\x2d\xcb\x4c\x4e\x8d\xcf\x4c\x89\x2f\x28\x4d\xca\xc9\x4c\x8e\xcf\x4e\xad\x8c\xcf\x4c\xa9\xb0\xe6\xe2\x72\x0e\x72\x75\x0c\x71\x55\x08\xf5\xf3\x0c\x0c\x75\x45\x18\xe0\xe7\x1f\x82\xdb\x90\xe4\xfc\xdc\xdc\xd2\xbc\xcc\x12\x90\x19\x20\x63\xb8\x14\x14\xfc\xfd\x60\xea\x34\xe0\xea\x74\x14\x90\x15\x6a\x5a\x03\x02\x00\x00\xff\xff\x7c\xaf\x0f\xbc\xa3\x00\x00\x00")

func _20190315225536_change_stream_unique_indexUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20190315225536_change_stream_unique_indexUpSql,
		"20190315225536_change_stream_unique_index.up.sql",
	)
}

func _20190315225536_change_stream_unique_indexUpSql() (*asset, error) {
	bytes, err := _20190315225536_change_stream_unique_indexUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20190315225536_change_stream_unique_index.up.sql", size: 163, mode: os.FileMode(420), modTime: time.Unix(1552691070, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb2, 0x17, 0x4d, 0xb5, 0x68, 0x88, 0x65, 0x74, 0x3f, 0x57, 0xaf, 0xc4, 0x5a, 0x8, 0x2b, 0x43, 0x14, 0xb3, 0xfc, 0x4e, 0x22, 0xc, 0xb9, 0x77, 0x2, 0x2d, 0x54, 0x47, 0xfd, 0x96, 0x3e, 0xe1}}
	return a, nil
}

var __20190512204433_add_device_labelDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x49\x2d\xcb\x4c\x4e\x2d\xe6\x52\x50\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x83\x8a\xc6\xe7\x24\x26\xa5\xe6\x58\x03\x02\x00\x00\xff\xff\x8c\xd1\x34\xbd\x2f\x00\x00\x00")

func _20190512204433_add_device_labelDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20190512204433_add_device_labelDownSql,
		"20190512204433_add_device_label.down.sql",
	)
}

func _20190512204433_add_device_labelDownSql() (*asset, error) {
	bytes, err := _20190512204433_add_device_labelDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20190512204433_add_device_label.down.sql", size: 47, mode: os.FileMode(420), modTime: time.Unix(1557703783, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x81, 0x81, 0x1e, 0x1e, 0x3, 0xd6, 0x5a, 0xed, 0x46, 0xa4, 0xd, 0xdf, 0x7, 0x1e, 0xd9, 0xf9, 0x34, 0xfb, 0x13, 0x79, 0x53, 0x55, 0x41, 0x6f, 0xdc, 0x7b, 0xd3, 0x8e, 0xe0, 0xc2, 0xc7, 0x53}}
	return a, nil
}

var __20190512204433_add_device_labelUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x49\x2d\xcb\x4c\x4e\x2d\xe6\x52\x50\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x83\x0a\xc6\xe7\x24\x26\xa5\xe6\x28\x84\xb8\x46\x84\x28\xf8\xf9\x87\x28\xf8\x85\xfa\xf8\x28\xb8\xb8\xba\x39\x86\xfa\x84\x28\xa8\xab\x5b\x03\x02\x00\x00\xff\xff\x04\xb5\x14\x14\x47\x00\x00\x00")

func _20190512204433_add_device_labelUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20190512204433_add_device_labelUpSql,
		"20190512204433_add_device_label.up.sql",
	)
}

func _20190512204433_add_device_labelUpSql() (*asset, error) {
	bytes, err := _20190512204433_add_device_labelUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20190512204433_add_device_label.up.sql", size: 71, mode: os.FileMode(420), modTime: time.Unix(1557703783, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xef, 0x4, 0x75, 0x44, 0x35, 0xa, 0x4b, 0x83, 0x71, 0x8e, 0xc7, 0x72, 0x20, 0x59, 0x4, 0x67, 0x22, 0x44, 0x11, 0xce, 0xf, 0x52, 0xb2, 0x40, 0xf6, 0x93, 0xc6, 0xe, 0x90, 0xdd, 0x9e, 0x5d}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"20180525115614_create_device_table.down.sql": _20180525115614_create_device_tableDownSql,

	"20180525115614_create_device_table.up.sql": _20180525115614_create_device_tableUpSql,

	"20180526232618_add_streams_table.down.sql": _20180526232618_add_streams_tableDownSql,

	"20180526232618_add_streams_table.up.sql": _20180526232618_add_streams_tableUpSql,

	"20181202133704_add_operations.down.sql": _20181202133704_add_operationsDownSql,

	"20181202133704_add_operations.up.sql": _20181202133704_add_operationsUpSql,

	"20190306164350_remove_broker_col.down.sql": _20190306164350_remove_broker_colDownSql,

	"20190306164350_remove_broker_col.up.sql": _20190306164350_remove_broker_colUpSql,

	"20190306170548_add_certificate_table.down.sql": _20190306170548_add_certificate_tableDownSql,

	"20190306170548_add_certificate_table.up.sql": _20190306170548_add_certificate_tableUpSql,

	"20190308144957_rename_policy_id.down.sql": _20190308144957_rename_policy_idDownSql,

	"20190308144957_rename_policy_id.up.sql": _20190308144957_rename_policy_idUpSql,

	"20190315170620_add_uuid_column_to_stream.down.sql": _20190315170620_add_uuid_column_to_streamDownSql,

	"20190315170620_add_uuid_column_to_stream.up.sql": _20190315170620_add_uuid_column_to_streamUpSql,

	"20190315225536_change_stream_unique_index.down.sql": _20190315225536_change_stream_unique_indexDownSql,

	"20190315225536_change_stream_unique_index.up.sql": _20190315225536_change_stream_unique_indexUpSql,

	"20190512204433_add_device_label.down.sql": _20190512204433_add_device_labelDownSql,

	"20190512204433_add_device_label.up.sql": _20190512204433_add_device_labelUpSql,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"20180525115614_create_device_table.down.sql":        &bintree{_20180525115614_create_device_tableDownSql, map[string]*bintree{}},
	"20180525115614_create_device_table.up.sql":          &bintree{_20180525115614_create_device_tableUpSql, map[string]*bintree{}},
	"20180526232618_add_streams_table.down.sql":          &bintree{_20180526232618_add_streams_tableDownSql, map[string]*bintree{}},
	"20180526232618_add_streams_table.up.sql":            &bintree{_20180526232618_add_streams_tableUpSql, map[string]*bintree{}},
	"20181202133704_add_operations.down.sql":             &bintree{_20181202133704_add_operationsDownSql, map[string]*bintree{}},
	"20181202133704_add_operations.up.sql":               &bintree{_20181202133704_add_operationsUpSql, map[string]*bintree{}},
	"20190306164350_remove_broker_col.down.sql":          &bintree{_20190306164350_remove_broker_colDownSql, map[string]*bintree{}},
	"20190306164350_remove_broker_col.up.sql":            &bintree{_20190306164350_remove_broker_colUpSql, map[string]*bintree{}},
	"20190306170548_add_certificate_table.down.sql":      &bintree{_20190306170548_add_certificate_tableDownSql, map[string]*bintree{}},
	"20190306170548_add_certificate_table.up.sql":        &bintree{_20190306170548_add_certificate_tableUpSql, map[string]*bintree{}},
	"20190308144957_rename_policy_id.down.sql":           &bintree{_20190308144957_rename_policy_idDownSql, map[string]*bintree{}},
	"20190308144957_rename_policy_id.up.sql":             &bintree{_20190308144957_rename_policy_idUpSql, map[string]*bintree{}},
	"20190315170620_add_uuid_column_to_stream.down.sql":  &bintree{_20190315170620_add_uuid_column_to_streamDownSql, map[string]*bintree{}},
	"20190315170620_add_uuid_column_to_stream.up.sql":    &bintree{_20190315170620_add_uuid_column_to_streamUpSql, map[string]*bintree{}},
	"20190315225536_change_stream_unique_index.down.sql": &bintree{_20190315225536_change_stream_unique_indexDownSql, map[string]*bintree{}},
	"20190315225536_change_stream_unique_index.up.sql":   &bintree{_20190315225536_change_stream_unique_indexUpSql, map[string]*bintree{}},
	"20190512204433_add_device_label.down.sql":           &bintree{_20190512204433_add_device_labelDownSql, map[string]*bintree{}},
	"20190512204433_add_device_label.up.sql":             &bintree{_20190512204433_add_device_labelUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
