package mashup

import "testing"

func TestQueryStatus(t *testing.T) {
	data := QueryStatus("jaxleof")
	SaveStatus(data)
}
