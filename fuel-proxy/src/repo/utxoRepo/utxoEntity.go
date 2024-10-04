package utxoRepo

import (
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_scalars"
	"time"
)

type UtxoEntity struct {
	InternalUpdatedAtUnixMs int64  `redis:"InternalUpdatedAtUnixMs"`
	IsSpent                 bool   `redis:"IsSpent"`
	TxId                    string `redis:"TxId"`
	TxOutputIndex           string `redis:"TxOutputIndex"`
	Amount                  uint64 `redis:"Amount"`
	Owner                   string `redis:"Owner"`
	AssetId                 string `redis:"AssetId"`
	BlockCreated            uint64 `redis:"BlockCreated"`
	TxCreatedIdx            uint64 `redis:"TxCreatedIdx"`
}

func NewUtxoEntity(txId string, txOutputIndex string, owner string, assetId string, amount uint64, blockCreated uint64, txCreatedIdx uint64) *UtxoEntity {
	e := &UtxoEntity{
		TxId:          txId,
		TxOutputIndex: txOutputIndex,
		Owner:         owner,
		AssetId:       assetId,
		Amount:        amount,
		BlockCreated:  blockCreated,
		TxCreatedIdx:  txCreatedIdx,
	}
	e.UpdateInternalUpdatedAt()
	return e
}

// UtxoId 34 bytes utxo id build from TxId+TxOutputIndex
func (w *UtxoEntity) UtxoId() (*graphql_scalars.Bytes34, error) {
	txId, err := graphql_scalars.NewBytes32TryFromString(w.TxId)
	if err != nil {
		return nil, err
	}
	txOutputIndex, err := graphql_scalars.NewBytes32TryFromString(w.TxOutputIndex)
	if err != nil {
		return nil, err
	}
	txIdSlice := txId.Val()
	txOutputIndexSlice := txOutputIndex.Val()
	utxoId, err := graphql_scalars.NewBytes34TryFromSlice(append(txIdSlice[:], txOutputIndexSlice[30:]...))
	if err != nil {
		return nil, err
	}
	return utxoId, nil
}

func (w *UtxoEntity) SetIsSpent(v bool) {
	w.IsSpent = v
}

func (w *UtxoEntity) UpdateInternalUpdatedAt() {
	w.InternalUpdatedAtUnixMs = time.Now().UnixMilli()
}
