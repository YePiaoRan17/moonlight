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

package basic

import (
	"time"

	"github.com/ping-cloudnative/moonlight/apistructs"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/aop"
	"github.com/ping-cloudnative/moonlight/internal/tools/pipeline/aop/aoptypes"
	"github.com/ping-cloudnative/moonlight-utils/base/servicehub"
	"github.com/ping-cloudnative/moonlight/proto-go/core/pipeline/report/pb"
)

// +provider
type provider struct {
	aoptypes.PipelineBaseTunePoint
}

func (p *provider) Name() string { return "basic" }

func (p *provider) Handle(ctx *aoptypes.TuneContext) error {
	pipeline := ctx.SDK.Pipeline

	// make report content
	var content apistructs.PipelineBasicReport
	content.ID = pipeline.ID
	content.Status = pipeline.Status
	content.PipelineSource = pipeline.PipelineSource
	content.PipelineYmlName = pipeline.PipelineYmlName
	content.ClusterName = pipeline.ClusterName
	content.TimeCreated = pipeline.TimeCreated
	content.TimeBegin = pipeline.TimeBegin
	content.TimeEnd = pipeline.TimeEnd
	content.TotalCostTimeSec = pipeline.CostTimeSec

	tasks, err := ctx.SDK.DBClient.ListPipelineTasksByPipelineID(pipeline.ID)
	if err != nil {
		return err
	}
	for _, task := range tasks {
		taskReport := apistructs.TaskReportInfo{
			ID:               task.ID,
			Status:           task.Status,
			Name:             task.Name,
			ActionType:       task.Type,
			ActionVersion:    task.Extra.Action.Version,
			ExecutorType:     string(task.ExecutorKind),
			ClusterName:      task.Extra.ClusterName,
			TimeBegin:        getTimeOrNil(task.TimeBegin),
			TimeEnd:          getTimeOrNil(task.TimeEnd),
			TimeBeginQueue:   getTimeOrNil(task.Extra.TimeBeginQueue),
			TimeEndQueue:     getTimeOrNil(task.Extra.TimeEndQueue),
			QueueCostTimeSec: task.QueueTimeSec,
			RunCostTimeSec:   task.CostTimeSec,
			MachineStat:      task.Inspect.MachineStat,
			Meta: func() map[string]string {
				result := make(map[string]string)
				for _, meta := range task.GetMetadata() {
					result[meta.Name] = meta.Value
				}
				return result
			}(),
		}
		content.TaskInfos = append(content.TaskInfos, taskReport)
	}

	pbMeta, err := ctx.SDK.Report.MakePBMeta(map[string]interface{}{
		"data": content,
	})
	if err != nil {
		return err
	}
	_, err = ctx.SDK.Report.Create(&pb.PipelineReportCreateRequest{
		PipelineID: ctx.SDK.Pipeline.ID,
		Type:       string(apistructs.PipelineReportTypeBasic),
		Meta:       pbMeta,
	})
	return err
}

func getTimeOrNil(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}

func (p *provider) Init(ctx servicehub.Context) error {
	err := aop.RegisterTunePoint(p)
	if err != nil {
		panic(err)
	}
	return nil
}

func init() {
	servicehub.Register(aop.NewProviderNameByPluginName(&provider{}), &servicehub.Spec{
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
