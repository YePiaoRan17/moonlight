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

package openapi_rule

import (
	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/pkg/transport"
	"github.com/ping-cloudnative/moonlight/internal/tools/orchestrator/hepa/common"
	"github.com/ping-cloudnative/moonlight/internal/tools/orchestrator/hepa/services/openapi_rule/impl"
	zoneI "github.com/ping-cloudnative/moonlight/internal/tools/orchestrator/hepa/services/zone/impl"
	"github.com/ping-cloudnative/moonlight/pkg/common/apis"
	perm "github.com/ping-cloudnative/moonlight/pkg/common/permission"
	"github.com/ping-cloudnative/moonlight/proto-go/core/hepa/openapi_rule/pb"
)

type config struct {
}

// +provider
type provider struct {
	Cfg                *config
	Log                logs.Logger
	Register           transport.Register
	openapiRuleService *openapiRuleService
	Perm               perm.Interface `autowired:"permission"`
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.openapiRuleService = &openapiRuleService{p}
	err := zoneI.NewGatewayZoneServiceImpl()
	if err != nil {
		return err
	}
	err = impl.NewGatewayOpenapiRuleServiceImpl()
	if err != nil {
		return err
	}
	if p.Register != nil {
		type ruleService = pb.OpenapiRuleServiceServer
		pb.RegisterOpenapiRuleServiceImp(p.Register, p.openapiRuleService, apis.Options(), p.Perm.Check(
			perm.Method(ruleService.CreateLimit, perm.ScopeOrg, "org", perm.ActionGet, perm.OrgIDValue()),
			perm.Method(ruleService.DeleteLimit, perm.ScopeOrg, "org", perm.ActionGet, perm.OrgIDValue()),
			perm.Method(ruleService.GetLimits, perm.ScopeOrg, "org", perm.ActionGet, perm.OrgIDValue()),
			perm.Method(ruleService.UpdateLimit, perm.ScopeOrg, "org", perm.ActionGet, perm.OrgIDValue()),
		), common.AccessLogWrap(common.AccessLog))
	}
	return nil
}

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	switch {
	case ctx.Service() == "erda.core.hepa.openapi_rule.OpenapiRuleService" || ctx.Type() == pb.OpenapiRuleServiceServerType() || ctx.Type() == pb.OpenapiRuleServiceHandlerType():
		return p.openapiRuleService
	}
	return p
}

func init() {
	servicehub.Register("erda.core.hepa.openapi_rule", &servicehub.Spec{
		Services:             pb.ServiceNames(),
		Types:                pb.Types(),
		OptionalDependencies: []string{"service-register"},
		Dependencies: []string{
			"hepa",
			"erda.core.hepa.domain.DomainService",
		},
		Description: "",
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
