# appdynamics-firehose-nozzle


## Edit the following manifest.yml
```
      NOZZLE_UAA_URL: 'https://uaa.<ops domain>' # cf curl /v2/info | jq .doppler_logging_endpoint
      NOZZLE_TRAFFIC_CONTROLLER_URL: 'wss://doppler.<domain>:443' # cf curl /v2/info | jq .token_endpoint 
      NOZZLE_USERNAME: opentsdb-firehose-nozzle
      NOZZLE_PASSWORD: <password>
      APPD_CONTROLLER_HOST: <controller>
      APPD_ACCESS_KEY: <accesskey>
      APPD_CONTROLLER_PORT: <port>
      APPD_ACCOUNT: <account>
```

## Push the application

```
cf push
```

## How to find Credentials


- [NOZZLE_TRAFFIC_CONTROLLER_URL](https://docs.pivotal.io/pivotalcf/2-2/loggregator/architecture.html#components)

```
(master)$ cf curl /v2/info | jq .doppler_logging_endpoint
"wss://doppler.sys.pie-20.cfplatformeng.com:443"
```

- *NOZZLE_UAA_URL* 
```
(master)$ cf curl /v2/info | jq .token_endpoint
"https://uaa.sys.pie-multi-az-blue.cfplatformeng.com"
```

- *NOZZLE_USERNAME/NOZZLE_PASSWORD* 
   * **Easy way** login to opsmanager and go to PAS tile -> credentials -> UAA -> Opentsdb Nozzle Credentials -> Copy the username and password. CF environments usually ships with `opentsdb-firehose-nozzle` account which already belongs to `doppler.firehose` group.  
   
   * (or) create a new account in `doppler.firehose` group with permissions, https://github.com/cf-platform-eng/firehose-nozzle#option-2-uaa-client
