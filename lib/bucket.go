package lib

import (
	"github.com/spaolacci/murmur3"
)

type Bucket struct {
	experimentID   string
	projectKey     string
	variantVersion string
	IDRange        map[string]uint64
}

// Spilt function
func GetVersion(buckets []Bucket, userID string) string {
	var variantVersion string

	bucketID := getBucketID(userID)

	for _, bucket := range buckets {
		if (bucket.IDRange["min"] <= bucketID) && (bucketID <= bucket.IDRange["max"]) {
			variantVersion = bucket.variantVersion
		}
	}
	return variantVersion
}

// C
func CreateBucket(experimentID string, projectKey string, config map[string]uint64) []Bucket {
	var newBuckets []Bucket
	startID := uint64(0)

	for key, value := range config {
		IDRange := make(map[string]uint64)
		IDRange["min"] = startID
		IDRange["max"] = startID + value - 1
		newBuckets = append(newBuckets, Bucket{
			experimentID,
			projectKey,
			key,
			IDRange,
		})

		startID = IDRange["max"] + 1
	}

	return newBuckets
}

func getBucketID(customerID string) uint64 {

	customerIDint := murmur3.Sum64([]byte(customerID))

	bucketID := customerIDint % 100
	return bucketID
}
