const express = require("express");
const app = express();
const axios = require("axios");

app.get("/", (req, res) => {
  console.log(req.query);
  req.query.url.map((u) => {
    console.log(u);
  });
  console.log("user hit the resource");
  axios.get("http://localhost:8090/primes").then(function (response) {
    console.log(response.data);
  });
  res.status(200).send("Home Page");
});

app.listen(8080, () => {
  console.log("server is listening on port 8080...");
});