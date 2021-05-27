module order

go 1.16

replace geo => ../geo

require (
	geo v0.0.0
	github.com/go-kit/kit v0.10.0
	github.com/gogo/protobuf v1.3.2
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/metaverse/truss v0.2.1
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.38.0
)