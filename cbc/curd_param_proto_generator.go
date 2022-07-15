package cbc

import(
	"path"
	tpl "text/template"

	"os"
)

type CURDParaProtoGenerator struct{
	Adaptor *Adaptor
}

func NewCURDParaProtoGenerator(adaptor *Adaptor) *CURDParaProtoGenerator{
	return &CURDParaProtoGenerator{
		Adaptor: adaptor,
	}
}

func (this *CURDParaProtoGenerator) Generate() error {
	var targetDirPath = "api"
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

func (this *CURDParaProtoGenerator) Generated_PkgPath() string{
	return "api"
}
