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
	CURDParamGenerator *CURDParamProtoGenerator 
	ServiceApiGenerator *ServiceApiGenerator
	BizGenerator *BizGenerator
}

func NewServiceGenerator(adaptor *Adaptor, targetNode *ent.Type,curdParamGenerator *CURDParamProtoGenerator) *ServiceGenerator {
	s := &ServiceGenerator{
		Adaptor:    adaptor,
		TargetNode: targetNode,
		CURDParamGenerator: curdParamGenerator,
	}
	s.ServiceApiGenerator = NewServiceApiGenerator(s)
	s.BizGenerator = NewBizGenerator(s)
	return s
}


func (this *ServiceGenerator) Generate() error {
	// { api define
	apiGenerator := this.ServiceApiGenerator
	err := apiGenerator.Generate()
	if err != nil {
		return err
	}
	// }

	// { service layer
	const targetSvcDirPath = "internal/service"
	err = os.MkdirAll(targetSvcDirPath,os.ModePerm)
	if err != nil {
		return err
	}

	generatedSvcFile, err := os.Create(path.Join(targetSvcDirPath,fmt.Sprintf("%s.go",strings.ToLower(this.TargetNode.Name()))))
	if err != nil {
		return err
	}
	defer generatedSvcFile.Close()

	svcTpl,err := tpl.New("service_layer.tmpl").ParseFS(templates,"tpls/service_layer.tmpl")
	if err != nil {
		return err
	}
	//svcTpl.ParseFS(this.Adaptor.Core.TemplatesFS(),"")
	err = svcTpl.Execute(generatedSvcFile,this)
	if err != nil {
		return err
	}

	// { generate biz layer
	err = this.BizGenerator.Generate()
	if err != nil {
		return err
	}

	// }

	return nil
	// }
}

func (this *ServiceGenerator) Generated_SvcStructName() string{
	return fmt.Sprintf("%sSvc",this.TargetNode.Name())
}

func (this *ServiceGenerator) Imports() []string{
	return []string{
		`"context"`,
		//fmt.Sprintf("\"%s\"",this.Adaptor.Core.Generated_PkgPath()),
		fmt.Sprintf("\"%s\"",this.CURDParamGenerator.Generated_PkgPath()),
		fmt.Sprintf("\"%s\"",this.ServiceApiGenerator.Generated_PkgPath()),
		"structpb \"google.golang.org/protobuf/types/known/structpb\"",
		fmt.Sprintf("\"%s\"",this.BizGenerator.Generated_PkgPath()),
	}
}

func (this *ServiceGenerator) ApiLayer() *ServiceApiGenerator {
	return this.ServiceApiGenerator
}

func (this *ServiceGenerator) BizLayer() *BizGenerator {
	return this.BizGenerator
}

