package server

import (
	"fmt"
	"reflect"
	"testing"
)

const (
	expectedInfoJSONResult = `{"version":"1.2.3","name":"Test Server","hostname":"localhost",` +
		`"UUID":"ABCDEFGHIJKLMNOPQRSTUVWXYZ","port":1001,"profPort":1002,"maxConnections":9999,` +
		`"debugEnabled":false,"namespace":"dev23","alivePath":"/pathalive/","notifyPath":"/pathnotify/",` +
		`"statusPath":"/pathstatus/","targetHost":"myjenkins.test.com","targetPort":1003,` +
		`"targetPath":"/pathtarget/"}`
)

func TestInfoNew(t *testing.T) {
	info := InfoNew(func(i *Info) {
		i.Version = "1.2.3"
		i.Name = "Test Server"
		i.Hostname = "localhost"
		i.UUID = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i.Port = 1001
		i.ProfPort = 1002
		i.MaxConn = 9999
		i.Namespace = "dev23"
		i.AlivePath = "/pathalive/"
		i.NotifyPath = "/pathnotify/"
		i.StatusPath = "/pathstatus/"
		i.TargetHost = "myjenkins.test.com"
		i.TargetPort = 1003
		i.TargetPath = "/pathtarget/"
	})
	tp := reflect.TypeOf(info)

	if tp.Kind() != reflect.Ptr {
		t.Fatalf("Info not created as a pointer.")
	}

	tp = tp.Elem()
	if tp.Kind() != reflect.Struct {
		t.Fatalf("Info not created as a struct.")
	}
	if tp.Name() != "Info" {
		t.Fatalf("Info struct is not named correctly.")
	}
	if !(tp.NumField() > 0) {
		t.Fatalf("Info struct is empty.")
	}
}

func TestInfoString(t *testing.T) {
	t.Parallel()
	info := InfoNew(func(i *Info) {
		i.Version = "1.2.3"
		i.Name = "Test Server"
		i.Hostname = "localhost"
		i.UUID = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i.Port = 1001
		i.ProfPort = 1002
		i.MaxConn = 9999
		i.Namespace = "dev23"
		i.AlivePath = "/pathalive/"
		i.NotifyPath = "/pathnotify/"
		i.StatusPath = "/pathstatus/"
		i.TargetHost = "myjenkins.test.com"
		i.TargetPort = 1003
		i.TargetPath = "/pathtarget/"
	})
	actual := fmt.Sprint(info)
	if actual != expectedInfoJSONResult {
		t.Errorf("Info not converted to json string.\n\nExpected: %s\n\nActual: %s\n",
			expectedInfoJSONResult, actual)
	}
}
