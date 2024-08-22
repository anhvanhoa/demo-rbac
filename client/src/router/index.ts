import { createRouter, createWebHistory } from 'vue-router'
import LayoutMain from '@/layouts/layout-main/LayoutMain.vue'
import AuthView from '@/views/AuthView.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            component: LayoutMain,
            children: [
                {
                    path: '',
                    name: 'home',
                    component: () => import('@/views/HomeView.vue')
                },
                {
                    path: '/roles',
                    name: 'roles',
                    component: () => import('@/views/RoleView.vue')
                },
                {
                    path: '/users',
                    name: 'users',
                    component: () => import('@/views/UserView.vue')
                },
                {
                    path: '/users/:id',
                    name: 'user',
                    component: () => import('@/views/DetailUser.vue')
                },
                {
                    path: '/roles/add',
                    name: 'role-add',
                    component: () => import('@/views/FormRole.vue')
                },
            ]
        },
        {
            path: '/auth',
            name: 'auth',
            component: AuthView
        }
    ]
})

export default router
