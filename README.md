# Golang-Tutorials

This is a simple web service that accepts the below mentioned JSON at a URL (`/date/update`) and updates the date according the JSON content (`intervals`).

JSON msg:

```
{
    "start_date": "2020-09-20T15:45:00Z",
    "intervals": [
        "20h",
        "140h25s",
        "40m"
    ]
}
```

The `start_date` will be updated based on the `intervals` content, and the web service should return the following results:

```
{
    "start_date": "2020-09-20T15:45:00Z",
    "intervals": [
        "20h",
        "140h25s",
        "40m"
    ],
    "results": [
        "2020-09-21T11:45:00Z",
        "2020-09-26T11:45:25Z",
        "2020-09-20T16:25:00Z"
    ]
}
```
