FROM ubuntu:latest
LABEL authors="wang"

ENTRYPOINT ["top", "-b"]