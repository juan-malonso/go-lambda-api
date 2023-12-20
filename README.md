# go-lambda-api

```sh
# AWS DEPLOYMENT
task up
task down

# CURL LAMDA_CONFIG
curl --location 'http://localhost:4566/restapis/<uuid>/test-app/_user_request_/<step_function>' \
--header 'Content-Type: text/plain' \
--header 'origin: http://localhost:4566' \
--data '{}'

# LAMBDA LOGS
task apps:lambdas:endpoint_config:logs

# LAMBDA TEST
task apps:lambdas:endpoint_config:dev
task apps:lambdas:endpoint_config:start
task apps:lambdas:endpoint_config:invoke
```


```mermaid
graph TD


    subgraph Leyenda[ ]
            direction RL
            
            ENDED:::ended
            DOING:::doing
    end

    ApiGateway[<h3>Api Gateway</h3>]:::ended
    
    ApiGateway--POST-->StepFunction1
    subgraph StepFunction1[<h3>Step Function</h3>]
        subgraph StepFunction1_container[ ]
            Lambda1[<h3>Endpoint Config</h3>Lambda 1]:::ended
            Lambda2[<h3>External Request</h3>Lambda 2]
            Lambda3[<h3>Event Store</h3>Lambda 3]:::doing
            Lambda4[<h3>Asset Store</h3>Lambda 4]
            Lambda5[<h3>Info Process</h3>Lambda 5]
        end
    end

    CloudWatch[<h3>Log Groups</h3>CloudWatch]:::ended

    EventBridge1[<h3>Events Bus</h3>Event Bridge]:::doing
    EventBridge2[<h3>Asset Bus</h3>Event Bridge]

    S31[<h3>Asset Bucket</h3>S3]

    DynamoDB2[<h3>Events Table</h3>Dynamo DB]:::ended
    DynamoDB1[<h3>Config Table</h3>Dynamo DB]:::ended
    
    Lambda1-.import.->Library2
    Lambda1-.import.->Library3

    Lambda2-.import.->Library2

    Lambda3-.import.->Library4
    Lambda3-.import.->Library2

    Lambda4-.import.->Library5
    Lambda4-.import.->Library2

    Lambda5-.import.->Library7
    Lambda5-.import.->Library6
    Lambda5-.import.->Library2

    subgraph Libraries[<h3>Libraries</h3>]
        subgraph Libraries_container[ ]
            Library4[<h3>Track Event</h3>]:::doing
            Library5[<h3>Track Asset</h3>]
            Library2[<h3>JSON Log</h3>]:::ended
            Library1[<h3>Env Var</h3>]:::ended
            Library3[<h3>Dynamo Config</h3>]:::ended
            Library6[<h3>Dynamo Event</h3>]:::ended
            Library7[<h3>S3 Asset</h3>]
        end
    end

    Library5-->EventBridge2
    EventBridge2-->AuxLambda2[<h3>S3 Store</h3> Lambda]
    AuxLambda2-.import.->Library7
    AuxLambda2-.import.->Library4

 
    Library4--->EventBridge1
    EventBridge1-->AuxLambda1[<h3>Dynamo Store</h3> Lambda]:::doing
    AuxLambda1-.import.->Library6

    Library2---->CloudWatch
    Library5-.import.->Library7
    Library5-.import.->Library1
    Library4-.import.->Library6
    Library4-.import.->Library1
    Library3-.import.->Library1
    Library3-->DynamoDB1
    Library6-->DynamoDB2
    Library6-.import.->Library1
    Library7-->S31
    Library7-.import.->Library1

    classDef doing fill:#cc7e22
    classDef ended fill:#057c46
```
