// Copyright 2020 CUE Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pkg

import (
	_ "github.com/solo-io/cue/pkg/crypto/hmac"
	_ "github.com/solo-io/cue/pkg/crypto/md5"
	_ "github.com/solo-io/cue/pkg/crypto/sha1"
	_ "github.com/solo-io/cue/pkg/crypto/sha256"
	_ "github.com/solo-io/cue/pkg/crypto/sha512"
	_ "github.com/solo-io/cue/pkg/encoding/base64"
	_ "github.com/solo-io/cue/pkg/encoding/csv"
	_ "github.com/solo-io/cue/pkg/encoding/hex"
	_ "github.com/solo-io/cue/pkg/encoding/json"
	_ "github.com/solo-io/cue/pkg/encoding/yaml"
	_ "github.com/solo-io/cue/pkg/html"

	_ "github.com/solo-io/cue/pkg/list"
	_ "github.com/solo-io/cue/pkg/math"
	_ "github.com/solo-io/cue/pkg/math/bits"
	_ "github.com/solo-io/cue/pkg/net"
	_ "github.com/solo-io/cue/pkg/path"
	_ "github.com/solo-io/cue/pkg/regexp"
	_ "github.com/solo-io/cue/pkg/strconv"
	_ "github.com/solo-io/cue/pkg/strings"
	_ "github.com/solo-io/cue/pkg/struct"
	_ "github.com/solo-io/cue/pkg/text/tabwriter"
	_ "github.com/solo-io/cue/pkg/text/template"
	_ "github.com/solo-io/cue/pkg/time"
	_ "github.com/solo-io/cue/pkg/tool"
	_ "github.com/solo-io/cue/pkg/tool/cli"
	_ "github.com/solo-io/cue/pkg/tool/exec"
	_ "github.com/solo-io/cue/pkg/tool/file"
	_ "github.com/solo-io/cue/pkg/tool/http"
	_ "github.com/solo-io/cue/pkg/tool/os"
	_ "github.com/solo-io/cue/pkg/uuid"
)
