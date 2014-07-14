multicurrency payment system
============================

Introduction
------------

The multicurrency payment system enables payments in multiple crypto-currencies by providing a backend structure to organize payment addresses. Payment listeners are created to constantly scan the blockchain and fire off callbacks when certain criteria is met. The multicurrency payment system facilitates easy access to cross-blockchain functionality. It is written in golang.

Key Features
------------

* Provides multicurrency exchange rates based on specified exchange APIs.
* Handles requests for payment addresses.
* Creates a payment listener to fire off callbacks when conditions are met.
* Forwards funds to a specified address (only for homogeneous currency requests).
* Records all completed transactions into a ledger for historical purposes.

Installation
------------

Installing multicurrency payment system:

    go get github.com/blocktech/multicurrency_payment_system

Configuring URL endpoints can be done by modifying `config.go`.

    quote_endpoint=/quote/
    payment_endpoint=/payment/
    status_endpoint=/status/


Communication
-------------

All communication is done via HTTP POST with the POST body in JSON format. These are the necessary HTTP headers to communicate with the multicurrency payment system:

    POST 
    Accept: application/json
    Content-Type: application/json

Below you can find descriptions of all requests and how they are handled by the program. The responses are also described in detail. They vary depending on the input received in the initial request.

Every parameter that is marked with an `*` is optional.

---

### quote

#### quote request

     currency      string      ex: "USD"
     amount        int         ex: 1405
    *convert       string      ex: "BTC"

* `currency` specifies the base currency.
* `amount` specifies the base currency amount that will be compared to the list of currencies in the response.
* `convert` specifies a single currency the requestor wants to know about (this option reduces HTTP traffic and simplifies code).

Future versions will allow multiple `convert` options.

#### quote response

    currency:amount      string:int      ex: "BTC":3172000

The response is determined by whether or not `convert` is specified. In general, the system responds with the currency conversion rates the requestor is interested in.

#### quote examples

If `convert` is not specified, the response will be a list of key:value pairs for all available currencies. Here is an example of a quote request and response without `convert` specified:

*quote request:* 

    {
        "currency": "USD",
        "amount": 2000
    }

*quote response:*

    {
       "BTC":3172000,
       "LTC":2370000000,
       "DOGE":737900000000000,
       "FLO":3000000000000
    }

However, if `convert` is specified, the request will look like this:

*quote request:*

    {
        "currency": "USD"
        "amount": 2000
        "convert": "FLO"
    }

The response will simply be a double value with the converted amount in the specified currency.

*quote response:*

    { 
        3000000000000
    }

---

### payment

#### payment request

     currency        string        ex: "BTC"
    *convert         string        ex: "FLO"
     amount          int           ex: 100000
     min_confirms    int           ex: 5
     timeout         int           ex: 59000, 1405230924
    *fee_quote       boolean       ex: true
    *forward_to      string        ex: "1AYvYfub9BsLDSF9CqShphKD23VUvvL6Cm"

    *callback        JSON object       
       method        string        ex: "HTTP_POST", "BLOCKCHAIN_WRITE"
       max_confirms  int           ex: 20
       params        JSON object
         HTTP_POST PARAMS:         See below.
         BLOCKCHAIN_WRITE PARAMS:  See below.

* `currency` defines which currency the requestor wants to pay in.
* `convert`, when set to a known currency, will convert `amount` to the currency set in `currency`.
* `amount` defines the amount of satoshis of that currency that marks this payment as "complete". If not specified, there is no minimum, and even 1 satoshi will mark the payment as "complete".
* `min_confirms` is the minimum amount of confirms needed to consider the payment complete.
* `fee_quote`, when set to `true`, will prevent this program from creating a new address or payment listener. The payment response will be the cost (in satoshis of `currency`) of setting up that listener.
* `forward_to` if specified, the exact amount paid (minus fees) will be forwarded to the address specified.
* `timeout` defines the amount of time in which this payment listener will expire. It is always an int, but behaves differently given different inputs. When given a block number as input, it will timeout when that block is reached. When given a timestamp, it will timeout when that timestamp is reached. `timeout` cannot be zero.
* `callback` is a JSON object containing information about the callback requested. It defines the requestor's constraints for the payment listener. The payment listener serves data determined by the requestor's JSON parameters. This is the payment system's callback service. It is explained in detail in the examples section below.
* `callback->method` determines the method of callback. Currently supporting HTTP_POST and BLOCKCHAIN_WRITE callbacks.
* `callback->max_confirms` is the maximum amount of confirms that will fire off a callback. After `max_confirms` has passed, no more callbacks will be sent.
* `callback->params` define the callback service.

A successful payment request will create a payment listener. Whenever a payment listener detects certain conditions, a callback is fired off. Callbacks are fired off when both `amount` and `min_confirms` are reached. At the moment, `HTTP_POST` and `BLOCKCHAIN_WRITE` are the two options for callbacks served by the multi-currency-api.

###### HTTP\_POST PARAMS

HTTP\_POST callbacks are programmable. This makes use of the application's connectivity to the network and cuts down on unnecessary HTTP traffic.

    url             string      ex: "http://florincoin.info/mucua/callback/
    tx_notify       boolean     ex: false
    data            JSON object
      block         boolean     ex: true
      time          boolean     ex: true
      hash          boolean     ex: true
      confirms      boolean     ex: true
      custom        JSON object

