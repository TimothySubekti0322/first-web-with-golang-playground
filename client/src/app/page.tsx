"use client";

import { useState } from "react";
import axios from "axios";

export default function Home() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = async () => {
    const formData = {
      email: email,
      password: password,
    };
    try {
      const response = await axios.post("http://localhost:8080/auth", formData);

      if (response.data.status == 200) {
        localStorage.setItem(
          "token-first-web-with-golang-playground",
          response.data.data.token
        );
        window.location.href = "/listUser";
      }
    } catch (error) {
      alert(error);
      console.log(error);
    }
  };
  return (
    <main className="flex w-screen min-h-screen flex-col items-center justify-center bg-[#98D2EB]">
      {/* Form */}
      <div className="w-1/4 rounded-xl flex flex-col justify-center bg-[#E1F2FE] text-black px-6 py-10">
        <p className="self-center pb-4 font-bold text-2xl">Login</p>
        {/* email*/}
        <label htmlFor="email" className="text-lg">
          Email
        </label>
        <input
          type="text"
          id="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className="bg-[#B2B1CF] rounded-sm mt-2 py-2 px-4"
        />
        <label htmlFor="password" className="text-lg mt-8">
          Password
        </label>
        {/* Password */}
        <input
          type="text"
          id="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className="bg-[#B2B1CF] rounded-sm mt-2 py-2 px-4"
        />
        <button
          className="bg-[#77625C] hover:bg-[#483b38] w-1/2 px-8 py-2 rounded-full text-white mt-12 self-center font-semibold"
          onClick={() => handleLogin()}
        >
          Login
        </button>
      </div>
    </main>
  );
}
