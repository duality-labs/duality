#!/bin/sh

# dualityd init duality

# dualityd add-consumer-section


dualityd --log_level ${LOG_LEVEL:-info} start

if [ "$KEEP_RUNNING" != "false" ]
then
    tail -f /dev/null;
fi
