import { createRouter, createWebHistory } from 'vue-router';
import DefaultLayout from './components/DefaultLayout.vue';
import Tasks from './pages/task/Tasks.vue';
import Login from './pages/Login.vue';
import NotFound from './pages/NotFound.vue';
import TasksCreate from './pages/task/TasksCreate.vue';
import TasksEdit from './pages/task/TasksEdit.vue';

const routes = [
    {
        path: '/',
        component: DefaultLayout,
        redirect: '/tasks',
        children: [
            {
                path: '/tasks',
                name: 'Tasks',
                component: Tasks,
                meta: { requiresAuth: true, title: 'Tasks' }
            },
            {
                path: '/task/create',
                name: 'TasksCreate',
                component: TasksCreate,
                meta: { requiresAuth: true, title: 'Create Task' }
            },
            {
                path: '/task/:id/edit',
                name: 'TasksEdit',
                component: TasksEdit,
                meta: { requiresAuth: true, title: 'Edit Task' }
            },
        ]
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
        meta: { guestOnly: true }
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: NotFound,
    }
];

const router = createRouter({
   history: createWebHistory(),
   routes 
});

function isLoggedIn() {
  return !!localStorage.getItem("token");
}

router.beforeEach((to) => {
  if (to.meta.requiresAuth && !isLoggedIn()) {
    return { name: "Login" }
  }

  if (to.meta.guestOnly && isLoggedIn()) {
    return { name: "Tasks" }
  }
})
 
 export default router;