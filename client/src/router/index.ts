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
                    path: '/about',
                    name: 'about',
                    // route level code-splitting
                    // this generates a separate chunk (About.[hash].js) for this route
                    // which is lazy-loaded when the route is visited.
                    component: () => import('../views/AboutView.vue')
                }
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
