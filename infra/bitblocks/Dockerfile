FROM ubuntu:xenial

RUN apt-get update --fix-missing && apt-get install --yes software-properties-common wget

RUN wget -c https://https://github.com/ReniR256/bitblocks/archive/refs/tags/v0.1.tar.gz -O - | tar xz
RUN mv ./bitblocks /app
RUN chmod +x /app/bin/bitblocksd
RUN chmod +x /app/bin/bitblocks-cli

COPY bitblocks.conf /root/.bitblocks/bitblocks.conf
COPY run.sh /root/run.sh
RUN chmod +x /root/run.sh

EXPOSE 18332

ENTRYPOINT ["./root/run.sh"]
