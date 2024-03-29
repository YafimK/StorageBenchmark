package main

import (
	"fmt"
	"github.com/ReneKroon/ttlcache"
	"github.com/bluele/gcache"
	"github.com/coocood/freecache"
	gocache "github.com/patrickmn/go-cache"
	"log"
	"math/rand"
	"testing"
	"time"
)

const maxEntrySize = 512
const ttl = timeUnit * time.Second
const timeUnit =30
const defaultGoCacheCleanWindow = 30 * time.Second
const defaultTTl = 168 * time.Second

func BenchmarkFreeCacheSet(b *testing.B) {
	cache := freecache.NewCache(b.N * maxEntrySize)
	for i := 0; i < b.N; i++ {
		if err := cache.Set([]byte(key(i)), value(), timeUnit); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkGoCacheSet(b *testing.B) {
	cache := gocache.New(defaultTTl, defaultGoCacheCleanWindow)
	for i := 0; i < b.N; i++ {
		cache.Set(key(i), value(), ttl)
	}
}

func BenchmarkTTLCacheSet(b *testing.B) {
	cache := ttlcache.NewCache()

	for i := 0; i < b.N; i++ {
		cache.SetWithTTL(key(i), value(), ttl)
	}
}

func BenchmarkGCacheSet(b *testing.B) {
	gc := gcache.New(b.N).ARC().Build()
	for i := 0; i < b.N; i++ {
		if err := gc.SetWithExpire(key(i), value(), ttl); err != nil {
			log.Fatal(err)
		}
	}
}

//func BenchmarkBadgerSet(b *testing.B) {
//	db := startBadger()
//	defer db.Close()
//	for i := 0; i < b.N; i++ {
//		if err :=db.Update(func(txn *badger.Txn) error {
//			e := badger.NewEntry([]byte(key(i)), value()).WithTTL(ttl)
//			err := txn.SetEntry(e)
//			return err
//		}); err != nil {
//			log.Fatal(err)
//		}
//	}
//}

func BenchmarkFreeCacheGet(b *testing.B) {
	b.StopTimer()
	cache := freecache.NewCache(b.N * maxEntrySize)
	for i := 0; i < b.N; i++ {
		if err := cache.Set([]byte(key(i)), value(), timeUnit); err != nil {
			log.Fatal(err)
		}
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if _, err := cache.Get([]byte(key(i))); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkGoCacheGet(b *testing.B) {
	b.StopTimer()
	cache := gocache.New(defaultTTl, defaultGoCacheCleanWindow)
	for i := 0; i < b.N; i++ {
		cache.Set(key(i), value(), ttl)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(key(i))
	}
}

func BenchmarkTTLCacheGet(b *testing.B) {
	b.StopTimer()
	cache := ttlcache.NewCache()

	for i := 0; i < b.N; i++ {
		cache.SetWithTTL(key(i), value(), ttl)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(key(i))
	}
}

func BenchmarkGCacheGet(b *testing.B) {
	b.StopTimer()
	gc := gcache.New(b.N).ARC().Build()
	for i := 0; i < b.N; i++ {
		if err := gc.SetWithExpire(key(i), value(), ttl); err != nil {
			log.Fatal(err)
		}
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if _, err := gc.Get(key(i)); err != nil {
			log.Fatal(err)
		}
	}
}

//func BenchmarkBadgerGet(b *testing.B) {
//	b.StopTimer()
//	db := startBadger()
//	defer db.Close()
//	for i := 0; i < b.N; i++ {
//		if err :=db.Update(func(txn *badger.Txn) error {
//			e := badger.NewEntry([]byte(key(i)), value()).WithTTL(ttl)
//			err := txn.SetEntry(e)
//			return err
//		}); err != nil {
//			log.Fatal(err)
//		}
//
//	}
//
//	b.StartTimer()
//	for i := 0; i < b.N; i++ {
//		if err := db.View(func(txn *badger.Txn) error {
//			_, err := txn.Get([]byte(key(i)))
//			return err
//		}); err != nil {
//			log.Fatal(err)
//		}
//	}
//}

func BenchmarkFreeCacheSetParallel(b *testing.B) {
	cache := freecache.NewCache(b.N * maxEntrySize)
	rand.Seed(time.Now().Unix())

	b.RunParallel(func(pb *testing.PB) {
		id := rand.Intn(1000)
		counter := 0
		for pb.Next() {
			if err := cache.Set([]byte(parallelKey(id, counter)), value(), timeUnit); err != nil {
				log.Fatal(err)
			}
			counter += 1
		}
	})
}

func BenchmarkGoCacheSetParallel(b *testing.B) {
	cache := gocache.New(defaultTTl, defaultGoCacheCleanWindow)
	rand.Seed(time.Now().Unix())

	b.RunParallel(func(pb *testing.PB) {
		id := rand.Intn(1000)
		counter := 0
		for pb.Next() {
			cache.Set(parallelKey(id, counter), value(), ttl)
			counter += 1
		}
	})
}

func BenchmarkTTLCacheSetParallel(b *testing.B) {
	cache := ttlcache.NewCache()

	rand.Seed(time.Now().Unix())

	b.RunParallel(func(pb *testing.PB) {
		id := rand.Intn(1000)
		counter := 0
		for pb.Next() {
			cache.SetWithTTL(parallelKey(id, counter), value(), ttl)
			counter += 1
		}
	})
}

func BenchmarkGCacheSetParallel(b *testing.B) {
	gc := gcache.New(b.N).ARC().Build()

	rand.Seed(time.Now().Unix())

	b.RunParallel(func(pb *testing.PB) {
		id := rand.Intn(1000)
		counter := 0
		for pb.Next() {
			if err := gc.SetWithExpire(parallelKey(id, counter), value(), ttl); err != nil {
				log.Fatal(err)
			}
			counter += 1
		}
	})
}

//func BenchmarkBadgerSetParallel(b *testing.B) {
//	db := startBadger()
//	defer db.Close()
//	rand.Seed(time.Now().Unix())
//
//	b.RunParallel(func(pb *testing.PB) {
//		id := rand.Intn(1000)
//		counter := 0
//		for pb.Next() {
//			if err :=db.Update(func(txn *badger.Txn) error {
//				e := badger.NewEntry([]byte(parallelKey(id, counter)), value()).WithTTL(ttl)
//				err := txn.SetEntry(e)
//				return err
//			}); err != nil {
//				log.Fatal(err)
//			}
//
//			counter += 1
//		}
//	})
//}

func BenchmarkFreeCacheGetParallel(b *testing.B) {
	b.StopTimer()
	cache := freecache.NewCache(b.N * maxEntrySize)
	//wg := sync.WaitGroup{}
	//wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		//go func(i int) {
		//	defer func() { wg.Done() }()
			if err := cache.Set([]byte(key(i)), value(), timeUnit); err != nil {
				log.Fatal(err)
			}
		//}(i)
	}
	//wg.Wait()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		counter := 0
		for pb.Next() {
			if _, err := cache.Get([]byte(key(counter))); err != nil {
				log.Fatal(err)
			}
			counter += 1
		}
	})
}

func BenchmarkGoCacheGetParallel(b *testing.B) {
	b.StopTimer()
	cache := gocache.New(defaultTTl, defaultGoCacheCleanWindow)
	for i := 0; i < b.N; i++ {
		cache.Set(key(i), value(), ttl)
	}

	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		counter := 0
		for pb.Next() {
			cache.Get(key(counter))
			counter += 1
		}
	})
}

