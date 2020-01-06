package android

import "C"
import (
	"sync"
	"unsafe"
)

// Ref returns a reference to C object as it is.
func (x *AssetManager) Ref() *C.AAssetManager {
	if x == nil {
		return nil
	}
	return (*C.AAssetManager)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *AssetManager) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewAssetManagerRef converts the C object reference into a raw struct reference without wrapping.
func NewAssetManagerRef(ref unsafe.Pointer) *AssetManager {
	return (*AssetManager)(ref)
}

// NewAssetManager allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewAssetManager() *AssetManager {
	return (*AssetManager)(allocAssetManagerMemory(1))
}

// allocAssetManagerMemory allocates memory for type C.AAssetManager in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAssetManagerMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAssetManagerValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfAssetManagerValue = unsafe.Sizeof([1]C.AAssetManager{})

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *AssetManager) PassRef() *C.AAssetManager {
	if x == nil {
		x = (*AssetManager)(allocAssetManagerMemory(1))
	}
	return (*C.AAssetManager)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *AssetDir) Ref() *C.AAssetDir {
	if x == nil {
		return nil
	}
	return (*C.AAssetDir)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *AssetDir) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewAssetDirRef converts the C object reference into a raw struct reference without wrapping.
func NewAssetDirRef(ref unsafe.Pointer) *AssetDir {
	return (*AssetDir)(ref)
}

// NewAssetDir allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewAssetDir() *AssetDir {
	return (*AssetDir)(allocAssetDirMemory(1))
}

// allocAssetDirMemory allocates memory for type C.AAssetDir in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAssetDirMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAssetDirValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfAssetDirValue = unsafe.Sizeof([1]C.AAssetDir{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *AssetDir) PassRef() *C.AAssetDir {
	if x == nil {
		x = (*AssetDir)(allocAssetDirMemory(1))
	}
	return (*C.AAssetDir)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *Asset) Ref() *C.AAsset {
	if x == nil {
		return nil
	}
	return (*C.AAsset)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *Asset) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewAssetRef converts the C object reference into a raw struct reference without wrapping.
func NewAssetRef(ref unsafe.Pointer) *Asset {
	return (*Asset)(ref)
}

// NewAsset allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewAsset() *Asset {
	return (*Asset)(allocAssetMemory(1))
}

// allocAssetMemory allocates memory for type C.AAsset in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAssetMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAssetValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfAssetValue = unsafe.Sizeof([1]C.AAsset{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *Asset) PassRef() *C.AAsset {
	if x == nil {
		x = (*Asset)(allocAssetMemory(1))
	}
	return (*C.AAsset)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *JNINativeMethod) Ref() *C.JNINativeMethod {
	if x == nil {
		return nil
	}
	return (*C.JNINativeMethod)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *JNINativeMethod) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewJNINativeMethodRef converts the C object reference into a raw struct reference without wrapping.
func NewJNINativeMethodRef(ref unsafe.Pointer) *JNINativeMethod {
	return (*JNINativeMethod)(ref)
}

// NewJNINativeMethod allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewJNINativeMethod() *JNINativeMethod {
	return (*JNINativeMethod)(allocJNINativeMethodMemory(1))
}

// allocJNINativeMethodMemory allocates memory for type C.JNINativeMethod in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJNINativeMethodMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJNINativeMethodValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJNINativeMethodValue = unsafe.Sizeof([1]C.JNINativeMethod{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *JNINativeMethod) PassRef() *C.JNINativeMethod {
	if x == nil {
		x = (*JNINativeMethod)(allocJNINativeMethodMemory(1))
	}
	return (*C.JNINativeMethod)(unsafe.Pointer(x))
}

// allocJavaVMAttachArgsMemory allocates memory for type C.JavaVMAttachArgs in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJavaVMAttachArgsMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJavaVMAttachArgsValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJavaVMAttachArgsValue = unsafe.Sizeof([1]C.JavaVMAttachArgs{})

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

type stringHeader struct {
	Data uintptr
	Len  int
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = uintptr(unsafe.Pointer(p))
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
	}
	return
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(unsafe.Pointer(h.Data)), C.int(h.Len))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *JavaVMAttachArgs) Ref() *C.JavaVMAttachArgs {
	if x == nil {
		return nil
	}
	return x.refab42f020
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *JavaVMAttachArgs) Free() {
	if x != nil && x.allocsab42f020 != nil {
		x.allocsab42f020.(*cgoAllocMap).Free()
		x.refab42f020 = nil
	}
}

// NewJavaVMAttachArgsRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewJavaVMAttachArgsRef(ref unsafe.Pointer) *JavaVMAttachArgs {
	if ref == nil {
		return nil
	}
	obj := new(JavaVMAttachArgs)
	obj.refab42f020 = (*C.JavaVMAttachArgs)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *JavaVMAttachArgs) PassRef() (*C.JavaVMAttachArgs, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refab42f020 != nil {
		return x.refab42f020, nil
	}
	memab42f020 := allocJavaVMAttachArgsMemory(1)
	refab42f020 := (*C.JavaVMAttachArgs)(memab42f020)
	allocsab42f020 := new(cgoAllocMap)
	allocsab42f020.Add(memab42f020)

	var cversion_allocs *cgoAllocMap
	refab42f020.version, cversion_allocs = (C.jint)(x.Version), cgoAllocsUnknown
	allocsab42f020.Borrow(cversion_allocs)

	var cname_allocs *cgoAllocMap
	refab42f020.name, cname_allocs = unpackPCharString(x.Name)
	allocsab42f020.Borrow(cname_allocs)

	var cgroup_allocs *cgoAllocMap
	refab42f020.group, cgroup_allocs = (C.jobject)(x.Group), cgoAllocsUnknown
	allocsab42f020.Borrow(cgroup_allocs)

	x.refab42f020 = refab42f020
	x.allocsab42f020 = allocsab42f020
	return refab42f020, allocsab42f020

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x JavaVMAttachArgs) PassValue() (C.JavaVMAttachArgs, *cgoAllocMap) {
	if x.refab42f020 != nil {
		return *x.refab42f020, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *JavaVMAttachArgs) Deref() {
	if x.refab42f020 == nil {
		return
	}
	x.Version = (Jint)(x.refab42f020.version)
	x.Name = packPCharString(x.refab42f020.name)
	x.Group = (Jobject)(x.refab42f020.group)
}

// Ref returns a reference to C object as it is.
func (x *JavaVMOption) Ref() *C.JavaVMOption {
	if x == nil {
		return nil
	}
	return (*C.JavaVMOption)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *JavaVMOption) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewJavaVMOptionRef converts the C object reference into a raw struct reference without wrapping.
func NewJavaVMOptionRef(ref unsafe.Pointer) *JavaVMOption {
	return (*JavaVMOption)(ref)
}

// NewJavaVMOption allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewJavaVMOption() *JavaVMOption {
	return (*JavaVMOption)(allocJavaVMOptionMemory(1))
}

