import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import Snippets from "./Snippets";
import MenuItem from "@material-ui/core/MenuItem";
import axios from "axios";
import Editor from "react-simple-code-editor";
import { highlight, languages } from "prismjs/components/prism-core";
import "prismjs/components/prism-clike";
import "prismjs/components/prism-javascript";
import "prismjs/components/prism-java";
import "prismjs/components/prism-python";
import "prismjs/components/prism-go";
import "prismjs/themes/prism.css";

const useStyles = makeStyles((theme) => ({
  textBox: {
    height: "100%",
    display: "flex",
    flexWrap: "wrap",
  },
  textBox__codeBox: {
    textAlign: "left",
    color: theme.palette.text.secondary,
    height: "100%",
    width: "100%",
  },
  Selector: {
    marginTop: 70,
    marginBottom: 15,
    width: "100%",
    display: "flex",
    flexWrap: "wrap",
  },
  Selector__grid: {
    width: "100%",
  },
  Selector__lang: {
    width: "100%",
  },
  Selector__formatter: {
    width: "80%",
  },
  Selector__formatterButton: {
    width: "20%",
  },
}));

const supported_languages = [
  {
    value: "python",
    label: "Python",
  },
  { value: "java", label: "Java" },
  {
    value: "javascript",
    label: "Javascript",
  },
  { value: "golang", label: "Go" },
];

const formatters = {
  python: ["black", "autopep8", "yapf"],
  java: ["google-java-format"],
  javascript: ["prettier", "js-beautify"],
  golang: ["gofmt"],
};

export default function Body() {
  const classes = useStyles();
  const [language, setLanguage] = React.useState("python");
  const [formatter, setFormatter] = React.useState("black");
  const [originalCode, setOriginalCode] = React.useState(Snippets[language]);
  const [formattedCode, setFormattedCode] = React.useState(Snippets[language]);

  const handleLangChange = (event) => {
    setLanguage(event.target.value);
    setOriginalCode(Snippets[event.target.value]);
  };
  const handleFormatterChange = (event) => {
    setFormatter(event.target.value);
  };
  const handleInputChange = (event) => {
    setOriginalCode(event.target.value);
  };

  return (
    <React.Fragment>
      <Grid container spacing={3} className={classes.Selector}>
        <Grid item xs={1} />
        <Grid item xs={5} className={classes.Selector__grid}>
          <TextField
            id="standard-select-language-native"
            select
            label="Language"
            value={language}
            onChange={handleLangChange}
            helperText="Please select your language"
            className={classes.Selector__lang}
          >
            {supported_languages.map((option) => (
              <MenuItem
                key={option.value}
                value={option.value}
                formatter={option.formatter}
              >
                {option.label}
              </MenuItem>
            ))}
          </TextField>
        </Grid>
        <Grid item xs={5} className={classes.Selector__grid}>
          <TextField
            id="standard-select-formatter-native"
            select
            label="Formatter"
            value={formatter}
            onChange={handleFormatterChange}
            helperText="Please select your formatter"
            className={classes.Selector__formatter}
          >
            {formatters[language].map((option) => (
              <MenuItem key={option} value={option}>
                {option}
              </MenuItem>
            ))}
          </TextField>
          <Button
            variant="contained"
            color="secondary"
            className={classes.Selector__formatterButton}
            onClick={() => {
              const host = `api/${language}/${formatter}/${encodeURI(
                originalCode
              )}`;
              console.log("Host : ", host);

              axios.post(host).then(function (response) {
                console.log("Response : ", response.data.result);
                setFormattedCode(response.data.result);
              });
            }}
          >
            Format
          </Button>
          <Grid item xs={1} />
        </Grid>
      </Grid>

      <Grid container spacing={3} className={classes.textBox}>
        <Grid item xs={1} />
        <Grid item xs={5}>
          <Editor
            className={classes.textBox__codeBox}
            value={originalCode}
            onChange={handleInputChange}
            onValueChange={setOriginalCode}
            highlight={(originalCode) =>
              highlight(
                originalCode,
                languages[language !== "golang" ? language : "go"]
              )
            }
            padding={10}
            style={{
              fontFamily: '"Fira code", "Fira Mono", monospace',
              fontSize: 16,
            }}
            tabSize={4}
          />
        </Grid>
        <Grid item xs={5}>
          <Editor
            className={classes.textBox__codeBox}
            value={formattedCode}
            onChange={handleInputChange}
            onValueChange={setOriginalCode}
            highlight={(formattedCode) =>
              highlight(
                formattedCode,
                languages[language !== "golang" ? language : "go"]
              )
            }
            padding={10}
            style={{
              fontFamily: '"Fira code", "Fira Mono", monospace',
              fontSize: 16,
            }}
            tabSize={4}
          />
        </Grid>
        <Grid item xs={1} />
      </Grid>
    </React.Fragment>
  );
}
