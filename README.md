# Shamir secret sharing with Go

**[forked from dsprenkels/sss-go](https://github.com/dsprenkels/sss-go)**

[![Build Status](https://api.travis-ci.com/minias/sss-go.svg?branch=dev)](https://travis-ci.org/minias/sss-go)

`sss-go` contains Go bindings for my [Shamir secret sharing library][sss].
This library allows users to split secret data into a number of different
shares. With the possession of some or all of these shares, the original secret
can be restored.

An example use case is a beer brewery which has a vault which contains their
precious super secret recipe. The 5 board members of this brewery do not trust
all the others well enough that they won't secretly break into the vault and
sell the recipe to a competitor. So they split the code into 5 shares, and
allow 4 shares to restore the original code. Now they are sure that the
majority of the staff will know when the vault is opened, but they can still
open the vault when one of the staff members is abroad or sick at home.

## Macos Required Installation Requirements

```shell
brew install coreutils
```

## Installation

```shell
git clone github.com/minias/sss-go
go build
```

## Usage

```shell
./sss-go
```

## Changelog

### Version 0.1.1

- Remove an unintended side channel which allows a participating attacker with
  access to a accurate timing channel to iteratively guess shares during the
  execution of `combine_shares`.

## Questions

Feel free to send me an email on my Github associated e-mail address.

[randombytes]: https://github.com/dsprenkels/randombytes
[sss]: https://github.com/dsprenkels/sss


