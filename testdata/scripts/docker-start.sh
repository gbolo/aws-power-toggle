#!/usr/bin/env bash

docker stop aws-power-toggle && docker rm aws-power-toggle

docker run -d --name "aws-power-toggle" \
  -p 8080:8080 \
  -e "POWER_TOGGLE_SERVER_ACCESS_LOG=true" \
  -e "AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID}" \
  -e "AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY}" \
  gbolo/aws-power-toggle:3.1

