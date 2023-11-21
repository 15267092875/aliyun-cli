// Copyright (c) 2009-present, Alibaba Cloud All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package i18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLibrary(t *testing.T) {
	//Test T(en string, zh string)*Text
	text := T("hello", "你好")
	assert.Equal(t, "hello", text.dic["en"])
	assert.Equal(t, "你好", text.dic["zh"])

	text = T("", "你好")
	assert.Equal(t, "", text.dic["en"])
	assert.Equal(t, "你好", text.dic["zh"])

	text = T("hello", "")
	assert.Equal(t, "hello", text.dic["en"])
	assert.Equal(t, "", text.dic["zh"])
}
