package cbc

import (
	tpl "text/template"
	"strings"
	"path"
	"os"
	"fmt"
	ent "github.com/pigfall/ent_utils"
)

type ServiceGenerator struct {
	Adaptor    *Adaptor
	TargetNode *ent.Type
}

func NewServiceGenerator(adaptor *Adaptor, targetNode *ent.Type) *ServiceGenerator {
	return &ServiceGenerator{
		Adaptor:    adaptor,
		TargetNode: targetNode,
	}
}


func (this *ServiceGenerator) Generate() error {
	const targetSvcDirPath = "internal/service"
	err := os.MkdirAll(targetSvcDirPath,os.ModePerm)
	if err != nil {
		return err
	}

	generatedSvcFile, err := os.Create(path.Join(targetSvcDirPath,fmt.Sprintf("%s.go",strings.ToLower(this.TargetNode.Name()))))
	if err != nil {
		return err
	}
	defer generatedSvcFile.Close()

	svcTpl,err := tpl.New("service.tmpl").ParseFS(templates,"tpls/service.tmpl")
	if err != nil {
		return err
	}
	err = svcTpl.Execute(generatedSvcFile,this)
	if err != nil {
		return err
	}

	return nil
}

func (this *ServiceGenerator) Generated_SvcStructName() string{
	return fmt.Sprintf("%sSvc",this.TargetNode.Name())
}

func (this *ServiceGenerator) Imports() []string{
	return []string{
		`"context"`,
		fmt.Sprintf("\"%s\"",this.Adaptor.Core.Generated_PkgPath()),
	}
}
