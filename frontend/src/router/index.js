import { createRouter, createWebHistory } from 'vue-router';
import { useAuthStore } from '../store/auth';

// Import komponen-komponen halaman
import Login from '../views/Login.vue';
import UserDashboard from '../views/user/UserDashboard.vue';
import CreateOrder from '../views/user/CreateOrder.vue';
import AdminDashboard from '../views/admin/AdminDashboard.vue';
import Register from '../views/Register.vue';

const routes = [
    {
        path: '/',
        redirect: '/dashboard',
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
    },
    {
        path: '/register',
        name: 'Register',
        component: Register,
    },
    {
        path: '/dashboard',
        name: 'UserDashboard',
        component: UserDashboard,
        meta: { requiresAuth: true }, 
    },
    {
        path: '/orders/create',
        name: 'CreateOrder',
        component: CreateOrder,
        meta: { requiresAuth: true }, 
    },
    {
        path: '/admin',
        name: 'AdminDashboard',
        component: AdminDashboard,
        meta: { requiresAuth: true, requiresAdmin: true },
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: () => import('../views/NotFound.vue') 
    }
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
});

router.beforeEach(async (to, from, next) => {
    const authStore = useAuthStore();

    if (authStore.token && !authStore.user) {
        await authStore.fetchUser();
    }

    const isAuthenticated = authStore.isAuthenticated;
    const isAdmin = authStore.isAdmin;

    // 1. Cek apakah rute tujuan membutuhkan login
    if (to.meta.requiresAuth && !isAuthenticated) {
        // Jika butuh login TAPI user belum login, lempar ke halaman login
        next({ name: 'Login' });
    }
    // 2. Cek apakah rute tujuan membutuhkan peran admin
    else if (to.meta.requiresAdmin && !isAdmin) {
        // Jika butuh admin TAPI user login bukan admin, lempar ke dashboard user
        next({ name: 'UserDashboard' });
    }
    // 3. Cek jika user yang sudah login mencoba mengakses halaman login
    else if (to.name === 'Login' && isAuthenticated) {
        // Arahkan ke dashboard yang sesuai dengan perannya
        next(isAdmin ? { name: 'AdminDashboard' } : { name: 'UserDashboard' });
    }
    // 4. Jika semua kondisi di atas tidak terpenuhi, izinkan navigasi
    else {
        next();
    }
});

export default router;