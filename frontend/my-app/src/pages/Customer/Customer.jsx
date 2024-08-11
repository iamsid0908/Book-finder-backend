import React, { useEffect, useState } from "react";
import axios from "axios";
import "./Cutomer.css";

function Customer() {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const token = localStorage.getItem("Bearer");

        const response = await axios.get(
          "http://localhost:8000/v1/customer/get",
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        );
        setData(response.data.data);
      } catch (error) {
        alert(error);
        console.error("Failed to fetch data", error);
      }
    };

    fetchData();
  }, []);
  return (
    <div>
      <div className="cutom">
        {data.map((item) => (
          <div className="cutomer-card">
            <h1>{item.ID}</h1>
            <h1>{item.Address}</h1>
            <p>{item.LastOrder}</p>
            <p>{item.Phone}</p>
            <p>{item.BillingIds}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Customer;
