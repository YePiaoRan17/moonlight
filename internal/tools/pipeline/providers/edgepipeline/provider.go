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

package edgepipeline

import (
	"context"
	"reflect"

	"github.com/ping-cloudnative/moonlight-utils/base/logs"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight/bundle"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/cancel"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/edgepipeline_register"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/edgereporter"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/providers/run"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/services/pipelinesvc"
	cronpb "github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/cron/pb"
)

type config struct {
	IsEdge      bool   `env:"DICE_IS_EDGE" default:"false"`
	ClusterName string `env:"DICE_CLUSTER_NAME"`
}

type provider struct {
	Log logs.Logger
	Cfg *config

	Cron           cronpb.CronServiceServer
	PipelineRun    run.Interface
	PipelineCancel cancel.Interface
	EdgeRegister   edgepipeline_register.Interface
	EdgeReporter   edgereporter.Interface

	bdl         *bundle.Bundle
	pipelineSvc *pipelinesvc.PipelineSvc
}

func (s *provider) Init(ctx servicehub.Context) error {
	s.bdl = bundle.New(bundle.WithClusterManager())
	return nil
}

func (s *provider) Run(ctx context.Context) error {
	return nil
}

func init() {
	interfaceType := reflect.TypeOf((*Interface)(nil)).Elem()
	servicehub.Register("edgepipeline", &servicehub.Spec{
		Services:     []string{"edgepipeline"},
		Types:        []reflect.Type{interfaceType},
		Dependencies: nil,
		Description:  "edge pipeline",
		ConfigFunc:   func() interface{} { return &config{} },
		Creator:      func() servicehub.Provider { return &provider{} },
	})
}
