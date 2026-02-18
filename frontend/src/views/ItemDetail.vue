<template>
  <div class="max-w-7xl mx-auto px-6 py-12">
    <!-- Breadcrumbs / Back Link -->
    <router-link to="/" class="inline-flex items-center gap-2 text-gray-500 hover:text-white transition-colors mb-8 group">
      <span class="group-hover:-translate-x-1 transition-transform">←</span>
      <span class="font-medium">Back to discovery</span>
    </router-link>

    <div v-if="loading" class="flex flex-col items-center justify-center py-32 space-y-4">
      <div class="w-12 h-12 border-4 border-blue-500/20 border-t-blue-500 rounded-full animate-spin"></div>
    </div>

    <div v-else-if="!item" class="text-center py-32 bg-gray-900/50 rounded-3xl border border-white/10">
      <div class="text-6xl mb-4">🔍</div>
      <h3 class="text-xl font-bold text-white mb-2">Item not found</h3>
      <router-link to="/" class="text-blue-500 hover:underline">Return home</router-link>
    </div>

    <div v-else class="grid grid-cols-1 lg:grid-cols-12 gap-12">
      <!-- Left Column: Gallery -->
      <div class="lg:col-span-7 space-y-6">
        <div class="relative aspect-[4/3] rounded-3xl overflow-hidden border border-white/5 shadow-2xl group">
          <img 
            :src="item.image_url || 'https://via.placeholder.com/800x600?text=No+Image'" 
            class="w-full h-full object-cover" 
          />
          <div class="absolute inset-0 bg-gradient-to-t from-gray-950/20 to-transparent"></div>
        </div>
      </div>

      <!-- Right Column: Info & Actions -->
      <div class="lg:col-span-5 space-y-8">
        <div class="bg-gray-900 rounded-3xl p-8 border border-white/5 shadow-xl space-y-6">
          <div class="space-y-4">
            <div class="flex items-center gap-3">
              <span class="px-3 py-1 bg-blue-500/10 text-blue-500 rounded-full text-[10px] font-bold uppercase tracking-wider border border-blue-500/20">
                {{ item.category?.name }}
              </span>
              <span class="px-3 py-1 bg-gray-800 text-gray-400 rounded-full text-[10px] font-bold uppercase tracking-wider border border-white/5">
                {{ item.condition }}
              </span>
            </div>
            
            <h1 class="text-4xl font-black text-white leading-tight">{{ item.title }}</h1>
            
            <div class="flex items-baseline gap-2">
              <span class="text-4xl font-black text-blue-500">${{ item.price }}</span>
              <span class="text-gray-500 text-sm font-medium">USD</span>
            </div>
          </div>

          <p class="text-gray-400 leading-relaxed text-lg">
            {{ item.description }}
          </p>

          <div class="pt-8 border-t border-white/5 space-y-4">
            <template v-if="isOwner">
              <router-link 
                :to="`/listings/${item.id}/edit`" 
                class="block w-full text-center bg-gray-800 text-white font-bold py-4 rounded-2xl hover:bg-gray-700 transition-all border border-white/5"
              >
                Edit Listing
              </router-link>
              <button 
                @click="handleDelete" 
                class="w-full text-center bg-red-500/10 text-red-500 font-bold py-4 rounded-2xl hover:bg-red-500 hover:text-white transition-all border border-red-500/20"
              >
                Delete Listing
              </button>
            </template>
            <template v-else-if="authStore.isAuthenticated">
              <button 
                @click="showModal = true" 
                class="w-full bg-blue-600 text-white font-black py-4 rounded-2xl hover:bg-blue-500 transition-all shadow-lg shadow-blue-500/20 text-lg translate-y-0 active:translate-y-1"
              >
                Contact Seller
              </button>
            </template>
            <template v-else>
              <router-link 
                to="/login" 
                class="block w-full text-center bg-blue-600 text-white font-black py-4 rounded-2xl shadow-lg shadow-blue-500/20"
              >
                Login to Message Seller
              </router-link>
            </template>
          </div>
        </div>

        <!-- Seller Profile Card -->
        <div class="bg-gray-900 rounded-3xl p-6 border border-white/5 shadow-xl flex items-center gap-4">
          <div class="w-14 h-14 rounded-2xl bg-gradient-to-br from-blue-500 to-indigo-600 flex items-center justify-center text-xl font-bold text-white shadow-lg shadow-blue-500/20">
            {{ item.seller?.name?.charAt(0) }}
          </div>
          <div>
            <h3 class="text-white font-bold text-lg">{{ item.seller?.name }}</h3>
            <p class="text-gray-500 text-xs font-medium uppercase tracking-widest">
              Seller since {{ new Date(item.seller?.created_at).getFullYear() }}
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Message Modal -->
    <div v-if="showModal" class="fixed inset-0 bg-black/90 backdrop-blur-sm flex items-center justify-center p-6 z-[100]">
      <div class="bg-gray-900 w-full max-w-lg p-8 rounded-3xl border border-white/10 shadow-2xl space-y-6">
        <div class="space-y-2">
          <h2 class="text-2xl font-bold text-white">Message {{ item.seller?.name }}</h2>
          <p class="text-gray-500 text-sm">Inquiry about: <span class="text-white font-medium">{{ item.title }}</span></p>
        </div>
        
        <textarea 
          v-model="messageContent" 
          rows="5" 
          placeholder="I'm interested in this item! Is it still available?" 
          class="w-full bg-gray-950 border border-white/5 rounded-2xl px-6 py-4 text-white focus:outline-none focus:ring-2 focus:ring-blue-500/50 focus:border-blue-500 transition-all placeholder:text-gray-700"
        ></textarea>
        
        <div class="flex gap-4">
          <button @click="showModal = false" class="flex-1 bg-gray-800 text-white font-bold py-4 rounded-2xl hover:bg-gray-700 transition-all">Cancel</button>
          <button 
            @click="sendMessage" 
            :disabled="!messageContent || sending" 
            class="flex-1 bg-blue-600 text-white font-black py-4 rounded-2xl hover:bg-blue-500 transition-all disabled:opacity-50 shadow-lg shadow-blue-500/20"
          >
            {{ sending ? 'Sending...' : 'Send Message' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useAuthStore } from '../stores/auth';
import api from '../services/api';

const route = useRoute();
const router = useRouter();
const authStore = useAuthStore();

const item = ref(null);
const loading = ref(true);
const showModal = ref(false);
const messageContent = ref('');
const sending = ref(false);

const isOwner = computed(() => {
  return authStore.user && item.value && authStore.user.id === item.value.seller_id;
});

const fetchItem = async () => {
  try {
    const response = await api.get(`/items/${route.params.id}`);
    item.value = response.data;
  } catch (error) {
    console.error('Failed to fetch item:', error);
  } finally {
    loading.value = false;
  }
};

const handleDelete = async () => {
  if (!confirm('Are you sure you want to delete this listing?')) return;
  try {
    await api.delete(`/items/${item.value.id}`);
    router.push('/');
  } catch (error) {
    alert('Failed to delete item');
  }
};

const sendMessage = async () => {
  sending.value = true;
  try {
    await api.post('/messages', {
      receiver_id: item.value.seller_id,
      item_id: item.value.id,
      content: messageContent.value,
    });
    alert('Message sent!');
    showModal.value = false;
    messageContent.value = '';
  } catch (error) {
    alert('Failed to send message');
  } finally {
    sending.value = false;
  }
};

onMounted(fetchItem);
</script>
