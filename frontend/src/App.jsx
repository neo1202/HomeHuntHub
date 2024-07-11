// 實際內容物display在Outlet
// router中的Link使得不用每次刷新整個頁面
import { useState } from "react";
import { Link, Outlet, useNavigate } from "react-router-dom";
import Alert from "./pages/Alert";

function App() {
  const [jwtToken, setJwtToken] = useState("");
  const [alertMessage, setAlertMessage] = useState("");
  const [alertClassName, setAlertClassName] = useState("hidden");

  const navigate = useNavigate();

  const logOut = () => {
    setJwtToken("");
    navigate("/login");
  };

  return (
    <div className="flex flex-col w-screen h-screen m-0 bg-slate-200">
      <div className="flex flex-row w-full h-18">
        <div className="w-5/6 mt-3 text-4xl text-neutral-950">
          Go Watch a Movie!
        </div>

        <div className="flex flex-col w-1/6 text-end">
          {jwtToken === "" ? (
            <Link to="/login">
              <span className="text-white bg-green-500">Login</span>
            </Link>
          ) : (
            <a href="#!" onClick={logOut}>
              <span className="text-white bg-red-500">Logout</span>
            </a>
          )}
        </div>
        <hr className="mb-12"></hr>
      </div>
      <hr className="h-0.5 bg-cyan-950"></hr>
      <div className="flex flex-row flex-1 w-full">
        <div className="flex flex-row w-1/6">
          <nav className="flex flex-col w-full text-sm font-medium text-gray-900 bg-white border border-gray-200 rounded-lg">
            <Link to="/" className="w-full px-4 py-2 border-b border-gray-200 ">
              Home
            </Link>
            <Link
              to="/movies"
              className="w-full px-4 py-2 border-b border-gray-200"
            >
              Movies
            </Link>
            <Link
              to="/genres"
              className="w-full px-4 py-2 border-b border-gray-200"
            >
              Genres
            </Link>
            {jwtToken !== "" && (
              <>
                <Link
                  to="/admin/movie/0"
                  className="w-full px-4 py-2 border-b border-gray-200"
                >
                  Add Movie
                </Link>
                <Link
                  to="/manage-catalogue"
                  className="w-full px-4 py-2 border-b border-gray-200"
                >
                  Manage Catalogue
                </Link>
                <Link
                  to="/graphql"
                  className="w-full px-4 py-2 border-b border-gray-200"
                >
                  GraphQL
                </Link>
              </>
            )}
          </nav>
        </div>
        <div className="flex flex-col w-5/6">
          <Alert message={alertMessage} className={alertClassName} />
          <Outlet
            context={{
              jwtToken,
              setJwtToken,
              setAlertClassName,
              setAlertMessage,
            }}
          />
        </div>
      </div>
    </div>
  );
}

export default App;
