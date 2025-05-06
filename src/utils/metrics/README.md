# prometheus metrics Package

The `prometheus metrics` package provides functionality for collecting and recording various metrics associated with HTTP requests in our backend service. It utilizes the Prometheus library for metric instrumentation. this package offers an efficient and robust solution for monitoring and analyzing service performance.

## Functions

### Init()

The function `Init` initializes two maps, counters and histograms, to store Prometheus Counter and Histogram vectors respectively.

### Inc()

The function `Inc` Increments the counter associated with the given key.

- **Parameters**:
  - `ctx *gin.Context` : Gin context representing the HTTP request being processed.
  - `key string` : Key used to identify the counter metric.
  - `service string` : service name like authentication,watchList etc.
  - `path string` : Information about the request path.
  - `method string` : HTTP method of the request.
  - `statusCode string` : Status code of the request.

### Recording Histograms

The function `Record` Records a metric value for a given key.

- **Parameters**:
  - `ctx *gin.Context` : Gin context representing the HTTP request being processed.
  - `key string` : Key used to identify the histogram metric.
  - `service string` : service name like authentication,watchList etc.
  - `timeInMillis float64`: Time taken for the request in milliseconds.
  - `path string` : Information about the request path.
  - `requestType string` :  Type of the request Like BFF,NEST,OVERALL etc.

### Register Counter

The function `registerCounter` registers a counter metric with the specified key. It utilizes the `prometheus.NewCounterVec` function to create a new counter vector with the given options. If a counter metric with the same key already exists in the counters map, the function returns without registering a new one. Otherwise, it registers the newly created counter vector using `prometheus.MustRegister`.

- **Parameters**:
  - `key string` : A string representing the key used to identify the counter metric.

### Register Histogram

The function `registerHistogram` registers a histogram metric with the specified key. Similar to registerCounter, it utilizes the `prometheus.NewHistogramVec` function to create a new histogram vector with the given options. If a histogram metric with the same key already exists in the histograms map, the function returns without registering a new one. Otherwise, it registers the newly created histogram vector using `prometheus.MustRegister`.

- **Parameters**:
  - `key string` : A string representing the key used to identify the histogram metric.

### Usage Notes:
1) Thread safety is ensured during access and modification of metrics maps through the use of a mutex.
2) Additional information extracted from the Gin context, such as timestamp, request ID, and hostname, is utilized as labels for Prometheus metrics.
3) Metrics are registered only once for each key, optimizing resource usage.

#### How To Use

To use the metrics package effectively, follow these steps:

1) `Increment Counters`: Use the `Inc()` function to increment counter metrics. For example, you can call this function to count the number of HTTP requests received by your server.

```
func Handler(c *gin.Context) {
    // Your handler logic here
    
    // Increment counter for the endpoint accessed
    metrics.Inc(c, "endpoint_accessed", "your_service_name", c.Request.URL.Path, c.Request.Method, "200")
    
    // Your handler logic here
}
```

2) `Record Histograms`: Use the `Record()` function to record histogram metrics. For example, you can use this function to record the latency of your HTTP requests.

```
func Handler(c *gin.Context) {
    // Your handler logic here
    
    startTime := time.Now()
    // Your handler logic here
    // Measure the time taken
    latency := time.Since(startTime)
    
    // Record latency histogram for the endpoint accessed
    metrics.Record(c, "endpoint_latency", latency.Seconds()*1000, c.Request.URL.Path, c.Request.Method, "your_service_name")
}
```

 