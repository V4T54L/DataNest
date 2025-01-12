import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { baseURL } from "../constants/apiConstants";

const ListDashboards = () => {
  const [dashboards, setDashboards] = useState([]);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(false);
  const [successMessage, setSuccessMessage] = useState(null);
  const navigate = useNavigate()

  const viewDashboardPage = (id) => {
    navigate("/" + id)
  }

  useEffect(() => {
    const fetchDashboards = async () => {
      const endpoint = baseURL+'/dashboards';
      setLoading(true);
      try {
        const response = await fetch(endpoint, {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: 'include',
        });

        const data = await response.json();

        if (response.ok && data.data) {
          console.log("Success:", data);
          setDashboards(data.data);
          setError(null);
        } else {
          setError(data.message || "Failed to load dashboards");
          console.error("Error:", data);
        }
      } catch (err) {
        setError("An unexpected error occurred");
        console.error("Fetch error:", err);
      } finally {
        setLoading(false);
      }
    };

    fetchDashboards();
  }, []);

  const createDashboard = async () => {
    const endpoint = baseURL+'/dashboards';
    const name = prompt("New Dashboard Name:", "").trim();

    if (!name) {
      return;
    }

    try {
      const response = await fetch(endpoint, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: 'include',
        body: JSON.stringify({ name }),
      });

      const data = await response.json();

      if (response.ok && data.message) {
        const arr = data.message.split(":")
        const id = parseInt(arr[arr.length - 1])
        setDashboards((prev) => [...prev, { id, name }]);
        setSuccessMessage(data.message);
        setError(null);
      } else {
        setError(data.error || "Failed to create dashboard");
        console.error("Error:", data);
      }
    } catch (err) {
      setError("An unexpected error occurred while creating the dashboard");
      console.error("Fetch error:", err);
    }
  };

  return (
    <>
      {error && <div className="text-red-500 mb-4">{error}</div>}
      {successMessage && <div className="text-green-500 mb-4">{successMessage}</div>} {/* Display success message */}
      {loading && <div className="spinner">Loading...</div>} {/* Optional loading spinner */}
      <button
        onClick={createDashboard}
        className="bg-blue-500 text-white py-2 px-4 rounded mb-4 hover:bg-blue-600 transition"
      >
        Create Dashboard
      </button>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        {dashboards.map(dashboard => (
          <div
          onClick={()=>viewDashboardPage(dashboard.id)}
            key={dashboard.id}
            className="bg-white shadow-md rounded-lg p-4 hover:shadow-lg transition"
          >
            <h2 className="text-lg font-semibold">{dashboard.name}</h2>
          </div>
        ))}
      </div>
    </>
  );
};

export default ListDashboards;