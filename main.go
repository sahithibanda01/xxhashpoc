package main

import (
	"fmt"

	//  "github.com/buraksezer/consistent"
	"github.com/cespare/xxhash"
)

// // In your distributed system, you probably have a custom data type
// // for your cluster members. Just add a String function to implement
// // consistent.Member interface.
// type Member string

// func (m Member) String() string {
// 	return string(m)
// }

// type hasher struct{}

// func (h hasher) Sum64(data []byte) uint64 {
// 	return xxhash.Sum64(data)
// }

// func main() {
// 	// Create a new consistent instance
// 	members := []consistent.Member{}
// 	for i := 0; i < 20; i++ {
// 		member := Member(fmt.Sprintf("Bucket_%d", i))
// 		members = append(members, member)
// 	}
// 	cfg := consistent.Config{
// 		PartitionCount:    100000,
// 		ReplicationFactor: 1,
// 		Load:              1.05,
// 		Hasher:            hasher{},
// 	}

// 	c := consistent.New(members, cfg)
// 	keyCounts := make(map[string]int)

// 	for i:=1; i<= 100000000; i++ {
// 		key := []byte(fmt.Sprintf("key%d", i))
// 		owner := c.LocateKey(key)
// 		member := owner.String()

// 		// Increment the count for this member
// 		keyCounts[member]++
// 	}

// 	// Print the counts for each member
// 	for member, count := range keyCounts {
// 		fmt.Printf("Member %s: %d keys\n", member, count)

// 	}
	
// 	// calculates partition id for the given key
// 	// partID := hash(key) % partitionCount
// 	// the partitions is already distributed among members by Add function.
	
// 	// Prints node2.olricmq.com
// }

// //random  seedvalue 
// // always key to bucket should be same 5k timnes
// // 1lkh keys 20 bukets uniform almost



// //store it in file run it again and again store it in file and compare


// //

// package main
// import (
//     "fmt"
//     "github.com/cespare/xxhash/v2"
// )

func main() {
    key := []byte("example-key")

    // Create a new hasher using the xxhash library
    hasher := xxhash.New()
    hasher.Write(key)
    hkey := hasher.Sum64()

    fmt.Printf("Hash value for key '%s': %d\n", key, hkey)
}

//create 2 vms in gcp with diff os and diff config {cpu/memory}
