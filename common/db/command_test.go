package db

import (
	"fmt"
	"github.com/mongodb/mongo-tools/common/options"
	"testing"
)

func sessionProvider(t *testing.T, uri string) *SessionProvider {
	enabled := options.EnabledOptions{
		Auth:       true,
		Connection: true,
		Namespace:  true,
		URI:        true,
	}
	toolOptions := options.New("test", "", "", "", true, enabled)
	_, err := toolOptions.ParseArgs([]string{"--uri", uri})
	if err != nil {
		t.Fatal(err)
	}
	sp, err := NewSessionProvider(*toolOptions)
	if err != nil {
		t.Fatal(err)
	}
	return sp
}

func TestSessionProvider_GetNodeType(t *testing.T) {
	sp := sessionProvider(t, "mongodb://127.0.0.1:27017,127.0.0.1:27018/foo?readPreference=primaryPreferred")
	nType, err := sp.GetNodeType()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(nType)
}
