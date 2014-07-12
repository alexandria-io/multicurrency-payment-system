multi-currency-api
==================

Golang implementation of a multi-currency payment API


Specification
-------------

The multi-currency-API has a few key features:

* Handles requests from external sources.
* Serves data and sends HTTP responses when conditions are met.
* Records all completed transactions into a ledger for historical purposes.

The following is an explination of how each request is structured:

`
quote POST
    currency      string      ex: "USD"
    amount        float       ex: 14.05
   *r_currency    string      ex: "BTC"
`
The response is   
    
