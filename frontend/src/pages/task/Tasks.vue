<script setup>

import {onMounted, ref, computed} from "vue";
import axiosClient from "../../api";

const tasks = ref([])
const errorMessage = ref('');
let limit = 10;
let page = 1;
const meta = ref({
    page: page,
    limit: limit,
    total: 0
});

const totalPages = computed(() => {
    const t = Number(meta.value.total || 0);
    const l = Number(meta.value.limit || limit);
    return Math.max(1, Math.ceil(t / l));
});

function setPage(p) {
    const np = Math.max(1, Math.min(totalPages.value, Number(p || 1)));
    page = np;
    return fetchTasks();
}
function prevPage() { if (page > 1) return setPage(page - 1); }
function nextPage() { if (page < totalPages.value) return setPage(page + 1); }

const searchValue = ref('');
const statusValue = ref('');
const sortValue = ref('created_at');
const orderValue = ref(-1);

function buildUrl() {
    const params = new URLSearchParams();
    params.append('limit', String(limit));
    params.append('page', String(page));

    if (searchValue.value) params.append('search', searchValue.value);
    if (statusValue.value) params.append('status', statusValue.value);
    if (sortValue.value) params.append('sort_by', sortValue.value);
    if (orderValue.value) params.append('order', orderValue.value);

    return `/tasks?${params.toString()}`;
}

function fetchTasks() {
    return axiosClient.get(buildUrl()).then((response) => {
        tasks.value = response.data.data ?? [];
        const m = response.data.meta ?? {};
        meta.value = {
            page: Number(m.page ?? page),
            limit: Number(m.limit ?? limit),
            total: Number(m.total ?? 0)
        };
        page = meta.value.page;
    });
}

function doFilter() {
    page = 1;
    return fetchTasks();
}

function deleteTask(id) {
  if (!confirm("Are you sure you want to delete this task?")) {
    return;
  }

  axiosClient.delete(`/tasks/${id}`)
      .then(response => {
        fetchTasks();
      }).catch(error => {
        console.log(error.response)
        errorMessage.value = error.response.data.error;
    });
}

onMounted(() => {
  fetchTasks();
})

</script>

