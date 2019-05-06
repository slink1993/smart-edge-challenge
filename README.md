# Smart Edge Code Challenge
## Meets the following requirements:
- [x] Given a string input of up to 250 characters, return a JSON response compliant to the schema defined below.
- [x] You are responsible for generating a public/private RSA or ECDSA keypair and persisting the keypair on the filesystem
- [x] Subsequent invocations of your application should read from the same files
- [x] Document your code, at a minimum defining parameter types and return values for any public methods
- [x] Include Unit Test(s) with instructions on how a Continuous Integration system can execute your test(s)
- [x] You may only use first order libraries, you may not use any third party libraries or packages


## Testing:

`docker build -f test/Dockerfile . -t test`

## To run the app:

1. `docker build . -t codechal`
2. `docker run codechal`

