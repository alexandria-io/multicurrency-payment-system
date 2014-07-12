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

**quote**

    quote POST
      currency      string      ex: "USD"
      amount        float       ex: 14.05
     *r_currency    string      ex: "BTC"

The response is determined by whether or not `r_currency` is specified.

If `r_currency` is not specified, the response will be a list of key:value pairs for all available currencies. For example, if a `quote` request is received with the following parameters

    {
        "currency": "USD"
        "amount": 20.00
    }
    
The response will be something along these lines:

    {
       "BTC":0.03172
       "LTC":2.370
       "DOGE":7379000
       "FLO":30000
    }

However, if `r_currency` is specified, the request will look like this:

    {
        "currency": "USD"
        "amount": 20.00
        "r_currency": "FLO"
    }

The response will simply be a double value with the converted amount in the specified currency.

    { 30000 }


**payment_address**
