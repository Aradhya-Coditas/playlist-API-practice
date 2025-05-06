package metrics

import (
	"fmt"
	"omnenest-backend/src/constants"
	"omnenest-backend/src/utils/configs"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	counters      map[string]*prometheus.CounterVec
	histograms    map[string]*prometheus.HistogramVec
	mu            sync.Mutex
	labels        = []string{"service", "path", "method", "statusCode", "requestId", "timestamp", "hostname"}
	latencyLabels = []string{"path", "requestType", "service", "requestId", "timestamp", "hostname"}
)

// Init initializes the counters and histograms maps.
func Init() {
	counters = make(map[string]*prometheus.CounterVec)
	histograms = make(map[string]*prometheus.HistogramVec)
}

// Inc increments the counter associated with the given key.
func Inc(ctx *gin.Context, key, service, path, method, statusCode string) {
	timestamp := ctx.Value(constants.APIRequestTime).(string)
	requestId := ctx.Value(constants.RequestIDHeader).(string)
	hostName := configs.GetHostName()
	if value, ok := counters[key]; ok {
		value.WithLabelValues(service, path, method, statusCode, requestId, timestamp, hostName).Inc()
	} else {
		registerCounter(key)
		counters[key].WithLabelValues(service, path, method, statusCode, requestId, timestamp, hostName).Inc()
	}
}

// Record is a function that records a metric value for a given key.
func Record(ctx *gin.Context, key string, timeInMillis float64, path, requestType, service string) {
	metricKey := fmt.Sprintf("%s_hist", key)
	timestamp := ctx.Value(constants.APIRequestTime).(string)
	requestId := ctx.Value(constants.RequestIDHeader).(string)
	hostName := configs.GetHostName()
	if value, ok := histograms[metricKey]; ok {
		value.WithLabelValues(path, requestType, service, requestId, timestamp, hostName).Observe(timeInMillis)
	} else {
		registerHistogram(metricKey)
		histograms[metricKey].WithLabelValues(path, requestType, service, requestId, timestamp, hostName).Observe(timeInMillis)
	}
}

// registerCounter is a function that registers a counter metric with the given key.
func registerCounter(key string) {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := counters[key]; ok {
		return
	}
	counters[key] = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: constants.OmnenestKey,
		Name:      key,
		Help:      "counter metric",
	}, labels)
	prometheus.MustRegister(counters[key])
}

// registerHistogram is a function that registers a histogram metric with the given key.
func registerHistogram(key string) {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := histograms[key]; ok {
		return
	}
	histograms[key] = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: constants.OmnenestKey,
		Name:      key,
		Help:      "histogram metric",
	}, latencyLabels)
	prometheus.MustRegister(histograms[key])
}
