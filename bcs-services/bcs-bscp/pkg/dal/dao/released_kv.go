/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dao

import (
	"fmt"

	"bscp.io/pkg/dal/gen"
	"bscp.io/pkg/dal/table"
	"bscp.io/pkg/kit"
)

// ReleasedKv supplies all the released kv related operations.
type ReleasedKv interface {
	// BulkCreateWithTx bulk create released config items with tx.
	BulkCreateWithTx(kit *kit.Kit, tx *gen.QueryTx, kvs []*table.ReleasedKv) error
}

var _ ReleasedKv = new(releasedKvDao)

type releasedKvDao struct {
	genQ     *gen.Query
	idGen    IDGenInterface
	auditDao AuditDao
}

func (dao *releasedKvDao) BulkCreateWithTx(kit *kit.Kit, tx *gen.QueryTx, kvs []*table.ReleasedKv) error {

	if len(kvs) == 0 {
		return nil
	}

	// validate released config item field.
	for _, kv := range kvs {
		if err := kv.ValidateCreate(); err != nil {
			return err
		}
	}

	// generate released config items id.
	ids, err := dao.idGen.Batch(kit, table.ReleasedKvTable, len(kvs))
	if err != nil {
		return err
	}

	start := 0
	for _, kv := range kvs {
		kv.ID = ids[start]
		start++
	}
	batchSize := 100

	q := tx.ReleasedKv.WithContext(kit.Ctx)
	if err := q.CreateInBatches(kvs, batchSize); err != nil {
		return fmt.Errorf("create released kv in batch failed, err: %v", err)
	}

	ad := dao.auditDao.DecoratorV2(kit, kvs[0].Attachment.BizID).PrepareCreate(table.RkvList(kvs))
	if err := ad.Do(tx.Query); err != nil {
		return err
	}

	return nil

}
