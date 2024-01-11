import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

const Movies = () => {
  const [movies, setMovies] = useState([]);

  useEffect(() => {
    let moviesList = [
      {
        id: 1,
        title: "Highlander",
        release_date: "1986-03-07",
        runtime: 116,
        mpaa_rating: "R",
        description: "Some long description",
      },
      {
        id: 2,
        title: "Raiders of the Lost Ark",
        release_date: "1981-06-12",
        runtime: 115,
        mpaa_rating: "PG-13",
        description: "Some long description",
      },
    ];

    setMovies(moviesList);
  }, []);

  return (
    <div>
      <h2 className="my-4 text-2xl font-bold">Movies</h2>
      <hr className="mb-4" />
      <table className="min-w-full divide-y divide-gray-200">
        <thead className="bg-gray-50">
          <tr>
            <th className="px-6 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase">
              Movie
            </th>
            <th className="px-6 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase">
              Release Date
            </th>
            <th className="px-6 py-3 text-xs font-medium tracking-wider text-left text-gray-500 uppercase">
              Rating
            </th>
          </tr>
        </thead>
        <tbody className="text-black bg-white divide-y divide-gray-200">
          {movies.map((m) => (
            <tr key={m.id}>
              <td className="px-6 py-4 whitespace-nowrap">
                <Link
                  to={`/movies/${m.id}`}
                  className="text-indigo-600 hover:text-indigo-900"
                >
                  {m.title}
                </Link>
              </td>
              <td className="px-6 py-4 whitespace-nowrap">{m.release_date}</td>
              <td className="px-6 py-4 whitespace-nowrap">{m.mpaa_rating}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Movies;
