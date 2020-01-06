package app

import "C"

import (
	"time"
	"unsafe"

	"android"
)

type LifecycleEvent struct {
	Activity *android.NativeActivity
	Kind     LifecycleEventKind
}

type LifecycleEventKind string

const (
	OnCreate  LifecycleEventKind = "onCreate"
	OnDestroy LifecycleEventKind = "onDestroy"
	OnStart   LifecycleEventKind = "onStart"
	OnStop    LifecycleEventKind = "onStop"
	OnPause   LifecycleEventKind = "onPause"
	OnResume  LifecycleEventKind = "onResume"
)

func onCreate(activity *C.ANativeActivity) {
	defaultApp.initWG.Wait()

	defaultApp.mux.Lock()
	defaultApp.activity = android.NewNativeActivityRef(unsafe.Pointer(activity))
	defaultApp.activity.Deref()
	defaultApp.mux.Unlock()

	event := LifecycleEvent{
		Activity: defaultApp.activity,
		Kind:     OnCreate,
	}
	defaultApp.lifecycleEvents <- event
}

func onDestroy(activity *C.ANativeActivity) {
	defaultApp.initWG.Wait()

	event := LifecycleEvent{
		Activity: android.NewNativeActivityRef(unsafe.Pointer(activity)),
		Kind:     OnDestroy,
	}
	defaultApp.lifecycleEvents <- event
}

func onStart(activity *C.ANativeActivity) {
	defaultApp.initWG.Wait()

	event := LifecycleEvent{
		Activity: android.NewNativeActivityRef(unsafe.Pointer(activity)),
		Kind:     OnStart,
	}
	defaultApp.lifecycleEvents <- event
}

func onStop(activity *C.ANativeActivity) {
	defaultApp.initWG.Wait()

	event := LifecycleEvent{
		Activity: android.NewNativeActivityRef(unsafe.Pointer(activity)),
		Kind:     OnStop,
	}
	defaultApp.lifecycleEvents <- event
}

func onPause(activity *C.ANativeActivity) {
	defaultApp.initWG.Wait()

	event := LifecycleEvent{
		Activity: android.NewNativeActivityRef(unsafe.Pointer(activity)),
		Kind:     OnPause,
	}
	defaultApp.lifecycleEvents <- event
}

func onResume(activity *C.ANativeActivity) {
	defaultApp.initWG.Wait()

	event := LifecycleEvent{
		Activity: android.NewNativeActivityRef(unsafe.Pointer(activity)),
		Kind:     OnResume,
	}
	defaultApp.lifecycleEvents <- event
}

type SaveStateFunc func(activity *android.NativeActivity, size *uint32) unsafe.Pointer

func onSaveInstanceState(activity *C.ANativeActivity, outSize *C.size_t) unsafe.Pointer {
	defaultApp.initWG.Wait()

	fn := defaultApp.getSaveInstanceStateFunc()
	if fn == nil {
		return nil
	}
	activityRef := android.NewNativeActivityRef(unsafe.Pointer(activity))
	result := fn(activityRef, (*uint32)(outSize))
	return result
}

type WindowFocusEvent struct {
	Activity *android.NativeActivity
	HasFocus bool
}

func onWindowFocusChanged(activity *C.ANativeActivity, hasFocus int) {
	defaultApp.initWG.Wait()

	out := defaultApp.getWindowFocusEventsOut()
	if out == nil {
		return
	}
	event := WindowFocusEvent{
		Activity: android.NewNativeActivityRef(unsafe.Pointer(activity)),
		HasFocus: hasFocus > 0,
	}
	select {
	case out <- event:
	case <-time.After(defaultApp.maxDispatchTime):
	}
}

type NativeWindowEvent struct {
	Activity *android.NativeActivity
	Window   *android.NativeWindow
	Kind     NativeWindowEventKind
}

type NativeWindowEventKind string

const (
	NativeWindowCreated      NativeWindowEventKind = "nativeWindowCreated"
	NativeWindowRedrawNeeded NativeWindowEventKind = "nativeWindowRedrawNeeded"
	NativeWindowDestroyed    NativeWindowEventKind = "nativeWindowDestroyed"
)

func onNativeWindowCreated(activity *C.ANativeActivity, window *C.ANativeWindow) {
	defaultApp.initWG.Wait()

	out := defaultApp.getNativeWindowEventsOut()
	if out == nil {
		return
	}
	event := NativeWindowEvent{
		Activity: android.NewNativeActivityRef(unsafe.Pointer(activity)),
		Window:   (*android.NativeWindow)(window),
		Kind:     NativeWindowCreated,
	}
	select {
	case out <- event:
	case <-time.After(defaultApp.maxDispatchTime):
	}
}

