package file

import (
	"strings"
)

type File struct {
	path   string
	name   string
	format string
}

func (file *File) ParseFull(s string) *File {
	arr := strings.Split(s, "/")
	length := len(arr)
	if length == 1 {
		file.Parse(arr[0])
	} else if length > 1 {
		file.path = strings.Join(arr[0:length-1], "/")
		file.Parse(arr[length-1])
	}
	return file
}
func (file *File) Parse(s string) (*File) {
	arr := strings.Split(s, ".")
	length := len(arr)
	if length == 1 {
		file.SetName(arr[0])
	} else if length > 1 {
		file.SetName(strings.Join(arr[0:length-1], "."))
		file.SetFormat(arr[length-1])
	}
	return file
}

func (file *File) Path() string {
	return file.path
}
func (file *File) Name() string {
	return file.name
}
func (file *File) FullName() string {
	return file.name + "." + file.format
}
func (file *File) Format() string {
	return file.format
}

func (file *File) SetName(name string) {
	file.name = name
}

func (file *File) SetPath(path string) {
	file.path = path
}
func (file *File) SetFormat(format string) {
	file.format = format
}

func (file *File) StringFull() string {
	return file.path + "/" + file.FullName()
}
