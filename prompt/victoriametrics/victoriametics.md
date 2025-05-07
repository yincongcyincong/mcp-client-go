`vm_data_write` – Write time series data to VM
- Write a data point to the VM database for metric `{job: "api", instance: "localhost:8080"}` with values `[0.5, 0.6]` and timestamps `[1714032000, 1714035600]`.

`vm_prometheus_write` – Import Prometheus exposition format data
- Import the following Prometheus metrics to VictoriaMetrics:
```
http_requests_total{job="api", instance="localhost:8080"} 5
```

1. `vm_query_range` – Query time series over a time range
   - Get the CPU usage (`rate(node_cpu_seconds_total[5m])`) from `1714032000` to `1714035600`, with a step of `60s`.

2. `vm_query` – Query the current value of a time series
   - Query the current value of `up{job="api"}` at time `1714035600`.

3. `vm_labels`
   - List all the unique label names in the metrics database.

4. `vm_label_values`
   - What are all the values for the label `job`?

