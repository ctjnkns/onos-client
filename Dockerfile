FROM onosproject/onos:2.7-latest

COPY ENTRYPOINT.sh /

RUN apt-get update && apt-get install -y --no-install-recommends \
    #curl \
    iproute2 \
    iputils-ping \
    mininet \
    net-tools \
    openvswitch-switch \
    openvswitch-testcontroller \
    #x11-xserver-utils \
    #xterm \
    openssh-client \
 && rm -rf /var/lib/apt/lists/* \
 && chmod +x /ENTRYPOINT.sh \
 && ln /usr/bin/ovs-testcontroller /usr/bin/controller


EXPOSE 6633 6653 6640 8101

#ENTRYPOINT ["/ENTRYPOINT.sh"]
#seems like this over-rides, so remove it and run it in the machine after the build.

