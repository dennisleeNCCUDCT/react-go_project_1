import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

const Movies = () => {
  const [movies, setMovies] = useState([]);

  useEffect(() => {
    const headers = new Headers();
    //
    
    headers.append("Accept", "application/json"); //insert headers to innitial api request
   
   //
    const requestOptions = {
      method: "Get", //remember to add "," after every parameter
      headers: headers, //remember to add "," after every parameter
    };
    //fetch api請求，使用requestOptions參數
    fetch(`http://localhost:8080/movies`, requestOptions)
  .then((response) => {
    if (!response.ok) {
      throw new Error("Network response was not ok");
    }
    return response.json();
  })
  .then((data) => {
    console.log(data); // Log the data to see what is received
    setMovies(data);
  })
  .catch((err) => {
    console.error("Fetch error:", err);
  });
  }, []);

  return (
    <div>
      <h2>Movies</h2>
      <hr />
      <table className="table table-striped table-hover">
        <thead>
          <tr>
            <th>Movie</th>
            <th>Release Date</th>
            <th>Rating</th>
          </tr>
        </thead>
        <tbody>
          {movies.map((m) => (
            <tr key={m.id}>
              <td>
                <Link to={`/movies/${m.id}`}>{m.title}</Link>
              </td>
              <td>{m.release_date}</td>
              <td>{m.mpaa_rating}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Movies;
