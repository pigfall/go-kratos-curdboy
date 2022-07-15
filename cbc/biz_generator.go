package cbc

import(
	"strings"
	"os"
	tpl "text/template"
	"path"
	"fmt"
)

type BizGenerator struct{
	ServiceGenerator *ServiceGenerator
}

func NewBizGenerator(svcGenerator *ServiceGenerator)*BizGenerator{
	return &BizGenerator{
		ServiceGenerator:svcGenerator,
	}
}

func (this *BizGenerator) Generate()error{
	tplIns,err := tpl.New("biz.tmpl").ParseFS(templates,"tpls/biz.tmpl")
	if err != nil {
		return err
	}

	err = os.MkdirAll(this.ReleativeTargetDirPath(),os.ModePerm)
	if err != nil {
		return err
	}

	generatedFile,err := os.Create(
			path.Join(this.ReleativeTargetDirPath(),fmt.Sprintf("%s.go",strings.ToLower(this.ServiceGenerator.TargetNode.Name()))),
	)
	if err != nil {
		return err
	}


	return  tplIns.Execute(generatedFile,this)
}

func (this *BizGenerator) Imports() []string{
	return []string {
		`"context"`,
	}
}

func (this *BizGenerator) StructName() string{
	return fmt.Sprintf("%sBiz",this.ServiceGenerator.TargetNode.Name())
}

func (this *BizGenerator) Generated_CreateFuncName() string{
	return fmt.Sprintf("Create")
}

func (this *BizGenerator) ReleativeTargetDirPath() string{
	return "internal/biz"
}

func (this *BizGenerator) Generated_PkgName() string{
	return "biz"
}


func (this *BizGenerator) Generated_PkgPath()string{
	return path.Join(this.ServiceGenerator.Adaptor.Core.Module.Path, this.ReleativeTargetDirPath())
}

func (this *BizGenerator) Generated_DataInterfaceName() string{
	return fmt.Sprintf("%sStorage",this.ServiceGenerator.TargetNode.Name())
}

func (this *BizGenerator) Generated_DataInterfaceCreateFuncName() string{
	return "Create"
}
