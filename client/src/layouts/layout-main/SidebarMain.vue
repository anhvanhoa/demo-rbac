<script setup lang="ts">
import { RouterLink } from 'vue-router'

const handleLogout = () => {
    fetch('http://localhost:8008/auth/logout', {
        credentials: 'include'
    })
        .then((response) => response.json())
        .then((json) => {
            if (json.message) {
                alert(json.message)
                window.location.href = '/auth'
                return
            }
            window.location.href = '/auth'
        })
}
</script>

<template>
    <aside
        id="sidebar-multi-level-sidebar"
        class="fixed top-0 left-0 z-40 w-64 h-screen transition-transform -translate-x-full sm:translate-x-0"
        aria-label="Sidebar"
    >
        <div class="h-full px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800">
            <div>
                <RouterLink to="/">
                    <div class="flex items-center mb-4 gap-2">
                        <img
                            class="size-9 rounded-md object-cover"
                            src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAOEAAADhCAMAAAAJbSJIAAAA1VBMVEX///8BsPEAsPAAsfoIq+IEsfZivO0AsvEArfEArfUDsOdZw+Sd3/oAq/Lt+vzg9vvC6v1mw/TZ7vMEtPAAp/L3/Pwst/225fBkzP3N7v+a3+4Ap+MiuN4wv+7s+f+h1/bR8fkjufAhsf0Aou4Ao+ng+fhey+Vt0uvu/fnr9P0ttPMAr9sFqNcItOUHuvwYnu8AuuUVnPADwPZmw+s7ruw0qeJKvOLJ7/YjxOolweAMqP7I5POP2PSw4PhVyvMFoepu0vSm2upcs+pewPyFy/YAvflhvuu4TYIcAAAEgElEQVR4nO3df1PaSBzHcbNJupscAWrWUhtQSDjtiRs8aOUOi9j70ef/kKpzk7DT3mZBYrpLP6//HL9i3ixKJqx4dAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAANvr14pVRLt/XXTjkrlnadRf+xjyjXJ3UXXg88U3i9mtfw2NOjOLUXxg5RiEo3LOQ7HAolbM73NA3X/iChcSjjh+qEZZcx8VwyFKXqGd9V2qkccQrbjgOfEobKXTyWPQDNTdN32yOOqciTFWjaShGuU+K5XR96qlv1594cd5IIZtOeaZ8ngoTIiZeUhxHNs0yqn5Wc3iWZcUqJtf9G0+9hpRPl7+zJgrpTMwGvvpkI4+lu3omhBipZ8UgFLNy2CcVN+v6YhSOaAOFhGXzTrujdnbWZcVdTYIPlbPt9vBDedclt8dnx1XDw4+8kUcpDfqa2VNe/vrw3mjOrtqpVz78R5qjPglC2wvPqmdRuJ9dCu8OvvDw13CnQrcsdFCIwv2gUIZCGQplKNwPCmUolKFQhsL9oFCGQhkKZSjcDwplKJSZWBjzFyok2sK0mUIS32pmF5tZ/oemsCPKF1sCobuqP2JNvDKzReFpWs5mf+rWUPDtCwdpM4XpR83sqVMcCGHaNbxKi4c00c2eXC0bKfSCda/bq3D3KS5f5U7C+bvVO6XW5f26/Jklk/Oq2+2uLrMsbqIwnK1v1lwpisZxUN7VsR95U4cpkCVj0aCYjVigvlnOl+TGF2EThY+Ys4vq/RZE+cF35G+L/TQoROGPh0IU/pyFW24jJN9/pPzKbz/Ltnqm/W8Lbe2FnceTla0wtg6D4uTKn7hJ7gXq6TDPRHHgQSTYeJvv8HSf5LUXHq1aW1qtHjbrfZ1/bi26SnefH9jmRKz/dKa6pV7tgbtY+OUjynutm/U2j9KHJg6uFous/HGir3Sz126xvzS6r/0vDF6KdBXDmWtmO0mx4JTr1tscUmEy0MwOefFbibBfGjm6OkiFRFs4LgsDFJoDhTIUmgmFMhSaCYUyFJoJhTIUmgmFMhSaCYUyFJpppytRthfG95rZofOv3YXrvzSznWlid2H+VjO7cC1fw+xvzWxnya0uZOGo96niVavVosUEs6+wG5WbsxzlfqhCOWht4dZQaBIUotB8KDy8wop3zpP2C1laGDLuO1RpHLOMWl1I3HwkwkB1PuMnHnHtXkPKGeU8UfKXfuxZXeiM+hcX92+V5v8MhN2PUudpF1DlTqdeWO6JsrNQe51m5QR2F2qvtZ0TFJpol8LWwReeb97N8mALycEXHv4aotBEKJSh0EwolKHQTCiUodBMKJSh0EwolKHQTCiUWV+ov+Yd2H0linDdXzp33dDqQie7bVfpvD/PbNybKBXmo1kolDu+Iu8m9m3cmygVsmC59GOlx89mqd1X9SfCoVT9f7hihwobd0HLezGengyq9mJs2Fq4PRSa5CcoHD+n0ImsKiTqfVBKFr23yWPhMwLtKqS6/ev/e4ZjT+Hpl/nr57j80QcOAAAAAAAAAAAAAAAAAAfuK5Df2CxAQwH7AAAAAElFTkSuQmCC"
                            alt=""
                        />
                        <p class="text-lg font-semibold text-white">Techmaster</p>
                    </div>
                </RouterLink>
            </div>
            <ul class="space-y-2 font-medium">
                <li>
                    <RouterLink
                        to="/"
                        class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
                    >
                        <svg
                            class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                            aria-hidden="true"
                            xmlns="http://www.w3.org/2000/svg"
                            fill="currentColor"
                            viewBox="0 0 22 21"
                        >
                            <path
                                d="M16.975 11H10V4.025a1 1 0 0 0-1.066-.998 8.5 8.5 0 1 0 9.039 9.039.999.999 0 0 0-1-1.066h.002Z"
                            />
                            <path
                                d="M12.5 0c-.157 0-.311.01-.565.027A1 1 0 0 0 11 1.02V10h8.975a1 1 0 0 0 1-.935c.013-.188.028-.374.028-.565A8.51 8.51 0 0 0 12.5 0Z"
                            />
                        </svg>
                        <span class="ms-3">Dashboard</span>
                    </RouterLink>
                </li>
                <li>
                    <RouterLink
                        to="/roles"
                        class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
                    >
                        <svg
                            class="flex-shrink-0 w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                            aria-hidden="true"
                            xmlns="http://www.w3.org/2000/svg"
                            fill="currentColor"
                            viewBox="0 0 20 18"
                        >
                            <path
                                d="M14 2a3.963 3.963 0 0 0-1.4.267 6.439 6.439 0 0 1-1.331 6.638A4 4 0 1 0 14 2Zm1 9h-1.264A6.957 6.957 0 0 1 15 15v2a2.97 2.97 0 0 1-.184 1H19a1 1 0 0 0 1-1v-1a5.006 5.006 0 0 0-5-5ZM6.5 9a4.5 4.5 0 1 0 0-9 4.5 4.5 0 0 0 0 9ZM8 10H5a5.006 5.006 0 0 0-5 5v2a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-2a5.006 5.006 0 0 0-5-5Z"
                            />
                        </svg>
                        <span class="flex-1 ms-3 whitespace-nowrap">Roles</span>
                    </RouterLink>
                </li>
                <li>
                    <RouterLink
                        to="/users"
                        class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
                    >
                        <svg
                            class="flex-shrink-0 w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                            xmlns="http://www.w3.org/2000/svg"
                            width="18"
                            height="18"
                            viewBox="0 0 36 36"
                        >
                            <path
                                fill="currentColor"
                                d="M30.61 24.52a17.16 17.16 0 0 0-25.22 0a1.5 1.5 0 0 0-.39 1v6A1.5 1.5 0 0 0 6.5 33h23a1.5 1.5 0 0 0 1.5-1.5v-6a1.5 1.5 0 0 0-.39-.98"
                                class="clr-i-solid clr-i-solid-path-1"
                            />
                            <circle
                                cx="18"
                                cy="10"
                                r="7"
                                fill="currentColor"
                                class="clr-i-solid clr-i-solid-path-2"
                            />
                            <path fill="none" d="M0 0h36v36H0z" />
                        </svg>
                        <span class="flex-1 ms-3 whitespace-nowrap">Users</span>
                    </RouterLink>
                </li>
                <li>
                    <a
                        href="#"
                        class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
                    >
                        <svg
                            class="flex-shrink-0 w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                            aria-hidden="true"
                            xmlns="http://www.w3.org/2000/svg"
                            fill="currentColor"
                            viewBox="0 0 18 20"
                        >
                            <path
                                d="M17 5.923A1 1 0 0 0 16 5h-3V4a4 4 0 1 0-8 0v1H2a1 1 0 0 0-1 .923L.086 17.846A2 2 0 0 0 2.08 20h13.84a2 2 0 0 0 1.994-2.153L17 5.923ZM7 9a1 1 0 0 1-2 0V7h2v2Zm0-5a2 2 0 1 1 4 0v1H7V4Zm6 5a1 1 0 1 1-2 0V7h2v2Z"
                            />
                        </svg>
                        <span class="flex-1 ms-3 whitespace-nowrap">Products</span>
                    </a>
                </li>
                <li>
                    <button
                        @click="handleLogout"
                        class="w-full flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
                    >
                        <svg
                            class="flex-shrink-0 w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                            aria-hidden="true"
                            xmlns="http://www.w3.org/2000/svg"
                            fill="none"
                            viewBox="0 0 18 16"
                        >
                            <path
                                stroke="currentColor"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="2"
                                d="M1 8h11m0 0L8 4m4 4-4 4m4-11h3a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-3"
                            />
                        </svg>
                        <span class="ms-3 whitespace-nowrap">Sign Out</span>
                    </button>
                </li>
            </ul>
        </div>
    </aside>
</template>
