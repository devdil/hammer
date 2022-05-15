```
                          \`. 
.--------------.___________) \
|//////////////|___________[ ]
`--------------'           ) (
                           '-'

█░█ ▄▀█ █▀▄▀█ █▀▄▀█ █▀▀ █▀█
█▀█ █▀█ █░▀░█ █░▀░█ ██▄ █▀▄
```
## Hammer
### A simple load testing tool built in go

## Usage
For help
```shell
./hammer -h

Usage of ./hammer:
  -dur int
    	Duration in seconds dur=1 where is 1 is a valid duration in seconds. (default 1)
  -headers {'Content-Type' : 'application/json'
    	headers in string format. example: {'Content-Type' : 'application/json' (default "{}")
  -method string
    	If protocol used is http, possible options are get, put, update, delete. Default is get. (default "get")
  -prot string
    	Protocol used http, tcp or udp. (default "http")
  -rps int
    	Requests per second. Usage rps=1 where 1 is a positive integer. (default 1)
  -url string
    	Target url for load testing. (default "url")
```
Example
```shell
./hammer -dur 2 -method GET -prot http -rps 1 -url http://diljitpr.net
```
Output
```shell
                          \`.
.--------------.___________) \
|//////////////|___________[ ]
`--------------'           ) (
                           '-'

█░█ ▄▀█ █▀▄▀█ █▀▄▀█ █▀▀ █▀█
█▀█ █▀█ █░▀░█ █░▀░█ ██▄ █▀▄
Hammer Load Test Tool
---------------------------------
-----Parameters------------------
Requests Per Second:  1
Duration:  2
---------------------------------
Hammering ...

---------------------------------
Summary Statistics
---------------------------------
P99(ms) :  259
P95(ms) :  259
P90(ms) :  259
Max(ms) :  418
Min(ms) :  100
---------------------------------
```