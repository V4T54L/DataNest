import { Bar, Pie, Line, Radar, Doughnut, Scatter, Bubble } from 'react-chartjs-2';
import { Chart as ChartJS, BarElement, CategoryScale, LinearScale, Title, Tooltip, Legend, ArcElement, PointElement, LineElement, RadialLinearScale, BubbleController } from 'chart.js';
import chartDataArray from '../constants/mockChartData';
import { useParams } from 'react-router-dom';
import { useEffect, useState } from 'react';
import { baseURL } from '../constants/apiConstants';

ChartJS.register(BarElement, CategoryScale, LinearScale, Title, Tooltip, Legend, ArcElement, PointElement, LineElement, RadialLinearScale, BubbleController);

const renderChart = (chart) => {
  const { type, data } = chart;

  switch (type) {
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
  const { dashboardId } = useParams()
  const [charts, setCharts] = useState([])
  const [dashboardName, setDashboardName] = useState("")
  const [codePreviewIdx, setCodePreviewIdx] = useState(0)


  useEffect(() => {
    const fetchCharts = async () => {
      const endpoint = baseURL + '/dashboards/' + dashboardId;
      try {
        const response = await fetch(endpoint, {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: 'include',
        });

        const data = await response.json();

        if (response.ok && data?.data) {
          if (data?.data?.name) {
            setDashboardName(data.data.name)
          }
          if (data?.data?.charts) {
            let chartsData = data.data.charts.map((chart) => {
              chart.data = JSON.parse(chart.data);
              return chart;
            });

            setCharts(chartsData);
          }
          console.log("Success:", data);
        } else {
          console.error("Error:", data);
        }
      } catch (err) {
        console.error("Fetch error:", err);
      }
    };

    fetchCharts();
  }, [dashboardId]);

  const createChart = async () => {
    const endpoint = baseURL + `/dashboards/${dashboardId}/charts`;
    const type = prompt("ChartType:", "").trim();
    const chartData = prompt("ChartData:", "").trim();

    if (!type || !chartData) {
      return;
    }

    try {
      const response = await fetch(endpoint, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: 'include',
        body: JSON.stringify({ type, data: chartData }),
      });

      const data = await response.json();

      if (response.ok && data.message) {
        const arr = data.message.split(":")
        const id = parseInt(arr[arr.length - 1])
        setCharts(prev => [...prev, { id, type, data: JSON.parse(chartData) }])
      } else {
        console.error("Error:", data);
      }
    } catch (err) {
      console.error("Fetch error:", err);
    }
  }

  useEffect(() => {
    console.log("Charts : ", charts)
  }, [charts])

  if (!dashboardName) {
    return <div>Loading...</div>;
  }

  return (
    <div className="p-4">
      <h2>{dashboardName}</h2>
      <button
        onClick={createChart}
        className="bg-blue-500 text-white py-2 px-4 rounded mb-4 hover:bg-blue-600 transition"
      >
        Create New Chart
      </button>

      {/* Render Example Chart Data For Reference */}
      <div>
        {
          chartDataArray &&
          <>
            <h2 className='m-2 text-lg font-semibold'>Examples : </h2>
            <div className='flex gap-2 my-2'>
              {
                chartDataArray.map((e, idx) => (
                  <div
                    className='px-2 py-0 bg-gray-300 rounded-md shadow-md'
                    onClick={() => setCodePreviewIdx(idx)} key={e.chartType}>
                    {e.chartType}
                  </div>
                ))
              }
            </div>
            <div className='bg-gray-200 px-4 py-2 w-1/2 mt-4 mb-8'>
              <pre className='text-xs text-wrap'>{JSON.stringify(chartDataArray[codePreviewIdx].data)}</pre>
            </div>
          </>
        }
      </div>

      {/* Render Chart */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        {
          charts &&
          charts.map((chart) => (
            <div className="bg-white shadow-md rounded-lg p-4" key={chart.id}>
              <h2 className="text-lg font-semibold mb-2">{chart.type.charAt(0).toUpperCase() + chart.type.slice(1)} Chart</h2>
              {renderChart(chart)}
            </div>
          ))}
      </div>
    </div>
  );
};

export default ViewDashboard;