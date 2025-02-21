// Copyright © 2020 Banzai Cloud
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

package vault

import "github.com/bank-vaults/vault-sdk/vault"

// Logger is a unified interface for various logging use cases and practices, including:
//   - leveled logging
//   - structured logging
//
// See the original repository for more information: https://github.com/logur/logur
//
// Deprecated: use [vault.Logger] instead.
type Logger = vault.Logger
