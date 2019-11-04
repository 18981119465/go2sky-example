#! /bin/sh
docker build -f Dockerfile-svca -t lish/go2sky-demo-svca:0.1 .
docker build -f Dockerfile-svcb -t lish/go2sky-demo-svcb:0.1 .
