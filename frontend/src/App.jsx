import "./App.css";
import Body from "./Body";
import Header from "./Header";
import React from "react";
import { Grid } from "@material-ui/core";

function App() {
  return (
    <section className="container">
      <Grid
        container
        direction="column"
        className="app"
        style={{ padding: "0 10px 10px" }}
      >
        <Header />
        <Body />
      </Grid>
    </section>
  );
}

export default App;
