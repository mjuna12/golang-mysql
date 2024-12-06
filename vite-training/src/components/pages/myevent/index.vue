<script setup lang="ts">
import { useAuth } from '../../../store/auth';
import Navbar from '../../../components/global/navbar/index.vue';
import { onMounted } from 'vue';
import Button from '../../global/button/index.vue'
import { useRouter } from 'vue-router';
// import Card from '../../global/card/index.vue'

const authStore = useAuth();
const router = useRouter();

const goToDetail = (id:number) => {
  router.push({ name: 'Detail', params: { id } });
}

onMounted(() => {
  authStore.getMyEvent();
});
</script>

<template>
  <Navbar />
  <div class="dark:bg-gray-900 pl-12 space-y-8">
    <h1 class="text-3xl font-bold mb-6 dark:text-white flex items-center justify-center">My Event</h1>
    <div v-for="(item, index) in authStore.list_myevent" :key="index" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6  transform transition-all duration-300 ease-in hover:scale-105">
        <a href="#" class="flex-shrink-0">
          <img class="rounded-t-lg w-full h-48 object-cover" :src="item?.event?.event_picture" alt="Event Image" />
        </a>
        <div class="flex-grow p-5 flex flex-col justify-between">
          <a href="#" class="mb-2">
            <h5 class="text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{{ item?.event?.event_name }}</h5>
          </a>
          <p class="mb-3 font-normal text-gray-700 dark:text-gray-400 line-clamp-6">{{ item?.event?.event_desc }}</p>
          <div class="space-y-4">
          <Button text="Read More" type="primary" customClass="w-full text-white bg-slate-700 rounded-lg hover:bg-slate-800 focus:ring-4 focus:outline-none focus:ring-slate-300 dark:bg-slate-600 dark:hover:bg-slate-700 dark:focus:ring-blue-800" @click="goToDetail(item?.event.event_id)"/>
        </div>  
        </div>
    </div>
  </div>
</template>
