# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project
aims to follow [Semantic Versioning](https://semver.org/spec/v2.0.0.html) once the first release is
tagged. Until then, source builds report the version `dev`.

## [0.4.0](https://github.com/pvyswiss/pvyai-coding-agent/compare/v0.3.1...v0.4.0) (2026-07-16)


### Features

* **acp:** Agent Client Protocol surface — drive ZERO as an editor backend ([#288](https://github.com/pvyswiss/pvyai-coding-agent/issues/288)) ([5b34805](https://github.com/pvyswiss/pvyai-coding-agent/commit/5b348056a4bbfc648c9efd75af286284b23c6be4))
* add --auto flag for LLM-generated commit messages ([#423](https://github.com/pvyswiss/pvyai-coding-agent/issues/423)) ([b0abde7](https://github.com/pvyswiss/pvyai-coding-agent/commit/b0abde7d0697e808480cd59d69a6f4d0c6320475))
* add `providers detect` and `sandbox check` (wire audit-found unwired features) + coverage ([#238](https://github.com/pvyswiss/pvyai-coding-agent/issues/238)) ([29fc29a](https://github.com/pvyswiss/pvyai-coding-agent/commit/29fc29a41d1ac1e82dbc379c0577d3dce481ba8a))
* add backend doctor diagnostics ([#194](https://github.com/pvyswiss/pvyai-coding-agent/issues/194)) ([0444eb0](https://github.com/pvyswiss/pvyai-coding-agent/commit/0444eb0e6fdb5311189501fed2370e5a760ef414))
* add backend lifecycle status command ([#192](https://github.com/pvyswiss/pvyai-coding-agent/issues/192)) ([6895cbb](https://github.com/pvyswiss/pvyai-coding-agent/commit/6895cbb4aabc9ed2d3ac3db2d727de40a18a21f1))
* add config validation diagnostics ([#35](https://github.com/pvyswiss/pvyai-coding-agent/issues/35)) ([bed4dd7](https://github.com/pvyswiss/pvyai-coding-agent/commit/bed4dd77524e5ded609e615d93b25bfa5560633e))
* add doctor fix flow ([#198](https://github.com/pvyswiss/pvyai-coding-agent/issues/198)) ([284a948](https://github.com/pvyswiss/pvyai-coding-agent/commit/284a948f00207a46ace96ba832ab5937893c7086))
* add go agent runtime ([9cb6803](https://github.com/pvyswiss/pvyai-coding-agent/commit/9cb680318f862d36cfbdcfdf97dfee4fc9c8ffc5))
* add Go command center ([#63](https://github.com/pvyswiss/pvyai-coding-agent/issues/63)) ([ef20e4b](https://github.com/pvyswiss/pvyai-coding-agent/commit/ef20e4bcae1f289d170fdc2ded43e8226521ba76))
* add Go headless protocol sessions ([#60](https://github.com/pvyswiss/pvyai-coding-agent/issues/60)) ([a710b0e](https://github.com/pvyswiss/pvyai-coding-agent/commit/a710b0e5f5fa621aba48d01f7a9405239702ca86))
* add Go MCP network transports ([95e6524](https://github.com/pvyswiss/pvyai-coding-agent/commit/95e652446dc6df6031ce51913ca2d02ddba1603f))
* add go provider config and tui foundations ([b4c374c](https://github.com/pvyswiss/pvyai-coding-agent/commit/b4c374c6a2c379bd970475283cc0f3274fb04a9b))
* add Go session lineage backend ([b5eec64](https://github.com/pvyswiss/pvyai-coding-agent/commit/b5eec64acf975e5310e805790427c733b61d1947))
* add Go worktree and verification backend ([#70](https://github.com/pvyswiss/pvyai-coding-agent/issues/70)) ([a58e6c4](https://github.com/pvyswiss/pvyai-coding-agent/commit/a58e6c4f575f8aa5d8114602e693d29e0763a4ee))
* add local plugin manifest loader ([#42](https://github.com/pvyswiss/pvyai-coding-agent/issues/42)) ([10f749e](https://github.com/pvyswiss/pvyai-coding-agent/commit/10f749e7a7a3ee57bb694bf67b4059fc38536a76))
* add m1 headless exec surface ([#17](https://github.com/pvyswiss/pvyai-coding-agent/issues/17)) ([291af70](https://github.com/pvyswiss/pvyai-coding-agent/commit/291af70174f8b86f187419110a970d4ac5390229))
* add M2 install scripts ([#19](https://github.com/pvyswiss/pvyai-coding-agent/issues/19)) ([f9de264](https://github.com/pvyswiss/pvyai-coding-agent/commit/f9de264e9bd838c454e2f9865b4c57b3120645fa))
* add M2 performance benchmark harness ([#20](https://github.com/pvyswiss/pvyai-coding-agent/issues/20)) ([0ac9009](https://github.com/pvyswiss/pvyai-coding-agent/commit/0ac90092707959fedd4b83d114c454f683fb2eb7))
* add M2 release checksum verification ([#25](https://github.com/pvyswiss/pvyai-coding-agent/issues/25)) ([50374f4](https://github.com/pvyswiss/pvyai-coding-agent/commit/50374f449089c0fe736b8181514794dd1976a665))
* add M2 update check ([#16](https://github.com/pvyswiss/pvyai-coding-agent/issues/16)) ([b801c53](https://github.com/pvyswiss/pvyai-coding-agent/commit/b801c532ef91ff597fcc5b833e3ca3ee04544eb2))
* add M5 git workflow backend ([#72](https://github.com/pvyswiss/pvyai-coding-agent/issues/72)) ([4eacb5b](https://github.com/pvyswiss/pvyai-coding-agent/commit/4eacb5b6d30f9276eaa48f0bfde9610d9ee830e5))
* add M5 self verification backend ([#76](https://github.com/pvyswiss/pvyai-coding-agent/issues/76)) ([b973084](https://github.com/pvyswiss/pvyai-coding-agent/commit/b9730849692757b7176a907551188dffda25cb46))
* add M5 test runner backend ([#74](https://github.com/pvyswiss/pvyai-coding-agent/issues/74)) ([8e1e74a](https://github.com/pvyswiss/pvyai-coding-agent/commit/8e1e74ad4bb9cb3dab64750e7d4fc59e56fd68cc))
* add M6 sandbox backend ([#77](https://github.com/pvyswiss/pvyai-coding-agent/issues/77)) ([29aec1b](https://github.com/pvyswiss/pvyai-coding-agent/commit/29aec1b190c48583ef235229e0b807638cef69bc))
* add MCP manager UX ([#188](https://github.com/pvyswiss/pvyai-coding-agent/issues/188)) ([250bf24](https://github.com/pvyswiss/pvyai-coding-agent/commit/250bf24c9d43a760841f9ed66c07e9d5d8e00274))
* add MCP permission grants ([#40](https://github.com/pvyswiss/pvyai-coding-agent/issues/40)) ([6af3c8c](https://github.com/pvyswiss/pvyai-coding-agent/commit/6af3c8ceaeb8d2c16c2732fbfd398f21bb8765fd))
* add MCP server protocol ([bb88c84](https://github.com/pvyswiss/pvyai-coding-agent/commit/bb88c84359d8ca8d818209a88441c060546242cd))
* add premium startup splash ([5c58ea3](https://github.com/pvyswiss/pvyai-coding-agent/commit/5c58ea33fc60c745ef6f39cb380ece3322f00fd6))
* add runtime permission decisions ([#95](https://github.com/pvyswiss/pvyai-coding-agent/issues/95)) ([52fbb6d](https://github.com/pvyswiss/pvyai-coding-agent/commit/52fbb6d9e73ce8c8e541261d81ed80f04c1ad5b1))
* add session search indexing ([#36](https://github.com/pvyswiss/pvyai-coding-agent/issues/36)) ([fd978f8](https://github.com/pvyswiss/pvyai-coding-agent/commit/fd978f8589c4b509158c4add69ae1c690074024e))
* add TUI diagnostics center ([#197](https://github.com/pvyswiss/pvyai-coding-agent/issues/197)) ([7153ec7](https://github.com/pvyswiss/pvyai-coding-agent/commit/7153ec7a7d434a9f70a5d237c52c0bad5284504b))
* add tui model selector ([#22](https://github.com/pvyswiss/pvyai-coding-agent/issues/22)) ([a8c8143](https://github.com/pvyswiss/pvyai-coding-agent/commit/a8c814352864e129c313881580300a31d0694dca))
* add TUI session controls ([4afacf7](https://github.com/pvyswiss/pvyai-coding-agent/commit/4afacf75275a1ccbcc5de617d72ad728ccb9f6d3))
* add Zero Anthropic provider ([#15](https://github.com/pvyswiss/pvyai-coding-agent/issues/15)) ([b28f497](https://github.com/pvyswiss/pvyai-coding-agent/commit/b28f497c72917ccce869717772b783abf3caab78))
* add zero changes push and pr subcommands, and extra repo-info metrics ([#391](https://github.com/pvyswiss/pvyai-coding-agent/issues/391)) ([2312abe](https://github.com/pvyswiss/pvyai-coding-agent/commit/2312abe5ddd95f4c6ef373cfb61cc03092f48cdd))
* add Zero config inspection ([#29](https://github.com/pvyswiss/pvyai-coding-agent/issues/29)) ([d249173](https://github.com/pvyswiss/pvyai-coding-agent/commit/d24917314891f792ea5852bb0776f553bcc432e1))
* add Zero doctor backend ([#28](https://github.com/pvyswiss/pvyai-coding-agent/issues/28)) ([30c8c32](https://github.com/pvyswiss/pvyai-coding-agent/commit/30c8c32eb83f170e734fa0f00c205f209e873fc3))
* add Zero exec resume and fork sessions ([#32](https://github.com/pvyswiss/pvyai-coding-agent/issues/32)) ([5b01fbb](https://github.com/pvyswiss/pvyai-coding-agent/commit/5b01fbb49edf017abbd758fb35c532680bd66619))
* add Zero Gemini provider ([#18](https://github.com/pvyswiss/pvyai-coding-agent/issues/18)) ([b909b6f](https://github.com/pvyswiss/pvyai-coding-agent/commit/b909b6f105829a8dad79e339462f39d7d78499fd))
* add Zero hook backend ([#43](https://github.com/pvyswiss/pvyai-coding-agent/issues/43)) ([a6ab274](https://github.com/pvyswiss/pvyai-coding-agent/commit/a6ab2744fd51f7ace084d93f64685452a1af7b33))
* add Zero MCP client backend ([#37](https://github.com/pvyswiss/pvyai-coding-agent/issues/37)) ([cf84860](https://github.com/pvyswiss/pvyai-coding-agent/commit/cf8486067ab6ed927051ac1d074d73d75c2d774c))
* add Zero model registry ([#9](https://github.com/pvyswiss/pvyai-coding-agent/issues/9)) ([abab103](https://github.com/pvyswiss/pvyai-coding-agent/commit/abab103d61a6b515fdb901b85a972d2da5abc5af))
* add Zero provider runtime resolver ([#12](https://github.com/pvyswiss/pvyai-coding-agent/issues/12)) ([5526d39](https://github.com/pvyswiss/pvyai-coding-agent/commit/5526d3904e9541cfe56f2ba95a6abad90c78802a))
* add Zero secret redaction ([#26](https://github.com/pvyswiss/pvyai-coding-agent/issues/26)) ([a681f2b](https://github.com/pvyswiss/pvyai-coding-agent/commit/a681f2bbd164d35138511456d30a4d3c45db9053))
* add Zero session event store ([#21](https://github.com/pvyswiss/pvyai-coding-agent/issues/21)) ([16d9213](https://github.com/pvyswiss/pvyai-coding-agent/commit/16d9213b344517f38f5a4f2d0ac80081120bf43a))
* add Zero session search ([#27](https://github.com/pvyswiss/pvyai-coding-agent/issues/27)) ([ab1da27](https://github.com/pvyswiss/pvyai-coding-agent/commit/ab1da27f5552f0d25aa4e744020f411ea2578402))
* add Zero stream-json protocol ([#31](https://github.com/pvyswiss/pvyai-coding-agent/issues/31)) ([aee8b82](https://github.com/pvyswiss/pvyai-coding-agent/commit/aee8b824e0cef1b3e687b01f93ef3e8b4c9094a9))
* add Zero usage tracking backend ([#23](https://github.com/pvyswiss/pvyai-coding-agent/issues/23)) ([6b76879](https://github.com/pvyswiss/pvyai-coding-agent/commit/6b76879c606fd6f986508b0dbbd77d9547e03382))
* advance Go-first TUI and npm wrapper ([#57](https://github.com/pvyswiss/pvyai-coding-agent/issues/57)) ([37f3cc1](https://github.com/pvyswiss/pvyai-coding-agent/commit/37f3cc17a5205cdb38887432bd79553b7bc20060))
* agent quality, caching, retry, and tooling upgrades ([#506](https://github.com/pvyswiss/pvyai-coding-agent/issues/506)) ([3c81fea](https://github.com/pvyswiss/pvyai-coding-agent/commit/3c81fea22873ee3df7fc97b10cb4f77792706c4b))
* **agent:** curb over-engineering the solution in the editing discipline ([#517](https://github.com/pvyswiss/pvyai-coding-agent/issues/517)) ([f4c998a](https://github.com/pvyswiss/pvyai-coding-agent/commit/f4c998ac30c4f07ff313a2d706791e857293be49)), closes [#516](https://github.com/pvyswiss/pvyai-coding-agent/issues/516)
* **agent:** inject per-user config.UserConfigDir()/zero/ZERO.md guidelines into system prompt ([#475](https://github.com/pvyswiss/pvyai-coding-agent/issues/475)) ([7b10aab](https://github.com/pvyswiss/pvyai-coding-agent/commit/7b10aab74bf14a01166b2cea22deab79bba9850b))
* **agent:** preamble before multi-step tasks for a readable chat ([#317](https://github.com/pvyswiss/pvyai-coding-agent/issues/317)) ([ca41ac2](https://github.com/pvyswiss/pvyai-coding-agent/commit/ca41ac2b6c09c574df6f475b27bbbee6edd3feb5))
* clipboard image paste + provider 400 handler for image rejection ([#268](https://github.com/pvyswiss/pvyai-coding-agent/issues/268)) ([c572791](https://github.com/pvyswiss/pvyai-coding-agent/commit/c572791edf282c97f271ba4f64af833931afd496))
* context-aware compaction + usage gauge for all models ([#332](https://github.com/pvyswiss/pvyai-coding-agent/issues/332)) ([ca5e0b9](https://github.com/pvyswiss/pvyai-coding-agent/commit/ca5e0b9bbb2353eee0b9df30f9df0e1bbf9fb5a4))
* **daemon:** TLS remote bridge — authenticated network access to the daemon ([#211](https://github.com/pvyswiss/pvyai-coding-agent/issues/211)) ([f857c4a](https://github.com/pvyswiss/pvyai-coding-agent/commit/f857c4afb4ca347a4267cc49d3ff2a14c0b248d0))
* **daemon:** zero daemon — worker pool + session supervisor over a local control socket ([#203](https://github.com/pvyswiss/pvyai-coding-agent/issues/203)) ([135e3c5](https://github.com/pvyswiss/pvyai-coding-agent/commit/135e3c56572709657b395e198cccb8fa245168c6))
* encrypted provider credentials + provider-aware /model ([#329](https://github.com/pvyswiss/pvyai-coding-agent/issues/329)) ([a4b2e01](https://github.com/pvyswiss/pvyai-coding-agent/commit/a4b2e01fe69d0342e48f6bdec0fe3d936709026e))
* expose MCP server CLI ([2a5b89c](https://github.com/pvyswiss/pvyai-coding-agent/commit/2a5b89c095e08606bc22feb2a7e2ae5b6b10c175))
* five capability features from the reference-agent comparison ([#276](https://github.com/pvyswiss/pvyai-coding-agent/issues/276)) ([b2bb6c4](https://github.com/pvyswiss/pvyai-coding-agent/commit/b2bb6c4b41107405833a3bd80167fc61b274978e))
* free out-of-the-box web (keyless Firecrawl default) + resilient MCP startup ([#239](https://github.com/pvyswiss/pvyai-coding-agent/issues/239)) ([c0fedad](https://github.com/pvyswiss/pvyai-coding-agent/commit/c0fedade8d3f0a3d0c8e07f97e650f34170b5d31))
* harden sandbox command execution ([#79](https://github.com/pvyswiss/pvyai-coding-agent/issues/79)) ([623c53e](https://github.com/pvyswiss/pvyai-coding-agent/commit/623c53e3240cd2686de6a7e24b2c6ae3f332a94d))
* make Go runtime the app path ([#71](https://github.com/pvyswiss/pvyai-coding-agent/issues/71)) ([43c2560](https://github.com/pvyswiss/pvyai-coding-agent/commit/43c256068ac2ba2f9063ffb58720430adb7324a6))
* **npm:** publish @gitlawb/zero with a postinstall binary installer ([#284](https://github.com/pvyswiss/pvyai-coding-agent/issues/284)) ([33da6b6](https://github.com/pvyswiss/pvyai-coding-agent/commit/33da6b6454ed18f3098821d1ad142ec68645ffac))
* **oauth:** Hugging Face + ChatGPT Codex as first-class built-in OAuth presets ([#245](https://github.com/pvyswiss/pvyai-coding-agent/issues/245)) ([ab3d09d](https://github.com/pvyswiss/pvyai-coding-agent/commit/ab3d09d78e0f55c45241c724ac6207c5ce3aa31e))
* **oauth:** reusable provider OAuth engine + zero auth CLI ([#210](https://github.com/pvyswiss/pvyai-coding-agent/issues/210)) ([bb7de29](https://github.com/pvyswiss/pvyai-coding-agent/commit/bb7de29c80a6156763842696f9c0c2a940c23621))
* **openai:** forward prompt_cache_key for server-side prefix cache routing ([#515](https://github.com/pvyswiss/pvyai-coding-agent/issues/515)) ([87e7e69](https://github.com/pvyswiss/pvyai-coding-agent/commit/87e7e69afd18b5539579856f3a61c6a95bc445ae))
* persist TUI sessions ([#69](https://github.com/pvyswiss/pvyai-coding-agent/issues/69)) ([8a621fc](https://github.com/pvyswiss/pvyai-coding-agent/commit/8a621fc46eb5fa3804bd3b6fe10ca9697928d709))
* **providers:** add `zero providers models` to discover a provider's models ([#386](https://github.com/pvyswiss/pvyai-coding-agent/issues/386)) ([0bc8074](https://github.com/pvyswiss/pvyai-coding-agent/commit/0bc8074c97b0310e4a9d70c3f967003ee5e8a59f))
* **providers:** add KiloCode and OpenCode provider support ([#388](https://github.com/pvyswiss/pvyai-coding-agent/issues/388)) ([b1ccb6d](https://github.com/pvyswiss/pvyai-coding-agent/commit/b1ccb6d9c1875377f5e5ea81a1304edd1e41ab4f))
* **providers:** add Meituan LongCat catalog preset ([#424](https://github.com/pvyswiss/pvyai-coding-agent/issues/424)) ([b4275e3](https://github.com/pvyswiss/pvyai-coding-agent/commit/b4275e350472b2490212bf814709819d354c1216))
* **providers:** split minimax zai into intl cn ([#398](https://github.com/pvyswiss/pvyai-coding-agent/issues/398)) ([aaad4d2](https://github.com/pvyswiss/pvyai-coding-agent/commit/aaad4d271270f41af837b6f3b60ae80beba0c645))
* publish zero to npm via release-please ([#367](https://github.com/pvyswiss/pvyai-coding-agent/issues/367)) ([8eccc26](https://github.com/pvyswiss/pvyai-coding-agent/commit/8eccc2669887bc38d35bc16a315c888e4d9ec43a))
* **reasoning:** typed per-provider reasoning capability from a models.dev catalog ([#338](https://github.com/pvyswiss/pvyai-coding-agent/issues/338)) ([8c5898d](https://github.com/pvyswiss/pvyai-coding-agent/commit/8c5898d5dba86808ae4fa3be5e0203bd61dc529b))
* redesign tui shell ([59bd917](https://github.com/pvyswiss/pvyai-coding-agent/commit/59bd917ca812a19e05613d726c48814d31c5e22e))
* require manual approval before npm publish + drop release-as pin ([#369](https://github.com/pvyswiss/pvyai-coding-agent/issues/369)) ([bd89a1f](https://github.com/pvyswiss/pvyai-coding-agent/commit/bd89a1f451643c1b65ec803070abc7b116631ebe))
* **sandbox:** complete sandbox breadth — WSL fallback + opt-in MITM TLS inspection ([#206](https://github.com/pvyswiss/pvyai-coding-agent/issues/206)) ([ba8fbb8](https://github.com/pvyswiss/pvyai-coding-agent/commit/ba8fbb8b71e3062c8fc586f5522334afc9e783d5))
* **sandbox:** harden internal/sandbox — re-entrancy guard, SOCKS5 egress, sandbox-aware search, fine-grained path lists ([#199](https://github.com/pvyswiss/pvyai-coding-agent/issues/199)) ([671af59](https://github.com/pvyswiss/pvyai-coding-agent/commit/671af594fb8133099a570e2282ebac449861dd67))
* **sandbox:** online/offline Windows network identity — run approved network commands sandboxed ([#297](https://github.com/pvyswiss/pvyai-coding-agent/issues/297)) ([1b1abd5](https://github.com/pvyswiss/pvyai-coding-agent/commit/1b1abd5bf17601a4a93ce8b3f811141ee900c58c))
* **sandbox:** unelevated Windows fallback tier instead of prompts-only degrade ([#427](https://github.com/pvyswiss/pvyai-coding-agent/issues/427)) ([b9ddd6f](https://github.com/pvyswiss/pvyai-coding-agent/commit/b9ddd6f42138312a1fee8d8bb67c46c8eb1dea2f))
* **specialist:** add exec metadata flags ([#104](https://github.com/pvyswiss/pvyai-coding-agent/issues/104)) ([ae6c959](https://github.com/pvyswiss/pvyai-coding-agent/commit/ae6c9595dbeaaa2d032ab6dafa000316389b3f82))
* **specialist:** add profile management ([#100](https://github.com/pvyswiss/pvyai-coding-agent/issues/100)) ([af6a0ef](https://github.com/pvyswiss/pvyai-coding-agent/commit/af6a0ef204b2f1a18a008dac7c07edc3a80e2eee))
* **subagent:** proactive specialist delegation + swarm execution fixes ([#243](https://github.com/pvyswiss/pvyai-coding-agent/issues/243)) ([995b908](https://github.com/pvyswiss/pvyai-coding-agent/commit/995b9086c7bc32664d4b55e0ec468591d1272701))
* support shift enter for composer newlines ([#462](https://github.com/pvyswiss/pvyai-coding-agent/issues/462)) ([daf65e0](https://github.com/pvyswiss/pvyai-coding-agent/commit/daf65e0af9a040314d4ab337b0ad59c55416b7bc))
* **swarm:** let swarm members build (write + sandboxed shell) ([#313](https://github.com/pvyswiss/pvyai-coding-agent/issues/313)) ([46c44a8](https://github.com/pvyswiss/pvyai-coding-agent/commit/46c44a87532510e8bc7f517a7256c9065dea5561))
* **swarm:** multi-agent swarm over the specialist sub-agent ([#207](https://github.com/pvyswiss/pvyai-coding-agent/issues/207)) ([9ddef0e](https://github.com/pvyswiss/pvyai-coding-agent/commit/9ddef0e9570a7ac40205454e666cc97d73bc5d2f))
* **tools:** show red/green diff preview on write_file and edit_file ([#318](https://github.com/pvyswiss/pvyai-coding-agent/issues/318)) ([f9a26e7](https://github.com/pvyswiss/pvyai-coding-agent/commit/f9a26e7148396d6ce9901f856dc576891a351005))
* **tui:** `?` keyboard-shortcut overlay ([#271](https://github.com/pvyswiss/pvyai-coding-agent/issues/271)) ([ead2b92](https://github.com/pvyswiss/pvyai-coding-agent/commit/ead2b9239eeb9fbe38076cf54f48664ff753d0fc))
* **tui:** add Claude-Code-style mode status line to splash ([fbd0bcb](https://github.com/pvyswiss/pvyai-coding-agent/commit/fbd0bcb238c718fca770d3d1aade979ff57cf1d7))
* **tui:** add search/filter to provider picker in setup wizard ([#400](https://github.com/pvyswiss/pvyai-coding-agent/issues/400)) ([2fcea71](https://github.com/pvyswiss/pvyai-coding-agent/commit/2fcea71778d23e050c93409c471aef45b68c1621))
* **tui:** age-based fade for streaming assistant text ([#223](https://github.com/pvyswiss/pvyai-coding-agent/issues/223)) ([1e25fde](https://github.com/pvyswiss/pvyai-coding-agent/commit/1e25fde3b102c328abd7e80be947d45297a6e01b))
* **tui:** animate live dot and running tool rows ([5e89e6d](https://github.com/pvyswiss/pvyai-coding-agent/commit/5e89e6d38e1a36614036b1b5bcccb0627462b646))
* **tui:** ask_user suggested answers with a recommended default ([#326](https://github.com/pvyswiss/pvyai-coding-agent/issues/326)) ([1d1fa7c](https://github.com/pvyswiss/pvyai-coding-agent/commit/1d1fa7c51435c422a052df3fd8bf2983fbf4376d))
* **tui:** chat + context sidebar redesign, AGENTS panel, animations ([#300](https://github.com/pvyswiss/pvyai-coding-agent/issues/300)) ([63c0e27](https://github.com/pvyswiss/pvyai-coding-agent/commit/63c0e278ff1d65e95b581f893321df6305ba4d6d))
* **tui:** click a swarm member in the sidebar to open its subchat ([#311](https://github.com/pvyswiss/pvyai-coding-agent/issues/311)) ([11926cc](https://github.com/pvyswiss/pvyai-coding-agent/commit/11926cc5d29655747c6e8a4c7ade2d492d6ba73a))
* **tui:** clickable plan-step detail + plan/sidebar UX & rendering polish ([#315](https://github.com/pvyswiss/pvyai-coding-agent/issues/315)) ([0e7318f](https://github.com/pvyswiss/pvyai-coding-agent/commit/0e7318ff727bb815d2855175935f1e8b7df00133))
* **tui:** compact one-line /model switch confirmation ([#331](https://github.com/pvyswiss/pvyai-coding-agent/issues/331)) ([8aae0be](https://github.com/pvyswiss/pvyai-coding-agent/commit/8aae0be6db43257420941e246d9d9281cc835684))
* **tui:** Ctrl+T cycle + divider display for reasoning effort ([#229](https://github.com/pvyswiss/pvyai-coding-agent/issues/229)) ([922a5ba](https://github.com/pvyswiss/pvyai-coding-agent/commit/922a5bab5f90c4ba2cc5745c4fb06ad37093060f))
* **tui:** declutter the chat — drop the gimmicky billboard chrome ([#291](https://github.com/pvyswiss/pvyai-coding-agent/issues/291)) ([0aa4569](https://github.com/pvyswiss/pvyai-coding-agent/commit/0aa45695a8936bcda0b53aaa37a01b240da67669))
* **tui:** dim the transcript backdrop behind overlays for visibility ([#333](https://github.com/pvyswiss/pvyai-coding-agent/issues/333)) ([653cee6](https://github.com/pvyswiss/pvyai-coding-agent/commit/653cee60c1f1165caf8df9efcd6ed13b210f09d5))
* **tui:** FILES sidebar panel with click-to-select and file drill-in ([#365](https://github.com/pvyswiss/pvyai-coding-agent/issues/365)) ([142c548](https://github.com/pvyswiss/pvyai-coding-agent/commit/142c548c89a8652ce300e64ddf1228ee36df7606))
* **tui:** interactive theme picker + color theme catalog ([#354](https://github.com/pvyswiss/pvyai-coding-agent/issues/354)) ([49bec02](https://github.com/pvyswiss/pvyai-coding-agent/commit/49bec0274f44073d3c4f7d1fd9d5606624afb432))
* **tui:** keep finished swarm members visible during the run ([#312](https://github.com/pvyswiss/pvyai-coding-agent/issues/312)) ([1c94168](https://github.com/pvyswiss/pvyai-coding-agent/commit/1c9416824521443a06bc10e4028b28f8ec90f9c1))
* **tui:** live token counter on the working line ([#308](https://github.com/pvyswiss/pvyai-coding-agent/issues/308)) ([8acc1be](https://github.com/pvyswiss/pvyai-coding-agent/commit/8acc1be7b246bc10ea415f1120ca000ca509729e))
* **tui:** make tool results and selections easier to scan at a glance ([#287](https://github.com/pvyswiss/pvyai-coding-agent/issues/287)) ([e52fc3a](https://github.com/pvyswiss/pvyai-coding-agent/commit/e52fc3aebb1910ef7927fbd12c8d86e25844b643))
* **tui:** polish command output ([#98](https://github.com/pvyswiss/pvyai-coding-agent/issues/98)) ([ef34364](https://github.com/pvyswiss/pvyai-coding-agent/commit/ef343643ae4ce0285825edd85ea57bae2c4a7d0c))
* **tui:** premium splash + working-view redesign ([#83](https://github.com/pvyswiss/pvyai-coding-agent/issues/83)) ([4f4f965](https://github.com/pvyswiss/pvyai-coding-agent/commit/4f4f96511eca5f162ba06bc9197158db2b8197e1))
* **tui:** rebuild startup splash into reusable components ([5165efb](https://github.com/pvyswiss/pvyai-coding-agent/commit/5165efbd963e657525069f1457e6d4d3ef822e16))
* **tui:** redesign in-session working view ([5d1889b](https://github.com/pvyswiss/pvyai-coding-agent/commit/5d1889b8865ed55ab6c932e13505263986909def))
* **tui:** remove the /mode command (superseded by /model + /effort) ([#337](https://github.com/pvyswiss/pvyai-coding-agent/issues/337)) ([7a09a42](https://github.com/pvyswiss/pvyai-coding-agent/commit/7a09a42a1776d6e677276f24adcdb9c3f5d12bcd))
* **tui:** rotating "working word" for the liveness spinner ([#221](https://github.com/pvyswiss/pvyai-coding-agent/issues/221)) ([6092303](https://github.com/pvyswiss/pvyai-coding-agent/commit/6092303e73fa6b07016b8f67fe57496f4cdb312b))
* **tui:** show provider on each model-picker row ([#305](https://github.com/pvyswiss/pvyai-coding-agent/issues/305)) ([215170b](https://github.com/pvyswiss/pvyai-coding-agent/commit/215170b967b20428395ff50e5ba495ce76f6f1c5))
* **tui:** show slash-command menu on splash, drop keycap hints ([029ea53](https://github.com/pvyswiss/pvyai-coding-agent/commit/029ea535b92849010f95bfef1afd18d6ca091c6e))
* **tui:** simplify shortcut hints to a clean single line ([382dadf](https://github.com/pvyswiss/pvyai-coding-agent/commit/382dadf0ca090e7f47c5340cc657d2cc34b5b521))
* **tui:** sticky plan panel + specialist task cards + task table overlay ([#255](https://github.com/pvyswiss/pvyai-coding-agent/issues/255)) ([b237e98](https://github.com/pvyswiss/pvyai-coding-agent/commit/b237e9849a9045b93f833d73919fa80b3b77527d))
* **tui:** surface sandbox permission events ([#93](https://github.com/pvyswiss/pvyai-coding-agent/issues/93)) ([5b1b7ed](https://github.com/pvyswiss/pvyai-coding-agent/commit/5b1b7edf74d808c5596309ca6f2bb2a9927e20af))
* **tui:** switch splash wordmark to figlet Rebel with two-tone cyan ([f9d6c77](https://github.com/pvyswiss/pvyai-coding-agent/commit/f9d6c773d2309d0f29381e41df3220a571d4a140))
* **tui:** syntax-highlighted code, word-level diffs, reduced-motion gate ([#289](https://github.com/pvyswiss/pvyai-coding-agent/issues/289)) ([efe48ad](https://github.com/pvyswiss/pvyai-coding-agent/commit/efe48ade0c894c9fbc4be31c98d0bdbb7ec3b0a4))
* **tui:** tabbed ask_user questionnaire in the composer with suggested answers ([#327](https://github.com/pvyswiss/pvyai-coding-agent/issues/327)) ([fb3277f](https://github.com/pvyswiss/pvyai-coding-agent/commit/fb3277f46eb19f9fca5dea0e476d7dbea7c03a6e))
* **tui:** transcript polish — left-rule cards, tool-call density, pinned plan ([#269](https://github.com/pvyswiss/pvyai-coding-agent/issues/269)) ([96a851d](https://github.com/pvyswiss/pvyai-coding-agent/commit/96a851df2dce79bd5d0f49c984c614f7c78b1d1a))
* **tui:** use figlet Electronic wordmark for ZERO splash logo ([24498d3](https://github.com/pvyswiss/pvyai-coding-agent/commit/24498d3bf890904bf2fb3e7e255b0b84b42bbe36))
* unify go runtime provider contracts ([a222765](https://github.com/pvyswiss/pvyai-coding-agent/commit/a222765a290c039a92405f2ae5405fc997106de2))
* **update:** add zero upgrade command to apply self-updates ([#461](https://github.com/pvyswiss/pvyai-coding-agent/issues/461)) ([5f36349](https://github.com/pvyswiss/pvyai-coding-agent/commit/5f36349c1884e81fa9bc66bb5fe813b627e897b7))
* **update:** expose release check target options ([#96](https://github.com/pvyswiss/pvyai-coding-agent/issues/96)) ([3e4add9](https://github.com/pvyswiss/pvyai-coding-agent/commit/3e4add95ab544190e6fed250a39dc931cfec38b7))
* **update:** verify release metadata by target ([#97](https://github.com/pvyswiss/pvyai-coding-agent/issues/97)) ([699df01](https://github.com/pvyswiss/pvyai-coding-agent/commit/699df010e2a5e7b713a4d6ba0dd0ea728f818649))
* wire Go CLI to real provider flow ([#54](https://github.com/pvyswiss/pvyai-coding-agent/issues/54)) ([95700c0](https://github.com/pvyswiss/pvyai-coding-agent/commit/95700c089826063dfb590109bdfb6c43537c0234))
* **zerocommands:** add typed MCP, hooks, and plugins snapshots ([#90](https://github.com/pvyswiss/pvyai-coding-agent/issues/90)) ([1ddbe9d](https://github.com/pvyswiss/pvyai-coding-agent/commit/1ddbe9def0b7189e49ed9bfd1a2f3a371cef907f))
* **zerocommands:** add typed sandbox policy, risk, violation, and decision snapshots ([#92](https://github.com/pvyswiss/pvyai-coding-agent/issues/92)) ([0c77dad](https://github.com/pvyswiss/pvyai-coding-agent/commit/0c77dad1905beb5cf9476b49aa064ac35bd773c2))
* **zerocommands:** add typed SandboxGrantSnapshot contract ([#89](https://github.com/pvyswiss/pvyai-coding-agent/issues/89)) ([b70fea7](https://github.com/pvyswiss/pvyai-coding-agent/commit/b70fea7d6ffd4ef085aef71cb95997ce960d86e9))


### Bug Fixes

* **action:** keep provider key scoped to zero step ([#448](https://github.com/pvyswiss/pvyai-coding-agent/issues/448)) ([407a927](https://github.com/pvyswiss/pvyai-coding-agent/commit/407a92739ff508cba32d2c12b3f36f0efcdd54c3))
* add android platform support for Termux npm install ([#455](https://github.com/pvyswiss/pvyai-coding-agent/issues/455)) ([9bd93c6](https://github.com/pvyswiss/pvyai-coding-agent/commit/9bd93c62f8d57fb74057284aa66a1b6e1429dcdd)), closes [#449](https://github.com/pvyswiss/pvyai-coding-agent/issues/449)
* address coderabbit review feedback ([e9b3d7b](https://github.com/pvyswiss/pvyai-coding-agent/commit/e9b3d7b331c0d9d43644bfe22c6e5d3e01dd39bc))
* address MCP review feedback ([4e19bd9](https://github.com/pvyswiss/pvyai-coding-agent/commit/4e19bd992fcf833cbccb12abaac28080fe0bf14e))
* address session lineage review feedback ([53097ab](https://github.com/pvyswiss/pvyai-coding-agent/commit/53097ab2206fea8c4977d0c70e098ff3d19289bd))
* agent runtime resilience — stream stalls, weak-model tool args, and sub-agent reliability ([#349](https://github.com/pvyswiss/pvyai-coding-agent/issues/349)) ([7fb0092](https://github.com/pvyswiss/pvyai-coding-agent/commit/7fb0092de747559c34d45da298a601c7254005d0))
* **agent:** gate headless run completion on honesty (no false-success on no-tool-call / self-reported-incomplete turns) ([#325](https://github.com/pvyswiss/pvyai-coding-agent/issues/325)) ([f8ac5ec](https://github.com/pvyswiss/pvyai-coding-agent/commit/f8ac5eceb090c10278612218fc9528e55eff55c2))
* **agent:** reject a malformed additional_permissions payload before prompting ([#453](https://github.com/pvyswiss/pvyai-coding-agent/issues/453)) ([e4f760e](https://github.com/pvyswiss/pvyai-coding-agent/commit/e4f760ee8bd57299cd2fcb37e8e23130037c2607))
* align tui with compact agent stream ([4c37ea5](https://github.com/pvyswiss/pvyai-coding-agent/commit/4c37ea558c40a859677607d4a4cbe8c2c113f652))
* allow non-TLS connections to private-network provider endpoints ([#444](https://github.com/pvyswiss/pvyai-coding-agent/issues/444)) ([1d86384](https://github.com/pvyswiss/pvyai-coding-agent/commit/1d8638466ca31517eb9db2b9353d3dce1cbeeabc))
* **auth:** propagate credentials to every provider-build surface and pin children to the live provider ([#366](https://github.com/pvyswiss/pvyai-coding-agent/issues/366)) ([6e0a665](https://github.com/pvyswiss/pvyai-coding-agent/commit/6e0a665118fe0e09c4b07d482dd18f86045acd2b))
* **auth:** route zero auth login chatgpt to the dedicated ChatGPT flow ([#443](https://github.com/pvyswiss/pvyai-coding-agent/issues/443)) ([305a62c](https://github.com/pvyswiss/pvyai-coding-agent/commit/305a62c954ca6cec00bc58d5398f933415156aff))
* behavioral-audit follow-ups (exec -p, cron/version/specialist UX, update_plan coercion, Makefile) ([#286](https://github.com/pvyswiss/pvyai-coding-agent/issues/286)) ([bd7cc56](https://github.com/pvyswiss/pvyai-coding-agent/commit/bd7cc5665fba310406f78bded577023fe0ee1918))
* cap unbounded exec_command output buffer to prevent OOM kill ([#353](https://github.com/pvyswiss/pvyai-coding-agent/issues/353)) ([2a8802e](https://github.com/pvyswiss/pvyai-coding-agent/commit/2a8802e39ccdf7dd866fe07cef4be1e0ca691ab2))
* **config:** fall back to a usable saved provider instead of forcing full re-onboarding ([#410](https://github.com/pvyswiss/pvyai-coding-agent/issues/410)) ([c60ad87](https://github.com/pvyswiss/pvyai-coding-agent/commit/c60ad8729f79bb841114d352ee2d2fe29d5d0e41))
* **config:** let a gateway ANTHROPIC_BASE_URL resolve as anthropic-compatible ([#497](https://github.com/pvyswiss/pvyai-coding-agent/issues/497)) ([30dd7c3](https://github.com/pvyswiss/pvyai-coding-agent/commit/30dd7c3112ad22d42fa12b5addd4e38f4beda42a)), closes [#479](https://github.com/pvyswiss/pvyai-coding-agent/issues/479)
* **config:** skip an unresolvable non-active provider instead of crashing startup ([#283](https://github.com/pvyswiss/pvyai-coding-agent/issues/283)) ([a2d902c](https://github.com/pvyswiss/pvyai-coding-agent/commit/a2d902c38f13a838ca8afe743a6251809df10112))
* **config:** unbrick first-run setup — default google/anthropic models, enter setup on fixable config errors ([#385](https://github.com/pvyswiss/pvyai-coding-agent/issues/385)) ([72eed06](https://github.com/pvyswiss/pvyai-coding-agent/commit/72eed06b4f94c43d75d31fe54a58d2f566de059e))
* **config:** use ~/.config on macOS and enter setup when no provider ([#371](https://github.com/pvyswiss/pvyai-coding-agent/issues/371)) ([#372](https://github.com/pvyswiss/pvyai-coding-agent/issues/372)) ([027a8f2](https://github.com/pvyswiss/pvyai-coding-agent/commit/027a8f2768b17b89f5c8270887f156e2ccda69ea))
* cut the content-stall dead-wait and auto-retry a mid-tool-call stall ([#362](https://github.com/pvyswiss/pvyai-coding-agent/issues/362)) ([02f0c09](https://github.com/pvyswiss/pvyai-coding-agent/commit/02f0c0929bec65f5fc41bdcab3eb518d2b0b329a))
* **docs:** rename AGENTS.MD &gt; AGENTS.md ([#438](https://github.com/pvyswiss/pvyai-coding-agent/issues/438)) ([4266baf](https://github.com/pvyswiss/pvyai-coding-agent/commit/4266baf222df583ed2078b776687f12d496475b5))
* **gemini:** strip unsupported JSON Schema fields from tool declarations ([#374](https://github.com/pvyswiss/pvyai-coding-agent/issues/374)) ([39e7100](https://github.com/pvyswiss/pvyai-coding-agent/commit/39e7100674150144a1152e3110c64c7cf0321d64)), closes [#373](https://github.com/pvyswiss/pvyai-coding-agent/issues/373)
* gracefully close stdio MCP clients ([ec639c1](https://github.com/pvyswiss/pvyai-coding-agent/commit/ec639c10e445e8bc43e1fd04e86bb0415dcaae71))
* guard concurrent MCP client close ([87bf655](https://github.com/pvyswiss/pvyai-coding-agent/commit/87bf65569fc80bba328972b196ee04d71dd38b76))
* **install:** persist install dir to user PATH on Windows ([#407](https://github.com/pvyswiss/pvyai-coding-agent/issues/407)) ([bdb1b0e](https://github.com/pvyswiss/pvyai-coding-agent/commit/bdb1b0ecd15859b1712a6037d296dace7f9c3c3f))
* lazy load tui for cli commands ([048d8de](https://github.com/pvyswiss/pvyai-coding-agent/commit/048d8decd51e696ed8866e0bdde2961031fb507d))
* Manual renamed Sessions remain untouched from Retitle ([8a3bc37](https://github.com/pvyswiss/pvyai-coding-agent/commit/8a3bc37136691f3413943a23dc7a9e2c798d35d4))
* **mcp:** block cross-origin credential redirects ([#396](https://github.com/pvyswiss/pvyai-coding-agent/issues/396)) ([f915f70](https://github.com/pvyswiss/pvyai-coding-agent/commit/f915f70e5a3096e2419fa8d961a0f84a626fa4a9))
* **model:** expose reasoning effort for GPT-5 / Codex / o-series ([#330](https://github.com/pvyswiss/pvyai-coding-agent/issues/330)) ([8ea8831](https://github.com/pvyswiss/pvyai-coding-agent/commit/8ea8831c30e0ff1af157d8a7d2d1305921ea9eb0))
* **modelregistry:** one source of truth for a model's reasoning efforts ([#335](https://github.com/pvyswiss/pvyai-coding-agent/issues/335)) ([aa45eda](https://github.com/pvyswiss/pvyai-coding-agent/commit/aa45edacc06efe017963fac9f4fb10c8622a002e))
* narrow TUI usage fallback ([1864c42](https://github.com/pvyswiss/pvyai-coding-agent/commit/1864c42278516c69df45d9ccea6f3c08d1ff3bc3))
* **oauth:** make ChatGPT (Codex) login work end-to-end ([#264](https://github.com/pvyswiss/pvyai-coding-agent/issues/264)) ([72a4229](https://github.com/pvyswiss/pvyai-coding-agent/commit/72a4229c2a8124f278fc9fffd1d328c4c1c263f7))
* **oauth:** make xAI / Hugging Face OAuth actually launch from the wizard ([#355](https://github.com/pvyswiss/pvyai-coding-agent/issues/355)) ([c4c14d6](https://github.com/pvyswiss/pvyai-coding-agent/commit/c4c14d62d4a20ae6191e544c796895a84f7de713))
* **oauth:** treat Windows ERROR_ACCESS_DENIED as lock contention in createSecretFile ([#445](https://github.com/pvyswiss/pvyai-coding-agent/issues/445)) ([d05e914](https://github.com/pvyswiss/pvyai-coding-agent/commit/d05e9148a7f79f67d1d3c31fca2775f21fbd331e))
* **openai:** always send message content so strict OpenAI-compatible servers accept it ([#343](https://github.com/pvyswiss/pvyai-coding-agent/issues/343)) ([cfd065b](https://github.com/pvyswiss/pvyai-coding-agent/commit/cfd065b565100058905a25b8c683de41da99ea89))
* **openai:** forward reasoning effort to the Codex Responses API ([#336](https://github.com/pvyswiss/pvyai-coding-agent/issues/336)) ([d99ac4d](https://github.com/pvyswiss/pvyai-coding-agent/commit/d99ac4d1e334c92555b09fb51fb6cab302a3bdb5))
* **openai:** handle Ollama reasoning stream deltas ([#486](https://github.com/pvyswiss/pvyai-coding-agent/issues/486)) ([f6c0606](https://github.com/pvyswiss/pvyai-coding-agent/commit/f6c060631e18e082dda24cc4dc0903c31c2120d6))
* **opengateway:** correct gateway host, make it the recommended default ([#341](https://github.com/pvyswiss/pvyai-coding-agent/issues/341)) ([072f387](https://github.com/pvyswiss/pvyai-coding-agent/commit/072f3877b52015b90b66ed46b43326ea71070b75))
* preserve conversation context in exec prompts ([#460](https://github.com/pvyswiss/pvyai-coding-agent/issues/460)) ([949ee43](https://github.com/pvyswiss/pvyai-coding-agent/commit/949ee43f71e5cb7fab4695c5cb7b442fe4ecfbf7))
* preserve MCP response close errors ([a09a721](https://github.com/pvyswiss/pvyai-coding-agent/commit/a09a7213b60aab67941f7255903dca2ffddd2f80))
* **provider-wizard:** allow multiple custom OpenAI-compatible providers ([#403](https://github.com/pvyswiss/pvyai-coding-agent/issues/403)) ([3fbbd28](https://github.com/pvyswiss/pvyai-coding-agent/commit/3fbbd28e4c586822cc4312c86232d94befe56e87))
* **providerio:** bound heartbeat-but-no-output streams so they can't hang forever ([#347](https://github.com/pvyswiss/pvyai-coding-agent/issues/347)) ([02a3d12](https://github.com/pvyswiss/pvyai-coding-agent/commit/02a3d12f3468eba44ef94749a353755332b8c569))
* **providerio:** disable HTTP keep-alive reuse on macOS to defeat degraded-connection stalls ([#360](https://github.com/pvyswiss/pvyai-coding-agent/issues/360)) ([85273ea](https://github.com/pvyswiss/pvyai-coding-agent/commit/85273ea72d86a2c9886c20d580c86214728b5af8))
* **providers:** clear error when a model host is unreachable ([#233](https://github.com/pvyswiss/pvyai-coding-agent/issues/233)) ([7c36805](https://github.com/pvyswiss/pvyai-coding-agent/commit/7c36805505f76dae42bfefcb7501891a72d1acdc))
* **providers:** make the stream idle timeout global, configurable, and less aggressive ([#285](https://github.com/pvyswiss/pvyai-coding-agent/issues/285)) ([29246e5](https://github.com/pvyswiss/pvyai-coding-agent/commit/29246e56592f5125a0af60abd6122a9a799e44a9))
* redesign tui as agent console ([b3349dd](https://github.com/pvyswiss/pvyai-coding-agent/commit/b3349dd32bdbdabaae1d7b9f3576e63bc3abd87c))
* reject empty exec prompts ([b8e194d](https://github.com/pvyswiss/pvyai-coding-agent/commit/b8e194d34d6e5d8849345990426fccc8424b5cd7))
* resolve 15 findings from the 2026-06-20 deep audit ([#281](https://github.com/pvyswiss/pvyai-coding-agent/issues/281)) ([f7ee189](https://github.com/pvyswiss/pvyai-coding-agent/commit/f7ee189ba5eb7a04f69c6d792a6f35521b6f8917))
* resolve product/UX audit findings for OSS launch ([#282](https://github.com/pvyswiss/pvyai-coding-agent/issues/282)) ([9047977](https://github.com/pvyswiss/pvyai-coding-agent/commit/904797789569e5524fa9686e120801905f6cb3b7))
* restore startup shortcut hints ([c586502](https://github.com/pvyswiss/pvyai-coding-agent/commit/c5865028a78586a3b48ee339878f65380532949e))
* **sandbox:** fix nested pipe creation under the Windows restricted token ([#456](https://github.com/pvyswiss/pvyai-coding-agent/issues/456)) ([563a6db](https://github.com/pvyswiss/pvyai-coding-agent/commit/563a6dbe91e65d5daeefd7626e8a77e30a6d8fb2))
* **sandbox:** gate /tmp test assertions on GOOS, not path existence ([#426](https://github.com/pvyswiss/pvyai-coding-agent/issues/426)) ([f653dca](https://github.com/pvyswiss/pvyai-coding-agent/commit/f653dcac363fb69ad7be5b35e6e0fa6d2bce476d))
* **sandbox:** grant ancestor metadata so cd into the workspace works ([#302](https://github.com/pvyswiss/pvyai-coding-agent/issues/302)) ([b776a92](https://github.com/pvyswiss/pvyai-coding-agent/commit/b776a9281fd1d92c407bcc1582eeef57afcb91bd))
* **sandbox:** let macOS sandbox run Homebrew/usr-local tools (node, python3) ([#296](https://github.com/pvyswiss/pvyai-coding-agent/issues/296)) ([b46e952](https://github.com/pvyswiss/pvyai-coding-agent/commit/b46e952282b9784e63e95484f161d1226a33fc42))
* **sandbox:** retry approved network shell commands ([#301](https://github.com/pvyswiss/pvyai-coding-agent/issues/301)) ([6be50cd](https://github.com/pvyswiss/pvyai-coding-agent/commit/6be50cd3a75e53bfc435e2e5f36500ee153b0a87))
* **sandbox:** self-dispatch the Windows sandbox helpers so they work in dev ([#280](https://github.com/pvyswiss/pvyai-coding-agent/issues/280)) ([990e3c4](https://github.com/pvyswiss/pvyai-coding-agent/commit/990e3c439ba67f8687340e73862811cf25613c27))
* **sandbox:** self-heal a corrupt unelevated setup marker ([#437](https://github.com/pvyswiss/pvyai-coding-agent/issues/437)) ([8d0c5fe](https://github.com/pvyswiss/pvyai-coding-agent/commit/8d0c5feccb8bdbfb015df0508aa6e3bcbd1fd0e8))
* **sandbox:** stop bricking Windows when the sandbox isn't set up (degrade like Linux/macOS) ([#295](https://github.com/pvyswiss/pvyai-coding-agent/issues/295)) ([ffaa3a9](https://github.com/pvyswiss/pvyai-coding-agent/commit/ffaa3a9a48ef6aff9d1cdfd2522cf46cf76c9473))
* **sandbox:** stop the interactive guard from blocking 'node --version' & friends ([#294](https://github.com/pvyswiss/pvyai-coding-agent/issues/294)) ([3c2cd1d](https://github.com/pvyswiss/pvyai-coding-agent/commit/3c2cd1d491dd9f7f65c2699c6fa0066bcbec1b94))
* **specialist:** cap max specialist nesting depth ([#491](https://github.com/pvyswiss/pvyai-coding-agent/issues/491)) ([177442c](https://github.com/pvyswiss/pvyai-coding-agent/commit/177442cfe4015bd8df04cc9894f98b468ee796d4))
* stream Codex reasoning + drop swarm agents on their own completion ([#345](https://github.com/pvyswiss/pvyai-coding-agent/issues/345)) ([83a371d](https://github.com/pvyswiss/pvyai-coding-agent/commit/83a371d669892a1e4188ec1339765871e9ced22f))
* strengthen tui shell visual hierarchy ([62d322a](https://github.com/pvyswiss/pvyai-coding-agent/commit/62d322a6384eb93471770c77e9cba9e83b950797))
* **swarm:** constrain agent_type to the registered roster via a dynamic enum ([#344](https://github.com/pvyswiss/pvyai-coding-agent/issues/344)) ([84823d7](https://github.com/pvyswiss/pvyai-coding-agent/commit/84823d7be15794147afbcdb9b56c140d846c6733))
* **swarm:** make swarm_collect wait for members instead of polling ([#310](https://github.com/pvyswiss/pvyai-coding-agent/issues/310)) ([258f1a3](https://github.com/pvyswiss/pvyai-coding-agent/commit/258f1a31c9004774094aa994592d20dacef3a603))
* Termux/Android support — PRoot scroll, SIGSYS sandbox, build docs ([#509](https://github.com/pvyswiss/pvyai-coding-agent/issues/509)) ([0f69d99](https://github.com/pvyswiss/pvyai-coding-agent/commit/0f69d995e9b586b774f66c066b21abab5e03024a))
* **tools:** CRLF line ending mismatch in edit_file tool on Windows ([#378](https://github.com/pvyswiss/pvyai-coding-agent/issues/378)) ([33dc7ae](https://github.com/pvyswiss/pvyai-coding-agent/commit/33dc7ae2cc82c5389675531e1416856dae7151ce))
* **tools:** fix cmd.exe /S/C corrupting commands with embedded quotes ([#465](https://github.com/pvyswiss/pvyai-coding-agent/issues/465)) ([190241b](https://github.com/pvyswiss/pvyai-coding-agent/commit/190241bd593f43211b766e0b13c8e89802d4bb37))
* **tools:** flag piped POSIX utilities before running on Windows ([#412](https://github.com/pvyswiss/pvyai-coding-agent/issues/412)) ([5658a36](https://github.com/pvyswiss/pvyai-coding-agent/commit/5658a366274fc59a9d5336b06a21019c9c25cbf1))
* **tools:** make grep and glob respect run cancellation ([#464](https://github.com/pvyswiss/pvyai-coding-agent/issues/464)) ([ba6c026](https://github.com/pvyswiss/pvyai-coding-agent/commit/ba6c0264697b7d7ed479f6e782fba9700a481e3d))
* **tools:** require permission before web_search requests ([#382](https://github.com/pvyswiss/pvyai-coding-agent/issues/382)) ([960db96](https://github.com/pvyswiss/pvyai-coding-agent/commit/960db9660e4e31dc588fe8f7d6f116ff5e225566))
* **tui:** cap the expanded live reasoning to ~half screen ([#303](https://github.com/pvyswiss/pvyai-coding-agent/issues/303)) ([c498a8b](https://github.com/pvyswiss/pvyai-coding-agent/commit/c498a8b8705ed6d94973196d16e87817a365b3d3))
* **tui:** compose help overlay through the viewport overlay pipeline ([#421](https://github.com/pvyswiss/pvyai-coding-agent/issues/421)) ([5b2b4de](https://github.com/pvyswiss/pvyai-coding-agent/commit/5b2b4dea1aaf9e0f68baa25e97e83296fb17b1a2))
* **tui:** context-fill gauge for custom/local Ollama models + related fixes ([#356](https://github.com/pvyswiss/pvyai-coding-agent/issues/356)) ([c0254e3](https://github.com/pvyswiss/pvyai-coding-agent/commit/c0254e3266ff6d8a7125b06c96788af427598104))
* **tui:** explicitly request the initial window size at startup ([#358](https://github.com/pvyswiss/pvyai-coding-agent/issues/358)) ([6866907](https://github.com/pvyswiss/pvyai-coding-agent/commit/68669073c2b1c95670b0ea9daf78d6b589548c56))
* **tui:** give command-info screens real contrast (not flat grey) ([#272](https://github.com/pvyswiss/pvyai-coding-agent/issues/272)) ([17e7be2](https://github.com/pvyswiss/pvyai-coding-agent/commit/17e7be2743b7a611970539cac9a81f1568e59bc9))
* **tui:** keep running swarm subagents in the AGENTS sidebar ([#309](https://github.com/pvyswiss/pvyai-coding-agent/issues/309)) ([4a5a417](https://github.com/pvyswiss/pvyai-coding-agent/commit/4a5a417156341b6f742e22e40db719216efeda1e))
* **tui:** keep the profile name on /model switch so the stored key resolves ([#441](https://github.com/pvyswiss/pvyai-coding-agent/issues/441)) ([9134148](https://github.com/pvyswiss/pvyai-coding-agent/commit/9134148f4df3e4e556fba6c2f8babfdf6fcfeee1)), closes [#440](https://github.com/pvyswiss/pvyai-coding-agent/issues/440)
* **tui:** make the stall watchdog visible before it fires, not just after ([#357](https://github.com/pvyswiss/pvyai-coding-agent/issues/357)) ([33410d8](https://github.com/pvyswiss/pvyai-coding-agent/commit/33410d87304f05a4ed9b778a5243183a4285b075))
* **tui:** paint transcript selection once, not twice ([#328](https://github.com/pvyswiss/pvyai-coding-agent/issues/328)) ([95aaca6](https://github.com/pvyswiss/pvyai-coding-agent/commit/95aaca6055f1fff44c81de54a890161e50ca4d7f))
* **tui:** require a second Esc press to cancel a running turn ([#352](https://github.com/pvyswiss/pvyai-coding-agent/issues/352)) ([bcc65da](https://github.com/pvyswiss/pvyai-coding-agent/commit/bcc65da2a0ce0cc84cbd7c53ab900885d12e7fa3))
* **tui:** resolve every permission request so the agent can't deadlock ([#397](https://github.com/pvyswiss/pvyai-coding-agent/issues/397)) ([952788f](https://github.com/pvyswiss/pvyai-coding-agent/commit/952788f72d32957659fe004521fcc8372b9ba9b4))
* **tui:** selectable + clickable permission popup ([#232](https://github.com/pvyswiss/pvyai-coding-agent/issues/232)) ([b1cf248](https://github.com/pvyswiss/pvyai-coding-agent/commit/b1cf24831d90fcc1b561ed600355abbe704bf62f))
* **tui:** show an M suffix for million-scale token counts ([#457](https://github.com/pvyswiss/pvyai-coding-agent/issues/457)) ([0562e3b](https://github.com/pvyswiss/pvyai-coding-agent/commit/0562e3bef7df2328610a48a1e81632a8da4aec64))
* **tui:** show the queued-message preview above the composer, not below ([#361](https://github.com/pvyswiss/pvyai-coding-agent/issues/361)) ([3fde03a](https://github.com/pvyswiss/pvyai-coding-agent/commit/3fde03a9000f8789445199b6d39c1f5e8d15b086))
* **tui:** stop ANSI escape leaks + drop redundant update_plan cards ([#316](https://github.com/pvyswiss/pvyai-coding-agent/issues/316)) ([01fed03](https://github.com/pvyswiss/pvyai-coding-agent/commit/01fed03a5903f6bebae801fdb4eb091c8bce6128))
* **tui:** stop showing the token count on both sides ([#306](https://github.com/pvyswiss/pvyai-coding-agent/issues/306)) ([bc5c24e](https://github.com/pvyswiss/pvyai-coding-agent/commit/bc5c24ef873fcda9ed8a1eb71853960b3f85df30))
* **tui:** sub-agents inherit the parent's live provider after a /model switch ([#348](https://github.com/pvyswiss/pvyai-coding-agent/issues/348)) ([5aced72](https://github.com/pvyswiss/pvyai-coding-agent/commit/5aced721b31634e033f7f952b64e76c3cf3bfb9f))
* **tui:** subchat mouse selection, hover highlighting, and drag-to-edge auto-scroll ([#350](https://github.com/pvyswiss/pvyai-coding-agent/issues/350)) ([1c65f59](https://github.com/pvyswiss/pvyai-coding-agent/commit/1c65f59be60f21cf16810eb53ef9c647f4925a19))
* **tui:** title /model rows by model name, not the catalog description ([#395](https://github.com/pvyswiss/pvyai-coding-agent/issues/395)) ([cdf9d83](https://github.com/pvyswiss/pvyai-coding-agent/commit/cdf9d839ae57a729f292f36f7c5b0c67b41b288d))
* update ChatGPT Codex OAuth request handling ([#299](https://github.com/pvyswiss/pvyai-coding-agent/issues/299)) ([a2e0f3c](https://github.com/pvyswiss/pvyai-coding-agent/commit/a2e0f3c1feb285c976994a6dff81931b9243ebf7))


### Performance Improvements

* cache TUI model registry ([#496](https://github.com/pvyswiss/pvyai-coding-agent/issues/496)) ([e7d88b4](https://github.com/pvyswiss/pvyai-coding-agent/commit/e7d88b4b518049733da25a8447c00144bd1da518))
* **tools:** lower default deferThreshold to stop eager MCP-schema token waste ([#241](https://github.com/pvyswiss/pvyai-coding-agent/issues/241)) ([b903825](https://github.com/pvyswiss/pvyai-coding-agent/commit/b903825139c6e95350cb11e7b717f4361a5ef860))
* universal tool-output ceiling with spill + async post-edit diagnostics ([#518](https://github.com/pvyswiss/pvyai-coding-agent/issues/518)) ([95ccd5b](https://github.com/pvyswiss/pvyai-coding-agent/commit/95ccd5bc327f6fb464ff0239f7229de789f578dc))


### Reverts

* **local:** restore the original bright-lime accent color ([#351](https://github.com/pvyswiss/pvyai-coding-agent/issues/351)) ([d530520](https://github.com/pvyswiss/pvyai-coding-agent/commit/d53052071a59aa41bd95549c91b6e969cebb051a))

## 0.1.0 (2026-07-02)


### Features

* publish pvyai to npm via release-please ([#367](https://github.com/pvyswiss/pvyai-coding-agent/issues/367)) ([8eccc26](https://github.com/pvyswiss/pvyai-coding-agent/commit/8eccc2669887bc38d35bc16a315c888e4d9ec43a))
* **tui:** FILES sidebar panel with click-to-select and file drill-in ([#365](https://github.com/pvyswiss/pvyai-coding-agent/issues/365)) ([142c548](https://github.com/pvyswiss/pvyai-coding-agent/commit/142c548c89a8652ce300e64ddf1228ee36df7606))


### Bug Fixes

* **auth:** propagate credentials to every provider-build surface and pin children to the live provider ([#366](https://github.com/pvyswiss/pvyai-coding-agent/issues/366)) ([6e0a665](https://github.com/pvyswiss/pvyai-coding-agent/commit/6e0a665118fe0e09c4b07d482dd18f86045acd2b))

## [Unreleased]

### Added
- `SECURITY.md` with a private vulnerability-reporting path, `CODE_OF_CONDUCT.md`, this changelog, and
  GitHub issue/PR templates.
- Interactive `/theme` picker: bare `/theme` opens a popup that live-previews each palette as you move
  and applies on select (Esc reverts).
- Ten built-in color themes alongside the `dark`/`light` built-ins — `dracula`, `nord`, `gruvbox`,
  `tokyo-night`, `catppuccin`, `one-dark`, `solarized-dark`, `rose-pine`, `everforest`, and
  `solarized-light` — selectable via `/theme <name>`, `--theme <name>`, or `PVYAI_THEME`. Every palette
  is contrast-audited to WCAG AA. The built-in light theme was reworked for legibility.
- `--theme <name>` flag for the TUI, accepting `auto` or any registered theme (previously only the
  `PVYAI_THEME` env var existed).
- "Accessibility / Appearance" section in the README documenting `NO_COLOR`, `PVYAI_THEME`, `/theme`,
  and `PVYAI_NO_FADE`.

### Changed
- Provider connectivity health checks now allow loopback hosts for explicitly user-configured local
  providers (Ollama / LM Studio), so the keyless local-model path verifies instead of failing with
  "localhost hosts are blocked". The SSRF guard for fetched/remote URLs is unchanged.
- Auth (401/403) errors now show a curated, actionable message pointing at `pvyai auth` / setup; the
  raw upstream body is shown only under a verbose/debug flag.
- No-provider / missing-key errors now point at `pvyai setup` and `pvyai auth`, and distinguish a
  missing key from a rejected key.
- `pvyai doctor` no longer reports "Overall: pass" when no provider credential is configured, and
  formats the missing-language-server list for humans (no raw Go `map[...]`).
- Raised the `faint`/`faintest` theme tokens (and the light-theme accent) to meet WCAG AA contrast for
  the content they carry.
- `NO_COLOR` is now honored for any non-empty value, per the no-color.org spec.

### Removed
- The inert `/input-style` slash command (it had no backend).

### Fixed
- README/`go.mod` Go-version mismatch and other stale public-release docs claims.
