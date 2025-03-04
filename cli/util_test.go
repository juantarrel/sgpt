// Copyright (c) 2023 Tim <tbckr>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//
// SPDX-License-Identifier: MIT

package cli

import (
	"testing"

	"github.com/spf13/viper"
)

type exitMemento struct {
	code int
}

func (e *exitMemento) Exit(i int) {
	e.code = i
}

func mockIsPipedShell(isPiped bool, err error) func() (bool, error) {
	return func() (bool, error) {
		return isPiped, err
	}
}

func createTestConfig(t *testing.T) *viper.Viper {
	configDir := t.TempDir()
	cacheDir := t.TempDir()
	personasDir := t.TempDir()

	config := viper.New()
	config.AddConfigPath(configDir)
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.Set("cacheDir", cacheDir)
	config.Set("personas", personasDir)
	config.Set("TESTING", 1)

	return config
}
