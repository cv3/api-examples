# CV3 Advanaced API Examples

This collection is under development and doesn't contain much yet.

## Configuration

The following Globals are required to be set in Postman:

- REST_KEY
- REST_SECRET
- OAUTH2_TOKEN (not required, but this collection will create it)

For more Postman details see [general usage](../). and
[specific CV3 usage](../01-intro-postman-collection/).

## Quick Start

Open up the various requests in the Collection and click `Send`

## How does Auth Work?

Requests are grouped into Postman Folders based on endpoint, and
each folder has a Pre-request Script that runs prior to every
request inside the folder. This script authenticates
using our OAuth2 service. So while using this collection you
don't have to worry about authentication at all.

If you want to see individual authentication requests, check
out our [Intro Collection](../01-intro-postman-collection/)
