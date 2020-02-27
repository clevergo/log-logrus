// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package logrusadapter

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNew(t *testing.T) {
	logrusLogger := logrus.New()
	_ := New(logrusLogger)
}
