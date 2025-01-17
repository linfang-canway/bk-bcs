/*
Tencent is pleased to support the open source community by making Basic Service Configuration Platform available.
Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at
http://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/

package pbcs

import (
	"bscp.io/pkg/criteria/errf"
	"bscp.io/pkg/criteria/validator"
)

// Validate delete strategy request param.
func (r *DeleteStrategySetReq) Validate() error {
	if r.Id == 0 {
		return errf.New(errf.InvalidParameter, "invalid id, id should > 0")
	}

	if r.BizId == 0 {
		return errf.New(errf.InvalidParameter, "invalid biz_id, biz_id should > 0")
	}

	if r.AppId == 0 {
		return errf.New(errf.InvalidParameter, "invalid app_id, app_id should > 0")
	}

	return nil
}

// Validate 新建服务校验
func (r *CreateAppReq) Validate() error {
	if err := validator.ValidateAppName(r.Name); err != nil {
		return err
	}
	return nil
}
