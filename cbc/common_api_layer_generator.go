package cbc

import(
	"path"
	tpl "text/template"
	"io"
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
	generatedFile,err := os.Create(path.Join(targetDirPath,"common.proto"))
	if err != nil {
		return err
	}
	defer generatedFile.Close()

	tplIns,err := tpl.New("common_api_layer.tmpl").ParseFS(templates,"tpls/common_api_layer.tmpl")
	if err != nil {
		return err
	}

	err =  tplIns.Execute(generatedFile,this)
	if err != nil{
		return err
	}

	// { put openapi proto to third_party
	openApiDirPath := path.Join("third_party","protoc-gen-openapiv2","options")
	err = os.MkdirAll(openApiDirPath,os.ModePerm)
	if err != nil {
		return err
	}
	var embedFSBasePath = "third_party/openapi_protos"
	entries,err := thirdParty.ReadDir(embedFSBasePath)
	if err != nil {
		return err
	}
	for _,e := range entries{
		f,err := thirdParty.Open(path.Join(embedFSBasePath,e.Name()))
		if err != nil {
			return err
		}
		
		dstFile,err := os.Create(path.Join(openApiDirPath,e.Name()))
		if err != nil {
			return err
		}
		_,err = io.Copy(dstFile,f)
		if err != nil {
			return err
		}
		err = dstFile.Close()
		if err != nil {
			return err
		}
	}
	// }
	return nil
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
