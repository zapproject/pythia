#!/bin/sh
cd hardhat; npm i; ./start.sh & 
cd ..; sleep 20; ./release_build.sh; ./zap-miner dataserver

