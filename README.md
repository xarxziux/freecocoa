# Freecocoa - a FreeCiv combat calculator #

To start, run the command `make run` from the root folder.  This exposes an API on port 7123 with a single endpoint: `/api/v1/LTT/attack`.

Some sample queries:

```
> curl --request POST --url http://localhost:7123/api/v1/LTT/attack --header 'Content-Type: application/json' --data '{"attacker":{"name":"archers"},"defender":{"name":"warriors","terrain":{"type":"plains"}}}'
{"Attacker":{"AP":3,"HP":10,"FP":1},"Defender":{"DP":1,"HP":10,"FP":1}}
```

```
> curl --request POST --url http://localhost:7123/api/v1/LTT/attack --header 'Content-Type: application/json' --data '{"attacker":{"name":"howitzer","vetLevel":3,"hp":25},"defender":{"name":"destroyer","vetLevel":2,"hasCity":true,"city":{"size":12,"hasWalls":true,"hasCoastalDefence":true},"terrain":{"type":"hills","hasRiver":true}}}'
{"Attacker":{"AP":24,"HP":25,"FP":4},"Defender":{"DP":21,"HP":30,"FP":1}}
```
