import React from "react";
import { Link } from "react-router-dom";
import Ticket from "../assets/movie_tickets.jpg";

function Home() {
  return (
    <>
      <div className="text-center text-black">
        <div>Find a movie to watch tonight!</div>
        <hr />
        <Link to="/movies">
          <img src={Ticket} alt="movie tickets"></img>
        </Link>
      </div>
    </>
  );
}

export default Home;
