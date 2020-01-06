package android

import (
	"errors"
	"fmt"
)

func (a *NativeActivity) JNICall(fn func(env *JNIEnv,
	activity Jobject, activityClass, contextClass *Jclass) error) (err error) {
	a.Deref()
	if ret := JNIAttachCurrentThread(a.Vm, &a.Env, &JavaVMAttachArgs{
		Version: JNIVersion16,
		Name:    s("NativeThread"),
	}); ret == JNIErr {
		return errors.New("JNICall: JNIAttachCurrentThread failed")
	}
	defer JNIDetachCurrentThread(a.Vm)
	activityClass := JNIEnvGetObjectClass(a.Env, a.Clazz)
	contextClass := JNIEnvFindClass(a.Env, JClassContext.Name())

	defer func(err *error) {
		if v := recover(); v != nil {
			*err = fmt.Errorf("JNICall handler failed: %v", v)
		}
	}(&err)
	err = fn(a.Env, a.Clazz, activityClass, contextClass)
	return err
}

type JNITypeSpec struct {
	Signature string
	Class className
	IsArray bool
}

func (j JNITypeSpec) Sig() string {
	if len(j.Signature) > 0 {
		return j.Signature + "\x00"
	}
	return JNITypeSig(string(j.Class), j.IsArray) + "\x00"
}

func JNIMethodSig(ret JNITypeSpec, args ...JNITypeSpec) string {
	var argsStr string
	for _, arg := range args {
		if len(arg.Signature) > 0 {
			argsStr += arg.Signature
			continue
		}
		if len(arg.Class) > 0 {
			argsStr += JNITypeSig(string(arg.Class), arg.IsArray)
		}
	}
	var retStr string
	if len(ret.Signature) > 0 {
		retStr = ret.Signature
	} else if len(ret.Class) > 0 {
		retStr = JNITypeSig(string(ret.Class), ret.IsArray)
	} else {
		retStr = JNITypeSig(string(JVoid))
	}
	sig := fmt.Sprintf("(%s)%s\x00", argsStr, retStr)
	return sig
}

func JNITypeSig(name string, arr ...bool) string {
	var prefix string
	if len(arr) > 0 && arr[0] {
		prefix = "["
	}
	switch className(name) {
	case JVoid:
		return prefix + "V"
	case JBoolean:
		return prefix + "Z"
	case JByte:
		return prefix + "B"
	case JChar:
		return prefix + "C"
	case JShort:
		return prefix + "S"
	case JInt:
		return prefix + "I"
	case JLong:
		return prefix + "J"
	case JFloat:
		return prefix + "F"
	case JDouble:
		return prefix + "D"
	default:
		return fmt.Sprintf("%sL%s;", prefix, name)
	}
}

type className string

func (c className) Name() string {
	return string(c) + "\x00"
}

func (c className) Spec(arr ...bool) JNITypeSpec {
	return JNITypeSpec{
		Class:   c,
		IsArray: len(arr) > 0 && arr[0],
	}
}
