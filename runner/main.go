// Copyright 2016 Atelier Disko. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package runner

type Runnable interface {
	Enable() error
	Disable() error
	Clean() error
	Generate() error
}
