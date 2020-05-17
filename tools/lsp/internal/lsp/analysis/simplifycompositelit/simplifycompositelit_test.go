// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package simplifycompositelit_test

import (
	"testing"

	"cuelang.org/go/tools/lsp/internal/lsp/analysis/simplifycompositelit"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, simplifycompositelit.Analyzer, "a")
}
