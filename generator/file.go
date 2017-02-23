// Copyright 2017, TCN Inc.
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:

//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of TCN Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"

	"golang.org/x/tools/imports"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type FileStruct struct {
	*bytes.Buffer
	Name             string
	ProtoFileName    string
	PackageName      string
	OriginalFileDesc *descriptor.FileDescriptorProto
	ErrorList        string
	ImportList       []*Import
}

func NewFileStruct(proto *descriptor.FileDescriptorProto) *FileStruct {
	ret := new(FileStruct)
	ret.Name = strings.Replace(proto.GetName(), ".proto", ".persist.go", -1)
	ret.ProtoFileName = proto.GetName()
	ret.PackageName = proto.GetPackage()
	ret.OriginalFileDesc = proto
	ret.ErrorList = ""
	ret.Buffer = new(bytes.Buffer)
	ret.ImportList = []*Import{
		&Import{ProtoFileName: "", ProtoPackageName: "", GoPackageName: "fmt", GoImportPath: "fmt"},
		&Import{ProtoFileName: "", ProtoPackageName: "", GoPackageName: "sql", GoImportPath: "database/sql"},
		&Import{ProtoFileName: "", ProtoPackageName: "", GoPackageName: "driver", GoImportPath: "database/sql/driver"},
		&Import{ProtoFileName: "", ProtoPackageName: "", GoPackageName: "jsonpb", GoImportPath: "github.com/golang/protobuf/jsonpb"},
	}

	return ret
}

func (f *FileStruct) AddImport(goPkg, goPath string, file *descriptor.FileDescriptorProto) string {
	for _, im := range f.ImportList {
		if im.GoImportPath == goPath {
			return im.GoPackageName
		}
	}
	i := &Import{
		GoImportPath: goPath,
		GoPackageName: func() string {
			for _, im := range f.ImportList {
				if im.GoPackageName == goPkg {
					return "_" + goPkg
				}
			}
			return goPkg
		}(),
		ProtoFileName:    file.GetName(),
		ProtoPackageName: file.GetPackage(),
	}
	f.ImportList = append(f.ImportList, i)
	return i.GoPackageName
}

func (f *FileStruct) AddHeader() string {
	return fmt.Sprintf(`// This file is generated by protoc-gen-persist 
// Source File: %s 
// DO NOT EDIT !

`, f.OriginalFileDesc.GetName())
}

func (f *FileStruct) AddPackage() string {
	return fmt.Sprintf("package %s\n", f.GetPackage())
}

func (f *FileStruct) AddImports() string {
	ret := "import (\n"
	for _, im := range f.ImportList {
		ret += im.GetImportString() + "\n"
	}
	ret += ")\n"
	return ret

}

func (f *FileStruct) GetContent() string {
	var buffer *bytes.Buffer = new(bytes.Buffer)
	buffer.WriteString(f.AddHeader())
	buffer.WriteString(f.AddPackage())
	buffer.WriteString(f.AddImports())
	buffer.Write(f.Buffer.Bytes())

	buf, err := imports.Process(f.Name, buffer.Bytes(), &imports.Options{
		TabWidth:  8,
		TabIndent: true,
		Comments:  true,
		Fragment:  true,
	})
	if err != nil {
		// should I panic here ?
		logrus.WithError(err).WithField("content", buffer).Error("Error on optimizing source code !")
		buf, err = format.Source(f.Buffer.Bytes())
		if err != nil {
			logrus.WithError(err).WithField("content", buffer).Error("Error formatting source")
			return string(buf)
		}
		return string(buf)

	}
	return string(buf)
}

func (f *FileStruct) GetPackage() string {
	if f.OriginalFileDesc.Options != nil && f.OriginalFileDesc.GetOptions().GetGoPackage() != "" {
		vals := strings.Split(f.OriginalFileDesc.GetOptions().GetGoPackage(), ";")
		return vals[len(vals)-1]
	} else {
		return f.PackageName[strings.LastIndex(f.PackageName, ".")+1 : len(f.PackageName)]
	}
}

func (f *FileStruct) Error(msg string) {
	f.ErrorList = f.ErrorList + "\n" + msg
}

func (f *FileStruct) P(str ...interface{}) {
	for _, v := range str {
		switch s := v.(type) {
		case string:
			f.WriteString(s)
		case *string:
			f.WriteString(*s)
		case bool:
			fmt.Fprintf(f, "%t", s)
		case *bool:
			fmt.Fprintf(f, "%t", *s)
		case int:
			fmt.Fprintf(f, "%d", s)
		case *int32:
			fmt.Fprintf(f, "%d", *s)
		case *int64:
			fmt.Fprintf(f, "%d", *s)
		case float64:
			fmt.Fprintf(f, "%g", s)
		case *float64:
			fmt.Fprintf(f, "%g", *s)
		default:
			f.Error(fmt.Sprintf("unknown type in printer: %T", v))
		}
	}
	f.WriteByte('\n')
}

type FileList struct {
	List []*FileStruct
}

func NewFileList() *FileList {
	ret := new(FileList)
	return ret
}

func (fl *FileList) ProtoFileExist(fileName string) bool {
	for _, f := range fl.List {
		if f.ProtoFileName == fileName {
			return true
		}
	}
	return false
}

func (fl *FileList) GetFileByName(protoFileName string) *FileStruct {
	for _, f := range fl.List {
		if f.ProtoFileName == protoFileName {
			return f
		}
	}
	return nil
}

func (fl *FileList) GetFileByDesc(file *descriptor.FileDescriptorProto) *FileStruct {
	for _, f := range fl.List {
		if f.ProtoFileName == file.GetName() {
			return f
		}
	}
	return nil
}

func (fl *FileList) NewOrGetFile(proto *descriptor.FileDescriptorProto) *FileStruct {
	if f := fl.GetFileByDesc(proto); f != nil {
		return f
	}

	ret := NewFileStruct(proto)
	fl.List = append(fl.List, ret)

	return ret
}
