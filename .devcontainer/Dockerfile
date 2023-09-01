FROM mcr.microsoft.com/devcontainers/go

# WASMEDGE
WORKDIR /opt/wasmedge

RUN curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash

RUN echo 'source "$HOME/.wasmedge/env"' >> $HOME/.bash_profile

# WASMTIME
WORKDIR /opt/wasmtime

RUN curl https://wasmtime.dev/install.sh -sSf | bash

ENV PATH "$PATH:/root/.wasmtime/bin/;" ls