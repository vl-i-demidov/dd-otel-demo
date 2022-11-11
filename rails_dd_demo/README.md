# Instructions

# configurations

- can be cound in `config/initializers/datadog.rb`

# Run Instructions

* RUN `docker build -t dd_rails_demo .`
* RUN `docker run --name dd_rails_demo -e DD_AGENT_HOST=datadog-agent -p 3000:3000 dd_rails_demo`
* Access endpoint `http://localhost:3000/vendors`
