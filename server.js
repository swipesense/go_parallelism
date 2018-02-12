const express = require("express");
const app = express();

app.get("/", (req, res) => {
  setTimeout(() => res.send("Hello World!"), 1000);
});

app.listen(3000, () => console.log("Example server started on port 3000."));
