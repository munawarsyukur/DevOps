# Setup Kibana in AWS EC2

Below steps are to setup Kibana locally inside the AWS EC2

### Update the OS

```bash
$sudo yum -y update
```

### Download the Kibana-OSS

Due to license and the AWS Elastic Search API version currently available is 7.1, we are going to use Kibana OSS 7.1.1.

```bash
$wget https://artifacts.elastic.co/downloads/kibana/kibana-oss-7.1.1-x86_64.rpm
```

### Install the Kibana-OSS

```bash
$rpm -ivh kibana-oss-7.1.1-x86_64.rpm
```

### Enable the Kibana-OSS as Systemd Service and auto start.

```bash
1. $sudo /bin/systemctl status kibana.services --> check the status
2. $sudo /bin/systemctl enable --now kibana.services --> enable the service immediately
```

### Configure the Kibana-OSS

Change the kibana.services configuration and add the AWS ES endpoint

```bash
$sudo vi /etc/kibana/kibana.yml
```
Inside kibana.yml remove the # and save it.
```bash
server.port: 5601
server.host: "0.0.0.0"
elasticsearch.url: "http://ec2-18-197-147-233.eu-central-1.compute.amazonaws.com:9200/"
```
### Ensure the Kibana-OSS up & running

Open browser and type AWS EC2 public IP with port 5601 or (sample: http://ec2-54-255-154-134.ap-southeast-1.compute.amazonaws.com:5601/). Or can use CURL

```bash
$curl https://localhost:5601
```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
Copyright 2019 Munawar Syukur

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.