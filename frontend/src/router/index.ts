import { createRouter, createWebHistory } from 'vue-router';
import { useAuthStore } from '../stores/auth';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: () => import('../views/Login.vue'),
    },
    {
      path: '/register',
      name: 'Register',
      component: () => import('../views/Register.vue'),
    },
    {
      path: '/',
      name: 'Dashboard',
      component: () => import('../views/Dashboard.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/inventory',
      name: 'Inventory',
      component: () => import('../views/Inventory.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/recipes',
      name: 'Recipe Management',
      component: () => import('../views/Recipes.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/products',
      name: 'Production & Stock',
      component: () => import('../views/Products.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/pos',
      name: 'POS',
      component: () => import('../views/POS.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/preparations',
      name: 'Preparation History',
      component: () => import('../views/PreparationHistory.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/settings',
      name: 'Settings',
      component: () => import('../views/Settings.vue'),
      meta: { requiresAuth: true },
    },
  ],
});

router.beforeEach(async (to, _from, next) => {
  const authStore = useAuthStore();
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next('/login');
  } else if ((to.path === '/login' || to.path === '/register') && authStore.isAuthenticated) {
    next('/');
  } else {
    next();
  }
});

export default router;
