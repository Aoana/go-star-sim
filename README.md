# go-star-sim
Star system simulation in Go

<img src="assets/images/screenshot.png" alt="Star Simulator Screenshot" width="100%" />

## Build

```
go install ./cmd/star-calc/
go install ./cmd/star-sim/
```
or
```
go build ./cmd/star-calc/
go build ./cmd/star-sim/
```

## Run

Example:
```
./star-calc -stars 1000 -steps 500 -file /tmp/simdata
./star-sim -file /tmp/simdata
./star-calc -model BarnesHutGR -stars 1000 -steps 500 -file /tmp/simdata2
./star-sim -file /tmp/simdata2
```

## Calculation time based on method

<img src="assets/images/calctime.png" alt="Star Simulator Screenshot" width="100%" />
