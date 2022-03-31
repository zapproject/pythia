./pythia approve 500000000 0x0165878a594ca255338adfa4d48449f69242eb8f
./pythia transfer 50000000 0x0165878a594ca255338adfa4d48449f69242eb8f

sleep 1

start powershell -NoExit -Mta  "./pythia dataserver"

sleep 1

./pythia --config=local_cfgs/config1.json stake deposit
./pythia --config=local_cfgs/config2.json stake deposit
./pythia --config=local_cfgs/config3.json stake deposit
./pythia --config=local_cfgs/config4.json stake deposit
./pythia --config=local_cfgs/config5.json stake deposit

sleep 1

nohup ./pythia --config=local_cfgs/config1.json mine -r > logs/1.log &
sleep 1
nohup ./pythia --config=local_cfgs/config2.json mine -r > logs/2.log &
sleep 1
nohup ./pythia --config=local_cfgs/config3.json mine -r > logs/3.log &
sleep 1
nohup ./pythia --config=local_cfgs/config4.json mine -r > logs/4.log &
sleep 1
nohup ./pythia --config=local_cfgs/config5.json mine -r > logs/5.log &
