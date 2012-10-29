gotools
=======

Various tools for the go language, designed to pick up where the standard library let me down. One example is the json-rpc library, where net/rpc/jsonrpc only allowed to receive calls to methods named type.Method, where Method started with upper case.

The implementation found here allow for any string to be used as method, including spaces. The official json-rpc, indeed, does not mention any such restriction.

Also included is a implementation of USC, Universal Server Controller, a json-rpc interface to servers, allowing connection from a USC client that will retrieve the list of available commands, offer completion to the user, and send the commands to the server in a pre-parsed format.
