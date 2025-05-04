# solar-zero-scrape-golang

Solar Zero Scrape re-written and improved in go

## Config - Docker

* Copy config-template.json to config.json, edit. change username/password for solar zero and set the influx and mqtt sections.
* Run docker
  eg: `docker run -v $(pwd)/solar-zero-scrape.json:/config/solar-zero-scrape.json ghcr.io/mwinters-stuff/solar-zero-scrape-golang:latest`

* Or use --help to get command line parameters and pass as parameters.

## Config - Helm Chart

* make a directory,
* add in `repository.yaml`
```
---
apiVersion: source.toolkit.fluxcd.io/v1beta2
kind: HelmRepository
metadata:
  name: solar-zero-scrape-golang
  namespace: apps
spec:
  interval: 24h
  url: https://mwinters-stuff.github.io/solar-zero-scrape-golang/

```

* add a `release.yaml`
* fill with
```
apiVersion: helm.toolkit.fluxcd.io/v2beta1
kind: HelmRelease
metadata:
    name: solar-zero-scrape
    namespace: apps
spec:
    chart:
        spec:
            chart: solar-zero-scrape
            reconcileStrategy: ChartVersion
            sourceRef:
                kind: HelmRepository
                name: solar-zero-scrape-golang
                namespace: apps
    interval: 12h
    values:
        image:
            tag: 1.17.2
        replicaCount: 1
        TimeZone: Pacific/Auckland
        host: 0.0.0.0
        service:
            enabled: true
            annotations: {}
            type: ClusterIP
            httpPort: 9898
            externalPort: 9898
        serviceAccount:
            # Specifies whether a service account should be created
            enabled: false
            # The name of the service account to use.
            # If not set and create is true, a name is generated using the fullname template
            # List of image pull secrets if pulling from private registries
            # imagePullSecrets: []
        securityContext: {}
        resources:
            limits:
                cpu: 0.5
                memory: 16Mi
            requests:
                cpu: 0.25
                memory: 16Mi
        probes:
            readiness:
                initialDelaySeconds: 60
                timeoutSeconds: 5
                failureThreshold: 3
                successThreshold: 1
                periodSeconds: 60
            liveness:
                initialDelaySeconds: 60
                timeoutSeconds: 5
                failureThreshold: 3
                successThreshold: 1
                periodSeconds: 60
        tls:
            enabled: false
            # the name of the secret used to mount the certificate key pair
            secretName: null
            # the path where the certificate key pair will be mounted
            certPath: /data/cert
            # the port used to host the tls endpoint on the service
            port: 9899
            # the port used to bind the tls port to the host
            # NOTE: requires privileged container with NET_BIND_SERVICE capability -- this is useful for testing
            # in local clusters such as kind without port forwarding
            hostPort: null
        # create a certificate manager certificate (cert-manager required)
        # may not be necessary.
        certificate:
            create: false
            # the issuer used to issue the certificate
            issuerRef:
                kind: ClusterIssuer
                name: self-signed
            # the hostname / subject alternative names for the certificate
            dnsNames:
                - solar-zero-scrape
        solarZero:
            username: your@username
            password: yourpassword
        influxdb:
            hostUrl: url
            token: atoken
            org: place.nz
            bucket: solarzero
            measurement: solarzero
        mqtt:
            hosturl: mqtt://mosquitto.apps.svc:1883
            topic: solar-zero
            username: solarzero
            password: zerosolar

```

If you dont want influxdb, or mqtt those can be removed from the configuration.
