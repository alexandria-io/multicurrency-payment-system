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

--------

#### quote request

    POST
      currency      string      ex: "USD"
      amount        int         ex: 1405
     *r_currency    string      ex: "BTC"

* `currency` specifies the base currency, `r_currency` is the currency pairing we are interested in.

* `amount` specifies the base currency amount that will be compared to the list of currencies in the response.

* `r_currency` specifies a single currency the requstor wants (to reduce HTTP traffic and simplify code).

Future versions will allow multiple `r_currency` options.

#### quote response

    currency:amount      string:int      ex: "BTC":3172000

The response is determined by whether or not `r_currency` is specified. In general, the API responds with the currency conversion rates the requestor is interested in.

#### quote examples

If `r_currency` is not specified, the response will be a list of key:value pairs for all available currencies. For example, if a `quote` request is received with the following parameters

*quote request:* 

    {
        "currency": "USD",
        "amount": 20.00
    }
    
*quote response:*

    {
       "BTC":3172000,
       "LTC":2370000000,
       "DOGE":737900000000000,
       "FLO":3000000000000
    }

However, if `r_currency` is specified, the request will look like this:

*quote request:*

    {
        "currency": "USD"
        "amount": 20.00
        "r_currency": "FLO"
    }

The response will simply be a double value with the converted amount in the specified currency.

*quote response:*

    { 30000 }


#### payment request

    POST
      currency      string      ex: "BTC"
      amount        int         ex: 100000
     *timeout       int         ex: 59000, 1405230924
     *callback      JSON object       
        method      string      ex: "HTTP_POST", "BLOCKCHAIN_WRITE"
                                
        params      JSON object
          HTTP_POST PARAMS:
          url       string      ex: "http://florincoin.info/mucua/callback/
          data      string      ex: "{'success':true,'sender':'florincoin.info'}"
          BLOCKCHAIN_WRITE PARAMS:
          data      string      ex: "Hello world! I love freedom of speech."
          binary    string      ex: "01001000" 

* `currency` defines which currency the requestor wants to pay in.

* `amount` defines the amount of satoshis of that currency that marks this payment as "complete". If not specified, there is no minimum, and even 1 satoshi will mark the payment as "complete".

* `timeout` defines the amount of time in which this payment will expire. It is always an int, but behaves differently given different inputs. When given a block number as input, it will timeout when that block is reached. When given a timestamp, it will timeout when that timestamp is reached.

* `callback` is a JSON object containing information about the callback requested via the payment API's callback service. It is explained in detail in the example section below.

The payment api includes many options to serve callback data to your application. Specifically, the HTTP_POST params are programmable in a way that makes use of the application's connectivity to the network and cuts down on unnecessary API calls.

Below is the future specification for this API call in full.

*NOTE: the basic API v0.1 will have only the above examples. The following examples will be implemented in future versions:*

###### HTTP_POST PARAMS

    data            JSON object
      block         boolean 
      time          boolean
      hash          boolean
      min_confirms  int
      confirms      int
      custom        JSON object


#### payment response 

    address      string      ex: "17qfT3hssK5mx7km7QtuogiXeka9Spo1VK"
    timeout      int         ex: 10000

The payment API responds based upon on the input of the received request. It will always, at the very least, respond with a pamynet address for the currency specified.

#### payment examples

If no `callback` is specified, the API will remain silent and ping nothing when the payment is complete. Requestors will have no choice but to use the `status` API call to examine their payment.

If `callback` is specified, the payment API (which runs as a service) will begin serving callbacks when certain conditions are met. Here is a very general example that would be seen in the wild:

*payment request:*

    {
        "currency":"BTC",

