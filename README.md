# lispylicense

[![Report](https://goreportcard.com/badge/github.com/lispysnake/lispylicense)](https://goreportcard.com/report/github.com/lispysnake/lispylicense) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)


Work in progress project to handle product license generation and
management.

Our initial scope is to facilitate the pregeneration of 1,000 license
keys and allocate them to Lispy Snake 2D Tite Lifetime License holders.

### Initial alpha

Scope for initial alpha is a CLI only tool (rather than fleshed server)
that will allocate a UUID on the CLI, stash to DB and store it. In future
we'll add job-based UUID allocation in a microservice.

### Key Format

Nothing special here. We simply generate a UUID, and ensure it is completely
unique due to the Birthday Problem. At that point it becomes a valid license
key, not due to format, simply because its a key our DB will recognise.

Additionally the key will require knowledge of the email address to
be 'valid'.

### Proposal

Microservice to handle license key generation and allocation.
Allocation is performed based on an ACCOUNT_ID field, to allow us to
perform a future migration from the initial signup email to a hub
account ID.

For now, this is a manual process of allocation, but in future we may
work out something more automatic:

 - Acknowledge transaction begin
 - Send an email saying we're waiting for funds to clear
 - Acknowledge transaction completion
 - Allocate specific license key and email the user
 - Stash to DB

### Security

To ensure minimal security risk, we recommend running the license
server internally (behind the firewall). Then the account service,
userfacing, can internally be updated with license state from this
service.

### License

Copyright © 2019 Lispy Snake, Ltd.

`lispylicense` is available under the terms of the `Apache-2.0` license
