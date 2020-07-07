import React from "react";
import history from "./history";
import HomePage from "../Views/HomePage/HomePage";
import SignIn from "../Views/SignIn/SignIn";
import Profile from "../Views/Profile/Profile";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import {
  HOMEPAGE_ROUTE,
  SIGN_IN_ROUTE,
  PROFILE_ROUTE
} from "../Constants/RouteConstant";

export default function Routes(props) {
  return (
    <Router history={history}>
      <Switch>
        <Route
          path={HOMEPAGE_ROUTE}
          exact
          render={homePageProps => <HomePage {...props} {...homePageProps} />}
        />
        <Route
          path={SIGN_IN_ROUTE}
          exact
          render={SignInProps => <SignIn {...props} {...SignInProps} />}
        />
        <Route
          path={PROFILE_ROUTE}
          exact
          render={ProfileProps => <Profile {...props} {...ProfileProps} />}
        />
      </Switch>
    </Router>
  );
}
