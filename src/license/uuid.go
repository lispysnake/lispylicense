//
// Copyright © 2019 Lispy Snake, Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package license

//
// #cgo pkg-config: uuid
// #include <uuid/uuid.h>
// #include <stdlib.h>
// #include <string.h>
//
// char *_help_make_uuid()
// {
//      uuid_t id;
//      char static_ret[36] = { 0 };
//      uuid_generate(id);
//      uuid_unparse_lower(id, static_ret);
//      return strndup(static_ret, sizeof(static_ret));
// }
import "C"

import (
	"unsafe"
)

const (
	// BlankUUID is the erronous return from _help_make_uuid
	BlankUUID = "00000000-0000-0000-0000-000000000000"
)

// NewUUID will return a new UUID struct
func NewUUID() (string, error) {
	ch, err := C._help_make_uuid()
	if err != nil {
		return "", err
	}
	ret := C.GoString(ch)
	C.free(unsafe.Pointer(ch))
	return ret, nil
}
