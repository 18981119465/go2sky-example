#!/bin/bash
cd svca/cmd/ && go build -tags netgo -o svca
cd ../../svcb/cmd/ && go build -tags netgo -o svcb
