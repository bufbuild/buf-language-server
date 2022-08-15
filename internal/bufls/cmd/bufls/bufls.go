// Copyright 2022 Buf Technologies, Inc.
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

package bufls

import (
	"context"

	"github.com/bufbuild/buf-language-server/internal/bufls"
	"github.com/bufbuild/buf-language-server/internal/bufls/cmd/bufls/command/definition"
	"github.com/bufbuild/buf-language-server/internal/bufls/cmd/bufls/command/serve"
	"github.com/bufbuild/buf/private/pkg/app/appcmd"
	"github.com/bufbuild/buf/private/pkg/app/appflag"
)

// Main is the entrypoint to the buf CLI.
func Main(name string) {
	appcmd.Main(context.Background(), NewRootCommand(name))
}

// NewRootCommand returns a new root command.
//
// This is public for use in testing.
func NewRootCommand(name string) *appcmd.Command {
	builder := appflag.NewBuilder(
		name,
		appflag.BuilderWithTracing(),
	)
	return &appcmd.Command{
		Use:                 name,
		Short:               "The Protobuf Language Server",
		Long:                "A tool that's compatible with any editor that speaks the Language Server Protocol (LSP).",
		Version:             bufls.Version,
		BindPersistentFlags: appcmd.BindMultiple(builder.BindRoot),
		SubCommands: []*appcmd.Command{
			definition.NewCommand("definition", builder),
			serve.NewCommand("serve", builder),
		},
	}
}
