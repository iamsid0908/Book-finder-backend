import React, { useEffect, useState } from "react";
import axios from "axios";

function User() {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const token = localStorage.getItem("Bearer");
        const response = await axios.get("http://localhost:8000/v1/user/list", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

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
              <h1>{item.id}</h1>
              <h1>{item.email}</h1>
              <p>{item.name}</p>
              <p>{item.language}</p>
              <p>{item.roles}</p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default User;
