<script setup lang="ts">
import { useAuth } from '../../../store/auth';
import Navbar from '../../../components/global/navbar/index.vue';

const authStore = useAuth();
const handleRegisterEvent = (id: any) => {
  try {
    const randomNumber1 = Math.floor(Math.random() * 100) + 1;
    const randomString = (Math.random() + 1);
    authStore.registerEvent(id, randomNumber1, randomString);
  } catch (error) {
    console.log(error)
  }
}

</script>

<template>
  <Navbar />
  <div class="dark:bg-gray-900">
    <div class="max-w-screen-xl mx-auto p-4">
      <div class="bg-white rounded-lg shadow-md dark:bg-gray-800">
        <div class="flex flex-col md:flex-row">
          <img class="w-full md:w-1/2 h-64 object-cover rounded-t-lg md:rounded-l-lg" :src="authStore.selectedEvent.event_picture" alt="Event Image" />
          <div class="p-6 flex flex-col justify-between">
            <h1 class="text-3xl font-bold dark:text-white">{{ authStore.selectedEvent.event_name }}</h1>
            <p class="mt-4 text-lg text-gray-700 dark:text-gray-400">{{ authStore.selectedEvent.event_desc }}</p>
            <p class="mt-4 text-gray-600 dark:text-gray-300">Event Start: {{ authStore.selectedEvent.event_start }}</p>
            <p class="mt-4 text-gray-600 dark:text-gray-300">Event End: {{ authStore.selectedEvent.event_end }}</p>
            <div class="mt-6">
              <button class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 rounded-lg py-2" @click="handleRegisterEvent(authStore.selectedEvent.event_id)">
                Register for Event
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
