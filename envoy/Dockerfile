
FROM envoyproxy/envoy:36f39c746eb7d03b762099b206403935b11972d8
RUN apt-get update
COPY envoy.yaml /etc/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy.yaml