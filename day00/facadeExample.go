// Package day00 is an example of the Facade Pattern
package day00

import "strings"

type SomeActions struct {
	action1 *Action1
	action2 *Action2
	action3 *Action3
}

type Action1 struct {
}

type Action2 struct {
}

type Action3 struct {
}

func InitActions() *SomeActions {
	return &SomeActions{
		action1: &Action1{},
		action2: &Action2{},
		action3: &Action3{},
	}
}

func (a *SomeActions) DoActions() string {
	res := []string{
		a.action1.Action1(),
		a.action2.Action2(),
		a.action3.Action3(),
	}
	return strings.Join(res, "\n")
}

func (a *Action1) Action1() string {
	return "Simple text1 to show in example"
}

func (a *Action2) Action2() string {
	return "Simple text2 to show in example"
}

func (a *Action3) Action3() string {
	return "Simple text3 to show in example"
}
