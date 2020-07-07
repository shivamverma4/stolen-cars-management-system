import React, { useState, useEffect } from "react";
import "./App.css";
import Routes from "./Routes/Routes";
import appContext from "./Components/Context/Context";
import { CircularProgress } from "@material-ui/core";
// import fetchUserProfileData from "./API/UserProfile/FetchUserProfile";
// import Cookie from "./Utils/Cookie";
// import Config from "./Config/Config";
import history from "./Routes/history";
import { SIGN_IN_ROUTE, HOMEPAGE_ROUTE } from "./Constants/RouteConstant";
import HomePage from "./Views/HomePage/HomePage";
import SignIn from "./Views/SignIn/SignIn";

function App() {
  const [userState, setUserState] = useState({
    userProfile: null,
    loggedInUser: false
  });

  function handleLoggedInUserState(state) {
    setUserState({ ...state });
  }

  useEffect(() => {}, []);

  return (
    <div className="App" style={{ height: "100vh" }}>
      <appContext.Provider value={{ userState, handleLoggedInUserState }}>
        <Routes />
      </appContext.Provider>
    </div>
  );
  // <div className="circularProgressContainer">
  //   <CircularProgress />
  // </div>
}

export default App;
