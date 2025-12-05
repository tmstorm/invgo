# Changelog

## [0.3.1](https://github.com/tmstorm/invgo/compare/v0.3.0...v0.3.1) (2025-12-05)


### 🐛 Bug Fixes

* invgate json tags got removed ([2f3dee5](https://github.com/tmstorm/invgo/commit/2f3dee5eed6e5afed5737ede43d6add2291f70c4))

## [0.3.0](https://github.com/tmstorm/invgo/compare/v0.2.1...v0.3.0) (2025-12-05)


### 🚀 Features

* add url parsing util, AllowHTTP, and comments to Invgate struct ([fc89ae4](https://github.com/tmstorm/invgo/commit/fc89ae4845427f9d645303976aa7000dd98e50ed))
* expose InvgateAPIPath to public API ([faac90f](https://github.com/tmstorm/invgo/commit/faac90fa8c773fd5ea8983259c89252a83a2df41))


### 📝 Documentation

* add CONTRIBUTING.md ([78f5099](https://github.com/tmstorm/invgo/commit/78f50995cbbad0f017474f06018144d16c071bc1))
* add doc.go ([4bb8fe6](https://github.com/tmstorm/invgo/commit/4bb8fe68b4f25fe9346d1b2cf82ed47814e12947))
* clean up comments ([68ddf89](https://github.com/tmstorm/invgo/commit/68ddf89f76e2dfcd741032aec1f1b7be9deaa17a))
* correct missing AllowHTTP description ([b02b555](https://github.com/tmstorm/invgo/commit/b02b555ba10d474624cc214b629fada5339ad56c))
* fix comments and ensure packag is fully documented ([d3d4ded](https://github.com/tmstorm/invgo/commit/d3d4ded807092ec2ae42141945628849e711ef3f))
* update docs for AllowHTTP ([b041a9e](https://github.com/tmstorm/invgo/commit/b041a9e54e81eac909275ede73ff82b982a4ae14))


### 🔨 Refactoring

* move scopes check to methods.go ([2c1425b](https://github.com/tmstorm/invgo/commit/2c1425b504a6362744f3ee2b3fbce3f0d193f793))


### 🧹 Maintenance

* eliminate endpoint boilerplate with NewPublicMethod ([e758668](https://github.com/tmstorm/invgo/commit/e758668c18e537088a7dd7fec1357d3120fd6f0d))
* move client methods to endpoint_methods.go ([57fac43](https://github.com/tmstorm/invgo/commit/57fac437cf9028e3cf87df8060bd0b7704fe3630))

## [0.2.1](https://github.com/tmstorm/invgo/compare/v0.2.0...v0.2.1) (2025-12-04)


### 🐛 Bug Fixes

* correct public endpoint parameter structs ([a392670](https://github.com/tmstorm/invgo/commit/a39267088ab4e06ef8ff2162276990fa6f5b1b18))
* endpoint struct could not be imported to client code ([8894ca6](https://github.com/tmstorm/invgo/commit/8894ca6779e3bfb2ef689a44161aa7bada3b46c9))

## [0.2.0](https://github.com/tmstorm/invgo/compare/v0.1.6...v0.2.0) (2025-12-04)


### 🚀 Features

* refactor code base ([87ed8ea](https://github.com/tmstorm/invgo/commit/87ed8ea7db04fa73fdbc48a155a82c1d875034c3))


### 🧹 Maintenance

* add clarity to endpoints.go ([fd652b9](https://github.com/tmstorm/invgo/commit/fd652b99012fffa5e9725df240cf7ed5d61c000a))
* correct var and field names ([d182ca2](https://github.com/tmstorm/invgo/commit/d182ca222ab7ddfe2ea195a4820223b027ef8687))
* modernize for loops ([bc942c1](https://github.com/tmstorm/invgo/commit/bc942c1916830cf46d1088df3b9042b73aaba0a6))

## [0.1.6](https://github.com/tmstorm/invgo/compare/v0.1.5...v0.1.6) (2025-12-03)


### 🐛 Bug Fixes

* **utils:** nested structs were not being parsed causing malformed calls ([65897d7](https://github.com/tmstorm/invgo/commit/65897d7d07d38efadf6cd43e2c5721ed95a47733))

## [0.1.5](https://github.com/tmstorm/invgo/compare/v0.1.4...v0.1.5) (2025-12-03)


### 🐛 Bug Fixes

* misused err assignment and fix naming ([74a52db](https://github.com/tmstorm/invgo/commit/74a52db4f7f75fadb210305849b8fd998e9815bc))


### 📝 Documentation

* add clarity from refactor ([d6a9a48](https://github.com/tmstorm/invgo/commit/d6a9a48caa8561e3c9d061e1c0820f8ce6b97b62))
* add go pkg badges ([0895e70](https://github.com/tmstorm/invgo/commit/0895e70caff4180640c6f01d9f2c8a5cd17acc50))


### 🔨 Refactoring

* use reflect to create query params ([1b07adb](https://github.com/tmstorm/invgo/commit/1b07adbe6b783a3982d6bd7c21f008c1acfdb8ca))


### 🧹 Maintenance

* correct comments and spelling ([1a75562](https://github.com/tmstorm/invgo/commit/1a755624c461b187044e472db844503144a2afab))

## [0.1.4](https://github.com/tmstorm/invgo/compare/v0.1.3...v0.1.4) (2025-12-01)


### 🧹 Maintenance

* add license and badges ([6ec7e87](https://github.com/tmstorm/invgo/commit/6ec7e87b2b016361b29199d51c7a9777224a726a))
* change repo refs ([29adcac](https://github.com/tmstorm/invgo/commit/29adcac7e5f6098bf423db299dea5005837d31a9))

## [0.1.3](https://github.com/tmstorm/invgo/compare/v0.1.2...v0.1.3) (2025-12-01)


### 📝 Documentation

* spelling and grammar ([9c23c5c](https://github.com/tmstorm/invgo/commit/9c23c5c2a8030847ca563b1e0bf9c54a058c88b8))


### 🧹 Maintenance

* **ci:** update go test versions ([068ac59](https://github.com/tmstorm/invgo/commit/068ac590ed6add9e201a290d49925926110bca88))
* set min go version to base 1.24 ([af27f39](https://github.com/tmstorm/invgo/commit/af27f39366df84af4fc1aee5ec3c109b02b482f0))

## [0.1.2](https://github.com/tmstorm/invgo/compare/v0.1.1...v0.1.2) (2025-11-24)


### 🐛 Bug Fixes

* **code scanning:** workflow does not contain permissions ([#2](https://github.com/tmstorm/invgo/issues/2)) ([798a3a6](https://github.com/tmstorm/invgo/commit/798a3a69d16aecc50804ba7f46ffd6040b7f1833))
* **code scanning:** workflow does not contain permissions ([#4](https://github.com/tmstorm/invgo/issues/4)) ([1973306](https://github.com/tmstorm/invgo/commit/1973306d1e3257175f7865c010049bb3d483a1af))


### 🧹 Maintenance

* **deps-dev:** bump js-yaml from 4.1.0 to 4.1.1 ([#5](https://github.com/tmstorm/invgo/issues/5)) ([cbbb0ee](https://github.com/tmstorm/invgo/commit/cbbb0eea5f6caa5d6d7529ea326d66ed9d609778))

## [0.1.1](https://github.com/tmstorm/invgo/compare/v0.1.0...v0.1.1) (2025-06-04)


### 🧹 Maintenance

* initial commit ([45c6f27](https://github.com/tmstorm/invgo/commit/45c6f27780978d47e5c7990e4e62594e6720886c))
* set inital version to 0.1.0 ([12f303f](https://github.com/tmstorm/invgo/commit/12f303ff1d42b652862d0bf1354e5b2ae7085eb5))
