import React from "react";
import { Link } from "react-router-dom";
import './Header.css';

const Header = () => {
  return (
    <header className="header-container">
      <Link to="/">
        <button className="header-button">Home</button>
      </Link>
      <Link to="/posts">
        <button className="header-button">Posts</button>
      </Link>
      <Link to="/likedposts">
        <button className="header-button">Liked Posts</button>
      </Link>
      <Link to="/createpost">
        <button className="header-button">Create Post</button>
      </Link>
    </header>
  );
};

export default Header;