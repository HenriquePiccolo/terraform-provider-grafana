module github.com/HenriquePiccolo/terraform-provider-grafana-v2

go 1.16

replace (
	github.com/grafana/grafana-api-golang-client v0.0.0-20210720012848-3049c6914b86 => github.com/HenriquePiccolo/grafana-api-golang-client v1.0.0
	k8s.io/client-go v12.0.0+incompatible => k8s.io/client-go v0.21.3
)

require (
	github.com/grafana/grafana-api-golang-client v0.0.0-20210720012848-3049c6914b86
	github.com/grafana/synthetic-monitoring-agent v0.0.24
	github.com/grafana/synthetic-monitoring-api-go-client v0.0.2
	github.com/hashicorp/go-cleanhttp v0.5.2
	github.com/hashicorp/terraform-plugin-docs v0.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
)
