import { makeStyles } from "@material-ui/core/styles";

const profilePageStyles = makeStyles(theme => ({
  row: {
    padding: "20px",
    backgroundColor: "#3F01BC",
    color: "white"
  },
  logoImage: {
    width: "130px",
    maxWidth: "100%"
  },
  mainArea: {
    backgroundColor: "white"
  },
  avatar: {
    width: "40px",
    maxWidth: "100%"
  },
  userContainer: {
    display: "inline-flex",
    justifyContent: "flex-end"
  },
  mainAreaBlocks: {
    textAlign: "center",
    height: "100vh",
    textAlign: "center",
    border: "solid",
    borderRadius: "11px",
    padding: "20px",
    overflowY: "scroll"
  },
  greenYellow: { backgroundColor: "greenyellow" },
  yellow: { backgroundColor: "yellow" },
  red: { backgroundColor: "red" },
  darkGreen: { backgroundColor: "darkgreen" },
  greenYellow: { backgroundColor: "greenyellow" }
}));

export default profilePageStyles;
