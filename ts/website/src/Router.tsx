import { MemoryRouter, Switch, Route } from "react-router-dom";
import { createMemoryHistory } from "history";

import Header from "./components/nav/Header";

import HomePage from "./pages/HomePage";
import AboutPage from "./pages/AboutPage";

const Router = () => {
  const history = createMemoryHistory();

  return (
    // This needs to be changed!
    <MemoryRouter>
      <Header />
      <Switch>
        <Route path="/about">
          <AboutPage />
        </Route>
        <Route path="/">
          <HomePage />
        </Route>
      </Switch>
    </MemoryRouter>
  );
};

export default Router;
