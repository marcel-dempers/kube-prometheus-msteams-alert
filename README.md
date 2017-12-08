# kube-prometheus-msteams-alert

Forwarding application that forwards requests from Prometheus alertmanager into Microsoft Teams

## Use Case

* You are running a Kubernetes cluster
* You monitor your environment using Prometheus
* You have Prometheus Alert Manager running in your cluster and configured to generate alerts
* You would like to have alerts sent to Microsoft Teams


## What does it do ?

This app runs inside Kubernetes and listens on a port for Prometheus Alertmanager messages. </br>
When alert comes in, the application will parse the alert and send it on the the configured receiver </br>

It translates alert manager messages into `JSON` format that MS Teams will understand <br/>

## Usage

### Build it!

Build the docker image and push it to your registry!

```
cd src
docker build . -t alertmanager-msteams

docker push alertmanager-msteams

```

### Edit your values !

Go to the values file at the root of this repo and fill out your values.
Key values are:

* image : Your docker image and registry
* Your MS Teams webhook URL under configmap section

### Deploy it !

Check your `kubectl` is configured to point to your cluster <br/>
Deploy using `helm` 
```
helm install alertmanager-msteams -f values.yaml
```

### Test

Simply POST `json` data to the /alert endpoint of this application.
The `json` format needs to represent Prometheus structure described [here](https://prometheus.io/docs/alerting/configuration/#webhook_config) </br>

Your Prometheus Alertmanager needs to be configured to push alerts to external web hook:
```
<webhook_config>
The webhook receiver allows configuring a generic receiver.

```

Configuring your Prometheus Alert Manager to push to this endpoint will allow this app to forward convert the alert manager `json` into MS teams friendly JSON and pass it onto MS Teams webhook endpoint



