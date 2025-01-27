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

package dop

import (
	"context"
	"embed"
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight-utils/base/version"
	componentprotocol "github.com/ping-cloudnative/moonlight-utils/providers/component-protocol"
	"github.com/ping-cloudnative/moonlight-utils/providers/component-protocol/protocol"
	"github.com/ping-cloudnative/moonlight-utils/providers/etcd"
	"github.com/ping-cloudnative/moonlight-utils/providers/httpserver"
	"github.com/ping-cloudnative/moonlight-utils/providers/i18n"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/internal/apps/devflow/flow"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/bdl"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/component-protocol/types"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/conf"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/metrics"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/autotest/testplan"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/cms"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/devflowrule"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/guide"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/issue/core"
	issuequery "github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/issue/core/query"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/issue/stream"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/issue/sync"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/projectpipeline"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/qa/unittest"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/queue"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/rule/actions/pipeline"
	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/taskerror"
	"github.com/ping-cloudnative/moonlight/internal/core/org"
	"github.com/ping-cloudnative/moonlight/internal/pkg/metrics/query"
	"github.com/ping-cloudnative/moonlight/pkg/dumpstack"
	"github.com/ping-cloudnative/moonlight/pkg/http/httpclient"
	dashboardPb "github.com/ping-cloudnative/moonlight/proto-go/cmp/dashboard/pb"
	clusterpb "github.com/ping-cloudnative/moonlight/proto-go/core/clustermanager/cluster/pb"
	dicehubpb "github.com/ping-cloudnative/moonlight/proto-go/core/dicehub/release/pb"
	cmspb "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/cms/pb"
	cronpb "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/cron/pb"
	definitionpb "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/definition/pb"
	graphpb "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/graph/pb"
	queuepb "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/queue/pb"
	sourcepb "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/source/pb"
	errboxpb "github.com/ping-cloudnative/moonlight/proto-go/core/services/errorbox/pb"
	tokenpb "github.com/ping-cloudnative/moonlight/proto-go/core/token/pb"
	userpb "github.com/ping-cloudnative/moonlight/proto-go/core/user/pb"
	rulepb "github.com/ping-cloudnative/moonlight/proto-go/dop/rule/pb"
	addonmysqlpb "github.com/ping-cloudnative/moonlight/proto-go/orchestrator/addon/mysql/pb"
)

//go:embed component-protocol/scenarios
var scenarioFS embed.FS

