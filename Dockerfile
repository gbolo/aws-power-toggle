# github.com/gbolo/aws-power-toggle

#
#  BUILD CONTAINER -------------------------------------------------------------
#

FROM gbolo/builder:alpine as builder

COPY . ${GOPATH}/src/github.com/gbolo/go-util/aws-power-toggle

# Building
RUN   set -xe; \
      SRC_DIR=${GOPATH}/src/github.com/gbolo/go-util/aws-power-toggle; \
      cd ${SRC_DIR}; \
      mkdir -p /tmp/build && \
      make all && \
      cp -rp www testdata/sampleconfig/power-toggle-config.yaml bin/aws-power-toggle /tmp/build/


#
#  FINAL BASE CONTAINER --------------------------------------------------------
#

FROM  gbolo/baseos:alpine

# prepare homedir
RUN   mkdir -p /opt/aws-pt

# Copy in from builder
COPY  --from=builder /tmp/build/ /opt/aws-pt/

# Run as non-privileged user by default
USER  65534

# Inherit gbolo/baseos entrypoint and pass it this argument
WORKDIR /opt/aws-pt/
CMD     ["/opt/aws-pt/aws-power-toggle", "-config", "/opt/aws-pt/power-toggle-config.yaml"]
