package cache

import (
	"github.com/go-redis/redis"
	"log"
	"github.com/headend/share-module/configuration"
	"strconv"
	"time"
)

type Cache struct {
	Cli *redis.Client
	Err error
}

func (r *Cache) Connect(config *configuration.Conf) *Cache {
	config.LoadConf()
	r.Cli = redis.NewClient(&redis.Options{
		Addr: config.Redis.Host +":"+ strconv.Itoa(config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.DB, 
		DialTimeout:  time.Duration(config.Redis.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.Redis.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Redis.WriteTimeout) * time.Second,
		PoolSize:     config.Redis.PoolSize,
		PoolTimeout:  time.Duration(config.Redis.PoolTimeout) * time.Second,
	})
	_, err := r.Cli.Ping().Result()
	if err != nil{
		r.Err = err
	}
	return r
}

func (r *Cache) CloseConnect() error {
	err := r.Cli.Close()
	return err
}

func (r *Cache) SetMessage(key string, msg string) *Cache{
	r.Err = r.Cli.Set(key, msg, 0).Err()
	if r.Err != nil {
		log.Print(r.Err)
	}
	return r
}

func (r *Cache) GetMessage(key string) (string, error){
	res, err := r.Cli.Get(key).Result()
	if err != nil {
		log.Print(err)
		return res,err
	}
	return res,err
}

func (r *Cache) GetAllKeys() ([]string, error){
	return r.FilterKeys("*")
}

func (r *Cache) FilterKeys(pattern string) ([]string, error){
	res, err := r.Cli.Keys(pattern).Result()
	if err != nil {
		log.Print(err)
		return res, err
	}
	return res,err
}

func (r *Cache) DeleteKey(pattern string) (error){
	err := r.Cli.Del(pattern).Err()
	if err != nil {
		log.Print(err)
		return err
	} else {
		return nil
	}
	return err
}
