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

package query

import (
	"reflect"
	"testing"

	"github.com/ping-cloudnative/moonlight/internal/apps/dop/providers/issue/dao"
)

func Test_getStageMap(t *testing.T) {
	type args struct {
		stages []dao.IssueStage
	}
	w := map[IssueStage]string{
		{"TASK", "1"}: "2",
	}
	tests := []struct {
		name string
		args args
		want map[IssueStage]string
	}{
		{
			args: args{
				[]dao.IssueStage{{IssueType: "TASK", Value: "1", Name: "2"}},
			},
			want: w,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStageMap(tt.args.stages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStageMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
