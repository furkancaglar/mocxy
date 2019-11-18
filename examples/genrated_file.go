package examples

import "time"

type (
	FakeMyGeniusIdea struct {
		smartArgs []*smartArgs
	}

	smartArgs struct {
		awesome struct {
			value string
			any   bool
		}
		worst struct {
			value interface{}
			any   bool
		}
		greatReturn struct {
			value time.Time
			any   bool
		}
		errReturn struct {
			value error
			any   bool
		}
	}
)

func NewFakeMyGeniusIdea() *FakeMyGeniusIdea {
	return &FakeMyGeniusIdea{}
}

// Smart fake implementation of interface.
func (f *FakeMyGeniusIdea) Smart(awesome string, worst interface{}) (greatReturn time.Time, errReturn error) {

	// TODO: implement general default logic to match the expected calls from field `smartArgs []*smartArgs`.

	return
}

// Prepare for conventional fit. For example: `&FakeMyGeniusIdea{}.Prepare().SmartAllAny()`
func (f *FakeMyGeniusIdea) Prepare() *FakeMyGeniusIdea {
	return f
}

// SmartAllAny creates a new `smartArgs` and sets all args to any.
//
// IMPORTANT NOTE: This means anything will be accepted when the `Smart` method called.
//
func (f *FakeMyGeniusIdea) SmartAllAny() *smartArgs {
	var args = &smartArgs{
		awesome: struct {
			value string
			any   bool
		}{any: true},
		worst: struct {
			value interface{}
			any   bool
		}{any: true},
		greatReturn: struct {
			value time.Time
			any   bool
		}{any: true},
		errReturn: struct {
			value error
			any   bool
		}{any: true},
	}

	f.smartArgs = append(f.smartArgs, args)

	return args
}

func (args *smartArgs) SetAwesome(awesome string) *smartArgs {
	args.awesome.any = false
	args.awesome.value = awesome
	return args
}

func (args *smartArgs) SetWorst(worst interface{}) *smartArgs {
	args.worst.any = false
	args.worst.value = worst
	return args
}

func (args *smartArgs) SetGreatReturn(greatReturn time.Time) *smartArgs {
	args.greatReturn.any = false
	args.greatReturn.value = greatReturn
	return args
}

func (args *smartArgs) SetErrReturn(errReturn error) *smartArgs {
	args.errReturn.any = false
	args.errReturn.value = errReturn
	return args
}