// allocJavaVMOptionMemory allocates memory for type C.JavaVMOption in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJavaVMOptionMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJavaVMOptionValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJavaVMOptionValue = unsafe.Sizeof([1]C.JavaVMOption{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *JavaVMOption) PassRef() *C.JavaVMOption {
	if x == nil {
		x = (*JavaVMOption)(allocJavaVMOptionMemory(1))
	}
	return (*C.JavaVMOption)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *JavaVMInitArgs) Ref() *C.JavaVMInitArgs {
	if x == nil {
		return nil
	}
	return (*C.JavaVMInitArgs)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *JavaVMInitArgs) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewJavaVMInitArgsRef converts the C object reference into a raw struct reference without wrapping.
func NewJavaVMInitArgsRef(ref unsafe.Pointer) *JavaVMInitArgs {
	return (*JavaVMInitArgs)(ref)
}

// NewJavaVMInitArgs allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewJavaVMInitArgs() *JavaVMInitArgs {
	return (*JavaVMInitArgs)(allocJavaVMInitArgsMemory(1))
}

// allocJavaVMInitArgsMemory allocates memory for type C.JavaVMInitArgs in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJavaVMInitArgsMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJavaVMInitArgsValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJavaVMInitArgsValue = unsafe.Sizeof([1]C.JavaVMInitArgs{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *JavaVMInitArgs) PassRef() *C.JavaVMInitArgs {
	if x == nil {
		x = (*JavaVMInitArgs)(allocJavaVMInitArgsMemory(1))
	}
	return (*C.JavaVMInitArgs)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *Configuration) Ref() *C.AConfiguration {
	if x == nil {
		return nil
	}
	return (*C.AConfiguration)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *Configuration) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewConfigurationRef converts the C object reference into a raw struct reference without wrapping.
func NewConfigurationRef(ref unsafe.Pointer) *Configuration {
	return (*Configuration)(ref)
}

// NewConfiguration allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewConfiguration() *Configuration {
	return (*Configuration)(allocConfigurationMemory(1))
}

// allocConfigurationMemory allocates memory for type C.AConfiguration in C.
// The caller is responsible for freeing the this memory via C.free.
func allocConfigurationMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfConfigurationValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfConfigurationValue = unsafe.Sizeof([1]C.AConfiguration{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *Configuration) PassRef() *C.AConfiguration {
	if x == nil {
		x = (*Configuration)(allocConfigurationMemory(1))
	}
	return (*C.AConfiguration)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *Looper) Ref() *C.ALooper {
	if x == nil {
		return nil
	}
	return (*C.ALooper)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *Looper) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewLooperRef converts the C object reference into a raw struct reference without wrapping.
func NewLooperRef(ref unsafe.Pointer) *Looper {
	return (*Looper)(ref)
}

// NewLooper allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewLooper() *Looper {
	return (*Looper)(allocLooperMemory(1))
}

