package server

import (
	"fmt"
	"testing"
)

const (
	expectedOptionsJSONResult = `{"name":"App-name","hostname":"localhost","port":1001,` +
		`"profPort":1002,"maxConnections":9999,"maxProcs":9998,"debugEnabled":true,` +
		`"namespace":"qa","alivePath":"alivepath","notifyPath":"notifypath","statusPath":"statuspath",` +
		`"targetHost":"myjenkins.test.com","targetPort":9997,"targetPath":"/testtargetpath/"}`
)

func TestOptionsString(t *testing.T) {
	t.Parallel()

	opts := &Options{
		Name:        "App-name",
		Hostname:    "localhost",
		Port:        1001,
		ProfPort:    1002,
		MaxConn:     9999,
		MaxProcs:    9998,
		Debug:       true,
		ValidTokens: []string{"token1", "token2", "token3"},
		Namespace:   "qa",
		AlivePath:   "alivepath",
		NotifyPath:  "notifypath",
		StatusPath:  "statuspath",
		TargetHost:  "myjenkins.test.com",
		TargetPort:  9997,
		TargetPath:  "/testtargetpath/",
		TargetToken: "atokentester",
	}
	actual := fmt.Sprint(opts)
	if actual != expectedOptionsJSONResult {
		t.Errorf("Options not converted to json string.\n\nExpected: %s\n\nActual: %s\n",
			expectedOptionsJSONResult, actual)
	}
}
