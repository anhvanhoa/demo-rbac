<script setup lang="ts">
import router from '@/router'
import { computed } from 'vue'
import { ref } from 'vue'
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const modal = ref(false)

const user = ref<{
    id: number
    username: string
    roles: {
        id: number
        name: string
    }[]
    email: string
}>({
    id: 0,
    username: '',
    roles: [],
    email: ''
})

const rules = ref<{
    id: number
        name: string
        path: string
        method: string
        status: boolean
        auth_type: string
        service: string
}[]>([])

const roleIds = ref<number[]>([])

const userRoles = computed(() => {
    return user.value.roles.map((role) => role.id)
})

const dataRoles = ref<
    {
        id: number
        name: string
    }[]
>([])

const setModal = (type: boolean) => {
    modal.value = type
}

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

const getUser = () => {
    fetch(`http://localhost:8008/users/${route.params.id}`, {
        credentials: 'include'
    })
        .then((response) => response.json())
        .then((json) => {
            user.value = json.user
            rules.value = json.rules
        })
}

const handleAddRole = () => {
    fetch(`http://localhost:8008/users/add-roles`, {
        credentials: 'include',
        method: 'POST',
        body: JSON.stringify({
            userId: user.value.id,
            roleIds: roleIds.value
        })
    })
        .then((response) => response.json())
        .then((json) => {
            roleIds.value = []
            setModal(false)
            if (json.message) {
                alert(json.message)
            }
            getUser()
        })
}

const handleDeleteRole = (id: number) => {
    if (confirm('Bạn có chắc chắn muốn xóa role này không?')) {
        fetch(`http://localhost:8008/users/delete-roles`, {
            credentials: 'include',
            method: 'POST',
            body: JSON.stringify({
                userId: user.value.id,
                roleIds: [id]
            })
        }).then((response) => {
            if (response.ok) {
                return getUser()
            }
            response.json().then((json) => {
                alert(json.message)
            })
        })
    }
}

onMounted(() => {
    getUser()
    getDataRole()
})
</script>
<template>
    <div v-if="modal" class="fixed z-50 inset-0">
        <div class="bg-black/70 absolute inset-0" @click="setModal(false)"></div>
        <div class="absolute top-1/2 -translate-y-1/2 z-10 flex justify-center w-full">
            <div class="bg-white shadow-lg rounded-lg min-w-96">
                <div class="p-6">
                    <div class="">
                        <form class="flex-1 flex items-start justify-between gap-8">
                            <div
                                v-for="item in dataRoles"
                                :key="item.id"
                                class="flex items-center gap-1"
                            >
                                <label
                                    :for="item.name"
                                    class="block mb-2 font-medium text-gray-900 cursor-pointer"
                                    :class="{
                                        'line-through cursor-default': userRoles.includes(item.id)
                                    }"
                                >
                                    {{ item.name }}
                                </label>
                                <input
                                    :id="item.name"
                                    :value="item.id"
                                    type="checkbox"
                                    :disabled="userRoles.includes(item.id)"
                                    v-model="roleIds"
                                    class="mb-1"
                                />
                            </div>
                        </form>
                        <button
                            @click="handleAddRole"
                            :class="{
                                'opacity-80': roleIds.length === 0,
                                'pointer-events-none': roleIds.length === 0
                            }"
                            class="justify-center w-full ppercase mt-4 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-700"
                        >
                            Add roles
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div>
        <div class="px-4 sm:px-0 mt-8">
            <h3 class="text-2xl font-semibold leading-7 text-gray-900">Chi tiết User</h3>
        </div>
        <div class="mt-6 border-t border-gray-100">
            <dl class="divide-y divide-gray-100">
                <div class="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
                    <dt class="font-medium leading-6 text-gray-900">Username</dt>
                    <dd class="mt-1 uppercase leading-6 text-gray-700 sm:col-span-2 sm:mt-0">
                        {{ user.username }}
                    </dd>
                </div>
                <div class="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
                    <dt class="font-medium leading-6 text-gray-900">Email address</dt>
                    <dd class="mt-1 uppercase leading-6 text-gray-700 sm:col-span-2 sm:mt-0">
                        {{ user.email }}
                    </dd>
                </div>
                <div class="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
                    <dt class="font-medium leading-6 text-gray-900">Roles</dt>
                    <dd
                        class="mt-1 leading-6 text-gray-700 sm:col-span-2 sm:mt-0 flex items-center"
                    >
                        <span
                            v-for="role in user.roles"
                            :key="role.id"
                            class="bg-blue-100 text-blue-700 px-3 py-2 rounded-md mr-4 flex items-center relative"
                        >
                            {{ role.name }}
                            <span
                                v-if="role.name !== 'ADMIN'"
                                @click="handleDeleteRole(role.id)"
                                class="absolute -top-2 -right-2 cursor-pointer"
                            >
                                <svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    width="18"
                                    height="18"
                                    viewBox="0 0 32 32"
                                >
                                    <path
                                        fill="currentColor"
                                        d="M16 2C8.2 2 2 8.2 2 16s6.2 14 14 14s14-6.2 14-14S23.8 2 16 2m5.4 21L16 17.6L10.6 23L9 21.4l5.4-5.4L9 10.6L10.6 9l5.4 5.4L21.4 9l1.6 1.6l-5.4 5.4l5.4 5.4z"
                                    />
                                </svg>
                            </span>
                        </span>
                        <button
                            @click="setModal(true)"
                            class="inline-flex items-center px-4 py-2 border border-transparent font-medium rounded-md shadow-sm text-white bg-sky-600 hover:bg-sky-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                        >
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                width="18"
                                height="18"
                                viewBox="0 0 24 24"
                            >
                                <path fill="currentColor" d="M11 13H5v-2h6V5h2v6h6v2h-6v6h-2z" />
                            </svg>
                        </button>
                    </dd>
                </div>
            </dl>
        </div>
        <div class="text-xl mt-6">Danh sách các path được truy cập</div>
        <table class="w-full text-sm text-left rtl:text-right text-gray-500 mt-4">
            <thead class="text-xs text-gray-700 uppercase bg-gray-50">
                <tr>
                    <th scope="col" class="px-6 py-3">#</th>
                    <th scope="col" class="px-6 py-3">Mô tả</th>
                    <th scope="col" class="px-6 py-3">Path</th>
                    <th scope="col" class="px-6 py-3">Method</th>
                    <th scope="col" class="px-6 py-3">Type Auth</th>
                    <th scope="col" class="px-6 py-3">Service</th>
                    <th scope="col" class="px-6 py-3">Status</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in rules" :key="item.id">
                    <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                        {{ item.id }}
                    </th>
                    <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                        {{ item.name }}
                    </th>
                    <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
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
                </tr>
            </tbody>
        </table>
    </div>
</template>
