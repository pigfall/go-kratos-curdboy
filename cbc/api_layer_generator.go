package cbc

import(
	"fmt"
	"strings"
	"path"
	"os"
	tpl "text/template"
)


type ServiceApiGenerator struct {
	SvcGenerator *ServiceGenerator
}

func NewServiceApiGenerator(svcGen *ServiceGenerator) *ServiceApiGenerator{
	return &ServiceApiGenerator{
		SvcGenerator: svcGen,
	}
}


func (this *ServiceApiGenerator) Generate() error {
	var relativeDirPath = this.RelativeTargetDirPath()

	err := os.MkdirAll(relativeDirPath,os.ModePerm)
	if err != nil {
		return err
	}

	var generatedFilePath = this.Generated_FilePath()
	generatedFile,err := os.Create(generatedFilePath)
	if err != nil {
		return err
	}
	defer generatedFile.Close()

	tplIns, err := tpl.New("api_layer.tmpl").Funcs(map[string]any{
		"toLowerCase":func(input string)string{
			return strings.ToLower(input)
		},
	}).ParseFS(templates,"tpls/api_layer.tmpl")
	if err != nil {
		return err
	}
	return tplIns.Execute(generatedFile,this)
}

func (this *ServiceApiGenerator) Generated_FilePath() string{
	return path.Join(this.RelativeTargetDirPath(),fmt.Sprintf("%s_gen.proto",strings.ToLower(this.SvcGenerator.TargetNode.Name())))
}

func (this *ServiceApiGenerator) ApiVersion() string{
	// TODO configable
	return "v1"
}

func (this *ServiceApiGenerator) ProtoPkgName() string{
	return fmt.Sprintf("%s.%s",strings.ToLower(this.SvcGenerator.TargetNode.Name()),this.ApiVersion())
}

func (this *ServiceApiGenerator) Generated_PkgName() string{
	return fmt.Sprintf("pb%s",strings.ToLower(this.SvcGenerator.TargetNode.Name()))
}

func (this *ServiceApiGenerator) Generated_PkgPath() string{
	return path.Join(this.SvcGenerator.Adaptor.Core.Module.Path,this.RelativeTargetDirPath())
}



func (this *ServiceApiGenerator) RelativeTargetDirPath() string {
	return fmt.Sprintf("api/%s/%s",strings.ToLower(this.SvcGenerator.TargetNode.Name()),this.ApiVersion()) 
}

func (this *ServiceApiGenerator) Generated_Proto_QueryResponseMessageName() string{
	return fmt.Sprintf("%sQueryResponse",this.SvcGenerator.TargetNode.Name())
}

func (this *ServiceApiGenerator) Generated_Proto_QueryResponseMetaMessageName() string{
	return fmt.Sprintf("%sQueryResponseMeta",this.SvcGenerator.TargetNode.Name())
}