type provider struct {
	Log logs.Logger

	bdl *bundle.Bundle

	Router    httpserver.Router
	RouterMgr httpserver.RouterManager

	PipelineCms           cmspb.CmsServiceServer                  `autowired:"erda.core.pipeline.cms.CmsService" optional:"true"`
	PipelineSource        sourcepb.SourceServiceServer            `autowired:"erda.core.pipeline.source.SourceService" required:"true"`
	PipelineDefinition    definitionpb.DefinitionServiceServer    `autowired:"erda.core.pipeline.definition.DefinitionService" required:"true"`
	TestPlanSvc           *testplan.TestPlanService               `autowired:"erda.core.dop.autotest.testplan.TestPlanService"`
	Cmp                   dashboardPb.ClusterResourceServer       `autowired:"erda.cmp.dashboard.resource.ClusterResource"`
	TaskErrorSvc          *taskerror.TaskErrorService             `autowired:"erda.core.dop.taskerror.TaskErrorService"`
	ErrorBoxSvc           errboxpb.ErrorBoxServiceServer          `autowired:"erda.core.services.errorbox.ErrorBoxService" optional:"true"`
	ProjectPipelineSvc    *projectpipeline.ProjectPipelineService `autowired:"erda.dop.projectpipeline.ProjectPipelineService"`
	PipelineCron          cronpb.CronServiceServer                `autowired:"erda.core.pipeline.cron.CronService" required:"true"`
	PipelineQueue         queuepb.QueueServiceServer              `autowired:"erda.core.pipeline.queue.QueueService" required:"true"`
	QueryClient           query.MetricQuery                       `autowired:"metricq-client"`
	CommentIssueStreamSvc *stream.CommentIssueStreamService       `autowired:"erda.dop.issue.stream.CommentIssueStreamService"`
	IssueSyncSvc          *sync.IssueSyncService                  `autowired:"erda.dop.issue.sync.IssueSyncService"`
	GuideSvc              *guide.GuideService                     `autowired:"erda.dop.guide.GuideService"`
	AddonMySQLSvc         addonmysqlpb.AddonMySQLServiceServer    `autowired:"erda.orchestrator.addon.mysql.AddonMySQLService"`
	DicehubReleaseSvc     dicehubpb.ReleaseServiceServer          `autowired:"erda.core.dicehub.release.ReleaseService"`
	CICDCmsSvc            *cms.CICDCmsService                     `autowired:"erda.dop.cms.CICDCmsService"`
	UnitTestService       *unittest.UnitTestService               `autowired:"erda.dop.qa.unittest.UnitTestService"`
	DevFlowRule           devflowrule.Interface
	TokenService          tokenpb.TokenServiceServer     `autowired:"erda.core.token.TokenService"`
	ClusterSvc            clusterpb.ClusterServiceServer `autowired:"erda.core.clustermanager.cluster.ClusterService"`
	DevFlowSvc            *flow.Service                  `autowired:"erda.apps.devflow.flow.FlowService"`
	IssueCoreSvc          *core.IssueService             `autowired:"erda.dop.issue.core.IssueCoreService"`
	GraphSvc              graphpb.GraphServiceServer     `autowired:"erda.core.pipeline.graph.GraphService"`
	Query                 issuequery.Interface
	Org                   org.Interface `required:"true"`
	Identity              userpb.UserServiceServer
	RuleService           rulepb.RuleServiceServer
	PipelineAction        pipeline.Interface
	Queue                 queue.Interface

	Protocol      componentprotocol.Interface
	CPTran        i18n.I18n        `autowired:"i18n"`
	ResourceTrans i18n.Translator  `translator:"resource-trans"`
	APIMTrans     i18n.Translator  `translator:"api-management-trans"`
	DB            *gorm.DB         `autowired:"mysql-client"`
	ETCD          etcd.Interface   // autowired
	EtcdClient    *clientv3.Client // autowired
}

func (p *provider) Init(ctx servicehub.Context) error {
	p.Log.Info("init dop")

	// component-protocol
	p.Log.Info("init component-protocol")
	p.Protocol.SetI18nTran(p.CPTran) // use custom i18n translator
	// compatible for legacy protocol context bundle

	metrics.Client = p.QueryClient

	bdl.Init(
		// bundle.WithDOP(), // TODO change to internal method invoke in component-protocol
		bundle.WithHepa(),
		bundle.WithOrchestrator(),
		bundle.WithGittar(),
		bundle.WithPipeline(),
		bundle.WithMonitor(),
		bundle.WithCollector(),
		bundle.WithKMS(),
		bundle.WithErdaServer(),
		bundle.WithHTTPClient(
			httpclient.New(
				httpclient.WithTimeout(time.Second, time.Duration(conf.BundleTimeoutSecond())*time.Second),
				httpclient.WithEnableAutoRetry(false),
			)),
	)
	p.bdl = bdl.Bdl
	p.Protocol.WithContextValue(types.GlobalCtxKeyBundle, bdl.Bdl)
	protocol.MustRegisterProtocolsFromFS(scenarioFS)
	p.Log.Info("init component-protocol done")

	p.Protocol.WithContextValue(types.AddonMySQLService, p.AddonMySQLSvc)
	p.Protocol.WithContextValue(types.DicehubReleaseService, p.DicehubReleaseSvc)

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000000000",
	})
	logrus.SetOutput(os.Stdout)

	dumpstack.Open()
	logrus.Infoln(version.String())

	return p.Initialize(ctx)
}

func (p *provider) Run(ctx context.Context) error {
	<-p.RouterMgr.Started()
	if err := registerWebHook(bdl.Bdl); err != nil {
		return err
	}

	if err := deleteWebhook(bdl.Bdl); err != nil {
		logrus.Errorf("failed to delete webhook, err: %v", err)
		return err
	}

	// 注册 hook
	if err := p.RegisterEvents(); err != nil {
		return err
	}

	return nil
}

func init() {
	servicehub.Register("dop", &servicehub.Spec{
		Services:     []string{"dop"},
		Dependencies: []string{"etcd", "http-server"},
		Creator:      func() servicehub.Provider { return &provider{} },
	})
}
