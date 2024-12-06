<script setup lang="ts">
import { useAuth } from '../../../store/auth';
import Navbar from '../../../components/global/navbar/index.vue';
import { onMounted } from 'vue';
import Button from '../../global/button/index.vue'
import { useRouter } from 'vue-router';

const authStore = useAuth();
const router = useRouter();

const goToDetail = (id: number) => {
  router.push({ name: 'Detail', params: { id } });
}
const handleRegisterEvent = (id: any) => {
  try {
    const randomNumber1 = Math.floor(Math.random() * 100) + 1;
    const randomString = (Math.random() + 1);
    authStore.registerEvent(id, randomNumber1, randomString);
  } catch (error) {
    console.log(error)
  }
}

onMounted(() => {
  authStore.getEventList();
});
</script>

<template>
  <Navbar />
  <div class="dark:bg-gray-900 pl-12">
    <h1 class="text-3xl font-bold mb-6 dark:text-white flex items-center justify-center">Public Webinars</h1>
    <div v-for="(item, index) in authStore.list_event" :key="index"
      class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div v-for="(items, index) in item.results as any" :key="index"
        class="max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 transform transition-all duration-300 ease-in hover:scale-110">
        <a href="#" class="flex-shrink-0">
          <img class="rounded-t-lg w-full h-48 object-cover" :src="items?.event_picture" alt="Event Image" />
        </a>
        <div class="flex-grow p-5 flex flex-col justify-between">
          <a href="#" class="mb-2">
            <h5 class="text-2xl font-bold tracking-tight text-gray-900 dark:text-white truncate">{{ items?.event_name }}
            </h5>
          </a>
          <p class="mb-3 font-normal text-gray-700 dark:text-gray-400 line-clamp-3">{{ items?.event_desc }}</p>
          <div class="space-y-4">
            <Button text="Read More" type="primary"
              customClass="w-full text-white bg-slate-700 rounded-lg hover:bg-slate-800 focus:ring-4 focus:outline-none focus:ring-slate-300 dark:bg-slate-600 dark:hover:bg-slate-700 dark:focus:ring-blue-800"
              @click="() => {
                goToDetail(items.event_id)
                authStore.setSelectedEvent(items)
              }" />
            <Button text="Register" type="primary"
              customClass="w-full text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
              @click="handleRegisterEvent(items.event_id)" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
