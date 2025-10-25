<script setup>
import { ref } from 'vue'
import { RouterLink } from 'vue-router';
import axiosClient from "../../axios";
import router from '../../router';

const errorMessage = ref('');
const data = ref({
  title: '',
  description: '',
  status: 'backlog',
})

function submit() {
  axiosClient.post("/tasks", data.value)
        .then(response => {
          router.push({name: 'Tasks'});
        })
        .catch(error => {
          console.log(error.response)
          errorMessage.value = error.response.data.error;
        });
}

const status = [
  {
    code: "backlog",
    name: 'Backlog',
  },
  {
    code: "in-progress",
    name: 'In Progress',
  },
  {
    code: "completed",
    name: 'Completed',
  },
]
const selected = ref(status[0])

</script>

<template>
    <header class="relative bg-white shadow-sm">
      <div class="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
        <div>
            <nav aria-label="breadcrumb" class="w-max">
                <ol class="flex w-full flex-wrap items-center rounded-md bg-slate-50 px-4 py-2">
                    <li class="flex cursor-pointer items-center text-sm text-slate-500 transition-colors duration-300 hover:text-slate-800">
                    <RouterLink :to="{ name: 'Tasks' }">Tasks</RouterLink>
                    <span class="pointer-events-none mx-2 text-slate-800">
                        /
                    </span>
                    </li>
                    <li class="flex cursor-pointer items-center text-sm text-slate-500 transition-colors duration-300 hover:text-slate-800">
                        <a href="#">Create</a>
                    </li>
                </ol>
            </nav>
        </div>
      </div>
    </header>
    <main>
      <div class="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
        <div v-if="errorMessage" class="mt-4 py-2 px-3 rounded text-white bg-red-400">
            {{errorMessage}}
        </div>
        <div class="mt-5 sm:mx-auto bg-white dark:bg-gray-800 shadow-lg rounded-xl p-6 w-full mx-auto">
            <form @submit.prevent="submit" class="space-y-6" action="#" method="POST">
                <div>
                    <label for="title" class="block mb-2 text-sm text-slate-600">Title</label>
                    <div class="mt-2">
                        <input v-model="data.title" type="text" name="title" id="title" autocomplete="" required="" class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow" />
                    </div>
                </div>

                <div>
                    <label for="description" class="block mb-2 text-sm text-slate-600">Description</label>
                    <div class="mt-2">
                        <textarea v-model="data.description" name="description" id="description" autocomplete="" required="" class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow" />
                    </div>
                </div>

                <div>
                    <label for="status" class="block mb-2 text-sm text-slate-600">Status</label>
                    <div class="mt-2 relative">
                        <select v-model="data.status" id="status" name="status" class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded pl-3 pr-8 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-400 shadow-sm focus:shadow-md appearance-none cursor-pointer">
                            <option
                                v-for="s in status"
                                :key="s.code"
                                :value="s.code"
                            >
                                {{ s.name }}
                            </option>
                        </select>
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.2" stroke="currentColor" class="h-5 w-5 ml-1 absolute top-2.5 right-2.5 text-slate-700">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M8.25 15 12 18.75 15.75 15m-7.5-6L12 5.25 15.75 9" />
                        </svg>
                    </div>
                </div>

                <div class="flex items-center gap-3">
                    <RouterLink
                        :to="{ name: 'Tasks' }"
                        class="flex select-none items-center gap-2 rounded bg-neutral-400 py-2.5 px-4 text-xs font-semibold text-white shadow-md shadow-neutral-500/10 transition-all hover:shadow-lg hover:shadow-neutral-500/20 focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none disabled:pointer-events-none disabled:opacity-50 disabled:shadow-none"
                    >
                        Back
                    </RouterLink>

                    <button
                        type="submit"
                        class="flex select-none items-center gap-2 rounded bg-blue-400 py-2.5 px-4 text-xs font-semibold text-white shadow-md shadow-blue-500/10 transition-all hover:shadow-lg hover:shadow-blue-500/20 focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none disabled:pointer-events-none disabled:opacity-50 disabled:shadow-none"
                    >
                        Create
                    </button>
                </div>
            </form>
        </div>
      </div>
    </main>

  
</template>

<style scoped></style>