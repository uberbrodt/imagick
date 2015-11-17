// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <wand/magick_wand.h>
*/
import "C"

type AffineMatrix struct {
	am *C.AffineMatrix
}
