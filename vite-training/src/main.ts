import { createApp } from 'vue';
import App from './App.vue';
import './index.css';
import router from './router';  // Import your router
import { createPinia } from 'pinia'; // Import createPinia from Pinia

// Create Vue app
const app = createApp(App);

// Initialize Pinia
const pinia = createPinia();
app.use(pinia);

// Initialize the router and use it
app.use(router);

// Mount the app to the DOM
app.mount('#app');
