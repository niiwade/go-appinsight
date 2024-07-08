package main

import (
	routes "appinsights/test/router"
	"fmt"
	"os"
	"time"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

// setup appinsights
var telemetryClient = appinsights.NewTelemetryClient(os.Getenv("InstrumentationKey"))

func main() {

	fmt.Println(telemetryClient)
	r := routes.SetupRouter()

	/*Set role instance name globally -- this is usually the name of the service submitting the telemetry*/
	telemetryClient.Context().Tags.Cloud().SetRole("Student")

	telemetryClient.TrackTrace("Starting Staudent Api", appinsights.Information)

	appinsights.NewDiagnosticsMessageListener(func(msg string) error {
		fmt.Printf("[%s] %s\n", time.Now().Format(time.UnixDate), msg)
		return nil
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
