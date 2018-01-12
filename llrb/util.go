// Copyright 2010 Petar Maymounkov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package llrb

type Int int

func (x Int) Less(_ string, than Item) bool {
	return x < than.(Int)
}

type String string

func (x String) Less(_ string, than Item) bool {
	return x < than.(String)
}
