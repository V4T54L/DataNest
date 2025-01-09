// src/Charts.jsx
const ViewDashboard = ({dashboardId}) => {
  console.log("Dashboard ID : ", dashboardId)

  const charts = [
    { id: 1, title: 'Chart 1' },
    { id: 2, title: 'Chart 2' },
    { id: 3, title: 'Chart 3' },
    { id: 4, title: 'Chart 4' },
    { id: 5, title: 'Chart 5' },
    { id: 6, title: 'Chart 6' },
  ];

  const handleCreateChart = () => {
    // Add your create chart logic here
    console.log('Create a new chart');
  };

  return (
    <div className="p-4">
      <button
        onClick={handleCreateChart}
        className="bg-blue-500 text-white py-2 px-4 rounded mb-4 hover:bg-blue-600 transition"
      >
        Create New Chart
      </button>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        {charts.map((chart) => (
          <div
            key={chart.id}
            className="bg-white shadow-md rounded-lg p-4 hover:shadow-lg transition"
          >
            <h2 className="text-lg font-semibold">{chart.title}</h2>
            <div className="h-32 bg-gray-300 flex items-center justify-center rounded">
              {/* Placeholder for a chart */}
              <span className="text-gray-600">Chart Placeholder</span>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default ViewDashboard;