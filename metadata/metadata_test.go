package metadata

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/IPFS-eX/go-graphsync/testbridge"

	"github.com/ipld/go-ipld-prime/linking/cid"

	"github.com/IPFS-eX/go-graphsync/testutil"
)

func TestDecodeEncodeMetadata(t *testing.T) {
	cids := testutil.GenerateCids(10)
	initialMetadata := make(Metadata, 0, 10)
	for _, k := range cids {
		link := cidlink.Link{Cid: k}
		blockPresent := rand.Int31()%2 == 0
		initialMetadata = append(initialMetadata, Item{link, blockPresent})
	}
	bridge := testbridge.NewMockIPLDBridge()
	encoded, err := EncodeMetadata(initialMetadata, bridge)
	if err != nil {
		t.Fatal("Error encoding")
	}
	decodedMetadata, err := DecodeMetadata(encoded, bridge)
	if err != nil {
		t.Fatal("Error decoding")
	}
	if !reflect.DeepEqual(initialMetadata, decodedMetadata) {
		t.Fatal("Metadata changed during encoding and decoding")
	}
}
