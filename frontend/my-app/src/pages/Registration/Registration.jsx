import React, { useState } from "react";
import "./Registration.css";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";
import { Link } from "react-router-dom";

function Registration() {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [name, setName] = useState("");
  const [password, setPassword] = useState("");
  const [role, setRole] = useState("");
  const user = {
    email,
    name,
    password,
    role: parseInt(role, 10),
  };

  const RegisterUser = (e) => {
    e.preventDefault();
    fetch("http://localhost:8000/v1/auth/register", {
      method: "POST",
      body: JSON.stringify(user),
      mode: "cors",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((response) => {
        if (!response.ok) {
          return response.json().then((errorData) => {
            throw new Error(errorData.message || "Something went wrong!");
          });
        }
        return response.json();
      })
      .then((data) => {
        navigate("/login");
      })
      .catch((err) => {
        alert(err.message);
      });
  };
  return (
    <div className="login">
      <div className="loginform">
        <h1>Registration form</h1>

        <label>Email</label>
        <Input
          value={email}
          onChange={(e) => {
            setEmail(e.target.value);
          }}
        />
        <label>Name</label>
        <Input
          value={name}
          onChange={(e) => {
            setName(e.target.value);
          }}
        />
        <label>Password</label>
        <Input
          value={password}
          onChange={(e) => {
            setPassword(e.target.value);
          }}
        />
        <label>Select role</label>
        <select
          name="role"
          id="roles"
          onChange={(e) => {
            setRole(e.target.value);
          }}
        >
          <option value="1">Sales</option>
          <option value="2">Accountant</option>
          <option value="3">HR</option>
          <option value="4">Administrator</option>
          <option value="5">Customer</option>
        </select>

        <Button onClick={RegisterUser}> SignUp</Button>
        <p>Already have an account?</p>

        <Link to="/login">
          <Button>Login</Button>
        </Link>
      </div>
    </div>
  );
}

export default Registration;
