# Overview

This is a command line util for generating the 6 digit One Time Pin code used commonly in 2 factor auth.

**Justification**
Two factor auth = what you know (password) + what you have (laptop/dev machine).


## History
I use to use a docker container for this but my Chromebook can't run docker (no Crostini support yet).

A golang binary is much faster in generating the OTP code and doesn't require a running instance of docker.

## Usage

(Assuming you've `go get` dependencies and built the damn binary)

`./otp-codegen sweetOTPSecret`
`123432`


## Recommendations
- Put the binary somewhere on your $PATH if you want to call it from any shell location.
- Use a script installed on your $PATH (e.g. a github one called `github-otp`) which contains your OTP secret and call the `otp-codegen` binary from that to spit out 6 digits to STDOUT.

I like the one liner `github-otp | pbcopy` on a Macbook which simplifies the login process.