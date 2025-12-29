import { useEffect, useState } from "react";
import "./App.css";

function App() {
  const [error, setError] = useState<undefined | string>(undefined);
  const [control, setControl] = useState(false);

  useEffect(() => {
    async function Verify() {
      try {
        const response = await fetch("/api/verify", {
          method: "POST",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        });
        if (!response.ok) {
          const data = await response.json();
          setError(data.error);
          return;
        }
        const data = await response.json();
        console.log(data);
        setControl(true);
        setError(undefined);
      } catch (err) {
        console.log(err);
        setError("error occured");
      }
    }
    Verify();
  }, []);

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    const form = e.currentTarget as HTMLFormElement;
    const formData = new FormData(form);
    const reqData = Object.fromEntries(formData.entries());

    const config = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(reqData),
    };

    try {
      console.log(reqData);
      const response = await fetch("/api/auth/register", config);
      if (!response.ok) {
        const data = await response.json();
        setError(data.error);
        return;
      }
      const data = await response.json();
      console.log("userSignIn", data);
      setError(undefined);
    } catch (err) {
      console.log(`Error: ${err}`);
      setError(`An error occurred  ${err}`);
    }
  }

  async function handleLogin(e: React.FormEvent<HTMLFormElement>) {
    console.log("handleLogin Called");
    e.preventDefault();
    const form = e.currentTarget as HTMLFormElement;
    const formData = new FormData(form);
    const reqData = Object.fromEntries(formData.entries());

    const config = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(reqData),
    };

    try {
      console.log(reqData);
      const response = await fetch("/api/auth/login/", config);
      if (!response.ok) {
        const data = await response.json();
        setError(data.error);
        return;
      }
      const data = await response.json();
      console.log(data);
      localStorage.setItem("token", data.token);
      setError(undefined);
      setControl(true);
    } catch (err) {
      console.log(`Error: ${err}`);
      setError("An error occurred");
    }
  }

  return (
    <>
      {error && <div>{error}</div>}
      <h1>Register</h1>
      <form onSubmit={handleSubmit}>
        <input type="text" name="username" placeholder="Enter your username" />
        <input
          type="password"
          name="password"
          id="password"
          placeholder="Enter Your Password"
        />
        <button type="submit">Submit</button>
      </form>
      <h1>Login</h1>
      <form onSubmit={handleLogin}>
        <input type="text" name="username" placeholder="Enter your username" />
        <input
          type="password"
          name="password"
          id="password"
          placeholder="Enter Your Password"
        />
        <button type="submit">Submit</button>
      </form>
      {control && <div>secret</div>}
    </>
  );
}

export default App;
