// Copyright 2019 The zhangjyr Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package main

import "github.com/zhangjyr/gopy/bind"

const (
	libExt       = ".pyd"
	extraGccArgs = ""
)

func init() {
	bind.WindowsOS = true
}
