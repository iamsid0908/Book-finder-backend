import React, { useEffect, useState } from "react";
import axios from "axios";
// import "./Cutomer.css";

function Billing() {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const token = localStorage.getItem("Bearer");
        const response = await axios.get(
          "http://localhost:8000/v1/billing/get",
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        );
        setData(response.data.data);
        console.log(response.data);
      } catch (error) {
        alert(error);
        console.error("Failed to fetch data", error);
      }
    };

    fetchData();
  });
  console.log(data);
  return (
    <div>
      <div className="cutom">
        {data.map((item) => (
          <div className="cutomer-card">
            <h1>{item.id}</h1>
            <h1>{item.amount}</h1>
            <p>{item.number}</p>
            <p>{item.payment}</p>
            <p>{item.status}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Billing;
