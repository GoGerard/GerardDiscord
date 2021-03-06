# GerardDiscord
A modern bot for Discord - The Discord bot.

[![Build Status](https://travis-ci.org/GoGerard/GerardDiscord.svg)](https://travis-ci.org/GoGerard/GerardDiscord)
[![Go Report](http://goreportcard.com/badge/GoGerard/GerardDiscord)](http://goreportcard.com/report/GoGerard/GerardDiscord)
[![Issues](https://img.shields.io/github/issues/GoGerard/GerardDiscord.svg)](https://github.com/GoGerard/GerardDiscord/issues)
[![Coverage Status](https://coveralls.io/repos/GoGerard/GerardDiscord/badge.svg?branch=master&service=github)](https://coveralls.io/github/GoGerard/GerardDiscord?branch=master)

----------

## GoGerard

Project details can be found on the [main repo.](https://github.com/GoGerard/GoGerard)


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
