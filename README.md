
# alertmanager-docker-compose : alertmanager setup for local tests and exploration of features

## How-to
1. cd docker-compose
2. execute `docker-compose up -d`
3. execute `docker ps -a` to see if all containers are up
4. execute `docker logs -f receiver` to see the results
5. you can also open browser and go to `localhost:9093` to loging into alertmanager
6. execute following to post variuos alerts to alertmanager from your terminal or postman.


Note that, alert might appear with delay (upto 1 minute) on the logs terminal.

to test if webhook server is working correctly or not :

```
curl --location 'http://127.0.0.1:8888' \
--header 'Content-Type: application/json' \
--data '[
    {
        "labels": {
            "alertname": "alert_test_team1",
            "team": "team1",
            "container": "vector",
            "severity": "warning"
        },
        "endsAt": "2029-04-19T16:00:00+00:00"
    }
]'
```

to test vector alert :

```
curl --location 'http://127.0.0.1:9093/api/v1/alerts' \
--header 'Content-Type: application/json' \
--data '[
    {
        "labels": {
            "alertname": "alert_test_team1_vector",
            "team": "team1",
            "container": "vector"
        },
        "endsAt": "2029-04-19T16:00:00+00:00"
    }
]'
```

to test kafka alerts :

```
curl --location 'http://127.0.0.1:9093/api/v1/alerts' \
--header 'Content-Type: application/json' \
--data '[
    {
        "labels": {
            "alertname": "alert_test_team1_kafka",
            "team": "team1",
            "container": "kafka"
        },
        "endsAt": "2029-04-19T16:00:00+00:00"
    }
]'
```

to test team1 only alert :

```
curl --location 'http://127.0.0.1:9093/api/v1/alerts' \
--header 'Content-Type: application/json' \
--data '[
    {
        "labels": {
            "alertname": "alert_test_team1_random",
            "team": "team1"
        },
        "endsAt": "2029-04-20T16:00:00+00:00"
    }
]'
```

to test critical alert :
```
curl --location 'http://127.0.0.1:9093/api/v1/alerts' \
--header 'Content-Type: application/json' \
--data '[
    {
        "labels": {
            "alertname": "alert_test_team1_critical",
            "team": "team1",
            "severity": "critical"
        },
        "endsAt": "2029-04-19T16:00:00+00:00"
    }
]'
```

to test minor alert (should not be displayed) :
```
curl --location 'http://127.0.0.1:9093/api/v1/alerts' \
--header 'Content-Type: application/json' \
--data '[
    {
        "labels": {
            "alertname": "alert_test_team1_minor",
            "team": "team1",
            "container": "vector",
            "severity": "minor"
        },
        "endsAt": "2029-04-19T16:00:00+00:00"
    }
]'
```


to resolve alerts, execute same above coomands and set the endsAt to nearest current timestamp.
