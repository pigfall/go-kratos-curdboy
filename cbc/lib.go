package cbc

import(
	cbc_core "github.com/pigfall/curdboy/curdboyc"
	"embed"
)

//go:embed tpls/*
var templates embed.FS

//go:embed third_party/openapi_protos/*
var thirdParty embed.FS



type Adaptor struct {
	Core *cbc_core.CURDGraphGenerator
}

func NewAdaptor(core *cbc_core.CURDGraphGenerator)*Adaptor{
	return &Adaptor{
		Core: core,
	}
}


func (this *Adaptor) Generate() error {

	// { proto for curd param

	curdParamProtoGenerator :=NewCURDParamProtoGenerator(this)
	if err := curdParamProtoGenerator.Generate();err != nil{
		return err
	}
	// }

	// { service for nodes
	for _,node := range this.Core.Graph.GetNodes(){
		err := NewServiceGenerator(this,node,curdParamProtoGenerator).Generate()
		if err != nil {
			return err
		}
	}
	// }

	return nil
}
