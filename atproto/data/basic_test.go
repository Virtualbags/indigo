package data

import (
	"encoding/json"
	"testing"

	"github.com/bluesky-social/indigo/atproto/syntax"

	"github.com/stretchr/testify/assert"
)

func TestSimpleValidation(t *testing.T) {
	assert := assert.New(t)

	s := "a string"
	assert.NoError(Validate(map[string]interface{}{
		"a": 5,
		"b": 123,
		"c": s,
		"d": &s,
	}))
	assert.NoError(Validate(map[string]interface{}{
		"$type": "com.example.thing",
		"a":     5,
	}))
	assert.Error(Validate(map[string]interface{}{
		"$type": 123,
		"a":     5,
	}))
	assert.Error(Validate(map[string]interface{}{
		"$type": "",
		"a":     5,
	}))
}

func TestSyntaxSerialize(t *testing.T) {
	assert := assert.New(t)

	atid, err := syntax.ParseAtIdentifier("did:web:example.com")
	assert.NoError(err)
	obj := map[string]interface{}{
		"at-identifier": atid,
		"at-uri":        syntax.ATURI("at://did:abc:123/io.nsid.someFunc/record-key"),
		"cid-string":    syntax.CID("bafyreidfayvfuwqa7qlnopdjiqrxzs6blmoeu4rujcjtnci5beludirz2a"),
		"datetime":      syntax.Datetime("2023-10-30T22:25:23Z"),
		"did":           syntax.DID("did:web:example.com"),
		"handle":        syntax.Handle("blah.example.com"),
		"language":      syntax.Language("us"),
		"nsid":          syntax.NSID("com.example.blah"),
		"recordkey":     syntax.RecordKey("self"),
		"tid":           syntax.TID("3kao2cl6lyj2p"),
		"uri":           syntax.URI("https://example.com/file"),
	}
	assert.NoError(Validate(obj))
	_, err = MarshalCBOR(obj)
	assert.NoError(err)
	_, err = json.Marshal(obj)
	assert.NoError(err)
}
