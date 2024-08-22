<script setup lang="ts">
import router from '@/router'
import { onMounted, ref } from 'vue'
import { RouterLink } from 'vue-router'

const dataUsers = ref<
    {
        id: number
        username: string
        email: string
        roles: string[]
    }[]
>([])

const dataRoles = ref<
    {
        id: number
        name: string
    }[]
>([])

const getDataUser = () => {
    fetch('http://localhost:8008/users', {
        credentials: 'include'
    })
        .then((response) => response.json())
        .then((json) => {
            if (json.message) {
                router.push('/auth')
                return
            }
            dataUsers.value = json
        })
}

// ---------------------

const getDataRole = () => {
    fetch('http://localhost:8008/roles', {
        credentials: 'include'
    })
        .then((response) => response.json())
        .then((json) => {
            if (json.message) {
                router.push('/auth')
                return
            }
            dataRoles.value = json
        })
}

onMounted(() => {
    getDataUser()
    getDataRole()
})
</script>
<template>
    <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
        <div class="p-4 flex justify-between">
            <button
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            >
                Add new user
            </button>
            <div>
            </div>
        </div>
        <table class="w-full text-sm text-left rtl:text-right text-gray-500">
            <thead class="text-xs text-gray-700 uppercase bg-gray-50">
                <tr>
                    <th scope="col" class="px-6 py-3">#</th>
                    <th scope="col" class="px-6 py-3">Id</th>
                    <th scope="col" class="px-6 py-3">User_Name</th>
                    <th scope="col" class="px-6 py-3">Roles</th>
                    <th scope="col" class="px-6 py-3">Action</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in dataUsers" :key="item.id">
                    <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                        <input type="checkbox" />
                    </th>
                    <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                        {{ item.id }}
                    </th>
                    <td class="px-6 py-4">
                        <span
                            class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800"
                        >
                            {{ item.username }}
                        </span>
                    </td>
                    <td class="px-6 py-4">
                        <span
                            v-for="role in item.roles"
                            :key="role"
                            class="inline-flex mr-1 items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800"
                        >
                            {{ role }}
                        </span>
                    </td>
                    <td class="px-6 py-4">
                        <RouterLink
                            :to="`/users/${item.id}`"
                            class="font-medium mr-2 text-blue-600 dark:text-blue-500 hover:underline"
                            >Edit</RouterLink
                        >
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>