// allocLooperMemory allocates memory for type C.ALooper in C.
// The caller is responsible for freeing the this memory via C.free.
func allocLooperMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfLooperValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfLooperValue = unsafe.Sizeof([1]C.ALooper{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *Looper) PassRef() *C.ALooper {
	if x == nil {
		x = (*Looper)(allocLooperMemory(1))
	}
	return (*C.ALooper)(unsafe.Pointer(x))
}

func (x LooperCallbackFunc) PassRef() (ref *C.ALooper_callbackFunc, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if looperCallbackFunc7E6B484CFunc == nil {
		looperCallbackFunc7E6B484CFunc = x
	}
	return (*C.ALooper_callbackFunc)(C.ALooper_callbackFunc_7e6b484c), nil
}

func (x LooperCallbackFunc) PassValue() (ref C.ALooper_callbackFunc, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if looperCallbackFunc7E6B484CFunc == nil {
		looperCallbackFunc7E6B484CFunc = x
	}
	return (C.ALooper_callbackFunc)(C.ALooper_callbackFunc_7e6b484c), nil
}

func NewLooperCallbackFuncRef(ref unsafe.Pointer) *LooperCallbackFunc {
	return (*LooperCallbackFunc)(ref)
}

//export looperCallbackFunc7E6B484C
func looperCallbackFunc7E6B484C(cfd C.int, cevents C.int, cdata unsafe.Pointer) C.int {
	if looperCallbackFunc7E6B484CFunc != nil {
		fd7e6b484c := (int32)(cfd)
		events7e6b484c := (int32)(cevents)
		data7e6b484c := (unsafe.Pointer)(unsafe.Pointer(cdata))
		ret7e6b484c := looperCallbackFunc7E6B484CFunc(fd7e6b484c, events7e6b484c, data7e6b484c)
		ret, _ := (C.int)(ret7e6b484c), cgoAllocsUnknown
		return ret
	}
	panic("callback func has not been set (race?)")
}

var looperCallbackFunc7E6B484CFunc LooperCallbackFunc

// allocNativeActivityMemory allocates memory for type C.ANativeActivity in C.
// The caller is responsible for freeing the this memory via C.free.
func allocNativeActivityMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfNativeActivityValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfNativeActivityValue = unsafe.Sizeof([1]C.ANativeActivity{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *NativeActivity) Ref() *C.ANativeActivity {
	if x == nil {
		return nil
	}
	return x.ref2cc295bf
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *NativeActivity) Free() {
	if x != nil && x.allocs2cc295bf != nil {
		x.allocs2cc295bf.(*cgoAllocMap).Free()
		x.ref2cc295bf = nil
	}
}

// NewNativeActivityRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewNativeActivityRef(ref unsafe.Pointer) *NativeActivity {
	if ref == nil {
		return nil
	}
	obj := new(NativeActivity)
	obj.ref2cc295bf = (*C.ANativeActivity)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *NativeActivity) PassRef() (*C.ANativeActivity, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref2cc295bf != nil {
		return x.ref2cc295bf, nil
	}
	mem2cc295bf := allocNativeActivityMemory(1)
	ref2cc295bf := (*C.ANativeActivity)(mem2cc295bf)
	allocs2cc295bf := new(cgoAllocMap)
	allocs2cc295bf.Add(mem2cc295bf)

	var ccallbacks_allocs *cgoAllocMap
	ref2cc295bf.callbacks, ccallbacks_allocs = *(**C.struct_ANativeActivityCallbacks)(unsafe.Pointer(&x.Callbacks)), cgoAllocsUnknown
	allocs2cc295bf.Borrow(ccallbacks_allocs)

	var cvm_allocs *cgoAllocMap
	ref2cc295bf.vm, cvm_allocs = *(**C.JavaVM)(unsafe.Pointer(&x.Vm)), cgoAllocsUnknown
	allocs2cc295bf.Borrow(cvm_allocs)

	var cenv_allocs *cgoAllocMap
	ref2cc295bf.env, cenv_allocs = *(**C.JNIEnv)(unsafe.Pointer(&x.Env)), cgoAllocsUnknown
	allocs2cc295bf.Borrow(cenv_allocs)

	var cclazz_allocs *cgoAllocMap
	ref2cc295bf.clazz, cclazz_allocs = (C.jobject)(x.Clazz), cgoAllocsUnknown
	allocs2cc295bf.Borrow(cclazz_allocs)

	var cinternalDataPath_allocs *cgoAllocMap
	ref2cc295bf.internalDataPath, cinternalDataPath_allocs = unpackPCharString(x.InternalDataPath)
	allocs2cc295bf.Borrow(cinternalDataPath_allocs)

	var cexternalDataPath_allocs *cgoAllocMap
	ref2cc295bf.externalDataPath, cexternalDataPath_allocs = unpackPCharString(x.ExternalDataPath)
	allocs2cc295bf.Borrow(cexternalDataPath_allocs)

	var csdkVersion_allocs *cgoAllocMap
	ref2cc295bf.sdkVersion, csdkVersion_allocs = (C.int32_t)(x.SdkVersion), cgoAllocsUnknown
	allocs2cc295bf.Borrow(csdkVersion_allocs)

	var cinstance_allocs *cgoAllocMap
	ref2cc295bf.instance, cinstance_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.Instance)), cgoAllocsUnknown
	allocs2cc295bf.Borrow(cinstance_allocs)

	var cassetManager_allocs *cgoAllocMap
	ref2cc295bf.assetManager, cassetManager_allocs = *(**C.AAssetManager)(unsafe.Pointer(&x.AssetManager)), cgoAllocsUnknown
	allocs2cc295bf.Borrow(cassetManager_allocs)

	var cobbPath_allocs *cgoAllocMap
	ref2cc295bf.obbPath, cobbPath_allocs = unpackPCharString(x.ObbPath)
	allocs2cc295bf.Borrow(cobbPath_allocs)

	x.ref2cc295bf = ref2cc295bf
	x.allocs2cc295bf = allocs2cc295bf
	return ref2cc295bf, allocs2cc295bf

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x NativeActivity) PassValue() (C.ANativeActivity, *cgoAllocMap) {
	if x.ref2cc295bf != nil {
		return *x.ref2cc295bf, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *NativeActivity) Deref() {
	if x.ref2cc295bf == nil {
		return
	}
	x.Callbacks = (*NativeActivityCallbacks)(unsafe.Pointer(x.ref2cc295bf.callbacks))
	x.Vm = (*JavaVM)(unsafe.Pointer(x.ref2cc295bf.vm))
	x.Env = (*JNIEnv)(unsafe.Pointer(x.ref2cc295bf.env))
	x.Clazz = (Jobject)(x.ref2cc295bf.clazz)
	x.InternalDataPath = packPCharString(x.ref2cc295bf.internalDataPath)
	x.ExternalDataPath = packPCharString(x.ref2cc295bf.externalDataPath)
	x.SdkVersion = (int32)(x.ref2cc295bf.sdkVersion)
	x.Instance = (unsafe.Pointer)(unsafe.Pointer(x.ref2cc295bf.instance))
	x.AssetManager = (*AssetManager)(unsafe.Pointer(x.ref2cc295bf.assetManager))
	x.ObbPath = packPCharString(x.ref2cc295bf.obbPath)
}

// Ref returns a reference to C object as it is.
func (x *NativeActivityCallbacks) Ref() *C.ANativeActivityCallbacks {
	if x == nil {
		return nil
	}
	return (*C.ANativeActivityCallbacks)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *NativeActivityCallbacks) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewNativeActivityCallbacksRef converts the C object reference into a raw struct reference without wrapping.
func NewNativeActivityCallbacksRef(ref unsafe.Pointer) *NativeActivityCallbacks {
	return (*NativeActivityCallbacks)(ref)
}

// NewNativeActivityCallbacks allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewNativeActivityCallbacks() *NativeActivityCallbacks {
	return (*NativeActivityCallbacks)(allocNativeActivityCallbacksMemory(1))
}

