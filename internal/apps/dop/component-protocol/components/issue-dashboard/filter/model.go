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

package filter

import (
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/cptype"
	"github.com/ping-cloudnative/moonlight/apistructs"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/component-protocol/components/issue-dashboard/common"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/issue/core/query"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/issue/dao"
	"github.com/ping-cloudnative/moonlight/internal/core/openapi/legacy/component-protocol/components/filter"
	"github.com/ping-cloudnative/moonlight/proto-go/dop/issue/core/pb"
)

type ComponentFilter struct {
	sdk      *cptype.SDK
	bdl      *bundle.Bundle
	issueSvc query.Interface
	filter.CommonFilter
	State    State    `json:"state,omitempty"`
	InParams InParams `json:"-"`

	// local vars
	Iterations     []apistructs.Iteration `json:"-"`
	Members        []apistructs.Member    `json:"-"`
	IssueList      []dao.IssueItem        `json:"-"`
	IssueStateList []dao.IssueState       `json:"-"`
	Stages         []*pb.IssueStage       `json:"-"`
}

type State struct {
	Conditions           []filter.PropCondition    `json:"conditions,omitempty"`
	Values               common.FrontendConditions `json:"values,omitempty"`
	Base64UrlQueryParams string                    `json:"filter__urlQuery,omitempty"`
}

type InParams struct {
	FrontEndProjectID      string `json:"projectId,omitempty"`
	FrontendUrlQuery       string `json:"filter__urlQuery,omitempty"`
	ProjectID              uint64
	FrontendFixedIteration string `json:"fixedIteration,omitempty"`
	IterationID            int64
}

const OperationKeyFilter filter.OperationKey = "filter"
