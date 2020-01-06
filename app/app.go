package app

import "C"

import (
	"os"
	"sync"
	"time"
	"unsafe"

	"android"
	"internal/callfn"
)

func callMain(mainPC uintptr) {
	for _, name := range []string{"TMPDIR", "PATH", "LD_LIBRARY_PATH"} {
		n := C.CString(name)
		os.Setenv(name, C.GoString(C.getenv(n)))
		C.free(unsafe.Pointer(n))
	}

	var curtime C.time_t
	var curtm C.struct_tm
	C.time(&curtime)
	C.localtime_r(&curtime, &curtm)
	tzOffset := int(curtm.tm_gmtoff)
	tz := C.GoString(curtm.tm_zone)
	time.Local = time.FixedZone(tz, tzOffset)

	go callfn.CallFn(mainPC)
}

type NativeActivity interface {
	InitDone()
	NativeWindowRedrawDone()
	InputQueueHandled()
	NativeActivity() *android.NativeActivity
	LifecycleEvents() <-chan LifecycleEvent
	HandleSaveInstanceState(fn SaveStateFunc)
	HandleWindowFocusEvents(out chan<- WindowFocusEvent)
	HandleNativeWindowEvents(out chan<- NativeWindowEvent)
	HandleInputQueueEvents(out chan<- InputQueueEvent)
	HandleContentRectEvents(out chan<- ContentRectEvent)
	HandleActivityEvents(out chan<- ActivityEvent)
	GetAsset(name string) ([]byte, error)
	OpenAsset(name string) (*AssetReader, error)
}

var defaultApp = &nativeActivity{
	lifecycleEvents:        make(chan LifecycleEvent),
	nativeWindowRedrawDone: make(chan Signal, 1),
	inputQueueHandled:      make(chan Signal, 1),
	maxDispatchTime:        1 * time.Second,

	initWG: new(sync.WaitGroup),
	mux:    new(sync.RWMutex),
}

func init() {
	defaultApp.initWG.Add(1)
}

func Main(fn func(app NativeActivity)) {
	fn(defaultApp)
}

type nativeActivity struct {
	activity *android.NativeActivity

	lifecycleEvents chan LifecycleEvent

	maxDispatchTime time.Duration

	windowFocusEvents  chan<- WindowFocusEvent
	nativeWindowEvents chan<- NativeWindowEvent
	inputQueueEvents   chan<- InputQueueEvent
	contentRectEvents  chan<- ContentRectEvent
	activityEvents     chan<- ActivityEvent

	saveInstanceStateFunc  SaveStateFunc
	nativeWindowRedrawDone chan Signal
	inputQueueHandled      chan Signal

	initWG *sync.WaitGroup
	mux    *sync.RWMutex
}

func (a *nativeActivity) InitDone() {
	a.initWG.Done()
}

func (a *nativeActivity) NativeActivity() *android.NativeActivity {
	a.mux.RLock()
	activity := a.activity
	a.mux.RUnlock()
	return activity
}

func (a *nativeActivity) LifecycleEvents() <-chan LifecycleEvent {
	return a.lifecycleEvents
}

type Signal struct{}

func (a *nativeActivity) NativeWindowRedrawDone() {
	select {
	case a.nativeWindowRedrawDone <- Signal{}:
	default:
	}
}

func (a *nativeActivity) InputQueueHandled() {
	select {
	case a.inputQueueHandled <- Signal{}:
	default:
	}
}

func (a *nativeActivity) HandleWindowFocusEvents(out chan<- WindowFocusEvent) {
	a.mux.Lock()
	a.windowFocusEvents = out
	a.mux.Unlock()
}

func (a *nativeActivity) getWindowFocusEventsOut() chan<- WindowFocusEvent {
	a.mux.RLock()
	out := a.windowFocusEvents
	a.mux.RUnlock()
	return out
}

func (a *nativeActivity) HandleNativeWindowEvents(out chan<- NativeWindowEvent) {
	a.mux.Lock()
	a.nativeWindowEvents = out
	a.mux.Unlock()
}

func (a *nativeActivity) getNativeWindowEventsOut() chan<- NativeWindowEvent {
	a.mux.RLock()
	out := a.nativeWindowEvents
	a.mux.RUnlock()
	return out
}

func (a *nativeActivity) HandleInputQueueEvents(out chan<- InputQueueEvent) {
	a.mux.Lock()
	a.inputQueueEvents = out
	a.mux.Unlock()
}

func (a *nativeActivity) getInputQueueEventsOut() chan<- InputQueueEvent {
	a.mux.RLock()
	out := a.inputQueueEvents
	a.mux.RUnlock()
	return out
}

func (a *nativeActivity) HandleContentRectEvents(out chan<- ContentRectEvent) {
	a.mux.Lock()
	a.contentRectEvents = out
	a.mux.Unlock()
}

func (a *nativeActivity) getContentRectEventsOut() chan<- ContentRectEvent {
	a.mux.RLock()
	out := a.contentRectEvents
	a.mux.RUnlock()
	return out
}

func (a *nativeActivity) HandleActivityEvents(out chan<- ActivityEvent) {
	a.mux.Lock()
	a.activityEvents = out
	a.mux.Unlock()
}

func (a *nativeActivity) getActivityEventsOut() chan<- ActivityEvent {
	a.mux.RLock()
	out := a.activityEvents
	a.mux.RUnlock()
	return out
}

func (a *nativeActivity) HandleSaveInstanceState(fn SaveStateFunc) {
	a.mux.Lock()
	a.saveInstanceStateFunc = fn
	a.mux.Unlock()
}

func (a *nativeActivity) getSaveInstanceStateFunc() SaveStateFunc {
	a.mux.RLock()
	fn := a.saveInstanceStateFunc
	a.mux.RUnlock()
	return fn
}
