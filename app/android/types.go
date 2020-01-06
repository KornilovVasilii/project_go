package android

import "C"
import "unsafe"

// AssetManager as declared in android/asset_manager.h:28
type AssetManager C.AAssetManager

// AssetDir as declared in android/asset_manager.h:31
type AssetDir C.AAssetDir

// Asset as declared in android/asset_manager.h:34
type Asset C.AAsset

// Jboolean type as declared in include/jni.h:44
type Jboolean byte

// Jbyte type as declared in include/jni.h:45
type Jbyte byte

// Jchar type as declared in include/jni.h:46
type Jchar uint16

// Jshort type as declared in include/jni.h:47
type Jshort int16

// Jint type as declared in include/jni.h:48
type Jint int32

// Jlong type as declared in include/jni.h:49
type Jlong int64

// Jfloat type as declared in include/jni.h:50
type Jfloat float32

// Jdouble type as declared in include/jni.h:51
type Jdouble float64

// Jsize type as declared in include/jni.h:55
type Jsize int32

// Jobject type as declared in include/jni.h:98
type Jobject unsafe.Pointer

// Jclass type as declared in include/jni.h:99
type Jclass unsafe.Pointer

// Jstring type as declared in include/jni.h:100
type Jstring unsafe.Pointer

// Jarray type as declared in include/jni.h:101
type Jarray unsafe.Pointer

// JobjectArray type as declared in include/jni.h:102
type JobjectArray unsafe.Pointer

// JbooleanArray type as declared in include/jni.h:103
type JbooleanArray unsafe.Pointer

// JbyteArray type as declared in include/jni.h:104
type JbyteArray unsafe.Pointer

// JcharArray type as declared in include/jni.h:105
type JcharArray unsafe.Pointer

// JshortArray type as declared in include/jni.h:106
type JshortArray unsafe.Pointer

// JintArray type as declared in include/jni.h:107
type JintArray unsafe.Pointer

// JlongArray type as declared in include/jni.h:108
type JlongArray unsafe.Pointer

// JfloatArray type as declared in include/jni.h:109
type JfloatArray unsafe.Pointer

// JdoubleArray type as declared in include/jni.h:110
type JdoubleArray unsafe.Pointer

// Jthrowable type as declared in include/jni.h:111
type Jthrowable unsafe.Pointer

// Jweak type as declared in include/jni.h:112
type Jweak unsafe.Pointer

// JfieldID as declared in include/jni.h:117
type JfieldID C.jfieldID

// JmethodID as declared in include/jni.h:120
type JmethodID C.jmethodID

// Jvalue as declared in include/jni.h:134
const sizeofJvalue = unsafe.Sizeof(C.jvalue{})

type Jvalue [sizeofJvalue]byte

// JNINativeMethod as declared in include/jni.h:147
type JNINativeMethod C.JNINativeMethod

// JNIEnv as declared in include/jni.h:157
type JNIEnv C.JNIEnv

// JavaVM as declared in include/jni.h:158
type JavaVM C.JavaVM

// JavaVMAttachArgs as declared in include/jni.h:1106
type JavaVMAttachArgs struct {
	Version        Jint
	Name           string
	Group          Jobject
	refab42f020    *C.JavaVMAttachArgs
	allocsab42f020 interface{}
}

// JavaVMOption as declared in include/jni.h:1115
type JavaVMOption C.JavaVMOption

// JavaVMInitArgs as declared in include/jni.h:1123
type JavaVMInitArgs C.JavaVMInitArgs

// Configuration as declared in android/configuration.h:27
type Configuration C.AConfiguration

// Looper as declared in android/looper.h:39
type Looper C.ALooper

// LooperCallbackFunc type as declared in android/looper.h:159
type LooperCallbackFunc func(fd int32, events int32, data unsafe.Pointer) int32

// NativeActivity as declared in android/native_activity.h:108
type NativeActivity struct {
	Callbacks        *NativeActivityCallbacks
	Vm               *JavaVM
	Env              *JNIEnv
	Clazz            Jobject
	InternalDataPath string
	ExternalDataPath string
	SdkVersion       int32
	Instance         unsafe.Pointer
	AssetManager     *AssetManager
	ObbPath          string
	ref2cc295bf      *C.ANativeActivity
	allocs2cc295bf   interface{}
}

// NativeActivityCallbacks as declared in android/native_activity.h:225
type NativeActivityCallbacks C.ANativeActivityCallbacks

// NativeActivityCreateFunc type as declared in android/native_activity.h:235
type NativeActivityCreateFunc func(activity *NativeActivity, savedState unsafe.Pointer, savedStateSize uint32)

// InputEvent as declared in android/input.h:136
type InputEvent C.AInputEvent

// InputQueue as declared in android/input.h:803
type InputQueue C.AInputQueue

// NativeWindow as declared in android/native_window.h:36
type NativeWindow C.ANativeWindow

// NativeWindowBuffer as declared in android/native_window.h:57
type NativeWindowBuffer struct {
	Width          int32
	Height         int32
	Stride         int32
	Format         int32
	Bits           unsafe.Pointer
	Reserved       [6]uint32
	ref3db2646c    *C.ANativeWindow_Buffer
	allocs3db2646c interface{}
}

// Rect as declared in android/rect.h:35
type Rect struct {
	Left           int32
	Top            int32
	Right          int32
	Bottom         int32
	ref9511c547    *C.ARect
	allocs9511c547 interface{}
}

// ObbInfo as declared in android/obb.h:28
type ObbInfo C.AObbInfo

// SensorVector as declared in android/sensor.h:119
type SensorVector struct {
	Status         byte
	Reserved       [3]byte
	refdc35e822    *C.ASensorVector
	allocsdc35e822 interface{}
}

// MetaDataEvent as declared in android/sensor.h:124
type MetaDataEvent struct {
	What           int32
	Sensor         int32
	ref4f7ec3e5    *C.AMetaDataEvent
	allocs4f7ec3e5 interface{}
}

// UncalibratedEvent as declared in android/sensor.h:143
type UncalibratedEvent struct {
	ref193d9a7d    *C.AUncalibratedEvent
	allocs193d9a7d interface{}
}

// HeartRateEvent as declared in android/sensor.h:148
type HeartRateEvent struct {
	Bpm            float32
	Status         byte
	ref1342bcf5    *C.AHeartRateEvent
	allocs1342bcf5 interface{}
}

// SensorEvent as declared in android/sensor.h:181
type SensorEvent struct {
	Version        int32
	Sensor         int32
	Type           int32
	Reserved0      int32
	Timestamp      int64
	Flags          uint32
	Reserved1      [3]int32
	refb7a6c7dc    *C.ASensorEvent
	allocsb7a6c7dc interface{}
}

// SensorManager as declared in android/sensor.h:184
type SensorManager C.ASensorManager

// SensorEventQueue as declared in android/sensor.h:187
type SensorEventQueue C.ASensorEventQueue

// Sensor as declared in android/sensor.h:190
type Sensor C.ASensor

// SensorRef as declared in android/sensor.h:191
type SensorRef C.ASensor

// SensorList as declared in android/sensor.h:192
type SensorList C.ASensorRef

// StorageManager as declared in android/storage_manager.h:28
type StorageManager C.AStorageManager

// StorageManagerObbCallbackFunc type as declared in android/storage_manager.h:98
type StorageManagerObbCallbackFunc func(filename string, state int32, data unsafe.Pointer)
