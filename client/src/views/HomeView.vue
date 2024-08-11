<script setup lang="ts">
import router from '@/router'
import { ref } from 'vue'
import { onMounted } from 'vue'

const roles = ['User', 'Teacher', 'Admin']

const data = ref<
    {
        id: number
        name: string
        path: string
        method: string
        roles: number[]
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

const handleRefresh = () => {
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
                        @click="handleRefresh"
                        class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                    >
                        Refresh client
                    </button>
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
                                class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-red-100 text-red-800"
                            >
                                {{ roles[role - 1] }}</span
                            >
                        </td>
                        <td class="px-6 py-4">
                            <span
                                class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-orange-100 text-orange-800"
                                >{{ item.service }}</span
                            >
                        </td>
                        <td class="px-6 py-4">
                            {{ item.status ? 'Auth' : 'Not auh' }}
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
