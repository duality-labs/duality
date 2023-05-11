ARG BASE_IMAGE_TAG=latest

FROM ghcr.io/strangelove-ventures/infra-toolkit:v0.0.7 AS infra-toolkit

FROM ghcr.io/duality-labs/duality:${BASE_IMAGE_TAG}

COPY --from=infra-toolkit /usr/bin /usr/bin
COPY --from=infra-toolkit /bin/ /bin/
COPY --from=infra-toolkit /lib /lib/
COPY --from=infra-toolkit /usr/lib /usr/lib/
COPY --from=infra-toolkit /usr/local/bin /usr/local/bin

COPY scripts scripts
COPY networks networks

USER root

CMD ["sh", "./scripts/startup.sh"]