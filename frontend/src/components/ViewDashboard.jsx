import { Bar, Pie, Line, Radar, Doughnut, Scatter, Bubble } from 'react-chartjs-2';
import { Chart as ChartJS, BarElement, CategoryScale, LinearScale, Title, Tooltip, Legend, ArcElement, PointElement, LineElement, RadialLinearScale, BubbleController } from 'chart.js';
import chartDataArray from '../constants/mockChartData';

ChartJS.register(BarElement, CategoryScale, LinearScale, Title, Tooltip, Legend, ArcElement, PointElement, LineElement, RadialLinearScale, BubbleController);

const renderChart = (chart) => {
  const { chartType, data } = chart;

  switch (chartType) {
    case 'bar':
      return <Bar data={data} options={{ responsive: true }} />;
    case 'pie':
      return <Pie data={data} options={{ responsive: true }} />;
    case 'line':
      return <Line data={data} options={{ responsive: true }} />;
    case 'radar':
      return <Radar data={data} options={{ responsive: true }} />;
    case 'doughnut':
      return <Doughnut data={data} options={{ responsive: true }} />;
    case 'scatter':
      return <Scatter data={data} options={{ responsive: true }} />;
    case 'bubble':
      return <Bubble data={data} options={{ responsive: true }} />;
    default:
      return null;
  }
};

const ViewDashboard = () => {

  const handleCreateChart = () => {
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
        {chartDataArray.map((chart, index) => (
          <div className="bg-white shadow-md rounded-lg p-4" key={index}>
            <h2 className="text-lg font-semibold mb-2">{chart.chartType.charAt(0).toUpperCase() + chart.chartType.slice(1)} Chart</h2>
            {renderChart(chart)}
          </div>
        ))}
      </div>
    </div>
  );
};

export default ViewDashboard;