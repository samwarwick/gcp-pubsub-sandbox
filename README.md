# gcp-pubsub-sandbox
Experiments with Google Cloud Pub/Sub.

## Stack

* [Google Cloud Pub/Sub](https://cloud.google.com/pubsub)
* go
* Node/JavaScript
* c# .NET

## Docker Pub/Sub emulator

Start the dockerized Pub/Sub Emulator.

```docker run --rm -ti -p 8681:8681 -e PUBSUB_PROJECT1=gps-demo,demo-topic:demo-sub messagebird/gcloud-pubsub-emulator:latest```

or

```./gps.sh docker```


## Clients

All client sessions most set the PUBSUB_EMULATOR_HOST enviroment variable to the emulator url.

```export PUBSUB_EMULATOR_HOST=localhost:8681```

Clients can be executed from their individual subfolders, or via `gps.sh`.
Publisher and subscribers from different SDKs can be mixed.

Supported clients:

* Go (go)
* C# .NET (dotnet)
* JavaScript (node)

See the README files in the individual client folders for additional information.

### Start a subscriber

```./gps.sh <client> sub```

e.g.

```./gps.sh go sub```

The subscriber can be halted by killing the process or sending a "quit" message.

### Publish messages

```./gps.sh <client> pub [message]```

If no message is provided then a default message is used.

e.g.

```./gps.sh node pub "My first message"```

Terminate subscriber:

```./gps.sh node pub quit```

