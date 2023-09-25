/*-
 * ‌
 * Hedera Mirror Node
 * ​
 * Copyright (C) 2019 - 2023 Hedera Hashgraph, LLC
 * ​
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * ‍
 */

package domain

import (
	"github.com/hashgraph/hedera-mirror-node/hedera-mirror-rosetta/app/interfaces"
	"github.com/hashgraph/hedera-mirror-node/hedera-mirror-rosetta/app/persistence/domain"
)

type AccountBalanceSnapshotBuilder struct {
	accountBalances    []domain.AccountBalance
	consensusTimestamp int64
	dbClient           interfaces.DbClient
}

func (b *AccountBalanceSnapshotBuilder) AddAccountBalance(accountId, balance int64) *AccountBalanceSnapshotBuilder {
	b.accountBalances = append(b.accountBalances, domain.AccountBalance{
		AccountId:          domain.MustDecodeEntityId(accountId),
		Balance:            balance,
		ConsensusTimestamp: b.consensusTimestamp,
	})
	return b
}

func (b *AccountBalanceSnapshotBuilder) Persist() {
	db := b.dbClient.GetDb()
	if len(b.accountBalances) != 0 {
		db.Create(b.accountBalances)
	}
}

func NewAccountBalanceSnapshotBuilder(dbClient interfaces.DbClient, consensusTimestamp int64) *AccountBalanceSnapshotBuilder {
	return &AccountBalanceSnapshotBuilder{
		accountBalances:    make([]domain.AccountBalance, 0),
		consensusTimestamp: consensusTimestamp,
		dbClient:           dbClient,
	}
}