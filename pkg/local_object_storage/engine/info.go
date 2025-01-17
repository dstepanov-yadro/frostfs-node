package engine

import (
	"github.com/TrueCloudLab/frostfs-node/pkg/local_object_storage/shard"
)

// Info groups the information about StorageEngine.
type Info struct {
	Shards []shard.Info
}

// DumpInfo returns information about the StorageEngine.
func (e *StorageEngine) DumpInfo() (i Info) {
	e.mtx.RLock()
	defer e.mtx.RUnlock()

	i.Shards = make([]shard.Info, 0, len(e.shards))

	for _, sh := range e.shards {
		info := sh.DumpInfo()
		info.ErrorCount = sh.errorCount.Load()
		i.Shards = append(i.Shards, info)
	}

	return
}
