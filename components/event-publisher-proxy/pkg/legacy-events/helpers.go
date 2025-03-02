package legacy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/kyma-project/kyma/components/event-publisher-proxy/pkg/legacy-events/api"
)

const (
	// eventTypeFormat is driven by BEB specification.
	// An EventType must have at least 4 segments separated by dots in the form of:
	// <domainNamespace>.<businessObjectName>.<operation>.<version>
	eventTypeFormat = "%s.%s.%s.%s"

	// eventTypeFormatWithoutPrefix must have at least 3 segments separated by dots in the form of:
	// <businessObjectName>.<operation>.<version>
	eventTypeFormatWithoutPrefix = "%s.%s.%s"
)

// ParseApplicationNameFromPath returns application name from the URL.
// The format of the URL is: /:application-name/v1/...
func ParseApplicationNameFromPath(path string) string {
	// Assumption: Clients(application validator which has a flag for the path (https://github.com/kyma-project/kyma/blob/main/components/application-connectivity-validator/cmd/applicationconnectivityvalidator/applicationconnectivityvalidator.go#L49) using this endpoint must be sending request to path /:application/v1/events
	// Hence it should be safe to return 0th index as the application name
	pathSegments := make([]string, 0)
	for _, segment := range strings.Split(path, "/") {
		if strings.TrimSpace(segment) != "" {
			pathSegments = append(pathSegments, segment)
		}
	}
	return pathSegments[0]
}

// is2XXStatusCode checks whether status code is a 2XX status code
func is2XXStatusCode(statusCode int) bool {
	return statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices
}

// writeJSONResponse writes a JSON response
func writeJSONResponse(w http.ResponseWriter, resp *api.PublishEventResponses) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", ContentTypeApplicationJSON)

	if resp.Error != nil {
		w.WriteHeader(resp.Error.Status)
		_ = encoder.Encode(resp.Error)
		return
	}

	if resp.Ok != nil {
		_ = encoder.Encode(resp.Ok)
		return
	}

	log.Errorf("received an empty response")
}

// formatEventType4BEB format eventType as per BEB spec
func formatEventType4BEB(eventTypePrefix, app, eventType, version string) string {
	if len(strings.TrimSpace(eventTypePrefix)) == 0 {
		return fmt.Sprintf(eventTypeFormatWithoutPrefix, app, eventType, version)
	}
	return fmt.Sprintf(eventTypeFormat, eventTypePrefix, app, eventType, version)
}
