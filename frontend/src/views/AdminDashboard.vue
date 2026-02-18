<template>
  <div class="py-8">
    <div class="flex justify-between items-center mb-8">
      <h1 class="text-3xl font-bold text-white">Admin Dashboard</h1>
      <button @click="showCategoryModal = true" class="bg-blue-600 px-4 py-2 rounded-lg font-semibold hover:bg-blue-700">Add Category</button>
    </div>

    <!-- Tabs -->
    <div class="flex border-b border-gray-700 mb-8 mt-4">
      <button @click="tab = 'items'" :class="['px-6 py-3 font-semibold transition-all', tab === 'items' ? 'text-blue-500 border-b-2 border-blue-500' : 'text-gray-500 hover:text-white']">Listings</button>
      <button @click="tab = 'categories'" :class="['px-6 py-3 font-semibold transition-all', tab === 'categories' ? 'text-blue-500 border-b-2 border-blue-500' : 'text-gray-500 hover:text-white']">Categories</button>
    </div>

    <!-- Listings Table -->
    <div v-if="tab === 'items'" class="bg-gray-800 rounded-2xl border border-gray-700 overflow-hidden">
      <table class="w-full text-left">
        <thead class="bg-gray-900/50 text-gray-400 text-sm uppercase">
          <tr>
            <th class="px-6 py-4">Item</th>
            <th class="px-6 py-4">Seller</th>
            <th class="px-6 py-4">Category</th>
            <th class="px-6 py-4">Price</th>
            <th class="px-6 py-4">Status</th>
            <th class="px-6 py-4 text-right">Actions</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-700 text-white">
          <tr v-for="item in items" :key="item.id">
            <td class="px-6 py-4 font-medium">{{ item.title }}</td>
            <td class="px-6 py-4">{{ item.seller?.name }}</td>
            <td class="px-6 py-4">{{ item.category?.name }}</td>
            <td class="px-6 py-4 text-blue-400">${{ item.price }}</td>
            <td class="px-6 py-4">
              <span :class="['text-xs px-2 py-1 rounded', item.status === 'active' ? 'bg-green-900/30 text-green-500' : 'bg-red-900/30 text-red-500']">{{ item.status }}</span>
            </td>
            <td class="px-6 py-4 text-right space-x-2">
              <button @click="toggleStatus(item)" class="text-xs text-yellow-500 hover:underline">{{ item.status === 'active' ? 'Deactivate' : 'Activate' }}</button>
              <button @click="deleteItem(item.id)" class="text-xs text-red-500 hover:underline">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Categories Table -->
    <div v-if="tab === 'categories'" class="bg-gray-800 rounded-2xl border border-gray-700 overflow-hidden">
      <table class="w-full text-left">
        <thead class="bg-gray-900/50 text-gray-400 text-sm uppercase">
          <tr>
            <th class="px-6 py-4">ID</th>
            <th class="px-6 py-4">Name</th>
            <th class="px-6 py-4">Created By</th>
            <th class="px-6 py-4 text-right">Actions</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-700 text-white">
          <tr v-for="cat in categories" :key="cat.id">
            <td class="px-6 py-4">{{ cat.id }}</td>
            <td class="px-6 py-4 font-medium">{{ cat.name }}</td>
            <td class="px-6 py-4 text-gray-400 text-sm">User {{ cat.created_by_login_id }} ({{ cat.created_by_role_id }})</td>
            <td class="px-6 py-4 text-right space-x-4 text-sm">
              <button @click="editCategory(cat)" class="text-blue-500 hover:underline">Edit</button>
              <button @click="deleteCategory(cat.id)" class="text-red-500 hover:underline">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Category Modal -->
    <div v-if="showCategoryModal" class="fixed inset-0 bg-black/80 flex items-center justify-center p-4 z-50">
      <div class="bg-gray-800 w-full max-w-sm p-6 rounded-2xl border border-gray-700">
        <h2 class="text-xl font-bold text-white mb-6">{{ editingCategory ? 'Edit Category' : 'Add New Category' }}</h2>
        <input v-model="categoryForm.name" placeholder="Category name" class="w-full bg-gray-900 border border-gray-700 rounded-lg px-4 py-3 text-white focus:outline-none mb-6" />
        <div class="flex gap-4">
          <button @click="closeCategoryModal" class="flex-1 bg-gray-700 font-semibold py-2 rounded-lg">Cancel</button>
          <button @click="saveCategory" class="flex-1 bg-blue-600 font-semibold py-2 rounded-lg hover:bg-blue-700">Save</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue';
import api from '../services/api';

const tab = ref('items');
const items = ref([]);
const categories = ref([]);
const showCategoryModal = ref(false);
const editingCategory = ref(null);
const categoryForm = reactive({ name: '' });

const fetchData = async () => {
    const [itemsRes, catRes] = await Promise.all([
        api.get('/admin/items'),
        api.get('/categories/')
    ]);
    items.value = itemsRes.data.data;
    categories.value = catRes.data;
};

const toggleStatus = async (item) => {
    const newStatus = item.status === 'active' ? 'inactive' : 'active';
    await api.put(`/items/${item.id}`, { status: newStatus });
    fetchData();
};

const deleteItem = async (id) => {
    if (!confirm('Delete this listing?')) return;
    await api.delete(`/items/${id}`);
    fetchData();
};

const editCategory = (cat) => {
    editingCategory.value = cat;
    categoryForm.name = cat.name;
    showCategoryModal.value = true;
};

const closeCategoryModal = () => {
    showCategoryModal.value = false;
    editingCategory.value = null;
    categoryForm.name = '';
};

const saveCategory = async () => {
    if (editingCategory.value) {
        await api.put(`/admin/categories/${editingCategory.value.id}`, categoryForm);
    } else {
        await api.post('/admin/categories', categoryForm);
    }
    closeCategoryModal();
    fetchData();
};

const deleteCategory = async (id) => {
    if (!confirm('Deleting category will also delete all its listings. Proceed?')) return;
    await api.delete(`/admin/categories/${id}`);
    fetchData();
};

onMounted(fetchData);
</script>
