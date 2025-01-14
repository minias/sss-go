# Shamir secret sharing with Go

**[forked from dsprenkels/sss-go](https://github.com/dsprenkels/sss-go)**

[![Build Status](https://app.travis-ci.com/minias/sss-go.svg?branch=dev)](https://app.travis-ci.com/minias/sss-go)

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
# 64 byte character
./sss-go 0123456789abcdefghijklmnopqrstuvwxyz0123456789abcdefghijklmnopqr
#CreateShares 2 of 3 [226 byte]
#CreateShares[0] :
01f39ec54f7fc0557301ba99a05813d4be4a1b908ab3ae19da4b90256a602820255afb91313112417966f655521d974217b0f54d37596a1f53147fcf5519a201117bb4b05a2f0218e38688d79f0a4e217774080de8a2388624bf512d4a869842d05e39087af42f3d7f79d11450fb8aab62
#CreateShares[1] :
0254519244ad1b72f765d08cf8a75a1d3019bfde764da44e4d33c1d4152320ad515afb91313112417966f655521d974217b0f54d37596a1f53147fcf5519a201117bb4b05a2f0218e38688d79f0a4e217774080de8a2388624bf512d4a869842d05e39087af42f3d7f79d11450fb8aab62
#CreateShares[2] :
03c01456b4e3526f8bb0f67639f2945a4a282ae422eea28ac91b0772c9ebd1d67d5afb91313112417966f655521d974217b0f54d37596a1f53147fcf5519a201117bb4b05a2f0218e38688d79f0a4e217774080de8a2388624bf512d4a869842d05e39087af42f3d7f79d11450fb8aab62
#restored key:
0123456789abcdefghijklmnopqrstuvwxyz0123456789abcdefghijklmnopqr
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


