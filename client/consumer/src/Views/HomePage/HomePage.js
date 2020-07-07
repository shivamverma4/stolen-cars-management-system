import React, { useEffect, useRef, useContext, useState } from "react";
// import homePageStyles from "./Style";
import Constants from "../../Constants/Constants";
import appContext from "../../Components/Context/Context";
import history from "../../Routes/history";
import { Link } from "react-router-dom";
import { Grid } from "@material-ui/core";
import { HOMEPAGE_ROUTE, SIGN_IN_ROUTE } from "../../Constants/RouteConstant";
import logo from "../../Images/logo.jpeg";

const HomePage = () => {
  //   const classes = homePageStyles();

  useEffect(() => {}, []);

  return (
    <div>
      <div>
        <Grid container alignItems="center">
          <Grid item>
            {/* <Link to={HOMEPAGE_ROUTE}> */}
            <span>
              <img src={logo} alt="Logo" />
            </span>
            {/* </Link> */}
          </Grid>
          <Grid item>Login as Car Owner</Grid>
          <Grid item>Login as Police Officer</Grid>
        </Grid>
      </div>
      <div id="carouselContainer">Here we are for stolen cars project</div>
    </div>
  );
};

export default HomePage;
