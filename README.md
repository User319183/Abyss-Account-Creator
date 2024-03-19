# Abyss Account Creator

This is a Go program that creates user accounts by sending HTTP POST requests to https://abyssdigital.xyz/. It uses goroutines for concurrent execution and a proxy for the requests.

## Features

- Concurrent account creation
- Proxy support
- Randomized user data
- Error logging
- Response logging
- Account details storage

## How it works

The program starts by setting a seed for the random number generator. It then defines the URL for the account creation requests and the proxy to use.

A goroutine is started for each account to be created. Each goroutine generates a random user ID, constructs the payload for the request, sends the request, and logs any errors.

The payload for each request includes a randomly generated email and display name, along with a predefined password, photo URL, and Discord user.

The `sendRequest` function is responsible for sending the request. It marshals the payload into JSON, creates a new HTTP POST request with the JSON payload, sets the request header, parses the proxy URL, and sends the request using a client that uses the proxy.

The response from the server is read and logged, and the account details are stored in a file.

## Usage

To use this program, simply run the `main.go` file. The number of accounts to create and the level of concurrency can be adjusted by changing the values of `i` in the for loop and `concurrency`, respectively.

## Note

This program was used to fuck up the owner's firebase database with hundreds to thousands of accounts. This program was also skidded and stolen, so here is the source code if you never had access to it.