func onNativeWindowRedrawNeeded(activity *C.ANativeActivity, window *C.ANativeWindow) {
	defaultApp.initWG.Wait()

	out := defaultApp.getNativeWindowEventsOut()
	if out == nil {
		return
	}
	event := NativeWindowEvent{
		Activity: android.NewNativeActivityRef(unsafe.Pointer(activity)),
		Window:   (*android.NativeWindow)(window),
		Kind:     NativeWindowRedrawNeeded,
	}
	select {
	case <-defaultApp.nativeWindowRedrawDone:
	default:
	}
	select {
	case out <- event:
	case <-time.After(defaultApp.maxDispatchTime):
		return
	}

	<-defaultApp.nativeWindowRedrawDone
}

func onNativeWindowDestroyed(activity *C.ANativeActivity, window *C.ANativeWindow) {
	defaultApp.initWG.Wait()

	out := defaultApp.getNativeWindowEventsOut()
	if out == nil {
		return
	}
	event := NativeWindowEvent{
		Activity: android.NewNativeActivityRef(unsafe.Pointer(activity)),
		Window:   (*android.NativeWindow)(window),
		Kind:     NativeWindowDestroyed,
	}
	select {
	case out <- event:
	case <-time.After(defaultApp.maxDispatchTime):
	}
}

type InputQueueEvent struct {
	Activity *android.NativeActivity
	Queue    *android.InputQueue
	Kind     InputQueueEventKind
}

type InputQueueEventKind string

const (
	QueueCreated   InputQueueEventKind = "queueCreated"
	QueueDestroyed InputQueueEventKind = "queueDestroyed"
)

func onInputQueueCreated(activity *C.ANativeActivity, queue *C.AInputQueue) {
	defaultApp.initWG.Wait()

	out := defaultApp.getInputQueueEventsOut()
	if out == nil {
		return
	}
	event := InputQueueEvent{
		Activity: android.NewNativeActivityRef(unsafe.Pointer(activity)),
		Queue:    (*android.InputQueue)(queue),
		Kind:     QueueCreated,
	}
	select {
	case out <- event:
	case <-time.After(defaultApp.maxDispatchTime):
		return
	}

	<-defaultApp.inputQueueHandled
}

func onInputQueueDestroyed(activity *C.ANativeActivity, queue *C.AInputQueue) {
	defaultApp.initWG.Wait()

	out := defaultApp.getInputQueueEventsOut()
	if out == nil {
		return
	}
	event := InputQueueEvent{
		Activity: android.NewNativeActivityRef(unsafe.Pointer(activity)),
		Queue:    (*android.InputQueue)(queue),
		Kind:     QueueDestroyed,
	}
	select {
	case out <- event:
	case <-time.After(defaultApp.maxDispatchTime):
		return
	}

	<-defaultApp.inputQueueHandled
}

type ContentRectEvent struct {
	Activity *android.NativeActivity
	Rect     *android.Rect
}

func onContentRectChanged(activity *C.ANativeActivity, rect *C.ARect) {
	defaultApp.initWG.Wait()

	out := defaultApp.getContentRectEventsOut()
	if out == nil {
		return
	}
	event := ContentRectEvent{
		Activity: android.NewNativeActivityRef(unsafe.Pointer(activity)),
		Rect:     android.NewRectRef(unsafe.Pointer(rect)),
	}
	select {
	case out <- event:
	case <-time.After(defaultApp.maxDispatchTime):
	}
}

type ActivityEvent struct {
	Activity *android.NativeActivity
	Kind     ActivityEventKind
}

type ActivityEventKind string

const (
	OnConfigurationChanged ActivityEventKind = "onConfigurationChanged"
	OnLowMemory            ActivityEventKind = "onLowMemory"
)

func onConfigurationChanged(activity *C.ANativeActivity) {
	defaultApp.initWG.Wait()

	out := defaultApp.getActivityEventsOut()
	if out == nil {
		return
	}
	event := ActivityEvent{
		Activity: android.NewNativeActivityRef(unsafe.Pointer(activity)),
		Kind:     OnConfigurationChanged,
	}
	select {
	case out <- event:
	case <-time.After(defaultApp.maxDispatchTime):
	}
}

func onLowMemory(activity *C.ANativeActivity) {
	defaultApp.initWG.Wait()

	out := defaultApp.getActivityEventsOut()
	if out == nil {
		return
	}
	event := ActivityEvent{
		Activity: android.NewNativeActivityRef(unsafe.Pointer(activity)),
		Kind:     OnLowMemory,
	}
	select {
	case out <- event:
	case <-time.After(defaultApp.maxDispatchTime):
	}
}
