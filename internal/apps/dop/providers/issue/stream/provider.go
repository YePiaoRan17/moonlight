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

package stream

import (
	"github.com/jinzhu/gorm"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/issue/core/query"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/issue/dao"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/issue/stream/core"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	"github.com/ping-cloudnative/moonlight/pkg/database/dbengine"
	"github.com/ping-cloudnative/moonlight/proto-go/dop/issue/stream/pb"
)

type config struct {
}

type provider struct {
	Cfg      *config
	Log      logs.Logger
	Register transport.Register `autowired:"service-register" required:"true"`
	DB       *gorm.DB           `autowired:"mysql-client"`
	bundle   *bundle.Bundle
	Stream   core.Interface
	Query    query.Interface

	commentIssueStreamService *CommentIssueStreamService
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.bundle = bundle.New(bundle.WithErdaServer())
	p.commentIssueStreamService = &CommentIssueStreamService{
		db: &dao.DBClient{
			DBEngine: &dbengine.DBEngine{
				DB: p.DB,
			},
		},
		logger: p.Log,
		bdl:    bundle.New(bundle.WithErdaServer()),
		stream: p.Stream,
		query:  p.Query,
	}

	if p.Register != nil {
		pb.RegisterCommentIssueStreamServiceImp(p.Register, p.commentIssueStreamService, apis.Options())
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.dop.issue.stream.CommentIssueStreamService" || ctx.Type() == pb.CommentIssueStreamServiceServerType() || ctx.Type() == pb.CommentIssueStreamServiceHandlerType():
		return p.commentIssueStreamService
	}
	return p
}

func init() {
	servicehub.Register("erda.dop.issue.stream", &servicehub.Spec{
		Services:             pb.ServiceNames(),
		Types:                pb.Types(),
		OptionalDependencies: []string{"service-register"},
		Description:          "",
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
