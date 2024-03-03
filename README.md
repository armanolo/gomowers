
### CHALLENGE

Application game to control robotic mowers.

Mower positions are represented by coordinates and cardinal 
compass points (N, E, S, W), and they receive commands in the form of 
letters (L, R, M) to control their movement. "L" and "R" turn the mower without moving, 
while "M" makes it move forward one grid point in its current direction. 
The mowers will navigate on the plateau efficiently based on commands sent.

```
Input Test Case #1:
55
12N
LMLMLMLMM
33E
MMRMMRMRRM
```

```
Output Test Case #2:
13N
51E
```

CLI exmple:
```
go run ./cmd/cli -i file -c ./example.coord 
go run ./cmd/cli -i text -c $'55\n12N\nLMLMLMLMM\n33E\nMMRMMRMRRM'
```


swag init -d "./cmd/api" -g openapi.go