FROM namely/protoc-all as builder

RUN apt-get -y install git

WORKDIR /
RUN git clone https://github.com/Dopl-Technologies/api-protos.git
RUN cp /api-protos/dopl/api/* /defs

WORKDIR /defs
RUN GEN_LANG=go /usr/local/bin/entrypoint.sh -d /defs -o /gen

WORKDIR /output
RUN cp /gen/github.com.dopl.protos/*.go .