<template>
    <header class="relative bg-white shadow-sm">
      <div class="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
        <div>
            <nav aria-label="breadcrumb" class="w-max">
                <ol class="flex w-full flex-wrap items-center rounded-md bg-slate-50 px-4 py-2">
                    <li class="flex cursor-pointer items-center text-sm text-slate-500 transition-colors duration-300 hover:text-slate-800">
                        <RouterLink :to="{ name: 'Tasks' }">List of Tasks</RouterLink>
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
        <div class="w-full flex justify-between items-center mb-3 mt-1 pl-3">
            <div>
                <div class="flex flex-col gap-2 shrink-0 sm:flex-row">      
                    <div class="w-full max-w-sm min-w-[200px]">
                        <label class="block mb-1 text-sm text-slate-800">
                            Filter by Status
                        </label>
                        
                        <div class="relative">
                            <select
                                v-model="statusValue"
                                @change="doFilter"
                                class="w-full bg-white placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded pl-3 pr-8 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-400 shadow-sm focus:shadow-md appearance-none cursor-pointer">
                                <option value="">All</option>
                                <option value="backlog">Backlog</option>
                                <option value="in-progress">In Progress</option>
                                <option value="completed">Completed</option>
                            </select>
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.2" stroke="currentColor" class="h-5 w-5 ml-1 absolute top-2.5 right-2.5 text-slate-700">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M8.25 15 12 18.75 15.75 15m-7.5-6L12 5.25 15.75 9" />
                            </svg>
                        </div>
                    </div>
                    <div class="w-full max-w-sm min-w-[200px]">
                        <label class="block mb-1 text-sm text-slate-800">
                            Sort By
                        </label>
                        
                        <div class="relative">
                            <select
                                v-model="sortValue"
                                @change="doFilter"
                                class="w-full bg-white placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded pl-3 pr-8 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-400 shadow-sm focus:shadow-md appearance-none cursor-pointer">
                                <option value="">Default</option>
                                <option value="created_at">Created At</option>
                                <option value="title">Title</option>
                                <option value="status">Status</option>
                            </select>
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.2" stroke="currentColor" class="h-5 w-5 ml-1 absolute top-2.5 right-2.5 text-slate-700">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M8.25 15 12 18.75 15.75 15m-7.5-6L12 5.25 15.75 9" />
                            </svg>
                        </div>
                    </div>
                    <div class="w-full max-w-sm min-w-[200px]">
                        <label class="block mb-1 text-sm text-slate-800">
                            Order By
                        </label>
                        
                        <div class="relative">
                            <select
                                v-model="orderValue"
                                @change="doFilter"
                                class="w-full bg-white placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded pl-3 pr-8 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-400 shadow-sm focus:shadow-md appearance-none cursor-pointer">
                                <option value="1">Ascending</option>
                                <option value="-1">Descending</option>
                            </select>
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.2" stroke="currentColor" class="h-5 w-5 ml-1 absolute top-2.5 right-2.5 text-slate-700">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M8.25 15 12 18.75 15.75 15m-7.5-6L12 5.25 15.75 9" />
                            </svg>
                        </div>
                    </div>
                </div>
            </div>
            <div class="ml-3">
                <div class="flex flex-col gap-2 shrink-0 sm:flex-row">
                    <div class="relative">
                        <input
                            v-model="searchValue"
                            @keydown.enter.prevent="doFilter"
                            class="bg-white w-full pr-11 h-10 pl-3 py-2 bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded transition duration-200 ease focus:outline-none focus:border-slate-400 hover:border-slate-400 shadow-sm focus:shadow-md"
                            placeholder="Search for task..."
                        />
                        <button
                            class="absolute h-8 w-8 right-1 top-1 my-auto px-2 flex items-center bg-white rounded"
                            type="button"
                            @click="doFilter"
                        >
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="3" stroke="currentColor" class="w-8 h-8 text-slate-600">
                            <path stroke-linecap="round" stroke-linejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z" />
                        </svg>
                        </button>
                    </div>
                    
                    <RouterLink :to="{ name: 'TasksCreate' }">
                        <button
                    class="flex select-none items-center gap-2 rounded bg-blue-500 py-2.5 px-4 text-xs font-semibold text-white shadow-md shadow-blue-900/10 transition-all hover:shadow-lg hover:shadow-blue-900/20 focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none disabled:pointer-events-none disabled:opacity-50 disabled:shadow-none"
                    type="button">Create New</button></RouterLink>
                    
                </div>
                
            </div>
            
        </div>
        
        <div class="relative flex flex-col w-full h-full overflow-scroll text-gray-700 bg-white shadow-md rounded-lg bg-clip-border">
        <table class="w-full text-left table-auto min-w-max">
            <thead>
            <tr>
                <th class="p-4 border-b border-slate-200 bg-slate-50">
                <p class="text-sm font-normal leading-none text-slate-500">
                    ID
                </p>
                </th>
                <th class="p-4 border-b border-slate-200 bg-slate-50">
                <p class="text-sm font-normal leading-none text-slate-500">
                    Title
                </p>
                </th>
                <th class="p-4 border-b border-slate-200 bg-slate-50">
                <p class="text-sm font-normal leading-none text-slate-500">
                    Description
                </p>
                </th>
                <th class="p-4 border-b border-slate-200 bg-slate-50">
                <p class="text-sm font-normal leading-none text-slate-500">
                    Status
                </p>
                </th>
                <th class="p-4 border-b border-slate-200 bg-slate-50">
                <p class="text-sm font-normal leading-none text-slate-500">
                    Created At
                </p>
                </th>
                <th class="p-4 border-b border-slate-200 bg-slate-50">
                <p class="text-sm font-normal leading-none text-slate-500">
                    Action
                </p>
                </th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="task in tasks" :key="task.id" class="hover:bg-slate-50 border-b border-slate-200">
                <td class="p-4 py-5">
                <p class="block text-sm text-slate-800">{{ task.id }}</p>
                </td>
                <td class="p-4 py-5">
                <p class="block text-sm font-semibold text-slate-800">{{ task.title }}</p>
                </td>
                <td class="p-4 py-5">
                <p class="text-sm text-slate-500">{{ task.description }}</p>
                </td>
                <td class="p-4 py-5">
                    <div class="w-max">
                        <div>
                            <div
                                :class="[
                                    'relative grid items-center px-2 py-1 font-sans text-xs font-bold uppercase rounded-md select-none whitespace-nowrap',
                                    (task.status || '').toString().toLowerCase().includes('in-progress') ? 'text-yellow-900 bg-yellow-400/20' :
                                    (task.status || '').toString().toLowerCase().includes('completed') ? 'text-green-900 bg-green-500/20' :
                                    (task.status || '').toString().toLowerCase().includes('backlog') ? 'text-orange-900 bg-orange-400/20' :
                                    'text-slate-700 bg-slate-200/20'
                                ]">
                                <span>{{ task.status }}</span>
                            </div>
                        </div>
                    </div>
                </td>
                <td class="p-4 py-5">
                <p class="text-sm text-slate-500">
                    {{
                        (() => {
                            if (!task.created_at) return '';
                            const d = new Date(task.created_at);
                            const p = n => String(n).padStart(2, '0');
                            return `${d.getFullYear()}-${p(d.getMonth()+1)}-${p(d.getDate())} ${p(d.getHours())}:${p(d.getMinutes())}:${p(d.getSeconds())}`;
                        })()
                    }}
                </p>
                </td>
                <td class="p-4 py-5">
                    <div class="text-sm flex items-center">
                        <RouterLink
                            :to="{ name: 'TasksEdit', params: { id: task.id } }"
                            class="text-blue-600 hover:text-blue-800 mr-3 flex items-center"
                            title="Edit"
                            aria-label="Edit"
                        >
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="h-5 w-5">
                                <path stroke-linecap="round" stroke-linejoin="round" d="m16.862 4.487 1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L6.832 19.82a4.5 4.5 0 0 1-1.897 1.13l-2.685.8.8-2.685a4.5 4.5 0 0 1 1.13-1.897L16.863 4.487Zm0 0L19.5 7.125" />
                            </svg>

                        </RouterLink>

                        <button
                            @click="deleteTask(task.id)"
                            class="text-red-600 hover:text-red-800 flex items-center cursor-pointer"
                            title="Delete"
                            aria-label="Delete"
                        >
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
                                <path d="M3 6h18"></path>
                                <path d="M8 6v-.5A2.5 2.5 0 0 1 10.5 3h3A2.5 2.5 0 0 1 16 5.5V6"></path>
                                <path d="M19 6l-1 14a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2L5 6"></path>
                                <path d="M10 11v6"></path>
                                <path d="M14 11v6"></path>
                            </svg>
                        </button>
                    </div>
                </td>
            </tr>
            
            </tbody>
        </table>
        
        <div class="flex justify-between items-center px-4 py-3">
            <div class="text-sm text-slate-500">
            Showing <b>{{ (meta.page * 10) - 9 }}-{{ (meta.total < (meta.page * 10)) ? meta.total : meta.limit }}</b> of {{ meta.total}}
            </div>
            <div class="flex space-x-1">
            <button @click="prevPage" class="px-3 py-1 min-w-9 min-h-9 text-sm font-normal text-slate-500 bg-white border border-slate-200 rounded hover:bg-slate-50 hover:border-slate-400 transition duration-200 ease">
                Prev
            </button>
            <!-- First page -->
            <button
              v-if="totalPages > 5 && (Number(meta.page || 1) > 3)"
              @click="setPage(1)"
              class="px-3 py-1 min-w-9 min-h-9 text-sm font-normal text-slate-500 bg-white border border-slate-200 rounded hover:bg-slate-50 hover:border-slate-400 transition duration-200 ease"
            >
              1
            </button>
            <span
              v-if="totalPages > 5 && (Number(meta.page || 1) > 4)"
              class="px-3 py-1 min-w-9 min-h-9 text-sm font-normal text-slate-500 flex items-center"
            >…</span>

            <!-- Middle pages (max 5 total shown) -->
            <template v-for="n in (() => {
                const tp = Number(totalPages || 1);
                const cur = Number(meta.page || 1);
                const maxButtons = 5;
                if (tp <= maxButtons) return Array.from({ length: tp }, (_, i) => i + 1);
                const side = Math.floor(maxButtons / 2); // 2
                let start = Math.max(1, cur - side);
                let end = Math.min(tp, cur + side);
                if (start === 1) end = maxButtons;
                if (end === tp) start = tp - (maxButtons - 1);
                const arr = [];
                for (let i = start; i <= end; i++) arr.push(i);
                return arr;
              })()"
            >
              <button
                @click="setPage(n)"
                :class="[
                  'px-3 py-1 min-w-9 min-h-9 text-sm font-normal border rounded transition duration-200 ease',
                  (Number(meta.page || 1) === n)
                    ? 'text-white bg-slate-800 border-slate-800'
                    : 'text-slate-500 bg-white border-slate-200 hover:bg-slate-50 hover:border-slate-400'
                ]"
              >
                {{ n }}
              </button>
            </template>

            <!-- Trailing ellipsis + last page when needed -->
            <span
              v-if="totalPages > 5 && (Number(meta.page || 1) < (totalPages - 3))"
              class="px-3 py-1 min-w-9 min-h-9 text-sm font-normal text-slate-500 flex items-center"
            >…</span>
            <button
              v-if="totalPages > 5 && (Number(meta.page || 1) < (totalPages - 2))"
              @click="setPage(totalPages)"
              class="px-3 py-1 min-w-9 min-h-9 text-sm font-normal text-slate-500 bg-white border border-slate-200 rounded hover:bg-slate-50 hover:border-slate-400 transition duration-200 ease"
            >
              {{ totalPages }}
            </button>
            <button @click="nextPage" class="px-3 py-1 min-w-9 min-h-9 text-sm font-normal text-slate-500 bg-white border border-slate-200 rounded hover:bg-slate-50 hover:border-slate-400 transition duration-200 ease">
                Next
            </button>
            </div>
        </div>
        </div>
      </div>
    </main>

  
</template>

<style scoped></style>