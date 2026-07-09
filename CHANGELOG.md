# Changelog

## 0.1.0-alpha.6 (2026-07-05)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/hanzoai/go-sdk/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Features

* regenerate SDK from unified OpenAPI spec (retire Stainless) ([#15](https://github.com/hanzoai/go-sdk/issues/15)) ([3e1b9d4](https://github.com/hanzoai/go-sdk/commit/3e1b9d46954732686667895eeff13587c6c9bb9b))


### Chores

* sync repo ([c378d3b](https://github.com/hanzoai/go-sdk/commit/c378d3bf9382c04c06b1eceebb15fdb9157b9731))

## 0.1.0-alpha.5 (2026-05-13)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/hanzoai/go-sdk/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Features

* **api:** api update ([1bf632d](https://github.com/hanzoai/go-sdk/commit/1bf632d95dfdd174716124b6d297d52c749c701c))
* **internal:** support comma format in multipart form encoding ([5c9f5ed](https://github.com/hanzoai/go-sdk/commit/5c9f5edef47f48ac3fad7fd19ef0383296df073a))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([90dd39c](https://github.com/hanzoai/go-sdk/commit/90dd39cf828d73b03d410b1f84ccdecf73ba2d06))
* **client:** use correct format specifier for header serialization ([ee7d75f](https://github.com/hanzoai/go-sdk/commit/ee7d75f38404b1bccdeecc8c80fac25d047b24b5))
* correct test assertions for service Options and BaseURL ([5d2f80e](https://github.com/hanzoai/go-sdk/commit/5d2f80ed7087b3398e0045395bd75fdc9f64ca8f))
* fix for namespace collisions with client and resource test methods ([fb9589b](https://github.com/hanzoai/go-sdk/commit/fb9589b63aa8a8c64a9a07f27a4a68438a98fe16))
* prevent duplicate ? in query params ([6f6019e](https://github.com/hanzoai/go-sdk/commit/6f6019eb07d5abe8bf51a70b176f1390083286c5))
* remove nil Options assertion when no opts passed ([dae317e](https://github.com/hanzoai/go-sdk/commit/dae317efdfddadcc4937abb0b30321d37203dab3))


### Chores

* bump Go 1.26.1 ([013362e](https://github.com/hanzoai/go-sdk/commit/013362e3052a117f7ec272016319ffccf97535b1))
* **ci:** skip lint on metadata-only changes ([fb3f899](https://github.com/hanzoai/go-sdk/commit/fb3f899cc9bb2f75376027931010b82598462444))
* **ci:** skip uploading artifacts on stainless-internal branches ([c9241ca](https://github.com/hanzoai/go-sdk/commit/c9241ca409ad2bb249dee7e346d07ed4689b38da))
* **ci:** support opting out of skipping builds on metadata-only commits ([3410dda](https://github.com/hanzoai/go-sdk/commit/3410ddad11ccb636101075183bbc0f7d4964cb0b))
* **internal:** codegen related update ([94102cf](https://github.com/hanzoai/go-sdk/commit/94102cf2715ee5016727cef18f414a2db9bd2102))
* **internal:** codegen related update ([ec5ccf8](https://github.com/hanzoai/go-sdk/commit/ec5ccf8cf8f2a307075e4de60260f330eefb1a24))
* **internal:** minor cleanup ([0c6c76b](https://github.com/hanzoai/go-sdk/commit/0c6c76b06414068947c484e6e8d15cd86f86d32b))
* **internal:** move custom custom `json` tags to `api` ([93d3871](https://github.com/hanzoai/go-sdk/commit/93d3871c3b3e61982d6b3141b47f28ff55b9d904))
* **internal:** remove mock server code ([db4778c](https://github.com/hanzoai/go-sdk/commit/db4778ca66069a01a548e3a75aa693be1848b614))
* **internal:** tweak CI branches ([34f01e7](https://github.com/hanzoai/go-sdk/commit/34f01e71d27233c83e2e803c0646c698657d6c08))
* **internal:** update gitignore ([c155214](https://github.com/hanzoai/go-sdk/commit/c155214aa6887465df2f9bf27d83682d43bf7719))
* **internal:** use explicit returns ([f586bc6](https://github.com/hanzoai/go-sdk/commit/f586bc6dedbfa35859584b569887578741f9cb98))
* **internal:** use explicit returns in more places ([bb1ab9f](https://github.com/hanzoai/go-sdk/commit/bb1ab9f1dbe94d9069532155b459e09ed577dd6d))
* remove unnecessary error check for url parsing ([bdc727f](https://github.com/hanzoai/go-sdk/commit/bdc727fa03bda1befcc6dfddc25901f50d541b47))
* symlink AGENTS.md and CLAUDE.md to LLM.md ([00b6700](https://github.com/hanzoai/go-sdk/commit/00b67009e79ad14790fa8ec85d82d56379b56877))
* update dependencies to latest ([21022ba](https://github.com/hanzoai/go-sdk/commit/21022ba6958fec6c5f9016068c5a911efe6bab70))
* update Go toolchain to 1.26 ([813c7ad](https://github.com/hanzoai/go-sdk/commit/813c7ad26dda56fdd99c39171e97f0a3040f5702))
* update mock server docs ([f9f053a](https://github.com/hanzoai/go-sdk/commit/f9f053a90fd561dae2a4b16e2c2e6db1f187860c))
* update placeholder string ([1d6c069](https://github.com/hanzoai/go-sdk/commit/1d6c0691aecb36d825782040bc0b2b0b84468377))


### Documentation

* add LLM.md project guide ([9b4d0c4](https://github.com/hanzoai/go-sdk/commit/9b4d0c495e183f851f1f87ac78ac9a908a34e6de))

## 0.1.0-alpha.4 (2026-01-28)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/hanzoai/go-sdk/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** api update ([d24dacc](https://github.com/hanzoai/go-sdk/commit/d24dacc551627f091ed1f01e558d420caf24bb5b))
* **api:** api update ([4825c99](https://github.com/hanzoai/go-sdk/commit/4825c99fc68fa29e850cd04159dc14831f9c22ea))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([ed4aaea](https://github.com/hanzoai/go-sdk/commit/ed4aaea25e878550eab7bf2819955019ed2ef301))
* Complete Go SDK fixes for full compilation and testing ([40cd2e9](https://github.com/hanzoai/go-sdk/commit/40cd2e95c081f76b4eb812d411eb4f308056aad7))
* Resolve Go SDK compilation errors ([ac3c794](https://github.com/hanzoai/go-sdk/commit/ac3c7941b4c50a091951e9cff19b5e07db3aa45f))
* use slices.Concat instead of sometimes modifying r.Options ([865e39b](https://github.com/hanzoai/go-sdk/commit/865e39b45cdc3ae4de260945113bb6e18025b3eb))


### Chores

* bump minimum go version to 1.22 ([7f2076a](https://github.com/hanzoai/go-sdk/commit/7f2076abec37ce5e638500d0fc0e26e628c2a054))
* do not install brew dependencies in ./scripts/bootstrap by default ([85660de](https://github.com/hanzoai/go-sdk/commit/85660de8f532790f0348b2276f3570826fc9af1c))
* update more docs for 1.22 ([83a09fc](https://github.com/hanzoai/go-sdk/commit/83a09fceb3d2fd98447cafa27c6a5f4271b9af19))


### Documentation

* Add working example demonstrating SDK usage ([bd13718](https://github.com/hanzoai/go-sdk/commit/bd1371885540634ff56e1a58fcf5af9a5db5744d))

## 0.1.0-alpha.3 (2025-09-06)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/hanzoai/go-sdk/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** api update ([a7d4e2d](https://github.com/hanzoai/go-sdk/commit/a7d4e2d7f0a701c7b83a5fa62129f511bf78aca7))
* **client:** add support for endpoint-specific base URLs in python ([a1fced6](https://github.com/hanzoai/go-sdk/commit/a1fced627f9e55d536366505ef216745cefa2d6d))
* **client:** add support for reading base URL from environment variable ([fc8778c](https://github.com/hanzoai/go-sdk/commit/fc8778cc31f0d39da5bddb266b2c167e8385110c))
* **client:** support custom http clients ([#9](https://github.com/hanzoai/go-sdk/issues/9)) ([6cf2e49](https://github.com/hanzoai/go-sdk/commit/6cf2e49da411a8daa7cef3eab22e426be8630059))


### Bug Fixes

* **client:** clean up reader resources ([564bf4b](https://github.com/hanzoai/go-sdk/commit/564bf4b621de12be7a25d2431b0e1c08cf9c7f00))
* **client:** correctly update body in WithJSONSet ([2bfb47f](https://github.com/hanzoai/go-sdk/commit/2bfb47f27b1754a24701d795e8f3af2b3752ad51))
* handle empty bodies in WithJSONSet ([f828405](https://github.com/hanzoai/go-sdk/commit/f8284051703ca4348139004ec8b38c99b7e1904c))


### Chores

* **ci:** add timeout thresholds for CI jobs ([4d9cf13](https://github.com/hanzoai/go-sdk/commit/4d9cf137360c379f896289bd456c73a388cdea00))
* **ci:** only use depot for staging repos ([49f9a44](https://github.com/hanzoai/go-sdk/commit/49f9a444a42e40037262cf4f2c0af6d84b2e7bc2))
* configure new SDK language ([70c6b17](https://github.com/hanzoai/go-sdk/commit/70c6b171a497aef2590abc76eb8a98626595ce71))
* **docs:** document pre-request options ([f8f26dc](https://github.com/hanzoai/go-sdk/commit/f8f26dc587e783b6b87bd04879fae15b03da01f6))
* **docs:** grammar improvements ([f3c04c3](https://github.com/hanzoai/go-sdk/commit/f3c04c3b71d319b3ff30e68f93a3071e4c795ea2))
* improve devcontainer setup ([f516bcc](https://github.com/hanzoai/go-sdk/commit/f516bcc5c9565fd35e97cb197ced67c44fd5dba4))
* **internal:** codegen related update ([c0717b5](https://github.com/hanzoai/go-sdk/commit/c0717b5a3e07964d53e20676d72a2c644f35e5e8))
* **internal:** expand CI branch coverage ([a241eb5](https://github.com/hanzoai/go-sdk/commit/a241eb5b8464517c02ff4bd10178c727b2442849))
* **internal:** reduce CI branch coverage ([b380cd0](https://github.com/hanzoai/go-sdk/commit/b380cd00da61b2f85df12c4a35515f5db0a01b1a))
* make go mod tidy continue on error ([fa097f8](https://github.com/hanzoai/go-sdk/commit/fa097f8c401ceabad9d782c07abf94946e8c310a))


### Documentation

* update documentation links to be more uniform ([b54e5f0](https://github.com/hanzoai/go-sdk/commit/b54e5f0947bf859d22c9fdd3f9b2884a543ae531))

## 0.1.0-alpha.2 (2025-04-03)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/hanzoai/go-sdk/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Bug Fixes

* **client:** return error on bad custom url instead of panic ([#7](https://github.com/hanzoai/go-sdk/issues/7)) ([3ce9376](https://github.com/hanzoai/go-sdk/commit/3ce937641c2ece5c60ffcc7a43d6511ceb4c49d6))


### Chores

* update SDK settings ([#5](https://github.com/hanzoai/go-sdk/issues/5)) ([66a9f24](https://github.com/hanzoai/go-sdk/commit/66a9f242233318310c4f996c051c2e04b88f4e61))

## 0.1.0-alpha.1 (2025-03-27)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/hanzoai/go-sdk/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** api update ([#3](https://github.com/hanzoai/go-sdk/issues/3)) ([c10803d](https://github.com/hanzoai/go-sdk/commit/c10803d6ab4467d23a4fc874b0e704648a16291e))


### Chores

* configure new SDK language ([289851b](https://github.com/hanzoai/go-sdk/commit/289851bfe452a74b720a6895bf3fe9733948635a))
* go live ([#1](https://github.com/hanzoai/go-sdk/issues/1)) ([8b01e13](https://github.com/hanzoai/go-sdk/commit/8b01e1365f5eb1b2e771c5f561d65c336f118d78))
