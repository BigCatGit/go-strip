// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package buildinfo provides access to information embedded in a Go binary
// about how it was built. This includes the Go toolchain version, and the
// set of modules used (for binaries built in module mode).
//
// Build information is available for the currently running binary in
// runtime/debug.ReadBuildInfo.
package gore

import (
	"fmt"
	"gostrip/gore/extern"
	"runtime/debug"
)

// BuildInfo that was extracted from the file.
type BuildInfo struct {
	// Compiler version. Can be nil.
	Compiler *GoVersion
	// ModInfo holds information about the Go modules in this file.
	// Can be nil.
	ModInfo *debug.BuildInfo
}

func (f *GoFile) extractBuildInfo() (*BuildInfo, error) {
	info, err := extern.Read(f.fh.getFile())
	if err != nil {
		return nil, fmt.Errorf("error when extracting build information: %w", err)
	}

	result := &BuildInfo{
		Compiler: ResolveGoVersion(info.GoVersion),
		ModInfo:  info,
	}
	return result, nil
}
