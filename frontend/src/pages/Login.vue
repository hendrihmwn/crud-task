<script setup>
import axiosClient from '../api';
import GuestLayout from '../components/GuestLayout.vue';
import { ref } from 'vue';
import { useRouter } from "vue-router";

const router = useRouter();

const username = ref('');
const password = ref('');
const errorMessage = ref('');

function submit() {
    axiosClient.post('/login', {
        username: username.value,
        password: password.value,
    }).then(response => {
        localStorage.setItem('token', response.data.data.token);
        window.location.href = '/';
    }).catch(error => {
        console.log(error.response);
        errorMessage.value = error.response.data.error;
    });
}

</script>

<template>
    <GuestLayout>
        <h2 class="mt-10 text-center text-2xl/9 font-bold tracking-tight text-gray-900">Sign in to your account</h2>
        <div v-if="errorMessage" class="mt-4 py-2 px-3 rounded text-white bg-red-400">
            {{errorMessage}}
        </div>
        <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
            <form @submit.prevent="submit" class="space-y-6" action="#" method="POST">
                <div>
                <label for="username" class="block text-sm/6 font-medium text-gray-900">Username</label>
                <div class="mt-2">
                    <input type="text" v-model="username" name="username" id="username" autocomplete="" required="" class="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6" />
                </div>
                </div>

                <div>
                <div class="flex items-center justify-between">
                    <label for="password" class="block text-sm/6 font-medium text-gray-900">Password</label>
                    <div class="text-sm">
                    </div>
                </div>
                <div class="mt-2">
                    <input type="password" v-model="password" name="password" id="password" autocomplete="current-password" required="" class="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6" />
                </div>
                </div>

                <div>
                <button type="submit" class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm/6 font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Sign in</button>
                </div>
            </form>
        </div>
    </GuestLayout>
</template>

<style scoped></style>