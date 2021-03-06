// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pms

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRawTokenReader(t *testing.T) {
	raw := "token"
	reader := &RawTokenReader{
		Token: raw,
	}

	token, err := reader.ReadToken()
	require.Nil(t, err)
	assert.Equal(t, raw, token)
}

func TestReadToken(t *testing.T) {
	// nil reader
	_, err := readToken(nil)
	assert.NotNil(t, err)

	// empty
	reader := bytes.NewReader([]byte{})
	_, err = readToken(reader)
	assert.NotNil(t, err)

	// contains no "access_token"
	content := "key1=value\nkey2=value2"
	reader = bytes.NewReader([]byte(content))
	_, err = readToken(reader)
	assert.NotNil(t, err)

	// contains "access_token" but no "="
	content = "access_token value\nkey2=value2"
	reader = bytes.NewReader([]byte(content))
	_, err = readToken(reader)
	assert.NotNil(t, err)

	// contains "access_token" and "=", but no value
	content = "access_token=\nkey2=value2"
	reader = bytes.NewReader([]byte(content))
	token, err := readToken(reader)
	require.Nil(t, err)
	assert.Len(t, token, 0)

	// valid "access_token"
	content = "access_token=token\nkey2=value2"
	reader = bytes.NewReader([]byte(content))
	token, err = readToken(reader)
	require.Nil(t, err)
	assert.Equal(t, "token", token)
}

func TestFileTokenReader(t *testing.T) {
	// file not exist
	path := "/tmp/not_exist_file"
	reader := &FileTokenReader{
		Path: path,
	}

	_, err := reader.ReadToken()
	assert.NotNil(t, err)

	// file exist
	path = "/tmp/exist_file"
	err = ioutil.WriteFile(path, []byte("access_token=token"), 0x0666)
	require.Nil(t, err)
	defer os.Remove(path)

	reader = &FileTokenReader{
		Path: path,
	}

	token, err := reader.ReadToken()
	require.Nil(t, err)
	assert.Equal(t, "token", token)
}
