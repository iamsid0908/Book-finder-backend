import "./App.css";
import { Route, Routes } from "react-router-dom";
import Login from "./pages/Login/Login";
import Home from "./pages/Home/Home";
import Registration from "./pages/Registration/Registration";
import Customer from "./pages/Customer/Customer";
import Billing from "./pages/Billing/Billing";
import User from "./pages/User/User";
import Payroll from "./pages/Payroll/Payroll";

function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Registration />} />
        <Route path="/cutomer" element={<Customer />} />
        <Route path="/billing" element={<Billing />} />
        <Route path="/user" element={<User />} />
        <Route path="/payroll" element={<Payroll />} />
      </Routes>
    </>
  );
}

export default App;
