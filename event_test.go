package debeziummodels

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimeMS_Time(t *testing.T) {
	timeMS := TimeMS(1674268794990093)
	converted := timeMS.Time()
	require.Equal(t, time.Date(2023, 01, 21, 2, 39, 54, 990093, time.UTC), converted)
}
