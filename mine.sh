#!/bin/bash
status=`./pythia stake status`

not_staked="Not currently staked"

echo $status

if [[ "$status" == "$not_staked" ]]; then
    ./pythia stake deposit
fi

# should remove in favor of sleeping before mine comand is ran
sleep $TRACKER_CYCLE

./pythia mine -r
