module github.com/bsanzhiev/taco-delivery/gateway

go 1.23.4

require (
	github.com/bsanzhiev/taco-delivery/common v0.0.0-20250212111714-ab7521a48a67
	github.com/gorilla/mux v1.8.1
	google.golang.org/grpc v1.70.0
)

require (
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250207221924-e9438ea467c6 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)

replace github.com/bsanzhiev/taco-delivery/common => ../common
