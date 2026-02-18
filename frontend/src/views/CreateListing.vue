<template>
  <div class="max-w-2xl mx-auto py-12">
    <div class="bg-gray-800 p-8 rounded-2xl border border-gray-700 shadow-xl">
      <h1 class="text-2xl font-bold text-white mb-8">Post New Listing</h1>
      <form @submit.prevent="handleSubmit" class="space-y-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-400 mb-2">Item Title</label>
            <input v-model="form.title" type="text" required class="w-full bg-gray-900 border border-gray-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-blue-500" placeholder="e.g. iPhone 15 Pro Max" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-400 mb-2">Price ($)</label>
            <input v-model.number="form.price" type="number" step="0.01" required class="w-full bg-gray-900 border border-gray-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-blue-500" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-400 mb-2">Category</label>
            <select v-model="form.category_id" required class="w-full bg-gray-900 border border-gray-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-blue-500">
              <option value="">Select a category</option>
              <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-400 mb-2">Condition</label>
            <div class="flex gap-4">
              <label class="bg-gray-900 px-4 py-2 border rounded-lg cursor-pointer transition-all" :class="form.condition === 'new' ? 'border-blue-500 text-blue-500' : 'border-gray-700 text-gray-400'">
                <input type="radio" v-model="form.condition" value="new" class="hidden" /> New
              </label>
              <label class="bg-gray-900 px-4 py-2 border rounded-lg cursor-pointer transition-all" :class="form.condition === 'used' ? 'border-blue-500 text-blue-500' : 'border-gray-700 text-gray-400'">
                <input type="radio" v-model="form.condition" value="used" class="hidden" /> Used
              </label>
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-400 mb-2">Image URL</label>
            <input v-model="form.image_url" type="url" class="w-full bg-gray-900 border border-gray-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-blue-500" placeholder="https://..." />
          </div>
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-400 mb-2">Description</label>
            <textarea v-model="form.description" rows="4" class="w-full bg-gray-900 border border-gray-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-blue-500" placeholder="Tell us more about the item..."></textarea>
          </div>
        </div>

        <div class="pt-6">
          <button type="submit" :disabled="submitting" class="w-full bg-blue-600 text-white font-bold py-4 rounded-xl hover:bg-blue-700 transition-all shadow-lg hover:shadow-blue-500/20 disabled:opacity-50">
            {{ submitting ? 'Publishing...' : 'Publish Listing' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { reactive, onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import api from '../services/api';

const router = useRouter();
const categories = ref([]);
const submitting = ref(false);

const form = reactive({
  title: '',
  description: '',
  price: 0,
  category_id: '',
  condition: 'used',
  image_url: '',
});

const fetchCategories = async () => {
  const res = await api.get('/categories/');
  categories.value = res.data;
};

const handleSubmit = async () => {
  submitting.value = true;
  try {
    await api.post('/items/', form);
    router.push('/');
  } catch (err) {
    alert('Failed to post item');
  } finally {
    submitting.value = false;
  }
};

onMounted(fetchCategories);
</script>
