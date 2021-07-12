./zap-miner approve 100000 0x5f3f1dBD7B74C6B46e8c44f98792A1dAf8d69154
./zap-miner approve 100000 0xb7278a61aa25c888815afc32ad3cc52ff24fe575
./zap-miner transfer 100000 0xb7278a61aa25c888815afc32ad3cc52ff24fe575

gnome-terminal -e "./zap-miner dataserver"

sleep 3

nohup ./zap-miner --config=local_cfgs/config6.json stake deposit > logs/6.log &
nohup ./zap-miner --config=local_cfgs/config7.json stake deposit > logs/7.log &
nohup ./zap-miner --config=local_cfgs/config8.json stake deposit > logs/8.log &
nohup ./zap-miner --config=local_cfgs/config9.json stake deposit > logs/9.log &
nohup ./zap-miner --config=local_cfgs/config10.json stake deposit > logs/10.log &

sleep 20

nohup ./zap-miner --config=local_cfgs/config6.json mine -r > logs/6.log &
sleep 1
nohup ./zap-miner --config=local_cfgs/config7.json mine -r > logs/7.log &
sleep 1
nohup ./zap-miner --config=local_cfgs/config8.json mine -r > logs/8.log &
sleep 1
nohup ./zap-miner --config=local_cfgs/config9.json mine -r > logs/9.log &
sleep 1
nohup ./zap-miner --config=local_cfgs/config10.json mine -r > logs/10.log &
