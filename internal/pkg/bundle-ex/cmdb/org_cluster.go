// Copyright (c) 2021 Terminus, Inc.
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

package cmdb

import (
	"context"

	"google.golang.org/grpc/metadata"

	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight/pkg/http/httputil"
	orgpb "github.com/ping-cloudnative/moonlight/proto-go/core/org/pb"
)

// GetOrgClusterRelationsByOrg 获取所有的企业集群关联关系
func (c *Cmdb) GetOrgClusterRelationsByOrg(ctx context.Context, orgID string) (*orgpb.GetOrgClusterRelationsByOrgResponse, error) {
	ctx = transport.WithHeader(ctx, metadata.New(map[string]string{httputil.InternalHeader: "true"}))
	return c.orgServer.GetOrgClusterRelationsByOrg(ctx, &orgpb.GetOrgClusterRelationsByOrgRequest{OrgID: orgID})
}
