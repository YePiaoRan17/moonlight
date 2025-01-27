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

package page

import (
	protocol "github.com/ping-cloudnative/moonlight/internal/core/openapi/legacy/component-protocol"
)

type ComponentPage struct {
	CtxBdl protocol.ContextBundle
	State  State `json:"state"`
	Props  Props `json:"props"`
}

type State struct {
	ActiveKey string `json:"activeKey"`
}

type Props struct {
	TabMenu []Option `json:"tabMenu"`
}

type Option struct {
	Key        string                 `json:"key"`
	Name       string                 `json:"name"`
	Operations map[string]interface{} `json:"operations"`
}

type ClickOperation struct {
	Key     string  `json:"key"`
	Reload  bool    `json:"reload"`
	Command Command `json:"command"`
}

type Command struct {
	Key          string `json:"key"`
	ScenarioType string `json:"scenarioType"`
	ScenarioKey  string `json:"scenarioKey"`
}
