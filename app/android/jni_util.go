package android

import "errors"

type SoftKeyboardState int32

const (
	SoftKeyboardHidden SoftKeyboardState = iota
	SoftKeyboardVisible
)

func (a *NativeActivity) SetSoftKeyboardState(state SoftKeyboardState) error {
	return a.JNICall(func(env *JNIEnv, activity Jobject, activityClass, contextClass *Jclass) error {
		inputMethodService := JNIEnvGetStaticObjectField(env, contextClass,
			JNIEnvGetStaticFieldID(env, contextClass,
				s("INPUT_METHOD_SERVICE"), JClassString.Spec().Sig()),
		)
		if inputMethodService == nil {
			return errors.New("failed to get INPUT_METHOD_SERVICE")
		}

		getSystemService := JNIEnvGetMethodID(env, activityClass,
			s("getSystemService"), JNIMethodSig(JClassObject.Spec(), JClassString.Spec()))
		inputMethodManager := JNIEnvCallObjectMethod(env, activity, getSystemService, []Jvalue{
			JobjectV(inputMethodService),
		})
		if inputMethodManager == nil {
			return errors.New("failed to run getSystemService()")
		}
		defer JNIEnvDeleteLocalRef(env, inputMethodManager)

		getWindowMethod := JNIEnvGetMethodID(env, activityClass,
			s("getWindow"), JNIMethodSig(JClassWindow.Spec()))
		getDecorViewMethod := JNIEnvGetMethodID(env, JNIEnvFindClass(env, JClassWindow.Name()),
			s("getDecorView"), JNIMethodSig(JClassView.Spec()))
		window := JNIEnvCallObjectMethod(env, activity, getWindowMethod, nil)
		if window == nil {
			return errors.New("failed to run getWindow()")
		}
		defer JNIEnvDeleteLocalRef(env, window)
		decorView := JNIEnvCallObjectMethod(env, window, getDecorViewMethod, nil)
		if decorView == nil {
			return errors.New("failed to run getDecorView()")
		}
		defer JNIEnvDeleteLocalRef(env, decorView)

		switch state {
		case SoftKeyboardHidden:
			const flags = 0

			getWindowTokenMethod := JNIEnvGetMethodID(env,
				JNIEnvFindClass(env, JClassView.Name()),
				s("getWindowToken"), JNIMethodSig(JClassIBinder.Spec()))
			binder := JNIEnvCallObjectMethod(env, decorView, getWindowTokenMethod, nil)
			if binder == nil {
				return errors.New("failed to run getWindowToken()")
			}
			defer JNIEnvDeleteLocalRef(env, binder)

			hideSoftInputFromWindowMethod := JNIEnvGetMethodID(env,
				JNIEnvFindClass(env, JClassInputMethodManager.Name()),
				s("hideSoftInputFromWindow"), JNIMethodSig(JBoolean.Spec(), JClassIBinder.Spec(), JInt.Spec()))
			result := JNIEnvCallBooleanMethod(env, inputMethodManager,
				hideSoftInputFromWindowMethod, []Jvalue{
					JobjectV(binder), JintV(flags),
				})
			if result == JNIFalse {
				return errors.New("failed to run hideSoftInputFromWindow()")
			}

		case SoftKeyboardVisible:
			const flags = 0
			showSoftInputMethod := JNIEnvGetMethodID(env,
				JNIEnvFindClass(env, JClassInputMethodManager.Name()),
				s("showSoftInput"), JNIMethodSig(JBoolean.Spec(), JClassView.Spec(), JInt.Spec()))
			result := JNIEnvCallBooleanMethod(env, inputMethodManager,
				showSoftInputMethod, []Jvalue{
					JobjectV(decorView), JintV(flags),
				})
			if result == JNIFalse {
				return errors.New("failed to run showSoftInput()")
			}
		}
		return nil
	})
}

func (a *NativeActivity) KeyEventGetUnicodeChar(action, keyCode, metaState int32) (rune, error) {
	var result rune
	err := a.JNICall(func(env *JNIEnv, activity Jobject, activityClass, contextClass *Jclass) error {
		keyEventClass := JNIEnvFindClass(env, JClassKeyEvent.Name())
		keyEventInit := JNIEnvGetMethodID(env, keyEventClass,
			s("<init>"), JNIMethodSig(JVoid.Spec(), JInt.Spec(), JInt.Spec()))
		keyEvent := JNIEnvNewObject(env, keyEventClass, keyEventInit, []Jvalue{
			JintV(action), JintV(keyCode),
		})
		defer JNIEnvDeleteLocalRef(env, keyEvent)

		if keyEvent == nil {
			return errors.New("failed to construct new KeyEvent")
		}
		if metaState == 0 {
			getUnicodeCharMethod := JNIEnvGetMethodID(env, keyEventClass,
				s("getUnicodeChar"), JNIMethodSig(JInt.Spec()))
			result = rune(JNIEnvCallIntMethod(env, keyEvent, getUnicodeCharMethod, nil))
		} else {
			getUnicodeCharMethod := JNIEnvGetMethodID(env, keyEventClass,
				s("getUnicodeChar"), JNIMethodSig(JInt.Spec(), JInt.Spec()))
			result = rune(JNIEnvCallIntMethod(env, keyEvent, getUnicodeCharMethod, []Jvalue{
				JintV(metaState),
			}))
		}
		return nil
	})
	return result, err
}
