package main

import (
	"fmt"

	"context"

	"github.com/signalfx/golib/datapoint"
	"github.com/signalfx/golib/sfxclient"
)

const SignalFxToken = "*** YOUR TOKEN HERE ***"

func main() {
	fmt.Println(SignalFxToken)
	fmt.Println(sfxclient.ClientVersion)

	sink := sfxclient.NewHTTPDatapointSink()
	sink.AuthToken = SignalFxToken

	ctx := context.Background()
	err := sink.AddDatapoints(ctx, []*datapoint.Datapoint{sfxclient.Gauge("kit.test.golang", nil, 43)})
	if err != nil {
		fmt.Println(err)
	}
}
