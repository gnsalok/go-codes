# Testing 

## Test Coverage

#### To the coverage details
`go test -cover`

#### To get info of cover statements
`go test -coverprofile cover.out`  

`go tool cover -func cover.out`

#### If need more info

`go tool cover -html cover.out`


#### To get the high and low count in coverage report
`go test -coverprofile count.out -covermode count` 
`go tool cover -html count.out`



## Benchmarking and Profiling

#### Run the benchmaking test
`go test -bench .`

#### Run benchmark test, targeting the specified time
`go test -bench . -benchtime 10s`

>Note : By default benchmark have 1s benchtime set. If you have scenario where it needs to test long running process like end-to-end testing, database connection etc.. increasing the time will give the good benchmark result.
Here, 10s beanchtime meaning, average time taken by per operation for no. of iteration run within 10s time.

### Profiling Test

##### run the test name alloc and generate profile output with help for memprofile
`go test -bench Alloc -memprofile  profile.out`

##### Analyse output file with pprof tool 
> Note: it will open interactive window, choose _svg_ tool to see the graph 

`go tool pprof profile.out`    

