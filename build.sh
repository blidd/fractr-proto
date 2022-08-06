#!/bin/zsh

for dir in ./*/
do
    PROTO="${dir%/}/$(basename "$dir").proto"
    PB="${dir%/}/$(basename "$dir").pb.go"
    GRPC="${dir%/}/$(basename "$dir")_grpc.pb.go"

    echo "generating gRPC files for ${PROTO}..."
    
    if [ -f ${PB} ]
    then
        rm ${PB}
    fi

    if [ -f ${GRPC} ]
    then
        rm ${GRPC}
    fi

    protoc -I=. --go_out=$dir --go-grpc_out=$dir ${PROTO}

    GEN_PB="${dir%/}/github.com/blidd/fractr-proto/$(basename $dir)/$(basename $dir).pb.go"
    GEN_GRPC="${dir%/}/github.com/blidd/fractr-proto/$(basename $dir)/$(basename $dir)_grpc.pb.go"

    if [ -f ${GEN_PB} ]
    then
        mv ${GEN_PB} $dir
    fi
    
    if [ -f ${GEN_GRPC} ]
    then
        mv ${GEN_GRPC} $dir
    fi

    rm -r ${dir%/}/github.com
    
done
