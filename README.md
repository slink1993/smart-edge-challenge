# Smart Edge Code Challenge [![CircleCI](https://circleci.com/gh/stephenlink1993/smart-edge-challenge.svg?style=svg)](https://circleci.com/gh/stephenlink1993/smart-edge-challenge)

## Meets the following requirements:
- [x] Given a string input of up to 250 characters, return a JSON response compliant to the schema defined below.
- [x] You are responsible for generating a public/private RSA or ECDSA keypair and persisting the keypair on the filesystem
- [x] Subsequent invocations of your application should read from the same files
- [x] Document your code, at a minimum defining parameter types and return values for any public methods
- [x] Include Unit Test(s) with instructions on how a Continuous Integration system can execute your test(s)
- [x] You may only use first order libraries, you may not use any third party libraries or packages


## Testing:

### Locally:

#### Docker:

1. `docker build -f test/Dockerfile . -t testcodechal`
2. `docker run testcodechal`

#### On Local Machine:

1. `go test ./test -v`

### CI/CD (CircleCI):

The fastest way to trigger testing via CI/CD is to fork this repository, then connect it to your personal CircleCI account. You may also clone this repository then push it your personal github. It will also need to be connected to your personal github. 

To trigger a build, you can do a `git commit`. 
If there are no changes to commit, you may use the command:

` git commit --allow-empty -m "Trigger CircleCI testing"`

Then:

 `git push`

And that's it!

## To run the app:

1. `docker build . -t codechal`
2. `docker run codechal your@email.com`

Which gives the expected output:
```
{
  "message": "your@email.com",
  "signature": "au78sefVbtC4Mmzk2jMhqWB/DNMZfkpBfPRD/0tvZcxncmPHdoX8ImZMPw1Yw/uTPiv8+KA7oCDXJ2fY6Vf6YivmYxs4anuYL5MswwT/D9za55BuOHWiA6JBFOS2Y8XNinuZ9L3wvYewNalGO9/gWgQ6CbG8nI5ABF68ZqHIOXg2efJI8C4R5X+jsJpm3cm8150qG/HU55FZkTVYUCwmnC2Kw1bfbiLFUoSBusigBytNskHD6Khix6XnqOsM7rYggey07TkamWv6z9zwioR9w7FZJbZ/Pbp3Fo/0VRs7IdgxBV7I6kjFKI1h356dfx/fPJg8A5LiLy62z3MP+f6Xxg==",
  "pubkey": "-----BEGIN PUBLIC KEY-----\nMIIBCgKCAQEAsle1eogQgT/D2axTV+rLu8lRoQxZiyNCdBtgO8sTP+tUMsxaXca3\nxNJpgbBwqESNZrAEsZxLWJRLNvfomkm5WhxVsG1H7gsWfyDBS85Qm27LOQKHBWry\nuUhiETwC+lSmMO5WbJz0As/jaywHOEaDbQv7QbNl3rkJnObjOeYh3jM5BLMBy++e\nvCBJP6rLcKOrvo0AmG82ilmh7rWE33YeGm1TtaNFfh/p2sUxaKnXA7KHjiOwZikZ\n0ZKd/oaCUarZq4BJ3VVXWHeYUg6A6LF0dvjzDVrejC0lLkikTH1DnipVxuCIO5I2\njqzOpnBm4CYHXIhlm4U/bCuhp8HoAtBJ3QIDAQAB\n-----END PUBLIC KEY-----\n"
}
```
