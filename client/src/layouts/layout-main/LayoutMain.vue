<script setup lang="ts">
import { RouterView } from 'vue-router';
import SidebarMain from './SidebarMain.vue'
</script>

<template>
    <SidebarMain />
    <div class="px-8 py-4 sm:ml-64">
        <RouterView />
    </div>
</template>
