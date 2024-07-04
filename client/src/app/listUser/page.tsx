"use client";

import React, { useEffect, useState } from "react";
import axios from "axios";

interface User {
  id: string;
  name: string;
  email: string;
}

const ListUser = () => {
  const [user, setUser] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);

  // Check auth
  useEffect(() => {
    const fetchData = async () => {
      const token = localStorage.getItem(
        "token-first-web-with-golang-playground"
      );
      console.log("token : ", token);
      try {
        const response = await axios.get("http://localhost:8080/users", {
          headers: {
            Authorization: `Bearer ${localStorage.getItem(
              "token-first-web-with-golang-playground"
            )}`,
          },
        });
        console.log(response.data.data);
        setLoading(false);
        setUser(response.data.data);
      } catch (error) {
        console.log(error);
      }
    };
    const token = localStorage.getItem(
      "token-first-web-with-golang-playground"
    );
    if (!token) {
      window.location.href = "/";
    }
    fetchData();
  }, []);

  // Logout
  const handleLogout = () => {
    localStorage.removeItem("token-first-web-with-golang-playground");
    window.location.href = "/";
  };
  return (
    <div className="w-screen min-h-screen bg-[#B2B1CF] px-16 pt-8">
      <div className="relative h-16 flex items-center ">
        <p className="w-full text-center text-2xl font-bold text-black">
          List User
        </p>
        <button
          className="bg-red-600 hover:bg-red-800 absolute top-3 px-4 h-10 right-4 rounded-lg"
          onClick={() => handleLogout()}
        >
          Log out
        </button>
      </div>
      {loading ? (
        <div className="w-full h-24 flex justify-center items-center">
          <span className="loading loading-spinner loading-md"></span>
        </div>
      ) : (
        <table className="w-full text-center mt-8">
          <thead className="">
            <tr className="bg-[#77625C]">
              <th className="py-2">Id</th>
              <th className="py-2">Name</th>
              <th className="py-2">Email</th>
            </tr>
          </thead>
          <tbody>
            {user?.map((user) => (
              <tr key={user.id} className="bg-[#49392C] text-white border-t-2">
                <td className="py-4">{user.id}</td>
                <td className="py-4">{user.name}</td>
                <td className="py-4">{user.email}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
};

export default ListUser;
