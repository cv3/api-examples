# Intro Postman Collection

## What's in this?

- OAuth Authentication examples
- Get requests for orders, products and categories

## Quick Start

You can import this Collection in Postman under Workspace >> Import

## Configuration / Global Variables

It expects some global variables to be set in your active Postman
environment.

- REST_KEY
- REST_SECRET

You can see these being used in the "Oath Authentication" requests
under the "Authorization" tab. These values are unique to the API
user making the call.

## Usage

This collection uses separate OAuth authentication requests that
can be followed immediately by data requests. For instance, to use
the Product-based requests like `Get Products By Range` you would
first click on the `OAuth-Products` endpoint and click `Send` to
authenticate with the API.

Once this request is complete you will see the successful response
in the response window containing a valid access token. You don't
have to do anything with this token in Postman because the code
under the `Tests` tab will automatically run and make that token
a Collection variable that can be used in the other calls.
You can simply click on the `Get Products By Range`
endpoint and click `Send` to make a data request.

OAuth requests are good for an hour, so you may find calls fail
after a while and you need to go back to the OAuth request and
trigger it again for a new token.

OAuth requests are also good for a single endpoint, so if you
will be working with the entire Collection you likely want to
trigger each Oauth request first (for `/products`, `/orders`, etc.)
and then start working with data.

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
