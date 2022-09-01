module github.com/yolo-sh/hetzner-cloud-provider

go 1.18

replace github.com/yolo-sh/yolo v0.0.0 => ../yolo

replace github.com/yolo-sh/agent v0.0.0 => ../agent

require (
	github.com/golang/mock v1.6.0
	github.com/hetznercloud/hcloud-go v1.35.2
	github.com/mikesmitty/edkey v0.0.0-20170222072505-3356ea4e686a
	github.com/pelletier/go-toml v1.9.5
	github.com/yolo-sh/agent v0.0.0
	github.com/yolo-sh/yolo v0.0.0
	golang.org/x/crypto v0.0.0-20220313003712-b769efc7c000
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gosimple/slug v1.12.0 // indirect
	github.com/gosimple/unidecode v1.0.1 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
