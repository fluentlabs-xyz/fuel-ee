package utxoRepo

import (
	"context"
	"github.com/fluentlabs-xyz/fuel-ee/src/repo"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type UtxoRepo struct {
	redisClient *redis.Client
}

func NewUtxoRepo(redisClient *redis.Client) *UtxoRepo {
	r := UtxoRepo{
		redisClient: redisClient,
	}

	return &r
}

func (r *UtxoRepo) SaveMany(ctx context.Context, entities []*UtxoEntity) error {
	if entities == nil || len(entities) <= 0 {
		return errors.Errorf("param [entities] must be set and have at least 1 element")
	}
	if _, err := r.redisClient.Pipelined(ctx, func(p redis.Pipeliner) error {
		for _, w := range entities {
			w.UpdateInternalUpdatedAt()
			err := r.saveOneUsingPipeliner(ctx, p, w)
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *UtxoRepo) SaveOne(ctx context.Context, w *UtxoEntity) error {
	return r.SaveMany(ctx, []*UtxoEntity{w})
}

//func (r *utxoRepo) FindByEntityId(ctx context.Context, entityId string) (*UtxoEntity, error) {
//	entity := &UtxoEntity{}
//	key := r.GenerateKeyUsingParams(entityId)
//	stringStringMapCmd := r.redisClient.HGetAll(ctx, key)
//	if stringStringMapCmd.Err() != nil {
//		return nil, stringStringMapCmd.Err()
//	}
//	if len(stringStringMapCmd.Val()) <= 0 {
//		return nil, nil
//	}
//	if err := r.redisClient.HGetAll(ctx, key).Scan(entity); err != nil {
//		return nil, err
//	}
//
//	return entity, nil
//}

// FindAllByParams you can '*' as 'any' mask for a parameter
func (r *UtxoRepo) FindAllByParams(ctx context.Context, ownerId, TxId, TxOutputIdx string, includeSpent bool) (map[string]*UtxoEntity, error) {
	// use map to prevent possible entity duplication from redis
	entities := make(map[string]*UtxoEntity)

	scanKey := r.GenerateKeyUsingParams(ownerId, TxId, TxOutputIdx)
	var cursor uint64 = 0
	// TODO implement pagination to prevent long-term blocking
	iter := r.redisClient.Scan(ctx, cursor, scanKey, 0).Iterator()
	if iter.Err() != nil {
		return nil, iter.Err()
	}
	for iter.Next(ctx) {
		w := &UtxoEntity{}
		stringStringMapCmd := r.redisClient.HGetAll(ctx, iter.Val())
		//if err := stringStringMapCmd.Scan(w); err != nil {
		//	return nil, errors.Errorf("failed, reason '%s'", err)
		//}
		err := r.scanDataInto(stringStringMapCmd.Val(), w)
		if err != nil {
			return nil, err
		}
		if w.IsSpent && !includeSpent {
			continue
		}
		wUniqueId := r.GenerateKey(w)
		entities[wUniqueId] = w
	}

	return entities, nil
}

func (r *UtxoRepo) LastProcessedBlockNumber(ctx context.Context) (uint64, error) {
	stringCmd := r.redisClient.Get(ctx, repo.UtxoLastProcessedBlockHashmapKeyTemplate)
	if stringCmd.Err() != nil {
		if stringCmd.Err().Error() == "redis: nil" {
			return 0, nil
		}
		return 0, stringCmd.Err()
	}
	blockNumber, err := strconv.ParseUint(stringCmd.Val(), 10, 64)
	if err != nil {
		return 0, err
	}

	return blockNumber, nil
}

func (r *UtxoRepo) SaveLastProcessedBlockNumber(ctx context.Context, v uint64) error {
	stringCmd := r.redisClient.Set(ctx, repo.UtxoLastProcessedBlockHashmapKeyTemplate, strconv.FormatUint(v, 10), 0)
	if stringCmd.Err() != nil {
		return stringCmd.Err()
	}
	return nil
}

// TODO use reflection to extract field names
func (r *UtxoRepo) saveOneUsingPipeliner(ctx context.Context, p redis.Pipeliner, w *UtxoEntity) error {
	if w == nil {
		return errors.Errorf("param [w] must be set")
	}
	key := r.GenerateKey(w)
	p.HSet(ctx, key, "IsSpent", strconv.FormatBool(w.IsSpent))
	p.HSet(ctx, key, "TxId", strings.ToLower(w.TxId))
	p.HSet(ctx, key, "TxOutputIndex", strings.ToLower(w.TxOutputIndex))
	p.HSet(ctx, key, "Amount", strconv.FormatUint(w.Amount, 10))
	p.HSet(ctx, key, "Owner", strings.ToLower(w.Owner))
	p.HSet(ctx, key, "AssetId", w.AssetId)
	p.HSet(ctx, key, "BlockCreated", strconv.FormatUint(w.BlockCreated, 10))
	p.HSet(ctx, key, "TxCreatedIdx", strconv.FormatUint(w.TxCreatedIdx, 10))

	return nil
}

// TODO use reflection to extract field names
func (r *UtxoRepo) scanDataInto(data map[string]string, w *UtxoEntity) error {
	field := "IsSpent"
	val, ok := data[field] // hex string: 0x(66 chars here)
	if ok {
		v, err := strconv.ParseBool(val)
		if err != nil {
			return err
		}
		w.IsSpent = v
	}

	field = "TxId"
	val, ok = data[field] // hex string: 0x(66 chars here)
	if ok {
		w.TxId = val
	}

	field = "TxOutputIndex"
	val, ok = data[field] // hex string: 0x(66 chars here)
	if ok {
		w.TxOutputIndex = val
	}

	field = "Owner"
	val, ok = data[field] // hex string: 0x(64 chars here)
	if ok {
		w.Owner = val
	}

	field = "AssetId"
	val, ok = data[field] // hex string: 0x(64 chars here)
	if ok {
		w.AssetId = val
	}

	field = "Amount"
	val, ok = data[field]
	if ok {
		v, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return err
		}
		w.Amount = v
	}

	field = "BlockCreated"
	val, ok = data[field]
	if ok {
		v, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return err
		}
		w.BlockCreated = v
	}

	field = "TxCreatedIdx"
	val, ok = data[field]
	if ok {
		v, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return err
		}
		w.TxCreatedIdx = v
	}

	return nil
}

func (r *UtxoRepo) Delete(ctx context.Context, w *UtxoEntity) error {
	key := r.GenerateKey(w)
	if cmd := r.redisClient.Del(ctx, key); cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func (r *UtxoRepo) DeleteByKey(ctx context.Context, key string) error {
	if cmd := r.redisClient.Del(ctx, key); cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func (r *UtxoRepo) GenerateKeyUsingParams(ownerId, TxId, TxOutputIdx string) string {
	return repo.FormUtxoHashmapKeyTemplate(ownerId, TxId, TxOutputIdx)
}

func (r *UtxoRepo) GenerateKey(w *UtxoEntity) string {
	return repo.FormUtxoHashmapKeyTemplate(w.Owner, w.TxId, w.TxOutputIndex)
}
