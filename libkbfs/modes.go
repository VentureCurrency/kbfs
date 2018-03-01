// Copyright 2018 Keybase Inc. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

package libkbfs

import (
	"fmt"
)

func NewInitModeFromType(t InitModeType) InitMode {
	switch t.Mode() {
	case InitModeDefault:
		return modeDefault{t}
	case InitModeMinimal:
		return modeMinimal{t}
	case InitModeSingleOp:
		return modeSingleOp{modeDefault{t}}
	case InitModeConstrained:
		return modeConstrained{modeMinimal{t}}
	default:
		panic(fmt.Sprintf("Unknown mode: %s", t))
	}
}

// Default mode:

type modeDefault struct {
	originalMode InitModeType
}

func (md modeDefault) Type() InitModeType {
	return InitModeDefault
}

func (md modeDefault) BlockWorkers() int {
	return defaultBlockRetrievalWorkerQueueSize
}

func (md modeDefault) PrefetchWorkers() {
	return defaultPrefetchWorkerQueueSize
}

// Minimal mode:

type modeMinimal struct {
	originalMode InitModeType
}

func (md modeMinimal) Type() InitModeType {
	return InitModeMinimal
}

func (md modeMinimal) BlockWorkers() int {
	return 0
}

func (md modeMinimal) PrefetchWorkers() {
	return 0
}

// Single op mode:

type modeSingleOp struct {
	InitMode
}

func (md modeSingleOp) Type() InitModeType {
	return InitModeSingleOp
}

// Constrained mode:

type modeConstrained struct {
	InitMode
}

func (md modeConstrained) Type() InitModeType {
	return InitModeConstrained
}

func (md modeMinimal) BlockWorkers() int {
	return 1
}
