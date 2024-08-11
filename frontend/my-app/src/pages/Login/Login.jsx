import React, { useState } from "react";
import "./Login.css";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";
import { Link } from "react-router-dom";

function Login() {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  console.log(email);
  const user = {
    email,
    password,
  };

  const LoginUser = (e) => {
    e.preventDefault();
    fetch("http://localhost:8000/v1/auth/login", {
      method: "POST",
      body: JSON.stringify(user),
      mode: "cors",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((data) => data.json())
      .then((data) => {
        localStorage.setItem("id", data.data.id);
        localStorage.setItem("name", data.data.name);
        localStorage.setItem("role", data.data.role);
        localStorage.setItem("email", data.data.email);
        localStorage.setItem("Bearer", data.data.token);
        navigate("/");
      })
      .catch((err) => {
        alert(err);
      });
  };

  return (
    <div className="login">
      <div className="loginform">
        <h1>Login form</h1>

        <label>Email</label>
        <Input
          value={email}
          onChange={(e) => {
            setEmail(e.target.value);
          }}
        />
        <label>Password</label>
        <Input
          value={password}
          onChange={(e) => {
            setPassword(e.target.value);
          }}
        />
        <Button onClick={LoginUser}>Submit</Button>
        <p>Don't have an account</p>
        <Link to="/register">
          <Button>Register</Button>
        </Link>
      </div>
    </div>
  );
}

export default Login;
