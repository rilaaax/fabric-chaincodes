package main

var samples = `
{
    "contractState": {
        "activeAssets": [
            "The ID of a managed asset. The resource focal point for a smart contract."
        ],
        "nickname": "TRADELANE",
        "version": "The version number of the current contract"
    },
    "event": {
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "carrier": "transport entity currently in possession of asset",
        "extension": {},
        "humidity": 123.456,
        "location": {
            "accuracy": 123.456,
            "latitude": 123.456,
            "longitude": 123.456,
            "measuredDateTime": 123.456
        },
        "temperature": 123.456,
        "timestamp": "2017-07-19T21:35:28.287983913+02:00"
    },
    "initEvent": {
        "nickname": "TRADELANE",
        "version": "The ID of a managed asset. The resource focal point for a smart contract."
    },
    "state": {
        "alerts": {
            "active": [
                "OVERTTEMP"
            ],
            "cleared": [
                "OVERTTEMP"
            ],
            "raised": [
                "OVERTTEMP"
            ]
        },
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "carrier": "transport entity currently in possession of asset",
        "compliant": true,
        "extension": {},
        "lastEvent": {
            "args": [
                "parameters to the function, usually args[0] is populated with a JSON encoded event object"
            ],
            "function": "function that created this state object",
            "redirectedFromFunction": "function that originally received the event"
        },
        "location": {
            "accuracy": 123.456,
            "latitude": 123.456,
            "longitude": 123.456,
            "measuredDateTime": 123.456
        },
        "temperature": 123.456,
        "timestamp": "2017-07-19T21:35:28.288073264+02:00",
        "txntimestamp": "Transaction timestamp matching that in the blockchain.",
        "txnuuid": "Transaction UUID matching that in the blockchain."
    }
}`