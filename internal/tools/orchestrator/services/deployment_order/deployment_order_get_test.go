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

package deployment_order

import (
	"context"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	"github.com/ping-cloudnative/moonlight/apistructs"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/internal/tools/orchestrator/dbclient"
	"github.com/ping-cloudnative/moonlight/internal/tools/orchestrator/services/apierrors"
	"github.com/ping-cloudnative/moonlight/proto-go/core/dicehub/release/pb"
)

func TestComposeApplicationsInfo(t *testing.T) {
	type args struct {
		AppReleases [][]*pb.ApplicationReleaseSummary
		Params      map[string]apistructs.DeploymentOrderParam
		AppsStatus  apistructs.DeploymentOrderStatusMap
	}

	appStatus := apistructs.DeploymentOrderStatusMap{
		"app1": {
			DeploymentID:     10,
			DeploymentStatus: apistructs.DeploymentStatusDeploying,
		},
	}

	params := map[string]apistructs.DeploymentOrderParam{
		"app1": []*apistructs.DeploymentOrderParamData{
			{Key: "key1", Value: "value1", Type: "ENV", Encrypt: true, Comment: "test1"},
			{Key: "key2", Value: "value2", Type: "FILE", Comment: "test2"},
		},
	}

	tests := []struct {
		name string
		args args
		want [][]*apistructs.ApplicationInfo
	}{
		{
			name: "pipeline",
			args: args{
				AppReleases: [][]*pb.ApplicationReleaseSummary{
					{
						{
							ReleaseID:       "8d2385a088df415decdf6357147ed4a2",
							ApplicationName: "app1",
						},
					},
				},
				Params:     params,
				AppsStatus: appStatus,
			},
			want: [][]*apistructs.ApplicationInfo{
				{
					{
						Name:         "app1",
						DeploymentId: 10,
						ReleaseId:    "8d2385a088df415decdf6357147ed4a2",
						Params: &apistructs.DeploymentOrderParam{
							{
								Key:     "key1",
								Value:   "",
								Encrypt: true,
								Type:    "kv",
								Comment: "test1",
							},
							{
								Key:     "key2",
								Value:   "value2",
								Type:    "dice-file",
								Comment: "test2",
							},
						},
						Status: apistructs.DeploymentStatusDeploying,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := composeApplicationsInfo(tt.args.AppReleases, tt.args.Params, tt.args.AppsStatus)
			assert.NoError(t, err)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestGetDeploymentOrderAccessDenied(t *testing.T) {
	order := New()

	defer monkey.UnpatchAll()
	monkey.PatchInstanceMethod(reflect.TypeOf(order.db), "GetDeploymentOrder", func(*dbclient.DBClient, string) (*dbclient.DeploymentOrder, error) {
		return &dbclient.DeploymentOrder{}, nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(order.bdl), "CheckPermission", func(*bundle.Bundle, *apistructs.PermissionCheckRequest) (*apistructs.PermissionCheckResponseData, error) {
		return &apistructs.PermissionCheckResponseData{
			Access: false,
		}, nil
	})

	_, err := order.Get(context.Background(), "100000", "789418c6-0bd4-4186-bd41-45372984621f")
	assert.Equal(t, err, apierrors.ErrListDeploymentOrder.AccessDenied())
}
