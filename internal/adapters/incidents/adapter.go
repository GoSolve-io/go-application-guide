package incidents

// Adapter uses bikewise.org API for bike incidents data.
// See: https://www.bikewise.org/documentation/api_v2#!/incidents/GET_version_incidents_format_get_0
type Adapter struct {
	address string
}
