# Instructions

The purpose of this repository is to reproduce the following Datadog UI issues which occur when OTEL-instrumented app calls Datadog-instrumented app:
- There is a gap between spans in DD UI
- OTEL app is not connected with DD app on Service Map

The repository contains two web apps. `rails-dd-demo` - is a RubyOnRails app instrumented with Datadog tracing. `go-otel-demo` is a golang app instrumented with OTEL tracing.
`go-otel-demo` exposes an http endpoint `http://localhost:3002/vendors` which simply calls an http endpoint of `rails-demo-app` at `http://localhost:3000/vendors`.

# Configurations

- can be found in `rails_dd_demo_config/initializers/datadog.rb` (rails app), `go-otel-demo/.env` (golang app) and `docker-compose.yaml`.

# Run Instructions

* `DD_API_KEY=XXX make restart`
* Access endpoint `http://localhost:3002/vendors`
