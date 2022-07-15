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

	for _,node := range this.Core.Graph.GetNodes(){
		err := NewServiceGenerator(this,node).Generate()
		if err != nil {
			return err
		}
	}

	return nil
}
