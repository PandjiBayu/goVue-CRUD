import { createRouter, createWebHistory } from "vue-router"; 

const routes= [
        {
            path: "/",
            alias: "/movies",
            name: "Movies",
            component: () => import('../components/Movies.vue'),
        },
        {   
            path: "/movie",
            name: "Create",
            component: () => import('../components/MovieCreate.vue'),
        },
        {   
            path: "/movie/:id",
            name: "Update",
            component: () => import('../components/MovieUpdate.vue'),
        },
    ]
const router = createRouter({
        history: createWebHistory(),
        routes
      })

export default router;
