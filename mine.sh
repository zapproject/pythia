#!/bin/bash
status=`./zap-miner stake status`

not_staked="Not currently staked"

echo $status

if [[ "$status" == "$not_staked" ]]; then
    ./zap-miner stake deposit
fi

./zap-miner mine -r
