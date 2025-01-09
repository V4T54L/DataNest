const dashboards = [
  { id: 1, name: 'Dashboard 1' },
  { id: 2, name: 'Dashboard 2' },
  { id: 3, name: 'Dashboard 3' },
  { id: 4, name: 'Dashboard 4' },
  { id: 5, name: 'Dashboard 5' },
];

const ListDashboards = () => {

  return (
    <>
      <button className="bg-blue-500 text-white py-2 px-4 rounded mb-4 hover:bg-blue-600 transition">
        Create Dashboard
      </button>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        {dashboards.map((dashboard) => (
          <div
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