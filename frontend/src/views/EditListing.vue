<template>
  <div class="max-w-2xl mx-auto py-12">
    <div class="bg-gray-800 p-8 rounded-2xl border border-gray-700 shadow-xl">
      <h1 class="text-2xl font-bold text-white mb-8">Edit Listing</h1>
      <div v-if="loading" class="text-center py-10">Loading...</div>
      <form v-else @submit.prevent="handleSubmit" class="space-y-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-400 mb-2">Item Title</label>
            <input v-model="form.title" type="text" required class="w-full bg-gray-900 border border-gray-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-blue-500" />
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
            <select v-model="form.condition" class="w-full bg-gray-900 border border-gray-700 rounded-lg px-4 py-3 text-white">
              <option value="new">New</option>
              <option value="used">Used</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-400 mb-2">Status</label>
            <select v-model="form.status" class="w-full bg-gray-900 border border-gray-700 rounded-lg px-4 py-3 text-white">
              <option value="active">Active</option>
              <option value="inactive">Inactive</option>
              <option value="sold">Sold</option>
            </select>
          </div>
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-400 mb-2">Image URL</label>
            <input v-model="form.image_url" type="url" class="w-full bg-gray-900 border border-gray-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-blue-500" />
          </div>
          <div class="md:col-span-2">
            <label class="block text-sm font-medium text-gray-400 mb-2">Description</label>
            <textarea v-model="form.description" rows="4" class="w-full bg-gray-900 border border-gray-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-blue-500"></textarea>
          </div>
        </div>

        <div class="pt-6 flex gap-4">
          <button type="button" @click="$router.back()" class="flex-1 bg-gray-700 text-white font-bold py-4 rounded-xl hover:bg-gray-600 transition-all">Cancel</button>
          <button type="submit" :disabled="submitting" class="flex-1 bg-blue-600 text-white font-bold py-4 rounded-xl hover:bg-blue-700 transition-all disabled:opacity-50">
            {{ submitting ? 'Saving...' : 'Save Changes' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { reactive, onMounted, ref } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import api from '../services/api';

const router = useRouter();
const route = useRoute();
const categories = ref([]);
const loading = ref(true);
const submitting = ref(false);

const form = reactive({
  title: '',
  description: '',
  price: 0,
  category_id: '',
  condition: 'used',
  image_url: '',
  status: 'active',
});

const fetchData = async () => {
    try {
        const [catRes, itemRes] = await Promise.all([
            api.get('/categories/'),
            api.get(`/items/${route.params.id}`)
        ]);
        categories.value = catRes.data;
        Object.assign(form, {
            title: itemRes.data.title,
            description: itemRes.data.description,
            price: itemRes.data.price,
            category_id: itemRes.data.category_id,
            condition: itemRes.data.condition,
            image_url: itemRes.data.image_url,
            status: itemRes.data.status,
        });
    } catch (err) {
        alert('Failed to load listing details');
        router.push('/');
    } finally {
        loading.value = false;
    }
};

const handleSubmit = async () => {
  submitting.value = true;
  try {
    await api.put(`/items/${route.params.id}`, form);
    router.push(`/listings/${route.params.id}`);
  } catch (err) {
    alert('Failed to update listing');
  } finally {
    submitting.value = false;
  }
};

onMounted(fetchData);
</script>
