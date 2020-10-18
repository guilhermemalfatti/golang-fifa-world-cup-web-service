FROM ubuntu:bionic

# Set the proxy
#ENV http_proxy http://proxy.ubisoft.org:3128
#NV https_proxy http://proxy.ubisoft.org:3128
#ENV no_proxy .ubisoft.org,.ubisoft.onbe

# Install HTTPS transport support.
RUN apt-get update && apt-get install -y apt-transport-https

# Add references to Artifactory.
RUN echo 'deb [trusted=yes] https://artifactory/debian-local default main' >> /etc/apt/sources.list
RUN echo 'Acquire::https::artifactory::Verify-Peer "false";' >> /etc/apt/apt.conf

# Avoiding extension reinstalls on container rebuild
# https://code.visualstudio.com/docs/remote/containers-advanced#_avoiding-extension-reinstalls-on-container-rebuild
RUN mkdir -p /home/root/.vscode-server/extensions \
        /home/root/.vscode-server-insiders/extensions \
    && chown -R root \
        /home/root/.vscode-server \
        /home/root/.vscode-server-insiders

# Common tooling.
RUN apt-get update && apt-get upgrade -y && apt-get install -y --allow-unauthenticated \
 	apt-utils \
	build-essential \
	git \
	tar \
	nano \
	unzip \
	software-properties-common \
	make \
	curl \
	wget \
	jq \
	openssh-client \
	dos2unix \
	zip \
	ntpdate \
	dpkg \
    wamerican \
	gettext-base \
	docker.io \
	#ubisoft-ca-certificates \
	#ubisoft-authentication \
	#ubisoft-r6s-ubiservices-cli \
	#ubisoft-r6s-cli \
	#ubisoft-r6s-code-generator \
	graphviz \
	gv \
	iputils-ping

# Flush the cache.
RUN apt-get autoclean

# Go ENV
ENV GOROOT=/usr/local/go
ENV GOPATH=/usr/go
ENV PATH=${PATH}:${GOROOT}/bin:${GOPATH}/bin
ENV GOPRIVATE=go.r6.ubi.com

# Go install
RUN mkdir -p $GOPATH
ADD https://dl.google.com/go/go1.14.1.linux-amd64.tar.gz go.tar.gz
RUN tar -xvf go.tar.gz
RUN mv go ${GOROOT}
RUN	rm go.tar.gz

# Go imports
RUN go get -u golang.org/x/lint/golint
RUN go get -u golang.org/x/tools/cmd/goimports
RUN go get -u golang.org/x/tools/gopls
RUN go get -u github.com/go-delve/delve/cmd/dlv

# go tools
RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/dustin/go-humanize


# Create simlinks
RUN echo 'rm -rf ~/.ssh && cp -r /mnt/home/.ssh ~/ && chmod -R 600 ~/.ssh' >> ~/.bashrc
RUN echo 'rm -rf ~/.aws && ln -s /mnt/home/.aws ~/.aws' >> ~/.bashrc
RUN echo 'rm -rf ~/.kube && ln -s /mnt/home/.kube ~/.kube' >> ~/.bashrc
RUN echo 'ln -sf /mnt/home/.gitconfig ~/.gitconfig' >> ~/.bashrc
