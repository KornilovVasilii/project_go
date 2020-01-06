package android

import "C"
import "unsafe"

// AssetManagerOpenDir function as declared in android/asset_manager.h:52
func AssetManagerOpenDir(mgr *AssetManager, dirName string) *AssetDir {
	cmgr, _ := (*C.AAssetManager)(unsafe.Pointer(mgr)), cgoAllocsUnknown
	cdirName, _ := unpackPCharString(dirName)
	__ret := C.AAssetManager_openDir(cmgr, cdirName)
	__v := *(**AssetDir)(unsafe.Pointer(&__ret))
	return __v
}

// AssetManagerOpen function as declared in android/asset_manager.h:59
func AssetManagerOpen(mgr *AssetManager, filename string, mode int32) *Asset {
	cmgr, _ := (*C.AAssetManager)(unsafe.Pointer(mgr)), cgoAllocsUnknown
	cfilename, _ := unpackPCharString(filename)
	cmode, _ := (C.int)(mode), cgoAllocsUnknown
	__ret := C.AAssetManager_open(cmgr, cfilename, cmode)
	__v := *(**Asset)(unsafe.Pointer(&__ret))
	return __v
}

// AssetDirGetNextFileName function as declared in android/asset_manager.h:71
func AssetDirGetNextFileName(assetDir *AssetDir) string {
	cassetDir, _ := (*C.AAssetDir)(unsafe.Pointer(assetDir)), cgoAllocsUnknown
	__ret := C.AAssetDir_getNextFileName(cassetDir)
	__v := packPCharString(__ret)
	return __v
}

// AssetDirRewind function as declared in android/asset_manager.h:76
func AssetDirRewind(assetDir *AssetDir) {
	cassetDir, _ := (*C.AAssetDir)(unsafe.Pointer(assetDir)), cgoAllocsUnknown
	C.AAssetDir_rewind(cassetDir)
}

// AssetDirClose function as declared in android/asset_manager.h:81
func AssetDirClose(assetDir *AssetDir) {
	cassetDir, _ := (*C.AAssetDir)(unsafe.Pointer(assetDir)), cgoAllocsUnknown
	C.AAssetDir_close(cassetDir)
}

// AssetRead function as declared in android/asset_manager.h:88
func AssetRead(asset *Asset, buf unsafe.Pointer, count uint32) int32 {
	casset, _ := (*C.AAsset)(unsafe.Pointer(asset)), cgoAllocsUnknown
	cbuf, _ := buf, cgoAllocsUnknown
	ccount, _ := (C.size_t)(count), cgoAllocsUnknown
	__ret := C.AAsset_read(casset, cbuf, ccount)
	__v := (int32)(__ret)
	return __v
}

// AssetSeek function as declared in android/asset_manager.h:96
func AssetSeek(asset *Asset, offset int, whence int32) int {
	casset, _ := (*C.AAsset)(unsafe.Pointer(asset)), cgoAllocsUnknown
	coffset, _ := (C.off_t)(offset), cgoAllocsUnknown
	cwhence, _ := (C.int)(whence), cgoAllocsUnknown
	__ret := C.AAsset_seek(casset, coffset, cwhence)
	__v := (int)(__ret)
	return __v
}

// AssetSeek64 function as declared in android/asset_manager.h:107
func AssetSeek64(asset *Asset, offset int64, whence int32) int64 {
	casset, _ := (*C.AAsset)(unsafe.Pointer(asset)), cgoAllocsUnknown
	coffset, _ := (C.off64_t)(offset), cgoAllocsUnknown
	cwhence, _ := (C.int)(whence), cgoAllocsUnknown
	__ret := C.AAsset_seek64(casset, coffset, cwhence)
	__v := (int64)(__ret)
	return __v
}

// AssetClose function as declared in android/asset_manager.h:112
func AssetClose(asset *Asset) {
	casset, _ := (*C.AAsset)(unsafe.Pointer(asset)), cgoAllocsUnknown
	C.AAsset_close(casset)
}

// AssetGetBuffer function as declared in android/asset_manager.h:119
func AssetGetBuffer(asset *Asset) unsafe.Pointer {
	casset, _ := (*C.AAsset)(unsafe.Pointer(asset)), cgoAllocsUnknown
	__ret := C.AAsset_getBuffer(casset)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// AssetGetLength function as declared in android/asset_manager.h:124
func AssetGetLength(asset *Asset) int {
	casset, _ := (*C.AAsset)(unsafe.Pointer(asset)), cgoAllocsUnknown
	__ret := C.AAsset_getLength(casset)
	__v := (int)(__ret)
	return __v
}

// AssetGetLength64 function as declared in android/asset_manager.h:130
func AssetGetLength64(asset *Asset) int64 {
	casset, _ := (*C.AAsset)(unsafe.Pointer(asset)), cgoAllocsUnknown
	__ret := C.AAsset_getLength64(casset)
	__v := (int64)(__ret)
	return __v
}

// AssetGetRemainingLength function as declared in android/asset_manager.h:135
func AssetGetRemainingLength(asset *Asset) int {
	casset, _ := (*C.AAsset)(unsafe.Pointer(asset)), cgoAllocsUnknown
	__ret := C.AAsset_getRemainingLength(casset)
	__v := (int)(__ret)
	return __v
}

// AssetGetRemainingLength64 function as declared in android/asset_manager.h:142
func AssetGetRemainingLength64(asset *Asset) int64 {
	casset, _ := (*C.AAsset)(unsafe.Pointer(asset)), cgoAllocsUnknown
	__ret := C.AAsset_getRemainingLength64(casset)
	__v := (int64)(__ret)
	return __v
}

// AssetOpenFileDescriptor function as declared in android/asset_manager.h:152
func AssetOpenFileDescriptor(asset *Asset, outStart *int, outLength *int) int32 {
	casset, _ := (*C.AAsset)(unsafe.Pointer(asset)), cgoAllocsUnknown
	coutStart, _ := (*C.off_t)(unsafe.Pointer(outStart)), cgoAllocsUnknown
	coutLength, _ := (*C.off_t)(unsafe.Pointer(outLength)), cgoAllocsUnknown
	__ret := C.AAsset_openFileDescriptor(casset, coutStart, coutLength)
	__v := (int32)(__ret)
	return __v
}

// AssetOpenFileDescriptor64 function as declared in android/asset_manager.h:163
func AssetOpenFileDescriptor64(asset *Asset, outStart *int64, outLength *int64) int32 {
	casset, _ := (*C.AAsset)(unsafe.Pointer(asset)), cgoAllocsUnknown
	coutStart, _ := (*C.off64_t)(unsafe.Pointer(outStart)), cgoAllocsUnknown
	coutLength, _ := (*C.off64_t)(unsafe.Pointer(outLength)), cgoAllocsUnknown
	__ret := C.AAsset_openFileDescriptor64(casset, coutStart, coutLength)
	__v := (int32)(__ret)
	return __v
}

// AssetIsAllocated function as declared in android/asset_manager.h:169
func AssetIsAllocated(asset *Asset) int32 {
	casset, _ := (*C.AAsset)(unsafe.Pointer(asset)), cgoAllocsUnknown
	__ret := C.AAsset_isAllocated(casset)
	__v := (int32)(__ret)
	return __v
}

// AssetManagerFromJava function as declared in android/asset_manager_jni.h:34
func AssetManagerFromJava(env *JNIEnv, assetManager Jobject) *AssetManager {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cassetManager, _ := (C.jobject)(assetManager), cgoAllocsUnknown
	__ret := C.AAssetManager_fromJava(cenv, cassetManager)
	__v := *(**AssetManager)(unsafe.Pointer(&__ret))
	return __v
}

// ConfigurationNew function as declared in android/configuration.h:125
func ConfigurationNew() *Configuration {
	__ret := C.AConfiguration_new()
	__v := *(**Configuration)(unsafe.Pointer(&__ret))
	return __v
}

// ConfigurationDelete function as declared in android/configuration.h:131
func ConfigurationDelete(config *Configuration) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	C.AConfiguration_delete(cconfig)
}

// ConfigurationFromAssetManager function as declared in android/configuration.h:137
func ConfigurationFromAssetManager(out *Configuration, am *AssetManager) {
	cout, _ := (*C.AConfiguration)(unsafe.Pointer(out)), cgoAllocsUnknown
	cam, _ := (*C.AAssetManager)(unsafe.Pointer(am)), cgoAllocsUnknown
	C.AConfiguration_fromAssetManager(cout, cam)
}

// ConfigurationCopy function as declared in android/configuration.h:142
func ConfigurationCopy(dest *Configuration, src *Configuration) {
	cdest, _ := (*C.AConfiguration)(unsafe.Pointer(dest)), cgoAllocsUnknown
	csrc, _ := (*C.AConfiguration)(unsafe.Pointer(src)), cgoAllocsUnknown
	C.AConfiguration_copy(cdest, csrc)
}

// ConfigurationGetMcc function as declared in android/configuration.h:147
func ConfigurationGetMcc(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getMcc(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetMcc function as declared in android/configuration.h:152
func ConfigurationSetMcc(config *Configuration, mcc int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	cmcc, _ := (C.int32_t)(mcc), cgoAllocsUnknown
	C.AConfiguration_setMcc(cconfig, cmcc)
}

// ConfigurationGetMnc function as declared in android/configuration.h:157
func ConfigurationGetMnc(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getMnc(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetMnc function as declared in android/configuration.h:162
func ConfigurationSetMnc(config *Configuration, mnc int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	cmnc, _ := (C.int32_t)(mnc), cgoAllocsUnknown
	C.AConfiguration_setMnc(cconfig, cmnc)
}

// ConfigurationGetLanguage function as declared in android/configuration.h:169
func ConfigurationGetLanguage(config *Configuration, outLanguage *byte) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	coutLanguage, _ := (*C.char)(unsafe.Pointer(outLanguage)), cgoAllocsUnknown
	C.AConfiguration_getLanguage(cconfig, coutLanguage)
}

// ConfigurationSetLanguage function as declared in android/configuration.h:175
func ConfigurationSetLanguage(config *Configuration, language string) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	clanguage, _ := unpackPCharString(language)
	C.AConfiguration_setLanguage(cconfig, clanguage)
}

// ConfigurationGetCountry function as declared in android/configuration.h:182
func ConfigurationGetCountry(config *Configuration, outCountry *byte) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	coutCountry, _ := (*C.char)(unsafe.Pointer(outCountry)), cgoAllocsUnknown
	C.AConfiguration_getCountry(cconfig, coutCountry)
}

// ConfigurationSetCountry function as declared in android/configuration.h:188
func ConfigurationSetCountry(config *Configuration, country string) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	ccountry, _ := unpackPCharString(country)
	C.AConfiguration_setCountry(cconfig, ccountry)
}

