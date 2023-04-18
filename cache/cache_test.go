// Package cache implements a RedisCache struct that stores data in a Redis cache.
package cache

import "testing"

// TestRedisCache_Has() tests if the Redis cache doesn't have an entry for 'foo'.
// If an entry exists, it should be deleted and checked again.
func TestRedisCache_Has(t *testing.T) {
	err := testRedisCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo found in cache, and it shouldn't be there")
	}

	err = testRedisCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	inCache, err = testRedisCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("foo not found in cache, but it should be there")
	}
}

// TestRedisCache_Get() checks if the function returns the correct value for a given key.
func TestRedisCache_Get(t *testing.T) {
	//Add an entry for 'foo' in cache.
	err := testRedisCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	//Get the value for 'foo'.
	x, err := testRedisCache.Get("foo")
	if err != nil {
		t.Error(err)
	}

	//Check if the value obtained from Get() for 'foo' is correct.
	if x != "bar" {
		t.Error("did not get correct value from cache")
	}
}

// TestRedisCache_Forget() checks if an entry can be deleted from cache.
func TestRedisCache_Forget(t *testing.T) {
	//Add an entry for 'alpha' in cache.
	err := testRedisCache.Set("alpha", "beta")
	if err != nil {
		t.Error(err)
	}

	//Delete the entry for 'alpha'.
	err = testRedisCache.Forget("alpha")
	if err != nil {
		t.Error(err)
	}

	//Check if 'alpha' exists in cache after deleting it.
	inCache, err := testRedisCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	//If 'alpha' exists in cache, it should not be there.
	if inCache {
		t.Error("alpha found in cache, and it should not be there")
	}
}

// TestRedisCache_Empty() checks if all entries in cache can be deleted.
func TestRedisCache_Empty(t *testing.T) {
	//Add an entry for 'alpha' in cache.
	err := testRedisCache.Set("alpha", "beta")
	if err != nil {
		t.Error(err)
	}

	//Empty the cache.
	err = testRedisCache.Empty()
	if err != nil {
		t.Error(err)
	}

	//Check if 'alpha' exists in cache after emptying it.
	inCache, err := testRedisCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	//If 'alpha' exists in cache, it should not be there.
	if inCache {
		t.Error("alpha found in cache, and it should not be there")
	}

}

// TestRedisCache_EmptyByMatch() checks if entries that match a given pattern can be deleted from the cache.
func TestRedisCache_EmptyByMatch(t *testing.T) {
	err := testRedisCache.Set("alpha", "foo")
	if err != nil {
		t.Error(err)
	}

	err = testRedisCache.Set("alpha2", "foo")
	if err != nil {
		t.Error(err)
	}

	err = testRedisCache.Set("beta", "foo")
	if err != nil {
		t.Error(err)
	}

	err = testRedisCache.EmptyByMatch("alpha")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("alpha found in cache, and it should not be there")
	}

	inCache, err = testRedisCache.Has("alpha2")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("alpha2 found in cache, and it should not be there")
	}

	inCache, err = testRedisCache.Has("beta")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("beta not found in cache, and it should be there")
	}
}

func TestEncodeDecode(t *testing.T) {
	entry := Entry{}
	entry["foo"] = "bar"
	bytes, err := encode(entry)
	if err != nil {
		t.Error(err)
	}

	_, err = decode(string(bytes))
	if err != nil {
		t.Error(err)
	}

}
