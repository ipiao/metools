package jsonapi

import (
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/ipiao/metools/mutils"
)

type idType int

const (
	typeInvalid idType = iota
	typeInt
	typeString
	typeAnnoy
)

type idLoc struct {
	annoy bool
	index int
	idType
}

var (
	typeMap     = map[reflect.Type]string{}
	idIndexMap  = map[reflect.Type]*idLoc{}
	typeLock    = new(sync.RWMutex)
	idIndexLock = new(sync.RWMutex)
)

// Resource is Resource
type Resource struct {
	ID         string      `json:"id"`
	Type       string      `json:"type"`
	Attributes interface{} `json:"attributes,omitempty"`
}

// TransModelToData transform model to data
func TransModelToData(model interface{}) interface{} {
	v := reflect.ValueOf(model)
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	t := v.Type()
	idIndx := getIDLoc(t)
	typeName := getType(t)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		var attributes = []*Resource{}
		for i := 0; i < v.Len(); i++ {
			attributes = append(attributes, transModelToResourceWith2(v.Index(i), idIndx, typeName))
		}
		return attributes
	case reflect.Map:
		var attributes = map[interface{}]*Resource{}
		for _, i := range v.MapKeys() {
			attributes[i.Interface()] = transModelToResourceWith2(v.MapIndex(i), idIndx, typeName)
		}
		return attributes
	case reflect.Struct:
		return transModelToResourceWith2(v, idIndx, typeName)
	}
	return model
}

// v must be kind of reflect.Struct
func transModelToResourceWith2(v reflect.Value, idloc *idLoc, typeName string) *Resource {
	v = reflect.Indirect(v)
	var idStr = getIDStr(v, idloc)
	return &Resource{
		Type:       typeName,
		ID:         idStr,
		Attributes: v.Interface(),
	}
}

func getIDStr(v reflect.Value, idloc *idLoc) string {
	if idloc.idType == typeInvalid {
	} else if idloc.idType == typeInt {
		return strconv.FormatInt(v.Field(idloc.index).Int(), 10)
	} else if idloc.idType == typeString {
		return v.Field(idloc.index).String()
	} else if idloc.idType == typeAnnoy {
		return getIDStr(v.Field(idloc.index), getIDLoc(v.Field(idloc.index).Type()))
	}
	return ""
}

// GetType to export
func GetType(i interface{}) string {
	t := reflect.TypeOf(i)
	return getType(t)
}

func getType(t reflect.Type) string {

	var typename string

	typeLock.RLock()
	r, ok := typeMap[t]
	typename = r
	typeLock.RUnlock()

	typeLock.Lock()
	if !ok {
		rt := indirect(t)
		if rt.Kind() != reflect.Struct {
			typename = ""
		} else {
			typename = getTypeName(rt)
		}
		typeMap[t] = typename
	}
	typeLock.Unlock()

	return typename
}

func getTypeName(t reflect.Type) string {
	// typeLock.Lock()
	// defer typeLock.Unlock()
	v := reflect.New(t)
	var typename string
	m, exist := t.MethodByName("Type")
	if exist {
		var args = make([]reflect.Value, 0)
		typename = v.Method(m.Index).Call(args)[0].String()
	} else {
		typename = mutils.SnakeName(t.Name())
	}
	return typename
}

func getIDLoc(t reflect.Type) *idLoc {
	var idloc = new(idLoc)

	idIndexLock.RLock()
	r, ok := idIndexMap[t]
	idloc = r
	idIndexLock.RUnlock()

	idIndexLock.Lock()
	if !ok {
		rt := indirect(t)
		if rt.Kind() != reflect.Struct {
			idloc.idType = typeInvalid
		} else {
			idloc = getIDIndex(rt)
		}
		idIndexMap[t] = idloc
	}
	idIndexLock.Unlock()

	return idloc
}

func getIDIndex(t reflect.Type) *idLoc {
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Anonymous {
			if r := getIDLoc(t.Field(i).Type); r.idType != typeInvalid {
				return &idLoc{
					idType: typeAnnoy,
					index:  i,
				}
			}
		}
		if strings.EqualFold(t.Field(i).Tag.Get("jsonapi"), "id") || strings.ToLower(t.Field(i).Name) == "id" {
			if t.Field(i).Type.Kind() == reflect.String {
				return &idLoc{
					idType: typeString,
					index:  i,
				}
			} else if t.Field(i).Type.Kind() >= reflect.Int && t.Field(i).Type.Kind() <= reflect.Uint64 {
				return &idLoc{
					idType: typeInt,
					index:  i,
				}
			}
		}
	}
	return &idLoc{idType: typeInvalid}
}

func indirect(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Slice || t.Kind() == reflect.Array || t.Kind() == reflect.Map {
		t = t.Elem()
	}
	return t
}
