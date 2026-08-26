package hermes

import (
	"testing"

	"github.com/heliantheon/common/filter"
)

func TestApplicationFiltersAllowExactApplicationID(t *testing.T) {
	t.Parallel()

	allowed, ok := applicationFilters["app_id"]
	if !ok {
		t.Fatal("application filters must expose app_id")
	}
	for _, operation := range allowed {
		if operation == filter.Eq {
			return
		}
	}
	t.Fatal("application app_id filter must allow exact matching")
}
