import React, { useEffect, useContext, useState } from "react";
import profilePageStyles from "./Style";
import Constants from "../../Constants/Constants";
import appContext from "../../Components/Context/Context";
import { Link, useHistory } from "react-router-dom";
import { Card, CircularProgress, Grid } from "@material-ui/core";
import {
  HOMEPAGE_ROUTE,
  SIGN_IN_ROUTE,
  PROFILE_ROUTE
} from "../../Constants/RouteConstant";
import logo from "../../Images/logo.jpeg";
import PersonIcon from "@material-ui/icons/Person";
import Fab from "@material-ui/core/Fab";
import AddIcon from "@material-ui/icons/Add";
import ClearIcon from "@material-ui/icons/Clear";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import DoneIcon from "@material-ui/icons/Done";
import createStolenCarRecord from "./../../API/CreateSCRecord";
import getAllStolenCars from "./../../API/GetAllCars";
import RadioGroup from "@material-ui/core/RadioGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import FormControl from "@material-ui/core/FormControl";
import Radio from "@material-ui/core/Radio";
import changeStolenCarStatus from "./../../API/UpdateCarStatus";

const Profile = () => {
  const classes = profilePageStyles();
  const { userState, handleLoggedInUserState } = useContext(appContext);
  const history = useHistory();

  const [addStolenCar, setAddStolenCar] = useState(false);
  const [refreshStolenCars, setRefreshStolenCars] = useState(false);

  const [carDetails, setCarDetails] = useState({
    regnum: "",
    color: "",
    ownerID: "",
    poID: "",
    description: ""
  });

  const [allCarDetails, setAllCarDetails] = useState(0);
  const [carStatus, setCarStatus] = useState(0);

  useEffect(() => {
    if (!userState.loggedInUser) {
      history.push(SIGN_IN_ROUTE);
    }

    if (userState.loggedInUser) {
      let temp = carDetails;
      let userType = "";
      let uID = "";
      if (userState.userProfile.usertype === 9) {
        temp["poID"] = userState.userProfile.ID;
        uID = userState.userProfile.ID;
        userType = "police";
      } else if (userState.userProfile.usertype === 1) {
        temp["ownerID"] = userState.userProfile.ID;
        uID = userState.userProfile.ID;
        userType = "owner";
      }
      async function getCars() {
        var allUserStolenCars = await getAllStolenCars(uID, userType);
        setAllCarDetails(allUserStolenCars);
      }
      getCars();

      setCarDetails(temp);
    }
    setRefreshStolenCars(false);
    setAddStolenCar(false);
  }, [refreshStolenCars]);

  function handleOnInputChange(event, key) {
    var temp = carDetails;
    temp[key] = event.target.value;
    setCarDetails(temp);
  }

  async function submitButtonClicked() {
    var resp = await createStolenCarRecord(carDetails);
    if (resp.created === true) {
      setRefreshStolenCars(true);
    } else {
      alert(resp.message);
    }
  }

  async function changeStolenCarStatusFunc(carID) {
    if (carStatus != "0") {
      var resp = await changeStolenCarStatus({
        oID: carID,
        poID: userState.userProfile.ID,
        status: parseInt(carStatus)
      });
      if (resp.created === true) {
        setRefreshStolenCars(true);
      }
    }
  }

  function radioHandleChange(event) {
    setCarStatus(event.target.value);
  }

  function showCarDetailedCards(car, index) {
    return (
      <div key={index}>
        <Card style={{ margin: "2rem 0rem" }}>
          <Grid
            style={{
              padding: "1.2rem 1rem",
              boxShadow:
                "0 0 4px 0 rgba(0, 0, 0, 0.12), 0 4px 4px 0 rgba(0, 0, 0, 0.16)"
            }}
            className={
              car.status === 0
                ? classes.greenYellow
                : car.status === 1
                ? classes.yellow
                : car.status === 5
                ? classes.red
                : car.status === 9
                ? classes.darkGreen
                : classes.greenYellow
            }
            container
            alignItems="center"
          >
            <Grid
              style={{ display: "flex", margin: "0.5rem 0rem" }}
              xs={12}
              item
            >
              <Grid style={{ textAlign: "left" }} xs={6} item>
                {"Registration Number :"}
              </Grid>
              <Grid style={{ textAlign: "left" }} xs={6} item>
                {car.regnum}
              </Grid>
            </Grid>
            <Grid
              style={{ display: "flex", margin: "0.5rem 0rem" }}
              xs={12}
              item
            >
              <Grid style={{ textAlign: "left" }} xs={6} item>
                {"Color :"}
              </Grid>
              <Grid style={{ textAlign: "left" }} xs={6} item>
                {car.color}
              </Grid>
            </Grid>
            <Grid
              style={{ display: "flex", margin: "0.5rem 0rem" }}
              xs={12}
              item
            >
              <Grid style={{ textAlign: "left" }} xs={6} item>
                {"Description :"}
              </Grid>
              <Grid style={{ textAlign: "left" }} xs={6} item>
                {car.description}
              </Grid>
            </Grid>
            <Grid
              style={{ display: "flex", margin: "0.5rem 0rem" }}
              xs={12}
              item
            >
              <Grid style={{ textAlign: "left" }} xs={6} item>
                {"Status : "}
              </Grid>
              <Grid style={{ textAlign: "left" }} xs={6} item>
                {car.status === 0
                  ? "No Police Officer available"
                  : car.status === 1
                  ? "Assigned to Police Officer"
                  : car.status === 5
                  ? "Not Found"
                  : car.status === 9
                  ? "Found"
                  : "-"}
              </Grid>
            </Grid>
            {userState.loggedInUser &&
            userState.userProfile.usertype === 9 &&
            car.status === 1 ? (
              <Grid
                style={{ display: "flex", margin: "0.5rem 0rem" }}
                xs={12}
                item
              >
                <Grid style={{ textAlign: "left" }} xs={6} item>
                  {"Change Status : "}
                </Grid>
                <Grid style={{ textAlign: "left" }} xs={6} item>
                  <Grid style={{ textAlign: "center" }} xs={12} item>
                    <FormControl component="fieldset">
                      <RadioGroup
                        aria-label="position"
                        name="position"
                        value={carStatus}
                        onChange={radioHandleChange}
                        row
                      >
                        <FormControlLabel
                          value={"9"}
                          control={<Radio />}
                          label={"Found"}
                        />
                        <FormControlLabel
                          value={"5"}
                          control={<Radio />}
                          label={"Not Found"}
                        />
                      </RadioGroup>
                    </FormControl>
                  </Grid>
                  <Grid style={{ textAlign: "center" }} xs={12} item>
                    <Button
                      className={classes.confirmBtn}
                      size="small"
                      variant="contained"
                      id={car.ID}
                      onClick={e => {
                        changeStolenCarStatusFunc(car._id);
                      }}
                    >
                      <DoneIcon /> &nbsp;&nbsp;{"Submit"}
                    </Button>
                  </Grid>
                </Grid>
              </Grid>
            ) : (
              ""
            )}
          </Grid>
        </Card>
      </div>
    );
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
          {!userState.loggedInUser ? (
            <Grid
              className={classes.userContainer}
              item
              xs={6}
              onClick={() => {
                history.push(SIGN_IN_ROUTE);
              }}
            >
              {" "}
              | SignIn |
            </Grid>
          ) : (
            <>
              <Grid
                className={classes.userContainer}
                item
                xs={3}
                onClick={() => {
                  history.push(PROFILE_ROUTE);
                }}
              >
                <PersonIcon />
                &nbsp;&nbsp;
                {userState.userProfile.name}
              </Grid>
              <Grid
                className={classes.userContainer}
                item
                xs={3}
                onClick={() => {
                  window.location = window.location.href;
                }}
              >
                {"Logout"}
              </Grid>
            </>
          )}
        </Grid>
      </div>
      <div className={classes.mainArea} id="carouselContainer">
        <Grid
          container
          style={{ padding: "10px", justifyContent: "space-evenly" }}
          alignItems="center"
        >
          <Grid className={classes.mainAreaBlocks} xs={5} item>
            {userState.loggedInUser && userState.userProfile.usertype === 1 ? (
              <>
                {allCarDetails && allCarDetails.data.length > 0 ? (
                  <>
                    {"Car Details"}
                    <>{allCarDetails.data.map(showCarDetailedCards)}</>
                  </>
                ) : (
                  "No Detials of Stolen Cars are added by you"
                )}
              </>
            ) : (
              <>
                {allCarDetails && allCarDetails.data.length > 0 ? (
                  <>
                    {"Assigned Car Details"}
                    <>{allCarDetails.data.map(showCarDetailedCards)}</>
                  </>
                ) : (
                  "No Stolen Cars are assigned to you yet"
                )}
              </>
            )}
          </Grid>
          {userState.loggedInUser && userState.userProfile.usertype === 1 ? (
            <Grid className={classes.mainAreaBlocks} xs={5} item>
              <div>Add Details of Stolen Car</div>
              {addStolenCar ? (
                <>
                  <Grid container alignItems="center">
                    <Grid style={{ textAlign: "center" }} xs={12} item>
                      <TextField
                        id="regnum"
                        label="Registration Number"
                        variant="outlined"
                        margin="normal"
                        onChange={e => handleOnInputChange(e, "regnum")}
                      />
                    </Grid>
                    <Grid style={{ textAlign: "center" }} item xs={12}>
                      <TextField
                        id="color"
                        label="Color"
                        variant="outlined"
                        margin="normal"
                        onChange={e => handleOnInputChange(e, "color")}
                      />
                    </Grid>
                    <Grid style={{ textAlign: "center" }} item xs={12}>
                      <TextField
                        id="description"
                        label="Description"
                        variant="outlined"
                        margin="normal"
                        onChange={e => handleOnInputChange(e, "description")}
                      />
                    </Grid>
                    <Grid style={{ textAlign: "center" }} xs={12} item>
                      <Button
                        className={classes.confirmBtn}
                        size="small"
                        variant="contained"
                        onClick={e => {
                          submitButtonClicked();
                        }}
                      >
                        <DoneIcon /> &nbsp;&nbsp;Submit
                      </Button>
                    </Grid>
                    <Grid
                      style={{ textAlign: "center", margin: "10px" }}
                      xs={12}
                      item
                    >
                      <Fab
                        color="secondary"
                        aria-label="add"
                        onClick={e => {
                          setAddStolenCar(false);
                        }}
                      >
                        <ClearIcon />
                      </Fab>
                    </Grid>
                  </Grid>
                </>
              ) : (
                ""
              )}
              {!addStolenCar ? (
                <div style={{ textAlign: "right", paddingTop: "150px" }}>
                  <Fab
                    color="primary"
                    aria-label="add"
                    onClick={e => {
                      setAddStolenCar(true);
                    }}
                  >
                    <AddIcon />
                  </Fab>
                </div>
              ) : (
                ""
              )}
            </Grid>
          ) : (
            ""
          )}
        </Grid>
      </div>
    </div>
  );
};

export default Profile;