* `url` is the URL that will be served with callback POST data. *Note: when a callback fails, it will retry every 2, 4, 8, ... etc seconds, growing exponentially.*
* `tx_notify` is a boolean which determines whether or not a callback will be sent when a new transaction paying this address is detected.
* `data`: All `boolean` data values will assure a response from the system that contains the data requested. For example, setting `block` to `true` will cause the system to respond with a callback containing the block number.
* `custom` can be filled with whatever static JSON the requestor determines. It will be served to the callback endpoint URL specified in `url`.

###### BLOCKCHAIN\_WRITE PARAMS

Callbacks are not limited to HTTP\_POST. You can request writing data to the blockchain instead.

    data      string      ex: "Hello world! I love freedom of speech."
    binary    string      ex: "01001000" 

* `data` specifies utf-8 data stored to the Florincoin blockchain.
* `binary` specifies a string of binary code to be stored in the Florincoin blockchain.

#### payment response 

     id           int         ex: 11022047
     address      string      ex: "17qfT3hssK5mx7km7QtuogiXeka9Spo1VK"
     currency     string      ex: "BTC"
    *amount       int         ex: 3050
    *fees         int         ex: 50
     timeout      int         ex: 10000

The payment system responds based upon on the inputs given in the payment request. The system will always, at the very least, respond with an id and payment address for the currency specified as well as the amount that must be paid.

* `id` is a unique identifier that must be passed to the status request. It is a unique identifier that identifies a payment listener and its associated address.
* `address` is a crypto-currency address that payments will be made to. The payment listener associated with this address is created when this response is sent.
* `currency` is the currency the address is associated with, and is the currency payments must be made in to that address.
* `amount` will be an integer representation of the minimum amount of satoshis that must be spent in `currency` before the payment is considered complete.
* `fees` is the number of satoshis paid in fees.
* `timeout` is the number of seconds the listener will stay alive.

#### payment examples

If no `callback` is specified, the system will remain silent and ping nothing when the payment is complete. Requestors will have no choice but to make use of status requests to examine their payment.

If `callback` is specified, the payment system (which runs as a service) will begin serving callbacks when certain conditions are met. Here is a very general example that would be seen in the wild:

*payment request:*

    {
        "currency":"BTC",
        "convert":"FLO",
        "amount":2000000,
        "min_confirms":30
        "fee_quote":true,
        "timeout":59300,
        "callback": {
            "method":"HTTP_POST",
            "max_confirms":40,
            "params": {
                "url":"http://florincoin.info/mucua/callback",
                "tx_notify":true,
                "data": {
                    "block":true,
                    "time":true,
                    "hash":true
                    "confirms":true,
                    "custom":"{'local_id':'BZ99ML7'}"
                }
            }
        }
    }

*payment response:*

    {
       "id":11022047,
       "address":"17qfT3hssK5mx7km7QtuogiXeka9Spo1VK",
       "currency":"BTC",
       "amount":3050,
       "fees":50,
       "timeout":10000
    }

In the above example, a request is sent to get a Bitcoin payment address to watch for a payment equal to 2000000 satoshis of FLO. The response shows that 3050 satoshis of bitcoin are necessary to cover the cost of 2000000 satoshis of FLO. There is a minimum number of confirms set, and a fee quote is requested. The fee quote in this case is 50 satoshis of bitcoin.

---

### status

#### status request

     id             string      ex: "177-134-559"
     address        string      ex: "17qfT3hssK5mx7km7QtuogiXeka9Spo1VK"
    *callback_info  boolean     ex: true

* `id` is the unique identifier given to the requestor when they received a response from the payment system.
* `address` specifies the address the requestor wants information about. Each payment listener is identified by the id number and address created for it.
* `callback_info` specifies whether or not the requestor wants callback info in the response.

#### status response

    mempool        boolean        ex: true
    received       int            ex: 2000
    confirms       int            ex: 12
    confirmed      boolean        ex: true
    transactions   JSON object    ex: <List of transactions in JSON format>
    callback       JSON object    ex: "{'confirmed':true,'confirms':12}"

The response will contain information about the payment listener binded to the given payment address.

* `mempool` is true or false if the transaction has been listed in the mempool. This will be set to `true` automatically when `confirms` is greater than zero, regardless of whether or not the transaction is actually seen in the mempool.
* `received` is the amount of coins received on the address so far.
* `confirms` is an integer representation of the number of confirms since `received` has reached `min_amount`.
* `confirmed` will be set to `true` only when the address has accumulated enough `received` coins to satisfy `min_amount` and all transactions in the `transactions` list have above `min_confirms` confirmations.
* `transactions` is a JSON object containing a list of txids associated with this address along with the amount of coins sent to the payment address in each transaction..
* `callback` is a JSON object that has all relevant callback information.

#### status example

A simple example of a request and response:

*status request:*

    {
        "address":"17qfT3hssK5mx7km7QtuogiXeka9Spo1VK",
        "callback_info":true
    }

*status response:*

    {
        "mempool":true,
        "confirms":12,
        "confirmed":true,
        "transactions": {
            "28afb4933555aa6f1616abf7009af47c796713d50d144803cff87aa7c8ebaa47":30000,
            "1df7ed637d64cf1282468198c6ae7f988f5ec8b88ee85623954528dd3710f311":40000,
            "74110c887b09fc910e2f55772e2da7a6ac3f10d4157ca737725cd6e32d8ab75b":50000,
        }
        "callback": {
            "method":"HTTP_POST",
            "url":"http://florincoin.info/mucu/callback/",
            "callbacks_sent":6,
            "callback_history": {
                1: {
                    "success":true,
                    "time":1405292062,
                    "block":59000
                },
                /* etc... continues for all callbacks */
            }
        }
    }

This response is a JSON object with detailed information about the status of the payment request (identified by the address given in the original payment response) containing relevant info about the callback history.
