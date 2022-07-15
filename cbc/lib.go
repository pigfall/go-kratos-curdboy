package cbc

import(
	cbc_core "github.com/pigfall/curdboy/curdboyc"
	"embed"
)

//go:embed tpls/*
var templates embed.FS


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
	if err := NewCURDParaProtoGenerator(this).Generate();err != nil{
		return err
	}
	// }

	// { service for nodes
	for _,node := range this.Core.Graph.GetNodes(){
		err := NewServiceGenerator(this,node).Generate()
		if err != nil {
			return err
		}
	}
	// }

	return nil
}
