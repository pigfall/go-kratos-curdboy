package cbc

import(
	"path"
	tpl "text/template"

	"os"
)

type CURDParamProtoGenerator struct{
	Adaptor *Adaptor
}

func NewCURDParamProtoGenerator(adaptor *Adaptor) *CURDParamProtoGenerator {
	return &CURDParamProtoGenerator {
		Adaptor: adaptor,
	}
}

func (this *CURDParamProtoGenerator ) Generate() error {
	var targetDirPath = this.ReleativeTargetDirPath()
	err := os.MkdirAll(targetDirPath,os.ModePerm)
	if err != nil{
		return err
	}
	generatedFile,err := os.Create(path.Join(targetDirPath,"curd_param.proto"))
	if err != nil {
		return err
	}
	defer generatedFile.Close()

	tplIns,err := tpl.New("curd_param_proto.tmpl").ParseFS(templates,"tpls/curd_param_proto.tmpl")
	if err != nil {
		return err
	}

	return tplIns.Execute(generatedFile,this)
}

func (this *CURDParamProtoGenerator) Generated_PkgName()string {
	return "common"
}

func (this *CURDParamProtoGenerator) ProtoPackageName() string{
	return "curdboy.common"
}

func (this *CURDParamProtoGenerator) ReleativeTargetDirPath() string{
	return "api"
}

func (this *CURDParamProtoGenerator ) Generated_PkgPath() string{
	return path.Join(this.Adaptor.Core.Module.Path,this.ReleativeTargetDirPath())
}
