import { createRouter, createWebHistory } from "vue-router";

import Register from "@/components/Register.vue";
import Login from "@/components/Login.vue"

const routes = [
    { path: "/login", component: Login },
    { path: "/register", component: Register },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;