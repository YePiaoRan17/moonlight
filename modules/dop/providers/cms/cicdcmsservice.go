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

package cms

import (
	context "context"
	"fmt"
	"strconv"
	"strings"

	cmspb "github.com/erda-project/erda-proto-go/core/pipeline/cms/pb"
	pb "github.com/erda-project/erda-proto-go/dop/cms/pb"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/bundle"
	"github.com/erda-project/erda/modules/dop/services/apierrors"
	"github.com/erda-project/erda/modules/dop/services/permission"
	"github.com/erda-project/erda/modules/dop/utils"
	"github.com/erda-project/erda/modules/pipeline/providers/cms"
	"github.com/erda-project/erda/pkg/common/apis"
	"github.com/erda-project/erda/providers/audit"
)

type CICDCmsService struct {
	p          *provider
	bdl        *bundle.Bundle
	permission *permission.Permission
}

func (p *CICDCmsService) WithPermission(permission *permission.Permission) {
	p.permission = permission
}

type CICDCmsCreateOrUpdateRequest struct {
	Batch         bool
	NamespaceName string
	AppID         string
	Encrypt       bool
	Configs       []*pb.Config
}

func (s *CICDCmsService) CICDCmsCreateOrUpdate(ctx context.Context, req *CICDCmsCreateOrUpdateRequest) (bool, error) {
	if req.NamespaceName == "" {
		return false, fmt.Errorf(apierrors.ErrCreateOrUpdatePipelineCmsConfigs.MissingParameter("namespace_name").Error())
	}

	appID, err := strconv.ParseUint(req.AppID, 10, 64)
	if err != nil {
		return false, fmt.Errorf(apierrors.ErrCreateOrUpdatePipelineCmsConfigs.InvalidParameter("appID error").Error())
	}

	identity := apistructs.IdentityInfo{
		UserID:         apis.GetUserID(ctx),
		InternalClient: apis.GetInternalClient(ctx),
	}

	// check permission
	if err := s.permission.CheckAppConfig(identity, appID, apistructs.UpdateAction); err != nil {
		return false, err
	}

	var updateRequest = &cmspb.CmsNsConfigsUpdateRequest{Ns: req.NamespaceName}
	var valueMap = make(map[string]*cmspb.PipelineCmsConfigValue, len(req.Configs))
	var keys []string

	for _, config := range req.Configs {
		if req.Batch && config.Type == "FILE" {
			continue
		}

		keys = append(keys, config.Key)

		var operations = &cmspb.PipelineCmsConfigOperations{}
		switch config.Type {
		case cms.ConfigTypeDiceFile:
			operations.CanDelete = true
			operations.CanDownload = true
			operations.CanEdit = true
		default:
			operations.CanDelete = true
			operations.CanDownload = false
			operations.CanEdit = true
		}
		valueMap[config.Key] = &cmspb.PipelineCmsConfigValue{
			Value:       config.Value,
			EncryptInDB: config.Encrypt,
			Type:        config.Type,
			Operations:  operations,
			Comment:     config.Comment,
		}
	}
	updateRequest.KVs = valueMap

	appInfo, err := s.bdl.GetApp(appID)
	if err != nil {
		return false, err
	}

	// get pipelineSource
	updateRequest.PipelineSource, err = s.getPipelineSource(appInfo)
	if err != nil {
		return false, err
	}

	a := s.p.audit.Begin()
	defer func() {
		a.Record(ctx, audit.AppScope, appID, string(apistructs.UpdatePipelineKeyTemplate), audit.Entries(func(ctx context.Context) (map[string]interface{}, error) {
			return map[string]interface{}{
				"projectId":   strconv.FormatUint(appInfo.ProjectID, 10),
				"appId":       strconv.FormatUint(appInfo.ID, 10),
				"projectName": appInfo.ProjectName,
				"appName":     appInfo.Name,
				"namespace":   req.NamespaceName,
				"key":         strings.Join(keys, ","),
			}, nil
		}))
	}()

	if _, err = s.p.PipelineCms.UpdateCmsNsConfigs(utils.WithInternalClientContext(ctx), updateRequest); err != nil {
		return false, err
	}

	return true, nil
}

