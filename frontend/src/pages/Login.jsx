import { useState } from "react";
import { useNavigate, useOutletContext } from "react-router-dom"; //useOutletContext：用于向子route传递context；
import Input from "../components/form/Input";

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const { setJwtToken } = useOutletContext();
  const { setAlertClassName } = useOutletContext();
  const { setAlertMessage } = useOutletContext();

  const navigate = useNavigate();

  const handleSubmit = (event) => {
    event.preventDefault();
    console.log("email/pass", email, password);
    // build the request payload
    let payload = {
      email: email,
      password: password,
    };
    const requestOptions = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(payload),
    };
    fetch(`/api/authenticate`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        if (data.error) {
          setAlertClassName(
            "bg-red-500 text-white p-4 rounded border border-red-700"
          );
          setAlertMessage(data.message);
        } else {
          setJwtToken(data.access_token);
          setAlertClassName("hidden");
          setAlertMessage("");
          navigate("/");
        }
      })
      .catch((error) => {
        setAlertClassName(
          "bg-red-500 text-white p-4 rounded border border-red-700"
        );
        setAlertMessage(error);
      });
  };

  return (
    <div className="w-1/2 ml-[25%] bg-zinc-600">
      <h2>Login</h2>
      <hr />

      <form onSubmit={handleSubmit}>
        <Input
          title="Email Address"
          type="email"
          className="w-full px-3 py-2 text-sm border border-gray-300 rounded focus:border-blue-500 focus:ring focus:ring-blue-200"
          name="email"
          autoComplete="email-new"
          onChange={(event) => setEmail(event.target.value)}
        />

        <Input
          title="Password"
          type="password"
          className="w-full px-3 py-2 text-sm border border-gray-300 rounded focus:border-blue-500 focus:ring focus:ring-blue-200"
          name="password"
          autoComplete="password-new"
          onChange={(event) => setPassword(event.target.value)}
        />

        <hr />

        <input
          type="submit"
          className="px-4 py-2 text-sm text-white bg-blue-500 rounded hover:bg-blue-700"
          value="Login"
        />
      </form>
    </div>
  );
};

export default Login;
