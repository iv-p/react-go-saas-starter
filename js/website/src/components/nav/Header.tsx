import { useEffect } from "react";
import { useAuth0 } from "@auth0/auth0-react";
import { Link } from "react-router-dom";

import config from "config";

import LoginButton from "components/auth/LoginButton";
import LogoutButton from "components/auth/LogoutButton";
import Profile from "components/user/Profile";

const Header = () => {
  const { isAuthenticated, isLoading, getAccessTokenSilently } = useAuth0();
  useEffect(() => {
    const getUserMetadata = async () => {
  
      try {
        const accessToken = await getAccessTokenSilently({
          audience: `https://${config.AUTH0_DOMAIN}/api/v2/`,
          scope: "read:current_user",
        });

        console.log(accessToken);
      } catch (e) {
        console.log(e.message);
      }
    };
  
    getUserMetadata();
  }, [getAccessTokenSilently]);

  const userProfile = isLoading ? (
    <span>Loading profile</span>
  ) : (
    <>
      <Profile />
      <LogoutButton />
    </>
  );

  return (
    <nav>
      <ul>
        <li>
          <Link to="/">Home</Link>
        </li>
        <li>
          <Link to="/about">About</Link>
        </li>
      </ul>
      {isAuthenticated || isLoading ? userProfile : <LoginButton />}
    </nav>
  );
};

export default Header;
