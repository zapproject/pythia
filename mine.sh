#!/bin/bash
status=`./pythia stake status`

not_staked="Not currently staked"

echo $status

if [[ "$status" == "$not_staked" ]]; then
    ./pythia stake deposit
fi

./pythia mine -r