// ConfigurationGetOrientation function as declared in android/configuration.h:193
func ConfigurationGetOrientation(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getOrientation(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetOrientation function as declared in android/configuration.h:198
func ConfigurationSetOrientation(config *Configuration, orientation int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	corientation, _ := (C.int32_t)(orientation), cgoAllocsUnknown
	C.AConfiguration_setOrientation(cconfig, corientation)
}

// ConfigurationGetTouchscreen function as declared in android/configuration.h:203
func ConfigurationGetTouchscreen(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getTouchscreen(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetTouchscreen function as declared in android/configuration.h:208
func ConfigurationSetTouchscreen(config *Configuration, touchscreen int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	ctouchscreen, _ := (C.int32_t)(touchscreen), cgoAllocsUnknown
	C.AConfiguration_setTouchscreen(cconfig, ctouchscreen)
}

// ConfigurationGetDensity function as declared in android/configuration.h:213
func ConfigurationGetDensity(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getDensity(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetDensity function as declared in android/configuration.h:218
func ConfigurationSetDensity(config *Configuration, density int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	cdensity, _ := (C.int32_t)(density), cgoAllocsUnknown
	C.AConfiguration_setDensity(cconfig, cdensity)
}

// ConfigurationGetKeyboard function as declared in android/configuration.h:223
func ConfigurationGetKeyboard(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getKeyboard(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetKeyboard function as declared in android/configuration.h:228
func ConfigurationSetKeyboard(config *Configuration, keyboard int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	ckeyboard, _ := (C.int32_t)(keyboard), cgoAllocsUnknown
	C.AConfiguration_setKeyboard(cconfig, ckeyboard)
}

// ConfigurationGetNavigation function as declared in android/configuration.h:233
func ConfigurationGetNavigation(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getNavigation(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetNavigation function as declared in android/configuration.h:238
func ConfigurationSetNavigation(config *Configuration, navigation int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	cnavigation, _ := (C.int32_t)(navigation), cgoAllocsUnknown
	C.AConfiguration_setNavigation(cconfig, cnavigation)
}

// ConfigurationGetKeysHidden function as declared in android/configuration.h:243
func ConfigurationGetKeysHidden(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getKeysHidden(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetKeysHidden function as declared in android/configuration.h:248
func ConfigurationSetKeysHidden(config *Configuration, keysHidden int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	ckeysHidden, _ := (C.int32_t)(keysHidden), cgoAllocsUnknown
	C.AConfiguration_setKeysHidden(cconfig, ckeysHidden)
}

// ConfigurationGetNavHidden function as declared in android/configuration.h:253
func ConfigurationGetNavHidden(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getNavHidden(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetNavHidden function as declared in android/configuration.h:258
func ConfigurationSetNavHidden(config *Configuration, navHidden int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	cnavHidden, _ := (C.int32_t)(navHidden), cgoAllocsUnknown
	C.AConfiguration_setNavHidden(cconfig, cnavHidden)
}

// ConfigurationGetSdkVersion function as declared in android/configuration.h:263
func ConfigurationGetSdkVersion(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getSdkVersion(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetSdkVersion function as declared in android/configuration.h:268
func ConfigurationSetSdkVersion(config *Configuration, sdkVersion int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	csdkVersion, _ := (C.int32_t)(sdkVersion), cgoAllocsUnknown
	C.AConfiguration_setSdkVersion(cconfig, csdkVersion)
}

// ConfigurationGetScreenSize function as declared in android/configuration.h:273
func ConfigurationGetScreenSize(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getScreenSize(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetScreenSize function as declared in android/configuration.h:278
func ConfigurationSetScreenSize(config *Configuration, screenSize int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	cscreenSize, _ := (C.int32_t)(screenSize), cgoAllocsUnknown
	C.AConfiguration_setScreenSize(cconfig, cscreenSize)
}

// ConfigurationGetScreenLong function as declared in android/configuration.h:283
func ConfigurationGetScreenLong(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getScreenLong(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetScreenLong function as declared in android/configuration.h:288
func ConfigurationSetScreenLong(config *Configuration, screenLong int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	cscreenLong, _ := (C.int32_t)(screenLong), cgoAllocsUnknown
	C.AConfiguration_setScreenLong(cconfig, cscreenLong)
}

// ConfigurationGetUiModeType function as declared in android/configuration.h:293
func ConfigurationGetUiModeType(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getUiModeType(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetUiModeType function as declared in android/configuration.h:298
func ConfigurationSetUiModeType(config *Configuration, uiModeType int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	cuiModeType, _ := (C.int32_t)(uiModeType), cgoAllocsUnknown
	C.AConfiguration_setUiModeType(cconfig, cuiModeType)
}

// ConfigurationGetUiModeNight function as declared in android/configuration.h:303
func ConfigurationGetUiModeNight(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getUiModeNight(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetUiModeNight function as declared in android/configuration.h:308
func ConfigurationSetUiModeNight(config *Configuration, uiModeNight int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	cuiModeNight, _ := (C.int32_t)(uiModeNight), cgoAllocsUnknown
	C.AConfiguration_setUiModeNight(cconfig, cuiModeNight)
}

// ConfigurationGetScreenWidthDp function as declared in android/configuration.h:314
func ConfigurationGetScreenWidthDp(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getScreenWidthDp(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetScreenWidthDp function as declared in android/configuration.h:319
func ConfigurationSetScreenWidthDp(config *Configuration, value int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	cvalue, _ := (C.int32_t)(value), cgoAllocsUnknown
	C.AConfiguration_setScreenWidthDp(cconfig, cvalue)
}

// ConfigurationGetScreenHeightDp function as declared in android/configuration.h:325
func ConfigurationGetScreenHeightDp(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getScreenHeightDp(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetScreenHeightDp function as declared in android/configuration.h:330
func ConfigurationSetScreenHeightDp(config *Configuration, value int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	cvalue, _ := (C.int32_t)(value), cgoAllocsUnknown
	C.AConfiguration_setScreenHeightDp(cconfig, cvalue)
}

// ConfigurationGetSmallestScreenWidthDp function as declared in android/configuration.h:336
func ConfigurationGetSmallestScreenWidthDp(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getSmallestScreenWidthDp(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetSmallestScreenWidthDp function as declared in android/configuration.h:341
func ConfigurationSetSmallestScreenWidthDp(config *Configuration, value int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	cvalue, _ := (C.int32_t)(value), cgoAllocsUnknown
	C.AConfiguration_setSmallestScreenWidthDp(cconfig, cvalue)
}

// ConfigurationGetLayoutDirection function as declared in android/configuration.h:347
func ConfigurationGetLayoutDirection(config *Configuration) int32 {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.AConfiguration_getLayoutDirection(cconfig)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationSetLayoutDirection function as declared in android/configuration.h:352
func ConfigurationSetLayoutDirection(config *Configuration, value int32) {
	cconfig, _ := (*C.AConfiguration)(unsafe.Pointer(config)), cgoAllocsUnknown
	cvalue, _ := (C.int32_t)(value), cgoAllocsUnknown
	C.AConfiguration_setLayoutDirection(cconfig, cvalue)
}

// ConfigurationDiff function as declared in android/configuration.h:359
func ConfigurationDiff(config1 *Configuration, config2 *Configuration) int32 {
	cconfig1, _ := (*C.AConfiguration)(unsafe.Pointer(config1)), cgoAllocsUnknown
	cconfig2, _ := (*C.AConfiguration)(unsafe.Pointer(config2)), cgoAllocsUnknown
	__ret := C.AConfiguration_diff(cconfig1, cconfig2)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationMatch function as declared in android/configuration.h:366
func ConfigurationMatch(base *Configuration, requested *Configuration) int32 {
	cbase, _ := (*C.AConfiguration)(unsafe.Pointer(base)), cgoAllocsUnknown
	crequested, _ := (*C.AConfiguration)(unsafe.Pointer(requested)), cgoAllocsUnknown
	__ret := C.AConfiguration_match(cbase, crequested)
	__v := (int32)(__ret)
	return __v
}

// ConfigurationIsBetterThan function as declared in android/configuration.h:378
func ConfigurationIsBetterThan(base *Configuration, test *Configuration, requested *Configuration) int32 {
	cbase, _ := (*C.AConfiguration)(unsafe.Pointer(base)), cgoAllocsUnknown
	ctest, _ := (*C.AConfiguration)(unsafe.Pointer(test)), cgoAllocsUnknown
	crequested, _ := (*C.AConfiguration)(unsafe.Pointer(requested)), cgoAllocsUnknown
	__ret := C.AConfiguration_isBetterThan(cbase, ctest, crequested)
	__v := (int32)(__ret)
	return __v
}

// LogWrite function as declared in android/log.h:94
func LogWrite(prio int32, tag string, text string) int32 {
	cprio, _ := (C.int)(prio), cgoAllocsUnknown
	ctag, _ := unpackPCharString(tag)
	ctext, _ := unpackPCharString(text)
	__ret := C.__android_log_write(cprio, ctag, ctext)
	__v := (int32)(__ret)
	return __v
}

// LooperForThread function as declared in android/looper.h:45
func LooperForThread() *Looper {
	__ret := C.ALooper_forThread()
	__v := *(**Looper)(unsafe.Pointer(&__ret))
	return __v
}

// LooperPrepare function as declared in android/looper.h:65
func LooperPrepare(opts int32) *Looper {
	copts, _ := (C.int)(opts), cgoAllocsUnknown
	__ret := C.ALooper_prepare(copts)
	__v := *(**Looper)(unsafe.Pointer(&__ret))
	return __v
}

// LooperAcquire function as declared in android/looper.h:99
func LooperAcquire(looper *Looper) {
	clooper, _ := (*C.ALooper)(unsafe.Pointer(looper)), cgoAllocsUnknown
	C.ALooper_acquire(clooper)
}

// LooperRelease function as declared in android/looper.h:104
func LooperRelease(looper *Looper) {
	clooper, _ := (*C.ALooper)(unsafe.Pointer(looper)), cgoAllocsUnknown
	C.ALooper_release(clooper)
}

// LooperPollOnce function as declared in android/looper.h:187
func LooperPollOnce(timeoutMillis int32, outFd *int32, outEvents *int32, outData []unsafe.Pointer) int32 {
	ctimeoutMillis, _ := (C.int)(timeoutMillis), cgoAllocsUnknown
	coutFd, _ := (*C.int)(unsafe.Pointer(outFd)), cgoAllocsUnknown
	coutEvents, _ := (*C.int)(unsafe.Pointer(outEvents)), cgoAllocsUnknown
	coutData, _ := (*unsafe.Pointer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outData)).Data)), cgoAllocsUnknown
	__ret := C.ALooper_pollOnce(ctimeoutMillis, coutFd, coutEvents, coutData)
	__v := (int32)(__ret)
	return __v
}

// LooperPollAll function as declared in android/looper.h:194
func LooperPollAll(timeoutMillis int32, outFd *int32, outEvents *int32, outData []unsafe.Pointer) int32 {
	ctimeoutMillis, _ := (C.int)(timeoutMillis), cgoAllocsUnknown
	coutFd, _ := (*C.int)(unsafe.Pointer(outFd)), cgoAllocsUnknown
	coutEvents, _ := (*C.int)(unsafe.Pointer(outEvents)), cgoAllocsUnknown
	coutData, _ := (*unsafe.Pointer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outData)).Data)), cgoAllocsUnknown
	__ret := C.ALooper_pollAll(ctimeoutMillis, coutFd, coutEvents, coutData)
	__v := (int32)(__ret)
	return __v
}

// LooperWake function as declared in android/looper.h:202
func LooperWake(looper *Looper) {
	clooper, _ := (*C.ALooper)(unsafe.Pointer(looper)), cgoAllocsUnknown
	C.ALooper_wake(clooper)
}

// LooperAddFd function as declared in android/looper.h:230
func LooperAddFd(looper *Looper, fd int32, ident int32, events int32, callback LooperCallbackFunc, data unsafe.Pointer) int32 {
	clooper, _ := (*C.ALooper)(unsafe.Pointer(looper)), cgoAllocsUnknown
	cfd, _ := (C.int)(fd), cgoAllocsUnknown
	cident, _ := (C.int)(ident), cgoAllocsUnknown
	cevents, _ := (C.int)(events), cgoAllocsUnknown
	ccallback, _ := callback.PassValue()
	cdata, _ := data, cgoAllocsUnknown
	__ret := C.ALooper_addFd(clooper, cfd, cident, cevents, ccallback, cdata)
	__v := (int32)(__ret)
	return __v
}

// LooperRemoveFd function as declared in android/looper.h:250
func LooperRemoveFd(looper *Looper, fd int32) int32 {
	clooper, _ := (*C.ALooper)(unsafe.Pointer(looper)), cgoAllocsUnknown
	cfd, _ := (C.int)(fd), cgoAllocsUnknown
	__ret := C.ALooper_removeFd(clooper, cfd)
	__v := (int32)(__ret)
	return __v
}

// NativeActivityFinish function as declared in android/native_activity.h:252
func NativeActivityFinish(activity *NativeActivity) {
	cactivity, _ := activity.PassRef()
	C.ANativeActivity_finish(cactivity)
}

// NativeActivitySetWindowFormat function as declared in android/native_activity.h:260
func NativeActivitySetWindowFormat(activity *NativeActivity, format int32) {
	cactivity, _ := activity.PassRef()
	cformat, _ := (C.int32_t)(format), cgoAllocsUnknown
	C.ANativeActivity_setWindowFormat(cactivity, cformat)
}

// NativeActivitySetWindowFlags function as declared in android/native_activity.h:268
func NativeActivitySetWindowFlags(activity *NativeActivity, addFlags uint32, removeFlags uint32) {
	cactivity, _ := activity.PassRef()
	caddFlags, _ := (C.uint32_t)(addFlags), cgoAllocsUnknown
	cremoveFlags, _ := (C.uint32_t)(removeFlags), cgoAllocsUnknown
	C.ANativeActivity_setWindowFlags(cactivity, caddFlags, cremoveFlags)
}

// NativeActivityShowSoftInput function as declared in android/native_activity.h:286
func NativeActivityShowSoftInput(activity *NativeActivity, flags uint32) {
	cactivity, _ := activity.PassRef()
	cflags, _ := (C.uint32_t)(flags), cgoAllocsUnknown
	C.ANativeActivity_showSoftInput(cactivity, cflags)
}

// NativeActivityHideSoftInput function as declared in android/native_activity.h:303
func NativeActivityHideSoftInput(activity *NativeActivity, flags uint32) {
	cactivity, _ := activity.PassRef()
	cflags, _ := (C.uint32_t)(flags), cgoAllocsUnknown
	C.ANativeActivity_hideSoftInput(cactivity, cflags)
}

// InputEventGetType function as declared in android/input.h:499
func InputEventGetType(event *InputEvent) int32 {
	cevent, _ := (*C.AInputEvent)(unsafe.Pointer(event)), cgoAllocsUnknown
	__ret := C.AInputEvent_getType(cevent)
	__v := (int32)(__ret)
	return __v
}

// InputEventGetDeviceId function as declared in android/input.h:511
func InputEventGetDeviceId(event *InputEvent) int32 {
	cevent, _ := (*C.AInputEvent)(unsafe.Pointer(event)), cgoAllocsUnknown
	__ret := C.AInputEvent_getDeviceId(cevent)
	__v := (int32)(__ret)
	return __v
}

// InputEventGetSource function as declared in android/input.h:514
func InputEventGetSource(event *InputEvent) int32 {
	cevent, _ := (*C.AInputEvent)(unsafe.Pointer(event)), cgoAllocsUnknown
	__ret := C.AInputEvent_getSource(cevent)
	__v := (int32)(__ret)
	return __v
}

// KeyEventGetAction function as declared in android/input.h:519
func KeyEventGetAction(keyEvent *InputEvent) int32 {
	ckeyEvent, _ := (*C.AInputEvent)(unsafe.Pointer(keyEvent)), cgoAllocsUnknown
	__ret := C.AKeyEvent_getAction(ckeyEvent)
	__v := (int32)(__ret)
	return __v
}

// KeyEventGetFlags function as declared in android/input.h:522
func KeyEventGetFlags(keyEvent *InputEvent) int32 {
	ckeyEvent, _ := (*C.AInputEvent)(unsafe.Pointer(keyEvent)), cgoAllocsUnknown
	__ret := C.AKeyEvent_getFlags(ckeyEvent)
	__v := (int32)(__ret)
	return __v
}

// KeyEventGetKeyCode function as declared in android/input.h:526
func KeyEventGetKeyCode(keyEvent *InputEvent) int32 {
	ckeyEvent, _ := (*C.AInputEvent)(unsafe.Pointer(keyEvent)), cgoAllocsUnknown
	__ret := C.AKeyEvent_getKeyCode(ckeyEvent)
	__v := (int32)(__ret)
	return __v
}

// KeyEventGetScanCode function as declared in android/input.h:530
func KeyEventGetScanCode(keyEvent *InputEvent) int32 {
	ckeyEvent, _ := (*C.AInputEvent)(unsafe.Pointer(keyEvent)), cgoAllocsUnknown
	__ret := C.AKeyEvent_getScanCode(ckeyEvent)
	__v := (int32)(__ret)
	return __v
}

// KeyEventGetMetaState function as declared in android/input.h:533
func KeyEventGetMetaState(keyEvent *InputEvent) int32 {
	ckeyEvent, _ := (*C.AInputEvent)(unsafe.Pointer(keyEvent)), cgoAllocsUnknown
	__ret := C.AKeyEvent_getMetaState(ckeyEvent)
	__v := (int32)(__ret)
	return __v
}

// KeyEventGetRepeatCount function as declared in android/input.h:539
func KeyEventGetRepeatCount(keyEvent *InputEvent) int32 {
	ckeyEvent, _ := (*C.AInputEvent)(unsafe.Pointer(keyEvent)), cgoAllocsUnknown
	__ret := C.AKeyEvent_getRepeatCount(ckeyEvent)
	__v := (int32)(__ret)
	return __v
}

// KeyEventGetDownTime function as declared in android/input.h:546
func KeyEventGetDownTime(keyEvent *InputEvent) int64 {
	ckeyEvent, _ := (*C.AInputEvent)(unsafe.Pointer(keyEvent)), cgoAllocsUnknown
	__ret := C.AKeyEvent_getDownTime(ckeyEvent)
	__v := (int64)(__ret)
	return __v
}

// KeyEventGetEventTime function as declared in android/input.h:550
func KeyEventGetEventTime(keyEvent *InputEvent) int64 {
	ckeyEvent, _ := (*C.AInputEvent)(unsafe.Pointer(keyEvent)), cgoAllocsUnknown
	__ret := C.AKeyEvent_getEventTime(ckeyEvent)
	__v := (int64)(__ret)
	return __v
}

// MotionEventGetAction function as declared in android/input.h:555
func MotionEventGetAction(motionEvent *InputEvent) int32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	__ret := C.AMotionEvent_getAction(cmotionEvent)
	__v := (int32)(__ret)
	return __v
}

// MotionEventGetFlags function as declared in android/input.h:558
func MotionEventGetFlags(motionEvent *InputEvent) int32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	__ret := C.AMotionEvent_getFlags(cmotionEvent)
	__v := (int32)(__ret)
	return __v
}

// MotionEventGetMetaState function as declared in android/input.h:562
func MotionEventGetMetaState(motionEvent *InputEvent) int32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	__ret := C.AMotionEvent_getMetaState(cmotionEvent)
	__v := (int32)(__ret)
	return __v
}

// MotionEventGetButtonState function as declared in android/input.h:565
func MotionEventGetButtonState(motionEvent *InputEvent) int32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	__ret := C.AMotionEvent_getButtonState(cmotionEvent)
	__v := (int32)(__ret)
	return __v
}

// MotionEventGetEdgeFlags function as declared in android/input.h:570
func MotionEventGetEdgeFlags(motionEvent *InputEvent) int32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	__ret := C.AMotionEvent_getEdgeFlags(cmotionEvent)
	__v := (int32)(__ret)
	return __v
}

// MotionEventGetDownTime function as declared in android/input.h:574
func MotionEventGetDownTime(motionEvent *InputEvent) int64 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	__ret := C.AMotionEvent_getDownTime(cmotionEvent)
	__v := (int64)(__ret)
	return __v
}

// MotionEventGetEventTime function as declared in android/input.h:578
func MotionEventGetEventTime(motionEvent *InputEvent) int64 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	__ret := C.AMotionEvent_getEventTime(cmotionEvent)
	__v := (int64)(__ret)
	return __v
}

// MotionEventGetXOffset function as declared in android/input.h:584
func MotionEventGetXOffset(motionEvent *InputEvent) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	__ret := C.AMotionEvent_getXOffset(cmotionEvent)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetYOffset function as declared in android/input.h:590
func MotionEventGetYOffset(motionEvent *InputEvent) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	__ret := C.AMotionEvent_getYOffset(cmotionEvent)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetXPrecision function as declared in android/input.h:595
func MotionEventGetXPrecision(motionEvent *InputEvent) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	__ret := C.AMotionEvent_getXPrecision(cmotionEvent)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetYPrecision function as declared in android/input.h:600
func MotionEventGetYPrecision(motionEvent *InputEvent) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	__ret := C.AMotionEvent_getYPrecision(cmotionEvent)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetPointerCount function as declared in android/input.h:604
func MotionEventGetPointerCount(motionEvent *InputEvent) uint32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	__ret := C.AMotionEvent_getPointerCount(cmotionEvent)
	__v := (uint32)(__ret)
	return __v
}

// MotionEventGetPointerId function as declared in android/input.h:610
func MotionEventGetPointerId(motionEvent *InputEvent, pointerIndex uint32) int32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getPointerId(cmotionEvent, cpointerIndex)
	__v := (int32)(__ret)
	return __v
}

// MotionEventGetToolType function as declared in android/input.h:615
func MotionEventGetToolType(motionEvent *InputEvent, pointerIndex uint32) int32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getToolType(cmotionEvent, cpointerIndex)
	__v := (int32)(__ret)
	return __v
}

// MotionEventGetRawX function as declared in android/input.h:621
func MotionEventGetRawX(motionEvent *InputEvent, pointerIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getRawX(cmotionEvent, cpointerIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetRawY function as declared in android/input.h:627
func MotionEventGetRawY(motionEvent *InputEvent, pointerIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getRawY(cmotionEvent, cpointerIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetX function as declared in android/input.h:632
func MotionEventGetX(motionEvent *InputEvent, pointerIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getX(cmotionEvent, cpointerIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetY function as declared in android/input.h:637
func MotionEventGetY(motionEvent *InputEvent, pointerIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getY(cmotionEvent, cpointerIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetPressure function as declared in android/input.h:643
func MotionEventGetPressure(motionEvent *InputEvent, pointerIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getPressure(cmotionEvent, cpointerIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetSize function as declared in android/input.h:651
func MotionEventGetSize(motionEvent *InputEvent, pointerIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getSize(cmotionEvent, cpointerIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetTouchMajor function as declared in android/input.h:655
func MotionEventGetTouchMajor(motionEvent *InputEvent, pointerIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getTouchMajor(cmotionEvent, cpointerIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetTouchMinor function as declared in android/input.h:659
func MotionEventGetTouchMinor(motionEvent *InputEvent, pointerIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getTouchMinor(cmotionEvent, cpointerIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetToolMajor function as declared in android/input.h:665
func MotionEventGetToolMajor(motionEvent *InputEvent, pointerIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getToolMajor(cmotionEvent, cpointerIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetToolMinor function as declared in android/input.h:671
func MotionEventGetToolMinor(motionEvent *InputEvent, pointerIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getToolMinor(cmotionEvent, cpointerIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetOrientation function as declared in android/input.h:681
func MotionEventGetOrientation(motionEvent *InputEvent, pointerIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getOrientation(cmotionEvent, cpointerIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetAxisValue function as declared in android/input.h:684
func MotionEventGetAxisValue(motionEvent *InputEvent, axis int32, pointerIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	caxis, _ := (C.int32_t)(axis), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getAxisValue(cmotionEvent, caxis, cpointerIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetHistorySize function as declared in android/input.h:691
func MotionEventGetHistorySize(motionEvent *InputEvent) uint32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	__ret := C.AMotionEvent_getHistorySize(cmotionEvent)
	__v := (uint32)(__ret)
	return __v
}

// MotionEventGetHistoricalEventTime function as declared in android/input.h:695
func MotionEventGetHistoricalEventTime(motionEvent *InputEvent, historyIndex uint32) int64 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	chistoryIndex, _ := (C.size_t)(historyIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getHistoricalEventTime(cmotionEvent, chistoryIndex)
	__v := (int64)(__ret)
	return __v
}

// MotionEventGetHistoricalRawX function as declared in android/input.h:705
func MotionEventGetHistoricalRawX(motionEvent *InputEvent, pointerIndex uint32, historyIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	chistoryIndex, _ := (C.size_t)(historyIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getHistoricalRawX(cmotionEvent, cpointerIndex, chistoryIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetHistoricalRawY function as declared in android/input.h:715
func MotionEventGetHistoricalRawY(motionEvent *InputEvent, pointerIndex uint32, historyIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	chistoryIndex, _ := (C.size_t)(historyIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getHistoricalRawY(cmotionEvent, cpointerIndex, chistoryIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetHistoricalX function as declared in android/input.h:722
func MotionEventGetHistoricalX(motionEvent *InputEvent, pointerIndex uint32, historyIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	chistoryIndex, _ := (C.size_t)(historyIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getHistoricalX(cmotionEvent, cpointerIndex, chistoryIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetHistoricalY function as declared in android/input.h:729
func MotionEventGetHistoricalY(motionEvent *InputEvent, pointerIndex uint32, historyIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	chistoryIndex, _ := (C.size_t)(historyIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getHistoricalY(cmotionEvent, cpointerIndex, chistoryIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetHistoricalPressure function as declared in android/input.h:737
func MotionEventGetHistoricalPressure(motionEvent *InputEvent, pointerIndex uint32, historyIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	chistoryIndex, _ := (C.size_t)(historyIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getHistoricalPressure(cmotionEvent, cpointerIndex, chistoryIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetHistoricalSize function as declared in android/input.h:747
func MotionEventGetHistoricalSize(motionEvent *InputEvent, pointerIndex uint32, historyIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	chistoryIndex, _ := (C.size_t)(historyIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getHistoricalSize(cmotionEvent, cpointerIndex, chistoryIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetHistoricalTouchMajor function as declared in android/input.h:753
func MotionEventGetHistoricalTouchMajor(motionEvent *InputEvent, pointerIndex uint32, historyIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	chistoryIndex, _ := (C.size_t)(historyIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getHistoricalTouchMajor(cmotionEvent, cpointerIndex, chistoryIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetHistoricalTouchMinor function as declared in android/input.h:759
func MotionEventGetHistoricalTouchMinor(motionEvent *InputEvent, pointerIndex uint32, historyIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	chistoryIndex, _ := (C.size_t)(historyIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getHistoricalTouchMinor(cmotionEvent, cpointerIndex, chistoryIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetHistoricalToolMajor function as declared in android/input.h:767
func MotionEventGetHistoricalToolMajor(motionEvent *InputEvent, pointerIndex uint32, historyIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	chistoryIndex, _ := (C.size_t)(historyIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getHistoricalToolMajor(cmotionEvent, cpointerIndex, chistoryIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetHistoricalToolMinor function as declared in android/input.h:775
func MotionEventGetHistoricalToolMinor(motionEvent *InputEvent, pointerIndex uint32, historyIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	chistoryIndex, _ := (C.size_t)(historyIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getHistoricalToolMinor(cmotionEvent, cpointerIndex, chistoryIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetHistoricalOrientation function as declared in android/input.h:787
func MotionEventGetHistoricalOrientation(motionEvent *InputEvent, pointerIndex uint32, historyIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	chistoryIndex, _ := (C.size_t)(historyIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getHistoricalOrientation(cmotionEvent, cpointerIndex, chistoryIndex)
	__v := (float32)(__ret)
	return __v
}

// MotionEventGetHistoricalAxisValue function as declared in android/input.h:792
func MotionEventGetHistoricalAxisValue(motionEvent *InputEvent, axis int32, pointerIndex uint32, historyIndex uint32) float32 {
	cmotionEvent, _ := (*C.AInputEvent)(unsafe.Pointer(motionEvent)), cgoAllocsUnknown
	caxis, _ := (C.int32_t)(axis), cgoAllocsUnknown
	cpointerIndex, _ := (C.size_t)(pointerIndex), cgoAllocsUnknown
	chistoryIndex, _ := (C.size_t)(historyIndex), cgoAllocsUnknown
	__ret := C.AMotionEvent_getHistoricalAxisValue(cmotionEvent, caxis, cpointerIndex, chistoryIndex)
	__v := (float32)(__ret)
	return __v
}

// InputQueueAttachLooper function as declared in android/input.h:809
func InputQueueAttachLooper(queue *InputQueue, looper *Looper, ident int32, callback LooperCallbackFunc, data unsafe.Pointer) {
	cqueue, _ := (*C.AInputQueue)(unsafe.Pointer(queue)), cgoAllocsUnknown
	clooper, _ := (*C.ALooper)(unsafe.Pointer(looper)), cgoAllocsUnknown
	cident, _ := (C.int)(ident), cgoAllocsUnknown
	ccallback, _ := callback.PassValue()
	cdata, _ := data, cgoAllocsUnknown
	C.AInputQueue_attachLooper(cqueue, clooper, cident, ccallback, cdata)
}

// InputQueueDetachLooper function as declared in android/input.h:815
func InputQueueDetachLooper(queue *InputQueue) {
	cqueue, _ := (*C.AInputQueue)(unsafe.Pointer(queue)), cgoAllocsUnknown
	C.AInputQueue_detachLooper(cqueue)
}

// InputQueueHasEvents function as declared in android/input.h:822
func InputQueueHasEvents(queue *InputQueue) int32 {
	cqueue, _ := (*C.AInputQueue)(unsafe.Pointer(queue)), cgoAllocsUnknown
	__ret := C.AInputQueue_hasEvents(cqueue)
	__v := (int32)(__ret)
	return __v
}

// InputQueueGetEvent function as declared in android/input.h:828
func InputQueueGetEvent(queue *InputQueue, outEvent **InputEvent) int32 {
	cqueue, _ := (*C.AInputQueue)(unsafe.Pointer(queue)), cgoAllocsUnknown
	coutEvent, _ := (**C.AInputEvent)(unsafe.Pointer(outEvent)), cgoAllocsUnknown
	__ret := C.AInputQueue_getEvent(cqueue, coutEvent)
	__v := (int32)(__ret)
	return __v
}

// InputQueuePreDispatchEvent function as declared in android/input.h:838
func InputQueuePreDispatchEvent(queue *InputQueue, event *InputEvent) int32 {
	cqueue, _ := (*C.AInputQueue)(unsafe.Pointer(queue)), cgoAllocsUnknown
	cevent, _ := (*C.AInputEvent)(unsafe.Pointer(event)), cgoAllocsUnknown
	__ret := C.AInputQueue_preDispatchEvent(cqueue, cevent)
	__v := (int32)(__ret)
	return __v
}

// InputQueueFinishEvent function as declared in android/input.h:844
func InputQueueFinishEvent(queue *InputQueue, event *InputEvent, handled int32) {
	cqueue, _ := (*C.AInputQueue)(unsafe.Pointer(queue)), cgoAllocsUnknown
	cevent, _ := (*C.AInputEvent)(unsafe.Pointer(event)), cgoAllocsUnknown
	chandled, _ := (C.int)(handled), cgoAllocsUnknown
	C.AInputQueue_finishEvent(cqueue, cevent, chandled)
}

// NativeWindowAcquire function as declared in android/native_window.h:63
func NativeWindowAcquire(window *NativeWindow) {
	cwindow, _ := (*C.ANativeWindow)(unsafe.Pointer(window)), cgoAllocsUnknown
	C.ANativeWindow_acquire(cwindow)
}

// NativeWindowRelease function as declared in android/native_window.h:68
func NativeWindowRelease(window *NativeWindow) {
	cwindow, _ := (*C.ANativeWindow)(unsafe.Pointer(window)), cgoAllocsUnknown
	C.ANativeWindow_release(cwindow)
}

// NativeWindowGetWidth function as declared in android/native_window.h:74
func NativeWindowGetWidth(window *NativeWindow) int32 {
	cwindow, _ := (*C.ANativeWindow)(unsafe.Pointer(window)), cgoAllocsUnknown
	__ret := C.ANativeWindow_getWidth(cwindow)
	__v := (int32)(__ret)
	return __v
}

// NativeWindowGetHeight function as declared in android/native_window.h:80
func NativeWindowGetHeight(window *NativeWindow) int32 {
	cwindow, _ := (*C.ANativeWindow)(unsafe.Pointer(window)), cgoAllocsUnknown
	__ret := C.ANativeWindow_getHeight(cwindow)
	__v := (int32)(__ret)
	return __v
}

// NativeWindowGetFormat function as declared in android/native_window.h:86
func NativeWindowGetFormat(window *NativeWindow) int32 {
	cwindow, _ := (*C.ANativeWindow)(unsafe.Pointer(window)), cgoAllocsUnknown
	__ret := C.ANativeWindow_getFormat(cwindow)
	__v := (int32)(__ret)
	return __v
}

// NativeWindowSetBuffersGeometry function as declared in android/native_window.h:102
func NativeWindowSetBuffersGeometry(window *NativeWindow, width int32, height int32, format int32) int32 {
	cwindow, _ := (*C.ANativeWindow)(unsafe.Pointer(window)), cgoAllocsUnknown
	cwidth, _ := (C.int32_t)(width), cgoAllocsUnknown
	cheight, _ := (C.int32_t)(height), cgoAllocsUnknown
	cformat, _ := (C.int32_t)(format), cgoAllocsUnknown
	__ret := C.ANativeWindow_setBuffersGeometry(cwindow, cwidth, cheight, cformat)
	__v := (int32)(__ret)
	return __v
}

// NativeWindowLock function as declared in android/native_window.h:113
func NativeWindowLock(window *NativeWindow, outBuffer *NativeWindowBuffer, inOutDirtyBounds *Rect) int32 {
	cwindow, _ := (*C.ANativeWindow)(unsafe.Pointer(window)), cgoAllocsUnknown
	coutBuffer, _ := outBuffer.PassRef()
	cinOutDirtyBounds, _ := inOutDirtyBounds.PassRef()
	__ret := C.ANativeWindow_lock(cwindow, coutBuffer, cinOutDirtyBounds)
	__v := (int32)(__ret)
	return __v
}

// NativeWindowUnlockAndPost function as declared in android/native_window.h:120
func NativeWindowUnlockAndPost(window *NativeWindow) int32 {
	cwindow, _ := (*C.ANativeWindow)(unsafe.Pointer(window)), cgoAllocsUnknown
	__ret := C.ANativeWindow_unlockAndPost(cwindow)
	__v := (int32)(__ret)
	return __v
}

// NativeWindowFromSurface function as declared in android/native_window_jni.h:34
func NativeWindowFromSurface(env *JNIEnv, surface Jobject) *NativeWindow {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	csurface, _ := (C.jobject)(surface), cgoAllocsUnknown
	__ret := C.ANativeWindow_fromSurface(cenv, csurface)
	__v := *(**NativeWindow)(unsafe.Pointer(&__ret))
	return __v
}

// ObbScannerGetObbInfo function as declared in android/obb.h:37
func ObbScannerGetObbInfo(filename string) *ObbInfo {
	cfilename, _ := unpackPCharString(filename)
	__ret := C.AObbScanner_getObbInfo(cfilename)
	__v := *(**ObbInfo)(unsafe.Pointer(&__ret))
	return __v
}

// ObbInfoDelete function as declared in android/obb.h:42
func ObbInfoDelete(obbInfo *ObbInfo) {
	cobbInfo, _ := (*C.AObbInfo)(unsafe.Pointer(obbInfo)), cgoAllocsUnknown
	C.AObbInfo_delete(cobbInfo)
}

// ObbInfoGetPackageName function as declared in android/obb.h:47
func ObbInfoGetPackageName(obbInfo *ObbInfo) string {
	cobbInfo, _ := (*C.AObbInfo)(unsafe.Pointer(obbInfo)), cgoAllocsUnknown
	__ret := C.AObbInfo_getPackageName(cobbInfo)
	__v := packPCharString(__ret)
	return __v
}

// ObbInfoGetVersion function as declared in android/obb.h:52
func ObbInfoGetVersion(obbInfo *ObbInfo) int32 {
	cobbInfo, _ := (*C.AObbInfo)(unsafe.Pointer(obbInfo)), cgoAllocsUnknown
	__ret := C.AObbInfo_getVersion(cobbInfo)
	__v := (int32)(__ret)
	return __v
}

// ObbInfoGetFlags function as declared in android/obb.h:57
func ObbInfoGetFlags(obbInfo *ObbInfo) int32 {
	cobbInfo, _ := (*C.AObbInfo)(unsafe.Pointer(obbInfo)), cgoAllocsUnknown
	__ret := C.AObbInfo_getFlags(cobbInfo)
	__v := (int32)(__ret)
	return __v
}

// SensorManagerGetInstance function as declared in android/sensor.h:204
func SensorManagerGetInstance() *SensorManager {
	__ret := C.ASensorManager_getInstance()
	__v := *(**SensorManager)(unsafe.Pointer(&__ret))
	return __v
}

// SensorManagerGetDefaultSensor function as declared in android/sensor.h:216
func SensorManagerGetDefaultSensor(manager *SensorManager, kind int32) *Sensor {
	cmanager, _ := (*C.ASensorManager)(unsafe.Pointer(manager)), cgoAllocsUnknown
	ckind, _ := (C.int)(kind), cgoAllocsUnknown
	__ret := C.ASensorManager_getDefaultSensor(cmanager, ckind)
	__v := *(**Sensor)(unsafe.Pointer(&__ret))
	return __v
}

// SensorManagerGetDefaultSensorEx function as declared in android/sensor.h:222
func SensorManagerGetDefaultSensorEx(manager *SensorManager, kind int32, wakeUp bool) *Sensor {
	cmanager, _ := (*C.ASensorManager)(unsafe.Pointer(manager)), cgoAllocsUnknown
	ckind, _ := (C.int)(kind), cgoAllocsUnknown
	cwakeUp, _ := (C._Bool)(wakeUp), cgoAllocsUnknown
	__ret := C.ASensorManager_getDefaultSensorEx(cmanager, ckind, cwakeUp)
	__v := *(**Sensor)(unsafe.Pointer(&__ret))
	return __v
}

// SensorManagerCreateEventQueue function as declared in android/sensor.h:228
func SensorManagerCreateEventQueue(manager *SensorManager, looper *Looper, ident int32, callback LooperCallbackFunc, data unsafe.Pointer) *SensorEventQueue {
	cmanager, _ := (*C.ASensorManager)(unsafe.Pointer(manager)), cgoAllocsUnknown
	clooper, _ := (*C.ALooper)(unsafe.Pointer(looper)), cgoAllocsUnknown
	cident, _ := (C.int)(ident), cgoAllocsUnknown
	ccallback, _ := callback.PassValue()
	cdata, _ := data, cgoAllocsUnknown
	__ret := C.ASensorManager_createEventQueue(cmanager, clooper, cident, ccallback, cdata)
	__v := *(**SensorEventQueue)(unsafe.Pointer(&__ret))
	return __v
}

// SensorManagerDestroyEventQueue function as declared in android/sensor.h:234
func SensorManagerDestroyEventQueue(manager *SensorManager, queue *SensorEventQueue) int32 {
	cmanager, _ := (*C.ASensorManager)(unsafe.Pointer(manager)), cgoAllocsUnknown
	cqueue, _ := (*C.ASensorEventQueue)(unsafe.Pointer(queue)), cgoAllocsUnknown
	__ret := C.ASensorManager_destroyEventQueue(cmanager, cqueue)
	__v := (int32)(__ret)
	return __v
}

// SensorEventQueueEnableSensor function as declared in android/sensor.h:242
func SensorEventQueueEnableSensor(queue *SensorEventQueue, sensor *Sensor) int32 {
	cqueue, _ := (*C.ASensorEventQueue)(unsafe.Pointer(queue)), cgoAllocsUnknown
	csensor, _ := (*C.ASensor)(unsafe.Pointer(sensor)), cgoAllocsUnknown
	__ret := C.ASensorEventQueue_enableSensor(cqueue, csensor)
	__v := (int32)(__ret)
	return __v
}

// SensorEventQueueDisableSensor function as declared in android/sensor.h:247
func SensorEventQueueDisableSensor(queue *SensorEventQueue, sensor *Sensor) int32 {
	cqueue, _ := (*C.ASensorEventQueue)(unsafe.Pointer(queue)), cgoAllocsUnknown
	csensor, _ := (*C.ASensor)(unsafe.Pointer(sensor)), cgoAllocsUnknown
	__ret := C.ASensorEventQueue_disableSensor(cqueue, csensor)
	__v := (int32)(__ret)
	return __v
}

// SensorEventQueueSetEventRate function as declared in android/sensor.h:256
func SensorEventQueueSetEventRate(queue *SensorEventQueue, sensor *Sensor, usec int32) int32 {
	cqueue, _ := (*C.ASensorEventQueue)(unsafe.Pointer(queue)), cgoAllocsUnknown
	csensor, _ := (*C.ASensor)(unsafe.Pointer(sensor)), cgoAllocsUnknown
	cusec, _ := (C.int32_t)(usec), cgoAllocsUnknown
	__ret := C.ASensorEventQueue_setEventRate(cqueue, csensor, cusec)
	__v := (int32)(__ret)
	return __v
}

// SensorEventQueueHasEvents function as declared in android/sensor.h:263
func SensorEventQueueHasEvents(queue *SensorEventQueue) int32 {
	cqueue, _ := (*C.ASensorEventQueue)(unsafe.Pointer(queue)), cgoAllocsUnknown
	__ret := C.ASensorEventQueue_hasEvents(cqueue)
	__v := (int32)(__ret)
	return __v
}

// SensorEventQueueGetEvents function as declared in android/sensor.h:278
func SensorEventQueueGetEvents(queue *SensorEventQueue, events *SensorEvent, count uint32) int32 {
	cqueue, _ := (*C.ASensorEventQueue)(unsafe.Pointer(queue)), cgoAllocsUnknown
	cevents, _ := events.PassRef()
	ccount, _ := (C.size_t)(count), cgoAllocsUnknown
	__ret := C.ASensorEventQueue_getEvents(cqueue, cevents, ccount)
	__v := (int32)(__ret)
	return __v
}

// SensorGetName function as declared in android/sensor.h:287
func SensorGetName(sensor *Sensor) string {
	csensor, _ := (*C.ASensor)(unsafe.Pointer(sensor)), cgoAllocsUnknown
	__ret := C.ASensor_getName(csensor)
	__v := packPCharString(__ret)
	return __v
}

// SensorGetVendor function as declared in android/sensor.h:292
func SensorGetVendor(sensor *Sensor) string {
	csensor, _ := (*C.ASensor)(unsafe.Pointer(sensor)), cgoAllocsUnknown
	__ret := C.ASensor_getVendor(csensor)
	__v := packPCharString(__ret)
	return __v
}

// SensorGetType function as declared in android/sensor.h:297
func SensorGetType(sensor *Sensor) int32 {
	csensor, _ := (*C.ASensor)(unsafe.Pointer(sensor)), cgoAllocsUnknown
	__ret := C.ASensor_getType(csensor)
	__v := (int32)(__ret)
	return __v
}

// SensorGetResolution function as declared in android/sensor.h:302
func SensorGetResolution(sensor *Sensor) float32 {
	csensor, _ := (*C.ASensor)(unsafe.Pointer(sensor)), cgoAllocsUnknown
	__ret := C.ASensor_getResolution(csensor)
	__v := (float32)(__ret)
	return __v
}

// SensorGetMinDelay function as declared in android/sensor.h:309
func SensorGetMinDelay(sensor *Sensor) int32 {
	csensor, _ := (*C.ASensor)(unsafe.Pointer(sensor)), cgoAllocsUnknown
	__ret := C.ASensor_getMinDelay(csensor)
	__v := (int32)(__ret)
	return __v
}

// SensorGetFifoMaxEventCount function as declared in android/sensor.h:315
func SensorGetFifoMaxEventCount(sensor *Sensor) int32 {
	csensor, _ := (*C.ASensor)(unsafe.Pointer(sensor)), cgoAllocsUnknown
	__ret := C.ASensor_getFifoMaxEventCount(csensor)
	__v := (int32)(__ret)
	return __v
}

// SensorGetFifoReservedEventCount function as declared in android/sensor.h:320
func SensorGetFifoReservedEventCount(sensor *Sensor) int32 {
	csensor, _ := (*C.ASensor)(unsafe.Pointer(sensor)), cgoAllocsUnknown
	__ret := C.ASensor_getFifoReservedEventCount(csensor)
	__v := (int32)(__ret)
	return __v
}

// SensorGetStringType function as declared in android/sensor.h:325
func SensorGetStringType(sensor *Sensor) string {
	csensor, _ := (*C.ASensor)(unsafe.Pointer(sensor)), cgoAllocsUnknown
	__ret := C.ASensor_getStringType(csensor)
	__v := packPCharString(__ret)
	return __v
}

// SensorGetReportingMode function as declared in android/sensor.h:330
func SensorGetReportingMode(sensor *Sensor) int32 {
	csensor, _ := (*C.ASensor)(unsafe.Pointer(sensor)), cgoAllocsUnknown
	__ret := C.ASensor_getReportingMode(csensor)
	__v := (int32)(__ret)
	return __v
}

// SensorIsWakeUpSensor function as declared in android/sensor.h:335
func SensorIsWakeUpSensor(sensor *Sensor) bool {
	csensor, _ := (*C.ASensor)(unsafe.Pointer(sensor)), cgoAllocsUnknown
	__ret := C.ASensor_isWakeUpSensor(csensor)
	__v := (bool)(__ret)
	return __v
}

// StorageManagerNew function as declared in android/storage_manager.h:88
func StorageManagerNew() *StorageManager {
	__ret := C.AStorageManager_new()
	__v := *(**StorageManager)(unsafe.Pointer(&__ret))
	return __v
}

// StorageManagerDelete function as declared in android/storage_manager.h:93
func StorageManagerDelete(mgr *StorageManager) {
	cmgr, _ := (*C.AStorageManager)(unsafe.Pointer(mgr)), cgoAllocsUnknown
	C.AStorageManager_delete(cmgr)
}

// StorageManagerMountObb function as declared in android/storage_manager.h:103
func StorageManagerMountObb(mgr *StorageManager, filename string, key string, cb StorageManagerObbCallbackFunc, data unsafe.Pointer) {
	cmgr, _ := (*C.AStorageManager)(unsafe.Pointer(mgr)), cgoAllocsUnknown
	cfilename, _ := unpackPCharString(filename)
	ckey, _ := unpackPCharString(key)
	ccb, _ := cb.PassValue()
	cdata, _ := data, cgoAllocsUnknown
	C.AStorageManager_mountObb(cmgr, cfilename, ckey, ccb, cdata)
}

// StorageManagerUnmountObb function as declared in android/storage_manager.h:109
func StorageManagerUnmountObb(mgr *StorageManager, filename string, force int32, cb StorageManagerObbCallbackFunc, data unsafe.Pointer) {
	cmgr, _ := (*C.AStorageManager)(unsafe.Pointer(mgr)), cgoAllocsUnknown
	cfilename, _ := unpackPCharString(filename)
	cforce, _ := (C.int)(force), cgoAllocsUnknown
	ccb, _ := cb.PassValue()
	cdata, _ := data, cgoAllocsUnknown
	C.AStorageManager_unmountObb(cmgr, cfilename, cforce, ccb, cdata)
}

// StorageManagerIsObbMounted function as declared in android/storage_manager.h:115
func StorageManagerIsObbMounted(mgr *StorageManager, filename string) int32 {
	cmgr, _ := (*C.AStorageManager)(unsafe.Pointer(mgr)), cgoAllocsUnknown
	cfilename, _ := unpackPCharString(filename)
	__ret := C.AStorageManager_isObbMounted(cmgr, cfilename)
	__v := (int32)(__ret)
	return __v
}

// StorageManagerGetMountedObbPath function as declared in android/storage_manager.h:120
func StorageManagerGetMountedObbPath(mgr *StorageManager, filename string) string {
	cmgr, _ := (*C.AStorageManager)(unsafe.Pointer(mgr)), cgoAllocsUnknown
	cfilename, _ := unpackPCharString(filename)
	__ret := C.AStorageManager_getMountedObbPath(cmgr, cfilename)
	__v := packPCharString(__ret)
	return __v
}

// JNIDestroyJavaVM function as declared in android/jni_call.h:6
func JNIDestroyJavaVM(vm *JavaVM) Jint {
	cvm, _ := (*C.JavaVM)(unsafe.Pointer(vm)), cgoAllocsUnknown
	__ret := C.JNI_DestroyJavaVM(cvm)
	__v := (Jint)(__ret)
	return __v
}

// JNIAttachCurrentThread function as declared in android/jni_call.h:7
func JNIAttachCurrentThread(vm *JavaVM, pEnv **JNIEnv, thrArgs *JavaVMAttachArgs) Jint {
	cvm, _ := (*C.JavaVM)(unsafe.Pointer(vm)), cgoAllocsUnknown
	cpEnv, _ := (**C.JNIEnv)(unsafe.Pointer(pEnv)), cgoAllocsUnknown
	cthrArgs, _ := thrArgs.PassRef()
	__ret := C.JNI_AttachCurrentThread(cvm, cpEnv, cthrArgs)
	__v := (Jint)(__ret)
	return __v
}

// JNIDetachCurrentThread function as declared in android/jni_call.h:8
func JNIDetachCurrentThread(vm *JavaVM) Jint {
	cvm, _ := (*C.JavaVM)(unsafe.Pointer(vm)), cgoAllocsUnknown
	__ret := C.JNI_DetachCurrentThread(cvm)
	__v := (Jint)(__ret)
	return __v
}

// JNIGetEnv function as declared in android/jni_call.h:9
func JNIGetEnv(vm *JavaVM, pEnv **JNIEnv, version int32) Jint {
	cvm, _ := (*C.JavaVM)(unsafe.Pointer(vm)), cgoAllocsUnknown
	cpEnv, _ := (**C.JNIEnv)(unsafe.Pointer(pEnv)), cgoAllocsUnknown
	cversion, _ := (C.jint)(version), cgoAllocsUnknown
	__ret := C.JNI_GetEnv(cvm, cpEnv, cversion)
	__v := (Jint)(__ret)
	return __v
}

// JNIAttachCurrentThreadAsDaemon function as declared in android/jni_call.h:10
func JNIAttachCurrentThreadAsDaemon(vm *JavaVM, pEnv **JNIEnv, thrArgs unsafe.Pointer) Jint {
	cvm, _ := (*C.JavaVM)(unsafe.Pointer(vm)), cgoAllocsUnknown
	cpEnv, _ := (**C.JNIEnv)(unsafe.Pointer(pEnv)), cgoAllocsUnknown
	cthrArgs, _ := thrArgs, cgoAllocsUnknown
	__ret := C.JNI_AttachCurrentThreadAsDaemon(cvm, cpEnv, cthrArgs)
	__v := (Jint)(__ret)
	return __v
}

// JNIEnvGetVersion function as declared in android/jni_call.h:12
func JNIEnvGetVersion(env *JNIEnv) Jint {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetVersion(cenv)
	__v := (Jint)(__ret)
	return __v
}

// JNIEnvDefineClass function as declared in android/jni_call.h:14
func JNIEnvDefineClass(env *JNIEnv, name string, obj Jobject, buf []byte, bufLen int32) *Jclass {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cbuf, _ := (*C.jbyte)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buf)).Data)), cgoAllocsUnknown
	cbufLen, _ := (C.jsize)(bufLen), cgoAllocsUnknown
	__ret := C.JNIEnv_DefineClass(cenv, cname, cobj, cbuf, cbufLen)
	__v := *(**Jclass)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvFindClass function as declared in android/jni_call.h:15
func JNIEnvFindClass(env *JNIEnv, name string) *Jclass {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	__ret := C.JNIEnv_FindClass(cenv, cname)
	__v := *(**Jclass)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvFromReflectedMethod function as declared in android/jni_call.h:17
func JNIEnvFromReflectedMethod(env *JNIEnv, obj Jobject) JmethodID {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	__ret := C.JNIEnv_FromReflectedMethod(cenv, cobj)
	__v := *(*JmethodID)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvFromReflectedField function as declared in android/jni_call.h:18
func JNIEnvFromReflectedField(env *JNIEnv, obj Jobject) JfieldID {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	__ret := C.JNIEnv_FromReflectedField(cenv, cobj)
	__v := *(*JfieldID)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvToReflectedMethod function as declared in android/jni_call.h:19
func JNIEnvToReflectedMethod(env *JNIEnv, clazz *Jclass, id JmethodID, isStatic byte) Jobject {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cisStatic, _ := (C.jboolean)(isStatic), cgoAllocsUnknown
	__ret := C.JNIEnv_ToReflectedMethod(cenv, cclazz, cid, cisStatic)
	__v := (Jobject)(__ret)
	return __v
}

// JNIEnvGetSuperclass function as declared in android/jni_call.h:21
func JNIEnvGetSuperclass(env *JNIEnv, clazz *Jclass) *Jclass {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetSuperclass(cenv, cclazz)
	__v := *(**Jclass)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvIsAssignableFrom function as declared in android/jni_call.h:22
func JNIEnvIsAssignableFrom(env *JNIEnv, clazz1 *Jclass, clazz2 *Jclass) Jboolean {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz1, _ := (C.jclass)(unsafe.Pointer(clazz1)), cgoAllocsUnknown
	cclazz2, _ := (C.jclass)(unsafe.Pointer(clazz2)), cgoAllocsUnknown
	__ret := C.JNIEnv_IsAssignableFrom(cenv, cclazz1, cclazz2)
	__v := (Jboolean)(__ret)
	return __v
}

// JNIEnvToReflectedField function as declared in android/jni_call.h:24
func JNIEnvToReflectedField(env *JNIEnv, clazz *Jclass, id JfieldID, isStatic byte) Jobject {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cisStatic, _ := (C.jboolean)(isStatic), cgoAllocsUnknown
	__ret := C.JNIEnv_ToReflectedField(cenv, cclazz, cid, cisStatic)
	__v := (Jobject)(__ret)
	return __v
}

// JNIEnvThrow function as declared in android/jni_call.h:26
func JNIEnvThrow(env *JNIEnv, ex *Jthrowable) Jint {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cex, _ := (C.jthrowable)(unsafe.Pointer(ex)), cgoAllocsUnknown
	__ret := C.JNIEnv_Throw(cenv, cex)
	__v := (Jint)(__ret)
	return __v
}

// JNIEnvThrowNew function as declared in android/jni_call.h:27
func JNIEnvThrowNew(env *JNIEnv, clazz *Jclass, msg string) Jint {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cmsg, _ := unpackPCharString(msg)
	__ret := C.JNIEnv_ThrowNew(cenv, cclazz, cmsg)
	__v := (Jint)(__ret)
	return __v
}

// JNIEnvExceptionOccurred function as declared in android/jni_call.h:28
func JNIEnvExceptionOccurred(env *JNIEnv) *Jthrowable {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	__ret := C.JNIEnv_ExceptionOccurred(cenv)
	__v := *(**Jthrowable)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvExceptionDescribe function as declared in android/jni_call.h:29
func JNIEnvExceptionDescribe(env *JNIEnv) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	C.JNIEnv_ExceptionDescribe(cenv)
}

// JNIEnvExceptionClear function as declared in android/jni_call.h:30
func JNIEnvExceptionClear(env *JNIEnv) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	C.JNIEnv_ExceptionClear(cenv)
}

// JNIEnvFatalError function as declared in android/jni_call.h:31
func JNIEnvFatalError(env *JNIEnv, msg string) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cmsg, _ := unpackPCharString(msg)
	C.JNIEnv_FatalError(cenv, cmsg)
}

// JNIEnvPushLocalFrame function as declared in android/jni_call.h:33
func JNIEnvPushLocalFrame(env *JNIEnv, capacity int32) Jint {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	ccapacity, _ := (C.jint)(capacity), cgoAllocsUnknown
	__ret := C.JNIEnv_PushLocalFrame(cenv, ccapacity)
	__v := (Jint)(__ret)
	return __v
}

// JNIEnvPopLocalFrame function as declared in android/jni_call.h:34
func JNIEnvPopLocalFrame(env *JNIEnv, obj Jobject) Jobject {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	__ret := C.JNIEnv_PopLocalFrame(cenv, cobj)
	__v := (Jobject)(__ret)
	return __v
}

// JNIEnvNewGlobalRef function as declared in android/jni_call.h:36
func JNIEnvNewGlobalRef(env *JNIEnv, ref Jobject) Jobject {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cref, _ := (C.jobject)(ref), cgoAllocsUnknown
	__ret := C.JNIEnv_NewGlobalRef(cenv, cref)
	__v := (Jobject)(__ret)
	return __v
}

// JNIEnvDeleteGlobalRef function as declared in android/jni_call.h:37
func JNIEnvDeleteGlobalRef(env *JNIEnv, ref Jobject) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cref, _ := (C.jobject)(ref), cgoAllocsUnknown
	C.JNIEnv_DeleteGlobalRef(cenv, cref)
}

// JNIEnvDeleteLocalRef function as declared in android/jni_call.h:38
func JNIEnvDeleteLocalRef(env *JNIEnv, ref Jobject) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cref, _ := (C.jobject)(ref), cgoAllocsUnknown
	C.JNIEnv_DeleteLocalRef(cenv, cref)
}

// JNIEnvIsSameObject function as declared in android/jni_call.h:39
func JNIEnvIsSameObject(env *JNIEnv, ref1 Jobject, ref2 Jobject) Jboolean {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cref1, _ := (C.jobject)(ref1), cgoAllocsUnknown
	cref2, _ := (C.jobject)(ref2), cgoAllocsUnknown
	__ret := C.JNIEnv_IsSameObject(cenv, cref1, cref2)
	__v := (Jboolean)(__ret)
	return __v
}

// JNIEnvNewLocalRef function as declared in android/jni_call.h:41
func JNIEnvNewLocalRef(env *JNIEnv, obj Jobject) Jobject {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	__ret := C.JNIEnv_NewLocalRef(cenv, cobj)
	__v := (Jobject)(__ret)
	return __v
}

// JNIEnvEnsureLocalCapacity function as declared in android/jni_call.h:42
func JNIEnvEnsureLocalCapacity(env *JNIEnv, capacity int32) Jint {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	ccapacity, _ := (C.jint)(capacity), cgoAllocsUnknown
	__ret := C.JNIEnv_EnsureLocalCapacity(cenv, ccapacity)
	__v := (Jint)(__ret)
	return __v
}

// JNIEnvAllocObject function as declared in android/jni_call.h:44
func JNIEnvAllocObject(env *JNIEnv, clazz *Jclass) Jobject {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	__ret := C.JNIEnv_AllocObject(cenv, cclazz)
	__v := (Jobject)(__ret)
	return __v
}

// JNIEnvNewObject function as declared in android/jni_call.h:45
func JNIEnvNewObject(env *JNIEnv, clazz *Jclass, id JmethodID, args []Jvalue) Jobject {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_NewObject(cenv, cclazz, cid, cargs)
	__v := (Jobject)(__ret)
	return __v
}

// JNIEnvGetObjectClass function as declared in android/jni_call.h:47
func JNIEnvGetObjectClass(env *JNIEnv, obj Jobject) *Jclass {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	__ret := C.JNIEnv_GetObjectClass(cenv, cobj)
	__v := *(**Jclass)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvIsInstanceOf function as declared in android/jni_call.h:48
func JNIEnvIsInstanceOf(env *JNIEnv, obj Jobject, clazz *Jclass) Jboolean {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	__ret := C.JNIEnv_IsInstanceOf(cenv, cobj, cclazz)
	__v := (Jboolean)(__ret)
	return __v
}

// JNIEnvGetMethodID function as declared in android/jni_call.h:49
func JNIEnvGetMethodID(env *JNIEnv, clazz *Jclass, name string, sig string) JmethodID {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	csig, _ := unpackPCharString(sig)
	__ret := C.JNIEnv_GetMethodID(cenv, cclazz, cname, csig)
	__v := *(*JmethodID)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvCallObjectMethod function as declared in android/jni_call.h:51
func JNIEnvCallObjectMethod(env *JNIEnv, obj Jobject, id JmethodID, args []Jvalue) Jobject {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallObjectMethod(cenv, cobj, cid, cargs)
	__v := (Jobject)(__ret)
	return __v
}

// JNIEnvCallBooleanMethod function as declared in android/jni_call.h:52
func JNIEnvCallBooleanMethod(env *JNIEnv, obj Jobject, id JmethodID, args []Jvalue) Jboolean {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallBooleanMethod(cenv, cobj, cid, cargs)
	__v := (Jboolean)(__ret)
	return __v
}

// JNIEnvCallByteMethod function as declared in android/jni_call.h:53
func JNIEnvCallByteMethod(env *JNIEnv, obj Jobject, id JmethodID, args []Jvalue) Jbyte {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallByteMethod(cenv, cobj, cid, cargs)
	__v := (Jbyte)(__ret)
	return __v
}

// JNIEnvCallCharMethod function as declared in android/jni_call.h:54
func JNIEnvCallCharMethod(env *JNIEnv, obj Jobject, id JmethodID, args []Jvalue) Jchar {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallCharMethod(cenv, cobj, cid, cargs)
	__v := (Jchar)(__ret)
	return __v
}

// JNIEnvCallShortMethod function as declared in android/jni_call.h:55
func JNIEnvCallShortMethod(env *JNIEnv, obj Jobject, id JmethodID, args []Jvalue) Jshort {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallShortMethod(cenv, cobj, cid, cargs)
	__v := (Jshort)(__ret)
	return __v
}

// JNIEnvCallIntMethod function as declared in android/jni_call.h:56
func JNIEnvCallIntMethod(env *JNIEnv, obj Jobject, id JmethodID, args []Jvalue) Jint {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallIntMethod(cenv, cobj, cid, cargs)
	__v := (Jint)(__ret)
	return __v
}

// JNIEnvCallLongMethod function as declared in android/jni_call.h:57
func JNIEnvCallLongMethod(env *JNIEnv, obj Jobject, id JmethodID, args []Jvalue) Jlong {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallLongMethod(cenv, cobj, cid, cargs)
	__v := (Jlong)(__ret)
	return __v
}

// JNIEnvCallFloatMethod function as declared in android/jni_call.h:58
func JNIEnvCallFloatMethod(env *JNIEnv, obj Jobject, id JmethodID, args []Jvalue) Jfloat {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallFloatMethod(cenv, cobj, cid, cargs)
	__v := (Jfloat)(__ret)
	return __v
}

// JNIEnvCallDoubleMethod function as declared in android/jni_call.h:59
func JNIEnvCallDoubleMethod(env *JNIEnv, obj Jobject, id JmethodID, args []Jvalue) Jdouble {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallDoubleMethod(cenv, cobj, cid, cargs)
	__v := (Jdouble)(__ret)
	return __v
}

// JNIEnvCallVoidMethod function as declared in android/jni_call.h:60
func JNIEnvCallVoidMethod(env *JNIEnv, obj Jobject, id JmethodID, args []Jvalue) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	C.JNIEnv_CallVoidMethod(cenv, cobj, cid, cargs)
}

// JNIEnvCallNonvirtualObjectMethod function as declared in android/jni_call.h:62
func JNIEnvCallNonvirtualObjectMethod(env *JNIEnv, obj Jobject, clazz *Jclass, id JmethodID, args []Jvalue) Jobject {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallNonvirtualObjectMethod(cenv, cobj, cclazz, cid, cargs)
	__v := (Jobject)(__ret)
	return __v
}

// JNIEnvCallNonvirtualBooleanMethod function as declared in android/jni_call.h:63
func JNIEnvCallNonvirtualBooleanMethod(env *JNIEnv, obj Jobject, clazz *Jclass, id JmethodID, args []Jvalue) Jboolean {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallNonvirtualBooleanMethod(cenv, cobj, cclazz, cid, cargs)
	__v := (Jboolean)(__ret)
	return __v
}

// JNIEnvCallNonvirtualByteMethod function as declared in android/jni_call.h:64
func JNIEnvCallNonvirtualByteMethod(env *JNIEnv, obj Jobject, clazz *Jclass, id JmethodID, args []Jvalue) Jbyte {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallNonvirtualByteMethod(cenv, cobj, cclazz, cid, cargs)
	__v := (Jbyte)(__ret)
	return __v
}

// JNIEnvCallNonvirtualCharMethod function as declared in android/jni_call.h:65
func JNIEnvCallNonvirtualCharMethod(env *JNIEnv, obj Jobject, clazz *Jclass, id JmethodID, args []Jvalue) Jchar {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallNonvirtualCharMethod(cenv, cobj, cclazz, cid, cargs)
	__v := (Jchar)(__ret)
	return __v
}

// JNIEnvCallNonvirtualShortMethod function as declared in android/jni_call.h:66
func JNIEnvCallNonvirtualShortMethod(env *JNIEnv, obj Jobject, clazz *Jclass, id JmethodID, args []Jvalue) Jshort {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallNonvirtualShortMethod(cenv, cobj, cclazz, cid, cargs)
	__v := (Jshort)(__ret)
	return __v
}

// JNIEnvCallNonvirtualIntMethod function as declared in android/jni_call.h:67
func JNIEnvCallNonvirtualIntMethod(env *JNIEnv, obj Jobject, clazz *Jclass, id JmethodID, args []Jvalue) Jint {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallNonvirtualIntMethod(cenv, cobj, cclazz, cid, cargs)
	__v := (Jint)(__ret)
	return __v
}

// JNIEnvCallNonvirtualLongMethod function as declared in android/jni_call.h:68
func JNIEnvCallNonvirtualLongMethod(env *JNIEnv, obj Jobject, clazz *Jclass, id JmethodID, args []Jvalue) Jlong {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallNonvirtualLongMethod(cenv, cobj, cclazz, cid, cargs)
	__v := (Jlong)(__ret)
	return __v
}

// JNIEnvCallNonvirtualFloatMethod function as declared in android/jni_call.h:69
func JNIEnvCallNonvirtualFloatMethod(env *JNIEnv, obj Jobject, clazz *Jclass, id JmethodID, args []Jvalue) Jfloat {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallNonvirtualFloatMethod(cenv, cobj, cclazz, cid, cargs)
	__v := (Jfloat)(__ret)
	return __v
}

// JNIEnvCallNonvirtualDoubleMethod function as declared in android/jni_call.h:70
func JNIEnvCallNonvirtualDoubleMethod(env *JNIEnv, obj Jobject, clazz *Jclass, id JmethodID, args []Jvalue) Jdouble {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallNonvirtualDoubleMethod(cenv, cobj, cclazz, cid, cargs)
	__v := (Jdouble)(__ret)
	return __v
}

// JNIEnvCallNonvirtualVoidMethod function as declared in android/jni_call.h:71
func JNIEnvCallNonvirtualVoidMethod(env *JNIEnv, obj Jobject, clazz *Jclass, id JmethodID, args []Jvalue) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	C.JNIEnv_CallNonvirtualVoidMethod(cenv, cobj, cclazz, cid, cargs)
}

// JNIEnvGetFieldID function as declared in android/jni_call.h:73
func JNIEnvGetFieldID(env *JNIEnv, clazz *Jclass, name string, sig string) JfieldID {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	csig, _ := unpackPCharString(sig)
	__ret := C.JNIEnv_GetFieldID(cenv, cclazz, cname, csig)
	__v := *(*JfieldID)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvGetObjectField function as declared in android/jni_call.h:75
func JNIEnvGetObjectField(env *JNIEnv, obj Jobject, id JfieldID) Jobject {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetObjectField(cenv, cobj, cid)
	__v := (Jobject)(__ret)
	return __v
}

// JNIEnvGetBooleanField function as declared in android/jni_call.h:76
func JNIEnvGetBooleanField(env *JNIEnv, obj Jobject, id JfieldID) Jboolean {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetBooleanField(cenv, cobj, cid)
	__v := (Jboolean)(__ret)
	return __v
}

// JNIEnvGetByteField function as declared in android/jni_call.h:77
func JNIEnvGetByteField(env *JNIEnv, obj Jobject, id JfieldID) Jbyte {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetByteField(cenv, cobj, cid)
	__v := (Jbyte)(__ret)
	return __v
}

// JNIEnvGetCharField function as declared in android/jni_call.h:78
func JNIEnvGetCharField(env *JNIEnv, obj Jobject, id JfieldID) Jchar {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetCharField(cenv, cobj, cid)
	__v := (Jchar)(__ret)
	return __v
}

// JNIEnvGetShortField function as declared in android/jni_call.h:79
func JNIEnvGetShortField(env *JNIEnv, obj Jobject, id JfieldID) Jshort {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetShortField(cenv, cobj, cid)
	__v := (Jshort)(__ret)
	return __v
}

// JNIEnvGetIntField function as declared in android/jni_call.h:80
func JNIEnvGetIntField(env *JNIEnv, obj Jobject, id JfieldID) Jint {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetIntField(cenv, cobj, cid)
	__v := (Jint)(__ret)
	return __v
}

// JNIEnvGetLongField function as declared in android/jni_call.h:81
func JNIEnvGetLongField(env *JNIEnv, obj Jobject, id JfieldID) Jlong {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetLongField(cenv, cobj, cid)
	__v := (Jlong)(__ret)
	return __v
}

// JNIEnvGetFloatField function as declared in android/jni_call.h:82
func JNIEnvGetFloatField(env *JNIEnv, obj Jobject, id JfieldID) Jfloat {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetFloatField(cenv, cobj, cid)
	__v := (Jfloat)(__ret)
	return __v
}

// JNIEnvGetDoubleField function as declared in android/jni_call.h:83
func JNIEnvGetDoubleField(env *JNIEnv, obj Jobject, id JfieldID) Jdouble {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetDoubleField(cenv, cobj, cid)
	__v := (Jdouble)(__ret)
	return __v
}

// JNIEnvSetObjectField function as declared in android/jni_call.h:85
func JNIEnvSetObjectField(env *JNIEnv, obj Jobject, id JfieldID, val Jobject) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jobject)(val), cgoAllocsUnknown
	C.JNIEnv_SetObjectField(cenv, cobj, cid, cval)
}

// JNIEnvSetBooleanField function as declared in android/jni_call.h:86
func JNIEnvSetBooleanField(env *JNIEnv, obj Jobject, id JfieldID, val byte) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jboolean)(val), cgoAllocsUnknown
	C.JNIEnv_SetBooleanField(cenv, cobj, cid, cval)
}

// JNIEnvSetByteField function as declared in android/jni_call.h:87
func JNIEnvSetByteField(env *JNIEnv, obj Jobject, id JfieldID, val byte) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jbyte)(val), cgoAllocsUnknown
	C.JNIEnv_SetByteField(cenv, cobj, cid, cval)
}

// JNIEnvSetCharField function as declared in android/jni_call.h:88
func JNIEnvSetCharField(env *JNIEnv, obj Jobject, id JfieldID, val uint16) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jchar)(val), cgoAllocsUnknown
	C.JNIEnv_SetCharField(cenv, cobj, cid, cval)
}

// JNIEnvSetShortField function as declared in android/jni_call.h:89
func JNIEnvSetShortField(env *JNIEnv, obj Jobject, id JfieldID, val int16) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jshort)(val), cgoAllocsUnknown
	C.JNIEnv_SetShortField(cenv, cobj, cid, cval)
}

// JNIEnvSetIntField function as declared in android/jni_call.h:90
func JNIEnvSetIntField(env *JNIEnv, obj Jobject, id JfieldID, val int32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jint)(val), cgoAllocsUnknown
	C.JNIEnv_SetIntField(cenv, cobj, cid, cval)
}

// JNIEnvSetLongField function as declared in android/jni_call.h:91
func JNIEnvSetLongField(env *JNIEnv, obj Jobject, id JfieldID, val int64) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jlong)(val), cgoAllocsUnknown
	C.JNIEnv_SetLongField(cenv, cobj, cid, cval)
}

// JNIEnvSetFloatField function as declared in android/jni_call.h:92
func JNIEnvSetFloatField(env *JNIEnv, obj Jobject, id JfieldID, val float32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jfloat)(val), cgoAllocsUnknown
	C.JNIEnv_SetFloatField(cenv, cobj, cid, cval)
}

// JNIEnvSetDoubleField function as declared in android/jni_call.h:93
func JNIEnvSetDoubleField(env *JNIEnv, obj Jobject, id JfieldID, val float64) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jdouble)(val), cgoAllocsUnknown
	C.JNIEnv_SetDoubleField(cenv, cobj, cid, cval)
}

// JNIEnvGetStaticMethodID function as declared in android/jni_call.h:95
func JNIEnvGetStaticMethodID(env *JNIEnv, clazz *Jclass, name string, sig string) JmethodID {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	csig, _ := unpackPCharString(sig)
	__ret := C.JNIEnv_GetStaticMethodID(cenv, cclazz, cname, csig)
	__v := *(*JmethodID)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvCallStaticObjectMethod function as declared in android/jni_call.h:97
func JNIEnvCallStaticObjectMethod(env *JNIEnv, clazz *Jclass, id JmethodID, args []Jvalue) Jobject {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallStaticObjectMethod(cenv, cclazz, cid, cargs)
	__v := (Jobject)(__ret)
	return __v
}

// JNIEnvCallStaticBooleanMethod function as declared in android/jni_call.h:98
func JNIEnvCallStaticBooleanMethod(env *JNIEnv, clazz *Jclass, id JmethodID, args []Jvalue) Jboolean {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallStaticBooleanMethod(cenv, cclazz, cid, cargs)
	__v := (Jboolean)(__ret)
	return __v
}

// JNIEnvCallStaticByteMethod function as declared in android/jni_call.h:99
func JNIEnvCallStaticByteMethod(env *JNIEnv, clazz *Jclass, id JmethodID, args []Jvalue) Jbyte {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallStaticByteMethod(cenv, cclazz, cid, cargs)
	__v := (Jbyte)(__ret)
	return __v
}

// JNIEnvCallStaticCharMethod function as declared in android/jni_call.h:100
func JNIEnvCallStaticCharMethod(env *JNIEnv, clazz *Jclass, id JmethodID, args []Jvalue) Jchar {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallStaticCharMethod(cenv, cclazz, cid, cargs)
	__v := (Jchar)(__ret)
	return __v
}

// JNIEnvCallStaticShortMethod function as declared in android/jni_call.h:101
func JNIEnvCallStaticShortMethod(env *JNIEnv, clazz *Jclass, id JmethodID, args []Jvalue) Jshort {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallStaticShortMethod(cenv, cclazz, cid, cargs)
	__v := (Jshort)(__ret)
	return __v
}

// JNIEnvCallStaticIntMethod function as declared in android/jni_call.h:102
func JNIEnvCallStaticIntMethod(env *JNIEnv, clazz *Jclass, id JmethodID, args []Jvalue) Jint {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallStaticIntMethod(cenv, cclazz, cid, cargs)
	__v := (Jint)(__ret)
	return __v
}

// JNIEnvCallStaticLongMethod function as declared in android/jni_call.h:103
func JNIEnvCallStaticLongMethod(env *JNIEnv, clazz *Jclass, id JmethodID, args []Jvalue) Jlong {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallStaticLongMethod(cenv, cclazz, cid, cargs)
	__v := (Jlong)(__ret)
	return __v
}

// JNIEnvCallStaticFloatMethod function as declared in android/jni_call.h:104
func JNIEnvCallStaticFloatMethod(env *JNIEnv, clazz *Jclass, id JmethodID, args []Jvalue) Jfloat {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallStaticFloatMethod(cenv, cclazz, cid, cargs)
	__v := (Jfloat)(__ret)
	return __v
}

// JNIEnvCallStaticDoubleMethod function as declared in android/jni_call.h:105
func JNIEnvCallStaticDoubleMethod(env *JNIEnv, clazz *Jclass, id JmethodID, args []Jvalue) Jdouble {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	__ret := C.JNIEnv_CallStaticDoubleMethod(cenv, cclazz, cid, cargs)
	__v := (Jdouble)(__ret)
	return __v
}

// JNIEnvCallStaticVoidMethod function as declared in android/jni_call.h:106
func JNIEnvCallStaticVoidMethod(env *JNIEnv, clazz *Jclass, id JmethodID, args []Jvalue) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jmethodID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cargs, _ := (*C.jvalue)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&args)).Data)), cgoAllocsUnknown
	C.JNIEnv_CallStaticVoidMethod(cenv, cclazz, cid, cargs)
}

// JNIEnvGetStaticFieldID function as declared in android/jni_call.h:108
func JNIEnvGetStaticFieldID(env *JNIEnv, clazz *Jclass, name string, sig string) JfieldID {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cname, _ := unpackPCharString(name)
	csig, _ := unpackPCharString(sig)
	__ret := C.JNIEnv_GetStaticFieldID(cenv, cclazz, cname, csig)
	__v := *(*JfieldID)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvGetStaticObjectField function as declared in android/jni_call.h:110
func JNIEnvGetStaticObjectField(env *JNIEnv, clazz *Jclass, id JfieldID) Jobject {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetStaticObjectField(cenv, cclazz, cid)
	__v := (Jobject)(__ret)
	return __v
}

// JNIEnvGetStaticBooleanField function as declared in android/jni_call.h:111
func JNIEnvGetStaticBooleanField(env *JNIEnv, clazz *Jclass, id JfieldID) Jboolean {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetStaticBooleanField(cenv, cclazz, cid)
	__v := (Jboolean)(__ret)
	return __v
}

// JNIEnvGetStaticByteField function as declared in android/jni_call.h:112
func JNIEnvGetStaticByteField(env *JNIEnv, clazz *Jclass, id JfieldID) Jbyte {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetStaticByteField(cenv, cclazz, cid)
	__v := (Jbyte)(__ret)
	return __v
}

// JNIEnvGetStaticCharField function as declared in android/jni_call.h:113
func JNIEnvGetStaticCharField(env *JNIEnv, clazz *Jclass, id JfieldID) Jchar {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetStaticCharField(cenv, cclazz, cid)
	__v := (Jchar)(__ret)
	return __v
}

// JNIEnvGetStaticShortField function as declared in android/jni_call.h:114
func JNIEnvGetStaticShortField(env *JNIEnv, clazz *Jclass, id JfieldID) Jshort {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetStaticShortField(cenv, cclazz, cid)
	__v := (Jshort)(__ret)
	return __v
}

// JNIEnvGetStaticIntField function as declared in android/jni_call.h:115
func JNIEnvGetStaticIntField(env *JNIEnv, clazz *Jclass, id JfieldID) Jint {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetStaticIntField(cenv, cclazz, cid)
	__v := (Jint)(__ret)
	return __v
}

// JNIEnvGetStaticLongField function as declared in android/jni_call.h:116
func JNIEnvGetStaticLongField(env *JNIEnv, clazz *Jclass, id JfieldID) Jlong {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetStaticLongField(cenv, cclazz, cid)
	__v := (Jlong)(__ret)
	return __v
}

// JNIEnvGetStaticFloatField function as declared in android/jni_call.h:117
func JNIEnvGetStaticFloatField(env *JNIEnv, clazz *Jclass, id JfieldID) Jfloat {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetStaticFloatField(cenv, cclazz, cid)
	__v := (Jfloat)(__ret)
	return __v
}

// JNIEnvGetStaticDoubleField function as declared in android/jni_call.h:118
func JNIEnvGetStaticDoubleField(env *JNIEnv, clazz *Jclass, id JfieldID) Jdouble {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetStaticDoubleField(cenv, cclazz, cid)
	__v := (Jdouble)(__ret)
	return __v
}

// JNIEnvSetStaticObjectField function as declared in android/jni_call.h:120
func JNIEnvSetStaticObjectField(env *JNIEnv, clazz *Jclass, id JfieldID, val Jobject) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jobject)(val), cgoAllocsUnknown
	C.JNIEnv_SetStaticObjectField(cenv, cclazz, cid, cval)
}

// JNIEnvSetStaticBooleanField function as declared in android/jni_call.h:121
func JNIEnvSetStaticBooleanField(env *JNIEnv, clazz *Jclass, id JfieldID, val byte) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jboolean)(val), cgoAllocsUnknown
	C.JNIEnv_SetStaticBooleanField(cenv, cclazz, cid, cval)
}

// JNIEnvSetStaticByteField function as declared in android/jni_call.h:122
func JNIEnvSetStaticByteField(env *JNIEnv, clazz *Jclass, id JfieldID, val byte) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jbyte)(val), cgoAllocsUnknown
	C.JNIEnv_SetStaticByteField(cenv, cclazz, cid, cval)
}

// JNIEnvSetStaticCharField function as declared in android/jni_call.h:123
func JNIEnvSetStaticCharField(env *JNIEnv, clazz *Jclass, id JfieldID, val uint16) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jchar)(val), cgoAllocsUnknown
	C.JNIEnv_SetStaticCharField(cenv, cclazz, cid, cval)
}

// JNIEnvSetStaticShortField function as declared in android/jni_call.h:124
func JNIEnvSetStaticShortField(env *JNIEnv, clazz *Jclass, id JfieldID, val int16) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jshort)(val), cgoAllocsUnknown
	C.JNIEnv_SetStaticShortField(cenv, cclazz, cid, cval)
}

// JNIEnvSetStaticIntField function as declared in android/jni_call.h:125
func JNIEnvSetStaticIntField(env *JNIEnv, clazz *Jclass, id JfieldID, val int32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jint)(val), cgoAllocsUnknown
	C.JNIEnv_SetStaticIntField(cenv, cclazz, cid, cval)
}

// JNIEnvSetStaticLongField function as declared in android/jni_call.h:126
func JNIEnvSetStaticLongField(env *JNIEnv, clazz *Jclass, id JfieldID, val int64) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jlong)(val), cgoAllocsUnknown
	C.JNIEnv_SetStaticLongField(cenv, cclazz, cid, cval)
}

// JNIEnvSetStaticFloatField function as declared in android/jni_call.h:127
func JNIEnvSetStaticFloatField(env *JNIEnv, clazz *Jclass, id JfieldID, val float32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jfloat)(val), cgoAllocsUnknown
	C.JNIEnv_SetStaticFloatField(cenv, cclazz, cid, cval)
}

// JNIEnvSetStaticDoubleField function as declared in android/jni_call.h:128
func JNIEnvSetStaticDoubleField(env *JNIEnv, clazz *Jclass, id JfieldID, val float64) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cid, _ := *(*C.jfieldID)(unsafe.Pointer(&id)), cgoAllocsUnknown
	cval, _ := (C.jdouble)(val), cgoAllocsUnknown
	C.JNIEnv_SetStaticDoubleField(cenv, cclazz, cid, cval)
}

// JNIEnvNewString function as declared in android/jni_call.h:130
func JNIEnvNewString(env *JNIEnv, buf *uint16, bufLen int32) *Jstring {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cbuf, _ := (*C.jchar)(unsafe.Pointer(buf)), cgoAllocsUnknown
	cbufLen, _ := (C.jsize)(bufLen), cgoAllocsUnknown
	__ret := C.JNIEnv_NewString(cenv, cbuf, cbufLen)
	__v := *(**Jstring)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvGetStringLength function as declared in android/jni_call.h:131
func JNIEnvGetStringLength(env *JNIEnv, str *Jstring) Jsize {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cstr, _ := (C.jstring)(unsafe.Pointer(str)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetStringLength(cenv, cstr)
	__v := (Jsize)(__ret)
	return __v
}

// JNIEnvGetStringChars function as declared in android/jni_call.h:132
func JNIEnvGetStringChars(env *JNIEnv, str *Jstring, isCopy *byte) *Jchar {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cstr, _ := (C.jstring)(unsafe.Pointer(str)), cgoAllocsUnknown
	cisCopy, _ := (*C.jboolean)(unsafe.Pointer(isCopy)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetStringChars(cenv, cstr, cisCopy)
	__v := *(**Jchar)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvReleaseStringChars function as declared in android/jni_call.h:133
func JNIEnvReleaseStringChars(env *JNIEnv, str *Jstring, chars *uint16) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cstr, _ := (C.jstring)(unsafe.Pointer(str)), cgoAllocsUnknown
	cchars, _ := (*C.jchar)(unsafe.Pointer(chars)), cgoAllocsUnknown
	C.JNIEnv_ReleaseStringChars(cenv, cstr, cchars)
}

// JNIEnvNewStringUTF function as declared in android/jni_call.h:134
func JNIEnvNewStringUTF(env *JNIEnv, str string) *Jstring {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cstr, _ := unpackPCharString(str)
	__ret := C.JNIEnv_NewStringUTF(cenv, cstr)
	__v := *(**Jstring)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvGetStringUTFLength function as declared in android/jni_call.h:135
func JNIEnvGetStringUTFLength(env *JNIEnv, str *Jstring) Jsize {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cstr, _ := (C.jstring)(unsafe.Pointer(str)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetStringUTFLength(cenv, cstr)
	__v := (Jsize)(__ret)
	return __v
}

// JNIEnvGetStringUTFChars function as declared in android/jni_call.h:136
func JNIEnvGetStringUTFChars(env *JNIEnv, str *Jstring, isCopy *byte) string {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cstr, _ := (C.jstring)(unsafe.Pointer(str)), cgoAllocsUnknown
	cisCopy, _ := (*C.jboolean)(unsafe.Pointer(isCopy)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetStringUTFChars(cenv, cstr, cisCopy)
	__v := packPCharString(__ret)
	return __v
}

// JNIEnvReleaseStringUTFChars function as declared in android/jni_call.h:137
func JNIEnvReleaseStringUTFChars(env *JNIEnv, str *Jstring, utf string) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cstr, _ := (C.jstring)(unsafe.Pointer(str)), cgoAllocsUnknown
	cutf, _ := unpackPCharString(utf)
	C.JNIEnv_ReleaseStringUTFChars(cenv, cstr, cutf)
}

// JNIEnvGetArrayLength function as declared in android/jni_call.h:138
func JNIEnvGetArrayLength(env *JNIEnv, arr Jarray) Jsize {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jarray)(arr), cgoAllocsUnknown
	__ret := C.JNIEnv_GetArrayLength(cenv, carr)
	__v := (Jsize)(__ret)
	return __v
}

// JNIEnvNewObjectArray function as declared in android/jni_call.h:139
func JNIEnvNewObjectArray(env *JNIEnv, length int32, clazz *Jclass, obj Jobject) *JobjectArray {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	__ret := C.JNIEnv_NewObjectArray(cenv, clength, cclazz, cobj)
	__v := *(**JobjectArray)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvGetObjectArrayElement function as declared in android/jni_call.h:140
func JNIEnvGetObjectArrayElement(env *JNIEnv, arr *JobjectArray, index int32) Jobject {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jobjectArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cindex, _ := (C.jsize)(index), cgoAllocsUnknown
	__ret := C.JNIEnv_GetObjectArrayElement(cenv, carr, cindex)
	__v := (Jobject)(__ret)
	return __v
}

// JNIEnvSetObjectArrayElement function as declared in android/jni_call.h:141
func JNIEnvSetObjectArrayElement(env *JNIEnv, arr *JobjectArray, index int32, obj Jobject) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jobjectArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cindex, _ := (C.jsize)(index), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	C.JNIEnv_SetObjectArrayElement(cenv, carr, cindex, cobj)
}

// JNIEnvNewBooleanArray function as declared in android/jni_call.h:143
func JNIEnvNewBooleanArray(env *JNIEnv, length int32) *JbooleanArray {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	__ret := C.JNIEnv_NewBooleanArray(cenv, clength)
	__v := *(**JbooleanArray)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvNewByteArray function as declared in android/jni_call.h:144
func JNIEnvNewByteArray(env *JNIEnv, length int32) *JbyteArray {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	__ret := C.JNIEnv_NewByteArray(cenv, clength)
	__v := *(**JbyteArray)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvNewCharArray function as declared in android/jni_call.h:145
func JNIEnvNewCharArray(env *JNIEnv, length int32) *JcharArray {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	__ret := C.JNIEnv_NewCharArray(cenv, clength)
	__v := *(**JcharArray)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvNewShortArray function as declared in android/jni_call.h:146
func JNIEnvNewShortArray(env *JNIEnv, length int32) *JshortArray {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	__ret := C.JNIEnv_NewShortArray(cenv, clength)
	__v := *(**JshortArray)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvNewIntArray function as declared in android/jni_call.h:147
func JNIEnvNewIntArray(env *JNIEnv, length int32) *JintArray {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	__ret := C.JNIEnv_NewIntArray(cenv, clength)
	__v := *(**JintArray)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvNewLongArray function as declared in android/jni_call.h:148
func JNIEnvNewLongArray(env *JNIEnv, length int32) *JlongArray {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	__ret := C.JNIEnv_NewLongArray(cenv, clength)
	__v := *(**JlongArray)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvNewFloatArray function as declared in android/jni_call.h:149
func JNIEnvNewFloatArray(env *JNIEnv, length int32) *JfloatArray {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	__ret := C.JNIEnv_NewFloatArray(cenv, clength)
	__v := *(**JfloatArray)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvNewDoubleArray function as declared in android/jni_call.h:150
func JNIEnvNewDoubleArray(env *JNIEnv, length int32) *JdoubleArray {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	__ret := C.JNIEnv_NewDoubleArray(cenv, clength)
	__v := *(**JdoubleArray)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvGetBooleanArrayElements function as declared in android/jni_call.h:152
func JNIEnvGetBooleanArrayElements(env *JNIEnv, arr *JbooleanArray, isCopy *byte) *Jboolean {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jbooleanArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cisCopy, _ := (*C.jboolean)(unsafe.Pointer(isCopy)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetBooleanArrayElements(cenv, carr, cisCopy)
	__v := *(**Jboolean)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvGetByteArrayElements function as declared in android/jni_call.h:153
func JNIEnvGetByteArrayElements(env *JNIEnv, arr *JbyteArray, isCopy *byte) *Jbyte {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jbyteArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cisCopy, _ := (*C.jboolean)(unsafe.Pointer(isCopy)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetByteArrayElements(cenv, carr, cisCopy)
	__v := *(**Jbyte)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvGetCharArrayElements function as declared in android/jni_call.h:154
func JNIEnvGetCharArrayElements(env *JNIEnv, arr *JcharArray, isCopy *byte) *Jchar {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jcharArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cisCopy, _ := (*C.jboolean)(unsafe.Pointer(isCopy)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetCharArrayElements(cenv, carr, cisCopy)
	__v := *(**Jchar)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvGetShortArrayElements function as declared in android/jni_call.h:155
func JNIEnvGetShortArrayElements(env *JNIEnv, arr *JshortArray, isCopy *byte) *Jshort {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jshortArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cisCopy, _ := (*C.jboolean)(unsafe.Pointer(isCopy)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetShortArrayElements(cenv, carr, cisCopy)
	__v := *(**Jshort)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvGetIntArrayElements function as declared in android/jni_call.h:156
func JNIEnvGetIntArrayElements(env *JNIEnv, arr *JintArray, isCopy *byte) *Jint {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jintArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cisCopy, _ := (*C.jboolean)(unsafe.Pointer(isCopy)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetIntArrayElements(cenv, carr, cisCopy)
	__v := *(**Jint)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvGetLongArrayElements function as declared in android/jni_call.h:157
func JNIEnvGetLongArrayElements(env *JNIEnv, arr *JlongArray, isCopy *byte) *Jlong {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jlongArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cisCopy, _ := (*C.jboolean)(unsafe.Pointer(isCopy)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetLongArrayElements(cenv, carr, cisCopy)
	__v := *(**Jlong)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvGetFloatArrayElements function as declared in android/jni_call.h:158
func JNIEnvGetFloatArrayElements(env *JNIEnv, arr *JfloatArray, isCopy *byte) *Jfloat {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jfloatArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cisCopy, _ := (*C.jboolean)(unsafe.Pointer(isCopy)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetFloatArrayElements(cenv, carr, cisCopy)
	__v := *(**Jfloat)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvGetDoubleArrayElements function as declared in android/jni_call.h:159
func JNIEnvGetDoubleArrayElements(env *JNIEnv, arr *JdoubleArray, isCopy *byte) *Jdouble {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jdoubleArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cisCopy, _ := (*C.jboolean)(unsafe.Pointer(isCopy)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetDoubleArrayElements(cenv, carr, cisCopy)
	__v := *(**Jdouble)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvReleaseBooleanArrayElements function as declared in android/jni_call.h:161
func JNIEnvReleaseBooleanArrayElements(env *JNIEnv, arr *JbooleanArray, elems *byte, mode int32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jbooleanArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	celems, _ := (*C.jboolean)(unsafe.Pointer(elems)), cgoAllocsUnknown
	cmode, _ := (C.jint)(mode), cgoAllocsUnknown
	C.JNIEnv_ReleaseBooleanArrayElements(cenv, carr, celems, cmode)
}

// JNIEnvReleaseByteArrayElements function as declared in android/jni_call.h:162
func JNIEnvReleaseByteArrayElements(env *JNIEnv, arr *JbyteArray, elems *byte, mode int32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jbyteArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	celems, _ := (*C.jbyte)(unsafe.Pointer(elems)), cgoAllocsUnknown
	cmode, _ := (C.jint)(mode), cgoAllocsUnknown
	C.JNIEnv_ReleaseByteArrayElements(cenv, carr, celems, cmode)
}

// JNIEnvReleaseCharArrayElements function as declared in android/jni_call.h:163
func JNIEnvReleaseCharArrayElements(env *JNIEnv, arr *JcharArray, elems *uint16, mode int32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jcharArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	celems, _ := (*C.jchar)(unsafe.Pointer(elems)), cgoAllocsUnknown
	cmode, _ := (C.jint)(mode), cgoAllocsUnknown
	C.JNIEnv_ReleaseCharArrayElements(cenv, carr, celems, cmode)
}

// JNIEnvReleaseShortArrayElements function as declared in android/jni_call.h:164
func JNIEnvReleaseShortArrayElements(env *JNIEnv, arr *JshortArray, elems *int16, mode int32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jshortArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	celems, _ := (*C.jshort)(unsafe.Pointer(elems)), cgoAllocsUnknown
	cmode, _ := (C.jint)(mode), cgoAllocsUnknown
	C.JNIEnv_ReleaseShortArrayElements(cenv, carr, celems, cmode)
}

// JNIEnvReleaseIntArrayElements function as declared in android/jni_call.h:165
func JNIEnvReleaseIntArrayElements(env *JNIEnv, arr *JintArray, elems *int32, mode int32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jintArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	celems, _ := (*C.jint)(unsafe.Pointer(elems)), cgoAllocsUnknown
	cmode, _ := (C.jint)(mode), cgoAllocsUnknown
	C.JNIEnv_ReleaseIntArrayElements(cenv, carr, celems, cmode)
}

// JNIEnvReleaseLongArrayElements function as declared in android/jni_call.h:166
func JNIEnvReleaseLongArrayElements(env *JNIEnv, arr *JlongArray, elems *int64, mode int32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jlongArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	celems, _ := (*C.jlong)(unsafe.Pointer(elems)), cgoAllocsUnknown
	cmode, _ := (C.jint)(mode), cgoAllocsUnknown
	C.JNIEnv_ReleaseLongArrayElements(cenv, carr, celems, cmode)
}

// JNIEnvReleaseFloatArrayElements function as declared in android/jni_call.h:167
func JNIEnvReleaseFloatArrayElements(env *JNIEnv, arr *JfloatArray, elems *float32, mode int32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jfloatArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	celems, _ := (*C.jfloat)(unsafe.Pointer(elems)), cgoAllocsUnknown
	cmode, _ := (C.jint)(mode), cgoAllocsUnknown
	C.JNIEnv_ReleaseFloatArrayElements(cenv, carr, celems, cmode)
}

// JNIEnvReleaseDoubleArrayElements function as declared in android/jni_call.h:168
func JNIEnvReleaseDoubleArrayElements(env *JNIEnv, arr *JdoubleArray, elems *float64, mode int32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jdoubleArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	celems, _ := (*C.jdouble)(unsafe.Pointer(elems)), cgoAllocsUnknown
	cmode, _ := (C.jint)(mode), cgoAllocsUnknown
	C.JNIEnv_ReleaseDoubleArrayElements(cenv, carr, celems, cmode)
}

// JNIEnvGetBooleanArrayRegion function as declared in android/jni_call.h:170
func JNIEnvGetBooleanArrayRegion(env *JNIEnv, arr *JbooleanArray, start int32, length int32, buf *byte) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jbooleanArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jboolean)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_GetBooleanArrayRegion(cenv, carr, cstart, clength, cbuf)
}

// JNIEnvGetByteArrayRegion function as declared in android/jni_call.h:171
func JNIEnvGetByteArrayRegion(env *JNIEnv, arr *JbyteArray, start int32, length int32, buf *byte) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jbyteArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jbyte)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_GetByteArrayRegion(cenv, carr, cstart, clength, cbuf)
}

// JNIEnvGetCharArrayRegion function as declared in android/jni_call.h:172
func JNIEnvGetCharArrayRegion(env *JNIEnv, arr *JcharArray, start int32, length int32, buf *uint16) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jcharArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jchar)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_GetCharArrayRegion(cenv, carr, cstart, clength, cbuf)
}

// JNIEnvGetShortArrayRegion function as declared in android/jni_call.h:173
func JNIEnvGetShortArrayRegion(env *JNIEnv, arr *JshortArray, start int32, length int32, buf *int16) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jshortArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jshort)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_GetShortArrayRegion(cenv, carr, cstart, clength, cbuf)
}

// JNIEnvGetIntArrayRegion function as declared in android/jni_call.h:174
func JNIEnvGetIntArrayRegion(env *JNIEnv, arr *JintArray, start int32, length int32, buf *int32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jintArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jint)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_GetIntArrayRegion(cenv, carr, cstart, clength, cbuf)
}

// JNIEnvGetLongArrayRegion function as declared in android/jni_call.h:175
func JNIEnvGetLongArrayRegion(env *JNIEnv, arr *JlongArray, start int32, length int32, buf *int64) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jlongArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jlong)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_GetLongArrayRegion(cenv, carr, cstart, clength, cbuf)
}

// JNIEnvGetFloatArrayRegion function as declared in android/jni_call.h:176
func JNIEnvGetFloatArrayRegion(env *JNIEnv, arr *JfloatArray, start int32, length int32, buf *float32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jfloatArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jfloat)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_GetFloatArrayRegion(cenv, carr, cstart, clength, cbuf)
}

// JNIEnvGetDoubleArrayRegion function as declared in android/jni_call.h:177
func JNIEnvGetDoubleArrayRegion(env *JNIEnv, arr *JdoubleArray, start int32, length int32, buf *float64) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jdoubleArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jdouble)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_GetDoubleArrayRegion(cenv, carr, cstart, clength, cbuf)
}

// JNIEnvSetBooleanArrayRegion function as declared in android/jni_call.h:179
func JNIEnvSetBooleanArrayRegion(env *JNIEnv, arr *JbooleanArray, start int32, length int32, buf *byte) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jbooleanArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jboolean)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_SetBooleanArrayRegion(cenv, carr, cstart, clength, cbuf)
}

// JNIEnvSetByteArrayRegion function as declared in android/jni_call.h:180
func JNIEnvSetByteArrayRegion(env *JNIEnv, arr *JbyteArray, start int32, length int32, buf *byte) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jbyteArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jbyte)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_SetByteArrayRegion(cenv, carr, cstart, clength, cbuf)
}

// JNIEnvSetCharArrayRegion function as declared in android/jni_call.h:181
func JNIEnvSetCharArrayRegion(env *JNIEnv, arr *JcharArray, start int32, length int32, buf *uint16) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jcharArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jchar)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_SetCharArrayRegion(cenv, carr, cstart, clength, cbuf)
}

// JNIEnvSetShortArrayRegion function as declared in android/jni_call.h:182
func JNIEnvSetShortArrayRegion(env *JNIEnv, arr *JshortArray, start int32, length int32, buf *int16) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jshortArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jshort)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_SetShortArrayRegion(cenv, carr, cstart, clength, cbuf)
}

// JNIEnvSetIntArrayRegion function as declared in android/jni_call.h:183
func JNIEnvSetIntArrayRegion(env *JNIEnv, arr *JintArray, start int32, length int32, buf *int32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jintArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jint)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_SetIntArrayRegion(cenv, carr, cstart, clength, cbuf)
}

// JNIEnvSetLongArrayRegion function as declared in android/jni_call.h:184
func JNIEnvSetLongArrayRegion(env *JNIEnv, arr *JlongArray, start int32, length int32, buf *int64) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jlongArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jlong)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_SetLongArrayRegion(cenv, carr, cstart, clength, cbuf)
}

// JNIEnvSetFloatArrayRegion function as declared in android/jni_call.h:185
func JNIEnvSetFloatArrayRegion(env *JNIEnv, arr *JfloatArray, start int32, length int32, buf *float32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jfloatArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jfloat)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_SetFloatArrayRegion(cenv, carr, cstart, clength, cbuf)
}

// JNIEnvSetDoubleArrayRegion function as declared in android/jni_call.h:186
func JNIEnvSetDoubleArrayRegion(env *JNIEnv, arr *JdoubleArray, start int32, length int32, buf *float64) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jdoubleArray)(unsafe.Pointer(arr)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jdouble)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_SetDoubleArrayRegion(cenv, carr, cstart, clength, cbuf)
}

// JNIEnvRegisterNatives function as declared in android/jni_call.h:188
func JNIEnvRegisterNatives(env *JNIEnv, clazz *Jclass, methods *JNINativeMethod, nMethods int32) Jint {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	cmethods, _ := (*C.JNINativeMethod)(unsafe.Pointer(methods)), cgoAllocsUnknown
	cnMethods, _ := (C.jint)(nMethods), cgoAllocsUnknown
	__ret := C.JNIEnv_RegisterNatives(cenv, cclazz, cmethods, cnMethods)
	__v := (Jint)(__ret)
	return __v
}

// JNIEnvUnregisterNatives function as declared in android/jni_call.h:189
func JNIEnvUnregisterNatives(env *JNIEnv, clazz *Jclass) Jint {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cclazz, _ := (C.jclass)(unsafe.Pointer(clazz)), cgoAllocsUnknown
	__ret := C.JNIEnv_UnregisterNatives(cenv, cclazz)
	__v := (Jint)(__ret)
	return __v
}

// JNIEnvMonitorEnter function as declared in android/jni_call.h:190
func JNIEnvMonitorEnter(env *JNIEnv, obj Jobject) Jint {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	__ret := C.JNIEnv_MonitorEnter(cenv, cobj)
	__v := (Jint)(__ret)
	return __v
}

// JNIEnvMonitorExit function as declared in android/jni_call.h:191
func JNIEnvMonitorExit(env *JNIEnv, obj Jobject) Jint {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	__ret := C.JNIEnv_MonitorExit(cenv, cobj)
	__v := (Jint)(__ret)
	return __v
}

// JNIEnvGetJavaVM function as declared in android/jni_call.h:192
func JNIEnvGetJavaVM(env *JNIEnv, pVm **JavaVM) Jint {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cpVm, _ := (**C.JavaVM)(unsafe.Pointer(pVm)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetJavaVM(cenv, cpVm)
	__v := (Jint)(__ret)
	return __v
}

// JNIEnvGetStringRegion function as declared in android/jni_call.h:194
func JNIEnvGetStringRegion(env *JNIEnv, str *Jstring, start int32, length int32, buf *uint16) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cstr, _ := (C.jstring)(unsafe.Pointer(str)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.jchar)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_GetStringRegion(cenv, cstr, cstart, clength, cbuf)
}

// JNIEnvGetStringUTFRegion function as declared in android/jni_call.h:195
func JNIEnvGetStringUTFRegion(env *JNIEnv, str *Jstring, start int32, length int32, buf *byte) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cstr, _ := (C.jstring)(unsafe.Pointer(str)), cgoAllocsUnknown
	cstart, _ := (C.jsize)(start), cgoAllocsUnknown
	clength, _ := (C.jsize)(length), cgoAllocsUnknown
	cbuf, _ := (*C.char)(unsafe.Pointer(buf)), cgoAllocsUnknown
	C.JNIEnv_GetStringUTFRegion(cenv, cstr, cstart, clength, cbuf)
}

// JNIEnvGetPrimitiveArrayCritical function as declared in android/jni_call.h:197
func JNIEnvGetPrimitiveArrayCritical(env *JNIEnv, arr Jarray, isCopy *byte) unsafe.Pointer {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jarray)(arr), cgoAllocsUnknown
	cisCopy, _ := (*C.jboolean)(unsafe.Pointer(isCopy)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetPrimitiveArrayCritical(cenv, carr, cisCopy)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvReleasePrimitiveArrayCritical function as declared in android/jni_call.h:198
func JNIEnvReleasePrimitiveArrayCritical(env *JNIEnv, arr Jarray, carray unsafe.Pointer, mode int32) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	carr, _ := (C.jarray)(arr), cgoAllocsUnknown
	ccarray, _ := carray, cgoAllocsUnknown
	cmode, _ := (C.jint)(mode), cgoAllocsUnknown
	C.JNIEnv_ReleasePrimitiveArrayCritical(cenv, carr, ccarray, cmode)
}

// JNIEnvGetStringCritical function as declared in android/jni_call.h:200
func JNIEnvGetStringCritical(env *JNIEnv, str *Jstring, isCopy *byte) *Jchar {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cstr, _ := (C.jstring)(unsafe.Pointer(str)), cgoAllocsUnknown
	cisCopy, _ := (*C.jboolean)(unsafe.Pointer(isCopy)), cgoAllocsUnknown
	__ret := C.JNIEnv_GetStringCritical(cenv, cstr, cisCopy)
	__v := *(**Jchar)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvReleaseStringCritical function as declared in android/jni_call.h:201
func JNIEnvReleaseStringCritical(env *JNIEnv, str *Jstring, carray *uint16) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cstr, _ := (C.jstring)(unsafe.Pointer(str)), cgoAllocsUnknown
	ccarray, _ := (*C.jchar)(unsafe.Pointer(carray)), cgoAllocsUnknown
	C.JNIEnv_ReleaseStringCritical(cenv, cstr, ccarray)
}

// JNIEnvNewWeakGlobalRef function as declared in android/jni_call.h:203
func JNIEnvNewWeakGlobalRef(env *JNIEnv, obj Jobject) *Jweak {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	__ret := C.JNIEnv_NewWeakGlobalRef(cenv, cobj)
	__v := *(**Jweak)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvDeleteWeakGlobalRef function as declared in android/jni_call.h:204
func JNIEnvDeleteWeakGlobalRef(env *JNIEnv, obj *Jweak) {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jweak)(unsafe.Pointer(obj)), cgoAllocsUnknown
	C.JNIEnv_DeleteWeakGlobalRef(cenv, cobj)
}

// JNIEnvExceptionCheck function as declared in android/jni_call.h:206
func JNIEnvExceptionCheck(env *JNIEnv) Jboolean {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	__ret := C.JNIEnv_ExceptionCheck(cenv)
	__v := (Jboolean)(__ret)
	return __v
}

// JNIEnvNewDirectByteBuffer function as declared in android/jni_call.h:208
func JNIEnvNewDirectByteBuffer(env *JNIEnv, buf unsafe.Pointer, capacity int64) Jobject {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cbuf, _ := buf, cgoAllocsUnknown
	ccapacity, _ := (C.jlong)(capacity), cgoAllocsUnknown
	__ret := C.JNIEnv_NewDirectByteBuffer(cenv, cbuf, ccapacity)
	__v := (Jobject)(__ret)
	return __v
}

// JNIEnvGetDirectBufferAddress function as declared in android/jni_call.h:209
func JNIEnvGetDirectBufferAddress(env *JNIEnv, buf Jobject) unsafe.Pointer {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cbuf, _ := (C.jobject)(buf), cgoAllocsUnknown
	__ret := C.JNIEnv_GetDirectBufferAddress(cenv, cbuf)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// JNIEnvGetDirectBufferCapacity function as declared in android/jni_call.h:210
func JNIEnvGetDirectBufferCapacity(env *JNIEnv, buf Jobject) Jlong {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cbuf, _ := (C.jobject)(buf), cgoAllocsUnknown
	__ret := C.JNIEnv_GetDirectBufferCapacity(cenv, cbuf)
	__v := (Jlong)(__ret)
	return __v
}

// JNIEnvGetObjectRefType function as declared in android/jni_call.h:212
func JNIEnvGetObjectRefType(env *JNIEnv, obj Jobject) JobjectRefType {
	cenv, _ := (*C.JNIEnv)(unsafe.Pointer(env)), cgoAllocsUnknown
	cobj, _ := (C.jobject)(obj), cgoAllocsUnknown
	__ret := C.JNIEnv_GetObjectRefType(cenv, cobj)
	__v := (JobjectRefType)(__ret)
	return __v
}