package mysql

const (
	sqlMetricTable = "metrics"
)

type sqlMetric struct {
	ID             string `db:"id"`
	Url            string `db:"url"`
	ResponseStatus int    `db:"response_status"`
	RequestSize    int    `db:"request_size"`
	ResponseSize   int    `db:"response_size"`
	ResponseTime   int64  `db:"response_time"`
}
