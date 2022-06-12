#!/bin/bash

docker build -t test-goctl .
docker run test-goctl
docker image rm -f test-goctl