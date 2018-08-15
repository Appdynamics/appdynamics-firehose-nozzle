# appdynamics-firehose-nozzle

## How to find Credentials


- NOZZLE_TRAFFIC_CONTROLLER_URL: wss://doppler.sys.pie-20.cfplatformeng.com:443 [https://docs.pivotal.io/pivotalcf/2-2/loggregator/architecture.html#components]

```
(master)$ cf curl /v2/info | jq .doppler_logging_endpoint
"wss://doppler.sys.pie-20.cfplatformeng.com:443"
```

- NOZZLE_UAA_URL: https://uaa.{domain from the above ie.. (sys.pie-20.cfplatformeng.com)}  

- NOZZLE_USERNAME/NOZZLE_PASSWORD: 
   * easy way: login to opsmanager and go to PAS tile -> credentials -> UAA -> Opentsdb Nozzle Credentials -> Copy the username and password
   * (or) create a new account in doppler.firehose group with permissions, https://github.com/cf-platform-eng/firehose-nozzle#option-2-uaa-client

   


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
