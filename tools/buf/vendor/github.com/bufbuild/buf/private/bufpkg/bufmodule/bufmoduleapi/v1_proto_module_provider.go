// Copyright 2020-2025 Buf Technologies, Inc.
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

package bufmoduleapi

import (
	"context"
	"fmt"
	"log/slog"

	modulev1 "buf.build/gen/go/bufbuild/registry/protocolbuffers/go/buf/registry/module/v1"
	"connectrpc.com/connect"
	"github.com/bufbuild/buf/private/bufpkg/bufregistryapi/bufregistryapimodule"
	"github.com/bufbuild/buf/private/pkg/cache"
)

// v1ProtoModuleProvider provides a per-call provider of proto Modules.
//
// We don't want to persist these across calls - this could grow over time and this cache
// isn't an LRU cache, and the information also may change over time.
type v1ProtoModuleProvider struct {
	logger               *slog.Logger
	moduleClientProvider bufregistryapimodule.V1ModuleServiceClientProvider
	protoModuleCache     cache.Cache[string, *modulev1.Module]
}

func newV1ProtoModuleProvider(
	logger *slog.Logger,
	moduleClientProvider bufregistryapimodule.V1ModuleServiceClientProvider,
) *v1ProtoModuleProvider {
	return &v1ProtoModuleProvider{
		logger:               logger,
		moduleClientProvider: moduleClientProvider,
	}
}

func (a *v1ProtoModuleProvider) getV1ProtoModuleForProtoModuleID(
	ctx context.Context,
	registry string,
	// Dashless
	protoModuleID string,
) (*modulev1.Module, error) {
	return a.protoModuleCache.GetOrAdd(
		registry+"/"+protoModuleID,
		func() (*modulev1.Module, error) {
			response, err := a.moduleClientProvider.V1ModuleServiceClient(registry).GetModules(
				ctx,
				connect.NewRequest(
					&modulev1.GetModulesRequest{
						ModuleRefs: []*modulev1.ModuleRef{
							{
								Value: &modulev1.ModuleRef_Id{
									Id: protoModuleID,
								},
							},
						},
					},
				),
			)
			if err != nil {
				return nil, err
			}
			if len(response.Msg.Modules) != 1 {
				return nil, fmt.Errorf("expected 1 Module, got %d", len(response.Msg.Modules))
			}
			return response.Msg.Modules[0], nil
		},
	)
}
