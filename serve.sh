#!/bin/sh
cd hardhat; npm i; ./start.sh & 
cd ..; sleep 20; ./release_build.sh; ./pythia dataserver

