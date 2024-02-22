# This file is for use as a devcontainer and a runtime container
#
# The devcontainer should use the build target and run as root with podman
# or docker with user namespaces.
#
FROM golang:1.22.0-bookworm 

RUN echo "TEST"
ENTRYPOINT [ "bash" ]
# change this entrypoint if it is not the same as the repo
# ENTRYPOINT ["/container-startup.sh"]
# CMD ["serve"]