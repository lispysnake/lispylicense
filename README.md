# lispylicense

Work in progress project to handle product license generation and
management.

Our initial scope is to facilitate the pregeneration of 1,000 license
keys and allocate them to Lispy Snake 2D Tite Lifetime License holders.

### Key Format

Nothing special here. We simply generate a UUID, and ensure it is completely
unique due to the Birthday Problem. At that point it becomes a valid license
key, not due to format, simply because its a key our DB will recognise.

Additionally the key will require knowledge of the email address to
be 'valid'.

### Proposal

Simple web service that allows us to generate a set of license keys
on demand, and add new users into the system to allocate their license.
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
