import React from "react";
import Router from "./Router";
import { Auth0Provider } from "@auth0/auth0-react";
import config from "./config";

console.log(process.env);

const App = () => {
  return (
    <Auth0Provider
      domain={config.AUTH0_DOMAIN}
      clientId={config.AUTH0_CLIENT_ID}
      redirectUri={"test"}
      audience={`https://${config.AUTH0_DOMAIN}/api/v2/`}
      scope="read:current_user update:current_user_metadata"
    >
      <Router />
    </Auth0Provider>
  );
};

export default App;
