#!/usr/bin/env bash

function check_client_args () {
    if [[ $# -lt 2 ]]; then
        echo "usage: gps $1 <pub|sub>"
        exit 1
    fi
}

if [[ $# -lt 1 ]]; then
    echo "usage: gps <command>"
    exit 1
fi

GPSROOT=$(dirname $(readlink -f $0))

case $1 in
    docker)
        docker run --rm -ti -p 8681:8681 -e PUBSUB_PROJECT1=gps-demo,demo-topic:demo-sub messagebird/gcloud-pubsub-emulator:latest
        ;;
    dotnet)
        check_client_args $*
        if [[ $2 == pub* ]]; then
            cd $GPSROOT/dotnet/publisher
            dotnet run "$3"
        elif [[ $2 == sub* ]]; then
            cd $GPSROOT/dotnet/subscriber
            dotnet run
        else
            echo "Invalid dotnet command --" $2
            exit 1
        fi
        ;;
    go)
        check_client_args $* 
        if [[ $2 == pub* ]]; then
            cd $GPSROOT/go/publisher
            go run . "$3"
        elif [[ $2 == sub* ]]; then
            cd $GPSROOT/go/subscriber
            go run .
        else
            echo "Invalid go command --" $2
            exit 1
        fi
        ;;
    node)
        check_client_args $*
        cd $GPSROOT/node
        if [[ $2 == pub* ]]; then
            node publisher "$3"
        elif [[ $2 == sub* ]]; then
            node subscriber
        else
            echo "Invalid node command --" $2
            exit 1
        fi
        ;;
    *)
        echo "Invalid command --" $1
        ;;
esac

exit 0
