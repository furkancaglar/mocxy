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
		test struct {
			name    string
			prepare func(idea *FakeMyGeniusIdea)
			fields  fields
		}
	)
	tests := []test{
		{
			name: "success",
			prepare: func(idea *FakeMyGeniusIdea) {
				// each line means different possible future call to `Smart` method.
				idea.Prepare().SmartAllAny()
				idea.Prepare().SmartAllAny().SetAwesome("awesome")
				idea.Prepare().SmartAllAny().SetAwesome("awesome").SetGreatReturn(time.Now())
				idea.Prepare().SmartAllAny().SetAwesome("awesome").SetWorst(0).SetGreatReturn(time.Now())
				idea.Prepare().SmartAllAny().SetAwesome("awesome").SetWorst(0).SetErrReturn(errors.New("some error"))
			},
			fields: fields{idea: NewFakeMyGeniusIdea()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := NewAwesomeUsecase(tt.fields.idea)

			au.EverythingFunc()
		})
	}
}
