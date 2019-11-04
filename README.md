# go2sky-example
go2sky example, also see go-zipkin-example: https://github.com/iyabchen/go-zipkin-example

## how to build
please execute the shell script "build.sh".

``````shell
./build.sh
``````

## how to create docker images
please execute the shell script "create-docker.sh"

``````shell
./create-docker.sh
``````

## how to run docker
* for svca: 

  ``````shell
  docker run -d --name go-demo-svca -e SKYWALKING_BACKEND_SERVICE=docker_host_ip:11800 -e -e SVCA_NAME=go_demo_svca -e SVCA_PORT=18080 SVCB_HOST=docker_host_ip -SVCB_PORT=18081 -p 18080:18080 lish/go2sky-demo-svca:0.1
  ``````

* for svcb: 

  ``````shell
  docker run -d --name go-demo-svcb -e SKYWALKING_BACKEND_SERVICE=docker_host_ip:11800 -e SVCB_NAME=go_demo_svcb -e SVCB_PORT=18081 -p 18081:18081 lish/go2sky-demo-svcb:0.1
  ``````

## TIPS

*test result*：

* service  registry: passed
* service query: passed
* topology graph: passed
* tracing: failed (Cause unknown)