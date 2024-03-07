# Godgifu
2024 New Technologies mega-project. Godgifu is a worldbuilding themed around the Galactic Imperium from the 4X strategy game Stellaris.

## Description


## Goals
Godgifu's main goal is to be an *Overengineered Project*, meaning that it include technologies and libraries that are not necessary, but that I am interested in learning. It will prmarily feature new emerging technologies, and technologies that I have not worked with before. Currently, it will try to be an "everything app" featuring messaging and video chat, file sharing, and forum commenting.

### Overengineered by Design
Godgifu will include

- a user interface written in TypeScript using [Svelte](https://svelte.dev/) and [SvelteKit](https://kit.svelte.dev/) (only the neswest tech for the future). A browser version will be hosted online, with cross-platform Desktop clients also avaliable built using [Tauri](https://tauri.app/).
- Cypress and/or Playwright for Client E2E testing
- Desktop client will feature an offline mode backup by an SQLite database.
- a web server written in [Golang](https://go.dev/) using [Echo](https://echo.labstack.com/) as the web framework.
-a possible future rewrite of the web server using [Rust](https://www.rust-lang.org/) with the [Tokio runtime](https://tokio.rs/) and built using either [Actix-Web](https://actix.rs/docs/) or [Axum](https://docs.rs/axum/latest/axum/index.html#) (obviously everything in the future may be written in Rust).
- Redis caching.
- [Capt'n Proto](https://capnproto.org/), a binary messaging schema, will be used for Client-to-Server communication to keep data message sizes small (and to learn an alternative to JSON).
- One database will be used, [PostgreSQL](https://www.postgresql.org/docs/) for both relational data, and JSON messaging data. A second database might be used for either prototyping, or to increase the project complexity. If so it will be [CouchDB](https://docs.couchdb.org/en/stable/).
- Docker for building containers with multiple avenues of stress testing.
- Instead of JWTs token verification will use [PASETOs](https://paseto.io/).
- HTTPS for connections.
- Multi-Factor Authentication during the Login process.
- Application analytics software.
