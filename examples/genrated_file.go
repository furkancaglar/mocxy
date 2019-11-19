package examples

import (
	"reflect"
	"sort"
	"time"
)

type (
	FakeMyGeniusIdea struct {
		prepare *prepareStruct
	}

	prepareStruct struct {
		smartArgs []*smartArgs
	}

	smartArgs struct {
		howManyParamsSet int

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
	f.prepare.order()

	var (
		found      bool
		foundIndex int
	)

	//// TODO: implement general default logic to match the expected calls from field `smartArgs []*smartArgs`.
	for i, args := range f.prepare.smartArgs {
		if !args.awesome.any && !args.worst.any && reflect.DeepEqual(args.awesome.value, awesome) && reflect.DeepEqual(args.worst.value, worst) {
			foundIndex = i
			found = true
			break
		}
		if !args.awesome.any && args.worst.any && reflect.DeepEqual(args.awesome.value, awesome) {
			foundIndex = i
			found = true
			break
		}
		if args.awesome.any && !args.worst.any && reflect.DeepEqual(args.worst.value, worst) {
			foundIndex = i
			found = true
			break
		}

		if args.awesome.any && args.worst.any {
			foundIndex = i
			found = true
			break
		}
	}

	if found {
		return f.prepare.smartArgs[foundIndex].greatReturn.value, f.prepare.smartArgs[foundIndex].errReturn.value
	}

	panic("no matching prepared mock!")
}

// Prepare for conventional fit. For example: `&FakeMyGeniusIdea{}.Prepare().SmartAllAny()`
func (f *FakeMyGeniusIdea) Prepare() *prepareStruct {
	if nil == f.prepare {
		f.prepare = &prepareStruct{}
	}
	return f.prepare
}

// SmartAllAny creates a new `smartArgs` and sets all args to any.
//
// IMPORTANT NOTE: This means anything will be accepted when the `Smart` method called.
//
func (p *prepareStruct) SmartAllAny() *smartArgs {
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

	p.smartArgs = append(p.smartArgs, args)

	return args
}

func (args *smartArgs) SetAwesome(awesome string) *smartArgs {
	args.awesome.any = false
	args.awesome.value = awesome
	args.howManyParamsSet++
	return args
}

func (args *smartArgs) SetWorst(worst interface{}) *smartArgs {
	args.worst.any = false
	args.worst.value = worst
	args.howManyParamsSet++
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

func (p *prepareStruct) order() {
	sort.Slice(p.smartArgs, func(i, j int) bool {
		return p.smartArgs[i].howManyParamsSet > p.smartArgs[j].howManyParamsSet
	})
}
