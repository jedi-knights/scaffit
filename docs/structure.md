# Structuring a Golang Project

This document is going to be very opinionated. It is based on my miniscule experience with Go.  So be warned.

Although what I am proposing is likely to be wrong I am going to try to modularize things as much as possible.  I am going to try to keep things as simple as possible.  I am going to try to keep things as consistent as possible.

bin - a directory to dump the binaries

cmd - a directory to hold the commands
each of these commands shave a main.go file

handlers contains handlers
db contains database access layer
pkg contains business logic
types/common/shared contains enums and shared type definitions
util

cmd/api
cmd/drop
cmd/migrate/main.go
cmd/seed/main.go


go run cmd/seed/main.go

go run cmd/drop/main.go


his makefile references @go in the tasks figuring out what that's about


his db package has one file anmed database

with a CreateDatabase function that creates his database

and an Init function


Look at 

github.com/uptrace/bun
github.com/joho/godotenv
