# appdynamics-firehose-nozzle


## PRE-REQUISITE

-  Download [Appdynamics GoLang SDK for linux x64](https://download.appdynamics.com/download/#version=&apm=golang-sdk)
- `tar zxvf golang-sdk-x64-linux-4.5.1.0.tar`
-  `cp -r lib vendor/appdynamics/`

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

## See the metrics

```controller->appd-nozzle[default AppName]->MetricBrowser->Application Infrastructure Performance|appd-nozzle-tier[default tier name]|Individual Nodes```


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


## Overriding configuration

Although for the most part the nozzle application creates default configuration itself, one can override the configuration by setting the following environemnt variables and restaging the application 
   
     - cf set-env appdnozzle <ENVNAME> <NEW ENV VALUE>
     - cf restage appdnozzle
     
  | Environment Variable          	| Purpose                                                               	| Allowed Values                              	| Default Value    	|
|-------------------------------	|-----------------------------------------------------------------------	|---------------------------------------------	|------------------	|
| APPD_NOZZLE_APP          	     | Name of the Nozzle Application under which the metrics are recorded   	| Any string                                     	| appd-nozzle      	|
| APPD_NOZZLE_TIER            	| Name of the Nozzle Tier under which the metrics are recorded          	| Any String                                     	| appd-nozzle-tier 	|
| APPD_NOZZLE_NODE             	| Name of the Nozzle Node under which the metrics are recorded          	| Any String                                  	| appd-nozzle-node 	|
| APPD_SSL_ENABLED              	| Enable/Disable SSL to Controller                                      	| true/false                                  	| false            	|
| APPD_CONTROLLER_HOST          	| Hostname of Appdynamics Controller                                    	| host.appd.com                               	|                  	|
| APPD_CONTROLLER_PORT          	| Port on which Appdynamics Controller is listening                     	| port number                                 	| 8090             	|
| APPD_ACCOUNT                  	| Account name for the above controller                                 	| Account name                                	|                  	|
| APPD_ACCESS_KEY               	| Access Key                                                            	| Access Key                                  	|                  	|
| APPD_SINK                     	| Sink to which the metrics are to be pushed.                           	| stdout/Controller                           	| Controller       	|
| APPD_SAMPLING_RATE            	| Polling Interval in secs to Firehose Nozzle.                          	| number of seconds                           	| 2 secs           	|
| NOZZLE_UAA_URL                	| UAA Api endpoint URL                                                  	| cf curl /v2/info and record UAA endpoint    	|                  	|
| NOZZLE_TRAFFIC_CONTROLLER_URL 	| Doppler end point URL                                                 	| cf curl /v2/info and record doppler api url 	|                  	|
| NOZZLE_USERNAME               	| User name of account belonging to doppler.firehose group              	| user name, easy: use opentsb credentials    	|                  	|
| NOZZLE_PASSWORD               	| Password for the above account                                        	|                                             	|                  	|
      

