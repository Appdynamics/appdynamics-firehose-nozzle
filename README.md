# appdynamics-firehose-nozzle

## Edit the following manifest.yml
```
      NOZZLE_UAA_URL: 'https://uaa.<ops domain>'
      NOZZLE_TRAFFIC_CONTROLLER_URL: 'wss://doppler.<domain>:443'
      NOZZLE_USERNAME: opentsdb-firehose-nozzle
      NOZZLE_PASSWORD: <password>
      APPD_CONTROLLER: <controller>
      APPD_ACCESSKEY: <accesskey>
      APPD_PORT: <port>
      APPD_ACCOUNT: <account>
```

## push the application

```
cf push
```
