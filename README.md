# redis-orm

redis-orm fly orm up 

## features

## quick start

### generate command

````
$: go get github.com/ezbuy/redis-orm

$: redis-orm code -i example/yaml -o example/model

````

### read access usage

````
import "github.com/ezbuy/redis-orm/example/model"

# mysql
model.MySQLSetup(cf)


db := model.MySQL()
//! query (ids []string) by unique & index & range definitions
model.UserDBMgr(db).FindOne(unique)
model.UserDBMgr(db).Find(index)
model.UserDBMgr(db).Range(scope)
model.UserDBMgr(db).RangeRevert(scope)

//! fetch object 
model.UserDBMgr(db).Fetch(id string) (*User, error)
model.UserDBMgr(db).FetchByPrimaryKeys(pks []PrimaryKey) ([]*User, error)

# redis
model.RedisSetup(cf)

//! query (ids []string) by unique & index & range definitions
model.UserRedisMgr().FindOne(unique)
model.UserRedisMgr().Find(index)
model.UserRedisMgr().Range(scope)
model.UserRedisMgr().RangeRevert(scope)

//! fetch object 
model.UserRedisMgr().Fetch(id string) (*User, error)
model.UserRedisMgr().FetchByPrimaryKeys(pks []PrimaryKey) ([]*User, error)


````

### write access usage

````
import "github.com/ezbuy/redis-orm/example/model"

# mysql
model.MySQLSetup(cf)

db := model.MySQL()
tx, err := db.BeginTx()
defer tx.Close()

model.UserDBMgr(tx).Save(obj)
model.UserDBMgr(tx).Create(obj)
model.UserDBMgr(tx).Update(obj)
model.UserDBMgr(tx).Delete(obj)

model.UserDBMgr(tx).FindOne(unique)
model.UserDBMgr(tx).Find(index)
model.UserDBMgr(tx).Range(scope)
model.UserDBMgr(tx).RangeRevert(scope)

model.UserDBMgr(tx).Fetch(id string) (*User, error)
model.UserDBMgr(tx).FetchByPrimaryKeys(pks []PrimaryKey) ([]*User, error)

# redis
model.RedisSetup(cf)

model.UserRedisMgr().Save(obj)
model.UserRedisMgr().Create(obj)
model.UserRedisMgr().Update(obj)
model.UserRedisMgr().Delete(obj)

````

### sync data

````
import "github.com/ezbuy/redis-orm/example/model"

model.MySQLSetup(cf)
model.RedisSetup(cf)

model.UserRedisMgr().Load(model.UserMySQLMgr())

````

## bench redis vs mysql

enviroment:
  
  mysql-server, redis-server, test client all in the same machine (mac air)

*redis-orm.redis.bench*
  
    Ran 1000 samples:
    unique.runtime:
      Fastest Time: 0.000s
      Slowest Time: 0.001s
      Average Time: 0.000s ± 0.000s
    index.runtime:
      Fastest Time: 0.000s
      Slowest Time: 0.000s
      Average Time: 0.000s ± 0.000s
    range.runtime:
      Fastest Time: 0.000s
      Slowest Time: 0.000s
      Average Time: 0.000s ± 0.000s
    range.revert.runtime:
      Fastest Time: 0.000s
      Slowest Time: 0.000s
      Average Time: 0.000s ± 0.000s
    fetch.runtime:
      Fastest Time: 0.002s
      Slowest Time: 0.004s
      Average Time: 0.002s ± 0.000s

*redis-orm.mysql.bench*
  
    Ran 1000 samples:
    unique.runtime:
      Fastest Time: 0.002s
      Slowest Time: 0.106s
      Average Time: 0.003s ± 0.005s
    index.runtime:
      Fastest Time: 0.002s
      Slowest Time: 0.106s
      Average Time: 0.003s ± 0.005s
    range.runtime:
      Fastest Time: 0.002s
      Slowest Time: 0.105s
      Average Time: 0.002s ± 0.005s
    range.revert.runtime:
      Fastest Time: 0.002s
      Slowest Time: 0.105s
      Average Time: 0.002s ± 0.006s
    fetch.runtime:
      Fastest Time: 0.004s
      Slowest Time: 0.150s
      Average Time: 0.006s ± 0.009s

