<template>
  <div class="max-w-7xl mx-auto px-6 py-12">
    <!-- Hero / Search Section -->
    <div class="flex flex-col md:flex-row justify-between items-end gap-6 mb-12">
      <div class="space-y-2">
        <h1 class="text-4xl font-black tracking-tight text-white">
          Discover <span class="text-blue-500">Amazing</span> Items
        </h1>
        <p class="text-gray-500 font-medium">Browse thousands of products from your local community.</p>
      </div>
      
      <div class="flex flex-wrap gap-4 w-full md:w-auto">
        <div class="relative group flex-grow md:flex-grow-0">
          <input 
            v-model="filters.search" 
            placeholder="Search everything..." 
            class="w-full md:w-80 bg-gray-900 border border-white/5 rounded-2xl px-5 py-3.5 text-white placeholder:text-gray-600 focus:outline-none focus:ring-2 focus:ring-blue-500/50 focus:border-blue-500 transition-all shadow-xl" 
          />
          <span class="absolute right-4 top-1/2 -translate-y-1/2 text-gray-600 group-focus-within:text-blue-500 transition-colors">🔍</span>
        </div>
        
        <select 
          v-model="filters.category_id" 
          class="bg-gray-900 border border-white/5 rounded-2xl px-5 py-3.5 text-white focus:outline-none focus:ring-2 focus:ring-blue-500/50 shadow-xl cursor-pointer hover:bg-gray-800 transition-colors"
        >
          <option value="">All Categories</option>
          <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
        </select>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-32 space-y-4">
      <div class="w-12 h-12 border-4 border-blue-500/20 border-t-blue-500 rounded-full animate-spin"></div>
      <p class="text-gray-500 font-medium animate-pulse">Fetching latest listings...</p>
    </div>

    <!-- Empty State -->
    <div v-else-if="items.length === 0" class="text-center py-32 bg-gray-900/50 rounded-3xl border border-dashed border-white/10">
      <div class="text-6xl mb-4">📦</div>
      <h3 class="text-xl font-bold text-white mb-2">No items found</h3>
      <p class="text-gray-500 mb-8">Try adjusting your filters or be the first to post something!</p>
      <router-link to="/listings/create" class="px-8 py-3 bg-blue-600 text-white font-bold rounded-xl hover:bg-blue-500 transition-all shadow-lg shadow-blue-500/20">
        Create Listing
      </router-link>
    </div>

    <!-- Items Grid -->
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-8">
      <div 
        v-for="item in items" 
        :key="item.id" 
        class="group bg-gray-900 rounded-3xl overflow-hidden border border-white/5 hover:border-blue-500/50 transition-all cursor-pointer shadow-xl hover:shadow-blue-500/10 hover:-translate-y-2"
        @click="$router.push(`/listings/${item.id}`)"
      >
        <div class="aspect-[4/3] overflow-hidden relative">
          <img 
            :src="item.image_url || 'https://via.placeholder.com/400x300?text=No+Image'" 
            class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-500" 
          />
          <div class="absolute top-4 left-4">
            <span class="px-3 py-1 bg-black/60 backdrop-blur-md rounded-full text-[10px] font-bold uppercase tracking-wider text-white border border-white/10">
              {{ item.category?.name }}
            </span>
          </div>
          <div class="absolute inset-0 bg-gradient-to-t from-gray-950 via-transparent to-transparent opacity-60"></div>
        </div>
        
        <div class="p-6">
          <div class="flex justify-between items-start mb-3">
            <h3 class="text-lg font-bold text-white group-hover:text-blue-400 transition-colors line-clamp-1 flex-grow mr-4">
              {{ item.title }}
            </h3>
            <span class="text-xl font-black text-blue-500">${{ item.price }}</span>
          </div>
          
          <p class="text-gray-500 text-sm mb-6 line-clamp-2 leading-relaxed">
            {{ item.description }}
          </p>
          
          <div class="flex items-center justify-between pt-4 border-t border-white/5">
            <div class="flex items-center gap-2">
              <div class="w-8 h-8 rounded-full bg-gradient-to-br from-gray-700 to-gray-800 flex items-center justify-center text-xs font-bold text-gray-300">
                {{ item.seller?.name?.charAt(0) }}
              </div>
              <span class="text-xs font-semibold text-gray-400">{{ item.seller?.name || 'Anonymous' }}</span>
            </div>
            <button class="text-blue-500 opacity-0 group-hover:opacity-100 transition-all transform translate-x-4 group-hover:translate-x-0">
              Details →
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Pagination -->
    <div v-if="totalPages > 1" class="mt-20 flex justify-center items-center gap-4">
      <button 
        v-for="p in totalPages" 
        :key="p" 
        @click="filters.page = p" 
        :class="[
          'w-12 h-12 rounded-2xl font-bold transition-all shadow-lg', 
          filters.page === p 
            ? 'bg-blue-600 text-white shadow-blue-500/20' 
            : 'bg-gray-900 text-gray-500 hover:text-white hover:bg-gray-800'
        ]"
      >
        {{ p }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, reactive } from 'vue';
import api from '../services/api';

const items = ref([]);
const categories = ref([]);
const loading = ref(true);
const totalPages = ref(1);

const filters = reactive({
  category_id: '',
  search: '',
  page: 1,
});

const fetchItems = async () => {
  loading.value = true;
  try {
    const response = await api.get('/items/', { params: filters });
    items.value = response.data.data;
    totalPages.value = Math.ceil(response.data.total / response.data.page_size);
  } catch (error) {
    console.error('Failed to fetch items:', error);
  } finally {
    loading.value = false;
  }
};

const fetchCategories = async () => {
  try {
    const response = await api.get('/categories/');
    categories.value = response.data;
  } catch (error) {
    console.error('Failed to fetch categories:', error);
  }
};

onMounted(() => {
  fetchItems();
  fetchCategories();
});

watch(filters, () => {
  fetchItems();
});
</script>
