import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import { useAuth } from '../store/auth';
import LoginPage from '../components/pages/login/index.vue';
import Dashboard from '../components/pages/dashboard/index.vue';
import Registration from '../components/pages/registration/index.vue';
import Detail from '../components/pages/dashboard/detail.vue';
import MyEvent from '../components/pages/myevent/index.vue';

declare module "vue-router" {
  interface RouteMeta {
    requiresAuth?: boolean;
    isRedirect?: boolean;
    title?: string;
    icon?: string;
  }
}

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Dashboard',
    component: Dashboard,
    meta: { requiresAuth: true },  
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginPage,
  },
  {
    path: '/registration',
    name: 'Registration',
    component: Registration,
  },
  {
    path: '/detail/:id',
    name: 'Detail',
    component: Detail,
    meta: { requiresAuth: true },  
  },
  {
    path: '/myevent',
    name: 'MyEvent',
    component: MyEvent,
    meta: { requiresAuth: true },  
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, _, next) => {
  const authStore = useAuth();
  const isAuthenticated = authStore.getIsLogin; 
  
  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login');  
  } else {
    next();  
  }
});

export default router;
