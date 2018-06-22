package lib

import (
	"bytes"
	"encoding/binary"

	"github.com/zond/god/murmur"
)

// Spilt function
func Bucket(experimentID string, projectKey string, userID string) int64 {
	var customerIDBytes bytes.Buffer
	customerIDBytes.WriteString(experimentID)
	customerIDBytes.WriteString(projectKey)
	customerIDBytes.WriteString(userID)

	bucketID := getBucketID(customerIDBytes.String())

	return bucketID
}

func getBucketID(customerID string) int64 {
	var customerIDInt int64

	customerIDHash := murmur.HashString(customerID)
	binary.Read(bytes.NewReader(customerIDHash), binary.BigEndian, &customerIDInt)

	bucketID := customerIDInt & 0xFFFFFFFF % 100
	return bucketID
}
