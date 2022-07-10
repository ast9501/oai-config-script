FROM alan0415/oai-gnb:v0.1

ADD bin/setup-cu /oai
ADD bin/setup-du /oai

RUN apt update && \
    apt install -y net-tools && \
    rm -rf /var/lib/apt/lists/*
USER root
EXPOSE 2152 600 601 38412 30923 4043
