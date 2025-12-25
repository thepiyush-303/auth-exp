import { useState } from "react";
import "./App.css";

type Message = { type: "success" | "error"; text: string } | null;

function App() {
  const [activeTab, setActiveTab] = useState<"register" | "login">("register");
  const [message, setMessage] = useState<Message>(null);

  const submitRegister: React.FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();
    setMessage(null);

    const form = e.currentTarget as HTMLFormElement;
    const fd = new FormData(form);
    const params = new URLSearchParams();
    params.set("name", String(fd.get("name") ?? ""));
    params.set("email", String(fd.get("email") ?? ""));
    params.set("password", String(fd.get("password") ?? ""));

    try {
      const res = await fetch("/register", {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: params.toString(),
      });
      const text = await res.text();
      if (!res.ok) {
        setMessage({ type: "error", text });
      } else {
        setMessage({ type: "success", text });
        form.reset();
      }
    } catch (err: any) {
      setMessage({ type: "error", text: err?.message || "Network error" });
    }
  };

  const submitLogin: React.FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();
    setMessage(null);

    const form = e.currentTarget as HTMLFormElement;
    const fd = new FormData(form);
    const params = new URLSearchParams();
    params.set("email", String(fd.get("email") ?? ""));
    params.set("password", String(fd.get("password") ?? ""));

    try {
      const res = await fetch("/login", {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: params.toString(),
      });
      const text = await res.text();
      if (!res.ok) {
        setMessage({ type: "error", text });
      } else {
        setMessage({ type: "success", text });
        form.reset();
      }
    } catch (err: any) {
      setMessage({ type: "error", text: err?.message || "Network error" });
    }
  };

  return (
    <div className="container">
      <h1>Auth</h1>

      {message && (
        <div className={`banner ${message.type}`}>{message.text}</div>
      )}

      <div className="tabs">
        <button
          className={activeTab === "register" ? "active" : ""}
          onClick={() => setActiveTab("register")}
        >
          Register
        </button>
        <button
          className={activeTab === "login" ? "active" : ""}
          onClick={() => setActiveTab("login")}
        >
          Login
        </button>
      </div>

      {activeTab === "register" ? (
        <form className="form" onSubmit={submitRegister}>
          <label>
            <span>Name</span>
            <input name="name" type="text" required />
          </label>
          <label>
            <span>Email</span>
            <input name="email" type="email" required />
          </label>
          <label>
            <span>Password</span>
            <input name="password" type="password" required />
          </label>
          <button type="submit">Register</button>
        </form>
      ) : (
        <form className="form" onSubmit={submitLogin}>
          <label>
            <span>Email</span>
            <input name="email" type="email" required />
          </label>
          <label>
            <span>Password</span>
            <input name="password" type="password" required />
          </label>
          <button type="submit">Login</button>
        </form>
      )}
    </div>
  );
}

export default App;
