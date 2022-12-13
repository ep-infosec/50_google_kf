// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apps

import (
	"github.com/google/kf/v2/pkg/kf/apps"
	"github.com/google/kf/v2/pkg/kf/commands/config"
	"github.com/google/kf/v2/pkg/kf/internal/genericcli"
	"github.com/spf13/cobra"
)

// NewDeleteCommand creates a delete command.
func NewDeleteCommand(p *config.KfParams) *cobra.Command {
	cmd := genericcli.NewDeleteByNameCommand(
		apps.NewResourceInfo(),
		p,
		genericcli.WithDeleteByNameCommandName("delete"),
		genericcli.WithDeleteByNameAliases([]string{"d"}),
		genericcli.WithDeleteByNameAdditionalLongText(`
		Deleting an App will remove associated deployment history, cluster logs, and Builds.

		Things that won't be deleted:

		* source code
		* container images
		* Routes
		* ServiceInstances
		`),
	)

	return cmd
}
