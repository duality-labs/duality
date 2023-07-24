Duality is a decentralized exchange optimized for capital efficiency and composability while protecting users from harmful price manipulation. With Duality, liquidity providers and protocols have intuitive, fine grain control over parameters and distribution of liquidity.

Build Steps:

1. [Install the Ignite CLI](https://docs.ignite.com/guide/install)
    - if you are using VSCode for development: an environment with the Ignite CLI already installed is available using the [Dev Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extension and our .devcontainer configuration.
1. To build for your local machine run `ignite chain build`. You can customize the build target like so: `ignite chain build -t linux:amd64`. The binary will output to $GOPATH/bin/dualityd.

For more on duality's functionality and design check out our [documentation](https://duality.gitbook.io/duality-documentation/concepts)
