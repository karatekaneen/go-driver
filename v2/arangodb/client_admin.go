//
// DISCLAIMER
//
// Copyright 2023 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

package arangodb

import (
	"context"
)

type ClientAdmin interface {
	ClientAdminLog
	ClientAdminBackup
	// Health returns the cluster configuration & health.
	// It works in cluster or active fail-over mode.
	Health(ctx context.Context) (ClusterHealth, error)
}

type ClientAdminBackup interface {
}

type ClientAdminLog interface {
	// GetLogLevels returns log levels for topics.
	GetLogLevels(ctx context.Context, opts *LogLevelsGetOptions) (LogLevels, error)
	// SetLogLevels sets log levels for a given topics.
	SetLogLevels(ctx context.Context, logLevels LogLevels, opts *LogLevelsSetOptions) error
}