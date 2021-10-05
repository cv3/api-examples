/**
 *@NApiVersion 2.x
 *@NScriptType Suitelet
 *@NModuleScope SameAccount
 */
define(["N/crypto", "N/encode", "N/https"], function (crypto, encode, https) {
  function onRequest(params) {
    var output = new Object();
    var html = "";

    const key = "{CV3_KEY}";
    const secret = "{CV3_SECRET}";

    var data = "grant_type=client_credentials&scope=products";

    /////////////////////////////////// AUTH REQUEST
    var header = [];
    header["Content-Type"] = "application/x-www-form-urlencoded";

    var auth_string = key + ":" + secret;

    var auth_token = encode.convert({
      string: auth_string,
      inputEncoding: encode.Encoding.UTF_8,
      outputEncoding: encode.Encoding.BASE_64,
    });

    header["Authorization"] = "Basic " + auth_token;

    var cv3response = https.post({
      url: "https://service.commercev3.com/rest/oauth2/v2.0/token",
      body: data,
      headers: header,
    });

    var response_object = JSON.parse(cv3response["body"]);
    var access_token = response_object.access_token;

    /////////////////////////////////// DATA REQUEST
    var request_object = new Object();

    var header = [];
    header["Content-Type"] = "application/json";
    header["Authorization"] = "Bearer " + access_token;

    request_object["data"] = new Object();
    request_object["data"]["exportProductIDs"] = new Object();

    // SEND THE REQUEST
    var cv3response = https.post({
      url: "https://service.commercev3.com/rest/products",
      body: JSON.stringify(request_object),
      headers: header,
    });

    var response_object = JSON.parse(cv3response["body"]);
    output["prod_ids"] = response_object.exportProductIDs.prod_ids;

    // OUTPUT FOR DEBUGGING
    html = html + JSON.stringify(output, null, 4);

    params.response.write({
      output: html,
    });
  }

  return {
    onRequest: onRequest,
  };
});
