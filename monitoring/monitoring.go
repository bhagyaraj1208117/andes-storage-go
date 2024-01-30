package monitoring

import (
	"github.com/bhagyaraj1208117/andes-core-go/core"
	"github.com/bhagyaraj1208117/andes-core-go/core/atomic"
	logger "github.com/bhagyaraj1208117/andes-logger-go"
)

var log = logger.GetOrCreate("storage")

var cumulatedSizeInBytes atomic.Counter

// MonitorNewCache adds the size in the global cumulated size variable
func MonitorNewCache(tag string, sizeInBytes uint64) {
	cumulatedSizeInBytes.Add(int64(sizeInBytes))
	log.Debug("MonitorNewCache", "name", tag, "capacity", core.ConvertBytes(sizeInBytes), "cumulated", core.ConvertBytes(cumulatedSizeInBytes.GetUint64()))
}
