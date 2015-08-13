
     _   _ _   _ __  __  ___ _____ 
    | | | | | | |  \/  |/ _ \_   _|
    | |_| | | | | |\/| | | | || |  
    |  _  | |_| | |  | | |_| || |  
    |_| |_|\___/|_|  |_|\___/ |_|  

*[hew-moe] hypersonic http load tester*

## Usage ##

    humot <url> [<concurrency>]

 - *url*: URL to humot
 - *concurrency*: number of concurrent requests. default is 250
 - Prints stats every 2 seconds
 - Lists status codes received and total number of responses for each
 - Continues until you press Control-C

## Example ##

    $ humot http://10.70.6.83:31042 123
    Hitting http://10.70.6.83:31042 with 123 concurrent requests ... (^C to stop)
    2.004234597s | 200: 1281 | Total:  1281
    4.004963043s | 200: 2684 | Total:  2684
    6.00398053s | 200: 4024 | Total:  4024
    8.004568408s | 200: 5021 | Total:  5021
    10.004977115s | 200: 6332 | Total:  6332
    12.00363762s | 200: 7939 | Total:  7939
    14.004964808s | 200: 9693 | Total:  9693
    16.002970143s | 200: 11023 | Total:  11023
    18.00510027s | 200: 12641 | Total:  12641
    20.000053881s | 200: 13335 | Total:  13335
    22.0050679s | 200: 13949 | Total:  13949
    24.004964016s | 200: 14634 | Total:  14634
    26.00488832s | 200: 14900 | Total:  14900
    28.000276302s | 200: 14994 | Error(i/o timeout): 125 | Total:  15119
    30.005023182s | 200: 15853 | Error(i/o timeout): 166 | Total:  16019
    32.005028904s | 200: 16322 | Error(i/o timeout): 166 | Total:  16488

## Build from source ##

    go build humot.go

## tl;dr ##

    curl -O https://github.com/fotoetienne/humot/releases/download/v0.1/humot ; chmod +x humot
    ./humot http://myawesomewebsite.me 200000


