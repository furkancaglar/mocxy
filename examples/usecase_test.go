package examples

import (
	"errors"
	"testing"
	"time"
)

func TestAwesomeUsecase_EverythingFunc(t *testing.T) {
	type (
		fields struct {
			idea *FakeMyGeniusIdea
		}
		args struct {
			param1 string
			param2 interface{}
		}
		test struct {
			name    string
			prepare func(idea *FakeMyGeniusIdea)
			fields  fields
			args    args
		}
	)
	tests := []test{
		{
			name: "success",
			prepare: func(idea *FakeMyGeniusIdea) {
				// each line means different possible future call to `Smart` method.
				idea.Prepare().SmartAllAny()
				idea.Prepare().SmartAllAny().SetAwesome("awesome")
				idea.Prepare().SmartAllAny().SetWorst("worst").SetErrReturn(errors.New("expected error"))
				idea.Prepare().SmartAllAny().SetWorst("worst").SetAwesome("awesome")
				idea.Prepare().SmartAllAny().SetAwesome("awesome").SetGreatReturn(time.Now())
				idea.Prepare().SmartAllAny().SetAwesome("awesome").SetWorst(0).SetGreatReturn(time.Now())
			},
			fields: fields{idea: NewFakeMyGeniusIdea()},
			args: args{
				param1: "awesome",
				param2: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := NewAwesomeUsecase(tt.fields.idea)

			tt.prepare(tt.fields.idea)

			au.EverythingFunc(tt.args.param1, tt.args.param2)
		})
	}
}
