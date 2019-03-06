![Logo!](assets/hilo-small-logo.png)

[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://img.shields.io/badge/repo%20status-WIP-yellow.svg?style=flat-square)](https://www.repostatus.org/#wip)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue?style=flat-square&logo=go)](https://godoc.org/github.com/cicizeo/hilo)
[![Go Report Card](https://goreportcard.com/badge/github.com/cicizeo/hilo?style=flat-square)](https://goreportcard.com/report/github.com/cicizeo/hilo)
[![Version](https://img.shields.io/github/tag/cicizeo/hilo.svg?style=flat-square)](https://github.com/cicizeo/hilo/releases/latest)
[![License: Apache-2.0](https://img.shields.io/github/license/cicizeo/hilo.svg?style=flat-square)](https://github.com/cicizeo/hilo/blob/main/LICENSE)
[![Lines Of Code](https://img.shields.io/tokei/lines/github/cicizeo/hilo?style=flat-square)](https://github.com/cicizeo/hilo)

> A Golang implementation of the Hilo network, a decentralized universal capital
facility in the Cosmos ecosystem.

Hilo is a Universal Capital Facility that can collateralize assets on one blockchain
towards borrowing assets on another blockchain. The platform specializes in
allowing staked assets from PoS blockchains to be used as collateral for borrowing
across blockchains. The platform uses a combination of algorithmically determined
interest rates based on market driven conditions. As a cross chain DeFi protocol,
Hilo will allow a multitude of decentralized debt products.

## Table of Contents

- [Dependencies](#dependencies)
- [Active Networks](#active-networks)
- [Install](#install)

## Dependencies

- [Go 1.16+](https://golang.org/dl/)
- [Cosmos SDK v0.43.0+](https://github.com/cosmos/cosmos-sdk/releases)
- [Starport](https://docs.starport.network/intro/install.html)

## Active Networks

Coming soon...

## Install

To install the `hilod` binary:

```shell
$ make install
```