// allocNativeActivityCallbacksMemory allocates memory for type C.ANativeActivityCallbacks in C.
// The caller is responsible for freeing the this memory via C.free.
func allocNativeActivityCallbacksMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfNativeActivityCallbacksValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfNativeActivityCallbacksValue = unsafe.Sizeof([1]C.ANativeActivityCallbacks{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *NativeActivityCallbacks) PassRef() *C.ANativeActivityCallbacks {
	if x == nil {
		x = (*NativeActivityCallbacks)(allocNativeActivityCallbacksMemory(1))
	}
	return (*C.ANativeActivityCallbacks)(unsafe.Pointer(x))
}

func (x NativeActivityCreateFunc) PassRef() (ref *C.ANativeActivity_createFunc, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if nativeActivityCreateFunc76DCE4Func == nil {
		nativeActivityCreateFunc76DCE4Func = x
	}
	return (*C.ANativeActivity_createFunc)(C.ANativeActivity_createFunc_76dce4), nil
}

func NewNativeActivityCreateFuncRef(ref unsafe.Pointer) *NativeActivityCreateFunc {
	return (*NativeActivityCreateFunc)(ref)
}

//export nativeActivityCreateFunc76DCE4
func nativeActivityCreateFunc76DCE4(cactivity *C.ANativeActivity, csavedState unsafe.Pointer, csavedStateSize C.size_t) {
	if nativeActivityCreateFunc76DCE4Func != nil {
		activity76dce4 := NewNativeActivityRef(unsafe.Pointer(cactivity))
		savedState76dce4 := (unsafe.Pointer)(unsafe.Pointer(csavedState))
		savedStateSize76dce4 := (uint32)(csavedStateSize)
		nativeActivityCreateFunc76DCE4Func(activity76dce4, savedState76dce4, savedStateSize76dce4)
		return
	}
	panic("callback func has not been set (race?)")
}

var nativeActivityCreateFunc76DCE4Func NativeActivityCreateFunc

// Ref returns a reference to C object as it is.
func (x *InputEvent) Ref() *C.AInputEvent {
	if x == nil {
		return nil
	}
	return (*C.AInputEvent)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *InputEvent) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewInputEventRef converts the C object reference into a raw struct reference without wrapping.
func NewInputEventRef(ref unsafe.Pointer) *InputEvent {
	return (*InputEvent)(ref)
}

// NewInputEvent allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewInputEvent() *InputEvent {
	return (*InputEvent)(allocInputEventMemory(1))
}

// allocInputEventMemory allocates memory for type C.AInputEvent in C.
// The caller is responsible for freeing the this memory via C.free.
func allocInputEventMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfInputEventValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfInputEventValue = unsafe.Sizeof([1]C.AInputEvent{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *InputEvent) PassRef() *C.AInputEvent {
	if x == nil {
		x = (*InputEvent)(allocInputEventMemory(1))
	}
	return (*C.AInputEvent)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *InputQueue) Ref() *C.AInputQueue {
	if x == nil {
		return nil
	}
	return (*C.AInputQueue)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *InputQueue) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewInputQueueRef converts the C object reference into a raw struct reference without wrapping.
func NewInputQueueRef(ref unsafe.Pointer) *InputQueue {
	return (*InputQueue)(ref)
}

// NewInputQueue allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewInputQueue() *InputQueue {
	return (*InputQueue)(allocInputQueueMemory(1))
}

// allocInputQueueMemory allocates memory for type C.AInputQueue in C.
// The caller is responsible for freeing the this memory via C.free.
func allocInputQueueMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfInputQueueValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfInputQueueValue = unsafe.Sizeof([1]C.AInputQueue{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *InputQueue) PassRef() *C.AInputQueue {
	if x == nil {
		x = (*InputQueue)(allocInputQueueMemory(1))
	}
	return (*C.AInputQueue)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *NativeWindow) Ref() *C.ANativeWindow {
	if x == nil {
		return nil
	}
	return (*C.ANativeWindow)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *NativeWindow) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewNativeWindowRef converts the C object reference into a raw struct reference without wrapping.
func NewNativeWindowRef(ref unsafe.Pointer) *NativeWindow {
	return (*NativeWindow)(ref)
}

// NewNativeWindow allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewNativeWindow() *NativeWindow {
	return (*NativeWindow)(allocNativeWindowMemory(1))
}

