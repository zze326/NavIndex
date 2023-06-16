FROM alpine:3.18.2
ADD NavIndex /opt/NavIndex
COPY data.yaml /opt/conf/data.yaml
COPY dp2.png /opt/dp2.png
WORKDIR /opt
CMD ["/opt/NavIndex"]
