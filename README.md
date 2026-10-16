# Titan

[![Build Status](https://travis-ci.org/distributedio/titan.svg?branch=master)](https://travis-ci.org/distributedio/titan)
[![Go Report Card](https://goreportcard.com/badge/github.com/distributedio/titan)](https://goreportcard.com/report/github.com/distributedio/titan)
[![Coverage Status](https://coveralls.io/repos/github/distributedio/titan/badge.svg?branch=master)](https://coveralls.io/github/distributedio/titan?branch=master)
[![Coverage Status](https://img.shields.io/badge/version-v0.3.1-brightgreen.svg)](https://github.com/distributedio/titan/releases)
[![Discourse status](https://img.shields.io/discourse/https/meta.discourse.org/status.svg)](https://titan-tech-group.slack.com)

A distributed implementation of __Redis compatible layer__  based on [TiKV](https://github.com/tikv/tikv/)

## Why Titan?

* Completely compatible with Redis protocol
* Full distributed transaction with strong consistency
* Multi-tenancy support
* No painful scale out
* High availability

Thanks [TiKV](https://github.com/tikv/tikv/) for supporting the core features. The project is developed and open sourced by the Beijing Infrastructure Team at [Meitu](https://www.meitu.com/) and has been donated to [DistributedIO](https://github.com/distributedio) org.

## Architecture

![titan](docs/titan.jpeg)

## Quick start

Can't wait to experiment with Titan? Just follow 2 steps:

1. `curl -s -O https://raw.githubusercontent.com/distributedio/titan/master/docker-compose.yml`
2. `docker-compose up`

Then connect to Titan using `redis-cli`

```
redis-cli -p 7369
```

___Enjoy!___

## Installation

### SetUp TiKV cluster

Titan works with 2 TiDB components:

* TiKV
* PD

To setup TiKV and PD, please follow the official [instructions](https://pingcap.com/docs-cn/dev/how-to/deploy/orchestrated/ansible/)

### Run Titan

* Build the binary

```
go get github.com/distributedio/titan
cd $GOPATH/src/github.com/distributedio/titan
make
```

* Edit the configration file

```
pd-addrs="tikv://your-pd-addrs:port"
```

* Run Titan

```
./titan
```

For more details about [Deploy Titan](docs/ops/deploy.md), click here.

## Commands supporting status

See the details of the commands [supporting status](docs/command_list.md)

| command      | status                  |
| ------------ | ----------------------- |
| Connections  | Almost Fully Supported  |
| Transactions | Supported               |
| Server       | Almost Fully Supported  |
| Keys         | Supported               |
| Strings      | Almost Fully Supported  |
| List         | Almost Fully Supported  |
| Hashes       | Supported               |
| Sets         | Almost Fully Supported  |
| Sorted Sets  | Almost Fully Supported  |
| Geo          | Not Supported Yet       |
| Hyperloglog  | Not Supported Yet       |
| Pub/Sub      | Not Supported Yet       |
| Scripting    | Not Supported Yet       |
| Streams      | Not Supported Yet       |

## Benchmarks

Refer to the [benchmark docs](https://pan.baidu.com/s/1m5yp5LsvFjsDKvHtaXwWvg) for more details. It's shared on Baidu Disks, use the code `hzt6` to gain the permission.

Basic benchmarking result.

### Get

![Get command benchmark](docs/benchmark/get-benchmark.png)

### Set

![Set command benchmark](docs/benchmark/set-benchmark.png)

For more info, please vist here [Titan Benchmarks](docs/benchmark/benchmark.md)

## FAQ

[FAQ](https://github.com/distributedio/titan/issues?utf8=%E2%9C%93&q=+label%3A%22good+first+issue%22)

## Roadmap

View our [Roadmap](https://github.com/distributedio/titan/projects)

## Release Note
* 20.4.21: add support for rpop and rpoplpush
Hello World 2

---

## 🚲 Bicycle Advocacy 🚲

We believe that cycling is one of the most powerful tools we have for building healthier communities, reducing carbon emissions, and reclaiming our streets.

### Why Bicycles Matter

- **Climate action** — A bicycle produces zero direct emissions. Replacing even one car trip per day with a bike ride can save hundreds of kilograms of CO₂ per year.
- **Healthier cities** — Cycling reduces traffic congestion, improves air quality, and lowers noise pollution, making urban spaces more liveable for everyone.
- **Personal health** — Regular cycling improves cardiovascular fitness, mental well-being, and longevity.
- **Equity & access** — Bicycles are affordable, require no fuel, and open up mobility to people who cannot drive or afford a car.
- **Economic benefits** — Cyclists spend more at local businesses per kilometre travelled than motorists.

### How You Can Help

1. **Ride when you can.** Every trip made by bike instead of car counts.
2. **Advocate locally.** Attend city council meetings and support protected bike lane proposals.
3. **Be visible.** Use lights, wear bright clothing, and ride predictably.
4. **Welcome newcomers.** Offer to do a group ride with someone new to cycling.
5. **Support bike-friendly businesses.** Patronise shops that provide bike parking and commuter facilities.

### Resources

- [PeopleForBikes](https://www.peopleforbikes.org)
- [European Cyclists Federation](https://ecf.com)
- [Cycling UK](https://www.cyclinguk.org)
- [Bike to Work Day](https://bikeleague.org/bikemonth/)

> *"Life is like riding a bicycle. To keep your balance, you must keep moving."* — Albert Einstein