func BenchmarkTTLCacheGetParallel(b *testing.B) {
	b.StopTimer()
	cache := ttlcache.NewCache()

	for i := 0; i < b.N; i++ {
		cache.SetWithTTL(key(i), value(), ttl)
	}

	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		counter := 0
		for pb.Next() {
			cache.Get(key(counter))
			counter += 1
		}
	})
}

func BenchmarkGCacheGetParallel(b *testing.B) {
	b.StopTimer()
	gc := gcache.New(b.N).ARC().Build()

	for i := 0; i < b.N; i++ {
		if err := gc.SetWithExpire(key(i), value(), ttl); err != nil {
			log.Fatal(err)
		}
	}

	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		counter := 0
		for pb.Next() {
			if _, err := gc.Get(key(counter)); err != nil {
				log.Fatal(err)
			}
			counter += 1
		}
	})
}

//func BenchmarkBadgerGetParallel(b *testing.B) {
//	b.StopTimer()
//	db := startBadger()
//	defer db.Close()
//
//	for i := 0; i < b.N; i++ {
//		 if err :=db.Update(func(txn *badger.Txn) error {
//			e := badger.NewEntry([]byte(key(i)), value()).WithTTL(ttl)
//			err := txn.SetEntry(e)
//			return err
//		}); err != nil {
//			log.Fatal(err)
//		 }
//	}
//
//	b.StartTimer()
//	b.RunParallel(func(pb *testing.PB) {
//		counter := 0
//		for pb.Next() {
//			if err := db.View(func(txn *badger.Txn) error {
//				_, err := txn.Get([]byte(key(counter)))
//				return err
//			}); err != nil {
//				log.Fatal(err)
//			}
//			counter += 1
//		}
//	})
//}

func key(i int) string {
	return fmt.Sprintf("key-%010d", i)
}

func value() []byte {
	return make([]byte, 100)
}

func parallelKey(threadID int, counter int) string {
	return fmt.Sprintf("key-%04d-%06d", threadID, counter)
}

//func startBadger() *badger.DB {
//	options := badger.DefaultOptions("/tmp/badger").WithLogger(nil)
//	db, err := badger.Open(options)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return db
//}
