import { BrowserRouter, Switch, Route } from "react-router-dom";

import Header from "./components/nav/Header";

import HomePage from "./pages/HomePage";
import AboutPage from "./pages/AboutPage";

const Router = () => {
  return (
    <BrowserRouter>
      <Header />
      <Switch>
        <Route path="/about">
          <AboutPage />
        </Route>
        <Route path="/">
          <HomePage />
        </Route>
      </Switch>
    </BrowserRouter>
  );
};

export default Router;
