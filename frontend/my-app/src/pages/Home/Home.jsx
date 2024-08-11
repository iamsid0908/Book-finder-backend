import React, { useEffect, useState } from "react";
import "./Home.css";
import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";

function Home() {
  const [name, setName] = useState("");
  const [role, setRole] = useState("");
  const [email, setEmail] = useState("");

  useEffect(() => {
    setName(localStorage.getItem("name"));
    setRole(localStorage.getItem("role"));
    setEmail(localStorage.getItem("email"));
  }, []);

  return (
    <div className="home">
      <div className="home-con">
        <div>
          <p>Name:{name}</p>
          <p>Email:{email}</p>
          <p>Role:{role}</p>
        </div>
      </div>
      <Link to="/login">
        <Button>Login</Button>
      </Link>
      <Link to="/register">
        <Button>Register</Button>
      </Link>

      <div className="home-management">
        <p className="management">
          <Link to="/cutomer">Cutomer Management</Link>
        </p>
        <p className="management">
          <Link to="/billing">Billing Management</Link>
        </p>
        <p className="management">
          <Link to="/payroll">Payroll Management</Link>
        </p>
        <p className="management">
          {" "}
          <Link to="/user">User Management</Link>
        </p>
      </div>
    </div>
  );
}

export default Home;
