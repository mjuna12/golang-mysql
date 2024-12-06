<template>
  <Navbar/>
  <div class="chart-container">
    <div v-if="loading" class="loading">Loading...</div>
    <div v-else>
      <div v-if="chartData && chartData.labels.length">
        <BarChart :chart-data="chartData" :options="chartOptions" />
      </div>
      <div v-else>
        <p>No data available</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import Navbar from "../../global/navbar/index.vue";
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale } from "chart.js";
import axios from "axios";

// Register chart.js components
ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale);

const chartData = ref<any>(null);
const chartOptions = ref<any>({
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      position: "top",
    },
    tooltip: {
      callbacks: {
        label: function (tooltipItem: any) {
          return `${tooltipItem.dataset.label}: ${tooltipItem.raw}`;
        },
      },
    },
  },
});

const loading = ref(true);

const fetchData = async () => {
  try {
    const response = await axios.get("http://localhost:3000/soal/1");
    const data = response.data.data; // Assuming response contains the 'data' array as per the sample

    // Prepare chart data
    chartData.value = {
      labels: data.map((item: any) => item.course.course_name), // Courses
      datasets: [
        {
          label: "Jumlah Peserta Didik",
          data: data.map((item: any) => item.user_courses.length), // Number of students
          backgroundColor: "rgba(75, 192, 192, 0.2)", // Bar color
          borderColor: "rgba(75, 192, 192, 1)",
          borderWidth: 1,
        },
      ],
    };

    loading.value = false; // Data loaded
  } catch (error) {
    console.error("Error fetching data:", error);
    loading.value = false;
  }
};

const fetchData1 = async () => {
  try {
    const response = await axios.get("http://localhost:3000/soal/2");
    const data = response.data.data; // Assuming response contains the 'data' array as per the sample

    // Prepare chart data
    chartData.value = {
      labels: data.map((item: any) => item.course.course_name), // Courses
      datasets: [
        {
          label: "Jumlah Peserta Didik",
          data: data.map((item: any) => item.user_courses.length), // Number of students
          backgroundColor: "rgba(75, 192, 192, 0.2)", // Bar color
          borderColor: "rgba(75, 192, 192, 1)",
          borderWidth: 1,
        },
      ],
    };

    loading.value = false; // Data loaded
  } catch (error) {
    console.error("Error fetching data:", error);
    loading.value = false;
  }
};

const fetchData3 = async () => {
  try {
    const response = await axios.get("http://localhost:3000/soal/3");
    const data = response.data.data; // Assuming response contains the 'data' array as per the sample

    // Prepare chart data
    chartData.value = {
      labels: data.map((item: any) => item.course.course_name), // Courses
      datasets: [
        {
          label: "Jumlah Peserta Didik",
          data: data.map((item: any) => item.user_courses.length), // Number of students
          backgroundColor: "rgba(75, 192, 192, 0.2)", // Bar color
          borderColor: "rgba(75, 192, 192, 1)",
          borderWidth: 1,
        },
      ],
    };

    loading.value = false; // Data loaded
  } catch (error) {
    console.error("Error fetching data:", error);
    loading.value = false;
  }
};

onMounted(() => {
  fetchData();
  fetchData1();
  fetchData3();

});
</script>

<style scoped>
.chart-container {
  width: 100%;
  height: 400px;
}

.loading {
  font-size: 1.5em;
  color: #aaa;
}
</style>
