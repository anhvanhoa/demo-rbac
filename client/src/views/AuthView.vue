<script setup lang="ts">
import router from '@/router';
import { reactive } from 'vue'
const form = reactive({
    email: '',
    password: ''
})

const handleLogin = () => {
    fetch('http://localhost:8008/auth/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        credentials: 'include',
        body: JSON.stringify(form)
    })
        .then((response) => response.json())
        .then((json) => {
            if (json.message) {
                console.log('Login failed')
            } else {
                console.log('Login success')
                router.push('/')
            }
        })
}
</script>
<template>
    <div class="mt-8">
        <h1 class="text-center text-2xl font-semibold">Chào mừng đến với techmaster</h1>
        <div class="flex justify-center mt-4 max-w-sm mx-auto">
            <form class="w-full">
                <div class="mb-5">
                    <label for="email" class="block mb-2 text-sm font-medium text-gray-900"
                        >Your email</label
                    >
                    <input
                        v-model="form.email"
                        type="text"
                        id="email"
                        class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
                        placeholder="name@flowbite.com"
                        required
                    />
                </div>
                <div class="mb-5">
                    <label for="password" class="block mb-2 text-sm font-medium text-gray-900"
                        >Your password</label
                    >
                    <input
                        v-model="form.password"
                        type="password"
                        id="password"
                        class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
                        required
                    />
                </div>
                <button
                    @click.prevent="handleLogin"
                    type="submit"
                    class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                >
                    Đăng nhập
                </button>
            </form>
        </div>
    </div>
</template>
