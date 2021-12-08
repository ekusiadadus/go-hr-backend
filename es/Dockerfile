ARG ELASTIC_VER=7.8.1
ARG SUDACHI_PLUGIN_VER=2.0.3

FROM ibmjava:8-jre-alpine as dict_builder

ARG ELASTIC_VER
ARG SUDACHI_PLUGIN_VER

WORKDIR /home

RUN wget https://github.com/WorksApplications/elasticsearch-sudachi/releases/download/v${ELASTIC_VER}-${SUDACHI_PLUGIN_VER}/analysis-sudachi-${ELASTIC_VER}-${SUDACHI_PLUGIN_VER}.zip && \
    unzip analysis-sudachi-${ELASTIC_VER}-${SUDACHI_PLUGIN_VER}.zip && \
    wget http://sudachi.s3-website-ap-northeast-1.amazonaws.com/sudachidict/sudachi-dictionary-20210802-core.zip && \
    unzip sudachi-dictionary-20210802-core.zip && \
    mkdir -p /usr/share/elasticsearch/config/sudachi/ && \
    mv sudachi-dictionary-20210802/system_core.dic /usr/share/elasticsearch/config/sudachi/ && \
    rm -rf sudachi-dictionary-20210802-core.zip sudachi-dictionary-20210802/


FROM docker.elastic.co/elasticsearch/elasticsearch:${ELASTIC_VER}

ARG ELASTIC_VER
ARG SUDACHI_PLUGIN_VER

COPY es/sudachi/sudachi.json /usr/share/elasticsearch/config/sudachi/
COPY --from=dict_builder /home/analysis-sudachi-${ELASTIC_VER}-${SUDACHI_PLUGIN_VER}.zip /usr/share/elasticsearch/