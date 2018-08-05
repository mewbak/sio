// Copyright 2018 The ZikiChomgo Authors. All rights reserved.  Use of this source
// code is governed by a license that can be found in the License file.

// +build !darwin !cgo
// +build !linux !cgo

package sio

import (
	"fmt"

	"github.com/irifrance/snd"
	"github.com/irifrance/snd/sample"
)

func Devices() []*Dev {
	return nil
}

func (d *Dev) Input(v snd.Form, co sample.Codec, n int) (Input, error) {
	return nil, fmt.Errorf("unsupported\n")
}

func (d *Dev) Output(v snd.Form, co sample.Codec, n int) (Output, error) {
	return nil, fmt.Errorf("unsupported\n")
}
