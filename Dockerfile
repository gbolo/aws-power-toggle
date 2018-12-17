# github.com/gbolo/aws-power-toggle

#
#  BUILD CONTAINER -------------------------------------------------------------
#

FROM gbolo/builder:alpine as builder

COPY . ${GOPATH}/src/github.com/gbolo/aws-power-toggle

# Building
RUN   set -xe; \
      SRC_DIR=${GOPATH}/src/github.com/gbolo/aws-power-toggle; \
      cd ${SRC_DIR}; \
      mkdir -p /tmp/build && npm -v; \
      make all && \
      cp -rp frontend/dist /tmp/build/frontend; \
      cp -rp testdata/sampleconfig/power-toggle-config.yaml bin/aws-power-toggle /tmp/build/


#
#  FINAL BASE CONTAINER --------------------------------------------------------
#

FROM  gbolo/baseos:alpine

# prepare env vars
ENV   POWER_TOGGLE_SERVER_STATIC_FILES_DIR /opt/aws-pt/frontend

# prepare homedir
RUN   mkdir -p /opt/aws-pt

# Copy in from builder
COPY  --from=builder /tmp/build/ /opt/aws-pt/

# Run as non-privileged user by default
USER  65534

# Inherit gbolo/baseos entrypoint and pass it this argument
CMD     ["/opt/aws-pt/aws-power-toggle", "-config", "/opt/aws-pt/power-toggle-config.yaml"]
