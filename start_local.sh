./pythia approve 100000 0x5f3f1dBD7B74C6B46e8c44f98792A1dAf8d69154
./pythia approve 100000 0xb7278a61aa25c888815afc32ad3cc52ff24fe575
./pythia transfer 100000 0xb7278a61aa25c888815afc32ad3cc52ff24fe575

# gnome-terminal -e "./pythia dataserver"

# sleep 3

./pythia stake deposit
./pythia --config=local_cfgs/config2.json stake deposit
./pythia --config=local_cfgs/config3.json stake deposit
./pythia --config=local_cfgs/config4.json stake deposit
./pythia --config=local_cfgs/config5.json stake deposit

<<com
./pythia mine -r
./pythia --config=local_cfgs/config2.json mine -r
./pythia --config=local_cfgs/config3.json mine -r
./pythia --config=local_cfgs/config4.json mine -r
./pythia --config=local_cfgs/config5.json mine -r
com


# nohup ./pythia mine -r > logs/1.log &
# nohup ./pythia --config=local_cfgs/config2.json mine -r > logs/2.log &
# nohup ./pythia --config=local_cfgs/config3.json mine -r > logs/3.log &
# nohup ./pythia --config=local_cfgs/config4.json mine -r > logs/4.log &
# nohup ./pythia --config=local_cfgs/config5.json mine -r > logs/5.log &