func (s *CICDCmsService) CICDCmsUpdate(ctx context.Context, req *pb.CICDCmsUpdateRequest) (*pb.CICDCmsUpdateResponse, error) {
	ok, err := s.CICDCmsCreateOrUpdate(ctx, &CICDCmsCreateOrUpdateRequest{
		NamespaceName: req.NamespaceName,
		AppID:         req.AppID,
		Encrypt:       req.Encrypt,
		Configs:       req.Configs,
		Batch:         req.Batch,
	})
	if err != nil {
		return nil, err
	}

	if ok {
		return &pb.CICDCmsUpdateResponse{
			Data: "success",
		}, nil
	} else {
		return &pb.CICDCmsUpdateResponse{
			Data: "failed",
		}, nil
	}
}

func (s *CICDCmsService) CICDCmsCreate(ctx context.Context, req *pb.CICDCmsCreateRequest) (*pb.CICDCmsCreateResponse, error) {
	ok, err := s.CICDCmsCreateOrUpdate(ctx, &CICDCmsCreateOrUpdateRequest{
		NamespaceName: req.NamespaceName,
		AppID:         req.AppID,
		Encrypt:       req.Encrypt,
		Configs:       req.Configs,
	})
	if err != nil {
		return nil, err
	}

	if ok {
		return &pb.CICDCmsCreateResponse{
			Data: "success",
		}, nil
	} else {
		return &pb.CICDCmsCreateResponse{
			Data: "failed",
		}, nil
	}
}

func (s *CICDCmsService) CICDCmsDelete(ctx context.Context, req *pb.CICDCmsDeleteRequest) (*pb.CICDCmsDeleteResponse, error) {
	identity := apistructs.IdentityInfo{
		UserID:         apis.GetUserID(ctx),
		InternalClient: apis.GetInternalClient(ctx),
	}

	if req.Key == "" {
		return nil, fmt.Errorf(apierrors.ErrDeletePipelineCmsConfigs.MissingParameter("key").Error())
	}
	if req.NamespaceName == "" {
		return nil, fmt.Errorf(apierrors.ErrDeletePipelineCmsConfigs.MissingParameter("namespace_name").Error())
	}
	appID, err := strconv.ParseUint(req.AppID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf(apierrors.ErrDeletePipelineCmsConfigs.InvalidParameter("appID error").Error())
	}
	// check permission
	if err := s.permission.CheckAppConfig(identity, appID, apistructs.DeleteAction); err != nil {
		return nil, err
	}

	// bundle req
	var deleteReq = &cmspb.CmsNsConfigsDeleteRequest{
		Ns:         req.NamespaceName,
		DeleteKeys: []string{req.Key},
	}

	appInfo, err := s.bdl.GetApp(appID)
	if err != nil {
		return nil, err
	}

	// get pipelineSource
	deleteReq.PipelineSource, err = s.getPipelineSource(appInfo)
	if err != nil {
		return nil, fmt.Errorf(apierrors.ErrDeletePipelineCmsConfigs.InvalidParameter(err).Error())
	}

	a := s.p.audit.Begin()
	defer func() {
		a.Record(ctx, audit.AppScope, appID, string(apistructs.DeletePipelineKeyTemplate), audit.Entries(func(ctx context.Context) (map[string]interface{}, error) {
			return map[string]interface{}{
				"projectId":   strconv.FormatUint(appInfo.ProjectID, 10),
				"appId":       strconv.FormatUint(appInfo.ID, 10),
				"projectName": appInfo.ProjectName,
				"appName":     appInfo.Name,
				"namespace":   req.NamespaceName,
				"key":         req.Key,
			}, nil
		}))
	}()

	if _, err = s.p.PipelineCms.DeleteCmsNsConfigs(utils.WithInternalClientContext(ctx), deleteReq); err != nil {
		return nil, fmt.Errorf(apierrors.ErrDeletePipelineCmsNs.InternalError(err).Error())
	}

	return &pb.CICDCmsDeleteResponse{
		Data: "success",
	}, nil
}

func (s *CICDCmsService) getPipelineSource(appInfo *apistructs.ApplicationDTO) (string, error) {
	switch appInfo.Mode {
	case string(apistructs.ApplicationModeBigdata):
		return apistructs.PipelineSourceBigData.String(), nil
	default:
		return apistructs.PipelineSourceDice.String(), nil
	}
}