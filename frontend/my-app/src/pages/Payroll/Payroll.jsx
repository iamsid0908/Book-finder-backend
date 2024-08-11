import React, { useEffect, useState } from "react";
import axios from "axios";

function Payroll() {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const token = localStorage.getItem("Bearer");
        const response = await axios.get(
          "http://localhost:8000/v1/payment/get",
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
      <div>
        <div className="cutom">
          {data.map((item) => (
            <div className="cutomer-card">
              <p>UserID</p>
              <h1>{item.user_id}</h1>
              <p>Bill</p>
              <h1>{item.billing_id}</h1>
              <p>Amount</p>
              <p>{item.amount}</p>
              <p>Method</p>
              <p>{item.method}</p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
export default Payroll;
