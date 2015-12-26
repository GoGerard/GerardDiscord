# GerardDiscord
A modern bot for Discord - The Discord bot.

[![Build Status](https://travis-ci.org/GoGerard/GerardDiscord.svg)](https://travis-ci.org/GoGerard/GerardDiscord)
[![Go Report](http://goreportcard.com/badge/GoGerard/GerardDiscord)](http://goreportcard.com/report/GoGerard/GerardDiscord)
[![Issues](https://img.shields.io/github/issues/GoGerard/GerardDiscord.svg)](https://github.com/GoGerard/GerardDiscord/issues)

----------

## GoGerard
GoGerard is an opensource project that focuses on easy to adapt, community-driven chatbots for [Discord](https://discordapp.com/).

The application is written in three separated parts, which are all replaceable to adapt to your project's needs.

 - [GerardDiscord](https://github.com/GoGerard/GerardDiscord) - A client, written in Golang, that communicates with the Discord API.
 - [GerardAPI](https://github.com/GoGerard/GerardAPI) - An API Server, written in Golang, that is used to communicate with the client and database(s)
 - [GerardJS](https://github.com/GoGerard/GerardJS) - A web interface, powered by AngularJS, that serves the API to its end-users.

Note that in the current state the project is nowhere finished, dependent on unstable external libraries,  and breakable changes to the project will happen till a future release.

----------

### GerardDiscord ###

The Discord Client is written in Golang and is currently depended on the following libraries:

 1. [Gorm](https://github.com/jinzhu/gorm) - 'The fantastic ORM library for Golang, aims to be developer friendly.'
 2. [discordgo](https://github.com/bwmarrin/discordgo) - 'Golang bindings for Discord '


**How to use?**

 1. Clone repo
 2. Duplicate conf_sample.json to conf.json
 3. Fill in configuration with Discord account details
 4. `$ docker run --name POSTGRESCONTAINERNAME -e POSTGRES_PASSWORD=mysecretpassword -d postgres`
 5. Open terminal in project folder
 6. `$ docker build -t gerarddiscord .`
 7. `$ docker run -it --rm --name containername --link POSTGRESCONTAINERNAME:postgres  gerarddiscord`
