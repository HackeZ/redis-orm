package model

import (
	"fmt"
	redis "gopkg.in/redis.v5"

	"github.com/ezbuy/redis-orm/orm"
)

var (
	_ fmt.Formatter
	_ orm.VSet
)

//! relation
type UserLocation struct {
	Key       string  `db:"key" json:"key"`
	Longitude float64 `db:"longitude" json:"longitude"`
	Latitude  float64 `db:"latitude" json:"latitude"`
	Value     int32   `db:"value" json:"value"`
}

func (relation *UserLocation) GetClassName() string {
	return "UserLocation"
}

func (relation *UserLocation) GetIndexes() []string {
	idx := []string{}
	return idx
}

func (relation *UserLocation) GetStoreType() string {
	return "geo"
}

func (relation *UserLocation) GetPrimaryName() string {
	return "Key"
}

type _UserLocationRedisMgr struct {
	*orm.RedisStore
}

func UserLocationRedisMgr(stores ...*orm.RedisStore) *_UserLocationRedisMgr {
	if len(stores) > 0 {
		return &_UserLocationRedisMgr{stores[0]}
	}
	return &_UserLocationRedisMgr{_redis_store}
}

//! pipeline write
type _UserLocationRedisPipeline struct {
	*redis.Pipeline
	Err error
}

func (m *_UserLocationRedisMgr) BeginPipeline() *_UserLocationRedisPipeline {
	return &_UserLocationRedisPipeline{m.Pipeline(), nil}
}

func (m *_UserLocationRedisMgr) NewUserLocation(key string) *UserLocation {
	return &UserLocation{
		Key: key,
	}
}

//! redis relation pair
func (m *_UserLocationRedisMgr) LocationAdd(obj *UserLocation) error {
	return m.GeoAdd(geoOfClass(obj.GetClassName(), obj.Key), &redis.GeoLocation{
		Longitude: obj.Longitude,
		Latitude:  obj.Latitude,
		Name:      fmt.Sprint(obj.Value),
	}).Err()
}

func (m *_UserLocationRedisMgr) LocationRadius(key string, longitude float64, latitude float64, query *redis.GeoRadiusQuery) ([]*UserLocation, error) {
	locations, err := m.GeoRadius(geoOfClass("UserLocation", key), longitude, latitude, query).Result()
	if err != nil {
		return nil, err
	}

	objs := make([]*UserLocation, len(locations))
	for _, location := range locations {
		obj := m.NewUserLocation(key)
		obj.Longitude = location.Longitude
		obj.Latitude = location.Latitude
		if err := m.StringScan(location.Name, &obj.Value); err != nil {
			return nil, err
		}
		objs = append(objs, obj)
	}
	return objs, nil
}

func (m *_UserLocationRedisMgr) LocationRem(obj *UserLocation) error {
	return m.ZRem(geoOfClass(obj.GetClassName(), obj.Key), fmt.Sprint(obj.Value)).Err()
}
