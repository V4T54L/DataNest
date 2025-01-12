import { useState } from "react";
import {
  BrowserRouter as Router,
  Routes,
  Route,
} from "react-router-dom";
import Auth from "./components/Auth";
import Layout from "./components/Layout";
import ListDashboards from "./components/ListDashboards";
import ViewDashboard from "./components/ViewDashboard";

function App() {
  // State to hold the current user
  const [currentUser, setCurrentUser] = useState(null);

  // PrivateRoute component to handle unauthorized access
  const PrivateRoute = ({ children }) => {
    return currentUser ? children : <Auth setUser={setCurrentUser} />;
  };

  return (
    <Router>
      <Routes>
        {/* Main application routes */}
        <Route 
          path="/" 
          element={
            <PrivateRoute>
              <Layout />
            </PrivateRoute>
          }
        >
          <Route index element={<ListDashboards />} />
          <Route path=":dashboardId" element={<ViewDashboard />} />
        </Route>
      </Routes>
    </Router>
  );
}

export default App;