package app

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"unsafe"

	"android"
)

type AssetReader struct {
	mux   sync.Mutex
	asset *android.Asset
}

func (a *nativeActivity) OpenAsset(name string) (reader *AssetReader, err error) {
	a.mux.RLock()
	if a.activity == nil {
		err = errors.New("app: GetAsset must be called on initialized native activity")
	} else {
		reader = new(AssetReader)
		reader.asset = android.AssetManagerOpen(a.activity.AssetManager, name+"\x00", android.AssetModeStreaming)
		if reader.asset == nil {
			err = os.ErrNotExist
		}
	}
	a.mux.RUnlock()
	return
}

func (r *AssetReader) Read(b []byte) (n int, err error) {
	r.mux.Lock()
	if r.asset == nil {
		err = os.ErrClosed
	} else if len(b) > 0 {
		n = int(android.AssetRead(r.asset, unsafe.Pointer(&b[0]), uint32(len(b))))
		if n <= 0 {
			if n < 0 {
				err = fmt.Errorf("Read error %d", -n)
				n = 0
			} else {
				err = io.EOF
			}
			r.close()
		}
	}
	r.mux.Unlock()
	return
}

func (r *AssetReader) Close() error {
	r.mux.Lock()
	r.close()
	r.mux.Unlock()
	return nil
}

func (r *AssetReader) close() {
	if r.asset != nil {
		android.AssetClose(r.asset)
		r.asset = nil
	}
}

func (a *nativeActivity) GetAsset(name string) ([]byte, error) {
	var r, err = a.OpenAsset(name)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r)
}
