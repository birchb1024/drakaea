drakaea
======================

Generate test credit cards for provided prefixes (BIN numbers).

Based on the go-business-creditcard package (github.com/dsparling/go-business-creditcard)

## Installation

```bash
$ git clone git@github.com:birchb1024/drakaea.git
$ cd drakaea
$ go build cmd/drakaea.go
$ ./drakaea
```

## Usage

`drakaea` expects BIN number prefixes on stdin one per line. The remainder of the 
card number os padded with '1's followed by the Luhn checksum
For example:

```shell
$ tail -3 somebins.csv
549123
552123
554123
$ tail -3 somebins.csv | ./drakaea
5491 2311 1111 1110	 MasterCard
5521 2311 1111 1111	 MasterCard
5541 2311 1111 1111	 MasterCard
```

## Drakaea (Wikipedia)

*Pouyannian mimicry*

Many plants have evolved to appear like other organisms, most commonly insects. This can have wide-ranging benefits including increasing pollination. 
In Pouyannian mimicry, flowers mimic a potential female mate visually, but the key stimuli are often chemical and tactile. The hammer orchid 
(Drakaea spp., an endangered genus of orchid that is native to Australia) is one of the most notable examples. The orchid has both 
visual and olfactory mimics of a female wasp to lure males to both deposit and pick up pollen.[13][better source needed]