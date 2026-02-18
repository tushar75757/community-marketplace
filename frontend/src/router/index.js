import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import { useAuthStore } from '../stores/auth';

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home,
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('../views/Login.vue'),
        meta: { guestOnly: true },
    },
    {
        path: '/register',
        name: 'Register',
        component: () => import('../views/Register.vue'),
        meta: { guestOnly: true },
    },
    {
        path: '/listings/create',
        name: 'CreateListing',
        component: () => import('../views/CreateListing.vue'),
        meta: { requiresAuth: true },
    },
    {
        path: '/listings/:id',
        name: 'ItemDetail',
        component: () => import('../views/ItemDetail.vue'),
    },
    {
        path: '/listings/:id/edit',
        name: 'EditListing',
        component: () => import('../views/EditListing.vue'),
        meta: { requiresAuth: true },
    },
    {
        path: '/messages',
        name: 'Messages',
        component: () => import('../views/Messages.vue'),
        meta: { requiresAuth: true },
    },
    {
        path: '/admin',
        name: 'AdminDashboard',
        component: () => import('../views/AdminDashboard.vue'),
        meta: { requiresAdmin: true },
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    const authStore = useAuthStore();

    if (to.meta.requiresAuth && !authStore.isAuthenticated) {
        next('/login');
    } else if (to.meta.requiresAdmin && !authStore.isAdmin) {
        next('/');
    } else if (to.meta.guestOnly && authStore.isAuthenticated) {
        next('/');
    } else {
        next();
    }
});

export default router;
