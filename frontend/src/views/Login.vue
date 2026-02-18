<template>
  <div class="min-h-[80vh] flex items-center justify-center px-6 py-12">
    <div class="bg-gray-900 w-full max-w-md p-10 rounded-3xl border border-white/5 shadow-2xl space-y-8 relative overflow-hidden">
      <!-- Decorative background blur -->
      <div class="absolute -top-24 -right-24 w-48 h-48 bg-blue-600/10 blur-3xl rounded-full"></div>
      
      <div class="text-center space-y-2 relative">
        <h2 class="text-3xl font-black text-white tracking-tight">Welcome Back</h2>
        <p class="text-gray-500 font-medium">Enter your credentials to access the market.</p>
      </div>

      <form @submit.prevent="handleLogin" class="space-y-6 relative">
        <div class="space-y-2">
          <label class="block text-sm font-bold text-gray-400 ml-1">Email Address</label>
          <input 
            v-model="email" 
            type="email" 
            required 
            placeholder="name@example.com"
            class="w-full bg-gray-950 border border-white/5 rounded-2xl px-6 py-4 text-white focus:outline-none focus:ring-2 focus:ring-blue-500/50 focus:border-blue-500 transition-all placeholder:text-gray-700" 
          />
        </div>
        
        <div class="space-y-2">
          <div class="flex justify-between items-center ml-1">
            <label class="block text-sm font-bold text-gray-400">Password</label>
            <a href="#" class="text-xs font-bold text-blue-500 hover:text-blue-400 transition-colors">Forgot?</a>
          </div>
          <input 
            v-model="password" 
            type="password" 
            required 
            placeholder="••••••••"
            class="w-full bg-gray-950 border border-white/5 rounded-2xl px-6 py-4 text-white focus:outline-none focus:ring-2 focus:ring-blue-500/50 focus:border-blue-500 transition-all placeholder:text-gray-700" 
          />
        </div>

        <button 
          type="submit" 
          :disabled="loading" 
          class="w-full bg-blue-600 text-white font-black py-4 rounded-2xl hover:bg-blue-500 transition-all shadow-lg shadow-blue-500/20 disabled:opacity-50 text-lg translate-y-0 active:translate-y-1"
        >
          <span v-if="!loading">Sign In</span>
          <span v-else class="flex items-center justify-center gap-2">
            <span class="w-5 h-5 border-2 border-white/20 border-t-white rounded-full animate-spin"></span>
            Signing in...
          </span>
        </button>
      </form>

      <div v-if="error" class="bg-red-500/10 border border-red-500/20 p-4 rounded-2xl text-red-500 text-sm text-center font-medium animate-shake">
        {{ error }}
      </div>

      <p class="text-center text-gray-500 font-medium">
        New here? <router-link to="/register" class="text-blue-500 font-bold hover:underline decoration-2 underline-offset-4 transition-all">Create an account</router-link>
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useAuthStore } from '../stores/auth';
import { useRouter } from 'vue-router';

const authStore = useAuthStore();
const router = useRouter();

const email = ref('');
const password = ref('');
const loading = ref(false);
const error = ref('');

const handleLogin = async () => {
  loading.value = true;
  error.value = '';
  try {
    await authStore.login(email.value, password.value);
    router.push('/');
  } catch (err) {
    error.value = err.response?.data?.error || 'Login failed. Please try again.';
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
@keyframes shake {
  0%, 100% { transform: translateX(0); }
  25% { transform: translateX(-4px); }
  75% { transform: translateX(4px); }
}
.animate-shake {
  animation: shake 0.2s cubic-bezier(.36,.07,.19,.97) both;
}
</style>
