import { createRouter, createWebHistory } from 'vue-router';
import { useAuthStore } from '../stores/auth';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: () => import('../views/Auth.vue'),
    },
    {
      path: '/register',
      name: 'Register',
      component: () => import('../views/Auth.vue'),
    },
    {
      path: '/',
      component: () => import('../views/DashboardShell.vue'),
      meta: { requiresAuth: true },
      children: [
        {
          path: '',
          name: 'Dashboard',
          component: () => import('../views/Dashboard.vue'),
        },
        {
          path: 'inventory',
          name: 'Inventory',
          component: () => import('../views/Inventory.vue'),
        },
        {
          path: 'recipes',
          name: 'Recipe Management',
          component: () => import('../views/Recipes.vue'),
        },
        {
          path: 'products',
          name: 'Production & Stock',
          component: () => import('../views/Products.vue'),
        },
        {
          path: 'pos/order',
          name: 'POS Order',
          component: () => import('../views/POSOrder.vue'),
        },
        {
          path: 'pos/payment',
          name: 'POS Payment',
          component: () => import('../views/POSPayment.vue'),
        },
        {
          path: 'branches',
          name: 'Branch Manager',
          component: () => import('../views/Branches.vue'),
        },
        {
          path: 'tables',
          name: 'Table Management',
          component: () => import('../views/Tables.vue'),
        },
        {
          path: 'roles',
          name: 'Roles',
          component: () => import('../views/Roles.vue'),
        },
        {
          path: 'employees',
          name: 'Employees',
          component: () => import('../views/Employees.vue'),
        },
        {
          path: 'reports/orders',
          name: 'Order History',
          component: () => import('../views/OrderHistory.vue'),
        },
        {
          path: 'preparations',
          name: 'Production Log',
          component: () => import('../views/ProductionLog.vue'),
        },
        {
          path: 'settings',
          name: 'Settings',
          component: () => import('../views/Settings.vue'),
        },
        {
          path: 'guide',
          name: 'System Guide',
          component: () => import('../views/Guide.vue'),
        },
      ]
    },
  ],
});

router.beforeEach(async (to) => {
  const authStore = useAuthStore();
  
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    return '/login';
  } 
  
  if ((to.path === '/login' || to.path === '/register') && authStore.isAuthenticated) {
    return '/';
  }
});

export default router;
