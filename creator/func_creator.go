package creator

import (
	"reflect"
	"strings"
)

// FuncCreator is creator for Func
type FuncCreator struct {
	Name          string
	recvierType   reflect.Type
	Receiver      string
	ReceiverAlias string
	ReceiverAnnoy bool
	HasReceiver   bool
	ArgsIn        []Arg
	ArgsOut       []Arg
	FuncBody      string
}

// AddArgIn2 set the receiver
func (f *FuncCreator) AddArgIn2(alias string, t string) *FuncCreator {
	arg := Arg{
		Alias: alias,
		Type:  t,
	}
	f.ArgsIn = append(f.ArgsIn, arg)
	return f
}

// AddArgOut2 set the receiver
func (f *FuncCreator) AddArgOut2(t string) *FuncCreator {
	arg := Arg{
		Type:       t,
		AliasAnnoy: true,
	}
	f.ArgsOut = append(f.ArgsOut, arg)
	return f
}

// Arg for arg
type Arg struct {
	Alias      string
	Type       string
	AliasAnnoy bool
	t          reflect.Type
}

// NewFunc construct a FuncCreator
func NewFunc(name string) *FuncCreator {
	return &FuncCreator{
		Name:        name,
		recvierType: nil,
		ArgsIn:      make([]Arg, 0),
		ArgsOut:     make([]Arg, 0),
	}
}

// SetReceiver set the receiver
func (f *FuncCreator) SetReceiver(rec reflect.Type) *FuncCreator {
	if rec == nil {
		panic("receiver can not be nil")
	}
	f.recvierType = rec
	f.ReceiverAlias = strings.ToLower(rec.Name()[:0])
	f.Receiver = rec.String()
	f.HasReceiver = true
	return f
}

// SetReceiver2 set the receiver
func (f *FuncCreator) SetReceiver2(rec string) *FuncCreator {
	alais := rec[:1]
	if rec[:1] == "*" {
		alais = rec[1:2]
	}
	f.ReceiverAlias = strings.ToLower(alais)
	f.Receiver = rec
	f.HasReceiver = true
	return f
}

// SetReceiverAnnoy set the receiverAnnoy
func (f *FuncCreator) SetReceiverAnnoy(b ...bool) *FuncCreator {
	if len(b) == 0 {
		f.ReceiverAnnoy = true
	} else {
		f.ReceiverAnnoy = b[0]
	}
	return f
}
