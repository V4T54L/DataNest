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
  return (
    <>
      <Router>
        <Routes>
          <Route path="/" element={<Auth />} />

          <Route path="/dashboard" element={<Layout />}>
            <Route index element={<ListDashboards />} />
            <Route path=":dashboardId" element={<ViewDashboard />} />
          </Route>
        </Routes>
      </Router>
    </>
  )
}

export default App
