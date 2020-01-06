#include <android/api-level.h>
#include <android/asset_manager.h>
#include <android/asset_manager_jni.h>
#include <android/configuration.h>
#include <android/input.h>
#include <android/keycodes.h>
#include <android/log.h>
#include <android/looper.h>
#include <android/native_activity.h>
#include <android/native_window.h>
#include <android/native_window_jni.h>
#include <android/obb.h>
#include <android/rect.h>
#include <android/sensor.h>
#include <android/storage_manager.h>
#include <android/window.h>
#include <jni.h>
#include "jni_call.h"
#include <stdlib.h>
#pragma once

#define __CGOGEN 1

// ALooper_callbackFunc_7e6b484c is a proxy for callback ALooper_callbackFunc.
int ALooper_callbackFunc_7e6b484c(int fd, int events, void* data);

// ANativeActivity_createFunc_76dce4 is a proxy for callback ANativeActivity_createFunc.
void ANativeActivity_createFunc_76dce4(ANativeActivity* activity, void* savedState, unsigned int savedStateSize);

// AStorageManager_obbCallbackFunc_e8b15c3e is a proxy for callback AStorageManager_obbCallbackFunc.
void AStorageManager_obbCallbackFunc_e8b15c3e(char* filename, int state, void* data);

