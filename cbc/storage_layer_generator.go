package cbc

import(
	"os"
	"fmt"
	"strings"
	"path"
	tpl "text/template"
)

type StorageLayerGenerator struct{
	ServiceGenerator *ServiceGenerator
}

func NewStorageLayerGenerator(s *ServiceGenerator) *StorageLayerGenerator{
	return &StorageLayerGenerator{
		ServiceGenerator:s,
	}
}


func (this *StorageLayerGenerator) Generate() error{
	tplIns,err := tpl.New("storage_layer.tmpl").ParseFS(templates,"tpls/storage_layer.tmpl")
	if err != nil {
		return err
	}

	err = os.MkdirAll(this.ReleativeTargetDirPath(),os.ModePerm)
	if err != nil{
		return err
	}

	generatedFile,err := os.Create(path.Join(this.ReleativeTargetDirPath(),fmt.Sprintf("%s.go",strings.ToLower(this.ServiceGenerator.TargetNode.Name()))))
	if err != nil {
		return err
	}
	defer generatedFile.Close()

	return tplIns.Execute(generatedFile,this)
}

func (this *StorageLayerGenerator) ReleativeTargetDirPath() string{
	return "internal/data"
}

func (this *StorageLayerGenerator) Generated_PkgPath() string{
	return path.Join(this.ServiceGenerator.Adaptor.Core.Module.Path,this.ReleativeTargetDirPath())
}

func (this *StorageLayerGenerator) Generated_PkgName() string{
	return "data"
}


func (this *StorageLayerGenerator) Imports() []string{
	return []string{
		`"context"`,
	}
}

func (this *StorageLayerGenerator) Generated_StructName() string{
	return fmt.Sprintf("%sStorage",this.ServiceGenerator.TargetNode.Name())
}

func (this *StorageLayerGenerator) Generated_CreateFuncName() string{
	return "Create"
}

