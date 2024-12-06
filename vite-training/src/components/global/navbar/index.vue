<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useAuth } from '../../../store/auth';
import { useRouter } from 'vue-router';

const isDarkMode = ref(false);
const route = useRouter();
const toggleDarkMode = () => {
  isDarkMode.value = !isDarkMode.value;
  if (isDarkMode.value) {
    document.documentElement.classList.add('dark');
    localStorage.setItem('theme', 'dark');
  } else {
    document.documentElement.classList.remove('dark');
    localStorage.setItem('theme', 'light');
  }
};

const authStore = useAuth()

const handleLogOut = () => {
  try {
    authStore.signOut()
    route.push({name: 'Login'})
  } catch (error) {
    console.log(error)
  }
}

onMounted(() => {
  if (localStorage.getItem('theme') === 'dark') {
    isDarkMode.value = true;
    document.documentElement.classList.add('dark');
  }
});


</script>

<template>
  <nav class="bg-white border-gray-200 dark:bg-gray-900">
    <div class="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
      <a href="https://eduparx.id/home" class="flex items-center space-x-3 rtl:space-x-reverse">
        <img class="mr-4 w-40 object-cover" src="/public/logo-eduparx.png" alt="logo" />
      </a>
      <div class="flex items-center space-x-4">
        <div class="items-center justify-between hidden w-full md:flex md:w-auto md:order-1" id="navbar-user">
          <ul class="flex flex-col font-medium p-4 md:p-0 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:space-x-8 rtl:space-x-reverse md:flex-row md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
            <!-- Home Route -->
            <li>
              <router-link
                to="/"
                class="block py-2 px-3 dark:text-white rounded md:bg-transparent md:text-white-700 md:p-0 md:dark:text-white-500"
                aria-current="page"
              >
              <span class="hover:bg-gray-200 dark:hover:bg-gray-700 p-4 rounded-md">
                Home
              </span>
              </router-link>
            </li>
            <!-- MyEvent Route -->
            <li>
              <router-link
                to="/myevent"
                class="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:hover:text-white-700 md:p-0 dark:text-white md:dark:hover:text-white-500 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700"
              >
              <span class="hover:bg-gray-200 dark:hover:bg-gray-700 p-4 rounded-md">
                My Event
              </span>
              </router-link>
            </li>
            <li>
              <button
                @click="toggleDarkMode()"
                aria-label="Toggle dark mode"
              >
                <span v-if="isDarkMode" class="p-4  hover:bg-gray-200 dark:hover:bg-gray-700 dark:text-white">☀️ Light Mode</span>
                <span v-else class="p-4  hover:bg-gray-200 dark:hover:bg-gray-700 rounded-md dark:text-white">🌙 Dark Mode</span>
              </button>
            </li>
            <li>
              <button @click="handleLogOut()">
                <span class="hover:bg-gray-200 dark:hover:bg-gray-700 dark:text-white p-4 rounded-md">
                  Log Out
                </span>
              </button>
            </li>
            <!-- Add more links here if needed -->
          </ul>
        </div>
      </div>
    </div>
  </nav>
</template>
