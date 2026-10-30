# Dashboards

The Bookstore's dashboards, in each dashboard tool's native format:

- `bookstore-services.yaml` (Perses)
- `bookstore-services.json` (Grafana)

VKS Toolkit provisions them into the stack's dashboards provider when the app
is observed by OpenTelemetry eBPF Instrumentation, and lists them on the app's
row in the stack's Services tab.
