 
  ENV['DD_PROPAGATION_STYLE_INJECT'] = 'Datadog,B3'
  ENV['DD_PROPAGATION_STYLE_EXTRACT'] = 'Datadog,B3'
  
 # Configure datadog env variables if you want to set static variables
Datadog.configure do |c|
  # Add additional configuration here.
  # Activate integrations, change tracer settings, etc...
  # Agent conf
  c.agent.host = ENV['DD_AGENT_HOST']
  c.agent.port = 8126

  # App conf
  c.env = "development"
  c.version = "1.0.0"
  c.tracing.enabled = true
  c.tracing.test_mode.enabled = false
end
