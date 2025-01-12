// src/Auth.jsx
import { useState } from "react";
import { baseURL } from "../constants/apiConstants";

const Auth = ({ setUser }) => {
  const [isLogin, setIsLogin] = useState(true);
  const [formData, setFormData] = useState({
    name: "",
    username: "",
    email: "",
    password: "",
    confirmPassword: ""
  });

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const endpoint = baseURL + (isLogin ? "/auth/login" : "/auth/signup");
    console.log("Endpoint : ", endpoint)
    const response = await fetch(endpoint, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: 'include',
      body: JSON.stringify(formData),
    });

    const data = await response.json();

    if (response.ok) {
      console.log("Success:", data);
      if (isLogin) {
        setUser(data.data)
      }
    } else {
      console.error("Error:", data);
    }
  };

  return (
    <div className="flex items-center justify-center h-screen bg-gray-100">
      <div className="bg-white rounded-lg shadow-lg w-96 transition-transform duration-500 transform 
                      ${isLogin ? 'scale-100' : 'scale-110'}">
        <div className="p-6">
          <h2 className="text-2xl font-bold text-center mb-6">
            {isLogin ? "Login" : "Sign Up"}
          </h2>
          <form onSubmit={handleSubmit}>
            {!isLogin && (
              <>
                <div className="mb-4">
                  <label className="block mb-1" htmlFor="name">
                    Name
                  </label>
                  <input
                    type="text"
                    id="name"
                    name="name"
                    value={formData.name}
                    onChange={handleChange}
                    required
                    className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                  />
                </div>
              </>
            )}
            <div className="mb-4">
              <label className="block mb-1" htmlFor="username">
                Username
              </label>
              <input
                type="text"
                id="username"
                name="username"
                value={formData.username}
                onChange={handleChange}
                required
                className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
            {!isLogin && (
              <div className="mb-4">
                <label className="block mb-1" htmlFor="email">
                  Email
                </label>
                <input
                  type="email"
                  id="email"
                  name="email"
                  value={formData.email}
                  onChange={handleChange}
                  required
                  className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
            )}
            <div className="mb-4">
              <label className="block mb-1" htmlFor="password">
                Password
              </label>
              <input
                type="password"
                id="password"
                name="password"
                value={formData.password}
                onChange={handleChange}
                required
                className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
            {!isLogin && (
              <div className="mb-4">
                <label className="block mb-1" htmlFor="confirmPassword">
                  Confirm Password
                </label>
                <input
                  type="password"
                  id="confirmPassword"
                  name="confirmPassword"
                  value={formData.confirmPassword}
                  onChange={handleChange}
                  required
                  className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
            )}
            <button
              type="submit"
              className="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600 transition"
            >
              {isLogin ? "Login" : "Sign Up"}
            </button>
          </form>
          <div className="mt-4 text-center">
            <button
              onClick={() => setIsLogin(!isLogin)}
              className="text-blue-500 hover:underline"
            >
              {isLogin ? "Create an account" : "Already have an account?"}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Auth;