// allocNativeWindowMemory allocates memory for type C.ANativeWindow in C.
// The caller is responsible for freeing the this memory via C.free.
func allocNativeWindowMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfNativeWindowValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfNativeWindowValue = unsafe.Sizeof([1]C.ANativeWindow{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *NativeWindow) PassRef() *C.ANativeWindow {
	if x == nil {
		x = (*NativeWindow)(allocNativeWindowMemory(1))
	}
	return (*C.ANativeWindow)(unsafe.Pointer(x))
}

// allocNativeWindowBufferMemory allocates memory for type C.ANativeWindow_Buffer in C.
// The caller is responsible for freeing the this memory via C.free.
func allocNativeWindowBufferMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfNativeWindowBufferValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfNativeWindowBufferValue = unsafe.Sizeof([1]C.ANativeWindow_Buffer{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *NativeWindowBuffer) Ref() *C.ANativeWindow_Buffer {
	if x == nil {
		return nil
	}
	return x.ref3db2646c
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *NativeWindowBuffer) Free() {
	if x != nil && x.allocs3db2646c != nil {
		x.allocs3db2646c.(*cgoAllocMap).Free()
		x.ref3db2646c = nil
	}
}

// NewNativeWindowBufferRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewNativeWindowBufferRef(ref unsafe.Pointer) *NativeWindowBuffer {
	if ref == nil {
		return nil
	}
	obj := new(NativeWindowBuffer)
	obj.ref3db2646c = (*C.ANativeWindow_Buffer)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *NativeWindowBuffer) PassRef() (*C.ANativeWindow_Buffer, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref3db2646c != nil {
		return x.ref3db2646c, nil
	}
	mem3db2646c := allocNativeWindowBufferMemory(1)
	ref3db2646c := (*C.ANativeWindow_Buffer)(mem3db2646c)
	allocs3db2646c := new(cgoAllocMap)
	allocs3db2646c.Add(mem3db2646c)

	var cwidth_allocs *cgoAllocMap
	ref3db2646c.width, cwidth_allocs = (C.int32_t)(x.Width), cgoAllocsUnknown
	allocs3db2646c.Borrow(cwidth_allocs)

	var cheight_allocs *cgoAllocMap
	ref3db2646c.height, cheight_allocs = (C.int32_t)(x.Height), cgoAllocsUnknown
	allocs3db2646c.Borrow(cheight_allocs)

	var cstride_allocs *cgoAllocMap
	ref3db2646c.stride, cstride_allocs = (C.int32_t)(x.Stride), cgoAllocsUnknown
	allocs3db2646c.Borrow(cstride_allocs)

	var cformat_allocs *cgoAllocMap
	ref3db2646c.format, cformat_allocs = (C.int32_t)(x.Format), cgoAllocsUnknown
	allocs3db2646c.Borrow(cformat_allocs)

	var cbits_allocs *cgoAllocMap
	ref3db2646c.bits, cbits_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.Bits)), cgoAllocsUnknown
	allocs3db2646c.Borrow(cbits_allocs)

	var creserved_allocs *cgoAllocMap
	ref3db2646c.reserved, creserved_allocs = *(*[6]C.uint32_t)(unsafe.Pointer(&x.Reserved)), cgoAllocsUnknown
	allocs3db2646c.Borrow(creserved_allocs)

	x.ref3db2646c = ref3db2646c
	x.allocs3db2646c = allocs3db2646c
	return ref3db2646c, allocs3db2646c

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x NativeWindowBuffer) PassValue() (C.ANativeWindow_Buffer, *cgoAllocMap) {
	if x.ref3db2646c != nil {
		return *x.ref3db2646c, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *NativeWindowBuffer) Deref() {
	if x.ref3db2646c == nil {
		return
	}
	x.Width = (int32)(x.ref3db2646c.width)
	x.Height = (int32)(x.ref3db2646c.height)
	x.Stride = (int32)(x.ref3db2646c.stride)
	x.Format = (int32)(x.ref3db2646c.format)
	x.Bits = (unsafe.Pointer)(unsafe.Pointer(x.ref3db2646c.bits))
	x.Reserved = *(*[6]uint32)(unsafe.Pointer(&x.ref3db2646c.reserved))
}

// allocRectMemory allocates memory for type C.ARect in C.
// The caller is responsible for freeing the this memory via C.free.
func allocRectMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfRectValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfRectValue = unsafe.Sizeof([1]C.ARect{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Rect) Ref() *C.ARect {
	if x == nil {
		return nil
	}
	return x.ref9511c547
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Rect) Free() {
	if x != nil && x.allocs9511c547 != nil {
		x.allocs9511c547.(*cgoAllocMap).Free()
		x.ref9511c547 = nil
	}
}

// NewRectRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewRectRef(ref unsafe.Pointer) *Rect {
	if ref == nil {
		return nil
	}
	obj := new(Rect)
	obj.ref9511c547 = (*C.ARect)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Rect) PassRef() (*C.ARect, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref9511c547 != nil {
		return x.ref9511c547, nil
	}
	mem9511c547 := allocRectMemory(1)
	ref9511c547 := (*C.ARect)(mem9511c547)
	allocs9511c547 := new(cgoAllocMap)
	allocs9511c547.Add(mem9511c547)

	var cleft_allocs *cgoAllocMap
	ref9511c547.left, cleft_allocs = (C.int32_t)(x.Left), cgoAllocsUnknown
	allocs9511c547.Borrow(cleft_allocs)

	var ctop_allocs *cgoAllocMap
	ref9511c547.top, ctop_allocs = (C.int32_t)(x.Top), cgoAllocsUnknown
	allocs9511c547.Borrow(ctop_allocs)

	var cright_allocs *cgoAllocMap
	ref9511c547.right, cright_allocs = (C.int32_t)(x.Right), cgoAllocsUnknown
	allocs9511c547.Borrow(cright_allocs)

	var cbottom_allocs *cgoAllocMap
	ref9511c547.bottom, cbottom_allocs = (C.int32_t)(x.Bottom), cgoAllocsUnknown
	allocs9511c547.Borrow(cbottom_allocs)

	x.ref9511c547 = ref9511c547
	x.allocs9511c547 = allocs9511c547
	return ref9511c547, allocs9511c547

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x Rect) PassValue() (C.ARect, *cgoAllocMap) {
	if x.ref9511c547 != nil {
		return *x.ref9511c547, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Rect) Deref() {
	if x.ref9511c547 == nil {
		return
	}
	x.Left = (int32)(x.ref9511c547.left)
	x.Top = (int32)(x.ref9511c547.top)
	x.Right = (int32)(x.ref9511c547.right)
	x.Bottom = (int32)(x.ref9511c547.bottom)
}

// Ref returns a reference to C object as it is.
func (x *ObbInfo) Ref() *C.AObbInfo {
	if x == nil {
		return nil
	}
	return (*C.AObbInfo)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *ObbInfo) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewObbInfoRef converts the C object reference into a raw struct reference without wrapping.
func NewObbInfoRef(ref unsafe.Pointer) *ObbInfo {
	return (*ObbInfo)(ref)
}

// NewObbInfo allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewObbInfo() *ObbInfo {
	return (*ObbInfo)(allocObbInfoMemory(1))
}

// allocObbInfoMemory allocates memory for type C.AObbInfo in C.
// The caller is responsible for freeing the this memory via C.free.
func allocObbInfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfObbInfoValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfObbInfoValue = unsafe.Sizeof([1]C.AObbInfo{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *ObbInfo) PassRef() *C.AObbInfo {
	if x == nil {
		x = (*ObbInfo)(allocObbInfoMemory(1))
	}
	return (*C.AObbInfo)(unsafe.Pointer(x))
}

// allocSensorVectorMemory allocates memory for type C.ASensorVector in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSensorVectorMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSensorVectorValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSensorVectorValue = unsafe.Sizeof([1]C.ASensorVector{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *SensorVector) Ref() *C.ASensorVector {
	if x == nil {
		return nil
	}
	return x.refdc35e822
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *SensorVector) Free() {
	if x != nil && x.allocsdc35e822 != nil {
		x.allocsdc35e822.(*cgoAllocMap).Free()
		x.refdc35e822 = nil
	}
}

// NewSensorVectorRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewSensorVectorRef(ref unsafe.Pointer) *SensorVector {
	if ref == nil {
		return nil
	}
	obj := new(SensorVector)
	obj.refdc35e822 = (*C.ASensorVector)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *SensorVector) PassRef() (*C.ASensorVector, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refdc35e822 != nil {
		return x.refdc35e822, nil
	}
	memdc35e822 := allocSensorVectorMemory(1)
	refdc35e822 := (*C.ASensorVector)(memdc35e822)
	allocsdc35e822 := new(cgoAllocMap)
	allocsdc35e822.Add(memdc35e822)

	var cstatus_allocs *cgoAllocMap
	refdc35e822.status, cstatus_allocs = (C.int8_t)(x.Status), cgoAllocsUnknown
	allocsdc35e822.Borrow(cstatus_allocs)

	var creserved_allocs *cgoAllocMap
	refdc35e822.reserved, creserved_allocs = *(*[3]C.uint8_t)(unsafe.Pointer(&x.Reserved)), cgoAllocsUnknown
	allocsdc35e822.Borrow(creserved_allocs)

	x.refdc35e822 = refdc35e822
	x.allocsdc35e822 = allocsdc35e822
	return refdc35e822, allocsdc35e822

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x SensorVector) PassValue() (C.ASensorVector, *cgoAllocMap) {
	if x.refdc35e822 != nil {
		return *x.refdc35e822, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *SensorVector) Deref() {
	if x.refdc35e822 == nil {
		return
	}
	x.Status = (byte)(x.refdc35e822.status)
	x.Reserved = *(*[3]byte)(unsafe.Pointer(&x.refdc35e822.reserved))
}

// allocMetaDataEventMemory allocates memory for type C.AMetaDataEvent in C.
// The caller is responsible for freeing the this memory via C.free.
func allocMetaDataEventMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfMetaDataEventValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfMetaDataEventValue = unsafe.Sizeof([1]C.AMetaDataEvent{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *MetaDataEvent) Ref() *C.AMetaDataEvent {
	if x == nil {
		return nil
	}
	return x.ref4f7ec3e5
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *MetaDataEvent) Free() {
	if x != nil && x.allocs4f7ec3e5 != nil {
		x.allocs4f7ec3e5.(*cgoAllocMap).Free()
		x.ref4f7ec3e5 = nil
	}
}

// NewMetaDataEventRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewMetaDataEventRef(ref unsafe.Pointer) *MetaDataEvent {
	if ref == nil {
		return nil
	}
	obj := new(MetaDataEvent)
	obj.ref4f7ec3e5 = (*C.AMetaDataEvent)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *MetaDataEvent) PassRef() (*C.AMetaDataEvent, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref4f7ec3e5 != nil {
		return x.ref4f7ec3e5, nil
	}
	mem4f7ec3e5 := allocMetaDataEventMemory(1)
	ref4f7ec3e5 := (*C.AMetaDataEvent)(mem4f7ec3e5)
	allocs4f7ec3e5 := new(cgoAllocMap)
	allocs4f7ec3e5.Add(mem4f7ec3e5)

	var cwhat_allocs *cgoAllocMap
	ref4f7ec3e5.what, cwhat_allocs = (C.int32_t)(x.What), cgoAllocsUnknown
	allocs4f7ec3e5.Borrow(cwhat_allocs)

	var csensor_allocs *cgoAllocMap
	ref4f7ec3e5.sensor, csensor_allocs = (C.int32_t)(x.Sensor), cgoAllocsUnknown
	allocs4f7ec3e5.Borrow(csensor_allocs)

	x.ref4f7ec3e5 = ref4f7ec3e5
	x.allocs4f7ec3e5 = allocs4f7ec3e5
	return ref4f7ec3e5, allocs4f7ec3e5

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x MetaDataEvent) PassValue() (C.AMetaDataEvent, *cgoAllocMap) {
	if x.ref4f7ec3e5 != nil {
		return *x.ref4f7ec3e5, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *MetaDataEvent) Deref() {
	if x.ref4f7ec3e5 == nil {
		return
	}
	x.What = (int32)(x.ref4f7ec3e5.what)
	x.Sensor = (int32)(x.ref4f7ec3e5.sensor)
}

// allocUncalibratedEventMemory allocates memory for type C.AUncalibratedEvent in C.
// The caller is responsible for freeing the this memory via C.free.
func allocUncalibratedEventMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfUncalibratedEventValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfUncalibratedEventValue = unsafe.Sizeof([1]C.AUncalibratedEvent{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *UncalibratedEvent) Ref() *C.AUncalibratedEvent {
	if x == nil {
		return nil
	}
	return x.ref193d9a7d
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *UncalibratedEvent) Free() {
	if x != nil && x.allocs193d9a7d != nil {
		x.allocs193d9a7d.(*cgoAllocMap).Free()
		x.ref193d9a7d = nil
	}
}

// NewUncalibratedEventRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewUncalibratedEventRef(ref unsafe.Pointer) *UncalibratedEvent {
	if ref == nil {
		return nil
	}
	obj := new(UncalibratedEvent)
	obj.ref193d9a7d = (*C.AUncalibratedEvent)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *UncalibratedEvent) PassRef() (*C.AUncalibratedEvent, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref193d9a7d != nil {
		return x.ref193d9a7d, nil
	}
	mem193d9a7d := allocUncalibratedEventMemory(1)
	ref193d9a7d := (*C.AUncalibratedEvent)(mem193d9a7d)
	allocs193d9a7d := new(cgoAllocMap)
	allocs193d9a7d.Add(mem193d9a7d)

	x.ref193d9a7d = ref193d9a7d
	x.allocs193d9a7d = allocs193d9a7d
	return ref193d9a7d, allocs193d9a7d

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x UncalibratedEvent) PassValue() (C.AUncalibratedEvent, *cgoAllocMap) {
	if x.ref193d9a7d != nil {
		return *x.ref193d9a7d, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *UncalibratedEvent) Deref() {
	if x.ref193d9a7d == nil {
		return
	}
}

// allocHeartRateEventMemory allocates memory for type C.AHeartRateEvent in C.
// The caller is responsible for freeing the this memory via C.free.
func allocHeartRateEventMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfHeartRateEventValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfHeartRateEventValue = unsafe.Sizeof([1]C.AHeartRateEvent{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *HeartRateEvent) Ref() *C.AHeartRateEvent {
	if x == nil {
		return nil
	}
	return x.ref1342bcf5
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *HeartRateEvent) Free() {
	if x != nil && x.allocs1342bcf5 != nil {
		x.allocs1342bcf5.(*cgoAllocMap).Free()
		x.ref1342bcf5 = nil
	}
}

// NewHeartRateEventRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewHeartRateEventRef(ref unsafe.Pointer) *HeartRateEvent {
	if ref == nil {
		return nil
	}
	obj := new(HeartRateEvent)
	obj.ref1342bcf5 = (*C.AHeartRateEvent)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *HeartRateEvent) PassRef() (*C.AHeartRateEvent, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref1342bcf5 != nil {
		return x.ref1342bcf5, nil
	}
	mem1342bcf5 := allocHeartRateEventMemory(1)
	ref1342bcf5 := (*C.AHeartRateEvent)(mem1342bcf5)
	allocs1342bcf5 := new(cgoAllocMap)
	allocs1342bcf5.Add(mem1342bcf5)

	var cbpm_allocs *cgoAllocMap
	ref1342bcf5.bpm, cbpm_allocs = (C.float)(x.Bpm), cgoAllocsUnknown
	allocs1342bcf5.Borrow(cbpm_allocs)

	var cstatus_allocs *cgoAllocMap
	ref1342bcf5.status, cstatus_allocs = (C.int8_t)(x.Status), cgoAllocsUnknown
	allocs1342bcf5.Borrow(cstatus_allocs)

	x.ref1342bcf5 = ref1342bcf5
	x.allocs1342bcf5 = allocs1342bcf5
	return ref1342bcf5, allocs1342bcf5

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x HeartRateEvent) PassValue() (C.AHeartRateEvent, *cgoAllocMap) {
	if x.ref1342bcf5 != nil {
		return *x.ref1342bcf5, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *HeartRateEvent) Deref() {
	if x.ref1342bcf5 == nil {
		return
	}
	x.Bpm = (float32)(x.ref1342bcf5.bpm)
	x.Status = (byte)(x.ref1342bcf5.status)
}

// allocSensorEventMemory allocates memory for type C.ASensorEvent in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSensorEventMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSensorEventValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSensorEventValue = unsafe.Sizeof([1]C.ASensorEvent{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *SensorEvent) Ref() *C.ASensorEvent {
	if x == nil {
		return nil
	}
	return x.refb7a6c7dc
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *SensorEvent) Free() {
	if x != nil && x.allocsb7a6c7dc != nil {
		x.allocsb7a6c7dc.(*cgoAllocMap).Free()
		x.refb7a6c7dc = nil
	}
}

// NewSensorEventRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewSensorEventRef(ref unsafe.Pointer) *SensorEvent {
	if ref == nil {
		return nil
	}
	obj := new(SensorEvent)
	obj.refb7a6c7dc = (*C.ASensorEvent)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *SensorEvent) PassRef() (*C.ASensorEvent, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refb7a6c7dc != nil {
		return x.refb7a6c7dc, nil
	}
	memb7a6c7dc := allocSensorEventMemory(1)
	refb7a6c7dc := (*C.ASensorEvent)(memb7a6c7dc)
	allocsb7a6c7dc := new(cgoAllocMap)
	allocsb7a6c7dc.Add(memb7a6c7dc)

	var cversion_allocs *cgoAllocMap
	refb7a6c7dc.version, cversion_allocs = (C.int32_t)(x.Version), cgoAllocsUnknown
	allocsb7a6c7dc.Borrow(cversion_allocs)

	var csensor_allocs *cgoAllocMap
	refb7a6c7dc.sensor, csensor_allocs = (C.int32_t)(x.Sensor), cgoAllocsUnknown
	allocsb7a6c7dc.Borrow(csensor_allocs)

	var c_type_allocs *cgoAllocMap
	refb7a6c7dc._type, c_type_allocs = (C.int32_t)(x.Type), cgoAllocsUnknown
	allocsb7a6c7dc.Borrow(c_type_allocs)

	var creserved0_allocs *cgoAllocMap
	refb7a6c7dc.reserved0, creserved0_allocs = (C.int32_t)(x.Reserved0), cgoAllocsUnknown
	allocsb7a6c7dc.Borrow(creserved0_allocs)

	var ctimestamp_allocs *cgoAllocMap
	refb7a6c7dc.timestamp, ctimestamp_allocs = (C.int64_t)(x.Timestamp), cgoAllocsUnknown
	allocsb7a6c7dc.Borrow(ctimestamp_allocs)

	var cflags_allocs *cgoAllocMap
	refb7a6c7dc.flags, cflags_allocs = (C.uint32_t)(x.Flags), cgoAllocsUnknown
	allocsb7a6c7dc.Borrow(cflags_allocs)

	var creserved1_allocs *cgoAllocMap
	refb7a6c7dc.reserved1, creserved1_allocs = *(*[3]C.int32_t)(unsafe.Pointer(&x.Reserved1)), cgoAllocsUnknown
	allocsb7a6c7dc.Borrow(creserved1_allocs)

	x.refb7a6c7dc = refb7a6c7dc
	x.allocsb7a6c7dc = allocsb7a6c7dc
	return refb7a6c7dc, allocsb7a6c7dc

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x SensorEvent) PassValue() (C.ASensorEvent, *cgoAllocMap) {
	if x.refb7a6c7dc != nil {
		return *x.refb7a6c7dc, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *SensorEvent) Deref() {
	if x.refb7a6c7dc == nil {
		return
	}
	x.Version = (int32)(x.refb7a6c7dc.version)
	x.Sensor = (int32)(x.refb7a6c7dc.sensor)
	x.Type = (int32)(x.refb7a6c7dc._type)
	x.Reserved0 = (int32)(x.refb7a6c7dc.reserved0)
	x.Timestamp = (int64)(x.refb7a6c7dc.timestamp)
	x.Flags = (uint32)(x.refb7a6c7dc.flags)
	x.Reserved1 = *(*[3]int32)(unsafe.Pointer(&x.refb7a6c7dc.reserved1))
}

// Ref returns a reference to C object as it is.
func (x *SensorManager) Ref() *C.ASensorManager {
	if x == nil {
		return nil
	}
	return (*C.ASensorManager)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *SensorManager) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewSensorManagerRef converts the C object reference into a raw struct reference without wrapping.
func NewSensorManagerRef(ref unsafe.Pointer) *SensorManager {
	return (*SensorManager)(ref)
}

// NewSensorManager allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewSensorManager() *SensorManager {
	return (*SensorManager)(allocSensorManagerMemory(1))
}

// allocSensorManagerMemory allocates memory for type C.ASensorManager in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSensorManagerMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSensorManagerValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSensorManagerValue = unsafe.Sizeof([1]C.ASensorManager{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *SensorManager) PassRef() *C.ASensorManager {
	if x == nil {
		x = (*SensorManager)(allocSensorManagerMemory(1))
	}
	return (*C.ASensorManager)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *SensorEventQueue) Ref() *C.ASensorEventQueue {
	if x == nil {
		return nil
	}
	return (*C.ASensorEventQueue)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *SensorEventQueue) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewSensorEventQueueRef converts the C object reference into a raw struct reference without wrapping.
func NewSensorEventQueueRef(ref unsafe.Pointer) *SensorEventQueue {
	return (*SensorEventQueue)(ref)
}

// NewSensorEventQueue allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewSensorEventQueue() *SensorEventQueue {
	return (*SensorEventQueue)(allocSensorEventQueueMemory(1))
}

