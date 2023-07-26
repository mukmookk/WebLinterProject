import { AppBar, Toolbar } from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";
import React from "react";
import { Grid } from "@material-ui/core";
import Logo from "./favicon.ico";

const useStyles = makeStyles((theme) => ({
  root: {},
  header: {
    backgroundColor: "black",
    height: 60,
    boxShadow: "None",
  },
  header__toolbar: {
    padding: 0,
    margin: 0.1,
  },
  logo_container: {
    flexGrow: 1,
  },
  logo: {
    maxWidth: 30,
    display: "block",
    marginLeft: 10,
    marginRight: 10,
  },
  header__menuButton: {
    color: "#FE6B8B",
  },
  header__category: {
    marginLeft: "auto",
  },
  list: {
    width: 250,
  },
  icon: {
    color: "#FE6B8B",
  },
  link: {
    color: "black",
    textDecoration: "none",
  },
}));

export default function Header() {
  const classes = useStyles();
  return (
    <React.Fragment>
      <Grid container className={classes.root}>
        <AppBar className={classes.header}>
          <Toolbar className={classes.header__toolbar}>
            <img className={classes.logo} src={Logo} alt="WebLinterProject logo" />
            WebLinterProject
          </Toolbar>
        </AppBar>
      </Grid>
    </React.Fragment>
  );
}
