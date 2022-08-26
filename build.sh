#!/bin/zsh

# default directory for frontend js protos
FRONTEND_PB_DIR="$HOME/fractr-frontend/src/pb/"

while getopts ":d:" opt; do
    case $opt in 
        d)
            FRONTEND_PB_DIR="$OPTARG/src/pb/"
            ;;
        \?)
            echo "error occurred" >&2
            ;;
    esac
done

[ ! -d $FRONTEND_PB_DIR ] && echo "cannot find fractr-frontend directory: specify with -d [path/to/fractr-frontend]"

for dir in ./*/
do
    PROTO="${dir%/}/$(basename "$dir").proto"
    PB="${dir%/}/$(basename "$dir").pb.go"
    GRPC="${dir%/}/$(basename "$dir")_grpc.pb.go"

    # remove old files before re-generating
    rm $PB 2> /dev/null
    rm $GRPC 2> /dev/null

    echo "generating go grpc files for $PROTO..."
    protoc -I=. --go_out=$dir --go-grpc_out=$dir $PROTO

    GEN_PB="${dir%/}/github.com/blidd/fractr-proto/$(basename $dir)/$(basename $dir).pb.go"
    GEN_GRPC="${dir%/}/github.com/blidd/fractr-proto/$(basename $dir)/$(basename $dir)_grpc.pb.go"

    [ -f $GEN_PB ] && mv $GEN_PB $dir
    [ -f $GEN_GRPC ] && mv $GEN_GRPC $dir

    rm -r ${dir%/}/github.com 2> /dev/null

    # generate grpc js protos and move them to the proper directory 
    # echo $FRONTEND_PB_DIR
    # echo "${FRONTEND_PB_DIR%/}/$(basename "$dir")/"
    if [ -d "${FRONTEND_PB_DIR%/}/$(basename "$dir")/" ]
    then
        TARGET_FE_DIR="${FRONTEND_PB_DIR%/}/$(basename "$dir")"
        TARGET_FE_PB="$TARGET_FE_DIR/$(basename "$dir")_pb.js"
        TARGET_FE_GRPC_WEB="$TARGET_FE_DIR/$(basename "$dir")_grpc_web_pb.js"
        
        # echo $TARGET_FE_PB ; echo $TARGET_FE_GRPC_WEB

        rm $TARGET_FE_PB 2> /dev/null
        rm $TARGET_FE_GRPC_WEB 2> /dev/null

        echo "generating js grpc web files for $PROTO frontend..."
        protoc -I=. --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. ${PROTO}

        GEN_FE_PB="${dir%/}/$(basename "$dir")_pb.js"
        GEN_FE_GRPC_WEB="${dir%/}/$(basename "$dir")_grpc_web_pb.js"

        # echo "FILE: $GEN_FE_PB" ; echo "FILE: $GEN_FE_GRPC_WEB"

        [ -f $GEN_FE_PB ] && mv $GEN_FE_PB $TARGET_FE_DIR
        [ -f $GEN_FE_GRPC_WEB ] && mv $GEN_FE_GRPC_WEB $TARGET_FE_DIR
    fi
    
done
