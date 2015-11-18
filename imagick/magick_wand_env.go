// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <wand/magick_wand.h>
*/
import "C"

var (
	initialized bool
)

// Inicializes the MagickWand environment
func Initialize() {
	if initialized {
		return
	}
	C.InitializeMagick(C.CString("foobar"))
	initialized = true
}

// Terminates the MagickWand environment
func Terminate() {
	if initialized {
		C.DestroyMagick()
		initialized = false
	}
}
