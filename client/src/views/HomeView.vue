<script setup lang="ts">
import router from '@/router'
import { ref } from 'vue'
import { onMounted } from 'vue'

const data = ref<
    {
        id: number
        name: string
        path: string
        method: string
        roles: string[]
        status: boolean
        auth_type: string
        service: string
    }[]
>([])
const getData = () => {
    fetch('http://localhost:8008/rules', {
        credentials: 'include'
    })
        .then((response) => response.json())
        .then((json) => {
            if (json.message) {
                router.push('/auth')
                return
            }
            data.value = json
        })
}

const handleRefreshRule = () => {
    fetch('http://localhost:8000/refresh-rules', {
        credentials: 'include',
        method: 'POST'
    })
        .then((response) => response.json())
        .then((json) => {
            getData()
            if (json.message) {
                alert(json.message)
            }
        })
}
const handleRefreshRuleAdmin = () => {
    fetch('http://localhost:8008/refresh-rules', {
        credentials: 'include',
        method: 'POST'
    })
        .then((response) => response.json())
        .then((json) => {
            if (json.message) {
                alert(json.message)
            }
        })
}

onMounted(() => {
    getData()
})
</script>

<template>
    <main>
        <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
            <div class="p-4 flex justify-between">
                <button
                    class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                >
                    Add new rule
                </button>
                <div>
                    <button
                        @click="handleRefreshRule"
                        class="mx-2 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                    >
                        Refresh client
                    </button>
                    <button
                        @click="handleRefreshRuleAdmin"
                        class="mx-2 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-sky-600 hover:bg-sky-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-sky-500"
                    >
                        Refresh admin
                    </button>
                </div>
            </div>
            <div class="px-4 mb-2">
                <div class="grid grid-cols-4">
                    <div>
                        <p class="font-semibold">Method</p>
                        <div class="mt-1 grid grid-cols-5">
                            <div class="flex items-center gap-x-1 text-sm">
                                <label for="GET">GET</label>
                                <input id="GET" type="checkbox" class="mt-1" />
                            </div>
                            <div class="flex items-center gap-x-1 text-sm">
                                <label for="POST">POST</label>
                                <input id="POST" type="checkbox" class="mt-1" />
                            </div>
                            <div class="flex items-center gap-x-1 text-sm">
                                <label for="PUT">PUT</label>
                                <input id="PUT" type="checkbox" class="mt-1" />
                            </div>
                            <div class="flex items-center gap-x-1 text-sm">
                                <label for="DELETE">DELETE</label>
                                <input id="DELETE" type="checkbox" class="mt-1" />
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <table class="w-full text-sm text-left rtl:text-right text-gray-500">
                <thead class="text-xs text-gray-700 uppercase bg-gray-50">
                    <tr>
                        <th scope="col" class="px-6 py-3">#</th>
                        <th scope="col" class="px-6 py-3">Mô tả</th>
                        <th scope="col" class="px-6 py-3">Path</th>
                        <th scope="col" class="px-6 py-3">Method</th>
                        <th scope="col" class="px-6 py-3">Type Auth</th>
                        <th scope="col" class="px-6 py-3">Roles</th>
                        <th scope="col" class="px-6 py-3">Service</th>
                        <th scope="col" class="px-6 py-3">Status</th>
                        <th scope="col" class="px-6 py-3">Action</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="item in data" :key="item.id">
                        <th
                            scope="row"
                            class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
                        >
                            <input type="checkbox" />
                        </th>
                        <th
                            scope="row"
                            class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
                        >
                            {{ item.name }}
                        </th>
                        <th
                            scope="row"
                            class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
                        >
                            {{ item.path }}
                        </th>
                        <td class="px-6 py-4">
                            <span
                                class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800"
                            >
                                {{ item.method }}</span
                            >
                        </td>
                        <td class="px-6 py-4">
                            <p>
                                <span
                                    class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800"
                                    >{{ item.auth_type }}</span
                                >
                            </p>
                        </td>
                        <td class="px-6 py-4">
                            <span
                                v-for="role in item.roles"
                                :key="role"
                                class="inline-flex mx-1 items-center px-2.5 rounded-full text-xs font-medium bg-red-100 text-red-800"
                            >
                                {{ role }}</span
                            >
                        </td>
                        <td class="px-6 py-4">
                            <span
                                class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-orange-100 text-orange-800"
                                >{{ item.service }}</span
                            >
                        </td>
                        <td class="px-6 py-4">
                            <svg
                                v-if="item.status"
                                xmlns="http://www.w3.org/2000/svg"
                                width="18"
                                height="18"
                                viewBox="0 0 24 24"
                            >
                                <path
                                    fill="currentColor"
                                    d="m9.55 18l-5.7-5.7l1.425-1.425L9.55 15.15l9.175-9.175L20.15 7.4z"
                                />
                            </svg>
                            <svg
                                v-if="!item.status"
                                xmlns="http://www.w3.org/2000/svg"
                                width="18"
                                height="18"
                                viewBox="0 0 24 24"
                            >
                                <path
                                    fill="currentColor"
                                    d="M6.4 19L5 17.6l5.6-5.6L5 6.4L6.4 5l5.6 5.6L17.6 5L19 6.4L13.4 12l5.6 5.6l-1.4 1.4l-5.6-5.6z"
                                />
                            </svg>
                        </td>
                        <td class="px-6 py-4">
                            <a
                                href="#"
                                class="font-medium text-blue-600 dark:text-blue-500 hover:underline"
                                >Edit</a
                            >
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </main>
</template>
