<script setup lang="ts">
import router from '@/router'
import { ref } from 'vue'
import { onMounted } from 'vue'
import { RouterLink } from 'vue-router'

const notify = ref('')
const dataRoles = ref<
    {
        id: number
        name: string
    }[]
>([])

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
const handleRefreshRole = () => {
    fetch('http://localhost:8000/refresh-roles', {
        credentials: 'include',
        method: 'POST'
    })
        .then((response) => response.json())
        .then((json) => {
            getDataRole()
            if (json.message) {
                alert(json.message)
            }
        })
}
const handleRefreshRoleAdmin = () => {
    fetch('http://localhost:8008/refresh-roles', {
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

// ---------------------

const handleDeleteRole = (id: number) => {
    if (confirm('Bạn có chắc chắn muốn xóa role này không?')) {
        fetch(`http://localhost:8008/roles/delete/${id}`, {
            credentials: 'include',
            method: 'DELETE'
        })
            .then((response) => {
                return response.json()
                // const reader = response.body?.getReader()
                // const decoder = new TextDecoder('utf-8')
                // if (!reader) return
                // const loop = true
                // while (loop) {
                //     const { done, value } = await reader.read()
                //     if (done) break
                //     const chunk = JSON.parse(decoder.decode(value, { stream: true, }))
                //     if (chunk.message) {
                //         notify.value = chunk.message
                //     }
                // }
            })
            .then((json) => {
                if (json.message) {
                    notify.value = json.message
                    setTimeout(() => {
                        notify.value = ''
                    }, 3000)
                }
                getDataRole()
            })
    }
}
onMounted(() => {
    getDataRole()
})
</script>
<template>
    <div
        v-if="notify"
        class="absolute bottom-8 left-1/2 -translate-x-1/2 bg-slate-100 py-2 px-4 rounded-lg"
    >
        {{ notify }}
    </div>
    <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
        <div class="p-4 flex justify-between">
            <RouterLink
                to="/roles/add"
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            >
                Add new role
            </RouterLink>
            <div>
                <button
                    @click="handleRefreshRole"
                    class="inline-flex mx-2 items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                >
                    Refresh client
                </button>
                <button
                    @click="handleRefreshRoleAdmin"
                    class="inline-flex mx-2 items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-sky-600 hover:bg-sky-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-sky-500"
                >
                    Refresh admin
                </button>
            </div>
        </div>
        <table class="w-full text-sm text-left rtl:text-right text-gray-500">
            <thead class="text-xs text-gray-700 uppercase bg-gray-50">
                <tr>
                    <th scope="col" class="px-6 py-3">#</th>
                    <th scope="col" class="px-6 py-3">Id</th>
                    <th scope="col" class="px-6 py-3">Name</th>
                    <th scope="col" class="px-6 py-3">Action</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in dataRoles" :key="item.id">
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
                            {{ item.name }}
                        </span>
                    </td>
                    <td class="px-6 py-4" v-if="item.name !== 'ADMIN'">
                        <button
                            @click="handleDeleteRole(item.id)"
                            class="font-medium mr-2 text-red-600 dark:text-red-500 hover:underline"
                        >
                            Delete
                        </button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>
