# Changelog

## [0.9.0](https://github.com/tmstorm/invgo/compare/v0.8.0...v0.9.0) (2025-12-23)


### 🚀 Features

* add incidents.by.agent ([b623470](https://github.com/tmstorm/invgo/commit/b623470a63f286d22067cf600cfe3af3bd87749a))
* add incidents.by.cis ([d31229d](https://github.com/tmstorm/invgo/commit/d31229d919480a68f0ee02b0a422e0fa5d3a3541))
* add incidents.by.customer ([c1da53d](https://github.com/tmstorm/invgo/commit/c1da53dd13e0f45adc751f88727a4a504f3016e6))
* add incidents.by.helpdesk ([140c6e9](https://github.com/tmstorm/invgo/commit/140c6e93b1a1b109f994f8215726fb4969f16bd3))
* add incidents.by.sentiment ([08f3cde](https://github.com/tmstorm/invgo/commit/08f3cded6862f84295388cc2a40a1e97bcff411d))
* add incidents.by.view, .details.by.view, .last.hour ([7a05fb5](https://github.com/tmstorm/invgo/commit/7a05fb5312a98d582944e82796a0c316cee514d2))

## [0.8.0](https://github.com/tmstorm/invgo/compare/v0.7.0...v0.8.0) (2025-12-22)


### 🚀 Features

* add incident.link ([01b87e8](https://github.com/tmstorm/invgo/commit/01b87e8c37cde8e3efe3761484822480d3daf924))
* add incident.linked_cis.counters.from ([546c7e9](https://github.com/tmstorm/invgo/commit/546c7e946fcd22d387cc650b8a74657f1887b132))
* add incident.observer ([06502b5](https://github.com/tmstorm/invgo/commit/06502b55d17aed71bbd65164cde9ce96c0aea912))
* add incident.spontaneous_approval ([0d64d92](https://github.com/tmstorm/invgo/commit/0d64d924515dd3e871820b5a1f2dd01375a5eed6))
* add incident.tasks ([de1d16d](https://github.com/tmstorm/invgo/commit/de1d16d2ad904da9855df4ad7bd8abc756fee2d8))
* add incident.waitingfor.agent ([efd0693](https://github.com/tmstorm/invgo/commit/efd06931e6e821b3ba4c8087c44ff59c6bb572b6))
* add incident.waitingfor.customer, .date, .incident ([d6b3e4a](https://github.com/tmstorm/invgo/commit/d6b3e4af5a7e5645e2f4b109d367f99c6c452fd5))
* coverage script now generates json coverage file ([8e966ee](https://github.com/tmstorm/invgo/commit/8e966eeaa7be6cd3ac1f02d068aa6dcaecdd1432))

## [0.7.0](https://github.com/tmstorm/invgo/compare/v0.6.0...v0.7.0) (2025-12-19)


### 🚀 Features

* add incident.approval.add_voter ([8a57ccf](https://github.com/tmstorm/invgo/commit/8a57ccf04056034c415e5b1e5b9d5f85be3d752f))
* add incident.approval.possible_voters ([c997ec5](https://github.com/tmstorm/invgo/commit/c997ec5689c62169a1d821cd721c750cd630d9e6))
* add incident.attributes.priority, incident.attributes.source ([42ebac9](https://github.com/tmstorm/invgo/commit/42ebac9b365af01188878c828e32359dc116113c))
* add incident.cancel ([21e1dce](https://github.com/tmstorm/invgo/commit/21e1dce3faf78cdfb1faea11bfde47d5bad14554))
* add incident.collaborator ([8bc0549](https://github.com/tmstorm/invgo/commit/8bc0549d48e0428469bb657202aedd171d11d5cd))
* add incident.comment, incident.attachment ([53b5954](https://github.com/tmstorm/invgo/commit/53b59549305bd4fa5a0890e8102e28fdf5b88d1d))
* add incident.custom_approval ([c1027eb](https://github.com/tmstorm/invgo/commit/c1027eb059a832523cb91c4e92844f5d78005194))
* add incident.external_entity, incident.waitingfor.external_entity ([1e1e7b6](https://github.com/tmstorm/invgo/commit/1e1e7b6da9d4000fc8e0e9b952f9e474557f75e0))
* add incident.reassign ([9f51250](https://github.com/tmstorm/invgo/commit/9f51250e9dc6eb2b347625781dbc7d6eef64a9d2))
* add incident.reject ([46f1fee](https://github.com/tmstorm/invgo/commit/46f1fee0ca3b34b4785dc014747f2baa86b41ffb))
* add incident.reopen ([f2155cf](https://github.com/tmstorm/invgo/commit/f2155cf8de0f904a0d7cf38c20f68d5b30529ef6))
* add incident.solution.accept, incident.solution.reject ([195c004](https://github.com/tmstorm/invgo/commit/195c004f61c6f3eae4a6b08e4a6c73b94a5fa893))


### 🐛 Bug Fixes

* incident approval get return did not match docs ([a6fbf30](https://github.com/tmstorm/invgo/commit/a6fbf30236f2be08f3d2ed9060ab5210133ba41f))
* incident.reopen structs were wrong ([3f8a3ef](https://github.com/tmstorm/invgo/commit/3f8a3efd8149d9544c3b963745c3a932ac98d565))

## [0.6.0](https://github.com/tmstorm/invgo/compare/v0.5.0...v0.6.0) (2025-12-16)


### 🚀 Features

* add incident approval accept, cancel, reject ([96f53ad](https://github.com/tmstorm/invgo/commit/96f53ade5d04d93cee708ed64732b279f057dca7))
* add incident approval, approval.status, approval.type, approval.vote_status ([d4b37c7](https://github.com/tmstorm/invgo/commit/d4b37c726d53cf41aa6c4ddf9e46b5a2cd19cad1))
* add remaining users endpoints ([5352178](https://github.com/tmstorm/invgo/commit/5352178938e5263820d6734666c9ec381f1f7245))
* add timetracking endpoints ([4d0c41a](https://github.com/tmstorm/invgo/commit/4d0c41a6d8e1730afb4c8d50eba6f57510f081ce))
* add workflow endpoints ([3aaa0a2](https://github.com/tmstorm/invgo/commit/3aaa0a23ec765d352e1460e0f3bb85e28fc47d82))


### 🐛 Bug Fixes

* **coverage:** helpdesks was misspelled in coverage.go ([fcd1e3e](https://github.com/tmstorm/invgo/commit/fcd1e3ea37d0a61f49bb216d83381db1dfac4eba))
* incident endpoints were under incidents in coverage.go ([e405166](https://github.com/tmstorm/invgo/commit/e4051669789c36b632074a49b8c9903c715366fb))

## [0.5.0](https://github.com/tmstorm/invgo/compare/v0.4.1...v0.5.0) (2025-12-11)


### 🚀 Features

* add remaining user endpoints ([3eb6566](https://github.com/tmstorm/invgo/commit/3eb6566836875603c21b7c2c0e3334e3446aa100))


### 🐛 Bug Fixes

* tests were still using old newPublicMethod ([97cc503](https://github.com/tmstorm/invgo/commit/97cc5030ff68030d9304972efca4d70e1ae4c6ee))

## [0.4.1](https://github.com/tmstorm/invgo/compare/v0.4.0...v0.4.1) (2025-12-11)


### 🔨 Refactoring

* make NewPublicMethod private ([e0f87cc](https://github.com/tmstorm/invgo/commit/e0f87cc3ab48c6e9042b91b69d0b4f2027b8b21f))
* use reflect to create newPublicMethods ([68e7e1f](https://github.com/tmstorm/invgo/commit/68e7e1f88706b829c6eb6a6e67766233435634cf))

## [0.4.0](https://github.com/tmstorm/invgo/compare/v0.3.2...v0.4.0) (2025-12-09)


### 🚀 Features

* add Triggers endpoint ([1653aa4](https://github.com/tmstorm/invgo/commit/1653aa4692ee46601e5bd30eb52c667b9435e43c))
* partially implement user endpoints ([75a34d6](https://github.com/tmstorm/invgo/commit/75a34d68cadb279c9b6a00f06dc7c62ac3a10c4f))

## [0.3.2](https://github.com/tmstorm/invgo/compare/v0.3.1...v0.3.2) (2025-12-08)


### 🐛 Bug Fixes

* categories get was not returning response ([649b3cc](https://github.com/tmstorm/invgo/commit/649b3ccb114c7e494bd8511d38dcc3672399a9f3))


### 🧹 Maintenance

* add tests for core and endpoints ([8dd4014](https://github.com/tmstorm/invgo/commit/8dd4014257c2dccdbc5a5f71f4a5db7c9e732272))
* fix spelling in CONTRIBUTING.md ([396a9dc](https://github.com/tmstorm/invgo/commit/396a9dcbdf703c5fe53c1260a87c8b1599e8f804))
* make ServiceDeskVersionResponse public ([3a18c1e](https://github.com/tmstorm/invgo/commit/3a18c1e5133f66153c029fb2c51c6b9c514f97b9))

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
