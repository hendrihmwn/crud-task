<script setup>
import { computed } from 'vue'

const props = defineProps([
    'title',
    'description',
    'status',
])

const emit = defineEmits(['update:title', 'update:description', 'update:status'])
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

// expose a `data` object whose properties are computed getters/setters
// so the template's v-model="data.title" / "data.description" / "data.status" works
const data = {
    title: props.title,
    description: props.title,
    status: props.status
}
</script>

<template>
        <div>
        <label for="title" class="block mb-2 text-sm text-slate-600">Title</label>
        <div class="mt-2">
            <input v-model="data.title" type="text" name="title" id="title" autocomplete="" class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow" />
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
</template>

<style scoped></style>