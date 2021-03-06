package tangle

import (
	"github.com/gohornet/hornet/packages/datastructure"
	"github.com/gohornet/hornet/packages/profile"
)

var (
	BundleBucketCache *datastructure.LRUCache
)

func InitBundleCache() {
	opts := profile.GetProfile().Caches.Bundles
	BundleBucketCache = datastructure.NewLRUCache(opts.Size, &datastructure.LRUCacheOptions{
		EvictionCallback:  onEvictBundles,
		EvictionBatchSize: opts.EvictionSize,
	})
}

func onEvictBundles(_ interface{}, values interface{}) {
	valT := values.([]interface{})

	var bundles []*BundleBucket
	for _, obj := range valT {
		bundles = append(bundles, obj.(*BundleBucket))
	}

	err := StoreBundleBucketsInDatabase(bundles)
	if err != nil {
		panic(err)
	}
}

func FlushBundleCache() {
	BundleBucketCache.DeleteAll()
}
