import React, { useEffect, useContext, useState } from "react";
import SigninStyles from "./Style";
import appContext from "../../Components/Context/Context";
import { Link, useHistory } from "react-router-dom";
import { Grid } from "@material-ui/core";
import {
  HOMEPAGE_ROUTE,
  SIGN_IN_ROUTE,
  PROFILE_ROUTE
} from "../../Constants/RouteConstant";
import logo from "../../Images/logo.jpeg";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import Radio from "@material-ui/core/Radio";
import DoneIcon from "@material-ui/icons/Done";
import RadioGroup from "@material-ui/core/RadioGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import FormControl from "@material-ui/core/FormControl";
import createUser from "../../API/CreateUser";
import getUser from "../../API/GetUser";
import assignUnassignedStolenCar from "../../API/AssignStolenCar";

const SignIn = () => {
  const { userState, handleLoggedInUserState } = useContext(appContext);

  const classes = SigninStyles();
  const [selectedTab, setSelectedTab] = useState("SignIn");
  const [selectedUserType, setSelectedUserType] = useState("car_owner");

  const history = useHistory();

  useEffect(() => {
    if (!userState.loggedInUser) {
      history.push(SIGN_IN_ROUTE);
    }
  }, []);

  function radioHandleChange(event) {
    setSelectedUserType(event.target.value);
  }

  async function submitButtonClicked() {
    if (selectedTab == "SignIn") {
      // get user and verify
      var user = await getUser(document.getElementById("email").value);
      if (user.data) {
        handleLoggedInUserState({
          userProfile: user.data,
          loggedInUser: true
        });
        history.push(PROFILE_ROUTE);
      } else {
        alert(user.message);
      }
    } else if (selectedTab == "SignUp") {
      var userDetails = {
        name: document.getElementById("name").value,
        email: document.getElementById("email").value,
        usertype: selectedUserType == "car_owner" ? 1 : 9
      };

      var resp = await createUser(
        selectedUserType == "car_owner" ? "owner" : "police",
        userDetails
      );
      if (resp.created) {
        if (resp.created === true) {
          var user = await getUser(document.getElementById("email").value);
          if (user.data) {
            handleLoggedInUserState({
              userProfile: user.data,
              loggedInUser: true
            });

            if (user.data.usertype === 9) {
              var assignStolenCar = await assignUnassignedStolenCar(
                user.data.ID
              );
              if (
                assignStolenCar.data &&
                (assignStolenCar.data.created === true ||
                  (assignStolenCar.data.regnum &&
                    assignStolenCar.data.regnum.length > 0))
              ) {
                console.log("you're assigned to unassigned stolen car");
              } else {
                alert(assignStolenCar.message);
              }
            }
            history.push(PROFILE_ROUTE);
          } else {
            alert(user.message);
          }
        } else {
          alert(resp.message);
        }
      } else {
        alert(resp.message);
      }
    }
  }

  return (
    <div>
      <div>
        <Grid className={classes.row} container alignItems="center">
          <Grid style={{ textAlign: "left" }} xs={6} item>
            <Link to={HOMEPAGE_ROUTE}>
              <span>
                <img className={classes.logoImage} src={logo} alt="Logo" />
              </span>
            </Link>
          </Grid>
          <Grid
            item
            xs={3}
            onClick={() => {
              if (selectedTab == "SignUp") {
                setSelectedTab("SignIn");
              }
            }}
          >
            SignIn
          </Grid>
          <Grid
            item
            xs={3}
            onClick={() => {
              if (selectedTab == "SignIn") {
                setSelectedTab("SignUp");
              }
            }}
          >
            SignUp
          </Grid>
        </Grid>
      </div>
      <div className={classes.mainArea} id="carouselContainer">
        <div style={{ padding: "10px" }}>
          <div>{selectedTab}</div>
          <Grid container alignItems="center">
            <Grid style={{ textAlign: "center" }} xs={12} item>
              <TextField
                id="email"
                label="Email"
                variant="outlined"
                margin="normal"
              />
            </Grid>
            {selectedTab === "SignUp" ? (
              <Grid style={{ textAlign: "center" }} item xs={12}>
                <TextField
                  id="name"
                  label="Name"
                  variant="outlined"
                  margin="normal"
                />
              </Grid>
            ) : (
              ""
            )}
            {selectedTab === "SignUp" ? (
              <Grid style={{ textAlign: "center" }} xs={12} item>
                <FormControl component="fieldset">
                  <RadioGroup
                    aria-label="position"
                    name="position"
                    value={selectedUserType}
                    onChange={radioHandleChange}
                    row
                  >
                    <FormControlLabel
                      value="car_owner"
                      control={<Radio />}
                      label={"Car Owner"}
                    />
                    <FormControlLabel
                      value="police_officer"
                      control={<Radio />}
                      label={"Police Officer"}
                    />
                  </RadioGroup>
                </FormControl>
              </Grid>
            ) : (
              ""
            )}
            <Grid style={{ textAlign: "center" }} xs={12} item>
              <Button
                className={classes.confirmBtn}
                size="small"
                variant="contained"
                onClick={e => {
                  submitButtonClicked();
                }}
              >
                <DoneIcon /> &nbsp;&nbsp;{selectedTab}
              </Button>
            </Grid>
          </Grid>
        </div>
      </div>
    </div>
  );
};

export default SignIn;
