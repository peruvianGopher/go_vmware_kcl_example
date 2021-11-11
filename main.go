package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/vmware/vmware-go-kcl/clientlibrary/interfaces"
	"github.com/vmware/vmware-go-kcl/logger"
)

const (
	STREAM_NAME = "abc"
	REGION = "us-west-2"
)

func main() {


	fmt.Println(logger.Debug)
}

type consumer struct {

}

func (c consumer) Initialize(initializationInput *interfaces.InitializationInput) {
	initializationInput.ShardId = "a"
	panic("implement me")
}

func (c consumer) ProcessRecords(processRecordsInput *interfaces.ProcessRecordsInput) {
	panic("implement me")
}

func (c consumer) Shutdown(shutdownInput *interfaces.ShutdownInput) {
	panic("implement me")
}
