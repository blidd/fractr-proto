
if [ -f "./marketplace_common/marketplace_common.pb.go" ]
then
    rm ./marketplace_common/marketplace_common.pb.go
fi
protoc -I=. --go_out=./marketplace_common ./marketplace_common/marketplace_common.proto
mv ./marketplace_common/github.com/blidd/fractr-proto/marketplace_common/marketplace_common.pb.go ./marketplace_common
rm -r ./marketplace_common/github.com

if [ -f "./marketplace_secondary/marketplace_secondary.pb.go" ]
then
    rm ./marketplace_secondary/marketplace_secondary.pb.go
fi
protoc -I=. --go_out=./marketplace_secondary --go-grpc_out=./marketplace_secondary ./marketplace_secondary/marketplace_secondary.proto
mv ./marketplace_secondary/github.com/blidd/fractr-proto/marketplace_secondary/marketplace_secondary.pb.go ./marketplace_secondary
mv ./marketplace_secondary/github.com/blidd/fractr-proto/marketplace_secondary/marketplace_secondary_grpc.pb.go ./marketplace_secondary
rm -r ./marketplace_secondary/github.com