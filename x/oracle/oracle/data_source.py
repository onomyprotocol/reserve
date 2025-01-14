#!/usr/bin/env python3
from unittest import result
import requests
import json
import sys
import time

CONSTANTS = {
    "GATE_IO": {
        "SYMBOLS": {
            "NOM":"nom",
        },
        "URL": "https://api.gateio.ws/api/v4/spot/tickers"
    },

    "MEXC" : {
        "SYMBOLS": {
            "NOM":"nom",
        },
        "URL": "https://api.mexc.com/api/v3/ticker/price"

    }
}

    

def get_gate_price(symbols):
    if not symbols:
        return []

    # Set a header for api
    HEADER = {
        "Accept": "application/json",
    }

    PARAMETERS = {
        "currency_pair": "NOM_USDT",
    }

    # URL to retrieve the price of all tokens
    url = CONSTANTS["GATE_IO"]["URL"] 
    result = []
      # Set request parameters


    try:
        # Make api call
        response = requests.get(url, params=PARAMETERS, headers=HEADER, timeout=3)
        # Retrieve prices from response
        for idx, symbol in enumerate(symbols):
            if response.status_code == 200:
                result.append(float(response.json()[idx]["last"]))
            else:
                result.append(0)
        # Return prices
        return result
    except Exception as e:
        return [0 for i in range(len(symbols))]

def get_price_mecx(symbols):
    if not symbols:
        return []

    # Set a header for api
    HEADER = {
        "Accept": "application/json",
    }

    # URL to retrieve the price of all tokens
    url = CONSTANTS["MEXC"]["URL"]

    result = []
      # Set request parameters
    PARAMETERS = {
        "symbol": "FXUSDT",
    }


    try:
        # Make api call
        response = requests.get(url, params=PARAMETERS, headers=HEADER, timeout=3)

        # Retrieve prices from response
        for idx, symbol in enumerate(symbols):
            if response.status_code == 200:
                result.append(float(response.json()["price"]) )
            else:
                result.append(0)
        # Return prices
        return result
    except Exception as e:
        return [0 for i in range(len(symbols))]




def main(symbols):
    if len(symbols) == 0:
        return ""
    try:
        # Get price from bingx
        result_mecx = get_price_mecx(symbols)
        # Get price from gate
        result_gate = get_gate_price(symbols)

        # if lenghth of the results from all the sources is not same, then return 0
        if not len(result_gate) == len(result_mecx):
            return ",".join("0" for i in range(len(symbols)))

        result = []
        for item in zip(result_gate, result_mecx):
            different_sources_price = list(item)
            non_zero_sources = [price for price in different_sources_price if price != 0]
            if len(non_zero_sources) != 0:
                mean = sum(non_zero_sources) / len(non_zero_sources)
                result.append(mean)
            else:
                result.append(0)

        return ",".join(str(price) for price in result)
    except Exception as e:
        return ",".join("0" for i in range(len(symbols)))

if __name__ == "__main__":
    try:
        print(main(sys.argv[1:]))
    except Exception as e:
        print(e, file=sys.stderr)
        sys.exit(1)
