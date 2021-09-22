# HMAC Authentication Testing

This collection is under development and is not currently working. It appears to
not be creating the correct signature.

## Configuration

This collection requires the following Globals to be set in Postman's active environment:

- REST_KEY
- REST_SECRET

## Quick Start

Open up the `HMAC-GetOrdersByRange-v2` endpoint and review the Pre-request Script tab.

## Details

Current HMAC documentation is [here](https://docs.google.com/document/d/1-5jA_x_r4eLsnVABeqQto9BPjPoevnUcMD205gcV8bk/edit)

Here's some PHP code to construct a signature, this is what we are trying to duplicate
in Javascript in the Pre-requeste Script.

```php
$encryption = 'AES-256-CBC';
$ivlen = openssl_cipher_iv_length($encryption);
$iv = openssl_random_pseudo_bytes($ivlen);
$enc_raw = openssl_encrypt($service_id, $encryption, $service_token, $options=0, $iv);
$hmac = hash_hmac('sha256',$enc_raw, $service_token, $as_binary=true);
$signature = base64_encode($iv.$hmac.$enc_raw);
$json['auth'] = array('signature'=>$signature);
```

the v2 endpoint is my latest attempt and the active one we are trying to fix. It is trying
to follow the specific PHP code above.

I left v1 because that followed the published documentation but also doesn't work.
