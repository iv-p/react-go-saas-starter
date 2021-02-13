interface Config {
  AUTH0_DOMAIN: string;
  AUTH0_CLIENT_ID: string;
}

const config: Config = {
  AUTH0_DOMAIN: process.env.REACT_APP_AUTH0_DOMAIN || "",
  AUTH0_CLIENT_ID: process.env.REACT_APP_AUTH0_CLIENT_ID || "",
};

export default config;