// allocSensorEventQueueMemory allocates memory for type C.ASensorEventQueue in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSensorEventQueueMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSensorEventQueueValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSensorEventQueueValue = unsafe.Sizeof([1]C.ASensorEventQueue{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *SensorEventQueue) PassRef() *C.ASensorEventQueue {
	if x == nil {
		x = (*SensorEventQueue)(allocSensorEventQueueMemory(1))
	}
	return (*C.ASensorEventQueue)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *Sensor) Ref() *C.ASensor {
	if x == nil {
		return nil
	}
	return (*C.ASensor)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *Sensor) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewSensorRef converts the C object reference into a raw struct reference without wrapping.
func NewSensorRef(ref unsafe.Pointer) *Sensor {
	return (*Sensor)(ref)
}

// NewSensor allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewSensor() *Sensor {
	return (*Sensor)(allocSensorMemory(1))
}

// allocSensorMemory allocates memory for type C.ASensor in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSensorMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSensorValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSensorValue = unsafe.Sizeof([1]C.ASensor{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *Sensor) PassRef() *C.ASensor {
	if x == nil {
		x = (*Sensor)(allocSensorMemory(1))
	}
	return (*C.ASensor)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *StorageManager) Ref() *C.AStorageManager {
	if x == nil {
		return nil
	}
	return (*C.AStorageManager)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *StorageManager) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewStorageManagerRef converts the C object reference into a raw struct reference without wrapping.
