import { useAuth0 } from "@auth0/auth0-react";
import { Link } from "react-router-dom";

import LoginButton from "components/auth/LoginButton";
import LogoutButton from "components/auth/LogoutButton";
import Profile from "components/user/Profile";

const Header = () => {
  const { isAuthenticated, isLoading } = useAuth0();
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
