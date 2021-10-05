# CV3 API Examples

various API example calls for use in testing and developing against the CV3 APIs

## Quick Start

There are a variety of client-side technologies and toolkits that can be used to
access the CV3 REST API and exchange data. Here we collect examples so that
integration developers can review working code to get started with their own
integrations.

In most cases we suggest developers utilize the included Postman Collections. It
can be very useful throughout a project to have a quick way to test the
request/response details without the complications of any project-specific code.
In addition, when working with other developers or CV3 Support it can be used to
share code, problems and solutions among folks that may not have access to the
specific tools and technologies you're using in your project.

## CV3 REST API Credentials

Instructions for setting up the API for use are
[here](https://docs.google.com/document/d/1LG8mM1v5gmoIHL66H5szUVyeZz2chUlT6pGUv7y9-uU/edit).
Please note:

- API access credentials and use is limited to a sub-user and you should set permissions correctly for that sub-user
- Review the expiration information on credentials carefully.

## Postman Setup

[Postman](https://www.postman.com) is a popular application used to test and explore various APIs. Many of these
examples are offered as "Postman Collections" that can be easily imported and configured
for use. All of these collections work with the free & paid versions of Postman.

1. If you aren't using it already, [download](https://www.postman.com/downloads/) and install Postman.
2. If you don't have a Workspace you want to use, create a Postman Workspace
3. In the Workspace view, click `Import` in the Collections sidebar
4. Click `Upload` and choose one of our example collections (\*.postman_collection.json files)
5. Configure any necessary Global Postman variables for use with the collection

This last step is explained in more detail in each collection's README.md file. In general it will
involve making sure you have Global Variables set in Postman for the API authentication information.

## What's Here?

- [01-intro-postman-collection](./01-intro-postman-collection) - OAuth requests, basic fething of products and orders
- [02-advanced-postman-collection](02-advanced-postman-collection) - IN PROGRESS more detailed examples, data updates
- [03-hmac-postman-collection](03-hmac-postman-collection) - IN PROGRESS basic examples using HMAC Authentication

## Contributing

Contribute here or just send your Postman collections to CV3 support and we'll combine with what we have already!
