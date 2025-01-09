const chartDataArray = [
    {
      chartType: 'bar',
      data: {
        labels: ['January', 'February', 'March', 'April', 'May'],
        datasets: [{
          label: 'Sales Data',
          data: [12, 19, 3, 5, 2],
          backgroundColor: 'rgba(75, 192, 192, 0.6)',
          borderColor: 'rgba(75, 192, 192, 1)',
          borderWidth: 1,
        }],
      },
    },
    {
      chartType: 'pie',
      data: {
        labels: ['Red', 'Blue', 'Yellow', 'Green', 'Purple', 'Orange'],
        datasets: [{
          label: '# of Votes',
          data: [12, 19, 3, 5, 2, 3],
          backgroundColor: [
            'rgba(255, 99, 132, 0.6)',
            'rgba(54, 162, 235, 0.6)',
            'rgba(255, 206, 86, 0.6)',
            'rgba(75, 192, 192, 0.6)',
            'rgba(153, 102, 255, 0.6)',
            'rgba(255, 159, 64, 0.6)',
          ],
          borderColor: [
            'rgba(255, 99, 132, 1)',
            'rgba(54, 162, 235, 1)',
            'rgba(255, 206, 86, 1)',
            'rgba(75, 192, 192, 1)',
            'rgba(153, 102, 255, 1)',
            'rgba(255, 159, 64, 1)',
          ],
          borderWidth: 1,
        }],
      },
    },
    {
      chartType: 'line',
      data: {
        labels: ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'],
        datasets: [{
          label: 'Weekly Trends',
          data: [65, 59, 80, 81, 56, 55, 40],
          fill: false,
          backgroundColor: 'rgba(75, 192, 192, 0.4)',
          borderColor: 'rgba(75, 192, 192, 1)',
          borderWidth: 2,
        }],
      },
    },
    {
      chartType: 'radar',
      data: {
        labels: ['Running', 'Swimming', 'Eating', 'Cycling'],
        datasets: [{
          label: 'My Daily Activities',
          data: [20, 10, 4, 2],
          backgroundColor: 'rgba(255, 99, 132, 0.2)',
          borderColor: 'rgba(255, 99, 132, 1)',
          borderWidth: 2,
        }],
      },
    },
    {
      chartType: 'doughnut',
      data: {
        labels: ['Red', 'Blue', 'Yellow'],
        datasets: [{
          label: '# of Votes',
          data: [300, 50, 100],
          backgroundColor: [
            'rgba(255, 99, 132, 0.6)',
            'rgba(54, 162, 235, 0.6)',
            'rgba(255, 206, 86, 0.6)',
          ],
          borderColor: [
            'rgba(255, 99, 132, 1)',
            'rgba(54, 162, 235, 1)',
            'rgba(255, 206, 86, 1)',
          ],
          borderWidth: 1,
        }],
      },
    },
    {
      chartType: 'scatter',
      data: {
        datasets: [{
          label: 'Scatter Dataset',
          data: [
            { x: -10, y: 0 },
            { x: 0, y: 10 },
            { x: 10, y: 5 },
          ],
          backgroundColor: 'rgba(75, 192, 192, 0.6)',
        }],
      },
    },
    {
      chartType: 'bubble',
      data: {
        datasets: [{
          label: 'Bubble Dataset',
          data: [
            { x: 20, y: 30, r: 15 },
            { x: 30, y: 20, r: 10 },
            { x: 50, y: 50, r: 20 },
          ],
          backgroundColor: 'rgba(255, 206, 86, 0.6)',
        }],
      },
    },
  ];
  
  export default chartDataArray;