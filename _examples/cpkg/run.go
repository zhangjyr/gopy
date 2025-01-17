// Copyright 2015 The zhangjyr Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"fmt"

	"github.com/zhangjyr/gopy/_examples/cpkg"
)

func main() {
	fmt.Printf("hello from go\n")
	cpkg.Hello("me")
	fmt.Printf("bye me\n")
	cpkg.Hello("you")
	fmt.Printf("bye you\n")
}
