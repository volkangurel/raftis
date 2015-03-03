package ops

import (
	mdb "github.com/jbooth/gomdb"
	redis "github.com/jbooth/raftis/redis"
	utils "github.com/jbooth/raftis/utils"
	"io"
)

// 1 arg, key
func GET(args [][]byte, txn *mdb.Txn, w io.Writer) (int64, error) {
	key := args[0]
	println("GET " + string(key))
	val, err := utils.GetString(txn, key)
	if err == mdb.NotFound {
		// Not found is nil in redis
		return redis.NilReply.WriteTo(w)
	} else if err != nil {
		// write error
		return redis.NewError(err.Error()).WriteTo(w)
	}
	resp := &redis.BulkReply{[]byte(val)}
	return resp.WriteTo(w)
}

// 1 arg, key
func STRLEN(args [][]byte, txn *mdb.Txn, w io.Writer) (int64, error) {
	key := args[0]
	println("STRLEN " + string(key))
	val, err := utils.GetString(txn, key)
	if err == mdb.NotFound {
		resp := &redis.IntegerReply{0}
		return resp.WriteTo(w)
	} else if err != nil {
		return redis.NewError(err.Error()).WriteTo(w)
	}
	resp := &redis.IntegerReply{len(val)}
	return resp.WriteTo(w)
}
