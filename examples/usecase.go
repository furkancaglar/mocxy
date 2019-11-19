package examples

import "fmt"

type AwesomeUsecase struct {
	idea MyGeniusIdea
}

func NewAwesomeUsecase(idea MyGeniusIdea) *AwesomeUsecase {
	return &AwesomeUsecase{idea: idea}
}

func (au *AwesomeUsecase) EverythingFunc(param1 string, param2 interface{}) {

	greatReturn, err := au.idea.Smart(param1, param2)
	fmt.Printf("greatReturn --> %v   ,   err --> %v\n", greatReturn, err)
	if nil != err {
		return
	}

}
