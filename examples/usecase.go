package examples

import "fmt"

type AwesomeUsecase struct {
	idea MyGeniusIdea
}

func NewAwesomeUsecase(idea MyGeniusIdea) *AwesomeUsecase {
	return &AwesomeUsecase{idea: idea}
}

func (au *AwesomeUsecase) EverythingFunc() {
	var awesome string
	var worst interface{}

	greatReturn, err := au.idea.Smart(awesome, worst)
	if nil != err {
		return
	}

	fmt.Printf("greatReturn --> %v   ,   err --> %v\n", greatReturn, err)
}
