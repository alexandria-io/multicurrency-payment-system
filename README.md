multi-currency-api
==================

Golang implementation of a multi-currency payment API


Specification
-------------

The multi-currency-API has a few key features:

* Handles requests from external sources.
* Serves data and sends HTTP responses when conditions are met.
* Records all completed transactions into a ledger for historical purposes.

Requests
---

Below is a list of all requests and how they are handled by the program.

#### quote

    POST
      currency      string      ex: "USD"
      amount        int         ex: 1405
     *r_currency    string      ex: "BTC"

* `currency` specifies the 
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


#### payment_address

    POST
      currency      string      ex: "BTC"
      amount        int         ex: 100000
     *timeout       int         ex: 600
     *callback      JSON object       
        method      string      ex: "HTTP_POST", "BLOCKCHAIN_WRITE"
                                
        params      JSON object
          HTTP_POST PARAMS:
          url       string      ex: "http://florincoin.info/mucua/callback/
          data      string      ex: "{'success':true,'sender':'florincoin.info'}"
          BLOCKCHAIN_WRITE PARAMS:
          data      string      ex: "Hello world! I love freedom of speech."
          binary    string      ex: "01001000" 

The payment_address api includes many options to serve callback data to your application. Specifically, the HTTP_POST params are programmable in a way that makes use of the application's connectivity to the network and cuts down on unnecessary API calls.

Below is the future specification for this API call in full.

***NOTE: the basic API v0.1 will have only the above examples***

***The following examples will be implemented in further versions:***

###### HTTP_POST PARAMS

    data      JSON object
      block   boolean 
      time    boolean
      hash    boolean


#### Response

The payment_address API 