func NewStorageManagerRef(ref unsafe.Pointer) *StorageManager {
	return (*StorageManager)(ref)
}

// NewStorageManager allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewStorageManager() *StorageManager {
	return (*StorageManager)(allocStorageManagerMemory(1))
}

// allocStorageManagerMemory allocates memory for type C.AStorageManager in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStorageManagerMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStorageManagerValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStorageManagerValue = unsafe.Sizeof([1]C.AStorageManager{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *StorageManager) PassRef() *C.AStorageManager {
	if x == nil {
		x = (*StorageManager)(allocStorageManagerMemory(1))
	}
	return (*C.AStorageManager)(unsafe.Pointer(x))
}

func (x StorageManagerObbCallbackFunc) PassRef() (ref *C.AStorageManager_obbCallbackFunc, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if storageManagerObbCallbackFuncE8B15C3EFunc == nil {
		storageManagerObbCallbackFuncE8B15C3EFunc = x
	}
	return (*C.AStorageManager_obbCallbackFunc)(C.AStorageManager_obbCallbackFunc_e8b15c3e), nil
}

func (x StorageManagerObbCallbackFunc) PassValue() (ref C.AStorageManager_obbCallbackFunc, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if storageManagerObbCallbackFuncE8B15C3EFunc == nil {
		storageManagerObbCallbackFuncE8B15C3EFunc = x
	}
	return (C.AStorageManager_obbCallbackFunc)(C.AStorageManager_obbCallbackFunc_e8b15c3e), nil
}

func NewStorageManagerObbCallbackFuncRef(ref unsafe.Pointer) *StorageManagerObbCallbackFunc {
	return (*StorageManagerObbCallbackFunc)(ref)
}

//export storageManagerObbCallbackFuncE8B15C3E
func storageManagerObbCallbackFuncE8B15C3E(cfilename *C.char, cstate C.int32_t, cdata unsafe.Pointer) {
	if storageManagerObbCallbackFuncE8B15C3EFunc != nil {
		filenamee8b15c3e := packPCharString(cfilename)
		statee8b15c3e := (int32)(cstate)
		datae8b15c3e := (unsafe.Pointer)(unsafe.Pointer(cdata))
		storageManagerObbCallbackFuncE8B15C3EFunc(filenamee8b15c3e, statee8b15c3e, datae8b15c3e)
		return
	}
	panic("callback func has not been set (race?)")
}

var storageManagerObbCallbackFuncE8B15C3EFunc StorageManagerObbCallbackFunc

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
