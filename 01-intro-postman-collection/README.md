# Intro Postman Collection

## What's in this?

- Authentication Examples using OAuth
- Data requests to export orders, products and categories
- Combined Auth & Data Requests for orders and products

## Quick Start

You can import this Collection in Postman under Workspace >> Import

## Configuration / Global Variables

It expects some global variables to be set in Postman. Global variables
are not tied to a Postman Collection or Environment. It is important
to make sure these variables are NOT SHARED in any Postman exports.

- REST_KEY
- REST_SECRET

You can see these being used in the "Oath Authentication" requests
under the "Authorization" tab. These values are unique to the API
user making the call.

## Usage

This collection is comprised of a variety of requests grouped under
three folders:

- Auth Requests ~ contains discreet OAuth2 authentication requests for
  orders, products and categories endpoints.
- Data Requests ~ contains discreet data requests to get orders,
  products and categories. These do not include authentication.
- Auth & Data Requests ~ data requests that include the OAuth2
  authentication as part of the request.

## How do I learn from this

For each request in the collection, you can easily see the endpoint
at the top of the page and review the Authorization and Body tabs
of the request to see what is being sent. The response is also easily
viewed after the request is sent at the bottom of the page.

The response can be fully viewed in the response window on the lower
half of the page. More request details can be viewed in a number of
ways:

1. Review the "Documentation" icon in the right sidebar.
2. Review the "Code" icon in the right sidebar, which can be set
   to the language of your choice and will automatically create a
   code snippet for this particular request.
3. Click the `Console` toggle in the lower left corner to expand
   the console. You can then examine the raw request and response
   headers and bodies by expanding the appropriate arrows.

If your goal is to sucessfully construct an auth request or a data
request, you should look at the discreet examples for each. When
using the data requests, you'll need a valid token before clicking
the "Send" buttons. This collection is set up so that simply executing
the appropriate Auth request will then make data requests automatically
work until that token expires. Auth requests are good for an hour and for
a single endpoing (products, orders OR categories).

For instance, to use the Product-based requests like `Get Products By Range`
you would first click on the `OAuth-Products` endpoint in the Auth folder,
then and click `Send` to authenticate with the API. After that you can
click `Send` on the `Get Products By Range` endpoint and should see a
sucessful response.
