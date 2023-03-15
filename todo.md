# TODOs:

## Testing
I would like to have more test coverage on the API package that covers all the models, validation service, 
and API requests. These could be written in plain Go or use a framework like Ginkgo to help describe the tests a little better.

Ideally, we could use the `Client` interface to help us create mock structs and functions while testing, so we're 
not completely dependent on an actual API/internet response for testing.