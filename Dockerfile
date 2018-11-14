FROM golang:1.10.2-stretch

WORKDIR $GOPATH/src/github.com/iotexproject/iotex-core/

RUN apt-get install -y --no-install-recommends make

COPY ./ $GOPATH/src/github.com/iotexproject/iotex-core/

ARG SKIP_DEP=false

RUN if [ "$SKIP_DEP" != true ] ; \
    then \
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
        dep ensure --vendor-only; \
    fi

RUN mkdir -p $GOPATH/src/github.com/CoderZhi/go-ethereum/ && \
    mkdir -p $GOPATH/pkg/linux_amd64/github.com/CoderZhi/ && \
    rm -rf $GOPATH/src/github.com/iotexproject/iotex-core/vendor/github.com/CoderZhi/go-ethereum/ && \
    cd $GOPATH/src/github.com/iotexproject/iotex-core/pkg/ && \
    tar -xzvf go-ethereum.tar.gz && \
    cp -r $GOPATH/src/github.com/iotexproject/iotex-core/pkg/go-ethereum/binary_linux/* $GOPATH/pkg/linux_amd64/github.com/CoderZhi/ && \
    cp -r $GOPATH/src/github.com/iotexproject/iotex-core/pkg/go-ethereum/go-ethereum/* $GOPATH/src/github.com/CoderZhi/go-ethereum/ && \
    rm -rf $GOPATH/src/github.com/iotexproject/iotex-core/pkg/go-ethereum/ && \
    cd $GOPATH/src/github.com/iotexproject/iotex-core/ && \
    make clean build && \
    ln -s $GOPATH/src/github.com/iotexproject/iotex-core/bin/server /usr/local/bin/iotex-server  && \
    ln -s $GOPATH/src/github.com/iotexproject/iotex-core/bin/actioninjector /usr/local/bin/iotex-actioninjector && \
    mkdir -p /usr/local/lib/iotex/ && \
    ln -s $GOPATH/src/github.com/iotexproject/iotex-core/bin/addrgen /usr/local/bin/iotex-addrgen && \
    mkdir -p /usr/local/lib/iotex/ && \
    cp $GOPATH/src/github.com/iotexproject/iotex-core/crypto/lib/libsect283k1_ubuntu.so /usr/lib/ && \
    cp $GOPATH/src/github.com/iotexproject/iotex-core/crypto/lib/blslib/libtblsmnt_ubuntu.so /usr/lib/ && \
    mkdir -p /etc/iotex/ && \
    ln -s $GOPATH/src/github.com/iotexproject/iotex-core/blockchain/testnet_actions.yaml /etc/iotex/testnet_actions.yaml && \
    cp $GOPATH/src/github.com/iotexproject/iotex-core/e2etest/config_local_delegate.yaml /etc/iotex/config_local_delegate.yaml && \
    cp $GOPATH/src/github.com/iotexproject/iotex-core/e2etest/config_local_fullnode.yaml /etc/iotex/config_local_fullnode.yaml

CMD [ "iotex-server", "-config-path=/etc/iotex/config_local_fullnode.yaml"]
