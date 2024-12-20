// Code generated by protoc-gen-store. DO NOT EDIT.
// version: v0.0.1
// source: valid/message.proto

package pb

import (
	bolt "go.etcd.io/bbolt"
	"google.golang.org/protobuf/proto"
)

const (
	AssetKey string = "test.package.Asset"
	TestKey  string = "test.package.Test"
	PopeKey  string = "test.package.Pope"
)

type StoreSoul []byte
type Store struct {
	db *bolt.DB
}

func (store *Store) SetAsset(soul StoreSoul, data *Asset) error {
	return store.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(AssetKey))
		if err != nil {
			return err
		}

		dataBytes, err := proto.Marshal(data)
		if err != nil {
			return err
		}

		soulBucket, err := bucket.CreateBucketIfNotExists(soul)
		if err != nil {
			return err
		}

		return soulBucket.Put([]byte(data.GetId()), dataBytes)
	})
}

func (store *Store) GetAsset(soul StoreSoul) (result []*Asset, err error) {
	err = store.db.View(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(AssetKey))
		if err != nil {
			return err
		}

		soulBucket, err := bucket.CreateBucketIfNotExists(soul)
		if err != nil {
			return err
		}

		return soulBucket.ForEach(func(_, data []byte) error {
			value := &Asset{}

			err := proto.Unmarshal(data, value)
			if err != nil {
				return err
			}

			result = append(result, value)
			return nil
		})
	})

	return result, err
}

func (store *Store) GetAssetById(soul StoreSoul, id string) (result *Asset, err error) {
	err = store.db.View(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(AssetKey))
		if err != nil {
			return err
		}

		soulBucket, err := bucket.CreateBucketIfNotExists(soul)
		if err != nil {
			return err
		}

		data := soulBucket.Get([]byte(id))
		return proto.Unmarshal(data, result)
	})

	return result, err
}

func (store *Store) SetTest(soul StoreSoul, data *Test) error {
	return store.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(TestKey))
		if err != nil {
			return err
		}

		dataBytes, err := proto.Marshal(data)
		if err != nil {
			return err
		}

		return bucket.Put(soul, dataBytes)
	})
}

func (store *Store) GetTest(soul StoreSoul) (result *Test, err error) {
	err = store.db.View(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(TestKey))
		if err != nil {
			return err
		}

		data := bucket.Get(soul)
		return proto.Unmarshal(data, result)
	})

	return result, err
}

func (store *Store) SetPope(soul StoreSoul, data *Pope) error {
	return store.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(PopeKey))
		if err != nil {
			return err
		}

		dataBytes, err := proto.Marshal(data)
		if err != nil {
			return err
		}

		return bucket.Put(soul, dataBytes)
	})
}

func (store *Store) GetPope(soul StoreSoul) (result *Pope, err error) {
	err = store.db.View(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(PopeKey))
		if err != nil {
			return err
		}

		data := bucket.Get(soul)
		return proto.Unmarshal(data, result)
	})

	return result, err
}
