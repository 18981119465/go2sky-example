# go2sky-example
go2sky example, also see go-zipkin-example: https://github.com/iyabchen/go-zipkin-example

## how to build
execute shell script "build.sh".

## how to create docker images
execute shell script "create-docker.sh"

## how to run docker
* for svca: docker run -d --name go-demo-svca -e SKYWALKING_BACKEND_SERVICE=192.168.176.130:11800 -e SVCB_HOST=192.168.176.130 -p 18080:18080 lish/go2sky-demo-svca:0.1
* fro svcb: docker run -d --name go-demo-svcb -e SKYWALKING_BACKEND_SERVICE=192.168.176.130:11800 -p 18081:18081 lish/go2sky-demo-svcb:0.1
